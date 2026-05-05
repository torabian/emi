package unknownpackage
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.ClientContext
import emikot.MaybeField
import emikot.Maybe
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
/**
 * Action to communicate with the action BatchOfferPromotionPackageModificationAction
 */
data class BatchOfferPromotionPackageModificationActionMeta(
    val name: String = "BatchOfferPromotionPackageModificationAction",
    val url: String = "https://api.{environment}/sale/offers/promo-options-commands/{commandId}",
    val method: String = "put"
)
/*data class BatchOfferPromotionPackageModificationActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class BatchOfferPromotionPackageModificationActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object BatchOfferPromotionPackageModificationActionClient {
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
	): BatchOfferPromotionPackageModificationActionResponse =
        withContext(Dispatchers.IO) {
            val meta = BatchOfferPromotionPackageModificationActionMeta()
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
                BatchOfferPromotionPackageModificationActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for batchOfferPromotionPackageModificationActionReq
@Serializable
data class BatchOfferPromotionPackageModificationActionReq (
		@SerialName("offerCriteria")  val offerCriteria: List<BatchOfferPromotionPackageModificationActionReqOfferCriteria>  = emptyList(),
		@SerialName("modification")  val modification:  BatchOfferPromotionPackageModificationActionReqModification ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for offerCriteria
@Serializable
data class BatchOfferPromotionPackageModificationActionReqOfferCriteria (
		@SerialName("offers")  val offers: List<BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers>  = emptyList(),
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for offers
@Serializable
data class BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for modification
@Serializable
data class BatchOfferPromotionPackageModificationActionReqModification (
		@SerialName("basePackage")  val basePackage:  BatchOfferPromotionPackageModificationActionReqModificationBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<BatchOfferPromotionPackageModificationActionReqModificationExtraPackages>  = emptyList(),
		@SerialName("modificationTime")  val modificationTime: String  = "",
)
  // The base class definition for basePackage
@Serializable
data class BatchOfferPromotionPackageModificationActionReqModificationBasePackage (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class BatchOfferPromotionPackageModificationActionReqModificationExtraPackages (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("modification")  val modification:  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification ,
)
  // The base class definition for modification
@Serializable
data class BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification (
		@SerialName("basePackage")  val basePackage:  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages>  = emptyList(),
		@SerialName("modificationTime")  val modificationTime: String  = "",
)
  // The base class definition for basePackage
@Serializable
data class BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for batchOfferPromotionPackageModificationActionRes
@Serializable
data class BatchOfferPromotionPackageModificationActionRes (
		@SerialName("id")  val id: String  = "",
		@SerialName("taskCount")  val taskCount:  BatchOfferPromotionPackageModificationActionResTaskCount ,
)
  // The base class definition for taskCount
@Serializable
data class BatchOfferPromotionPackageModificationActionResTaskCount (
		@SerialName("failed")  val failed: Int  = 0,
		@SerialName("success")  val success: Int  = 0,
		@SerialName("total")  val total: Int  = 0,
)