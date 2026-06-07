package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for yokeMessageDto
@Serializable
data class YokeMessageDto (
		@SerialName("pitch")  val pitch: Double  = 0.0,
		@SerialName("role")  val role: Double  = 0.0,
)