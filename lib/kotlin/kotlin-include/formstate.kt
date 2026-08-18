package emikot

import kotlinx.serialization.json.JsonPrimitive

/**
 * Best-effort plain-string projection of a MaybeField's current value, used by every
 * generated <Dto>FormState.fromDto() (see kotlin-form-state.go, --tags android-forms)
 * to seed a Compose text-field-shaped MutableState<String> - regardless of whether the
 * MaybeField was built directly in Kotlin code (its Maybe.Value wraps the real typed
 * value) or produced by decoding JSON through MaybeFieldSerializer (which wraps a raw
 * JsonElement instead - see common.kt), both shapes unwrap to the same display string.
 * A structured value (list/map/dto) decoded from JSON unwraps to its exact JSON text
 * via JsonElement's own toString(); one built directly in code instead falls back to
 * Kotlin's default toString() for that value, which isn't guaranteed to be valid JSON
 * for every type (e.g. a Map) - fromDto() on a freshly network-decoded dto (the
 * realistic "load a form from a server response" case) always takes the JsonElement
 * path and is unaffected by this.
 */
fun <T> MaybeField<T>.toDisplayString(): String {
    return when (val v = this.value) {
        is Maybe.Value<*> -> when (val inner = v.v) {
            is JsonPrimitive -> inner.content
            null -> ""
            else -> inner.toString()
        }
        else -> ""
    }
}
