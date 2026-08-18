package org.example.inventory
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
 * Action to communicate with the action ProductGetAction
 */
data class ProductGetActionMeta(
    val name: String = "ProductGetAction",
    val url: String = "/product/:uniqueId",
    val method: String = "get"
)
/*data class ProductGetActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ProductGetActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<ProductDto>? = null
)
	/**
 * Path parameters for ProductGetAction
 */
data class ProductGetActionPathParameter (
	var UniqueId: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun ProductGetActionPathParameterApply(params: ProductGetActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":uniqueId", params.UniqueId)
	return url
}
object ProductGetActionClient {
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
    // into GResponse<ProductDto> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		path: ProductGetActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): ProductGetActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ProductGetActionMeta()
            // Falls back to the app-wide ClientContext.Default when this client's own
            // .context hasn't been set - see ClientContext's doc comment (common.kt)
            // and the "Client context & authentication" Kotlin doc page.
            val effectiveContext = context ?: ClientContext.Default
            var url = buildUrl(effectiveContext.baseUrl, meta.url, query)
            	url = ProductGetActionPathParameterApply(path, url)
            val body0: RequestBody? = null
            // Merges defaultHeaders under the explicit headers argument, then runs
            // ClientContext.onRequest if set (e.g. injecting a fresh auth token).
            val resolved = effectiveContext.resolve(url, headers)
            val requestBuilder = Request.Builder()
                .url(resolved.url)
                // HTTP methods are conventionally sent uppercase (meta.method itself
                // stays exactly as declared in the .emi.yml, e.g. "get"/"post", for
                // backwards-compat with anything already reading it) - some servers
                // are strict about it, and OkHttp itself special-cases request-body
                // permission (HttpMethod.permitsRequestBody/...) on the uppercase form.
                .method(meta.method.uppercase(), body0)
                .addHeader("Accept", "application/json")
            resolved.headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                val parsedPayload: GResponse<ProductDto>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<ProductDto>>(it)
                }
                ProductGetActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}