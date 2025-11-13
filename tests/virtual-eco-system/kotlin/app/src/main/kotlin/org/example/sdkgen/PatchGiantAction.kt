package unknownpackage
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
/**
 * Action to communicate with the action PatchGiantAction
 */
data class PatchGiantActionMeta(
    val name: String = "PatchGiantAction",
    val url: String = "/get/giant/:id",
    val method: String = "patch"
)
/*data class PatchGiantActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class PatchGiantActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object PatchGiantActionClient {
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
    suspend fun compute(jsonPayload: String): PatchGiantActionResponse =
        withContext(Dispatchers.IO) {
            val meta = PatchGiantActionMeta()
            val body = jsonPayload.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(meta.url)
                .method(meta.method, body)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                PatchGiantActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}