import { Entity4Dto } from "./Entity4Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity4Get
 */
/**
 * Entity4GetAction
 */
export class Entity4GetAction {
  //
  static URL = "/entity4/:id string";
  static NewUrl = (params, qs) => buildUrl(Entity4GetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity4GetAction.NewUrl(params, qs),
      {
        method: Entity4GetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params,
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity4Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity4Dto(item));
    const res = await Entity4GetAction.Fetch$(
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
    name: "entity4Get",
    url: "/entity4/:id string",
    method: "get",
    description: 'Looks up a single "entity4" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "Entity4Dto",
    },
  };
}
