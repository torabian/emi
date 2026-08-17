package unknownpackage
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
 * Action to communicate with the action TeamAwareDeletePreviewAction
 */
data class TeamAwareDeletePreviewActionMeta(
    val name: String = "TeamAwareDeletePreviewAction",
    val url: String = "/team/delete-preview",
    val method: String = "get"
)
/*data class TeamAwareDeletePreviewActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class TeamAwareDeletePreviewActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: GResponse<TeamAwareDeletePreviewActionRes>? = null
)
object TeamAwareDeletePreviewActionClient {
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
    // into GResponse<TeamAwareDeletePreviewActionRes> once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): TeamAwareDeletePreviewActionResponse =
        withContext(Dispatchers.IO) {
            val meta = TeamAwareDeletePreviewActionMeta()
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
                val parsedPayload: GResponse<TeamAwareDeletePreviewActionRes>? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<GResponse<TeamAwareDeletePreviewActionRes>>(it)
                }
                TeamAwareDeletePreviewActionResponse(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = parsedPayload
                )
            }
        }
}
  // The base class definition for teamAwareDeletePreviewActionRes
@Serializable
data class TeamAwareDeletePreviewActionRes (
		@SerialName("message")  val message: String  = "",
		@SerialName("affected")  val affected: List<TeamAwareDeletePreviewActionResAffected>  = emptyList(),
)
  // The base class definition for affected
@Serializable
data class TeamAwareDeletePreviewActionResAffected (
		@SerialName("relation")  val relation: String  = "",
		@SerialName("count")  val count: Long  = 0,
)