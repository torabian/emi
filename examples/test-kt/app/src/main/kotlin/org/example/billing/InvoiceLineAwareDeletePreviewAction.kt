package org.example.billing
import emikot.ClientContext
import emikot.GResponse
import emikot.Maybe
import emikot.MaybeField
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
/**
 * Action to communicate with the action InvoiceLineAwareDeletePreviewAction
 */
data class InvoiceLineAwareDeletePreviewActionMeta(
    val name: String = "InvoiceLineAwareDeletePreviewAction",
    val url: String = "/invoiceLine/delete-preview",
    val method: String = "get"
)
/*data class InvoiceLineAwareDeletePreviewActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class InvoiceLineAwareDeletePreviewActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<InvoiceLineAwareDeletePreviewActionRes>? = null
)
object InvoiceLineAwareDeletePreviewActionClient {
	public var context: ClientContext? = null
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
    fun buildUrl(base: String, path: String, query: Map<String, String>): String {
        val baseUrl = base.toHttpUrl()   // parses full URL like "http://asdasda/"
        val urlBuilder = baseUrl
            .newBuilder()
            .encodedPath(path)
        query.forEach { (k, v) ->
            urlBuilder.addQueryParameter(k, v)
        }
        return urlBuilder.build().toString()
    }
    // compute() actually serializes no body (this action has none)
    // via kotlinx.serialization before sending, and decodes the raw response body
    // into GResponse<InvoiceLineAwareDeletePreviewActionRes> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): InvoiceLineAwareDeletePreviewActionResponse =
        withContext(Dispatchers.IO) {
            val meta = InvoiceLineAwareDeletePreviewActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            val body0: RequestBody? = null
            val requestBuilder = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
            headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                val parsedPayload: GResponse<InvoiceLineAwareDeletePreviewActionRes>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<InvoiceLineAwareDeletePreviewActionRes>>(it)
                }
                InvoiceLineAwareDeletePreviewActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}
  // The base class definition for invoiceLineAwareDeletePreviewActionRes
@Serializable
data class InvoiceLineAwareDeletePreviewActionRes (
		@SerialName("message")  val message: String  = "",
		@SerialName("affected")  val affected: List<InvoiceLineAwareDeletePreviewActionResAffected>  = emptyList(),
)
  // The base class definition for affected
@Serializable
data class InvoiceLineAwareDeletePreviewActionResAffected (
		@SerialName("relation")  val relation: String  = "",
		@SerialName("count")  val count: Long  = 0,
)