import { Entity1Dto } from "./Entity1Dto";
import { Entity1OptionalDto } from "./Entity1OptionalDto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity1Update
 */
/**
 * Entity1UpdateAction
 */
export class Entity1UpdateAction {
  //
  static URL = "/entity1/:uniqueId";
  static NewUrl = (params, qs) => buildUrl(Entity1UpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity1UpdateAction.NewUrl(params, qs),
      {
        method: Entity1UpdateAction.Method,
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
    const res = await Entity1UpdateAction.Fetch$(
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
    name: "entity1Update",
    cliShort: "entity1-u",
    url: "/entity1/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "entity1" row by uniqueId.',
    in: {
      dto: "Entity1OptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity1Dto",
    },
  };
}
