package unknownpackage
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import emikot.MaybeField
import emikot.Maybe
import okhttp3.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import emikot.ClientContext
/**
 * Action to communicate with the action GetAllAvailableOfferPromotionPackagesAction
 */
data class GetAllAvailableOfferPromotionPackagesActionMeta(
    val name: String = "GetAllAvailableOfferPromotionPackagesAction",
    val url: String = "https://api.{environment}/sale/offer-promotion-packages",
    val method: String = "get"
)
/*data class GetAllAvailableOfferPromotionPackagesActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetAllAvailableOfferPromotionPackagesActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetAllAvailableOfferPromotionPackagesActionClient {
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
	): GetAllAvailableOfferPromotionPackagesActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetAllAvailableOfferPromotionPackagesActionMeta()
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
                GetAllAvailableOfferPromotionPackagesActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getAllAvailableOfferPromotionPackagesActionRes
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionRes (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackages")  val basePackages: List<GetAllAvailableOfferPromotionPackagesActionResBasePackages>  = emptyList(),
		@SerialName("extraPackages")  val extraPackages: List<GetAllAvailableOfferPromotionPackagesActionResExtraPackages>  = emptyList(),
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for basePackages
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionResBasePackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("cycleDuration")  val cycleDuration: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionResExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("cycleDuration")  val cycleDuration: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("basePackages")  val basePackages: List<GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages>  = emptyList(),
		@SerialName("extraPackages")  val extraPackages: List<GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages>  = emptyList(),
)
  // The base class definition for basePackages
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("cycleDuration")  val cycleDuration: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages (
		@SerialName("id")  val id: String  = "",
		@SerialName("name")  val name: String  = "",
		@SerialName("cycleDuration")  val cycleDuration: String  = "",
)