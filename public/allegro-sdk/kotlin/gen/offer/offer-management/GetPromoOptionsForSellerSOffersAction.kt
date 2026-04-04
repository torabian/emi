package unknownpackage
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import kotlinx.serialization.*
import emikot.MaybeField
import emikot.Maybe
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.json.*
/**
 * Action to communicate with the action GetPromoOptionsForSellerSOffersAction
 */
data class GetPromoOptionsForSellerSOffersActionMeta(
    val name: String = "GetPromoOptionsForSellerSOffersAction",
    val url: String = "https://api.{environment}/sale/offers/promo-options",
    val method: String = "get"
)
/*data class GetPromoOptionsForSellerSOffersActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetPromoOptionsForSellerSOffersActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetPromoOptionsForSellerSOffersActionClient {
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
	): GetPromoOptionsForSellerSOffersActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetPromoOptionsForSellerSOffersActionMeta()
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
                GetPromoOptionsForSellerSOffersActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getPromoOptionsForSellerSOffersActionRes
@Serializable
data class GetPromoOptionsForSellerSOffersActionRes (
		@SerialName("promoOptions")  val promoOptions: List<GetPromoOptionsForSellerSOffersActionResPromoOptions>  = emptyList(),
		@SerialName("count")  val count: Int  = 0,
		@SerialName("totalCount")  val totalCount: Int  = 0,
)
  // The base class definition for promoOptions
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptions (
		@SerialName("offerId")  val offerId: String  = "",
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for basePackage
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges (
		@SerialName("basePackage")  val basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackage")  val basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages>  = emptyList(),
		@SerialName("pendingChanges")  val pendingChanges:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges ,
)
  // The base class definition for basePackage
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)
  // The base class definition for pendingChanges
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges (
		@SerialName("basePackage")  val basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage ,
)
  // The base class definition for basePackage
@Serializable
data class GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage (
		@SerialName("id")  val id: String  = "",
		@SerialName("validFrom")  val validFrom: String  = "",
		@SerialName("validTo")  val validTo: String  = "",
		@SerialName("nextCycleDate")  val nextCycleDate: String  = "",
)