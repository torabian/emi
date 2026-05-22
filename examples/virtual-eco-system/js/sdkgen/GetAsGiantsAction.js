import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { GiantDto } from './GiantDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action getAsGiants
*/
	/**
 * GetAsGiantsAction
 */
export class GetAsGiantsAction { //
  static URL = '/get/giant/:id';
  static NewUrl = (
	params,
	qs
  ) => buildUrl(
		GetAsGiantsAction.URL,
		params,
		qs
	);
  static Method = 'get';
	static Fetch$ = async (
			params,
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
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
			params,
		init,
		{
			creatorFn,
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
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