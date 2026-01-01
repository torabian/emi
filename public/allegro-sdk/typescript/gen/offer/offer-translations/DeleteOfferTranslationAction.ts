import { FetchxContext, fetchx, handleFetchResponse, type TypedRequestInit } from './sdk/common/fetchx';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action Delete offer translation
*/
export type DeleteOfferTranslationActionOptions = {
	queryKey?: unknown[];
	qs?: URLSearchParams;
};
	/**
 * DeleteOfferTranslationAction
 */
export class DeleteOfferTranslationAction { //
  static URL = 'https://api.{environment}/sale/offers/{offerId}/translations/{language}';
  static NewUrl = (
	qs?: URLSearchParams
  ) => buildUrl(
		DeleteOfferTranslationAction.URL,
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
			overrideUrl ?? DeleteOfferTranslationAction.NewUrl(
				qs
			),
			{
				method: DeleteOfferTranslationAction.Method,
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
		const res = await DeleteOfferTranslationAction.Fetch$(
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
  "name": "Delete offer translation",
  "url": "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
  "method": "delete",
  "description": "Delete single element or entire manual translation. Read more: PL / EN."
}
}