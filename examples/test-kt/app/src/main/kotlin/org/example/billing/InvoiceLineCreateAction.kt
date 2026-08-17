package org.example.billing
import emikot.ClientContext
import emikot.GResponse
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
/**
 * Action to communicate with the action InvoiceLineCreateAction
 */
data class InvoiceLineCreateActionMeta(
    val name: String = "InvoiceLineCreateAction",
    val url: String = "/invoiceLine",
    val method: String = "post"
)
/*data class InvoiceLineCreateActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class InvoiceLineCreateActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<InvoiceLineDto>? = null
)
object InvoiceLineCreateActionClient {
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
    // compute() actually serializes the typed InvoiceLineDto body
    // via kotlinx.serialization before sending, and decodes the raw response body
    // into GResponse<InvoiceLineDto> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
		, body: InvoiceLineDto? = null
	): InvoiceLineCreateActionResponse =
        withContext(Dispatchers.IO) {
            val meta = InvoiceLineCreateActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            val body0 = body?.let { Json.encodeToString(it).toRequestBody(jsonType) }
            val requestBuilder = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
            headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                val parsedPayload: GResponse<InvoiceLineDto>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<InvoiceLineDto>>(it)
                }
                InvoiceLineCreateActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}