package php

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitive maps a leaf emi primitive type onto its PHP type. Returns
// "" when the field isn't a primitive, so callers fall through to
// structural resolution (array, object, one, collection...). enum is
// handled separately (see phpEnumBaseType) since it needs the full field,
// not just the type string.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "string"
	case "int", "int32", "int64", "int32?", "int64?", "int?":
		return "int"
	case "float32", "float64", "float32?", "float64?":
		return "float"
	case "bool", "bool?":
		return "bool"
	case "any", "complex", "any?", "complex?":
		return "mixed"
	default:
		return ""
	}
}

// phpDataStructureType resolves array/_list/slice/object/one/collection/
// class fields to their PHP type. `array`/`_list`/`slice`/`collection` are
// all represented as plain `array` (PHP has no generics - see the PHPDoc
// `@var X[]` annotation generated alongside, and Hydrator::arrayItemTypes()
// for the runtime-facing equivalent). `one`/`class`/`collection` always
// reference *another* generated class by name (`target:`) - nestedClassName
// is only ever a fallback for a malformed definition missing that target,
// and one that's never actually emitted anywhere (only object/array/_list
// fields get a flattened companion class), so single relations fall back to
// `mixed` instead of referencing a name that was never generated.
func phpDataStructureType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeOne, core.FieldTypeClass, core.FieldTypeOneNullable, core.FieldTypeClassNullable:
		if field.Target != "" {
			return field.Target
		}
		return "mixed"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable,
		core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable,
		core.FieldTypeSlice, core.FieldTypeSliceNullable,
		core.FieldTypeMap, core.FieldTypeMapNullable:
		return "array"
	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return nestedClassName
	default:
		return ""
	}
}

// phpArrayElementType computes the PHPDoc element type for an
// array/_list/slice/collection/map field - purely documentation/tooling
// (see phpArrayItemClass for the runtime-facing equivalent used by the
// Hydrator to know which elements need object coercion).
func phpArrayElementType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Target != "" {
			return field.Target
		}
		return "mixed"
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return nestedClassName
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "mixed"
		}
		return primitive
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		value := extractPrimitive(field.MapPairOf)
		if value == "" {
			value = "mixed"
		}
		return fmt.Sprintf("string, %v", value) // rendered as array<string, V>
	default:
		return "mixed"
	}
}

// phpArrayItemClass returns the fully-generated class name array elements
// should be coerced to during hydration - "" when elements are scalars (no
// coercion needed, json_decode's native type already matches).
func phpArrayItemClass(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		return field.Target
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return nestedClassName
	default:
		return ""
	}
}

// phpEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to
// a module-level `enums:` entry, or inline `of:` values (a flattened,
// generated companion backed-enum named `<prefixName><FieldName>`). Falls
// back to plain `string` with neither.
func phpEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	if len(field.OfType) > 0 {
		return prefixName + core.ToUpper(field.Name)
	}
	return "string"
}

// phpFieldType computes the fully-resolved PHP type hint for a field,
// including the `?` prefix for nullable emi types. nestedClassName is the
// PascalCase name already computed for this field, in case it's an
// object/array of objects.
func phpFieldType(field *core.EmiField, nestedClassName string, prefixName string) string {
	if field == nil {
		return "mixed"
	}

	if field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable {
		base := phpEnumBaseType(field, prefixName)
		if core.IsNullable(string(field.Type)) {
			return "?" + base
		}
		return base
	}

	base := extractPrimitive(string(field.Type))
	if base == "" {
		base = phpDataStructureType(field, nestedClassName)
	}
	if base == "" {
		base = "mixed"
	}
	if base == "mixed" {
		return base // `?mixed` is redundant - mixed already includes null
	}

	// A single `one`/`class` relation is never eagerly default-constructed
	// (see phpFieldDefault) even when the schema marks it required, so its
	// type always has to say nullable too.
	isSingleRelation := field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass
	if core.IsNullable(string(field.Type)) || isSingleRelation {
		return "?" + base
	}

	return base
}
