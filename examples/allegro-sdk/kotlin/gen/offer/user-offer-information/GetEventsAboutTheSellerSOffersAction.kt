package unknownpackage
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import emikot.MaybeField
import emikot.Maybe
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
/**
 * Action to communicate with the action GetEventsAboutTheSellerSOffersAction
 */
data class GetEventsAboutTheSellerSOffersActionMeta(
    val name: String = "GetEventsAboutTheSellerSOffersAction",
    val url: String = "https://api.{environment}/sale/offer-events",
    val method: String = "get"
)
/*data class GetEventsAboutTheSellerSOffersActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetEventsAboutTheSellerSOffersActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetEventsAboutTheSellerSOffersActionClient {
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
	): GetEventsAboutTheSellerSOffersActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetEventsAboutTheSellerSOffersActionMeta()
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
                GetEventsAboutTheSellerSOffersActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getEventsAboutTheSellerSOffersActionRes
@Serializable
data class GetEventsAboutTheSellerSOffersActionRes (
		  // List of events related to offer state changes
 @SerialName("offerEvents")  val offerEvents: List<GetEventsAboutTheSellerSOffersActionResOfferEvents>  = emptyList(),
)
  // The base class definition for offerEvents
@Serializable
data class GetEventsAboutTheSellerSOffersActionResOfferEvents (
		  // Unique event identifier (base64 encoded)
 @SerialName("id")  val id: String  = "",
		  // ISO8601 timestamp when the event occurred
 @SerialName("occurredAt")  val occurredAt: String  = "",
		  // Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
 @SerialName("type")  val type: String  = "",
		  // Basic offer information for which event occurred
 @SerialName("offer")  val offer:  GetEventsAboutTheSellerSOffersActionResOfferEventsOffer ,
)
  // The base class definition for offer
@Serializable
data class GetEventsAboutTheSellerSOffersActionResOfferEventsOffer (
		@SerialName("id")  val id: String  = "",
		@SerialName("external")  val external:  GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal ,
)
  // The base class definition for external
@Serializable
data class GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal (
		@SerialName("id")  val id: String  = "",
)