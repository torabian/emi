package unknownpackage
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
/**
 * Action to communicate with the action ModifyGiantDtoAction
 */
data class ModifyGiantDtoActionMeta(
    val name: String = "ModifyGiantDtoAction",
    val url: String = "/modify/dto",
    val method: String = "post"
)
/*data class ModifyGiantDtoActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ModifyGiantDtoActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ModifyGiantDtoActionClient {
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
    suspend fun compute(jsonPayload: String): ModifyGiantDtoActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ModifyGiantDtoActionMeta()
            val body = jsonPayload.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(meta.url)
                .method(meta.method, body)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                ModifyGiantDtoActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}