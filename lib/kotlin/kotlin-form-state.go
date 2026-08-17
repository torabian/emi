package kotlin

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// formFieldKind buckets a field's *base* type (nullability stripped, exactly like
// kotlinDataStructureType's own normalization) into how KotlinFormStateGenerator
// converts it to/from its Compose-bindable MutableState<String>. Every kind other than
// the plain scalars round-trips through kotlinx.serialization JSON text - a single text
// field is a basic but genuinely working edit surface for a list/map/relation/complex
// value, and keeps toDto()/fromDto() uniform instead of needing bespoke Compose widgets
// per structured type.
type formFieldKind string

const (
	formFieldKindString     formFieldKind = "string"
	formFieldKindInt        formFieldKind = "int"
	formFieldKindLong       formFieldKind = "long"
	formFieldKindFloat      formFieldKind = "float"
	formFieldKindDouble     formFieldKind = "double"
	formFieldKindBool       formFieldKind = "bool"
	// formFieldKindAny is `type: any` specifically - it resolves to a plain Kotlin
	// `Any` on the dto (see kotlinDataStructureType's final fallback), which
	// kotlinx.serialization can't en/decode without a registered contextual serializer
	// (the same reason the dto field itself is @Contextual), so it can't go through
	// the same Json.decodeFromString/encodeToString round trip formFieldKindStructured
	// uses - it's passed through as a plain string instead.
	formFieldKindAny        formFieldKind = "any"
	formFieldKindStructured formFieldKind = "structured"
)

func formFieldKindOf(field *core.EmiField) formFieldKind {
	base := core.FieldType(strings.TrimSuffix(string(field.Type), "?"))
	switch base {
	case core.FieldTypeString, core.FieldTypeEnum:
		return formFieldKindString
	case core.FieldTypeInt, core.FieldTypeInt32:
		return formFieldKindInt
	case core.FieldTypeInt64:
		return formFieldKindLong
	case core.FieldTypeFloat32:
		return formFieldKindFloat
	case core.FieldTypeFloat64:
		return formFieldKindDouble
	case core.FieldTypeBool:
		return formFieldKindBool
	case core.FieldTypeAny:
		return formFieldKindAny
	default:
		return formFieldKindStructured
	}
}

// isFormStateNestedObject reports whether a field gets its own nested
// <Prefix><Field>FormState sub-holder instead of a MutableState<String>. Scoped to
// plain (non-nullable) "object" fields with children - the common embedded-object
// case (e.g. examples/test-kt's ProfileDto.address). A nullable "object?" field is
// MaybeField<NestedType> on the dto, which doesn't have the "always constructible"
// property a sub-holder's own no-arg instance needs, so it's treated as a structured
// scalar (JSON round trip) like map/array/one/collection/complex instead.
func isFormStateNestedObject(field *core.EmiField) bool {
	return field.Type == core.FieldTypeObject && len(field.Fields) > 0
}

// formStateFieldName is the Kotlin property name for a field - matches renderField's
// own convention (core.ToLower(field.PublicName())) so it lines up 1:1 with the dto's
// own generated property.
func formStateFieldName(field *core.EmiField) string {
	return core.ToLower(field.PublicName())
}

// formStateConversion returns the toDto()/fromDto() expressions for a non-object-
// container field. computedType is the field's actual computed Kotlin dto type (from
// goFieldTypeOnNestedClasses(field, parentChain) - the same resolution renderField
// itself uses, so a structured field's bareType below matches the real generated
// nested class, e.g. "List<ProductDtoGallery>", not a guessed/unqualified name).
// stateAccessor is the Kotlin expression for the field's MutableState<String> value
// (e.g. "name.value"); dtoAccessor is the Kotlin expression for the dto's own field
// value (e.g. "dto.name").
func formStateConversion(field *core.EmiField, computedType, stateAccessor, dtoAccessor string) (toDto string, fromDto string) {
	nullable := core.IsNullable(string(field.Type))
	kind := formFieldKindOf(field)

	parseExpr := func(kind formFieldKind, text string) string {
		switch kind {
		case formFieldKindInt:
			return text + ".toIntOrNull()"
		case formFieldKindLong:
			return text + ".toLongOrNull()"
		case formFieldKindFloat:
			return text + ".toFloatOrNull()"
		case formFieldKindDouble:
			return text + ".toDoubleOrNull()"
		case formFieldKindBool:
			return text + ".toBooleanStrictOrNull()"
		default:
			return text
		}
	}

	if kind == formFieldKindAny {
		return stateAccessor, fmt.Sprintf("%s.toString()", dtoAccessor)
	}

	if kind == formFieldKindString {
		if nullable {
			return fmt.Sprintf("if (%s.isEmpty()) MaybeField(Maybe.Absent) else MaybeField(Maybe.Value(%s))", stateAccessor, stateAccessor),
				fmt.Sprintf("%s.toDisplayString()", dtoAccessor)
		}
		return stateAccessor, fmt.Sprintf("%s.toString()", dtoAccessor)
	}

	if kind == formFieldKindInt || kind == formFieldKindLong || kind == formFieldKindFloat || kind == formFieldKindDouble || kind == formFieldKindBool {
		parsed := parseExpr(kind, stateAccessor)
		if nullable {
			return fmt.Sprintf("%s?.let { MaybeField(Maybe.Value(it)) } ?: MaybeField(Maybe.Absent)", parsed),
				fmt.Sprintf("%s.toDisplayString()", dtoAccessor)
		}
		// KotlinSafeDefaultValue always has a real default for every non-nullable
		// scalar kind handled here (bool/int/int32/int64/float32/float64) - see
		// kotlin-field.go.
		defaultLiteral := KotlinSafeDefaultValue(field)
		return fmt.Sprintf("%s ?: %s", parsed, defaultLiteral),
			fmt.Sprintf("%s.toString()", dtoAccessor)
	}

	// Structured (map/slice/array/one/collection/complex): JSON round trip. bareType
	// is computedType with its MaybeField<...> wrapper stripped when nullable -
	// exactly the T a nullable field's MaybeField<T> wraps.
	bareType := computedType
	if nullable {
		bareType = strings.TrimSuffix(strings.TrimPrefix(computedType, "MaybeField<"), ">")
	}
	if bareType == "" {
		bareType = "Any"
	}

	if nullable {
		// MaybeField<T>'s T is invariant (unlike Maybe<out T>), so a bare
		// MaybeField(Maybe.Absent) inside a getOrElse {} lambda can't be unified
		// against the runCatching block's inferred MaybeField<bareType> - Kotlin
		// treats them as unrelated captured types. Spelling out <%s> on every
		// MaybeField(...) construction here removes the ambiguity entirely.
		toDto := fmt.Sprintf(
			"if (%s.isEmpty()) MaybeField<%s>(Maybe.Absent) else runCatching { MaybeField(Maybe.Value(Json.decodeFromString<%s>(%s))) }.getOrElse { MaybeField<%s>(Maybe.Absent) }",
			stateAccessor, bareType, bareType, stateAccessor, bareType,
		)
		return toDto, fmt.Sprintf("%s.toDisplayString()", dtoAccessor)
	}

	// Not every non-nullable structured type has a synthesizable safe default (e.g. a
	// required "one" relation, or a "complex" value like Money) - KotlinSafeDefaultValue
	// returns "" for those, so skip the runCatching/getOrElse wrapper entirely rather
	// than emit a syntactically-empty fallback block; a malformed value in the text
	// field throws at conversion time instead, which is the only sane behavior when
	// there's genuinely no default to fall back to.
	if safeDefault := KotlinSafeDefaultValue(field); safeDefault != "" {
		toDto = fmt.Sprintf("runCatching { Json.decodeFromString<%s>(%s) }.getOrElse { %s }", bareType, stateAccessor, safeDefault)
	} else {
		toDto = fmt.Sprintf("Json.decodeFromString<%s>(%s)", bareType, stateAccessor)
	}
	fromDto = fmt.Sprintf("Json.encodeToString(%s)", dtoAccessor)
	return toDto, fromDto
}

// formStateInitialValue is the mutableStateOf(...) seed for a fresh (no dto yet) form.
func formStateInitialValue(field *core.EmiField) string {
	switch formFieldKindOf(field) {
	case formFieldKindBool:
		return `mutableStateOf("false")`
	default:
		return `mutableStateOf("")`
	}
}

// KotlinFormStateGenerator renders a Compose-friendly <Dto>FormState class for a dto's
// fields - one mutableStateOf-backed MutableState<String> per scalar/structured field
// (see formFieldKind), a nested sub-holder per plain "object" field (see
// isFormStateNestedObject), a per-field validation error slot, and toDto()/fromDto()
// converters - so a caller building an Android form already has a state holder that
// plugs directly into the action that consumes the dto (see kotlin-action-render.go).
// Only emitted when the "android-forms" tag is set - see KotlinModuleFull. complexes
// is the same module-wide recognized-complexes list KotlinCommonStructGenerator
// already receives - a form-state field can decode a "complex" field's real class
// (e.g. Money) or another dto/entity's "target:" class exactly like the dto itself
// does, so it needs the identical import resolution.
func KotlinFormStateGenerator(fields []*core.EmiField, dtoClassName string, complexes []RecognizedComplex) (*core.CodeChunkCompiled, error) {
	rootPrefix := dtoClassName + "FormState"

	var b strings.Builder
	if err := renderFormStateClass(&b, fields, dtoClassName, rootPrefix); err != nil {
		return nil, err
	}

	var extraDeps []core.CodeChunkDependency
	goctx := commonClassContext{RecognizedComplexes: complexes}
	for _, item := range CollectComplexClasses(fields) {
		if location := findComplexLocation(item, goctx); location != "" {
			extraDeps = append(extraDeps, core.CodeChunkDependency{Location: location})
		}
	}
	extraDeps = append(extraDeps, kotlinCollectTargetDeps(fields, dtoClassName)...)

	deps := []core.CodeChunkDependency{
		{Location: "androidx.compose.runtime.MutableState"},
		{Location: "androidx.compose.runtime.mutableStateOf"},
		{Location: "kotlinx.serialization.json.Json"},
		{Location: "kotlinx.serialization.encodeToString"},
		{Location: "kotlinx.serialization.decodeFromString"},
		{Location: "emikot.MaybeField"},
		{Location: "emikot.Maybe"},
		{Location: "emikot.toDisplayString"},
	}
	deps = append(deps, extraDeps...)

	return &core.CodeChunkCompiled{
		ActualScript:          []byte(b.String()),
		SuggestedFileName:     rootPrefix,
		SuggestedExtension:    ".kt",
		CodeChunkDependensies: deps,
	}, nil
}

// renderFormStateClass writes one <prefix> class (root or nested) and recurses into
// every plain-object field's own nested class.
func renderFormStateClass(b *strings.Builder, fields []*core.EmiField, dtoClassName string, prefix string) error {
	type plan struct {
		field       *core.EmiField
		name        string
		isObject    bool
		nestedClass string
		toDto       string
		fromDto     string
	}

	var plans []plan
	var nested []*core.EmiField

	for _, f := range fields {
		if f == nil {
			continue
		}
		name := formStateFieldName(f)

		if isFormStateNestedObject(f) {
			nestedClass := prefix + core.ToUpper(f.Name)
			plans = append(plans, plan{field: f, name: name, isObject: true, nestedClass: nestedClass})
			nested = append(nested, f)
			continue
		}

		// dtoClassName doubles as the correct parentChain for type resolution here:
		// both the dto's own class-naming scheme (renderClasses) and
		// goFieldTypeOnNestedClasses accumulate nested prefixes identically
		// (parent + PascalCase(field name)), so the same string that names this
		// level's dto class is also the parentChain a field at this level resolves
		// against.
		computedType := goFieldTypeOnNestedClasses(f, dtoClassName)
		toDto, fromDto := formStateConversion(f, computedType, name+".value", "dto."+name)
		plans = append(plans, plan{field: f, name: name, toDto: toDto, fromDto: fromDto})
	}

	fmt.Fprintf(b, "class %s {\r\n", prefix)
	for _, p := range plans {
		if p.isObject {
			fmt.Fprintf(b, "\tval %s = %s()\r\n", p.name, p.nestedClass)
			continue
		}
		fmt.Fprintf(b, "\tvar %s: MutableState<String> = %s\r\n", p.name, formStateInitialValue(p.field))
	}
	b.WriteString("\r\n\tval errors: MutableState<Map<String, String>> = mutableStateOf(emptyMap())\r\n\r\n")
	b.WriteString("\tfun setError(field: String, message: String?) {\r\n")
	b.WriteString("\t\tval current = errors.value.toMutableMap()\r\n")
	b.WriteString("\t\tif (message == null) current.remove(field) else current[field] = message\r\n")
	b.WriteString("\t\terrors.value = current\r\n")
	b.WriteString("\t}\r\n\r\n")

	fmt.Fprintf(b, "\tfun toDto(): %s = %s(\r\n", dtoClassName, dtoClassName)
	for _, p := range plans {
		if p.isObject {
			fmt.Fprintf(b, "\t\t%s = %s.toDto(),\r\n", p.name, p.name)
			continue
		}
		fmt.Fprintf(b, "\t\t%s = %s,\r\n", p.name, p.toDto)
	}
	b.WriteString("\t)\r\n\r\n")

	fmt.Fprintf(b, "\tfun fromDto(dto: %s) {\r\n", dtoClassName)
	for _, p := range plans {
		if p.isObject {
			fmt.Fprintf(b, "\t\t%s.fromDto(dto.%s)\r\n", p.name, p.name)
			continue
		}
		fmt.Fprintf(b, "\t\t%s.value = %s\r\n", p.name, p.fromDto)
	}
	b.WriteString("\t}\r\n")
	b.WriteString("}\r\n\r\n")

	for _, f := range nested {
		nestedClass := prefix + core.ToUpper(f.Name)
		nestedDtoClass := dtoClassName + core.ToUpper(f.Name)
		if err := renderFormStateClass(b, f.Fields, nestedDtoClass, nestedClass); err != nil {
			return err
		}
	}

	return nil
}
