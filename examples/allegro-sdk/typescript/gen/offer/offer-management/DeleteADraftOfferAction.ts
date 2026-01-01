import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action Delete a draft offer
*/
export type DeleteADraftOfferActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * DeleteADraftOfferAction
 */
export class DeleteADraftOfferAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		DeleteADraftOfferAction.URL,
		 undefined,
		qs
	);
  static Method = 'delete';
	static Fetch$ = async (
		qs?: URLSearchParams,
		ctx?: FetchxContext,
		init?: TypedRequestInit<unknown, unknown>,
		overrideUrl?: string,
	) => {
		return fetchx<unknown, unknown, unknown>(
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
		init?: TypedRequestInit<unknown, unknown>,
		{
			qs,
			ctx,
			onMessage,
			overrideUrl
		} 
			: {
			qs?: URLSearchParams,
			ctx?: FetchxContext,
			onMessage?: (ev: MessageEvent) => void,
			overrideUrl?: string,		
		} 
			 = {
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