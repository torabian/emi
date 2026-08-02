import { Entity3Dto } from "./Entity3Dto";
import { GResponse } from "./sdk/envelopes/index";
import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action entity3Create
 */
/**
 * Entity3CreateAction
 */
export class Entity3CreateAction {
  //
  static URL = "/entity3";
  static NewUrl = (qs) => buildUrl(Entity3CreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? Entity3CreateAction.NewUrl(qs),
      {
        method: Entity3CreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init,
    { creatorFn, qs, ctx, onMessage, overrideUrl } = {
      creatorFn: (item) => new Entity3Dto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new Entity3Dto(item));
    const res = await Entity3CreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "entity3Create",
    cliShort: "entity3-c",
    url: "/entity3",
    method: "post",
    description: 'Creates a new "entity3" row.',
    in: {
      dto: "Entity3Dto",
    },
    out: {
      envelope: "GResponse",
      dto: "Entity3Dto",
    },
  };
}
