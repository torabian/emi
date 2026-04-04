package unknownpackage
import okhttp3.*
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.MaybeField
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.Maybe
/**
 * Action to communicate with the action GetOffersWithMissingParametersAction
 */
data class GetOffersWithMissingParametersActionMeta(
    val name: String = "GetOffersWithMissingParametersAction",
    val url: String = "https://api.{environment}/sale/offers/unfilled-parameters",
    val method: String = "get"
)
/*data class GetOffersWithMissingParametersActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetOffersWithMissingParametersActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetOffersWithMissingParametersActionClient {
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
	): GetOffersWithMissingParametersActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetOffersWithMissingParametersActionMeta()
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
                GetOffersWithMissingParametersActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getOffersWithMissingParametersActionRes
@Serializable
data class GetOffersWithMissingParametersActionRes (
		@SerialName("offers")  val offers: List<GetOffersWithMissingParametersActionResOffers>  = emptyList(),
		@SerialName("count")  val count: Int  = 0,
		@SerialName("totalCount")  val totalCount: Int  = 0,
)
  // The base class definition for offers
@Serializable
data class GetOffersWithMissingParametersActionResOffers (
		@SerialName("id")  val id: String  = "",
		@SerialName("parameters")  val parameters: List<GetOffersWithMissingParametersActionResOffersParameters>  = emptyList(),
		@SerialName("category")  val category:  GetOffersWithMissingParametersActionResOffersCategory ,
)
  // The base class definition for parameters
@Serializable
data class GetOffersWithMissingParametersActionResOffersParameters (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for category
@Serializable
data class GetOffersWithMissingParametersActionResOffersCategory (
		@SerialName("id")  val id: String  = "",
)