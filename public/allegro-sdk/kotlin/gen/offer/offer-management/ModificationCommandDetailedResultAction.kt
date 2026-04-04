package unknownpackage
import emikot.ClientContext
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.serialization.*
import emikot.MaybeField
import emikot.Maybe
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
/**
 * Action to communicate with the action ModificationCommandDetailedResultAction
 */
data class ModificationCommandDetailedResultActionMeta(
    val name: String = "ModificationCommandDetailedResultAction",
    val url: String = "https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks",
    val method: String = "get"
)
/*data class ModificationCommandDetailedResultActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ModificationCommandDetailedResultActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ModificationCommandDetailedResultActionClient {
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
	): ModificationCommandDetailedResultActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ModificationCommandDetailedResultActionMeta()
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
                ModificationCommandDetailedResultActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for modificationCommandDetailedResultActionRes
@Serializable
data class ModificationCommandDetailedResultActionRes (
		@SerialName("tasks")  val tasks: List<ModificationCommandDetailedResultActionResTasks>  = emptyList(),
		@SerialName("modification")  val modification:  ModificationCommandDetailedResultActionResModification ,
		@SerialName("additionalMarketplaces")  val additionalMarketplaces: List<ModificationCommandDetailedResultActionResAdditionalMarketplaces>  = emptyList(),
)
  // The base class definition for tasks
@Serializable
data class ModificationCommandDetailedResultActionResTasks (
		@SerialName("offer")  val offer:  ModificationCommandDetailedResultActionResTasksOffer ,
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("scheduledAt")  val scheduledAt: String  = "",
		@SerialName("finishedAt")  val finishedAt: String  = "",
		@SerialName("status")  val status: String  = "",
		@SerialName("errors")  val errors: List<ModificationCommandDetailedResultActionResTasksErrors>  = emptyList(),
)
  // The base class definition for offer
@Serializable
data class ModificationCommandDetailedResultActionResTasksOffer (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for errors
@Serializable
data class ModificationCommandDetailedResultActionResTasksErrors (
		@SerialName("code")  val code: String  = "",
		@SerialName("details")  val details: String  = "",
		@SerialName("message")  val message: String  = "",
		@SerialName("path")  val path: String  = "",
		@SerialName("userMessage")  val userMessage: String  = "",
		@SerialName("metadata")  val metadata:  ModificationCommandDetailedResultActionResTasksErrorsMetadata ,
)
  // The base class definition for metadata
@Serializable
data class ModificationCommandDetailedResultActionResTasksErrorsMetadata (
		@SerialName("productId")  val productId: String  = "",
)
  // The base class definition for modification
@Serializable
data class ModificationCommandDetailedResultActionResModification (
		@SerialName("basePackage")  val basePackage:  ModificationCommandDetailedResultActionResModificationBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<ModificationCommandDetailedResultActionResModificationExtraPackages>  = emptyList(),
		@SerialName("modificationTime")  val modificationTime: String  = "",
)
  // The base class definition for basePackage
@Serializable
data class ModificationCommandDetailedResultActionResModificationBasePackage (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class ModificationCommandDetailedResultActionResModificationExtraPackages (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for additionalMarketplaces
@Serializable
data class ModificationCommandDetailedResultActionResAdditionalMarketplaces (
		@SerialName("marketplaceId")  val marketplaceId: String  = "",
		@SerialName("modification")  val modification:  ModificationCommandDetailedResultActionResAdditionalMarketplacesModification ,
)
  // The base class definition for modification
@Serializable
data class ModificationCommandDetailedResultActionResAdditionalMarketplacesModification (
		@SerialName("basePackage")  val basePackage:  ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage ,
		@SerialName("extraPackages")  val extraPackages: List<ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages>  = emptyList(),
		@SerialName("modificationTime")  val modificationTime: String  = "",
)
  // The base class definition for basePackage
@Serializable
data class ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage (
		@SerialName("id")  val id: String  = "",
)
  // The base class definition for extraPackages
@Serializable
data class ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages (
		@SerialName("id")  val id: String  = "",
)