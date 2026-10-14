import {
  encodeMediaFrame,
  MEDIA_AUDIO,
  MEDIA_KEYFRAME,
  MEDIA_VIDEO,
} from "./MediaFrame";

export interface MediaStreamerOptions {
  /** Capture the microphone. Default true. */
  audio?: boolean;
  /** Capture the camera. Default true. */
  video?: boolean;
  /** Audio is sent as mono PCM16 little endian at this rate. Default 16000. */
  sampleRate?: number;
  /** Samples per audio frame (20 ms at 16 kHz = 320). Default sampleRate / 50. */
  samplesPerFrame?: number;
  /** Camera frames per second, sent as JPEG. Default 5. */
  fps?: number;
  /** Camera size hint. Default 640x480. */
  width?: number;
  height?: number;
  /** JPEG quality 0..1. Default 0.6. */
  jpegQuality?: number;
  /** A camera frame is skipped while the socket has more than this queued. Default 512 KiB. */
  maxBufferedBytes?: number;
  /** Use an existing stream instead of calling getUserMedia (tests, screen share). */
  stream?: MediaStream;
}

// Runs on the audio thread: forwards mono Float32 blocks to the main thread.
const WORKLET_SOURCE = `
class EmiPcmTap extends AudioWorkletProcessor {
  process(inputs) {
    const ch = inputs[0] && inputs[0][0];
    if (ch && ch.length) this.port.postMessage(ch.slice(0));
    return true;
  }
}
registerProcessor("emi-pcm-tap", EmiPcmTap);
`;

/**
 * Sends the browser's microphone (PCM16) and camera (JPEG) over a binary
 * websocket as media frames (see MediaFrame.ts). Works with the socket returned
 * by a generated `method: reactive` + `binaryType: arraybuffer` action's
 * `Create(...)`, or any WebSocket.
 *
 *   const socket = StreamMedia.Create(undefined, qs, { deviceId: "browser-1" });
 *   const streamer = new MediaStreamer(socket, { sampleRate: 16000 });
 *   socket.addEventListener("open", () => streamer.start());
 *   // ...later
 *   streamer.stop();
 *
 * Pass the same sampleRate in the action's query so the server knows the format.
 */
export class MediaStreamer {
  stats = { audioFrames: 0, videoFrames: 0, skippedVideoFrames: 0 };

  #socket: WebSocket;
  #opts: Required<Omit<MediaStreamerOptions, "stream">> & { stream?: MediaStream };
  #stream?: MediaStream;
  #ownsStream = false;
  #ctx?: AudioContext;
  #node?: AudioWorkletNode;
  #timer?: ReturnType<typeof setInterval>;
  #startedAt = 0;
  #audioSeq = 0;
  #videoSeq = 0;
  #pending = new Float32Array(0);

  constructor(socket: WebSocket, options: MediaStreamerOptions = {}) {
    const sampleRate = options.sampleRate ?? 16000;
    this.#socket = socket;
    this.#opts = {
      audio: options.audio ?? true,
      video: options.video ?? true,
      sampleRate,
      samplesPerFrame: options.samplesPerFrame ?? Math.round(sampleRate / 50),
      fps: options.fps ?? 5,
      width: options.width ?? 640,
      height: options.height ?? 480,
      jpegQuality: options.jpegQuality ?? 0.6,
      maxBufferedBytes: options.maxBufferedBytes ?? 512 * 1024,
      stream: options.stream,
    };
    if (socket.binaryType !== "arraybuffer" && socket.binaryType !== "blob") {
      socket.binaryType = "arraybuffer";
    }
  }

  get running() {
    return this.#startedAt !== 0;
  }

  async start(): Promise<void> {
    if (this.running) return;
    const o = this.#opts;

    if (o.stream) {
      this.#stream = o.stream;
    } else {
      this.#stream = await navigator.mediaDevices.getUserMedia({
        audio: o.audio ? { channelCount: 1, echoCancellation: true } : false,
        video: o.video ? { width: o.width, height: o.height } : false,
      });
      this.#ownsStream = true;
    }
    this.#startedAt = performance.now();

    // Independent pipelines: a slow audio start must not delay the camera.
    const stream = this.#stream;
    await Promise.all([
      o.audio && stream.getAudioTracks().length ? this.#startAudio(stream) : undefined,
      o.video && stream.getVideoTracks().length ? this.#startVideo(stream) : undefined,
    ]);
  }

  stop(): void {
    if (this.#timer) clearInterval(this.#timer);
    this.#timer = undefined;
    this.#node?.disconnect();
    this.#node = undefined;
    void this.#ctx?.close();
    this.#ctx = undefined;
    if (this.#ownsStream) this.#stream?.getTracks().forEach((t) => t.stop());
    this.#stream = undefined;
    this.#startedAt = 0;
  }

  #now() {
    return Math.round(performance.now() - this.#startedAt);
  }

  #send(kind: number, flags: number, seq: number, payload: Uint8Array) {
    if (this.#socket.readyState !== WebSocket.OPEN) return false;
    this.#socket.send(
      encodeMediaFrame({ kind, flags, seq, timestampMs: this.#now(), payload })
    );
    return true;
  }

  async #startAudio(stream: MediaStream) {
    const o = this.#opts;
    // Asking for the target rate lets the browser resample for us.
    this.#ctx = new AudioContext({ sampleRate: o.sampleRate });
    const url = URL.createObjectURL(new Blob([WORKLET_SOURCE], { type: "text/javascript" }));
    try {
      await this.#ctx.audioWorklet.addModule(url);
    } finally {
      URL.revokeObjectURL(url);
    }
    // Browsers create a context suspended unless start() ran inside a user gesture.
    if (this.#ctx.state === "suspended") await this.#ctx.resume();
    const source = this.#ctx.createMediaStreamSource(stream);
    this.#node = new AudioWorkletNode(this.#ctx, "emi-pcm-tap", { numberOfOutputs: 0 });
    this.#node.port.onmessage = (e: MessageEvent<Float32Array>) => this.#onSamples(e.data);
    source.connect(this.#node);
  }

  #onSamples(block: Float32Array) {
    const n = this.#opts.samplesPerFrame;
    const merged = new Float32Array(this.#pending.length + block.length);
    merged.set(this.#pending);
    merged.set(block, this.#pending.length);
    let offset = 0;
    while (merged.length - offset >= n) {
      const pcm = new Int16Array(n); // platform order: little endian on every browser target
      for (let i = 0; i < n; i++) {
        const s = Math.max(-1, Math.min(1, merged[offset + i]));
        pcm[i] = s < 0 ? s * 0x8000 : s * 0x7fff;
      }
      if (this.#send(MEDIA_AUDIO, 0, this.#audioSeq++, new Uint8Array(pcm.buffer))) {
        this.stats.audioFrames++;
      }
      offset += n;
    }
    this.#pending = merged.slice(offset);
  }

  async #startVideo(stream: MediaStream) {
    const o = this.#opts;
    const video = document.createElement("video");
    video.muted = true;
    video.playsInline = true;
    video.srcObject = new MediaStream(stream.getVideoTracks());
    await video.play();

    const canvas = document.createElement("canvas");
    const g = canvas.getContext("2d");
    if (!g) throw new Error("2d canvas is not available");
    let busy = false;

    this.#timer = setInterval(() => {
      if (busy || !video.videoWidth) return;
      if (this.#socket.bufferedAmount > o.maxBufferedBytes) {
        this.stats.skippedVideoFrames++; // slow link: drop the frame instead of queueing latency
        return;
      }
      canvas.width = video.videoWidth;
      canvas.height = video.videoHeight;
      g.drawImage(video, 0, 0);
      busy = true;
      canvas.toBlob(
        async (blob) => {
          busy = false;
          if (!blob) return;
          const jpeg = new Uint8Array(await blob.arrayBuffer());
          if (this.#send(MEDIA_VIDEO, MEDIA_KEYFRAME, this.#videoSeq++, jpeg)) {
            this.stats.videoFrames++;
          }
        },
        "image/jpeg",
        o.jpegQuality
      );
    }, 1000 / o.fps);
  }
}
