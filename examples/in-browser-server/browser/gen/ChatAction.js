import { WebSocketX } from "./sdk/common/WebSocketX.js";
import { buildUrl } from "./sdk/common/buildUrl.js";
/**
 * Action to communicate with the action chat
 */
/**
 * ChatAction
 */
export class ChatAction {
  //
  static URL = "/chat";
  static NewUrl = (qs) => buildUrl(ChatAction.URL, undefined, qs);
  static Method = "reactive";
  static Create = (overrideUrl, qs, options) => {
    const url = overrideUrl ?? ChatAction.NewUrl(qs);
    const Cls = options?.SocketClass ? options.SocketClass : WebSocketX;
    return new Cls(url, undefined, {
      MessageFactoryClass: undefined,
    });
  };
  static Definition = {
    name: "chat",
    url: "/chat",
    method: "reactive",
    description:
      "A Reactive chating application, that returns length of strings you've typed.",
  };
}
