package unknownpackage
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import emikot.Maybe
import emikot.MaybeField

  // The base class definition for giantDto
@Serializable
data class GiantDto (
		@SerialName("firstName")  val firstName: String  = "",
		@SerialName("firstNameNullable")  val firstNameNullable: MaybeField<String>  = MaybeField(Maybe.Absent),
		// @SerialName("array")  val array: List<GiantDtoArray>  = emptyList(),
		// @SerialName("arrayNullable")  val arrayNullable: List<GiantDtoArrayNullable> ,
		// @SerialName("booleanField")  val booleanField: Boolean ,
		// @SerialName("booleanFieldNullable")  val booleanFieldNullable: Boolean? ,
		// @SerialName("collectionItems")  val collectionItems: List<GiantDto>  = emptyList(),
		// @SerialName("collectionItemsNullable")  val collectionItemsNullable: List<GiantDto>? ,
		// @SerialName("dateObject")  @Contextual  val dateObject: Any ,
		// @SerialName("singleRefNullable")  @Contextual  val singleRefNullable: Any ,
		// @SerialName("enumeration")  val enumeration: String ,
		// @SerialName("enumerationNullable")  val enumerationNullable: String? ,
		// @SerialName("floatingPoint32")  val floatingPoint32: Float  = 0.0f,
		// @SerialName("floatingPoint32Nullable")  val floatingPoint32Nullable: Float? ,
		// @SerialName("floatingPoint64")  val floatingPoint64: Double  = 0.0f,
		// @SerialName("floatingPoint64Nullable")  val floatingPoint64Nullable: Double? ,
		// @SerialName("integerValue")  val integerValue: Int  = 0,
		// @SerialName("integer32ValueNullable")  val integer32ValueNullable: Int? ,
		// @SerialName("integer32Value")  val integer32Value: Int  = 0,
		// @SerialName("integer64ValueNullable")  val integer64ValueNullable: Long? ,
		// @SerialName("integer64Value")  val integer64Value: Long  = 0,
		// @SerialName("mapValue")  @Contextual  val mapValue: Any ,
		// @SerialName("mapValueNullable")  @Contextual  val mapValueNullable: Any ,
		// @SerialName("sliceValue")  val sliceValue: List<String>  = emptyList(),
		// @SerialName("sliceValueNullable")  @Contextual  val sliceValueNullable: Any ,
		// @SerialName("objectValue")  val objectValue:  GiantDtoObjectValue ,
)
  // The base class definition for array
@Serializable
data class GiantDtoArray (
		@SerialName("subItem1")  val subItem1: Int  = 0,
)
  // The base class definition for arrayNullable
@Serializable
data class GiantDtoArrayNullable (
		@SerialName("subItemNullable1")  val subItemNullable1: Int? = null ,
)
  // The base class definition for objectValue
@Serializable
data class GiantDtoObjectValue (
		@SerialName("innerObject")  val innerObject: List<GiantDtoObjectValueInnerObject>  = emptyList(),
)
  // The base class definition for innerObject
@Serializable
data class GiantDtoObjectValueInnerObject (
		@SerialName("innerObjText")  val innerObjText: String  = "",
)