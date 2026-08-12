package csharp

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitive maps a leaf emi primitive type onto its C# type. Returns
// "" when the field isn't a primitive, so callers fall through to
// structural resolution (array, object, one, collection...). enum is
// handled separately (see csEnumBaseType) since it needs the full field,
// not just the type string.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "string"
	case "int", "int32", "int32?":
		return "int"
	case "int64", "int64?":
		return "long"
	case "float32", "float32?":
		return "float"
	case "float64", "float64?":
		return "double"
	case "bool", "bool?":
		return "bool"
	case "any", "complex", "any?", "complex?":
		return "object"
	default:
		return ""
	}
}

// isValueType reports whether a bare (non-generic, non-reference) C# type
// csNullable appends the correct nullable suffix for the given C# type -
// value types need `?`, `object`/`dynamic`/reference types already accept
// null without one (and `object?` too, for consistency/clarity, is fine to
// use explicitly for reference types).
func csNullable(csType string) string {
	if csType == "" {
		return "object?"
	}
	return csType + "?"
}

func csMapType(field *core.EmiField) string {
	value := extractPrimitive(field.MapPairOf)
	if value == "" {
		value = "object?"
	}
	return fmt.Sprintf("Dictionary<string, %v>", value)
}

// csDataStructureType resolves array/_list/slice/object/one/collection/class
// fields to their C# type, given the already-computed name of the
// nested/target class. `one`/`class`/`collection` always reference *another*
// generated class by name (`target:`) - nestedClassName is only ever a
// fallback for a malformed definition missing that target, and one that's
// never actually emitted anywhere (only object/array/_list fields get a
// flattened companion class), so it falls back to `object` instead of
// referencing a name that was never generated.
func csDataStructureType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeOne, core.FieldTypeClass, core.FieldTypeOneNullable, core.FieldTypeClassNullable:
		if field.Target != "" {
			return field.Target
		}
		return "object"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Target != "" {
			return fmt.Sprintf("List<%v>", field.Target)
		}
		return "List<object?>"
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return fmt.Sprintf("List<%v>", nestedClassName)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "object?"
		}
		return fmt.Sprintf("List<%v>", primitive)
	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return nestedClassName
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return csMapType(field)
	default:
		return ""
	}
}

// csEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to
// a module-level `enums:` entry, or inline `of:` values (a flattened,
// generated companion enum named `<prefixName><FieldName>`). Falls back to
// plain `string` with neither.
func csEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	if len(field.OfType) > 0 {
		return prefixName + core.ToUpper(field.Name)
	}
	return "string"
}

// csFieldType computes the fully-resolved C# type for a field, including the
// `?` suffix for nullable emi types. nestedClassName is the PascalCase name
// already computed for this field, in case it's an object/array of objects.
func csFieldType(field *core.EmiField, nestedClassName string, prefixName string) string {
	if field == nil {
		return "object?"
	}

	if field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable {
		base := csEnumBaseType(field, prefixName)
		if core.IsNullable(string(field.Type)) {
			return csNullable(base)
		}
		return base
	}

	base := extractPrimitive(string(field.Type))
	if base == "" {
		base = csDataStructureType(field, nestedClassName)
	}
	if base == "" {
		base = "object"
	}

	// A single `one`/`class` relation is never eagerly default-constructed
	// (see csFieldDefault) even when the schema marks it required, so its
	// type always has to say nullable too. Likewise a bare `object` fallback
	// (any/complex/unresolved target) has no sensible non-null default
	// under nullable-reference-types, regardless of what the schema claims -
	// forcing it non-null would leave the property uninitialized and the
	// compiler would rightly flag it (CS8618).
	isSingleRelation := field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass
	if core.IsNullable(string(field.Type)) || isSingleRelation || base == "object" {
		return csNullable(base)
	}

	return base
}
