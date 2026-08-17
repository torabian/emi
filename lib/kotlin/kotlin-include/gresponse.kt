package emikot

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.JsonElement

/**
 * GResponse - the Google JSON Style Guide envelope every entity Create/Update/Get
 * action requests by default (see EmiAction.GetResponseEnvelopeClass /
 * lib/core/preprocess-entity-actions.go's entityActionResponseEnvelope), and any
 * hand-declared action can opt into via `out: { envelope: GResponse }`.
 *
 * Field set mirrors lib/js/ts-envelopes/google-json-style-guide/google-envelop.emi.yml
 * exactly - the canonical spec already shipped in this repo for the JS/TS compiler's
 * own GResponse (see lib/js/ts-envelopes/google-json-style-guide/index.ts) - so a
 * server/client pair can mix JS and Kotlin consumers of the same wire format.
 *
 * `data.item`/`data.items` are plain nullable `T?`/`List<T>?` rather than
 * `MaybeField<T>` - MaybeFieldSerializer (see common.kt) is written for
 * `MaybeField<Any?>` specifically, not proven generic-safe for an arbitrary `T` without
 * extra contextual-serializer wiring, and a plain nullable is all `data.item` vs
 * `data.items` (object vs array response) actually needs.
 */
@Serializable
data class GResponse<T>(
    @SerialName("apiVersion") val apiVersion: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("context") val context: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("id") val id: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("method") val method: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("params") val params: JsonElement? = null,
    @SerialName("data") val data: GResponseData<T>? = null,
    @SerialName("error") val error: GResponseError? = null,
)

@Serializable
data class GResponseData<T>(
    @SerialName("item") val item: T? = null,
    @SerialName("items") val items: List<T>? = null,
    @SerialName("editLink") val editLink: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("selfLink") val selfLink: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("kind") val kind: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("fields") val fields: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("etag") val etag: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("cursor") val cursor: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("id") val id: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("lang") val lang: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("updated") val updated: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("currentItemCount") val currentItemCount: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("itemsPerPage") val itemsPerPage: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("startIndex") val startIndex: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("totalItems") val totalItems: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("totalAvailableItems") val totalAvailableItems: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("pageIndex") val pageIndex: MaybeField<Int> = MaybeField(Maybe.Absent),
    @SerialName("totalPages") val totalPages: MaybeField<Int> = MaybeField(Maybe.Absent),
)

@Serializable
data class GResponseError(
    @SerialName("code") val code: Int = 0,
    @SerialName("message") val message: String = "",
    @SerialName("messageTranslated") val messageTranslated: String = "",
    @SerialName("errors") val errors: List<GResponseErrorDetail> = emptyList(),
)

@Serializable
data class GResponseErrorDetail(
    @SerialName("domain") val domain: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("reason") val reason: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("message") val message: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("messageTranslated") val messageTranslated: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("location") val location: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("locationType") val locationType: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("extendedHelp") val extendedHelp: MaybeField<String> = MaybeField(Maybe.Absent),
    @SerialName("sendReport") val sendReport: MaybeField<String> = MaybeField(Maybe.Absent),
)
