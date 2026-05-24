import { GiantDto } from "./GiantDto";
import { buildUrl } from "./sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
} from "./sdk/common/fetchx";
/**
 * Action to communicate with the action patchGiant
 */
export type PatchGiantActionOptions = {
  queryKey?: unknown[];
  params: PatchGiantActionPathParameter;
  qs?: URLSearchParams;
};
/**
 * Path parameters for PatchGiantAction
 */
export type PatchGiantActionPathParameter = {
  id: string;
};
/**
 * PatchGiantAction
 */
export class PatchGiantAction {
  //
  static URL = "/get/giant/:id";
  static NewUrl = (
    params: PatchGiantActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PatchGiantAction.URL, params, qs);
  static Method = "patch";
  static Fetch$ = async (
    params: PatchGiantActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<GiantDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<unknown, GiantDto, unknown>(
      overrideUrl ?? PatchGiantAction.NewUrl(params, qs),
      {
        method: PatchGiantAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PatchGiantActionPathParameter,
    init?: TypedRequestInit<GiantDto, unknown>,
    {
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {},
  ) => {
    const res = await PatchGiantAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "patchGiant",
    url: "/get/giant/:id",
    method: "patch",
    in: {
      dto: "GiantDto",
    },
  };
}
