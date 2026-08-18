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
class ProductOptionalDtoFormState {
	var uniqueId: MutableState<String> = mutableStateOf("")
	var title: MutableState<String> = mutableStateOf("")
	var sku: MutableState<String> = mutableStateOf("")
	var active: MutableState<String> = mutableStateOf("")
	var featured: MutableState<String> = mutableStateOf("")
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
	var specs: MutableState<String> = mutableStateOf("")
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
	fun toDto(): ProductOptionalDto = ProductOptionalDto(
		uniqueId = if (uniqueId.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(uniqueId.value)),
		title = if (title.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(title.value)),
		sku = if (sku.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(sku.value)),
		active = active.value.toBooleanStrictOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		featured = featured.value.toBooleanStrictOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		viewCount = viewCount.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		viewCountOpt = viewCountOpt.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		smallCount = smallCount.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		smallCountOpt = smallCountOpt.value.toIntOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		bigCount = bigCount.value.toLongOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		bigCountOpt = bigCountOpt.value.toLongOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio32 = ratio32.value.toFloatOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio32Opt = ratio32Opt.value.toFloatOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio64 = ratio64.value.toDoubleOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		ratio64Opt = ratio64Opt.value.toDoubleOrNull()?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent),
		status = if (status.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(status.value)),
		statusOpt = if (statusOpt.value.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(statusOpt.value)),
		tags = if (tags.value.isEmpty()) MaybeField<List<String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<String>>(tags.value))) }.getOrElse { MaybeField<List<String>>(Maybe.Absent) },
		tagsOpt = if (tagsOpt.value.isEmpty()) MaybeField<List<String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<String>>(tagsOpt.value))) }.getOrElse { MaybeField<List<String>>(Maybe.Absent) },
		labelsMeta = if (labelsMeta.value.isEmpty()) MaybeField<Map<String, String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<Map<String, String>>(labelsMeta.value))) }.getOrElse { MaybeField<Map<String, String>>(Maybe.Absent) },
		labelsMetaOpt = if (labelsMetaOpt.value.isEmpty()) MaybeField<Map<String, String>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<Map<String, String>>(labelsMetaOpt.value))) }.getOrElse { MaybeField<Map<String, String>>(Maybe.Absent) },
		gallery = if (gallery.value.isEmpty()) MaybeField<List<ProductOptionalDtoGallery>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<ProductOptionalDtoGallery>>(gallery.value))) }.getOrElse { MaybeField<List<ProductOptionalDtoGallery>>(Maybe.Absent) },
		galleryOpt = if (galleryOpt.value.isEmpty()) MaybeField<List<ProductOptionalDtoGalleryOpt>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<ProductOptionalDtoGalleryOpt>>(galleryOpt.value))) }.getOrElse { MaybeField<List<ProductOptionalDtoGalleryOpt>>(Maybe.Absent) },
		specs = if (specs.value.isEmpty()) MaybeField<ProductOptionalDtoSpecs>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<ProductOptionalDtoSpecs>(specs.value))) }.getOrElse { MaybeField<ProductOptionalDtoSpecs>(Maybe.Absent) },
		price = Json.decodeFromString<Money>(price.value),
		misc = misc.value,
		category = if (category.value.isEmpty()) MaybeField<CategoryDto>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<CategoryDto>(category.value))) }.getOrElse { MaybeField<CategoryDto>(Maybe.Absent) },
		invoiceLines = if (invoiceLines.value.isEmpty()) MaybeField<List<InvoiceLineDto>>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<List<InvoiceLineDto>>(invoiceLines.value))) }.getOrElse { MaybeField<List<InvoiceLineDto>>(Maybe.Absent) },
	)
	fun fromDto(dto: ProductOptionalDto) {
		uniqueId.value = dto.uniqueId.toDisplayString()
		title.value = dto.title.toDisplayString()
		sku.value = dto.sku.toDisplayString()
		active.value = dto.active.toDisplayString()
		featured.value = dto.featured.toDisplayString()
		viewCount.value = dto.viewCount.toDisplayString()
		viewCountOpt.value = dto.viewCountOpt.toDisplayString()
		smallCount.value = dto.smallCount.toDisplayString()
		smallCountOpt.value = dto.smallCountOpt.toDisplayString()
		bigCount.value = dto.bigCount.toDisplayString()
		bigCountOpt.value = dto.bigCountOpt.toDisplayString()
		ratio32.value = dto.ratio32.toDisplayString()
		ratio32Opt.value = dto.ratio32Opt.toDisplayString()
		ratio64.value = dto.ratio64.toDisplayString()
		ratio64Opt.value = dto.ratio64Opt.toDisplayString()
		status.value = dto.status.toDisplayString()
		statusOpt.value = dto.statusOpt.toDisplayString()
		tags.value = dto.tags.toDisplayString()
		tagsOpt.value = dto.tagsOpt.toDisplayString()
		labelsMeta.value = dto.labelsMeta.toDisplayString()
		labelsMetaOpt.value = dto.labelsMetaOpt.toDisplayString()
		gallery.value = dto.gallery.toDisplayString()
		galleryOpt.value = dto.galleryOpt.toDisplayString()
		specs.value = dto.specs.toDisplayString()
		price.value = Json.encodeToString(dto.price)
		misc.value = dto.misc.toString()
		category.value = dto.category.toDisplayString()
		invoiceLines.value = dto.invoiceLines.toDisplayString()
	}
}