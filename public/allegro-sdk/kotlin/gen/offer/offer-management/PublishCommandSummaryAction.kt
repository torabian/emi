package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import emikot.ClientContext
import kotlinx.serialization.json.*
import emikot.MaybeField
import emikot.Maybe
/**
 * Action to communicate with the action PublishCommandSummaryAction
 */
data class PublishCommandSummaryActionMeta(
    val name: String = "PublishCommandSummaryAction",
    val url: String = "https://api.{environment}/sale/offer-publication-commands/{commandId}",
    val method: String = "get"
)
/*data class PublishCommandSummaryActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class PublishCommandSummaryActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object PublishCommandSummaryActionClient {
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
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): PublishCommandSummaryActionResponse =
        withContext(Dispatchers.IO) {
            val meta = PublishCommandSummaryActionMeta()
            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)
            println(  url)
            val body0 = body?.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                PublishCommandSummaryActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for publishCommandSummaryActionRes
@Serializable
data class PublishCommandSummaryActionRes (
		@SerialName("id")  val id: String  = "",
		@SerialName("createdAt")  val createdAt: String  = "",
		@SerialName("completedAt")  val completedAt: String  = "",
		@SerialName("taskCount")  val taskCount:  PublishCommandSummaryActionResTaskCount ,
)
  // The base class definition for taskCount
@Serializable
data class PublishCommandSummaryActionResTaskCount (
		@SerialName("failed")  val failed: Int  = 0,
		@SerialName("success")  val success: Int  = 0,
		@SerialName("total")  val total: Int  = 0,
)