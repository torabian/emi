package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for teamOptionalDto
@Serializable
data class TeamOptionalDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("title")  val title: MaybeField<String>  = MaybeField(Maybe.Absent),
)