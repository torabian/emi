package unknownpackage
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
 * Action to communicate with the action MemberGetAction
 */
data class MemberGetActionMeta(
    val name: String = "MemberGetAction",
    val url: String = "/member/:uniqueId",
    val method: String = "get"
)
/*data class MemberGetActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class MemberGetActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<MemberDto>? = null
)
	/**
 * Path parameters for MemberGetAction
 */
data class MemberGetActionPathParameter (
	var UniqueId: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun MemberGetActionPathParameterApply(params: MemberGetActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":uniqueId", params.UniqueId)
	return url
}
object MemberGetActionClient {
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
    // into GResponse<MemberDto> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		path: MemberGetActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): MemberGetActionResponse =
        withContext(Dispatchers.IO) {
            val meta = MemberGetActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            	url = MemberGetActionPathParameterApply(path, url)
            val body0: RequestBody? = null
            val requestBuilder = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
            headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                val parsedPayload: GResponse<MemberDto>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<MemberDto>>(it)
                }
                MemberGetActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}