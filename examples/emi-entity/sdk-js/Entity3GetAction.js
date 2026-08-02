import { Entity3Dto } from "./Entity3Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity3Get
 */
/**
 * Entity3GetAction
 */
export class Entity3GetAction {
  //
  static URL = "/entity3/:uniqueId";
  static NewUrl = (params, qs) => buildUrl(Entity3GetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity3GetAction.NewUrl(params, qs),
      {
        method: Entity3GetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params,
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity3Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity3Dto(item));
    const res = await Entity3GetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "entity3Get",
    cliShort: "entity3-g",
    url: "/entity3/:uniqueId string",
    method: "get",
    description: 'Looks up a single "entity3" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "Entity3Dto",
    },
  };
}
