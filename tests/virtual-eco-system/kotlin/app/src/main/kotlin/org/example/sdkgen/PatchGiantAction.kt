package unknownpackage
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
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
		url = url.replace("id", params.Id)
	return url
}
object PatchGiantActionClient {
	public var context: ClientContext? = null
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
	fun buildUrl(base: String, path: String, query: Map<String, String>): String {
		val httpUrl = HttpUrl.Builder()
			.scheme("http")    // or dynamic
			.host("localhost") // or dynamic
			.encodedPath(path)
		query.forEach { (k, v) -> httpUrl.addQueryParameter(k, v) }
		return httpUrl.build().toString()
	}
    suspend fun compute(
		path: PatchGiantActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): PatchGiantActionResponse =
        withContext(Dispatchers.IO) {
            val meta = PatchGiantActionMeta()
            val body0 = body?.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(meta.url)
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