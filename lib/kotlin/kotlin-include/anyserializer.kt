package emikot

import kotlinx.serialization.KSerializer
import kotlinx.serialization.descriptors.SerialDescriptor
import kotlinx.serialization.descriptors.buildClassSerialDescriptor
import kotlinx.serialization.encoding.Decoder
import kotlinx.serialization.encoding.Encoder
import kotlinx.serialization.json.JsonDecoder
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonEncoder
import kotlinx.serialization.json.JsonNull
import kotlinx.serialization.json.JsonPrimitive
import kotlinx.serialization.json.JsonArray
import kotlinx.serialization.json.JsonObject

/**
 * AnySerializer - what a `type: any` field actually gets annotated with (see
 * kotlin-common-fields.go, ComputedType == "Any"), instead of a bare `@Contextual`.
 * `@Contextual` alone requires the caller to register a SerializersModule for `Any`
 * before encoding/decoding a dto that has one - a dead end for `Any` specifically,
 * since there's no single meaningful serializer for "whatever type this turns out to
 * be" to register. AnySerializer sidesteps that by working entirely in terms of
 * JsonElement: best-effort-converts common Kotlin values (String/number/Boolean/
 * List/Map/null) to a JsonElement on encode, and hands back the raw JsonElement on
 * decode - the same "expose the raw JsonElement, let the caller unwrap it" contract
 * MaybeFieldSerializer already uses for a structured MaybeField<T>'s Value payload
 * (see common.kt), so a `type: any` field behaves consistently with every other
 * "shape not known at codegen time" field this compiler produces.
 */
object AnySerializer : KSerializer<Any> {
    override val descriptor: SerialDescriptor = buildClassSerialDescriptor("Any")

    override fun serialize(encoder: Encoder, value: Any) {
        val jsonEncoder = encoder as? JsonEncoder ?: error("JsonEncoder required")
        jsonEncoder.encodeJsonElement(toJsonElement(value))
    }

    override fun deserialize(decoder: Decoder): Any {
        val jsonDecoder = decoder as? JsonDecoder ?: error("JsonDecoder required")
        return jsonDecoder.decodeJsonElement()
    }

    private fun toJsonElement(value: Any?): JsonElement = when (value) {
        null -> JsonNull
        is JsonElement -> value
        is String -> JsonPrimitive(value)
        is Boolean -> JsonPrimitive(value)
        is Number -> JsonPrimitive(value)
        is Map<*, *> -> JsonObject(value.entries.associate { (k, v) -> k.toString() to toJsonElement(v) })
        is Iterable<*> -> JsonArray(value.map { toJsonElement(it) })
        else -> JsonPrimitive(value.toString())
    }
}
