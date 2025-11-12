package com.example.flightstick

import androidx.compose.runtime.*
import kotlinx.coroutines.*
import okhttp3.*

enum class WebSocketStatus { CONNECTING, CONNECTED, FAILED, CLOSED }

@Composable
fun StreamYokeAction(
    serverUrl: String,
    gyroData: Triple<Float, Float, Float>
): String {
    val wsUrl = serverUrl.replace("http", "ws") + "/stream-yoke"
    val client = remember { OkHttpClient() }
    var socket: WebSocket? by remember { mutableStateOf(null) }
    var status by remember { mutableStateOf(WebSocketStatus.CONNECTING) }
    var statusText by remember { mutableStateOf("")}

    // Open WebSocket once
    LaunchedEffect(Unit) {
        val request = Request.Builder().url(wsUrl).build()
        socket = client.newWebSocket(request, object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: Response) {
                statusText = ("✅ Connected to WebSocket")
                status = WebSocketStatus.CONNECTED
            }

            override fun onMessage(webSocket: WebSocket, text: String) {
                statusText = ("📨 Received: $text")
            }

            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
                statusText=("❌ WebSocket error: ${t.message}, url: ${wsUrl}")
                status = WebSocketStatus.FAILED
            }

            override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
                statusText = ("⚠️ Closing WebSocket: $reason")
                webSocket.close(1000, null)
                status = WebSocketStatus.CLOSED
            }
        })
    }

    // Send gyro data whenever it changes
    LaunchedEffect(gyroData) {
        if (status == WebSocketStatus.CONNECTED) {
            val (x, y, z) = gyroData
            val json = """{"x": $x, "y": $y, "z": $z}"""
            socket?.send(json)
        }
    }

    return statusText
}
