package org.example.billing
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import org.example.inventory.ProductDto
import org.example.money.Money
  // The base class definition for invoiceLineDto
@Serializable
data class InvoiceLineDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("description")  val description: String  = "",
		@SerialName("quantity")  val quantity: Int  = 0,
		@SerialName("unitPrice")  val unitPrice: Money ,
		@SerialName("product")  val product: MaybeField<ProductDto>  = MaybeField(Maybe.Absent),
)