package unknownpackage
import emikot.ClientContext
import emikot.MaybeField
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.Maybe
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
/**
 * Action to communicate with the action GetAllDataOfTheParticularProductOfferAction
 */
data class GetAllDataOfTheParticularProductOfferActionMeta(
    val name: String = "GetAllDataOfTheParticularProductOfferAction",
    val url: String = "https://api.{environment}/sale/product-offers/{offerId}",
    val method: String = "get"
)
/*data class GetAllDataOfTheParticularProductOfferActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetAllDataOfTheParticularProductOfferActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetAllDataOfTheParticularProductOfferActionClient {
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
	): GetAllDataOfTheParticularProductOfferActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetAllDataOfTheParticularProductOfferActionMeta()
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
                GetAllDataOfTheParticularProductOfferActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getAllDataOfTheParticularProductOfferActionRes
@Serializable
data class GetAllDataOfTheParticularProductOfferActionRes (
		  // Unique offer identifier
 @SerialName("id")  val id: String  = "",
		  // Offer title
 @SerialName("name")  val name: String  = "",
		  // Offer language code (e.g. pl-PL)
 @SerialName("language")  val language: String  = "",
		  // Offer creation timestamp (ISO8601)
 @SerialName("createdAt")  val createdAt: String  = "",
		  // Offer last update timestamp (ISO8601)
 @SerialName("updatedAt")  val updatedAt: String  = "",
		@SerialName("category")  val category:  GetAllDataOfTheParticularProductOfferActionResCategory ,
		@SerialName("stock")  val stock:  GetAllDataOfTheParticularProductOfferActionResStock ,
		@SerialName("contact")  val contact:  GetAllDataOfTheParticularProductOfferActionResContact ,
		@SerialName("publication")  val publication:  GetAllDataOfTheParticularProductOfferActionResPublication ,
		@SerialName("sellingMode")  val sellingMode:  GetAllDataOfTheParticularProductOfferActionResSellingMode ,
		@SerialName("payments")  val payments:  GetAllDataOfTheParticularProductOfferActionResPayments ,
		@SerialName("delivery")  val delivery:  GetAllDataOfTheParticularProductOfferActionResDelivery ,
		@SerialName("afterSalesServices")  val afterSalesServices:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServices ,
		@SerialName("discounts")  val discounts:  GetAllDataOfTheParticularProductOfferActionResDiscounts ,
		@SerialName("description")  val description:  GetAllDataOfTheParticularProductOfferActionResDescription ,
		@SerialName("images")  val images: List<String>  = emptyList(),
		@SerialName("productSet")  val productSet: List<GetAllDataOfTheParticularProductOfferActionResProductSet>  = emptyList(),
		@SerialName("attachments")  val attachments: List<GetAllDataOfTheParticularProductOfferActionResAttachments>  = emptyList(),
		@SerialName("fundraisingCampaign")  val fundraisingCampaign:  GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign ,
		@SerialName("additionalServices")  val additionalServices:  GetAllDataOfTheParticularProductOfferActionResAdditionalServices ,
		@SerialName("additionalMarketplaces")  @Contextual  val additionalMarketplaces: Any ,
		@SerialName("b2b")  val b2b:  GetAllDataOfTheParticularProductOfferActionResB2b ,
		@SerialName("compatibilityList")  val compatibilityList:  GetAllDataOfTheParticularProductOfferActionResCompatibilityList ,
		@SerialName("validation")  val validation:  GetAllDataOfTheParticularProductOfferActionResValidation ,
		@SerialName("external")  val external:  GetAllDataOfTheParticularProductOfferActionResExternal ,
		@SerialName("sizeTable")  val sizeTable:  GetAllDataOfTheParticularProductOfferActionResSizeTable ,
		@SerialName("taxSettings")  val taxSettings:  GetAllDataOfTheParticularProductOfferActionResTaxSettings ,
		@SerialName("messageToSellerSettings")  val messageToSellerSettings:  GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings ,
)
  // The base class definition for category
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResCategory (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for stock
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResStock (
		@SerialName("available")  val available: Int  = 0,
		@SerialName("unit")  val unit: String  = "",
)
  // The base class definition for contact
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResContact (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for publication
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResPublication (
		@SerialName("duration")  val duration: String  = "",
		@SerialName("startingAt")  val startingAt: String  = "",
		@SerialName("endingAt")  val endingAt: String  = "",
		@SerialName("endedBy")  val endedBy: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("republish")  val republish: Boolean ,
		@SerialName("marketplaces")  val marketplaces:  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces ,
)
  // The base class definition for marketplaces
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces (
		@SerialName("base")  val base:  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase ,
		@SerialName("additional")  val additional: List<GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional>  = emptyList(),
)
  // The base class definition for base
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additional
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for sellingMode
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResSellingMode (
		@SerialName("format")  val format: String  = "",
		@SerialName("price")  val price:  GetAllDataOfTheParticularProductOfferActionResSellingModePrice ,
		@SerialName("minimalPrice")  val minimalPrice:  GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice ,
		@SerialName("startingPrice")  val startingPrice:  GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice ,
)
  // The base class definition for price
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResSellingModePrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for minimalPrice
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for startingPrice
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice (
		@SerialName("amount")  val amount: String  = "",
		@SerialName("currency")  val currency: String  = "",
)
  // The base class definition for payments
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResPayments (
		@SerialName("invoice")  val invoice: String  = "",
)
  // The base class definition for delivery
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDelivery (
		@SerialName("handlingTime")  val handlingTime: String  = "",
		@SerialName("additionalInfo")  val additionalInfo: String  = "",
		@SerialName("shipmentDate")  val shipmentDate: String  = "",
		@SerialName("shippingRates")  val shippingRates:  GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates ,
)
  // The base class definition for shippingRates
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for afterSalesServices
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAfterSalesServices (
		@SerialName("impliedWarranty")  val impliedWarranty:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty ,
		@SerialName("returnPolicy")  val returnPolicy:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy ,
		@SerialName("warranty")  val warranty:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty ,
)
  // The base class definition for impliedWarranty
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for returnPolicy
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for warranty
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for discounts
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDiscounts (
		@SerialName("wholesalePriceList")  val wholesalePriceList:  GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList ,
)
  // The base class definition for wholesalePriceList
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for description
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDescription (
		@SerialName("sections")  val sections: List<GetAllDataOfTheParticularProductOfferActionResDescriptionSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDescriptionSections (
		@SerialName("items")  val items: List<GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for productSet
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSet (
		@SerialName("quantity")  val quantity:  GetAllDataOfTheParticularProductOfferActionResProductSetQuantity ,
		@SerialName("product")  val product:  GetAllDataOfTheParticularProductOfferActionResProductSetProduct ,
		@SerialName("responsiblePerson")  val responsiblePerson:  GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson ,
		@SerialName("responsibleProducer")  val responsibleProducer:  GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer ,
		@SerialName("safetyInformation")  val safetyInformation:  GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation ,
		@SerialName("marketedBeforeGPSRObligation")  val marketedBeforeGPSRObligation: Boolean ,
		@SerialName("deposits")  val deposits: List<GetAllDataOfTheParticularProductOfferActionResProductSetDeposits>  = emptyList(),
)
  // The base class definition for quantity
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetQuantity (
		@SerialName("value")  val value: Int  = 0,
)
  // The base class definition for product
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProduct (
		@SerialName("id")  val id: String  = "",
		@SerialName("isAiCoCreated")  val isAiCoCreated: Boolean ,
		@SerialName("publication")  val publication:  GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication ,
		@SerialName("parameters")  val parameters: List<GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters>  = emptyList(),
)
  // The base class definition for publication
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication (
		@SerialName("status")  val status: String  = "",
)
  // The base class definition for parameters
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("rangeValue")  val rangeValue:  GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue ,
		@SerialName("values")  val values: List<GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues>  = emptyList(),
		@SerialName("valuesIds")  val valuesIds: List<GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds>  = emptyList(),
)
  // The base class definition for rangeValue
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue (
		@SerialName("from")  val from: String  = "",
		@SerialName("to")  val to: String  = "",
)
  // The base class definition for values
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues (
)
  // The base class definition for valuesIds
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds (
)
  // The base class definition for responsiblePerson
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for responsibleProducer
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation (
		@SerialName("type")  val type: String  = "",
		@SerialName("description")  val description: String  = "",
)
  // The base class definition for deposits
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResProductSetDeposits (
		@SerialName("id")  val id: String  = "",
		@SerialName("quantity")  val quantity: Int  = 0,
)
  // The base class definition for attachments
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAttachments (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for fundraisingCampaign
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalServices
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResAdditionalServices (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for b2b
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResB2b (
		@SerialName("buyableOnlyByBusiness")  val buyableOnlyByBusiness: Boolean ,
)
  // The base class definition for compatibilityList
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResCompatibilityList (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for validation
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResValidation (
		@SerialName("validatedAt")  val validatedAt: String  = "",
		@SerialName("errors")  val errors: List<GetAllDataOfTheParticularProductOfferActionResValidationErrors>  = emptyList(),
		@SerialName("warnings")  val warnings: List<GetAllDataOfTheParticularProductOfferActionResValidationWarnings>  = emptyList(),
)
  // The base class definition for errors
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResValidationErrors (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata (
		@SerialName("productId")  val productId: String  = "",
)
  // The base class definition for warnings
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResValidationWarnings (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata (
		@SerialName("productId")  val productId: String  = "",
)
  // The base class definition for external
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResExternal (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for sizeTable
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResSizeTable (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for taxSettings
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResTaxSettings (
		@SerialName("subject")  val subject: String  = "",
		@SerialName("exemption")  val exemption: String  = "",
		@SerialName("rates")  val rates: List<GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates>  = emptyList(),
)
  // The base class definition for rates
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates (
		@SerialName("rate")  val rate: String  = "",
		@SerialName("countryCode")  val countryCode: String  = "",
)
  // The base class definition for messageToSellerSettings
@Serializable
data class GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings (
		@SerialName("mode")  val mode: String  = "",
		@SerialName("hint")  val hint: String  = "",
)