package unknownpackage
  // The base class definition for giantDto
@Serializable
data class GiantDto (
		@SerialName("firstName") val firstName: string,
		@SerialName("firstNameNullable") val firstNameNullable: String?,
		@SerialName("array") val array: List<GiantDtoArray>,
		@SerialName("arrayNullable") val arrayNullable: List<GiantDtoArrayNullable>,
)