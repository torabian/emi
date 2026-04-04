import { buildUrl } from "./sdk/common/buildUrl";
import { fetchx, handleFetchResponse } from "./sdk/common/fetchx";
/**
 * Action to communicate with the action Delete offer translation
 */
/**
 * DeleteOfferTranslationAction
 */
export class DeleteOfferTranslationAction {
  //
  static URL =
    "https://api.{environment}/sale/offers/{offerId}/translations/{language}";
  static NewUrl = (qs) =>
    buildUrl(DeleteOfferTranslationAction.URL, undefined, qs);
  static Method = "delete";
  static Fetch$ = async (qs, ctx, init, overrideUrl) => {
    return fetchx(
      overrideUrl ?? DeleteOfferTranslationAction.NewUrl(qs),
      {
        method: DeleteOfferTranslationAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (init, { qs, ctx, onMessage, overrideUrl } = {}) => {
    const res = await DeleteOfferTranslationAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "Delete offer translation",
    url: "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
    method: "delete",
    description:
      "Delete single element or entire manual translation. Read more: PL / EN.",
  };
}
