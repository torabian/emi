package emikot

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

/**
 * A single HTTP/WebSocket request's url + headers, right before it's sent - the shape
 * ClientContext.onRequest is handed, and expected to return (possibly transformed).
 */
data class ClientRequestSpec(
    val url: String,
    val headers: Map<String, String>,
)

/**
 * Per-client (or, via ClientContext.Default, app-wide) request configuration for every
 * generated <Action>Client/reactive action object - the same role FetchxContext/
 * FetchxProvider play for the JS/TS SDK (see lib/js/ts-sdk/common/fetchx.ts): a base
 * URL prefix, headers merged into every request, and an `onRequest` interceptor hook
 * for anything more dynamic (a fresh auth token per request, request signing,
 * logging) - see the "Client context & authentication" Kotlin doc page and
 * examples/test-kt for a full worked example wiring this up with AuthState.
 *
 * Every generated action reads `context ?: ClientContext.Default` (see
 * kotlin-action-render.go/kotlin-action-reactive-render.go), so setting
 * `ClientContext.Default` once (e.g. at app startup) is enough for every action to
 * pick it up - assigning a client's own `.context` only when that one call needs to
 * deviate from the app-wide default.
 */
data class ClientContext(
    val baseUrl: String = "",
    val defaultHeaders: Map<String, String> = emptyMap(),
    val onRequest: ((ClientRequestSpec) -> ClientRequestSpec)? = null,
) {
    companion object {
        /**
         * The app-wide fallback every generated <Action>Client/reactive action object
         * uses when its own `context` var is left null.
         */
        var Default: ClientContext = ClientContext()
    }

    /**
     * Merges `headers` (the compute()/Create() call's own explicit headers argument)
     * over defaultHeaders - explicit call-site headers win on conflict, matching
     * FetchxContext.apply()'s merge order - then runs `onRequest` if set.
     */
    fun resolve(url: String, headers: Map<String, String>): ClientRequestSpec {
        val spec = ClientRequestSpec(url, defaultHeaders + headers)
        return onRequest?.invoke(spec) ?: spec
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
