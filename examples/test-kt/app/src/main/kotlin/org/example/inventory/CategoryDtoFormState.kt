package org.example.inventory
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import emikot.Maybe
import emikot.MaybeField
import emikot.toDisplayString
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
class CategoryDtoFormState {
	var uniqueId: MutableState<String> = mutableStateOf("")
	var name: MutableState<String> = mutableStateOf("")
	val errors: MutableState<Map<String, String>> = mutableStateOf(emptyMap())
	fun setError(field: String, message: String?) {
		val current = errors.value.toMutableMap()
		if (message == null) current.remove(field) else current[field] = message
		errors.value = current
	}
	fun toDto(): CategoryDto = CategoryDto(
		uniqueId = if (uniqueId.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(uniqueId.value)),
		name = name.value,
	)
	fun fromDto(dto: CategoryDto) {
		uniqueId.value = dto.uniqueId.toDisplayString()
		name.value = dto.name.toString()
	}
}