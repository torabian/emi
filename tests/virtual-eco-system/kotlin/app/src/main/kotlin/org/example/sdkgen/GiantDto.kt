package unknownpackage
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlinx.serialization.descriptors.SerialDescriptor
import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*


@Serializable
sealed class Maybe<out T> {
    @Serializable
    @SerialName("absent")
    data object Absent : Maybe<Nothing>()

    @Serializable
    @SerialName("null")
    data object Null : Maybe<Nothing>()

    @Serializable
    @SerialName("value")
    data class Value<T>(val v: T) : Maybe<T>()
}

object MaybeFieldSerializer : KSerializer<MaybeField<Any?>> {
    override val descriptor: SerialDescriptor =
        buildClassSerialDescriptor("MaybeField")

    override fun serialize(encoder: Encoder, value: MaybeField<Any?>) {
        val jsonEncoder = encoder as? JsonEncoder ?: error("JsonEncoder required")

        when (val v = value.value) {
            Maybe.Absent -> return   // <--- key will be omitted
            Maybe.Null -> jsonEncoder.encodeJsonElement(JsonNull)
            is Maybe.Value<*> -> {
                val el = v.v as? JsonElement ?: JsonPrimitive(v.v.toString())
                jsonEncoder.encodeJsonElement(el)
            }
        }
    }

    override fun deserialize(decoder: Decoder): MaybeField<Any?> {
        val jsonDecoder = decoder as? JsonDecoder ?: error("JsonDecoder required")
        val element = jsonDecoder.decodeJsonElement()
        val maybe = if (element is JsonNull) Maybe.Null else Maybe.Value(element)
        return MaybeField(maybe)
    }
}




// 2️⃣ Field wrapper with serializer
@Serializable(with = MaybeFieldSerializer::class)
data class MaybeField<T>(val value: Maybe<T> = Maybe.Absent) {
	override fun toString(): String = when (value) {
        is Maybe.Absent -> ""              // optional: treat absent as empty string
        is Maybe.Null -> "null"            // optional
        is Maybe.Value -> value.v.toString()
    }
}

// The base class definition for giantDto
@Serializable
data class GiantDto(
    @SerialName("firstName") var firstName: MaybeField<String> = MaybeField(Maybe.Absent)
) {
	fun toJson(pretty: Boolean = false): String {
        val json = Json {
            encodeDefaults = false      // omit Maybe.Absent
            prettyPrint = pretty
        }
        return json.encodeToString(this)
    }

	override fun toString(): String {
		return this.toJson()
	}

}
  // The base class definition for array
@Serializable
data class GiantDtoArray (
		@SerialName("subItem1")  val subItem1: Int,
)
  // The base class definition for arrayNullable
@Serializable
data class GiantDtoArrayNullable (
		@SerialName("subItemNullable1")  val subItemNullable1: Int?,
)
  // The base class definition for objectValue
@Serializable
data class GiantDtoObjectValue (
		@SerialName("innerObject")  val innerObject: List<GiantDtoObjectValueInnerObject>,
)
  // The base class definition for innerObject
@Serializable
data class GiantDtoObjectValueInnerObject (
		@SerialName("innerObjText")  val innerObjText: String,
)