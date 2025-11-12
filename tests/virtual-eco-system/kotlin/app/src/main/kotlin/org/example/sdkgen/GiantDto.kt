package unknownpackage

import kotlinx.serialization.*
import kotlinx.serialization.json.*


@Serializable
data class GiantDto (
		@SerialName("firstName") val firstName: String,
		@SerialName("firstNameNullable") val firstNameNullable: String?,
		@SerialName("array") val array: List<String>,
		@SerialName("arrayNullable") val arrayNullable: List<String>,
)