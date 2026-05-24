package unknownpackage
import emikot.MaybeField
import emikot.Maybe
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for userentityDto
@Serializable
data class UserentityDto (
		@SerialName("name")  val name: String  = "",
		@SerialName("email")  val email: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("age")  val age: Int  = 0,
		@SerialName("preferences")  val preferences:  UserentityDtoPreferences ,
		@SerialName("tags")  val tags: List<UserentityDtoTags>  = emptyList(),
		  // Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
 @SerialName("address")  val address:  UserentityDtoAddress ,
)
  // The base class definition for preferences
@Serializable
data class UserentityDtoPreferences (
		@SerialName("theme")  val theme: String  = "",
		@SerialName("locale")  val locale: String  = "",
)
  // The base class definition for tags
@Serializable
data class UserentityDtoTags (
		@SerialName("key")  val key: String  = "",
		@SerialName("value")  val value: String  = "",
)
  // The base class definition for address
@Serializable
data class UserentityDtoAddress (
		@SerialName("id")  val id: Long  = 0,
		@SerialName("userId")  val userId: Long  = 0,
		@SerialName("street")  val street: String  = "",
		@SerialName("city")  val city: String  = "",
		@SerialName("postcode")  val postcode: MaybeField<String>  = MaybeField(Maybe.Absent),
)