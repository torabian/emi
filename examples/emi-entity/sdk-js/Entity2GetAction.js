import { Entity2Dto } from "./Entity2Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity2Get
 */
/**
 * Entity2GetAction
 */
export class Entity2GetAction {
  //
  static URL = "/entity2/:uniqueId";
  static NewUrl = (params, qs) => buildUrl(Entity2GetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2GetAction.NewUrl(params, qs),
      {
        method: Entity2GetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params,
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity2Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity2Dto(item));
    const res = await Entity2GetAction.Fetch$(
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
    name: "entity2Get",
    cliShort: "entity2-g",
    url: "/entity2/:uniqueId string",
    method: "get",
    description: 'Looks up a single "entity2" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "Entity2Dto",
    },
  };
}
