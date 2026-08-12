package dart

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitive maps a leaf emi primitive type onto its dart type. Returns
// "" when the field isn't a primitive, so callers fall through to structural
// resolution (array, object, one, collection...). enum is handled separately
// (see dartEnumBaseType) since it needs the full field, not just the type
// string.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "String"
	case "int", "int32", "int64", "int?", "int32?", "int64?":
		return "int"
	case "float32", "float64", "float32?", "float64?":
		return "double"
	case "bool", "bool?":
		return "bool"
	case "any", "complex", "any?", "complex?":
		return "dynamic"
	default:
		return ""
	}
}

// dartMapType resolves `Map<String, V>` for `map`/`map?` fields. JSON object
// keys are always strings on the wire regardless of what `mapKeyOf` claims,
// so the key type is fixed; only the value type varies.
func dartMapType(field *core.EmiField) string {
	value := extractPrimitive(field.MapPairOf)
	if value == "" {
		value = "dynamic"
	}
	return fmt.Sprintf("Map<String, %v>", value)
}

// dartDataStructureType resolves array/_list/slice/object/one/collection/
// class fields to their dart type, given the already-computed name of the
// nested/target class. `one`/`class`/`collection` always reference *another*
// generated class by name (`target:`) - nestedClassName is only ever a
// fallback for a malformed definition missing that target, and one that's
// never actually emitted anywhere (only object/array/_list fields get a
// flattened companion class), so it deliberately falls back to `dynamic`
// instead - referencing a name that was never generated would fail to
// compile.
func dartDataStructureType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeOne, core.FieldTypeClass, core.FieldTypeOneNullable, core.FieldTypeClassNullable:
		if field.Target != "" {
			return field.Target
		}
		return "dynamic"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Target != "" {
			return fmt.Sprintf("List<%v>", field.Target)
		}
		return "List<dynamic>"
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return fmt.Sprintf("List<%v>", nestedClassName)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "dynamic"
		}
		return fmt.Sprintf("List<%v>", primitive)
	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return nestedClassName
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return dartMapType(field)
	default:
		return ""
	}
}

// dartEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to
// a module-level `enums:` entry, or inline `of:` values (a flattened,
// generated companion enum named `<prefixName><FieldName>` - see
// collectInlineEnums). Falls back to plain `String` with neither.
func dartEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	if len(field.OfType) > 0 {
		return prefixName + core.ToUpper(field.Name)
	}
	return "String"
}

// dartFieldType computes the fully-resolved dart type for a field, including
// the `?` suffix for nullable emi types. nestedClassName is the PascalCase
// name already computed for this field, in case it's an object/array of
// objects.
func dartFieldType(field *core.EmiField, nestedClassName string, prefixName string) string {
	if field == nil {
		return "dynamic"
	}

	if field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable {
		base := dartEnumBaseType(field, prefixName)
		if core.IsNullable(string(field.Type)) {
			return base + "?"
		}
		return base
	}

	base := extractPrimitive(string(field.Type))
	if base == "" {
		base = dartDataStructureType(field, nestedClassName)
	}
	if base == "" {
		base = "dynamic"
	}

	// A single `one`/`class` relation is never eagerly default-constructed
	// (see dartFieldDefault) even when the schema marks it required, so its
	// type always has to say nullable too.
	isSingleRelation := field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass
	if core.IsNullable(string(field.Type)) || isSingleRelation || base == "dynamic" {
		if base == "dynamic" {
			return base // dynamic is implicitly nullable, `dynamic?` is redundant
		}
		return base + "?"
	}

	return base
}
