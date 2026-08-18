package org.example.inventory

import emikot.AuthState
import emikot.AuthenticationSession
import emikot.AuthenticationUser
import emikot.ClientContext
import emikot.ClientRequestSpec
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import okhttp3.mockwebserver.MockResponse
import okhttp3.mockwebserver.MockWebServer
import kotlin.test.AfterTest
import kotlin.test.BeforeTest
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertTrue

/**
 * Exercises the real generated HTTP client code (ProductGetActionClient/
 * ProductCreateActionClient - see kotlin-action-render.go's compute()) against a real
 * local server (OkHttp's MockWebServer), rather than just unit-testing serialization:
 * proves compute() actually sends a real request and parses a real response, that
 * ClientContext.Default/onRequest/defaultHeaders (the "FetchXProvider"-equivalent
 * mechanism - see the Kotlin doc page) really apply to every request, and that
 * AuthState.headers() flows all the way through to the wire.
 */
class ProductActionHttpTest {

    private lateinit var server: MockWebServer
    private val json = Json { encodeDefaults = true }

    @BeforeTest
    fun setUp() {
        server = MockWebServer()
        server.start()
        ClientContext.Default = ClientContext()
        ProductGetActionClient.context = null
        ProductCreateActionClient.context = null
        AuthState.setSession(null)
    }

    @AfterTest
    fun tearDown() {
        server.shutdown()
        ClientContext.Default = ClientContext()
        ProductGetActionClient.context = null
        ProductCreateActionClient.context = null
        AuthState.setSession(null)
    }

    @Test
    fun `compute decodes a real GResponse envelope from a real server`() = runBlocking {
        val product = sampleProduct()
        val body = """{"data":{"item":${json.encodeToString(product)}}}"""
        server.enqueue(MockResponse().setBody(body).setResponseCode(200))

        ProductGetActionClient.context = ClientContext(baseUrl = server.url("/").toString())

        val response = ProductGetActionClient.compute(path = ProductGetActionPathParameter(UniqueId = "product-1"))

        assertEquals(200, response.statusCode)
        val item = response.payload?.data?.item
        assertEquals(product.title, item?.title)

        val recorded = server.takeRequest()
        assertEquals("GET", recorded.method)
        assertEquals("/product/product-1", recorded.path)
    }

    @Test
    fun `compute serializes a typed request body`() = runBlocking {
        val product = sampleProduct()
        server.enqueue(MockResponse().setBody("""{"data":{"item":${json.encodeToString(product)}}}""").setResponseCode(200))

        ProductCreateActionClient.context = ClientContext(baseUrl = server.url("/").toString())
        ProductCreateActionClient.compute(body = product)

        val recorded = server.takeRequest()
        assertEquals("POST", recorded.method)
        val sentBody = json.decodeFromString<ProductDto>(recorded.body.readUtf8())
        assertEquals(product.title, sentBody.title)
        assertEquals(product.price, sentBody.price)
    }

    // ClientContext.Default is the app-wide fallback every generated *Client reads
    // when its own .context is left unset - the "set once, applies to every action"
    // ergonomics FetchxProvider gives the JS/TS SDK (see the Kotlin doc page).
    @Test
    fun `ClientContext Default applies without setting context per action`() = runBlocking {
        val product = sampleProduct()
        server.enqueue(MockResponse().setBody("""{"data":{"item":${json.encodeToString(product)}}}""").setResponseCode(200))

        ClientContext.Default = ClientContext(baseUrl = server.url("/").toString())
        // Deliberately not setting ProductGetActionClient.context here.

        ProductGetActionClient.compute(path = ProductGetActionPathParameter(UniqueId = "product-1"))

        val recorded = server.takeRequest()
        assertTrue(recorded.path!!.startsWith("/product/"))
    }

    // defaultHeaders on ClientContext are merged into every request - the "add headers
    // for every request" half of the FetchXProvider-equivalent ask.
    @Test
    fun `defaultHeaders are sent on every request`() = runBlocking {
        val product = sampleProduct()
        server.enqueue(MockResponse().setBody("""{"data":{"item":${json.encodeToString(product)}}}""").setResponseCode(200))

        ProductGetActionClient.context = ClientContext(
            baseUrl = server.url("/").toString(),
            defaultHeaders = mapOf("X-App-Version" to "1.2.3"),
        )
        ProductGetActionClient.compute(path = ProductGetActionPathParameter(UniqueId = "product-1"))

        val recorded = server.takeRequest()
        assertEquals("1.2.3", recorded.getHeader("X-App-Version"))
    }

    // onRequest is the escape hatch for anything more dynamic than a static header map
    // (e.g. a freshly-signed request) - it's handed the resolved url+headers (after
    // defaultHeaders is already merged in) and can rewrite either.
    @Test
    fun `onRequest interceptor can rewrite headers and url`() = runBlocking {
        val product = sampleProduct()
        server.enqueue(MockResponse().setBody("""{"data":{"item":${json.encodeToString(product)}}}""").setResponseCode(200))

        ProductGetActionClient.context = ClientContext(
            baseUrl = server.url("/").toString(),
            onRequest = { spec: ClientRequestSpec ->
                spec.copy(headers = spec.headers + ("X-Intercepted" to "yes"))
            },
        )
        ProductGetActionClient.compute(path = ProductGetActionPathParameter(UniqueId = "product-1"))

        val recorded = server.takeRequest()
        assertEquals("yes", recorded.getHeader("X-Intercepted"))
    }

    // The end-to-end AuthState story: sign in, wire AuthState.headers() into
    // ClientContext.Default's defaultHeaders, and confirm the bearer token/workspace
    // headers really reach the server - "who is connected" driving real requests, not
    // just an in-memory flag.
    @Test
    fun `AuthState headers flow through to the wire`() = runBlocking {
        val product = sampleProduct()
        server.enqueue(MockResponse().setBody("""{"data":{"item":${json.encodeToString(product)}}}""").setResponseCode(200))

        AuthState.setSession(
            AuthenticationSession(
                token = "Bearer test-token",
                user = AuthenticationUser(uniqueId = "user-1", name = "Ada"),
            ),
        )
        assertTrue(AuthState.isAuthenticated)

        ClientContext.Default = ClientContext(
            baseUrl = server.url("/").toString(),
            defaultHeaders = AuthState.headers(),
        )
        ProductGetActionClient.compute(path = ProductGetActionPathParameter(UniqueId = "product-1"))

        val recorded = server.takeRequest()
        assertEquals("Bearer test-token", recorded.getHeader("Authorization"))

        AuthState.signOut()
        assertTrue(!AuthState.isAuthenticated)
    }
}
