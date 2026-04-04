package unknownpackage
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.Maybe
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import emikot.MaybeField
/**
 * Action to communicate with the action GetSelectedDataOfTheParticularProductOfferAction
 */
data class GetSelectedDataOfTheParticularProductOfferActionMeta(
    val name: String = "GetSelectedDataOfTheParticularProductOfferAction",
    val url: String = "https://api.{environment}/sale/product-offers/{offerId}/parts",
    val method: String = "get"
)
/*data class GetSelectedDataOfTheParticularProductOfferActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetSelectedDataOfTheParticularProductOfferActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetSelectedDataOfTheParticularProductOfferActionClient {
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
	): GetSelectedDataOfTheParticularProductOfferActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetSelectedDataOfTheParticularProductOfferActionMeta()
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
                GetSelectedDataOfTheParticularProductOfferActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getSelectedDataOfTheParticularProductOfferActionRes
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionRes (
		  // Unique offer identifier
 @SerialName("id")  val id: String  = "",
		@SerialName("stock")  val stock:  GetSelectedDataOfTheParticularProductOfferActionResStock ,
		@SerialName("sellingMode")  val sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResSellingMode ,
		  // Marketplace-specific price information
 @SerialName("additionalMarketplaces")  val additionalMarketplaces:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces ,
)
  // The base class definition for stock
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResStock (
		  // Number of available items in stock
 @SerialName("available")  val available: Int  = 0,
)
  // The base class definition for sellingMode
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResSellingMode (
		@SerialName("price")  val price:  GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice ,
)
  // The base class definition for price
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces (
		@SerialName("marketplaceId1")  val marketplaceId1:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 ,
		@SerialName("marketplaceId2")  val marketplaceId2:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 ,
)
  // The base class definition for marketplaceId1
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1 (
		@SerialName("sellingMode")  val sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode ,
)
  // The base class definition for sellingMode
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode (
		@SerialName("price")  val price:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice ,
)
  // The base class definition for price
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for marketplaceId2
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2 (
		@SerialName("sellingMode")  val sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode ,
)
  // The base class definition for sellingMode
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode (
		@SerialName("price")  val price:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice ,
)
  // The base class definition for price
@Serializable
data class GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)