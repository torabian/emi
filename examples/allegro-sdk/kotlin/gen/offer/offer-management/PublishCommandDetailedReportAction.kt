package unknownpackage
import emikot.ClientContext
import emikot.MaybeField
import okhttp3.*
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.Maybe
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
/**
 * Action to communicate with the action PublishCommandDetailedReportAction
 */
data class PublishCommandDetailedReportActionMeta(
    val name: String = "PublishCommandDetailedReportAction",
    val url: String = "https://api.{environment}/sale/offer-publication-commands/{commandId}/tasks",
    val method: String = "get"
)
/*data class PublishCommandDetailedReportActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class PublishCommandDetailedReportActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object PublishCommandDetailedReportActionClient {
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
	): PublishCommandDetailedReportActionResponse =
        withContext(Dispatchers.IO) {
            val meta = PublishCommandDetailedReportActionMeta()
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
                PublishCommandDetailedReportActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for publishCommandDetailedReportActionRes
@Serializable
data class PublishCommandDetailedReportActionRes (
		@SerialName("tasks")  val tasks: List<PublishCommandDetailedReportActionResTasks>  = emptyList(),
)
  // The base class definition for tasks
@Serializable
data class PublishCommandDetailedReportActionResTasks (
		@SerialName("field")  val field: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("offer")  val offer:  PublishCommandDetailedReportActionResTasksOffer ,
		@SerialName("status")  val status: String  = "",
		@SerialName("errors")  val errors: List<PublishCommandDetailedReportActionResTasksErrors>  = emptyList(),
)
  // The base class definition for offer
@Serializable
data class PublishCommandDetailedReportActionResTasksOffer (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for errors
@Serializable
data class PublishCommandDetailedReportActionResTasksErrors (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  PublishCommandDetailedReportActionResTasksErrorsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class PublishCommandDetailedReportActionResTasksErrorsMetadata (
		@SerialName("productId")  val productId: String  = "",
)