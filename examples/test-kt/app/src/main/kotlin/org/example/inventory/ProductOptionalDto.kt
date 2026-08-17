package org.example.inventory
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import org.example.billing.InvoiceLineDto
import org.example.money.Money
  // The base class definition for productOptionalDto
@Serializable
data class ProductOptionalDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("title")  val title: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("sku")  val sku: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("active")  val active: MaybeField<Boolean> ,
		@SerialName("featured")  val featured: MaybeField<Boolean> ,
		@SerialName("viewCount")  val viewCount: MaybeField<Int> ,
		@SerialName("viewCountOpt")  val viewCountOpt: MaybeField<Int> ,
		@SerialName("smallCount")  val smallCount: MaybeField<Int> ,
		@SerialName("smallCountOpt")  val smallCountOpt: MaybeField<Int> ,
		@SerialName("bigCount")  val bigCount: MaybeField<Long> ,
		@SerialName("bigCountOpt")  val bigCountOpt: MaybeField<Long> ,
		@SerialName("ratio32")  val ratio32: MaybeField<Float> ,
		@SerialName("ratio32Opt")  val ratio32Opt: MaybeField<Float> ,
		@SerialName("ratio64")  val ratio64: MaybeField<Double> ,
		@SerialName("ratio64Opt")  val ratio64Opt: MaybeField<Double> ,
		@SerialName("status")  val status: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("statusOpt")  val statusOpt: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("tags")  val tags: MaybeField<List<String>> ,
		@SerialName("tagsOpt")  val tagsOpt: MaybeField<List<String>> ,
		@SerialName("labelsMeta")  val labelsMeta: MaybeField<Map<String, String>>  = MaybeField(Maybe.Absent),
		@SerialName("labelsMetaOpt")  val labelsMetaOpt: MaybeField<Map<String, String>>  = MaybeField(Maybe.Absent),
		@SerialName("gallery")  val gallery: MaybeField<List<ProductOptionalDtoGallery>> ,
		@SerialName("galleryOpt")  val galleryOpt: MaybeField<List<ProductOptionalDtoGalleryOpt>> ,
		@SerialName("specs")  val specs: MaybeField<ProductOptionalDtoSpecs>  = MaybeField(Maybe.Absent),
		@SerialName("price")  val price: Money ,
		@SerialName("misc")  @Contextual  val misc: Any ,
		@SerialName("category")  val category: MaybeField<CategoryDto>  = MaybeField(Maybe.Absent),
		@SerialName("invoiceLines")  val invoiceLines: MaybeField<List<InvoiceLineDto>>  = MaybeField(Maybe.Absent),
)
  // The base class definition for gallery
@Serializable
data class ProductOptionalDtoGallery (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("url")  val url: String  = "",
)
  // The base class definition for galleryOpt
@Serializable
data class ProductOptionalDtoGalleryOpt (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("url")  val url: String  = "",
)
  // The base class definition for specs
@Serializable
data class ProductOptionalDtoSpecs (
		@SerialName("weightGrams")  val weightGrams: MaybeField<Int> ,
)