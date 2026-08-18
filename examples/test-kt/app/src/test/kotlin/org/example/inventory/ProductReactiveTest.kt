package org.example.inventory

import emikot.ClientContext
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withTimeout
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import okhttp3.mockwebserver.MockResponse
import okhttp3.mockwebserver.MockWebServer
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import kotlin.test.AfterTest
import kotlin.test.BeforeTest
import kotlin.test.Test
import kotlin.test.assertEquals

/**
 * Exercises the real generated WebSocket client (SubscribeProductUpdatesAction/
 * EmiWebSocketX - see kotlin-action-reactive-render.go and
 * kotlin-include/emiwebsocketx.kt) against a real local WebSocket server
 * (MockWebServer's WebSocket upgrade support), including that
 * SubscribeProductUpdatesAction.Create()'s http(s) -> ws(s) scheme swap
 * (webSocketBaseUrl()) actually produces a URL OkHttp's WebSocket client accepts.
 */
class ProductReactiveTest {

    private lateinit var server: MockWebServer

    @BeforeTest
    fun setUp() {
        server = MockWebServer()
        server.start()
        SubscribeProductUpdatesAction.context = null
    }

    @AfterTest
    fun tearDown() {
        server.shutdown()
        SubscribeProductUpdatesAction.context = null
    }

    @Test
    fun `send and receive a typed message over a real WebSocket`() = runBlocking {
        val json = Json { encodeDefaults = true }
        val received = Channel<String>(capacity = 1)
        var serverSocket: WebSocket? = null

        // Echoes back a typed SubscribeProductUpdatesActionRes as soon as the
        // handshake completes, so the test doesn't need a second round trip just to
        // prove the client can receive as well as send.
        val serverListener = object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: okhttp3.Response) {
                serverSocket = webSocket
                webSocket.send(json.encodeToString(SubscribeProductUpdatesActionRes(title = "Widget", status = "active")))
            }

            override fun onMessage(webSocket: WebSocket, text: String) {
                // Echo whatever the client sends, tagged, so the test can also assert
                // the client's own send() reached the server.
                webSocket.send("""{"title":"echo:$text","status":"active"}""")
            }
        }
        server.enqueue(MockResponse().withWebSocketUpgrade(serverListener))

        SubscribeProductUpdatesAction.context = ClientContext(baseUrl = server.url("/").toString())
        val socket = SubscribeProductUpdatesAction.Create(
            path = SubscribeProductUpdatesActionPathParameter(ProductId = "product-1"),
        )
        socket.onMessage = { msg -> received.trySend(msg.title) }
        socket.connect()

        val firstTitle = withTimeout(5_000) { received.receive() }
        assertEquals("Widget", firstTitle)

        socket.send(SubscribeProductUpdatesActionReq(ping = "hello"))
        val echoedTitle = withTimeout(5_000) { received.receive() }
        assertEquals(true, echoedTitle.startsWith("echo:"))

        socket.close()
        serverSocket?.close(1000, null)
    }
}
