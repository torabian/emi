import { WebSocketX } from "./sdk/common/WebSocketX.js";
import { YokeMessageDto } from "./YokeMessageDto.js";
import { buildUrl } from "./sdk/common/buildUrl.js";

export class SubscribeYokeAction {
  static URL = "ws://localhost:8080/subscribe-yoke";
  static NewUrl = (qs) => buildUrl(SubscribeYokeAction.URL, undefined, qs);
  static Method = "reactive";

  /**
   * Creates a WebSocket connection for Yoke updates.
   * @param {string} [overrideUrl] Optional override for the default URL.
   * @param {object} [qs] Optional query parameters.
   * @returns {WebSocketX<YokeMessageDto>} A WebSocketX instance configured for YokeMessageDto messages.
   */
  static Create = (overrideUrl, qs) => {
    const url = overrideUrl ?? SubscribeYokeAction.NewUrl(qs);
    return new WebSocketX(url, undefined, {
      MessageFactoryClass: YokeMessageDto,
    });
  };

  static Definition = {
    name: "subscribeYoke",
    url: "/subscribe-yoke",
    method: "reactive",
    description:
      "Used by the Unreal Engine to fetch yoke information and move elements.",
    out: {
      dto: "YokeMessageDto",
    },
  };
}
