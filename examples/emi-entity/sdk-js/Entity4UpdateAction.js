import { Entity4Dto } from "./Entity4Dto";
import { Entity4OptionalDto } from "./Entity4OptionalDto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity4Update
 */
/**
 * Entity4UpdateAction
 */
export class Entity4UpdateAction {
  //
  static URL = "/entity4/:uniqueId string";
  static NewUrl = (params, qs) => buildUrl(Entity4UpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (params, qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity4UpdateAction.NewUrl(params, qs),
      {
        method: Entity4UpdateAction.Method,
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
    const res = await Entity4UpdateAction.Fetch$(
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
    name: "entity4Update",
    cliShort: "entity4-u",
    url: "/entity4/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "entity4" row by uniqueId.',
    in: {
      dto: "Entity4OptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity4Dto",
    },
  };
}
