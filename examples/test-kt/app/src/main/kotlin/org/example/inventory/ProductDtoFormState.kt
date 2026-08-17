package org.example.inventory
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.mutableStateOf
import emikot.Maybe
import emikot.MaybeField
import emikot.toDisplayString
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import org.example.billing.InvoiceLineDto
import org.example.money.Money
class ProductDtoFormState {
	var uniqueId: MutableState<String> = mutableStateOf("")
	var title: MutableState<String> = mutableStateOf("")
	var sku: MutableState<String> = mutableStateOf("")
	var active: MutableState<String> = mutableStateOf("false")
	var featured: MutableState<String> = mutableStateOf("false")
	var viewCount: MutableState<String> = mutableStateOf("")
	var viewCountOpt: MutableState<String> = mutableStateOf("")
	var smallCount: MutableState<String> = mutableStateOf("")
	var smallCountOpt: MutableState<String> = mutableStateOf("")
	var bigCount: MutableState<String> = mutableStateOf("")
	var bigCountOpt: MutableState<String> = mutableStateOf("")
	var ratio32: MutableState<String> = mutableStateOf("")
	var ratio32Opt: MutableState<String> = mutableStateOf("")
	var ratio64: MutableState<String> = mutableStateOf("")
	var ratio64Opt: MutableState<String> = mutableStateOf("")
	var status: MutableState<String> = mutableStateOf("")
	var statusOpt: MutableState<String> = mutableStateOf("")
	var tags: MutableState<String> = mutableStateOf("")
	var tagsOpt: MutableState<String> = mutableStateOf("")
	var labelsMeta: MutableState<String> = mutableStateOf("")
	var labelsMetaOpt: MutableState<String> = mutableStateOf("")
	var gallery: MutableState<String> = mutableStateOf("")
	var galleryOpt: MutableState<String> = mutableStateOf("")
	val specs = ProductDtoFormStateSpecs()
	var price: MutableState<String> = mutableStateOf("")
	var misc: MutableState<String> = mutableStateOf("")
	var category: MutableState<String> = mutableStateOf("")
	var invoiceLines: MutableState<String> = mutableStateOf("")
	val errors: MutableState<Map<String, String>> = mutableStateOf(emptyMap())
	fun setError(field: String, message: String?) {
		val current = errors.value.toMutableMap()
		if (message == null) current.remove(field) else current[field] = message
		errors.value = current
	}
	fun toDto(): ProductDto = ProductDto(
		uniqueId = if (uniqueId.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(uniqueId.value)),
		title = title.value,
		sku = if (sku.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(sku.value)),
		active = active.value.toBooleanStrictOrNull() ?: false,
		featured = featured.value.toBooleanStrictOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		viewCount = viewCount.value.toIntOrNull() ?: 0,
		viewCountOpt = viewCountOpt.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		smallCount = smallCount.value.toIntOrNull() ?: 0,
		smallCountOpt = smallCountOpt.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		bigCount = bigCount.value.toLongOrNull() ?: 0,
		bigCountOpt = bigCountOpt.value.toLongOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio32 = ratio32.value.toFloatOrNull() ?: 0.0f,
		ratio32Opt = ratio32Opt.value.toFloatOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio64 = ratio64.value.toDoubleOrNull() ?: 0.0,
		ratio64Opt = ratio64Opt.value.toDoubleOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		status = status.value,
		statusOpt = if (statusOpt.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(statusOpt.value)),
		tags = runCatching { Json.decodeFromString<List<String>>(tags.value) }.getOrElse { emptyList() },
		tagsOpt = if (tagsOpt.value.isEmpty()) MaybeField<List<String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<String>>(tagsOpt.value))) }.getOrElse { MaybeField<List<String>>(Maybe.Absent) },
		labelsMeta = runCatching { Json.decodeFromString<Map<String, String>>(labelsMeta.value) }.getOrElse { emptyMap() },
		labelsMetaOpt = if (labelsMetaOpt.value.isEmpty()) MaybeField<Map<String, String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<Map<String, String>>(labelsMetaOpt.value))) }.getOrElse { MaybeField<Map<String, String>>(Maybe.Absent) },
		gallery = runCatching { Json.decodeFromString<List<ProductDtoGallery>>(gallery.value) }.getOrElse { emptyList() },
		galleryOpt = if (galleryOpt.value.isEmpty()) MaybeField<List<ProductDtoGalleryOpt>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<ProductDtoGalleryOpt>>(galleryOpt.value))) }.getOrElse { MaybeField<List<ProductDtoGalleryOpt>>(Maybe.Absent) },
		specs = specs.toDto(),
		price = Json.decodeFromString<Money>(price.value),
		misc = misc.value,
		category = if (category.value.isEmpty()) MaybeField<CategoryDto>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<CategoryDto>(category.value))) }.getOrElse { MaybeField<CategoryDto>(Maybe.Absent) },
		invoiceLines = if (invoiceLines.value.isEmpty()) MaybeField<List<InvoiceLineDto>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<InvoiceLineDto>>(invoiceLines.value))) }.getOrElse { MaybeField<List<InvoiceLineDto>>(Maybe.Absent) },
	)
	fun fromDto(dto: ProductDto) {
		uniqueId.value = dto.uniqueId.toDisplayString()
		title.value = dto.title.toString()
		sku.value = dto.sku.toDisplayString()
		active.value = dto.active.toString()
		featured.value = dto.featured.toDisplayString()
		viewCount.value = dto.viewCount.toString()
		viewCountOpt.value = dto.viewCountOpt.toDisplayString()
		smallCount.value = dto.smallCount.toString()
		smallCountOpt.value = dto.smallCountOpt.toDisplayString()
		bigCount.value = dto.bigCount.toString()
		bigCountOpt.value = dto.bigCountOpt.toDisplayString()
		ratio32.value = dto.ratio32.toString()
		ratio32Opt.value = dto.ratio32Opt.toDisplayString()
		ratio64.value = dto.ratio64.toString()
		ratio64Opt.value = dto.ratio64Opt.toDisplayString()
		status.value = dto.status.toString()
		statusOpt.value = dto.statusOpt.toDisplayString()
		tags.value = Json.encodeToString(dto.tags)
		tagsOpt.value = dto.tagsOpt.toDisplayString()
		labelsMeta.value = Json.encodeToString(dto.labelsMeta)
		labelsMetaOpt.value = dto.labelsMetaOpt.toDisplayString()
		gallery.value = Json.encodeToString(dto.gallery)
		galleryOpt.value = dto.galleryOpt.toDisplayString()
		specs.fromDto(dto.specs)
		price.value = Json.encodeToString(dto.price)
		misc.value = dto.misc.toString()
		category.value = dto.category.toDisplayString()
		invoiceLines.value = dto.invoiceLines.toDisplayString()
	}
}
class ProductDtoFormStateSpecs {
	var weightGrams: MutableState<String> = mutableStateOf("")
	val errors: MutableState<Map<String, String>> = mutableStateOf(emptyMap())
	fun setError(field: String, message: String?) {
		val current = errors.value.toMutableMap()
		if (message == null) current.remove(field) else current[field] = message
		errors.value = current
	}
	fun toDto(): ProductDtoSpecs = ProductDtoSpecs(
		weightGrams = weightGrams.value.toIntOrNull() ?: 0,
	)
	fun fromDto(dto: ProductDtoSpecs) {
		weightGrams.value = dto.weightGrams.toString()
	}
}