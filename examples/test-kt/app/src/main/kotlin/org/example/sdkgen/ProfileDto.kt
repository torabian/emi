package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for profileDto
@Serializable
data class ProfileDto (
		@SerialName("displayName")  val displayName: String  = "",
		@SerialName("bio")  val bio: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("address")  val address:  ProfileDtoAddress ,
		@SerialName("tags")  val tags: List<String>  = emptyList(),
		@SerialName("score")  val score: Double  = 0.0,
		@SerialName("level")  val level: String  = "",
)
  // The base class definition for address
@Serializable
data class ProfileDtoAddress (
		@SerialName("street")  val street: String  = "",
		@SerialName("city")  val city: String  = "",
)