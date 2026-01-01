import { FetchxContext, fetchx, handleFetchResponse } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action Delete a draft offer
*/
	/**
 * DeleteADraftOfferAction
 */
export class DeleteADraftOfferAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}';
  static NewUrl = (
	qs
  ) => buildUrl(
		DeleteADraftOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'delete';
	static Fetch$ = async (
		qs,
		ctx,
		init,
		overrideUrl,
	) => {
		return fetchx(
			overrideUrl ?? DeleteADraftOfferAction.NewUrl(
				qs
			),
			{
				method: DeleteADraftOfferAction.Method,
				...(init || {})
			},
			ctx
		)
	}
	static Fetch = async (
		init,
		{
			qs,
			ctx,
			onMessage,
			overrideUrl
		}  = {
		}
	) => {
		const res = await DeleteADraftOfferAction.Fetch$(
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
  "name": "Delete a draft offer",
  "url": "https://api.{environment}/sale/offers/{offerId}",
  "method": "delete",
  "description": "Use this resource to delete a draft offer. Read more: PL / EN."
}
}