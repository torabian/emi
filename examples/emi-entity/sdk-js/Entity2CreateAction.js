import { Entity2Dto } from "./Entity2Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity2Create
 */
/**
 * Entity2CreateAction
 */
export class Entity2CreateAction {
  //
  static URL = "/entity2";
  static NewUrl = (qs) => buildUrl(Entity2CreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity2CreateAction.NewUrl(qs),
      {
        method: Entity2CreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity2Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity2Dto(item));
    const res = await Entity2CreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "entity2Create",
    url: "/entity2",
    method: "post",
    description: 'Creates a new "entity2" row.',
    in: {
      dto: "Entity2Dto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity2Dto",
    },
  };
}
