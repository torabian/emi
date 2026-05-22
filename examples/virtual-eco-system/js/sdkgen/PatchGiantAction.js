import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { GiantDto } from './GiantDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action patchGiant
*/
	/**
 * PatchGiantAction
 */
export class PatchGiantAction { //
  static URL = '/get/giant/:id';
  static NewUrl = (
	params,
	qs
  ) => buildUrl(
		PatchGiantAction.URL,
		params,
		qs
	);
  static Method = 'patch';
	static Fetch$ = async (
			params,
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? PatchGiantAction.NewUrl(
				params,
				qs
			),
			{
				method: PatchGiantAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
			params,
		init,
		{
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
		}
	) => {
		const res = await PatchGiantAction.Fetch$(
			params,
			qs,
			ctx,
			init,
			overrideUrl,
			);
			return handleFetchResponse(
				res, 
				undefined,
				onMessage,
				init?.signal,
			);
	}
  static Definition = {
  "name": "patchGiant",
  "url": "/get/giant/:id",
  "method": "patch",
  "in": {
    "dto": "GiantDto"
  }
}
}