package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.withContext
import emikot.MaybeField
import emikot.Maybe
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import emikot.ClientContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
/**
 * Action to communicate with the action GetOfferPromotionPackagesAction
 */
data class GetOfferPromotionPackagesActionMeta(
    val name: String = "GetOfferPromotionPackagesAction",
    val url: String = "https://api.{environment}/sale/offers/{offerId}/promo-options",
    val method: String = "get"
)
/*data class GetOfferPromotionPackagesActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetOfferPromotionPackagesActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetOfferPromotionPackagesActionClient {
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
	): GetOfferPromotionPackagesActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetOfferPromotionPackagesActionMeta()
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
                GetOfferPromotionPackagesActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getOfferPromotionPackagesActionRes
@Serializable
data class GetOfferPromotionPackagesActionRes (
		@SerialName("offerId")  val offerId: String  = "",
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  GetOfferPromotionPackagesActionResBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<GetOfferPromotionPackagesActionResExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  GetOfferPromotionPackagesActionResPendingChanges ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<GetOfferPromotionPackagesActionResAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for basePackage
@Serializable
data class GetOfferPromotionPackagesActionResBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetOfferPromotionPackagesActionResExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class GetOfferPromotionPackagesActionResPendingChanges (
		@SerialName("basePackage")  val basePackage:  GetOfferPromotionPackagesActionResPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class GetOfferPromotionPackagesActionResPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class GetOfferPromotionPackagesActionResAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges ,
)
  // The base class definition for basePackage
@Serializable
data class GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges (
		@SerialName("basePackage")  val basePackage:  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)