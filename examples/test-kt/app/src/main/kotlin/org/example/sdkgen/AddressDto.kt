package unknownpackage
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for addressDto
@Serializable
data class AddressDto (
		@SerialName("street")  val street: String  = "",
		@SerialName("city")  val city: String  = "",
		@SerialName("postcode")  val postcode: MaybeField<String>  = MaybeField(Maybe.Absent),
)