package org.example.inventory
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for categoryDto
@Serializable
data class CategoryDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("name")  val name: String  = "",
)