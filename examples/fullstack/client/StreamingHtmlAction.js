import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action streamingHtml
 */
/**
 * StreamingHtmlAction
 */
export class StreamingHtmlAction {
  //
  static URL = "/stream/html";
  static NewUrl = (qs) => buildUrl(StreamingHtmlAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? StreamingHtmlAction.NewUrl(qs),
      {
        method: StreamingHtmlAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (init, { qs, ctx, onMessage, overrideUrl } = {}) => {
    const res = await StreamingHtmlAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "streamingHtml",
    url: "/stream/html",
    method: "get",
    out: {
      primitive: "bytes",
    },
  };
}
