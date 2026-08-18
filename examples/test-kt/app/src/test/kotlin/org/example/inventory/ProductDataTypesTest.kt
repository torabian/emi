package org.example.inventory

import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import kotlinx.serialization.json.JsonPrimitive
import org.example.billing.InvoiceLineDto
import org.example.money.Money
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertIs

private val json = Json { encodeDefaults = true }

/**
 * A fully-populated ProductDto - every field from inventory.emi.yml given a concrete
 * value, including every nullable field set to Maybe.Value(...) rather than left
 * Absent. Shared (top-level, same-package - no import needed) across this file's tests
 * and ProductFormStateTest/CrossModuleTest/ProductActionHttpTest.
 */
fun sampleProduct(): ProductDto = ProductDto(
    uniqueId = MaybeField(Maybe.Value("product-1")),
    title = "Widget",
    sku = MaybeField(Maybe.Value("WID-1")),
    active = true,
    featured = MaybeField(Maybe.Value(true)),
    viewCount = 10,
    viewCountOpt = MaybeField(Maybe.Value(11)),
    smallCount = 12,
    smallCountOpt = MaybeField(Maybe.Value(13)),
    bigCount = 14L,
    bigCountOpt = MaybeField(Maybe.Value(15L)),
    ratio32 = 1.5f,
    ratio32Opt = MaybeField(Maybe.Value(1.75f)),
    ratio64 = 2.5,
    ratio64Opt = MaybeField(Maybe.Value(2.75)),
    status = "active",
    statusOpt = MaybeField(Maybe.Value("discontinued")),
    tags = listOf("a", "b"),
    tagsOpt = MaybeField(Maybe.Value(listOf("c", "d"))),
    labelsMeta = mapOf("color" to "red"),
    labelsMetaOpt = MaybeField(Maybe.Value(mapOf("size" to "L"))),
    gallery = listOf(ProductDtoGallery(url = "https://example.com/1.png")),
    galleryOpt = MaybeField(Maybe.Value(listOf(ProductDtoGalleryOpt(url = "https://example.com/2.png")))),
    specs = ProductDtoSpecs(weightGrams = 500),
    price = Money(currency = "USD", minorUnits = 1999),
    misc = "anything",
    category = MaybeField(Maybe.Value(CategoryDto(uniqueId = MaybeField(Maybe.Value("cat-1")), name = "Gadgets"))),
    invoiceLines = MaybeField(Maybe.Value(listOf(
        InvoiceLineDto(description = "line 1", quantity = 1, unitPrice = Money(currency = "USD", minorUnits = 500))
    ))),
)

/**
 * Covers every non-Go-only field type in lib/core/EmiFieldType.go's
 * GetEmiFieldTypeCatalog, both nullable and non-nullable where a nullable variant
 * exists - see inventory.emi.yml's `product` entity, which was written to exercise
 * exactly this catalog for the Kotlin compiler.
 */
class ProductDataTypesTest {

    // Every non-nullable structured field (List/Map/nested-object/complex) decodes to
    // a real, directly comparable value with no MaybeField involved - a clean, full
    // round trip.
    @Test
    fun `non-nullable structured fields survive a JSON round trip exactly`() {
        val original = sampleProduct()
        val decoded = json.decodeFromString<ProductDto>(json.encodeToString(original))

        assertEquals(original.title, decoded.title)
        assertEquals(original.active, decoded.active)
        assertEquals(original.viewCount, decoded.viewCount)
        assertEquals(original.smallCount, decoded.smallCount)
        assertEquals(original.bigCount, decoded.bigCount)
        assertEquals(original.ratio32, decoded.ratio32)
        assertEquals(original.ratio64, decoded.ratio64)
        assertEquals(original.status, decoded.status)
        assertEquals(original.tags, decoded.tags)                 // slice
        assertEquals(original.labelsMeta, decoded.labelsMeta)     // map
        assertEquals(original.gallery, decoded.gallery)           // array of nested object
        assertEquals(original.specs, decoded.specs)               // object
        assertEquals(original.price, decoded.price)               // complex (Money)
    }

    // Nullable primitive fields round-trip their content the same way established last
    // session for MaybeField<String> - MaybeFieldSerializer.deserialize (common.kt)
    // decodes a present value into a raw JsonElement, so the content is unwrapped via
    // JsonPrimitive.content rather than a direct equality check.
    @Test
    fun `nullable primitive fields round trip their content`() {
        val decoded = json.decodeFromString<ProductDto>(json.encodeToString(sampleProduct()))

        fun <T> content(field: MaybeField<T>): String = ((field.value as Maybe.Value<*>).v as JsonPrimitive).content

        assertEquals("WID-1", content(decoded.sku))
        assertEquals("true", content(decoded.featured))
        assertEquals("11", content(decoded.viewCountOpt))
        assertEquals("13", content(decoded.smallCountOpt))
        assertEquals("15", content(decoded.bigCountOpt))
        assertEquals("1.75", content(decoded.ratio32Opt))
        assertEquals("2.75", content(decoded.ratio64Opt))
        assertEquals("discontinued", content(decoded.statusOpt))  // enum? resolves like string?
    }

    // Every nullable field on a fresh (all-defaults) ProductDto is Absent - proves the
    // MaybeField(Maybe.Absent) default KotlinSafeDefaultValue now generates for every
    // nullable structured type (one?/collection?/map?/enum?), not just object? as
    // before this change.
    @Test
    fun `nullable fields default to absent`() {
        val minimal = ProductDto(
            active = false,
            featured = MaybeField(Maybe.Absent),
            viewCountOpt = MaybeField(Maybe.Absent),
            smallCountOpt = MaybeField(Maybe.Absent),
            bigCountOpt = MaybeField(Maybe.Absent),
            ratio32Opt = MaybeField(Maybe.Absent),
            ratio64Opt = MaybeField(Maybe.Absent),
            tagsOpt = MaybeField(Maybe.Absent),
            galleryOpt = MaybeField(Maybe.Absent),
            specs = ProductDtoSpecs(),
            price = Money(),
            misc = Unit,
        )

        assertEquals(Maybe.Absent, minimal.sku.value)
        assertEquals(Maybe.Absent, minimal.statusOpt.value)
        assertEquals(Maybe.Absent, minimal.labelsMetaOpt.value)
        assertEquals(Maybe.Absent, minimal.category.value)        // one?
        assertEquals(Maybe.Absent, minimal.invoiceLines.value)    // collection?, cross-module
    }

    // one?/collection?/map? relation and structured fields resolve to their real
    // Kotlin type (MaybeField<CategoryDto>, MaybeField<List<InvoiceLineDto>>,
    // MaybeField<Map<String, String>>) rather than the `@Contextual Any` every one of
    // these fell back to before this change - this test wouldn't even compile
    // otherwise. Direct construction/access (not JSON) since MaybeField<Structured>'s
    // generic decode doesn't preserve the concrete type - see the class doc on
    // emikot.toDisplayString.
    @Test
    fun `relation and map-typed fields are strongly typed, not Any`() {
        val product = sampleProduct()

        val category = (product.category.value as Maybe.Value).v
        assertIs<CategoryDto>(category)
        assertEquals("Gadgets", category.name)

        val invoiceLines = (product.invoiceLines.value as Maybe.Value).v
        assertIs<List<InvoiceLineDto>>(invoiceLines)
        assertEquals("line 1", invoiceLines[0].description)

        val labelsMetaOpt = (product.labelsMetaOpt.value as Maybe.Value).v
        assertIs<Map<String, String>>(labelsMetaOpt)
        assertEquals("L", labelsMetaOpt["size"])
    }
}
