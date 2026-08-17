package unknownpackage
import emikot.ClientContext
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
 * Action to communicate with the action TeamAwareDeleteAction
 */
data class TeamAwareDeleteActionMeta(
    val name: String = "TeamAwareDeleteAction",
    val url: String = "/team/delete",
    val method: String = "post"
)
/*data class TeamAwareDeleteActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class TeamAwareDeleteActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: Any? = null
)
object TeamAwareDeleteActionClient {
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
    // compute() actually serializes the typed TeamAwareDeleteActionReq body
    // via kotlinx.serialization before sending, and decodes the raw response body
    // (left untyped - this action has no response shape) once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
		, body: TeamAwareDeleteActionReq? = null
	): TeamAwareDeleteActionResponse =
        withContext(Dispatchers.IO) {
            val meta = TeamAwareDeleteActionMeta()
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
                TeamAwareDeleteActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = null
                )
            }
        }
}
  // The base class definition for teamAwareDeleteActionReq
@Serializable
data class TeamAwareDeleteActionReq (
		@SerialName("uniqueIds")  val uniqueIds: List<String>  = emptyList(),
)