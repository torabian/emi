import { URLSearchParamsX } from "./sdk/common/URLSearchParamsX";
import { WebSocketX } from "./sdk/common/WebSocketX";
import { buildUrl } from "./sdk/common/buildUrl";
/**
 * Action to communicate with the action streamMedia
 */
export type StreamMediaActionOptions = {
  queryKey?: unknown[];
  params: StreamMediaActionPathParameter;
  qs?: StreamMediaActionQueryParams;
};
/**
 * Path parameters for StreamMediaAction
 */
export type StreamMediaActionPathParameter = {
  deviceId: string;
};
/**
 * StreamMediaAction
 */
export class StreamMediaAction {
  //
  static URL = "/media/stream/:deviceId";
  static NewUrl = (
    params: StreamMediaActionPathParameter,
    qs?: StreamMediaActionQueryParams,
  ) => buildUrl(StreamMediaAction.URL, params, qs);
  static Method = "REACTIVE";
  static Create = (
    overrideUrl?: string,
    qs?: StreamMediaActionQueryParams,
    params: StreamMediaActionPathParameter,
    options,
  ) => {
    const url = overrideUrl ?? StreamMediaAction.NewUrl(params, qs);
    const Cls = options?.SocketClass
      ? options.SocketClass
      : WebSocketX<unknown, unknown>;
    const socket = new Cls(url, undefined, {
      MessageFactoryClass: undefined,
    });
    socket.binaryType = "arraybuffer"; // raw binary frames, see MediaStreamer
    return socket;
  };
  static Definition = {
    name: "streamMedia",
    url: "/media/stream/:deviceId",
    method: "reactive",
    binaryType: "arraybuffer",
    qs: [
      {
        name: "sampleRate",
        type: "int",
      },
      {
        name: "channels",
        type: "int",
      },
      {
        name: "audioCodec",
        type: "string",
      },
      {
        name: "videoCodec",
        type: "string",
      },
    ],
    description:
      "Device -\u003e server raw audio (PCM/Opus) and video (MJPEG/H264) frames.",
  };
}
/**
 * StreamMediaActionQueryParams class
 * Auto-generated from EmiAction
 */
export class StreamMediaActionQueryParams extends URLSearchParamsX {
  /**
   *
   * @returns { number | null }
   */
  getSampleRate() {
    return this.getTyped("sampleRate", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setSampleRate(value: number | null) {
    this.set("sampleRate", value);
    return this;
  }
  /**
   *
   * @returns { number | null }
   */
  getChannels() {
    return this.getTyped("channels", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setChannels(value: number | null) {
    this.set("channels", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getAudioCodec() {
    return this.getTyped("audioCodec", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setAudioCodec(value: string | null) {
    this.set("audioCodec", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getVideoCodec() {
    return this.getTyped("videoCodec", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setVideoCodec(value: string | null) {
    this.set("videoCodec", value);
    return this;
  }
}
