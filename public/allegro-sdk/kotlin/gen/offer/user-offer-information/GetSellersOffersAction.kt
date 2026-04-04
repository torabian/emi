package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import emikot.ClientContext
import emikot.Maybe
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.MaybeField
/**
 * Action to communicate with the action GetSellersOffersAction
 */
data class GetSellersOffersActionMeta(
    val name: String = "GetSellersOffersAction",
    val url: String = "https://api.{environment}/sale/offers",
    val method: String = "get"
)
/*data class GetSellersOffersActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetSellersOffersActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetSellersOffersActionClient {
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
	): GetSellersOffersActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetSellersOffersActionMeta()
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
                GetSellersOffersActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getSellersOffersActionRes
@Serializable
data class GetSellersOffersActionRes (
		  // Number of offers in this page
 @SerialName("count")  val count: Int  = 0,
		  // Total number of offers available
 @SerialName("totalCount")  val totalCount: Int  = 0,
		@SerialName("offers")  val offers: List<GetSellersOffersActionResOffers>  = emptyList(),
)
  // The base class definition for offers
@Serializable
data class GetSellersOffersActionResOffers (
		  // Offer identifier
 @SerialName("id")  val id: String  = "",
		  // Offer name or title
 @SerialName("name")  val name: String  = "",
		@SerialName("category")  val category:  GetSellersOffersActionResOffersCategory ,
		@SerialName("primaryImage")  val primaryImage:  GetSellersOffersActionResOffersPrimaryImage ,
		@SerialName("sellingMode")  val sellingMode:  GetSellersOffersActionResOffersSellingMode ,
		@SerialName("saleInfo")  val saleInfo:  GetSellersOffersActionResOffersSaleInfo ,
		@SerialName("stock")  val stock:  GetSellersOffersActionResOffersStock ,
		@SerialName("stats")  val stats:  GetSellersOffersActionResOffersStats ,
		@SerialName("publication")  val publication:  GetSellersOffersActionResOffersPublication ,
		@SerialName("afterSalesServices")  val afterSalesServices:  GetSellersOffersActionResOffersAfterSalesServices ,
		@SerialName("additionalServices")  val additionalServices:  GetSellersOffersActionResOffersAdditionalServices ,
		@SerialName("external")  val external:  GetSellersOffersActionResOffersExternal ,
		@SerialName("delivery")  val delivery:  GetSellersOffersActionResOffersDelivery ,
		@SerialName("b2b")  val b2b:  GetSellersOffersActionResOffersB2b ,
		@SerialName("fundraisingCampaign")  val fundraisingCampaign:  GetSellersOffersActionResOffersFundraisingCampaign ,
		  // Marketplace-specific extensions for offer
 @SerialName("additionalMarketplaces")  @Contextual  val additionalMarketplaces: Any ,
)
  // The base class definition for category
@Serializable
data class GetSellersOffersActionResOffersCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for primaryImage
@Serializable
data class GetSellersOffersActionResOffersPrimaryImage (
		@SerialName("url")  val url: String  = "",
)
  // The base class definition for sellingMode
@Serializable
data class GetSellersOffersActionResOffersSellingMode (
		@SerialName("format")  val format: String  = "",
		@SerialName("price")  val price:  GetSellersOffersActionResOffersSellingModePrice ,
		@SerialName("priceAutomation")  val priceAutomation:  GetSellersOffersActionResOffersSellingModePriceAutomation ,
		@SerialName("minimalPrice")  val minimalPrice:  GetSellersOffersActionResOffersSellingModeMinimalPrice ,
		@SerialName("startingPrice")  val startingPrice:  GetSellersOffersActionResOffersSellingModeStartingPrice ,
)
  // The base class definition for price
@Serializable
data class GetSellersOffersActionResOffersSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for priceAutomation
@Serializable
data class GetSellersOffersActionResOffersSellingModePriceAutomation (
		@SerialName("rule")  val rule:  GetSellersOffersActionResOffersSellingModePriceAutomationRule ,
)
  // The base class definition for rule
@Serializable
data class GetSellersOffersActionResOffersSellingModePriceAutomationRule (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for minimalPrice
@Serializable
data class GetSellersOffersActionResOffersSellingModeMinimalPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for startingPrice
@Serializable
data class GetSellersOffersActionResOffersSellingModeStartingPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for saleInfo
@Serializable
data class GetSellersOffersActionResOffersSaleInfo (
		@SerialName("currentPrice")  val currentPrice:  GetSellersOffersActionResOffersSaleInfoCurrentPrice ,
		@SerialName("biddersCount")  val biddersCount: Int  = 0,
)
  // The base class definition for currentPrice
@Serializable
data class GetSellersOffersActionResOffersSaleInfoCurrentPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for stock
@Serializable
data class GetSellersOffersActionResOffersStock (
		@SerialName("available")  val available: Int  = 0,
		@SerialName("sold")  val sold: Int  = 0,
)
  // The base class definition for stats
@Serializable
data class GetSellersOffersActionResOffersStats (
		@SerialName("watchersCount")  val watchersCount: Int  = 0,
		@SerialName("visitsCount")  val visitsCount: Int  = 0,
)
  // The base class definition for publication
@Serializable
data class GetSellersOffersActionResOffersPublication (
		@SerialName("status")  val status: String  = "",
		@SerialName("startingAt")  val startingAt: String  = "",
		@SerialName("startedAt")  val startedAt: String  = "",
		@SerialName("endingAt")  val endingAt: String  = "",
		@SerialName("endedAt")  val endedAt: String  = "",
		@SerialName("marketplaces")  val marketplaces:  GetSellersOffersActionResOffersPublicationMarketplaces ,
)
  // The base class definition for marketplaces
@Serializable
data class GetSellersOffersActionResOffersPublicationMarketplaces (
		@SerialName("base")  val base:  GetSellersOffersActionResOffersPublicationMarketplacesBase ,
		@SerialName("additional")  val additional: List<GetSellersOffersActionResOffersPublicationMarketplacesAdditional>  = emptyList(),
)
  // The base class definition for base
@Serializable
data class GetSellersOffersActionResOffersPublicationMarketplacesBase (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additional
@Serializable
data class GetSellersOffersActionResOffersPublicationMarketplacesAdditional (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for afterSalesServices
@Serializable
data class GetSellersOffersActionResOffersAfterSalesServices (
		@SerialName("impliedWarranty")  val impliedWarranty:  GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty ,
		@SerialName("returnPolicy")  val returnPolicy:  GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy ,
		@SerialName("warranty")  val warranty:  GetSellersOffersActionResOffersAfterSalesServicesWarranty ,
)
  // The base class definition for impliedWarranty
@Serializable
data class GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for returnPolicy
@Serializable
data class GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for warranty
@Serializable
data class GetSellersOffersActionResOffersAfterSalesServicesWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalServices
@Serializable
data class GetSellersOffersActionResOffersAdditionalServices (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for external
@Serializable
data class GetSellersOffersActionResOffersExternal (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for delivery
@Serializable
data class GetSellersOffersActionResOffersDelivery (
		@SerialName("shippingRates")  val shippingRates:  GetSellersOffersActionResOffersDeliveryShippingRates ,
)
  // The base class definition for shippingRates
@Serializable
data class GetSellersOffersActionResOffersDeliveryShippingRates (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for b2b
@Serializable
data class GetSellersOffersActionResOffersB2b (
		@SerialName("buyableOnlyByBusiness")  val buyableOnlyByBusiness: Boolean ,
)
  // The base class definition for fundraisingCampaign
@Serializable
data class GetSellersOffersActionResOffersFundraisingCampaign (
		@SerialName("id")  val id: String  = "",
)