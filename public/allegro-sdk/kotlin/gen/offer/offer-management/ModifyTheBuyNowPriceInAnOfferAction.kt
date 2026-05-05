package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.MaybeField
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.Maybe
/**
 * Action to communicate with the action ModifyTheBuyNowPriceInAnOfferAction
 */
data class ModifyTheBuyNowPriceInAnOfferActionMeta(
    val name: String = "ModifyTheBuyNowPriceInAnOfferAction",
    val url: String = "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}",
    val method: String = "put"
)
/*data class ModifyTheBuyNowPriceInAnOfferActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ModifyTheBuyNowPriceInAnOfferActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ModifyTheBuyNowPriceInAnOfferActionClient {
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
	): ModifyTheBuyNowPriceInAnOfferActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ModifyTheBuyNowPriceInAnOfferActionMeta()
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
                ModifyTheBuyNowPriceInAnOfferActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionReq (
		@SerialName("id")  val id: String  = "",
		@SerialName("input")  val input:  ModifyTheBuyNowPriceInAnOfferActionReqInput ,
)
  // The base class definition for input
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionReqInput (
		@SerialName("buyNowPrice")  val buyNowPrice:  ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice ,
)
  // The base class definition for buyNowPrice
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionRes (
		@SerialName("id")  val id: String  = "",
		@SerialName("input")  val input:  ModifyTheBuyNowPriceInAnOfferActionResInput ,
		@SerialName("output")  val output:  ModifyTheBuyNowPriceInAnOfferActionResOutput ,
)
  // The base class definition for input
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionResInput (
		@SerialName("buyNowPrice")  val buyNowPrice:  ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice ,
)
  // The base class definition for buyNowPrice
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for output
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionResOutput (
		@SerialName("status")  val status: String  = "",
		@SerialName("errors")  val errors: List<ModifyTheBuyNowPriceInAnOfferActionResOutputErrors>  = emptyList(),
)
  // The base class definition for errors
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionResOutputErrors (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata (
		@SerialName("productId")  val productId: String  = "",
)