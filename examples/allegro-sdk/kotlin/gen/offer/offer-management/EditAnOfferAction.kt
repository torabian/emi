package unknownpackage
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.MaybeField
import emikot.Maybe
import okhttp3.*
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import okhttp3.MediaType.Companion.toMediaType
/**
 * Action to communicate with the action EditAnOfferAction
 */
data class EditAnOfferActionMeta(
    val name: String = "EditAnOfferAction",
    val url: String = "https://api.{environment}/sale/product-offers/{offerId}",
    val method: String = "patch"
)
/*data class EditAnOfferActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class EditAnOfferActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object EditAnOfferActionClient {
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
	): EditAnOfferActionResponse =
        withContext(Dispatchers.IO) {
            val meta = EditAnOfferActionMeta()
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
                EditAnOfferActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for editAnOfferActionReq
@Serializable
data class EditAnOfferActionReq (
		@SerialName("name")  val name: String  = "",
		@SerialName("language")  val language: String  = "",
		@SerialName("category")  val category:  EditAnOfferActionReqCategory ,
		@SerialName("productSet")  val productSet: List<EditAnOfferActionReqProductSet>  = emptyList(),
		@SerialName("stock")  val stock:  EditAnOfferActionReqStock ,
		@SerialName("sellingMode")  val sellingMode:  EditAnOfferActionReqSellingMode ,
		@SerialName("payments")  val payments:  EditAnOfferActionReqPayments ,
		@SerialName("delivery")  val delivery:  EditAnOfferActionReqDelivery ,
		@SerialName("publication")  val publication:  EditAnOfferActionReqPublication ,
		@SerialName("additionalMarketplaces")  @Contextual  val additionalMarketplaces: Any ,
		@SerialName("compatibilityList")  val compatibilityList:  EditAnOfferActionReqCompatibilityList ,
		@SerialName("images")  val images: List<String>  = emptyList(),
		@SerialName("description")  val description:  EditAnOfferActionReqDescription ,
		@SerialName("b2b")  val b2b:  EditAnOfferActionReqB2b ,
		@SerialName("attachments")  val attachments: List<EditAnOfferActionReqAttachments>  = emptyList(),
		@SerialName("fundraisingCampaign")  val fundraisingCampaign:  EditAnOfferActionReqFundraisingCampaign ,
		@SerialName("additionalServices")  val additionalServices:  EditAnOfferActionReqAdditionalServices ,
		@SerialName("afterSalesServices")  val afterSalesServices:  EditAnOfferActionReqAfterSalesServices ,
		@SerialName("sizeTable")  val sizeTable:  EditAnOfferActionReqSizeTable ,
		@SerialName("contact")  val contact:  EditAnOfferActionReqContact ,
		@SerialName("discounts")  val discounts:  EditAnOfferActionReqDiscounts ,
		@SerialName("location")  val location:  EditAnOfferActionReqLocation ,
		@SerialName("external")  val external:  EditAnOfferActionReqExternal ,
		@SerialName("taxSettings")  val taxSettings:  EditAnOfferActionReqTaxSettings ,
		@SerialName("messageToSellerSettings")  val messageToSellerSettings:  EditAnOfferActionReqMessageToSellerSettings ,
)
  // The base class definition for category
@Serializable
data class EditAnOfferActionReqCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for productSet
@Serializable
data class EditAnOfferActionReqProductSet (
		@SerialName("product")  val product:  EditAnOfferActionReqProductSetProduct ,
		@SerialName("quantity")  val quantity:  EditAnOfferActionReqProductSetQuantity ,
		@SerialName("responsiblePerson")  val responsiblePerson:  EditAnOfferActionReqProductSetResponsiblePerson ,
		@SerialName("responsibleProducer")  val responsibleProducer:  EditAnOfferActionReqProductSetResponsibleProducer ,
		@SerialName("safetyInformation")  val safetyInformation:  EditAnOfferActionReqProductSetSafetyInformation ,
		@SerialName("marketedBeforeGPSRObligation")  val marketedBeforeGPSRObligation: Boolean ,
		@SerialName("deposits")  val deposits: List<EditAnOfferActionReqProductSetDeposits>  = emptyList(),
)
  // The base class definition for product
@Serializable
data class EditAnOfferActionReqProductSetProduct (
		@SerialName("id")  val id: String  = "",
		@SerialName("idType")  val idType: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("category")  val category:  EditAnOfferActionReqProductSetProductCategory ,
		@SerialName("parameters")  val parameters: List<EditAnOfferActionReqProductSetProductParameters>  = emptyList(),
		@SerialName("images")  val images: List<String>  = emptyList(),
)
  // The base class definition for category
@Serializable
data class EditAnOfferActionReqProductSetProductCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for parameters
@Serializable
data class EditAnOfferActionReqProductSetProductParameters (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("rangeValue")  val rangeValue:  EditAnOfferActionReqProductSetProductParametersRangeValue ,
		@SerialName("values")  val values: List<String>  = emptyList(),
		@SerialName("valuesIds")  val valuesIds: List<String>  = emptyList(),
)
  // The base class definition for rangeValue
@Serializable
data class EditAnOfferActionReqProductSetProductParametersRangeValue (
		@SerialName("from")  val from: String  = "",
		@SerialName("to")  val to: String  = "",
)
  // The base class definition for quantity
@Serializable
data class EditAnOfferActionReqProductSetQuantity (
		@SerialName("value")  val value: Int  = 0,
)
  // The base class definition for responsiblePerson
@Serializable
data class EditAnOfferActionReqProductSetResponsiblePerson (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for responsibleProducer
@Serializable
data class EditAnOfferActionReqProductSetResponsibleProducer (
		@SerialName("id")  val id: String  = "",
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class EditAnOfferActionReqProductSetSafetyInformation (
		@SerialName("type")  val type: String  = "",
		@SerialName("description")  val description: String  = "",
)
  // The base class definition for deposits
@Serializable
data class EditAnOfferActionReqProductSetDeposits (
		@SerialName("id")  val id: String  = "",
		@SerialName("quantity")  val quantity: Int  = 0,
)
  // The base class definition for stock
@Serializable
data class EditAnOfferActionReqStock (
		@SerialName("available")  val available: Int  = 0,
		@SerialName("unit")  val unit: String  = "",
)
  // The base class definition for sellingMode
@Serializable
data class EditAnOfferActionReqSellingMode (
		@SerialName("format")  val format: String  = "",
		@SerialName("price")  val price:  EditAnOfferActionReqSellingModePrice ,
		@SerialName("minimalPrice")  val minimalPrice:  EditAnOfferActionReqSellingModeMinimalPrice ,
		@SerialName("startingPrice")  val startingPrice:  EditAnOfferActionReqSellingModeStartingPrice ,
)
  // The base class definition for price
@Serializable
data class EditAnOfferActionReqSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for minimalPrice
@Serializable
data class EditAnOfferActionReqSellingModeMinimalPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for startingPrice
@Serializable
data class EditAnOfferActionReqSellingModeStartingPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for payments
@Serializable
data class EditAnOfferActionReqPayments (
		@SerialName("invoice")  val invoice: String  = "",
)
  // The base class definition for delivery
@Serializable
data class EditAnOfferActionReqDelivery (
		@SerialName("handlingTime")  val handlingTime: String  = "",
		@SerialName("additionalInfo")  val additionalInfo: String  = "",
		@SerialName("shipmentDate")  val shipmentDate: String  = "",
		@SerialName("shippingRates")  val shippingRates:  EditAnOfferActionReqDeliveryShippingRates ,
)
  // The base class definition for shippingRates
@Serializable
data class EditAnOfferActionReqDeliveryShippingRates (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for publication
@Serializable
data class EditAnOfferActionReqPublication (
		@SerialName("duration")  val duration: String  = "",
		@SerialName("startingAt")  val startingAt: String  = "",
		@SerialName("endingAt")  val endingAt: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("republish")  val republish: Boolean ,
)
  // The base class definition for compatibilityList
@Serializable
data class EditAnOfferActionReqCompatibilityList (
		@SerialName("items")  val items: List<EditAnOfferActionReqCompatibilityListItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class EditAnOfferActionReqCompatibilityListItems (
		@SerialName("type")  val type: String  = "",
		@SerialName("text")  val text: String  = "",
)
  // The base class definition for description
@Serializable
data class EditAnOfferActionReqDescription (
		@SerialName("sections")  val sections: List<EditAnOfferActionReqDescriptionSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class EditAnOfferActionReqDescriptionSections (
		@SerialName("items")  val items: List<EditAnOfferActionReqDescriptionSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class EditAnOfferActionReqDescriptionSectionsItems (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for b2b
@Serializable
data class EditAnOfferActionReqB2b (
		@SerialName("buyableOnlyByBusiness")  val buyableOnlyByBusiness: Boolean ,
)
  // The base class definition for attachments
@Serializable
data class EditAnOfferActionReqAttachments (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for fundraisingCampaign
@Serializable
data class EditAnOfferActionReqFundraisingCampaign (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for additionalServices
@Serializable
data class EditAnOfferActionReqAdditionalServices (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for afterSalesServices
@Serializable
data class EditAnOfferActionReqAfterSalesServices (
		@SerialName("impliedWarranty")  val impliedWarranty:  EditAnOfferActionReqAfterSalesServicesImpliedWarranty ,
		@SerialName("returnPolicy")  val returnPolicy:  EditAnOfferActionReqAfterSalesServicesReturnPolicy ,
		@SerialName("warranty")  val warranty:  EditAnOfferActionReqAfterSalesServicesWarranty ,
)
  // The base class definition for impliedWarranty
@Serializable
data class EditAnOfferActionReqAfterSalesServicesImpliedWarranty (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for returnPolicy
@Serializable
data class EditAnOfferActionReqAfterSalesServicesReturnPolicy (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for warranty
@Serializable
data class EditAnOfferActionReqAfterSalesServicesWarranty (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for sizeTable
@Serializable
data class EditAnOfferActionReqSizeTable (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for contact
@Serializable
data class EditAnOfferActionReqContact (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for discounts
@Serializable
data class EditAnOfferActionReqDiscounts (
		@SerialName("wholesalePriceList")  val wholesalePriceList:  EditAnOfferActionReqDiscountsWholesalePriceList ,
)
  // The base class definition for wholesalePriceList
@Serializable
data class EditAnOfferActionReqDiscountsWholesalePriceList (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
)
  // The base class definition for location
@Serializable
data class EditAnOfferActionReqLocation (
		@SerialName("city")  val city: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
		@SerialName("postCode")  val postCode: String  = "",
		@SerialName("province")  val province: String  = "",
)
  // The base class definition for external
@Serializable
data class EditAnOfferActionReqExternal (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for taxSettings
@Serializable
data class EditAnOfferActionReqTaxSettings (
		@SerialName("subject")  val subject: String  = "",
		@SerialName("exemption")  val exemption: String  = "",
		@SerialName("rates")  val rates: List<EditAnOfferActionReqTaxSettingsRates>  = emptyList(),
)
  // The base class definition for rates
@Serializable
data class EditAnOfferActionReqTaxSettingsRates (
		@SerialName("rate")  val rate: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
)
  // The base class definition for messageToSellerSettings
@Serializable
data class EditAnOfferActionReqMessageToSellerSettings (
		@SerialName("mode")  val mode: String  = "",
		@SerialName("hint")  val hint: String  = "",
)
  // The base class definition for editAnOfferActionRes
@Serializable
data class EditAnOfferActionRes (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("language")  val language: String  = "",
		@SerialName("category")  val category:  EditAnOfferActionResCategory ,
		@SerialName("productSet")  val productSet: List<EditAnOfferActionResProductSet>  = emptyList(),
		@SerialName("stock")  val stock:  EditAnOfferActionResStock ,
		@SerialName("payments")  val payments:  EditAnOfferActionResPayments ,
		@SerialName("sellingMode")  val sellingMode:  EditAnOfferActionResSellingMode ,
		@SerialName("delivery")  val delivery:  EditAnOfferActionResDelivery ,
		@SerialName("publication")  val publication:  EditAnOfferActionResPublication ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces:  EditAnOfferActionResAdditionalMarketplaces ,
		@SerialName("b2b")  val b2b:  EditAnOfferActionResB2b ,
		@SerialName("compatibilityList")  val compatibilityList:  EditAnOfferActionResCompatibilityList ,
		@SerialName("validation")  val validation:  EditAnOfferActionResValidation ,
		@SerialName("warnings")  val warnings: List<String>  = emptyList(),
		@SerialName("afterSalesServices")  val afterSalesServices:  EditAnOfferActionResAfterSalesServices ,
		@SerialName("discounts")  val discounts:  EditAnOfferActionResDiscounts ,
		@SerialName("contact")  val contact:  EditAnOfferActionResContact ,
		@SerialName("attachments")  val attachments: List<EditAnOfferActionResAttachments>  = emptyList(),
		@SerialName("fundraisingCampaign")  val fundraisingCampaign:  EditAnOfferActionResFundraisingCampaign ,
		@SerialName("additionalServices")  val additionalServices:  EditAnOfferActionResAdditionalServices ,
		@SerialName("sizeTable")  val sizeTable:  EditAnOfferActionResSizeTable ,
		@SerialName("location")  val location:  EditAnOfferActionResLocation ,
		@SerialName("external")  val external:  EditAnOfferActionResExternal ,
		@SerialName("taxSettings")  val taxSettings:  EditAnOfferActionResTaxSettings ,
		@SerialName("messageToSellerSettings")  val messageToSellerSettings:  EditAnOfferActionResMessageToSellerSettings ,
		@SerialName("createdAt")  val createdAt: String  = "",
		@SerialName("updatedAt")  val updatedAt: String  = "",
		@SerialName("images")  val images: List<String>  = emptyList(),
		@SerialName("description")  val description:  EditAnOfferActionResDescription ,
)
  // The base class definition for category
@Serializable
data class EditAnOfferActionResCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for productSet
@Serializable
data class EditAnOfferActionResProductSet (
		@SerialName("quantity")  val quantity:  EditAnOfferActionResProductSetQuantity ,
		@SerialName("product")  val product:  EditAnOfferActionResProductSetProduct ,
		@SerialName("responsiblePerson")  val responsiblePerson:  EditAnOfferActionResProductSetResponsiblePerson ,
		@SerialName("responsibleProducer")  val responsibleProducer:  EditAnOfferActionResProductSetResponsibleProducer ,
		@SerialName("safetyInformation")  val safetyInformation:  EditAnOfferActionResProductSetSafetyInformation ,
		@SerialName("marketedBeforeGPSRObligation")  val marketedBeforeGPSRObligation: Boolean ,
		@SerialName("deposits")  val deposits: List<EditAnOfferActionResProductSetDeposits>  = emptyList(),
)
  // The base class definition for quantity
@Serializable
data class EditAnOfferActionResProductSetQuantity (
		@SerialName("value")  val value: Int  = 0,
)
  // The base class definition for product
@Serializable
data class EditAnOfferActionResProductSetProduct (
		@SerialName("id")  val id: String  = "",
		@SerialName("isAiCoCreated")  val isAiCoCreated: Boolean ,
		@SerialName("publication")  val publication:  EditAnOfferActionResProductSetProductPublication ,
		@SerialName("parameters")  val parameters: List<EditAnOfferActionResProductSetProductParameters>  = emptyList(),
)
  // The base class definition for publication
@Serializable
data class EditAnOfferActionResProductSetProductPublication (
		@SerialName("status")  val status: String  = "",
)
  // The base class definition for parameters
@Serializable
data class EditAnOfferActionResProductSetProductParameters (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("rangeValue")  val rangeValue:  EditAnOfferActionResProductSetProductParametersRangeValue ,
		@SerialName("values")  val values: List<String>  = emptyList(),
		@SerialName("valuesIds")  val valuesIds: List<String>  = emptyList(),
)
  // The base class definition for rangeValue
@Serializable
data class EditAnOfferActionResProductSetProductParametersRangeValue (
		@SerialName("from")  val from: String  = "",
		@SerialName("to")  val to: String  = "",
)
  // The base class definition for responsiblePerson
@Serializable
data class EditAnOfferActionResProductSetResponsiblePerson (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for responsibleProducer
@Serializable
data class EditAnOfferActionResProductSetResponsibleProducer (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class EditAnOfferActionResProductSetSafetyInformation (
		@SerialName("type")  val type: String  = "",
		@SerialName("description")  val description: String  = "",
)
  // The base class definition for deposits
@Serializable
data class EditAnOfferActionResProductSetDeposits (
		@SerialName("id")  val id: String  = "",
		@SerialName("quantity")  val quantity: Int  = 0,
)
  // The base class definition for stock
@Serializable
data class EditAnOfferActionResStock (
		@SerialName("available")  val available: Int  = 0,
		@SerialName("unit")  val unit: String  = "",
)
  // The base class definition for payments
@Serializable
data class EditAnOfferActionResPayments (
		@SerialName("invoice")  val invoice: String  = "",
)
  // The base class definition for sellingMode
@Serializable
data class EditAnOfferActionResSellingMode (
		@SerialName("format")  val format: String  = "",
		@SerialName("price")  val price:  EditAnOfferActionResSellingModePrice ,
		@SerialName("minimalPrice")  val minimalPrice:  EditAnOfferActionResSellingModeMinimalPrice ,
		@SerialName("startingPrice")  val startingPrice:  EditAnOfferActionResSellingModeStartingPrice ,
)
  // The base class definition for price
@Serializable
data class EditAnOfferActionResSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for minimalPrice
@Serializable
data class EditAnOfferActionResSellingModeMinimalPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for startingPrice
@Serializable
data class EditAnOfferActionResSellingModeStartingPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for delivery
@Serializable
data class EditAnOfferActionResDelivery (
		@SerialName("handlingTime")  val handlingTime: String  = "",
		@SerialName("additionalInfo")  val additionalInfo: String  = "",
		@SerialName("shipmentDate")  val shipmentDate: String  = "",
		@SerialName("shippingRates")  val shippingRates:  EditAnOfferActionResDeliveryShippingRates ,
)
  // The base class definition for shippingRates
@Serializable
data class EditAnOfferActionResDeliveryShippingRates (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for publication
@Serializable
data class EditAnOfferActionResPublication (
		@SerialName("duration")  val duration: String  = "",
		@SerialName("startingAt")  val startingAt: String  = "",
		@SerialName("endingAt")  val endingAt: String  = "",
		@SerialName("endedBy")  val endedBy: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("republish")  val republish: Boolean ,
		@SerialName("marketplaces")  val marketplaces:  EditAnOfferActionResPublicationMarketplaces ,
)
  // The base class definition for marketplaces
@Serializable
data class EditAnOfferActionResPublicationMarketplaces (
		@SerialName("base")  val base:  EditAnOfferActionResPublicationMarketplacesBase ,
		@SerialName("additional")  val additional: List<EditAnOfferActionResPublicationMarketplacesAdditional>  = emptyList(),
)
  // The base class definition for base
@Serializable
data class EditAnOfferActionResPublicationMarketplacesBase (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additional
@Serializable
data class EditAnOfferActionResPublicationMarketplacesAdditional (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class EditAnOfferActionResAdditionalMarketplaces (
		@SerialName("sellingMode")  val sellingMode:  EditAnOfferActionResAdditionalMarketplacesSellingMode ,
		@SerialName("publication")  val publication:  EditAnOfferActionResAdditionalMarketplacesPublication ,
)
  // The base class definition for sellingMode
@Serializable
data class EditAnOfferActionResAdditionalMarketplacesSellingMode (
		@SerialName("price")  val price:  EditAnOfferActionResAdditionalMarketplacesSellingModePrice ,
)
  // The base class definition for price
@Serializable
data class EditAnOfferActionResAdditionalMarketplacesSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for publication
@Serializable
data class EditAnOfferActionResAdditionalMarketplacesPublication (
		@SerialName("state")  val state: String  = "",
		@SerialName("refusalReasons")  val refusalReasons: List<EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons>  = emptyList(),
)
  // The base class definition for refusalReasons
@Serializable
data class EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons (
		@SerialName("code")  val code: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("parameters")  val parameters:  EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters ,
)
  // The base class definition for parameters
@Serializable
data class EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters (
		@SerialName("maxAllowedPriceDecreasePercent")  val maxAllowedPriceDecreasePercent: List<String>  = emptyList(),
)
  // The base class definition for b2b
@Serializable
data class EditAnOfferActionResB2b (
		@SerialName("buyableOnlyByBusiness")  val buyableOnlyByBusiness: Boolean ,
)
  // The base class definition for compatibilityList
@Serializable
data class EditAnOfferActionResCompatibilityList (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for validation
@Serializable
data class EditAnOfferActionResValidation (
		@SerialName("validatedAt")  val validatedAt: String  = "",
		@SerialName("errors")  val errors: List<EditAnOfferActionResValidationErrors>  = emptyList(),
		@SerialName("warnings")  val warnings: List<EditAnOfferActionResValidationWarnings>  = emptyList(),
)
  // The base class definition for errors
@Serializable
data class EditAnOfferActionResValidationErrors (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  EditAnOfferActionResValidationErrorsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class EditAnOfferActionResValidationErrorsMetadata (
		@SerialName("productId")  val productId: String  = "",
)
  // The base class definition for warnings
@Serializable
data class EditAnOfferActionResValidationWarnings (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  EditAnOfferActionResValidationWarningsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class EditAnOfferActionResValidationWarningsMetadata (
		@SerialName("productId")  val productId: String  = "",
)
  // The base class definition for afterSalesServices
@Serializable
data class EditAnOfferActionResAfterSalesServices (
		@SerialName("impliedWarranty")  val impliedWarranty:  EditAnOfferActionResAfterSalesServicesImpliedWarranty ,
		@SerialName("returnPolicy")  val returnPolicy:  EditAnOfferActionResAfterSalesServicesReturnPolicy ,
		@SerialName("warranty")  val warranty:  EditAnOfferActionResAfterSalesServicesWarranty ,
)
  // The base class definition for impliedWarranty
@Serializable
data class EditAnOfferActionResAfterSalesServicesImpliedWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for returnPolicy
@Serializable
data class EditAnOfferActionResAfterSalesServicesReturnPolicy (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for warranty
@Serializable
data class EditAnOfferActionResAfterSalesServicesWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for discounts
@Serializable
data class EditAnOfferActionResDiscounts (
		@SerialName("wholesalePriceList")  val wholesalePriceList:  EditAnOfferActionResDiscountsWholesalePriceList ,
)
  // The base class definition for wholesalePriceList
@Serializable
data class EditAnOfferActionResDiscountsWholesalePriceList (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for contact
@Serializable
data class EditAnOfferActionResContact (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for attachments
@Serializable
data class EditAnOfferActionResAttachments (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for fundraisingCampaign
@Serializable
data class EditAnOfferActionResFundraisingCampaign (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalServices
@Serializable
data class EditAnOfferActionResAdditionalServices (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for sizeTable
@Serializable
data class EditAnOfferActionResSizeTable (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for location
@Serializable
data class EditAnOfferActionResLocation (
		@SerialName("city")  val city: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
		@SerialName("postCode")  val postCode: String  = "",
		@SerialName("province")  val province: String  = "",
)
  // The base class definition for external
@Serializable
data class EditAnOfferActionResExternal (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for taxSettings
@Serializable
data class EditAnOfferActionResTaxSettings (
		@SerialName("subject")  val subject: String  = "",
		@SerialName("exemption")  val exemption: String  = "",
		@SerialName("rates")  val rates: List<EditAnOfferActionResTaxSettingsRates>  = emptyList(),
)
  // The base class definition for rates
@Serializable
data class EditAnOfferActionResTaxSettingsRates (
		@SerialName("rate")  val rate: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
)
  // The base class definition for messageToSellerSettings
@Serializable
data class EditAnOfferActionResMessageToSellerSettings (
		@SerialName("mode")  val mode: String  = "",
		@SerialName("hint")  val hint: String  = "",
)
  // The base class definition for description
@Serializable
data class EditAnOfferActionResDescription (
		@SerialName("sections")  val sections: List<EditAnOfferActionResDescriptionSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class EditAnOfferActionResDescriptionSections (
		@SerialName("items")  val items: List<EditAnOfferActionResDescriptionSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class EditAnOfferActionResDescriptionSectionsItems (
		@SerialName("type")  val type: String  = "",
)