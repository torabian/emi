import { Entity1Dto } from "./Entity1Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity1Create
 */
/**
 * Entity1CreateAction
 */
export class Entity1CreateAction {
  //
  static URL = "/entity1";
  static NewUrl = (qs) => buildUrl(Entity1CreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity1CreateAction.NewUrl(qs),
      {
        method: Entity1CreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity1Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity1Dto(item));
    const res = await Entity1CreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "entity1Create",
    cliShort: "entity1-c",
    url: "/entity1",
    method: "post",
    description: 'Creates a new "entity1" row.',
    in: {
      dto: "Entity1Dto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity1Dto",
    },
  };
}
