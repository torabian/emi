package unknownpackage
import kotlinx.serialization.json.*
import kotlinx.serialization.*
  // The base class definition for giantDto
@Serializable
data class GiantDto (
		@SerialName("firstName") val firstName: String,
		@SerialName("firstNameNullable") val firstNameNullable: String?,
		@SerialName("array") val array: List<GiantDtoArray>,
		@SerialName("arrayNullable") val arrayNullable: List<GiantDtoArrayNullable>,
		@SerialName("booleanField") val booleanField: Boolean,
		@SerialName("booleanFieldNullable") val booleanFieldNullable: Boolean?,
		@SerialName("collectionItems") val collectionItems: List<GiantDto>,
		@SerialName("collectionItemsNullable") val collectionItemsNullable: List<GiantDto>?,
		@SerialName("dateObject") val dateObject: Any,
		@SerialName("singleRefNullable") val singleRefNullable: Any,
		@SerialName("enumeration") val enumeration: String,
		@SerialName("enumerationNullable") val enumerationNullable: String?,
		@SerialName("floatingPoint32") val floatingPoint32: Float,
		@SerialName("floatingPoint32Nullable") val floatingPoint32Nullable: Float?,
		@SerialName("floatingPoint64") val floatingPoint64: Double,
		@SerialName("floatingPoint64Nullable") val floatingPoint64Nullable: Double?,
		@SerialName("integerValue") val integerValue: Int,
		@SerialName("integer32ValueNullable") val integer32ValueNullable: Int?,
		@SerialName("integer32Value") val integer32Value: Int,
		@SerialName("integer64ValueNullable") val integer64ValueNullable: Long?,
		@SerialName("integer64Value") val integer64Value: Long,
		@SerialName("mapValue") val mapValue: Any,
		@SerialName("mapValueNullable") val mapValueNullable: Any,
		@SerialName("sliceValue") val sliceValue: List<String>,
		@SerialName("sliceValueNullable") val sliceValueNullable: Any,
		@SerialName("objectValue") val objectValue:  GiantDtoObjectValue,
)
  // The base class definition for array
@Serializable
data class GiantDtoArray (
		@SerialName("subItem1") val subItem1: Int,
)
  // The base class definition for arrayNullable
@Serializable
data class GiantDtoArrayNullable (
		@SerialName("subItemNullable1") val subItemNullable1: Int?,
)
  // The base class definition for objectValue
@Serializable
data class GiantDtoObjectValue (
		@SerialName("innerObject") val innerObject: List<GiantDtoObjectValueInnerObject>,
)
  // The base class definition for innerObject
@Serializable
data class GiantDtoObjectValueInnerObject (
		@SerialName("innerObjText") val innerObjText: String,
)