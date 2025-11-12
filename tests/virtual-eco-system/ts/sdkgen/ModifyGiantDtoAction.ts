import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { GiantDto } from './GiantDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action modifyGiantDto
*/
export type ModifyGiantDtoActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * ModifyGiantDtoAction
 */
export class ModifyGiantDtoAction { //
  static URL = '/modify/dto';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		ModifyGiantDtoAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<GiantDto, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<GiantDto, GiantDto, unknown>(
			overrideUrl ?? ModifyGiantDtoAction.NewUrl(
				qs
			),
			{
				method: ModifyGiantDtoAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init?: TypedRequestInit<GiantDto, unknown>,
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
		const res = await ModifyGiantDtoAction.Fetch$(
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
  "name": "modifyGiantDto",
  "url": "/modify/dto",
  "method": "post",
  "in": {
    "dto": "GiantDto"
  },
  "out": {
    "dto": "GiantDto"
  }
}
}