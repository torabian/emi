package unknownpackage
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
/**
 * Action to communicate with the action ComputeApiAction
 */
data class ComputeApiActionMeta(
    val name: String = "ComputeApiAction",
    val url: String = "",
    val method: String = "post"
)
/*data class ComputeApiActionRequest(val call: io.ktor.server.application.ApplicationCall)*/
data class ComputeApiActionResponse(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)
object ComputeApiActionClient {
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()
    suspend fun compute(jsonPayload: String): ComputeApiActionResponse =
        withContext(Dispatchers.IO) {
            val meta = ComputeApiActionMeta()
            val body = jsonPayload.toRequestBody(jsonType)
            val request = Request.Builder()
                .url(meta.url)
                .method(meta.method, body)
                .addHeader("Accept", "application/json")
                .build()
            client.newCall(request).execute().use { resp ->
                ComputeApiActionResponse(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}
  // The base class definition for computeApiActionReq
@Serializable
data class ComputeApiActionReq (
		@SerialName("initialVector1") val initialVector1: List<Int>,
		@SerialName("value") val value: String?,
		@SerialName("initialVector2") val initialVector2: List<Int>,
)
  // The base class definition for computeApiActionRes
@Serializable
data class ComputeApiActionRes (
		@SerialName("outputVector") val outputVector: List<Int>,
)