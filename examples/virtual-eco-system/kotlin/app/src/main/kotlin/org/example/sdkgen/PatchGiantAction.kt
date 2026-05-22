package unknownpackage
import okhttp3.MediaType.Companion.toMediaType
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import okhttp3.*
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import emikot.ClientContext
import kotlinx.serialization.json.*
/**
 * Action to communicate with the action PatchGiantAction
 */
data class PatchGiantActionMeta(
    val name: String = "PatchGiantAction",
    val url: String = "/get/giant/:id",
    val method: String = "patch"
)
/*data class PatchGiantActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class PatchGiantActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
	/**
 * Path parameters for PatchGiantAction
 */
data class PatchGiantActionPathParameter (
	var Id: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun PatchGiantActionPathParameterApply(params: PatchGiantActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":id", params.Id)
	return url
}
object PatchGiantActionClient {
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
		path: PatchGiantActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): PatchGiantActionResponse =
        withContext(Dispatchers.IO) {
            val meta = PatchGiantActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            	url = PatchGiantActionPathParameterApply(path, url)
            println(  url)
            val body0 = body?.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                PatchGiantActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}