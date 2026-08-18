package org.example.inventory

import emikot.Maybe
import emikot.MaybeField
import org.example.billing.InvoiceLineDto
import org.example.money.Money
import kotlin.test.Test
import kotlin.test.assertEquals

/**
 * inventory.emi.yml and billing.emi.yml are compiled as two separate `emi kotlin`
 * invocations (see the Makefile's gen-inventory/gen-billing, different --pkg each) that
 * reference each other's entities: inventory's `product` has a `type: collection?`
 * into billing's InvoiceLineEntity, and billing's `invoiceLine` has a `type: one` back
 * into inventory's ProductEntity (see kotlinCollectTargetDeps in
 * kotlin-class-generator.go for how the cross-package import gets generated). Plus a
 * same-module one/collection pair (product.category) for comparison.
 */
class CrossModuleTest {

    @Test
    fun `inventory product collection-references billing InvoiceLineDto`() {
        val line = InvoiceLineDto(description = "widget line", quantity = 2, unitPrice = Money(minorUnits = 2500))
        val product = sampleProduct().copy(invoiceLines = MaybeField(Maybe.Value(listOf(line))))

        val lines = (product.invoiceLines.value as Maybe.Value).v
        assertEquals(1, lines.size)
        assertEquals("widget line", lines[0].description)
        assertEquals(2500L, lines[0].unitPrice.minorUnits)
    }

    // The other direction: billing's invoiceLine.product is `type: one` (non-nullable,
    // unlike inventory's `collection?` back-reference) into org.example.inventory.
    @Test
    fun `billing invoiceLine one-references inventory ProductDto`() {
        val product = sampleProduct()
        val line = InvoiceLineDto(
            description = "widget line",
            quantity = 1,
            unitPrice = Money(minorUnits = 1999),
            product = MaybeField(Maybe.Value(product)),
        )

        val linked = (line.product.value as Maybe.Value).v
        assertEquals(product.title, linked.title)
        assertEquals(product.uniqueId, linked.uniqueId)
    }

    // Same-module pair for comparison: product.category (one?) and, conceptually,
    // category would browse back its products via CategoryBrowseAction/the entity's
    // own query surface - not modeled as a field to avoid an unbounded cycle, same
    // reasoning examples/emi-entity/entity.emi.yml's entity1/entity2 pair documents.
    @Test
    fun `same-module product-category reference is strongly typed`() {
        val category = CategoryDto(uniqueId = MaybeField(Maybe.Value("cat-2")), name = "Tools")
        val product = sampleProduct().copy(category = MaybeField(Maybe.Value(category)))

        val linked = (product.category.value as Maybe.Value).v
        assertEquals("Tools", linked.name)
    }
}
