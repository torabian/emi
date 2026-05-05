package unknownpackage
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.serialization.json.*
import emikot.MaybeField
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import emikot.Maybe
/**
 * Action to communicate with the action GetSmartClassificationReportOfTheParticularOfferAction
 */
data class GetSmartClassificationReportOfTheParticularOfferActionMeta(
    val name: String = "GetSmartClassificationReportOfTheParticularOfferAction",
    val url: String = "https://api.{environment}/sale/offers/{offerId}/smart",
    val method: String = "get"
)
/*data class GetSmartClassificationReportOfTheParticularOfferActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetSmartClassificationReportOfTheParticularOfferActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetSmartClassificationReportOfTheParticularOfferActionClient {
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
	): GetSmartClassificationReportOfTheParticularOfferActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetSmartClassificationReportOfTheParticularOfferActionMeta()
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
                GetSmartClassificationReportOfTheParticularOfferActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getSmartClassificationReportOfTheParticularOfferActionRes
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionRes (
		  // Indicates if offer is queued for reclassification
 @SerialName("scheduledForReclassification")  val scheduledForReclassification: Boolean ,
		  // Offer classification status and last change date
 @SerialName("classification")  val classification:  GetSmartClassificationReportOfTheParticularOfferActionResClassification ,
		  // List of smart delivery method identifiers
 @SerialName("smartDeliveryMethods")  val smartDeliveryMethods: List<GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods>  = emptyList(),
		  // List of classification conditions with delivery method checks
 @SerialName("conditions")  val conditions: List<GetSmartClassificationReportOfTheParticularOfferActionResConditions>  = emptyList(),
)
  // The base class definition for classification
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionResClassification (
		  // Whether the classification conditions are fulfilled
 @SerialName("fulfilled")  val fulfilled: Boolean ,
		  // ISO8601 timestamp of last classification change
 @SerialName("lastChanged")  val lastChanged: String  = "",
)
  // The base class definition for smartDeliveryMethods
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for conditions
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionResConditions (
		  // Condition code identifier
 @SerialName("code")  val code: String  = "",
		  // Human-readable condition name
 @SerialName("name")  val name: String  = "",
		  // Detailed condition description
 @SerialName("description")  val description: String  = "",
		  // Indicates if this condition is fulfilled
 @SerialName("fulfilled")  val fulfilled: Boolean ,
		  // Delivery methods that passed validation for this condition
 @SerialName("passedDeliveryMethods")  val passedDeliveryMethods: List<GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods>  = emptyList(),
		  // Delivery methods that failed validation for this condition
 @SerialName("failedDeliveryMethods")  val failedDeliveryMethods: List<GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods>  = emptyList(),
)
  // The base class definition for passedDeliveryMethods
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for failedDeliveryMethods
@Serializable
data class GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods (
		@SerialName("id")  val id: String  = "",
)