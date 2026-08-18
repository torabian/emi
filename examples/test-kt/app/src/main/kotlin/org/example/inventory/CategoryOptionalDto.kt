package org.example.inventory
import emikot.Maybe
import emikot.MaybeField
import kotlinx.serialization.*
import kotlinx.serialization.json.*
  // The base class definition for categoryOptionalDto
@Serializable
data class CategoryOptionalDto (
		@SerialName("uniqueId")  val uniqueId: MaybeField<String>  = MaybeField(Maybe.Absent),
		@SerialName("name")  val name: MaybeField<String>  = MaybeField(Maybe.Absent),
)