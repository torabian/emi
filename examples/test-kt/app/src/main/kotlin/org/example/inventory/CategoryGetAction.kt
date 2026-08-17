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
 * Action to communicate with the action CategoryGetAction
 */
data class CategoryGetActionMeta(
    val name: String = "CategoryGetAction",
    val url: String = "/category/:uniqueId",
    val method: String = "get"
)
/*data class CategoryGetActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class CategoryGetActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<CategoryDto>? = null
)
	/**
 * Path parameters for CategoryGetAction
 */
data class CategoryGetActionPathParameter (
	var UniqueId: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun CategoryGetActionPathParameterApply(params: CategoryGetActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":uniqueId", params.UniqueId)
	return url
}
object CategoryGetActionClient {
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
    // into GResponse<CategoryDto> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		path: CategoryGetActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): CategoryGetActionResponse =
        withContext(Dispatchers.IO) {
            val meta = CategoryGetActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            	url = CategoryGetActionPathParameterApply(path, url)
            val body0: RequestBody? = null
            val requestBuilder = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
            headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                val parsedPayload: GResponse<CategoryDto>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<CategoryDto>>(it)
                }
                CategoryGetActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}