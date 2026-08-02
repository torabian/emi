import { Entity2Dto } from "./Entity2Dto";
import { Entity2OptionalDto } from "./Entity2OptionalDto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity2Update
 */
/**
 * Entity2UpdateAction
 */
export class Entity2UpdateAction {
  //
  static URL = "/entity2/:uniqueId string";
  static NewUrl = (params, qs) => buildUrl(Entity2UpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2UpdateAction.NewUrl(params, qs),
      {
        method: Entity2UpdateAction.Method,
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
    const res = await Entity2UpdateAction.Fetch$(
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
    name: "entity2Update",
    url: "/entity2/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "entity2" row by uniqueId.',
    in: {
      dto: "Entity2OptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity2Dto",
    },
  };
}
