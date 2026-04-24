package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.MaybeField
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.Maybe
/**
 * Action to communicate with the action CreateOfferBasedOnProductAction
 */
data class CreateOfferBasedOnProductActionMeta(
    val name: String = "CreateOfferBasedOnProductAction",
    val url: String = "https://api.{environment}/sale/product-offers",
    val method: String = "post"
)
/*data class CreateOfferBasedOnProductActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class CreateOfferBasedOnProductActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object CreateOfferBasedOnProductActionClient {
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
	): CreateOfferBasedOnProductActionResponse =
        withContext(Dispatchers.IO) {
            val meta = CreateOfferBasedOnProductActionMeta()
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
                CreateOfferBasedOnProductActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for createOfferBasedOnProductActionReq
@Serializable
data class CreateOfferBasedOnProductActionReq (
		  // Offer title
 @SerialName("name")  val name: String  = "",
		  // Offer language code (e.g., pl-PL)
 @SerialName("language")  val language: String  = "",
		@SerialName("category")  val category:  CreateOfferBasedOnProductActionReqCategory ,
		  // Product details and associated quantities
 @SerialName("productSet")  val productSet: List<CreateOfferBasedOnProductActionReqProductSet>  = emptyList(),
		@SerialName("stock")  val stock:  CreateOfferBasedOnProductActionReqStock ,
		@SerialName("sellingMode")  val sellingMode:  CreateOfferBasedOnProductActionReqSellingMode ,
		@SerialName("payments")  val payments:  CreateOfferBasedOnProductActionReqPayments ,
		@SerialName("delivery")  val delivery:  CreateOfferBasedOnProductActionReqDelivery ,
		@SerialName("publication")  val publication:  CreateOfferBasedOnProductActionReqPublication ,
		@SerialName("additionalMarketplaces")  @Contextual  val additionalMarketplaces: Any ,
		@SerialName("compatibilityList")  val compatibilityList:  CreateOfferBasedOnProductActionReqCompatibilityList ,
		@SerialName("images")  val images: List<String>  = emptyList(),
		@SerialName("description")  val description:  CreateOfferBasedOnProductActionReqDescription ,
		@SerialName("b2b")  val b2b:  CreateOfferBasedOnProductActionReqB2b ,
		@SerialName("attachments")  val attachments: List<CreateOfferBasedOnProductActionReqAttachments>  = emptyList(),
		@SerialName("fundraisingCampaign")  val fundraisingCampaign:  CreateOfferBasedOnProductActionReqFundraisingCampaign ,
		@SerialName("additionalServices")  val additionalServices:  CreateOfferBasedOnProductActionReqAdditionalServices ,
		@SerialName("afterSalesServices")  val afterSalesServices:  CreateOfferBasedOnProductActionReqAfterSalesServices ,
		@SerialName("sizeTable")  val sizeTable:  CreateOfferBasedOnProductActionReqSizeTable ,
		@SerialName("contact")  val contact:  CreateOfferBasedOnProductActionReqContact ,
		@SerialName("discounts")  val discounts:  CreateOfferBasedOnProductActionReqDiscounts ,
		@SerialName("location")  val location:  CreateOfferBasedOnProductActionReqLocation ,
		@SerialName("external")  val external:  CreateOfferBasedOnProductActionReqExternal ,
		@SerialName("taxSettings")  val taxSettings:  CreateOfferBasedOnProductActionReqTaxSettings ,
		@SerialName("messageToSellerSettings")  val messageToSellerSettings:  CreateOfferBasedOnProductActionReqMessageToSellerSettings ,
)
  // The base class definition for category
@Serializable
data class CreateOfferBasedOnProductActionReqCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for productSet
@Serializable
data class CreateOfferBasedOnProductActionReqProductSet (
		@SerialName("product")  val product:  CreateOfferBasedOnProductActionReqProductSetProduct ,
		@SerialName("quantity")  val quantity:  CreateOfferBasedOnProductActionReqProductSetQuantity ,
		@SerialName("responsiblePerson")  val responsiblePerson:  CreateOfferBasedOnProductActionReqProductSetResponsiblePerson ,
		@SerialName("responsibleProducer")  val responsibleProducer:  CreateOfferBasedOnProductActionReqProductSetResponsibleProducer ,
		@SerialName("safetyInformation")  val safetyInformation:  CreateOfferBasedOnProductActionReqProductSetSafetyInformation ,
		@SerialName("marketedBeforeGPSRObligation")  val marketedBeforeGPSRObligation: Boolean ,
		@SerialName("deposits")  val deposits: List<CreateOfferBasedOnProductActionReqProductSetDeposits>  = emptyList(),
)
  // The base class definition for product
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetProduct (
		@SerialName("id")  val id: String  = "",
		@SerialName("idType")  val idType: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("category")  val category:  CreateOfferBasedOnProductActionReqProductSetProductCategory ,
		@SerialName("parameters")  val parameters: List<CreateOfferBasedOnProductActionReqProductSetProductParameters>  = emptyList(),
		@SerialName("images")  val images: List<String>  = emptyList(),
)
  // The base class definition for category
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetProductCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for parameters
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetProductParameters (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("rangeValue")  val rangeValue:  CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue ,
		@SerialName("values")  val values: List<String>  = emptyList(),
		@SerialName("valuesIds")  val valuesIds: List<String>  = emptyList(),
)
  // The base class definition for rangeValue
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue (
		@SerialName("from")  val from: String  = "",
		@SerialName("to")  val to: String  = "",
)
  // The base class definition for quantity
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetQuantity (
		@SerialName("value")  val value: Int  = 0,
)
  // The base class definition for responsiblePerson
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetResponsiblePerson (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for responsibleProducer
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetResponsibleProducer (
		@SerialName("id")  val id: String  = "",
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetSafetyInformation (
		@SerialName("type")  val type: String  = "",
		@SerialName("description")  val description: String  = "",
)
  // The base class definition for deposits
@Serializable
data class CreateOfferBasedOnProductActionReqProductSetDeposits (
		@SerialName("id")  val id: String  = "",
		@SerialName("quantity")  val quantity: Int  = 0,
)
  // The base class definition for stock
@Serializable
data class CreateOfferBasedOnProductActionReqStock (
		@SerialName("available")  val available: Int  = 0,
		@SerialName("unit")  val unit: String  = "",
)
  // The base class definition for sellingMode
@Serializable
data class CreateOfferBasedOnProductActionReqSellingMode (
		@SerialName("format")  val format: String  = "",
		@SerialName("price")  val price:  CreateOfferBasedOnProductActionReqSellingModePrice ,
		@SerialName("minimalPrice")  val minimalPrice:  CreateOfferBasedOnProductActionReqSellingModeMinimalPrice ,
		@SerialName("startingPrice")  val startingPrice:  CreateOfferBasedOnProductActionReqSellingModeStartingPrice ,
)
  // The base class definition for price
@Serializable
data class CreateOfferBasedOnProductActionReqSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for minimalPrice
@Serializable
data class CreateOfferBasedOnProductActionReqSellingModeMinimalPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for startingPrice
@Serializable
data class CreateOfferBasedOnProductActionReqSellingModeStartingPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for payments
@Serializable
data class CreateOfferBasedOnProductActionReqPayments (
		@SerialName("invoice")  val invoice: String  = "",
)
  // The base class definition for delivery
@Serializable
data class CreateOfferBasedOnProductActionReqDelivery (
		@SerialName("handlingTime")  val handlingTime: String  = "",
		@SerialName("additionalInfo")  val additionalInfo: String  = "",
		@SerialName("shipmentDate")  val shipmentDate: String  = "",
		  // Optional; may be null
 @SerialName("shippingRates")  val shippingRates:  CreateOfferBasedOnProductActionReqDeliveryShippingRates ,
)
  // The base class definition for shippingRates
@Serializable
data class CreateOfferBasedOnProductActionReqDeliveryShippingRates (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for publication
@Serializable
data class CreateOfferBasedOnProductActionReqPublication (
		@SerialName("duration")  val duration: String  = "",
		@SerialName("startingAt")  val startingAt: String  = "",
		@SerialName("endingAt")  val endingAt: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("republish")  val republish: Boolean ,
)
  // The base class definition for compatibilityList
@Serializable
data class CreateOfferBasedOnProductActionReqCompatibilityList (
		@SerialName("items")  val items: List<CreateOfferBasedOnProductActionReqCompatibilityListItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class CreateOfferBasedOnProductActionReqCompatibilityListItems (
		@SerialName("type")  val type: String  = "",
		@SerialName("text")  val text: String  = "",
)
  // The base class definition for description
@Serializable
data class CreateOfferBasedOnProductActionReqDescription (
		@SerialName("sections")  val sections: List<CreateOfferBasedOnProductActionReqDescriptionSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class CreateOfferBasedOnProductActionReqDescriptionSections (
		@SerialName("items")  val items: List<CreateOfferBasedOnProductActionReqDescriptionSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class CreateOfferBasedOnProductActionReqDescriptionSectionsItems (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for b2b
@Serializable
data class CreateOfferBasedOnProductActionReqB2b (
		@SerialName("buyableOnlyByBusiness")  val buyableOnlyByBusiness: Boolean ,
)
  // The base class definition for attachments
@Serializable
data class CreateOfferBasedOnProductActionReqAttachments (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for fundraisingCampaign
@Serializable
data class CreateOfferBasedOnProductActionReqFundraisingCampaign (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for additionalServices
@Serializable
data class CreateOfferBasedOnProductActionReqAdditionalServices (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for afterSalesServices
@Serializable
data class CreateOfferBasedOnProductActionReqAfterSalesServices (
		@SerialName("impliedWarranty")  val impliedWarranty:  CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty ,
		@SerialName("returnPolicy")  val returnPolicy:  CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy ,
		@SerialName("warranty")  val warranty:  CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty ,
)
  // The base class definition for impliedWarranty
@Serializable
data class CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for returnPolicy
@Serializable
data class CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for warranty
@Serializable
data class CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for sizeTable
@Serializable
data class CreateOfferBasedOnProductActionReqSizeTable (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for contact
@Serializable
data class CreateOfferBasedOnProductActionReqContact (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for discounts
@Serializable
data class CreateOfferBasedOnProductActionReqDiscounts (
		@SerialName("wholesalePriceList")  val wholesalePriceList:  CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList ,
)
  // The base class definition for wholesalePriceList
@Serializable
data class CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for location
@Serializable
data class CreateOfferBasedOnProductActionReqLocation (
		@SerialName("city")  val city: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
		@SerialName("postCode")  val postCode: String  = "",
		@SerialName("province")  val province: String  = "",
)
  // The base class definition for external
@Serializable
data class CreateOfferBasedOnProductActionReqExternal (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for taxSettings
@Serializable
data class CreateOfferBasedOnProductActionReqTaxSettings (
		@SerialName("subject")  val subject: String  = "",
		@SerialName("exemption")  val exemption: String  = "",
		@SerialName("rates")  val rates: List<CreateOfferBasedOnProductActionReqTaxSettingsRates>  = emptyList(),
)
  // The base class definition for rates
@Serializable
data class CreateOfferBasedOnProductActionReqTaxSettingsRates (
		@SerialName("rate")  val rate: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
)
  // The base class definition for messageToSellerSettings
@Serializable
data class CreateOfferBasedOnProductActionReqMessageToSellerSettings (
		@SerialName("mode")  val mode: String  = "",
		@SerialName("hint")  val hint: String  = "",
)