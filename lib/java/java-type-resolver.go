package java

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitive maps a leaf emi primitive type onto its Java (boxed, so
// every field can be assigned `null`) type. Returns "" when the field isn't
// a primitive, so callers fall through to structural resolution (array,
// object, one, collection...). enum is handled separately (see
// javaEnumBaseType) since it needs the full field, not just the type
// string.
//
// Java has no `T?` syntax (unlike Dart/C#) - any reference type already
// accepts null, so nullability here only ever affects whether a *default
// value* is emitted (see javaFieldDefault), never the type string itself.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "String"
	case "int", "int32", "int32?":
		return "Integer"
	case "int64", "int64?":
		return "Long"
	case "float32", "float32?":
		return "Float"
	case "float64", "float64?":
		return "Double"
	case "bool", "bool?":
		return "Boolean"
	case "any", "complex", "any?", "complex?":
		return "Object"
	default:
		return ""
	}
}

func javaMapType(field *core.EmiField) string {
	value := extractPrimitive(field.MapPairOf)
	if value == "" {
		value = "Object"
	}
	return fmt.Sprintf("Map<String, %v>", value)
}

// javaDataStructureType resolves array/_list/slice/object/one/collection/
// class fields to their Java type, given the already-computed name of the
// nested/target class. `one`/`class`/`collection` always reference *another*
// generated class by name (`target:`) - nestedClassName is only ever a
// fallback for a malformed definition missing that target, and one that's
// never actually emitted anywhere (only object/array/_list fields get a
// flattened companion class), so it falls back to `Object` instead of
// referencing a name that was never generated.
func javaDataStructureType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeOne, core.FieldTypeClass, core.FieldTypeOneNullable, core.FieldTypeClassNullable:
		if field.Target != "" {
			return field.Target
		}
		return "Object"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Target != "" {
			return fmt.Sprintf("List<%v>", field.Target)
		}
		return "List<Object>"
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		core.FieldTypeList, core.FieldTypeListNullable:
		return fmt.Sprintf("List<%v>", nestedClassName)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "Object"
		}
		return fmt.Sprintf("List<%v>", primitive)
	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return nestedClassName
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return javaMapType(field)
	default:
		return ""
	}
}

// javaEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to
// a module-level `enums:` entry, or inline `of:` values (a flattened,
// generated companion enum named `<prefixName><FieldName>`). Falls back to
// plain `String` with neither.
func javaEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	if len(field.OfType) > 0 {
		return prefixName + core.ToUpper(field.Name)
	}
	return "String"
}

// javaFieldType computes the fully-resolved Java type for a field.
// nestedClassName is the PascalCase name already computed for this field, in
// case it's an object/array of objects.
func javaFieldType(field *core.EmiField, nestedClassName string, prefixName string) string {
	if field == nil {
		return "Object"
	}

	if field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable {
		return javaEnumBaseType(field, prefixName)
	}

	base := extractPrimitive(string(field.Type))
	if base == "" {
		base = javaDataStructureType(field, nestedClassName)
	}
	if base == "" {
		base = "Object"
	}
	return base
}
