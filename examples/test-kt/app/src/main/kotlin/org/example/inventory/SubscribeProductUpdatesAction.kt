package org.example.inventory
import emikot.ClientContext
import emikot.EmiWebSocketX
import emikot.Maybe
import emikot.MaybeField
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import okhttp3.*
import okhttp3.HttpUrl.Companion.toHttpUrl
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
/**
 * Reactive (WebSocket) action SubscribeProductUpdatesAction
 */
data class SubscribeProductUpdatesActionMeta(
    val name: String = "SubscribeProductUpdatesAction",
    val url: String = "/product/subscribe/:productId"
)
	/**
 * Path parameters for SubscribeProductUpdatesAction
 */
data class SubscribeProductUpdatesActionPathParameter (
	var ProductId: String,
)
// Converts a placeholder url, and applies the parameters to it.
fun SubscribeProductUpdatesActionPathParameterApply(params: SubscribeProductUpdatesActionPathParameter, templateUrl: String): String {
	var url = templateUrl
		url = url.replace(":productId", params.ProductId)
	return url
}
typealias SubscribeProductUpdatesActionSocket = EmiWebSocketX<SubscribeProductUpdatesActionReq, SubscribeProductUpdatesActionRes>
object SubscribeProductUpdatesAction {
	public var context: ClientContext? = null
	private val client = OkHttpClient()
	private fun webSocketBaseUrl(): String {
		val base = context?.baseUrl ?: ""
		return when {
			base.startsWith("https://") -> "wss://" + base.removePrefix("https://")
			base.startsWith("http://") -> "ws://" + base.removePrefix("http://")
			else -> base
		}
	}
	// Builds an unconnected socket - call .connect() on the result to actually open it,
	// same two-step shape (Create then connect) as Swift's EmiWebSocketX.
	fun Create(
		path: SubscribeProductUpdatesActionPathParameter,
		query: Map<String, String> = emptyMap()
	): SubscribeProductUpdatesActionSocket {
		val meta = SubscribeProductUpdatesActionMeta()
		val baseUrl = webSocketBaseUrl().toHttpUrl()
		val urlBuilder = baseUrl.newBuilder().encodedPath(meta.url)
		query.forEach { (k, v) -> urlBuilder.addQueryParameter(k, v) }
		var url = urlBuilder.build().toString()
		url = SubscribeProductUpdatesActionPathParameterApply(path, url)
		return SubscribeProductUpdatesActionSocket(
			client = client,
			url = url,
			sendSerializer = SubscribeProductUpdatesActionReq.serializer(),
			receiveSerializer = SubscribeProductUpdatesActionRes.serializer(),
		)
	}
}
  // The base class definition for subscribeProductUpdatesActionReq
@Serializable
data class SubscribeProductUpdatesActionReq (
		@SerialName("ping")  val ping: String  = "",
)
  // The base class definition for subscribeProductUpdatesActionRes
@Serializable
data class SubscribeProductUpdatesActionRes (
		@SerialName("title")  val title: String  = "",
		@SerialName("status")  val status: String  = "",
)