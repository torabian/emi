/**
 * One microphone / camera frame on a binary websocket (`method: reactive` +
 * `binaryType: arraybuffer`). Same 12 byte layout as emigo.MediaFrame (Go),
 * emi::EmiMediaFrameHeader (C++) and emimedia.h (C):
 *
 *   0 kind (1 audio, 2 video, 3 other) | 1 flags (bit0 keyframe) | 2-3 reserved
 *   4-7 seq (per kind, big endian)     | 8-11 timestampMs (big endian)  | 12.. payload
 */
export const MEDIA_AUDIO = 1;
export const MEDIA_VIDEO = 2;
export const MEDIA_OTHER = 3;
export const MEDIA_KEYFRAME = 0x01;
export const MEDIA_HEADER_SIZE = 12;

export interface MediaFrame {
  kind: number;
  flags: number;
  seq: number;
  timestampMs: number;
  payload: Uint8Array;
}

export function encodeMediaFrame(f: MediaFrame): Uint8Array<ArrayBuffer> {
  const out = new Uint8Array(MEDIA_HEADER_SIZE + f.payload.byteLength);
  const view = new DataView(out.buffer);
  out[0] = f.kind;
  out[1] = f.flags;
  view.setUint32(4, f.seq >>> 0);
  view.setUint32(8, f.timestampMs >>> 0);
  out.set(f.payload, MEDIA_HEADER_SIZE);
  return out;
}

/** Throws when the message is shorter than a header. `payload` aliases `data`. */
export function decodeMediaFrame(data: ArrayBuffer | ArrayBufferView): MediaFrame {
  const bytes = ArrayBuffer.isView(data)
    ? new Uint8Array(data.buffer, data.byteOffset, data.byteLength)
    : new Uint8Array(data);
  if (bytes.byteLength < MEDIA_HEADER_SIZE) {
    throw new Error("media frame shorter than its 12 byte header");
  }
  const view = new DataView(bytes.buffer, bytes.byteOffset, bytes.byteLength);
  return {
    kind: bytes[0],
    flags: bytes[1],
    seq: view.getUint32(4),
    timestampMs: view.getUint32(8),
    payload: bytes.subarray(MEDIA_HEADER_SIZE),
  };
}
