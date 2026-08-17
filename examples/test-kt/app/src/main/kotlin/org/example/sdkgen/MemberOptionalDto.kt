package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for memberOptionalDto
@Serializable
data class MemberOptionalDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("fullName")  val fullName: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("email")  val email: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("active")  val active: MaybeField<Boolean> ,
		@SerialName("score")  val score: MaybeField<Int> ,
		@SerialName("team")  val team: MaybeField<TeamDto>  = MaybeField(Maybe.Absent),
		@SerialName("badges")  val badges: MaybeField<List<MemberOptionalDtoBadges>> ,
)
  // The base class definition for badges
@Serializable
data class MemberOptionalDtoBadges (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("label")  val label: String  = "",
)