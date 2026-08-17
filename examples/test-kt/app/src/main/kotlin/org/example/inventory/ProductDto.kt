package org.example.inventory
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import org.example.billing.InvoiceLineDto
import org.example.money.Money
  // The base class definition for productDto
@Serializable
data class ProductDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("title")  val title: String  = "",
		@SerialName("sku")  val sku: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("active")  val active: Boolean  = false,
		@SerialName("featured")  val featured: MaybeField<Boolean> ,
		@SerialName("viewCount")  val viewCount: Int  = 0,
		@SerialName("viewCountOpt")  val viewCountOpt: MaybeField<Int> ,
		@SerialName("smallCount")  val smallCount: Int  = 0,
		@SerialName("smallCountOpt")  val smallCountOpt: MaybeField<Int> ,
		@SerialName("bigCount")  val bigCount: Long  = 0,
		@SerialName("bigCountOpt")  val bigCountOpt: MaybeField<Long> ,
		@SerialName("ratio32")  val ratio32: Float  = 0.0f,
		@SerialName("ratio32Opt")  val ratio32Opt: MaybeField<Float> ,
		@SerialName("ratio64")  val ratio64: Double  = 0.0,
		@SerialName("ratio64Opt")  val ratio64Opt: MaybeField<Double> ,
		@SerialName("status")  val status: String  = "",
		@SerialName("statusOpt")  val statusOpt: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("tags")  val tags: List<String>  = emptyList(),
		@SerialName("tagsOpt")  val tagsOpt: MaybeField<List<String>> ,
		@SerialName("labelsMeta")  val labelsMeta: Map<String, String>  = emptyMap(),
		@SerialName("labelsMetaOpt")  val labelsMetaOpt: MaybeField<Map<String, String>>  = MaybeField(Maybe.Absent),
		@SerialName("gallery")  val gallery: List<ProductDtoGallery>  = emptyList(),
		@SerialName("galleryOpt")  val galleryOpt: MaybeField<List<ProductDtoGalleryOpt>> ,
		@SerialName("specs")  val specs:  ProductDtoSpecs ,
		@SerialName("price")  val price: Money ,
		@SerialName("misc")  @Contextual  val misc: Any ,
		@SerialName("category")  val category: MaybeField<CategoryDto>  = MaybeField(Maybe.Absent),
		@SerialName("invoiceLines")  val invoiceLines: MaybeField<List<InvoiceLineDto>>  = MaybeField(Maybe.Absent),
)
  // The base class definition for gallery
@Serializable
data class ProductDtoGallery (
		@SerialName("url")  val url: String  = "",
)
  // The base class definition for galleryOpt
@Serializable
data class ProductDtoGalleryOpt (
		@SerialName("url")  val url: String  = "",
)
  // The base class definition for specs
@Serializable
data class ProductDtoSpecs (
		@SerialName("weightGrams")  val weightGrams: Int  = 0,
)