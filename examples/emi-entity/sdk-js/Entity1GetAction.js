import { Entity1Dto } from "./Entity1Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity1Get
 */
/**
 * Entity1GetAction
 */
export class Entity1GetAction {
  //
  static URL = "/entity1/:uniqueId";
  static NewUrl = (params, qs) => buildUrl(Entity1GetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity1GetAction.NewUrl(params, qs),
      {
        method: Entity1GetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params,
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity1Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity1Dto(item));
    const res = await Entity1GetAction.Fetch$(
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
    name: "entity1Get",
    cliShort: "entity1-g",
    url: "/entity1/:uniqueId string",
    method: "get",
    description: 'Looks up a single "entity1" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "Entity1Dto",
    },
  };
}
