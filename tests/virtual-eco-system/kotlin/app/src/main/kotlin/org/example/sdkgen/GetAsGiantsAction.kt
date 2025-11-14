package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
/**
 * Action to communicate with the action GetAsGiantsAction
 */
data class GetAsGiantsActionMeta(
    val name: String = "GetAsGiantsAction",
    val url: String = "/get/giant/:id",
    val method: String = "PATCH"
)
/*data class GetAsGiantsActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class GetAsGiantsActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)

object GetAsGiantsActionClient {
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
    suspend fun compute(jsonPayload: String): GetAsGiantsActionResponse =
        withContext(Dispatchers.IO) {
            val meta = GetAsGiantsActionMeta()
            val body = jsonPayload.toRequestBody(jsonType)
            val request = Request.Builder()
                .url("http://localhost:8080" + meta.url)
                .method(meta.method, body)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                GetAsGiantsActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}