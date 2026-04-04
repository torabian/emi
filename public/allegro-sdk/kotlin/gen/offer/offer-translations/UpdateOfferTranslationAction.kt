package unknownpackage
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.withContext
import emikot.ClientContext
import emikot.Maybe
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import kotlinx.coroutines.Dispatchers
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.MaybeField
/**
 * Action to communicate with the action UpdateOfferTranslationAction
 */
data class UpdateOfferTranslationActionMeta(
    val name: String = "UpdateOfferTranslationAction",
    val url: String = "https://api.{environment}/sale/offers/{offerId}/translations/{language}",
    val method: String = "patch"
)
/*data class UpdateOfferTranslationActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class UpdateOfferTranslationActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object UpdateOfferTranslationActionClient {
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
	): UpdateOfferTranslationActionResponse =
        withContext(Dispatchers.IO) {
            val meta = UpdateOfferTranslationActionMeta()
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
                UpdateOfferTranslationActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for updateOfferTranslationActionReq
@Serializable
data class UpdateOfferTranslationActionReq (
		@SerialName("description")  val description:  UpdateOfferTranslationActionReqDescription ,
		@SerialName("title")  val title:  UpdateOfferTranslationActionReqTitle ,
		@SerialName("safetyInformation")  val safetyInformation:  UpdateOfferTranslationActionReqSafetyInformation ,
)
  // The base class definition for description
@Serializable
data class UpdateOfferTranslationActionReqDescription (
		@SerialName("translation")  val translation:  UpdateOfferTranslationActionReqDescriptionTranslation ,
)
  // The base class definition for translation
@Serializable
data class UpdateOfferTranslationActionReqDescriptionTranslation (
		@SerialName("sections")  val sections: List<UpdateOfferTranslationActionReqDescriptionTranslationSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class UpdateOfferTranslationActionReqDescriptionTranslationSections (
		@SerialName("items")  val items: List<UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for title
@Serializable
data class UpdateOfferTranslationActionReqTitle (
		@SerialName("translation")  val translation: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class UpdateOfferTranslationActionReqSafetyInformation (
		@SerialName("products")  val products: List<UpdateOfferTranslationActionReqSafetyInformationProducts>  = emptyList(),
)
  // The base class definition for products
@Serializable
data class UpdateOfferTranslationActionReqSafetyInformationProducts (
		@SerialName("id")  val id: String  = "",
		@SerialName("translation")  val translation: String  = "",
)