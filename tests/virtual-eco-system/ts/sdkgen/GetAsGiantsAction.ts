import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { GiantDto } from './GiantDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action getAsGiants
*/
export type GetAsGiantsActionOptions = {
	queryKey?: unknown[];
	params: GetAsGiantsActionPathParameter;
	qs?: URLSearchParams;
};
	/**
 * Path parameters for GetAsGiantsAction
 */
export type GetAsGiantsActionPathParameter = {
	id: string | number | boolean;
}
	/**
 * GetAsGiantsAction
 */
export class GetAsGiantsAction { //
  static URL = '/get/giant/:id';
  static NewUrl = (
	params: GetAsGiantsActionPathParameter,
	qs?: URLSearchParams
  ) => buildUrl(
		GetAsGiantsAction.URL,
		params,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
			params: GetAsGiantsActionPathParameter,
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<GiantDto, unknown, unknown>(
			overrideUrl ?? GetAsGiantsAction.NewUrl(
				params,
				qs
			),
			{
				method: GetAsGiantsAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
			params: GetAsGiantsActionPathParameter,
		init?: TypedRequestInit<unknown, unknown>,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
				creatorFn?: ((item: unknown) => GiantDto) | undefined,
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
				creatorFn: (item) => new GiantDto(item),
		}
	) => {
		creatorFn = creatorFn || ((item) => new GiantDto(item))
		const res = await GetAsGiantsAction.Fetch$(
			params,
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				(item) => creatorFn(item),
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "getAsGiants",
  "url": "/get/giant/:id",
  "method": "get",
  "out": {
    "dto": "GiantDto"
  }
}
}