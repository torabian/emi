package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for memberDto
@Serializable
data class MemberDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("fullName")  val fullName: String  = "",
		@SerialName("email")  val email: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("active")  val active: Boolean  = false,
		@SerialName("score")  val score: Int  = 0,
		@SerialName("team")  val team: MaybeField<TeamDto>  = MaybeField(Maybe.Absent),
		@SerialName("badges")  val badges: List<MemberDtoBadges>  = emptyList(),
)
  // The base class definition for badges
@Serializable
data class MemberDtoBadges (
		@SerialName("label")  val label: String  = "",
)