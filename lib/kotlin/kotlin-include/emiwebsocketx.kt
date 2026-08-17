package emikot

import kotlinx.serialization.KSerializer
import kotlinx.serialization.json.Json
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener

/**
 * EmiWebSocketX - generated once per module, backs every "method: reactive" action (see
 * kotlin-action-reactive-render.go). A thin, explicitly-typed wrapper around OkHttp's
 * WebSocket, using kotlinx.serialization to JSON-encode/decode Send/Receive - a reactive
 * action is exactly as strongly-typed as a classic one.
 *
 * Because this is a plain (non-inline) generic class, it can't use Kotlin's reified
 * `Json.decodeFromString<Receive>()` the way a per-action generated file's classic
 * compute() can (see kotlin-action-render.go) - callers (the generated `<Action>.Create`
 * factory) supply the two KSerializers explicitly, using the `.serializer()` companion
 * every @Serializable class already gets from the kotlinx.serialization compiler plugin.
 */
class EmiWebSocketX<Send : Any, Receive : Any>(
    private val client: OkHttpClient,
    private val url: String,
    private val sendSerializer: KSerializer<Send>,
    private val receiveSerializer: KSerializer<Receive>,
    // Resolved through ClientContext.resolve() by the generated `<Action>.Create(...)`
    // factory (defaultHeaders merged in, then onRequest if set) - sent as the
    // WebSocket handshake's HTTP headers, same as a classic action's request headers.
    private val headers: Map<String, String> = emptyMap(),
) {
    private val json = Json { ignoreUnknownKeys = true }
    private var socket: WebSocket? = null

    var onMessage: ((Receive) -> Unit)? = null
    var onError: ((Throwable) -> Unit)? = null
    var onClose: ((code: Int, reason: String) -> Unit)? = null

    // Opens the connection. Safe to call at most once per instance - build a new one
    // via the generated `<Action>.Create(...)` factory to reconnect.
    fun connect(): EmiWebSocketX<Send, Receive> {
        val requestBuilder = Request.Builder().url(url)
        headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }
        val request = requestBuilder.build()
        socket = client.newWebSocket(request, object : WebSocketListener() {
            override fun onMessage(webSocket: WebSocket, text: String) {
                try {
                    val decoded = json.decodeFromString(receiveSerializer, text)
                    onMessage?.invoke(decoded)
                } catch (e: Exception) {
                    onError?.invoke(e)
                }
            }

            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
                onError?.invoke(t)
            }

            override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
                onClose?.invoke(code, reason)
            }
        })
        return this
    }

    fun send(value: Send) {
        val text = json.encodeToString(sendSerializer, value)
        socket?.send(text)
    }

    fun close(code: Int = 1000, reason: String = "") {
        socket?.close(code, reason)
    }
}
