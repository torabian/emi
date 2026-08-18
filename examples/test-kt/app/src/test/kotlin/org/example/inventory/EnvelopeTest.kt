package org.example.inventory

import emikot.GResponse
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.JsonPrimitive
import emikot.Maybe
import kotlin.test.Test
import kotlin.test.assertEquals

/**
 * Decodes a full Google JSON Style Guide response (see
 * lib/js/ts-envelopes/google-json-style-guide/google-envelop.emi.yml, the spec
 * lib/kotlin/kotlin-include/gresponse.kt's GResponse/GResponseData/GResponseError
 * mirror) directly - every documented field, not just the item/items payload other
 * tests exercise incidentally through an entity action's envelope.
 */
class EnvelopeTest {

    private val json = Json { ignoreUnknownKeys = true }

    private fun content(field: emikot.MaybeField<*>): String =
        ((field.value as Maybe.Value<*>).v as JsonPrimitive).content

    @Test
    fun `decodes every top-level and data field`() {
        // ProductDto has several required (no-default) fields, so a JSON `data.item`
        // has to be a *complete* ProductDto, not a `{"title": "Widget"}` stub - reuse
        // sampleProduct() (see ProductDataTypesTest.kt) encoded to JSON, spliced into
        // the rest of the literal Google-JSON-style-guide envelope below.
        val itemJson = json.encodeToString(sampleProduct())
        val raw = """
            {
              "apiVersion": "1.0",
              "context": "ctx-1",
              "id": "req-1",
              "method": "products.get",
              "data": {
                "item": $itemJson,
                "editLink": "https://example.com/edit",
                "selfLink": "https://example.com/self",
                "kind": "inventory#product",
                "fields": "title,status",
                "etag": "etag-1",
                "cursor": "cursor-1",
                "id": "product-1",
                "lang": "en",
                "updated": "2026-01-01T00:00:00Z",
                "currentItemCount": 1,
                "itemsPerPage": 20,
                "startIndex": 0,
                "totalItems": 1,
                "totalAvailableItems": 1,
                "pageIndex": 1,
                "totalPages": 1
              }
            }
        """.trimIndent()

        val response = json.decodeFromString<GResponse<ProductDto>>(raw)

        assertEquals("1.0", content(response.apiVersion))
        assertEquals("ctx-1", content(response.context))
        assertEquals("req-1", content(response.id))
        assertEquals("products.get", content(response.method))

        val data = response.data!!
        assertEquals("Widget", data.item?.title)
        assertEquals("https://example.com/edit", content(data.editLink))
        assertEquals("https://example.com/self", content(data.selfLink))
        assertEquals("inventory#product", content(data.kind))
        assertEquals("title,status", content(data.fields))
        assertEquals("etag-1", content(data.etag))
        assertEquals("cursor-1", content(data.cursor))
        assertEquals("product-1", content(data.id))
        assertEquals("en", content(data.lang))
        assertEquals("2026-01-01T00:00:00Z", content(data.updated))
        assertEquals("1", content(data.currentItemCount))
        assertEquals("20", content(data.itemsPerPage))
        assertEquals("0", content(data.startIndex))
        assertEquals("1", content(data.totalItems))
        assertEquals("1", content(data.totalAvailableItems))
        assertEquals("1", content(data.pageIndex))
        assertEquals("1", content(data.totalPages))
    }

    @Test
    fun `decodes a list response via data items`() {
        val itemA = json.encodeToString(sampleProduct().copy(title = "A"))
        val itemB = json.encodeToString(sampleProduct().copy(title = "B"))
        val raw = """{"data":{"items":[$itemA,$itemB]}}"""

        val response = json.decodeFromString<GResponse<ProductDto>>(raw)

        assertEquals(listOf("A", "B"), response.data?.items?.map { it.title })
    }

    @Test
    fun `decodes a full error object`() {
        val raw = """
            {
              "error": {
                "code": 404,
                "message": "Not Found",
                "messageTranslated": "Introuvable",
                "errors": [
                  {
                    "domain": "global",
                    "reason": "notFound",
                    "message": "Product not found",
                    "messageTranslated": "Produit introuvable",
                    "location": "uniqueId",
                    "locationType": "parameter",
                    "extendedHelp": "https://example.com/help",
                    "sendReport": "https://example.com/report"
                  }
                ]
              }
            }
        """.trimIndent()

        val response = json.decodeFromString<GResponse<ProductDto>>(raw)

        val error = response.error!!
        assertEquals(404, error.code)
        assertEquals("Not Found", error.message)
        assertEquals("Introuvable", error.messageTranslated)
        assertEquals(1, error.errors.size)

        val detail = error.errors[0]
        assertEquals("global", content(detail.domain))
        assertEquals("notFound", content(detail.reason))
        assertEquals("Product not found", content(detail.message))
        assertEquals("Produit introuvable", content(detail.messageTranslated))
        assertEquals("uniqueId", content(detail.location))
        assertEquals("parameter", content(detail.locationType))
        assertEquals("https://example.com/help", content(detail.extendedHelp))
        assertEquals("https://example.com/report", content(detail.sendReport))
    }
}
