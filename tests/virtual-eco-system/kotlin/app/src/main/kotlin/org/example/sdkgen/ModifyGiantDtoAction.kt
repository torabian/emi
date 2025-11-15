package unknownpackage
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
/**
 * Action to communicate with the action ModifyGiantDtoAction
 */
data class ModifyGiantDtoActionMeta(
    val name: String = "ModifyGiantDtoAction",
    val url: String = "/modify/dto",
    val method: String = "post"
)
/*data class ModifyGiantDtoActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ModifyGiantDtoActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ModifyGiantDtoActionClient {
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
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): ModifyGiantDtoActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ModifyGiantDtoActionMeta()
            val body0 = body?.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(meta.url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                ModifyGiantDtoActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}