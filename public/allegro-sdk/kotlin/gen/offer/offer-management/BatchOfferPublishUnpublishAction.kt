package unknownpackage
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.ClientContext
import kotlinx.serialization.json.*
import emikot.Maybe
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import emikot.MaybeField
/**
 * Action to communicate with the action BatchOfferPublishUnpublishAction
 */
data class BatchOfferPublishUnpublishActionMeta(
    val name: String = "BatchOfferPublishUnpublishAction",
    val url: String = "https://api.{environment}/sale/offer-publication-commands/{commandId}",
    val method: String = "put"
)
/*data class BatchOfferPublishUnpublishActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class BatchOfferPublishUnpublishActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object BatchOfferPublishUnpublishActionClient {
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
	): BatchOfferPublishUnpublishActionResponse =
        withContext(Dispatchers.IO) {
            val meta = BatchOfferPublishUnpublishActionMeta()
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
                BatchOfferPublishUnpublishActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for batchOfferPublishUnpublishActionReq
@Serializable
data class BatchOfferPublishUnpublishActionReq (
		@SerialName("offerCriteria")  val offerCriteria: List<BatchOfferPublishUnpublishActionReqOfferCriteria>  = emptyList(),
		@SerialName("publication")  val publication:  BatchOfferPublishUnpublishActionReqPublication ,
)
  // The base class definition for offerCriteria
@Serializable
data class BatchOfferPublishUnpublishActionReqOfferCriteria (
		@SerialName("offers")  val offers: List<BatchOfferPublishUnpublishActionReqOfferCriteriaOffers>  = emptyList(),
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for offers
@Serializable
data class BatchOfferPublishUnpublishActionReqOfferCriteriaOffers (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for publication
@Serializable
data class BatchOfferPublishUnpublishActionReqPublication (
		@SerialName("action")  val action: String  = "",
		@SerialName("scheduledFor")  val scheduledFor: String  = "",
)
  // The base class definition for batchOfferPublishUnpublishActionRes
@Serializable
data class BatchOfferPublishUnpublishActionRes (
		@SerialName("id")  val id: String  = "",
		@SerialName("createdAt")  val createdAt: String  = "",
		@SerialName("completedAt")  val completedAt: String  = "",
		@SerialName("taskCount")  val taskCount:  BatchOfferPublishUnpublishActionResTaskCount ,
)
  // The base class definition for taskCount
@Serializable
data class BatchOfferPublishUnpublishActionResTaskCount (
		@SerialName("failed")  val failed: Int  = 0,
		@SerialName("success")  val success: Int  = 0,
		@SerialName("total")  val total: Int  = 0,
)