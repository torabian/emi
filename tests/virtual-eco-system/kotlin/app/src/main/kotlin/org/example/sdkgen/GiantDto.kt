package unknownpackage
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for giantDto
@Serializable
data class GiantDto (
		@SerialName("firstName") val firstName: String,
		@SerialName("firstNameNullable") val firstNameNullable: String?,
		@SerialName("array") val array: List<GiantDtoArray>,
		@SerialName("arrayNullable") val arrayNullable: List<GiantDtoArrayNullable>,
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