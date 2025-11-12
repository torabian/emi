import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { GiantDto } from './GiantDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action modifyGiantDto
*/
	/**
 * ModifyGiantDtoAction
 */
export class ModifyGiantDtoAction { //
  static URL = '/modify/dto';
  static NewUrl = (
	qs
  ) => buildUrl(
		ModifyGiantDtoAction.URL,
		 undefined,
		qs
	);
  static Method = 'post';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
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