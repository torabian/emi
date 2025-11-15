package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.withContext
import kotlinx.serialization.json.*
import kotlinx.coroutines.Dispatchers
import emikot.ClientContext
import kotlinx.serialization.*
/**
 * Action to communicate with the action GetAsGiantsAction
 */
data class GetAsGiantsActionMeta(
    val name: String = "GetAsGiantsAction",
    val url: String = "/get/giant/:id",
    val method: String = "get"
)
/*data class GetAsGiantsActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetAsGiantsActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
	/**
 * Path parameters for GetAsGiantsAction
 */
data class GetAsGiantsActionPathParameter (
	var Id: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun GetAsGiantsActionPathParameterApply(params: GetAsGiantsActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":id", params.Id)
	return url
}
object GetAsGiantsActionClient {
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
    suspend fun compute(
		path: GetAsGiantsActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): GetAsGiantsActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetAsGiantsActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            	url = GetAsGiantsActionPathParameterApply(path, url)
            println(  url)
            val body0 = body?.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                GetAsGiantsActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}