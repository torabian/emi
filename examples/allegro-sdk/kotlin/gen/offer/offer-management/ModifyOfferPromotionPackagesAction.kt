package unknownpackage
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.ClientContext
import emikot.MaybeField
import emikot.Maybe
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
/**
 * Action to communicate with the action ModifyOfferPromotionPackagesAction
 */
data class ModifyOfferPromotionPackagesActionMeta(
    val name: String = "ModifyOfferPromotionPackagesAction",
    val url: String = "https://api.{environment}/sale/offers/{offerId}/promo-options-modification",
    val method: String = "post"
)
/*data class ModifyOfferPromotionPackagesActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ModifyOfferPromotionPackagesActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ModifyOfferPromotionPackagesActionClient {
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
	): ModifyOfferPromotionPackagesActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ModifyOfferPromotionPackagesActionMeta()
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
                ModifyOfferPromotionPackagesActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for modifyOfferPromotionPackagesActionReq
@Serializable
data class ModifyOfferPromotionPackagesActionReq (
		@SerialName("modifications")  val modifications: List<ModifyOfferPromotionPackagesActionReqModifications>  = emptyList(),
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for modifications
@Serializable
data class ModifyOfferPromotionPackagesActionReqModifications (
		@SerialName("modificationType")  val modificationType: String  = "",
		@SerialName("packageType")  val packageType: String  = "",
		@SerialName("packageId")  val packageId: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("modifications")  val modifications: List<ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications>  = emptyList(),
)
  // The base class definition for modifications
@Serializable
data class ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications (
		@SerialName("modificationType")  val modificationType: String  = "",
		@SerialName("packageType")  val packageType: String  = "",
		@SerialName("packageId")  val packageId: String  = "",
)
  // The base class definition for modifyOfferPromotionPackagesActionRes
@Serializable
data class ModifyOfferPromotionPackagesActionRes (
		@SerialName("offerId")  val offerId: String  = "",
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  ModifyOfferPromotionPackagesActionResBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<ModifyOfferPromotionPackagesActionResExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  ModifyOfferPromotionPackagesActionResPendingChanges ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<ModifyOfferPromotionPackagesActionResAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for basePackage
@Serializable
data class ModifyOfferPromotionPackagesActionResBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class ModifyOfferPromotionPackagesActionResExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class ModifyOfferPromotionPackagesActionResPendingChanges (
		@SerialName("basePackage")  val basePackage:  ModifyOfferPromotionPackagesActionResPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class ModifyOfferPromotionPackagesActionResPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class ModifyOfferPromotionPackagesActionResAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges ,
)
  // The base class definition for basePackage
@Serializable
data class ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges (
		@SerialName("basePackage")  val basePackage:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)