var __classPrivateFieldSet = (this && this.__classPrivateFieldSet) || function (receiver, state, value, kind, f) {
    if (kind === "m") throw new TypeError("Private method is not writable");
    if (kind === "a" && !f) throw new TypeError("Private accessor was defined without a setter");
    if (typeof state === "function" ? receiver !== state || !f : !state.has(receiver)) throw new TypeError("Cannot write private member to an object whose class did not declare it");
    return (kind === "a" ? f.call(receiver, value) : f ? f.value = value : state.set(receiver, value)), value;
};
var __classPrivateFieldGet = (this && this.__classPrivateFieldGet) || function (receiver, state, kind, f) {
    if (kind === "a" && !f) throw new TypeError("Private accessor was defined without a getter");
    if (typeof state === "function" ? receiver !== state || !f : !state.has(receiver)) throw new TypeError("Cannot read private member from an object whose class did not declare it");
    return kind === "m" ? f : kind === "a" ? f.call(receiver) : f ? f.value : state.get(receiver);
};
var _MediaStreamer_instances, _MediaStreamer_socket, _MediaStreamer_opts, _MediaStreamer_stream, _MediaStreamer_ownsStream, _MediaStreamer_ctx, _MediaStreamer_node, _MediaStreamer_timer, _MediaStreamer_startedAt, _MediaStreamer_audioSeq, _MediaStreamer_videoSeq, _MediaStreamer_pending, _MediaStreamer_now, _MediaStreamer_send, _MediaStreamer_startAudio, _MediaStreamer_onSamples, _MediaStreamer_startVideo;
import { encodeMediaFrame, MEDIA_AUDIO, MEDIA_KEYFRAME, MEDIA_VIDEO, } from "./MediaFrame";
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
    constructor(socket, options = {}) {
        var _a, _b, _c, _d, _e, _f, _g, _h, _j;
        _MediaStreamer_instances.add(this);
        this.stats = { audioFrames: 0, videoFrames: 0, skippedVideoFrames: 0 };
        _MediaStreamer_socket.set(this, void 0);
        _MediaStreamer_opts.set(this, void 0);
        _MediaStreamer_stream.set(this, void 0);
        _MediaStreamer_ownsStream.set(this, false);
        _MediaStreamer_ctx.set(this, void 0);
        _MediaStreamer_node.set(this, void 0);
        _MediaStreamer_timer.set(this, void 0);
        _MediaStreamer_startedAt.set(this, 0);
        _MediaStreamer_audioSeq.set(this, 0);
        _MediaStreamer_videoSeq.set(this, 0);
        _MediaStreamer_pending.set(this, new Float32Array(0));
        const sampleRate = (_a = options.sampleRate) !== null && _a !== void 0 ? _a : 16000;
        __classPrivateFieldSet(this, _MediaStreamer_socket, socket, "f");
        __classPrivateFieldSet(this, _MediaStreamer_opts, {
            audio: (_b = options.audio) !== null && _b !== void 0 ? _b : true,
            video: (_c = options.video) !== null && _c !== void 0 ? _c : true,
            sampleRate,
            samplesPerFrame: (_d = options.samplesPerFrame) !== null && _d !== void 0 ? _d : Math.round(sampleRate / 50),
            fps: (_e = options.fps) !== null && _e !== void 0 ? _e : 5,
            width: (_f = options.width) !== null && _f !== void 0 ? _f : 640,
            height: (_g = options.height) !== null && _g !== void 0 ? _g : 480,
            jpegQuality: (_h = options.jpegQuality) !== null && _h !== void 0 ? _h : 0.6,
            maxBufferedBytes: (_j = options.maxBufferedBytes) !== null && _j !== void 0 ? _j : 512 * 1024,
            stream: options.stream,
        }, "f");
        if (socket.binaryType !== "arraybuffer" && socket.binaryType !== "blob") {
            socket.binaryType = "arraybuffer";
        }
    }
    get running() {
        return __classPrivateFieldGet(this, _MediaStreamer_startedAt, "f") !== 0;
    }
    async start() {
        if (this.running)
            return;
        const o = __classPrivateFieldGet(this, _MediaStreamer_opts, "f");
        if (o.stream) {
            __classPrivateFieldSet(this, _MediaStreamer_stream, o.stream, "f");
        }
        else {
            __classPrivateFieldSet(this, _MediaStreamer_stream, await navigator.mediaDevices.getUserMedia({
                audio: o.audio ? { channelCount: 1, echoCancellation: true } : false,
                video: o.video ? { width: o.width, height: o.height } : false,
            }), "f");
            __classPrivateFieldSet(this, _MediaStreamer_ownsStream, true, "f");
        }
        __classPrivateFieldSet(this, _MediaStreamer_startedAt, performance.now(), "f");
        // Independent pipelines: a slow audio start must not delay the camera.
        const stream = __classPrivateFieldGet(this, _MediaStreamer_stream, "f");
        await Promise.all([
            o.audio && stream.getAudioTracks().length ? __classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_startAudio).call(this, stream) : undefined,
            o.video && stream.getVideoTracks().length ? __classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_startVideo).call(this, stream) : undefined,
        ]);
    }
    stop() {
        var _a, _b, _c;
        if (__classPrivateFieldGet(this, _MediaStreamer_timer, "f"))
            clearInterval(__classPrivateFieldGet(this, _MediaStreamer_timer, "f"));
        __classPrivateFieldSet(this, _MediaStreamer_timer, undefined, "f");
        (_a = __classPrivateFieldGet(this, _MediaStreamer_node, "f")) === null || _a === void 0 ? void 0 : _a.disconnect();
        __classPrivateFieldSet(this, _MediaStreamer_node, undefined, "f");
        void ((_b = __classPrivateFieldGet(this, _MediaStreamer_ctx, "f")) === null || _b === void 0 ? void 0 : _b.close());
        __classPrivateFieldSet(this, _MediaStreamer_ctx, undefined, "f");
        if (__classPrivateFieldGet(this, _MediaStreamer_ownsStream, "f"))
            (_c = __classPrivateFieldGet(this, _MediaStreamer_stream, "f")) === null || _c === void 0 ? void 0 : _c.getTracks().forEach((t) => t.stop());
        __classPrivateFieldSet(this, _MediaStreamer_stream, undefined, "f");
        __classPrivateFieldSet(this, _MediaStreamer_startedAt, 0, "f");
    }
}
_MediaStreamer_socket = new WeakMap(), _MediaStreamer_opts = new WeakMap(), _MediaStreamer_stream = new WeakMap(), _MediaStreamer_ownsStream = new WeakMap(), _MediaStreamer_ctx = new WeakMap(), _MediaStreamer_node = new WeakMap(), _MediaStreamer_timer = new WeakMap(), _MediaStreamer_startedAt = new WeakMap(), _MediaStreamer_audioSeq = new WeakMap(), _MediaStreamer_videoSeq = new WeakMap(), _MediaStreamer_pending = new WeakMap(), _MediaStreamer_instances = new WeakSet(), _MediaStreamer_now = function _MediaStreamer_now() {
    return Math.round(performance.now() - __classPrivateFieldGet(this, _MediaStreamer_startedAt, "f"));
}, _MediaStreamer_send = function _MediaStreamer_send(kind, flags, seq, payload) {
    if (__classPrivateFieldGet(this, _MediaStreamer_socket, "f").readyState !== WebSocket.OPEN)
        return false;
    __classPrivateFieldGet(this, _MediaStreamer_socket, "f").send(encodeMediaFrame({ kind, flags, seq, timestampMs: __classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_now).call(this), payload }));
    return true;
}, _MediaStreamer_startAudio = async function _MediaStreamer_startAudio(stream) {
    const o = __classPrivateFieldGet(this, _MediaStreamer_opts, "f");
    // Asking for the target rate lets the browser resample for us.
    __classPrivateFieldSet(this, _MediaStreamer_ctx, new AudioContext({ sampleRate: o.sampleRate }), "f");
    const url = URL.createObjectURL(new Blob([WORKLET_SOURCE], { type: "text/javascript" }));
    try {
        await __classPrivateFieldGet(this, _MediaStreamer_ctx, "f").audioWorklet.addModule(url);
    }
    finally {
        URL.revokeObjectURL(url);
    }
    // Browsers create a context suspended unless start() ran inside a user gesture.
    if (__classPrivateFieldGet(this, _MediaStreamer_ctx, "f").state === "suspended")
        await __classPrivateFieldGet(this, _MediaStreamer_ctx, "f").resume();
    const source = __classPrivateFieldGet(this, _MediaStreamer_ctx, "f").createMediaStreamSource(stream);
    __classPrivateFieldSet(this, _MediaStreamer_node, new AudioWorkletNode(__classPrivateFieldGet(this, _MediaStreamer_ctx, "f"), "emi-pcm-tap", { numberOfOutputs: 0 }), "f");
    __classPrivateFieldGet(this, _MediaStreamer_node, "f").port.onmessage = (e) => __classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_onSamples).call(this, e.data);
    source.connect(__classPrivateFieldGet(this, _MediaStreamer_node, "f"));
}, _MediaStreamer_onSamples = function _MediaStreamer_onSamples(block) {
    var _a, _b;
    const n = __classPrivateFieldGet(this, _MediaStreamer_opts, "f").samplesPerFrame;
    const merged = new Float32Array(__classPrivateFieldGet(this, _MediaStreamer_pending, "f").length + block.length);
    merged.set(__classPrivateFieldGet(this, _MediaStreamer_pending, "f"));
    merged.set(block, __classPrivateFieldGet(this, _MediaStreamer_pending, "f").length);
    let offset = 0;
    while (merged.length - offset >= n) {
        const pcm = new Int16Array(n); // platform order: little endian on every browser target
        for (let i = 0; i < n; i++) {
            const s = Math.max(-1, Math.min(1, merged[offset + i]));
            pcm[i] = s < 0 ? s * 0x8000 : s * 0x7fff;
        }
        if (__classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_send).call(this, MEDIA_AUDIO, 0, (__classPrivateFieldSet(this, _MediaStreamer_audioSeq, (_b = __classPrivateFieldGet(this, _MediaStreamer_audioSeq, "f"), _a = _b++, _b), "f"), _a), new Uint8Array(pcm.buffer))) {
            this.stats.audioFrames++;
        }
        offset += n;
    }
    __classPrivateFieldSet(this, _MediaStreamer_pending, merged.slice(offset), "f");
}, _MediaStreamer_startVideo = async function _MediaStreamer_startVideo(stream) {
    const o = __classPrivateFieldGet(this, _MediaStreamer_opts, "f");
    const video = document.createElement("video");
    video.muted = true;
    video.playsInline = true;
    video.srcObject = new MediaStream(stream.getVideoTracks());
    await video.play();
    const canvas = document.createElement("canvas");
    const g = canvas.getContext("2d");
    if (!g)
        throw new Error("2d canvas is not available");
    let busy = false;
    __classPrivateFieldSet(this, _MediaStreamer_timer, setInterval(() => {
        if (busy || !video.videoWidth)
            return;
        if (__classPrivateFieldGet(this, _MediaStreamer_socket, "f").bufferedAmount > o.maxBufferedBytes) {
            this.stats.skippedVideoFrames++; // slow link: drop the frame instead of queueing latency
            return;
        }
        canvas.width = video.videoWidth;
        canvas.height = video.videoHeight;
        g.drawImage(video, 0, 0);
        busy = true;
        canvas.toBlob(async (blob) => {
            var _a, _b;
            busy = false;
            if (!blob)
                return;
            const jpeg = new Uint8Array(await blob.arrayBuffer());
            if (__classPrivateFieldGet(this, _MediaStreamer_instances, "m", _MediaStreamer_send).call(this, MEDIA_VIDEO, MEDIA_KEYFRAME, (__classPrivateFieldSet(this, _MediaStreamer_videoSeq, (_b = __classPrivateFieldGet(this, _MediaStreamer_videoSeq, "f"), _a = _b++, _b), "f"), _a), jpeg)) {
                this.stats.videoFrames++;
            }
        }, "image/jpeg", o.jpegQuality);
    }, 1000 / o.fps), "f");
};
