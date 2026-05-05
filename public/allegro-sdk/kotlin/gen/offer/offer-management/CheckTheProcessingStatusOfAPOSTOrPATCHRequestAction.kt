package unknownpackage
import emikot.MaybeField
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.json.*
import emikot.Maybe
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.ClientContext
import kotlinx.serialization.*
/**
 * Action to communicate with the action CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction
 */
data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta(
    val name: String = "CheckTheProcessingStatusOfAPOSTOrPATCHRequestAction",
    val url: String = "https://api.{environment}/sale/product-offers/{offerId}/operations/{operationId}",
    val method: String = "get"
)
/*data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionClient {
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
	): CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse =
        withContext(Dispatchers.IO) {
            val meta = CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionMeta()
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
                CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for checkTheProcessingStatusOfAPOSTOrPATCHRequestActionRes
@Serializable
data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionRes (
		@SerialName("offer")  val offer:  CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer ,
		@SerialName("operation")  val operation:  CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation ,
)
  // The base class definition for offer
@Serializable
data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOffer (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for operation
@Serializable
data class CheckTheProcessingStatusOfAPOSTOrPATCHRequestActionResOperation (
		@SerialName("id")  val id: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("startedAt")  val startedAt: String  = "",
)