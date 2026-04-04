package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.HttpUrl.Companion.toHttpUrl
import kotlinx.coroutines.Dispatchers
import kotlinx.serialization.json.*
import emikot.MaybeField
import emikot.Maybe
import kotlinx.coroutines.withContext
import emikot.ClientContext
import kotlinx.serialization.*
/**
 * Action to communicate with the action GetOfferTranslationsAction
 */
data class GetOfferTranslationsActionMeta(
    val name: String = "GetOfferTranslationsAction",
    val url: String = "https://api.{environment}/sale/offers/{offerId}/translations",
    val method: String = "get"
)
/*data class GetOfferTranslationsActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetOfferTranslationsActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object GetOfferTranslationsActionClient {
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
	): GetOfferTranslationsActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetOfferTranslationsActionMeta()
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
                GetOfferTranslationsActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for getOfferTranslationsActionRes
@Serializable
data class GetOfferTranslationsActionRes (
		@SerialName("translations")  val translations: List<GetOfferTranslationsActionResTranslations>  = emptyList(),
)
  // The base class definition for translations
@Serializable
data class GetOfferTranslationsActionResTranslations (
		@SerialName("language")  val language: String  = "",
		@SerialName("title")  val title:  GetOfferTranslationsActionResTranslationsTitle ,
		@SerialName("description")  val description:  GetOfferTranslationsActionResTranslationsDescription ,
		@SerialName("safetyInformation")  val safetyInformation:  GetOfferTranslationsActionResTranslationsSafetyInformation ,
)
  // The base class definition for title
@Serializable
data class GetOfferTranslationsActionResTranslationsTitle (
		@SerialName("translation")  val translation: String  = "",
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for description
@Serializable
data class GetOfferTranslationsActionResTranslationsDescription (
		@SerialName("translation")  val translation:  GetOfferTranslationsActionResTranslationsDescriptionTranslation ,
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for translation
@Serializable
data class GetOfferTranslationsActionResTranslationsDescriptionTranslation (
		@SerialName("sections")  val sections: List<GetOfferTranslationsActionResTranslationsDescriptionTranslationSections>  = emptyList(),
)
  // The base class definition for sections
@Serializable
data class GetOfferTranslationsActionResTranslationsDescriptionTranslationSections (
		@SerialName("items")  val items: List<GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems>  = emptyList(),
)
  // The base class definition for items
@Serializable
data class GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems (
		@SerialName("type")  val type: String  = "",
)
  // The base class definition for safetyInformation
@Serializable
data class GetOfferTranslationsActionResTranslationsSafetyInformation (
		@SerialName("products")  val products: List<GetOfferTranslationsActionResTranslationsSafetyInformationProducts>  = emptyList(),
)
  // The base class definition for products
@Serializable
data class GetOfferTranslationsActionResTranslationsSafetyInformationProducts (
		@SerialName("id")  val id: String  = "",
		@SerialName("translation")  val translation: String  = "",
		@SerialName("type")  val type: String  = "",
)