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
	// Falls back to the app-wide ClientContext.Default when this object's own
	// .context hasn't been set - same fallback every classic *Client.compute() uses
	// (see kotlin-action-render.go).
	private fun webSocketBaseUrl(effectiveContext: ClientContext): String {
		val base = effectiveContext.baseUrl
		return when {
			base.startsWith("https://") -> "wss://" + base.removePrefix("https://")
			base.startsWith("http://") -> "ws://" + base.removePrefix("http://")
			else -> base
		}
	}
	// Builds an unconnected socket - call .connect() on the result to actually open it,
	// same two-step shape (Create then connect) as Swift's EmiWebSocketX. headers are
	// resolved through ClientContext.resolve() exactly like a classic action's
	// compute() (defaultHeaders merged in, then onRequest if set) and sent as the
	// WebSocket handshake's HTTP headers.
	fun Create(
		path: SubscribeProductUpdatesActionPathParameter,
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
	): SubscribeProductUpdatesActionSocket {
		val meta = SubscribeProductUpdatesActionMeta()
		val effectiveContext = context ?: ClientContext.Default
		val baseUrl = webSocketBaseUrl(effectiveContext).toHttpUrl()
		val urlBuilder = baseUrl.newBuilder().encodedPath(meta.url)
		query.forEach { (k, v) -> urlBuilder.addQueryParameter(k, v) }
		var url = urlBuilder.build().toString()
		url = SubscribeProductUpdatesActionPathParameterApply(path, url)
		val resolved = effectiveContext.resolve(url, headers)
		return SubscribeProductUpdatesActionSocket(
			client = client,
			url = resolved.url,
			headers = resolved.headers,
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