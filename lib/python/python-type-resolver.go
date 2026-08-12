package python

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitive maps a leaf emi primitive type onto its python type-hint
// equivalent. Returns "" when the field isn't a primitive, so callers can
// fall through to structural resolution (array, object, one, collection...).
// enum/enum? is deliberately not handled here - see pythonEnumBaseTypeHint,
// which needs the full field (Target / OfType), not just the type string.
func extractPrimitive(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "str"
	case "int", "int32", "int64", "int?", "int32?", "int64?":
		return "int"
	case "float32", "float64", "float32?", "float64?":
		return "float"
	case "bool", "bool?":
		return "bool"
	case "any", "complex", "any?", "complex?":
		return "Any"
	default:
		return ""
	}
}

// pythonMapType resolves the `Dict[key, value]` annotation for `map`/`map?` fields.
func pythonMapType(field *core.EmiField) string {
	key := extractPrimitive(field.MapKeyOf)
	if key == "" {
		key = "str"
	}

	value := extractPrimitive(field.MapPairOf)
	if value == "" {
		value = "Any"
	}

	return fmt.Sprintf("Dict[%v, %v]", key, value)
}

// pythonDataStructureType resolves array/_list/slice/object/one/collection/
// class fields to their python type-hint, given the already-computed name of
// the nested/target class (parentChain based, PascalCase - see
// pythonFieldTypeOnNestedClasses for how it's derived).
//
// `one`/`class`/`collection` always reference *another* generated class by
// name (`target:`) - nestedClassName is only ever a fallback for a malformed
// definition missing that target, and one that's never actually emitted as a
// class anywhere (only object/array/_list fields get a flattened subclass),
// so it deliberately does NOT fall back to nestedClassName - that would
// reference a name that doesn't exist, which blows up
// `typing.get_type_hints()` (used by the runtime's `from_dict`) with a
// NameError. `Any`/`List[Any]` is the safe, always-valid fallback instead.
func pythonDataStructureType(field *core.EmiField, nestedClassName string) string {
	switch field.Type {
	case core.FieldTypeOne, core.FieldTypeClass, core.FieldTypeOneNullable, core.FieldTypeClassNullable:
		if field.Target != "" {
			return field.Target
		}
		return "Any"
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Target != "" {
			return fmt.Sprintf("List[%v]", field.Target)
		}
		return "List[Any]"
	case core.FieldTypeArray, core.FieldTypeArrayNullable,
		// `_list`/`_list?` is the golang-entity-only twin of `array`/`array?`
		// (see EmiFieldType.go) - no other generator special-cases it either,
		// so it's rendered exactly the same way here.
		core.FieldTypeList, core.FieldTypeListNullable:
		return fmt.Sprintf("List[%v]", nestedClassName)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		primitive := extractPrimitive(field.Primitive)
		if primitive == "" {
			primitive = "Any"
		}
		return fmt.Sprintf("List[%v]", primitive)
	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return nestedClassName
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return pythonMapType(field)
	default:
		return ""
	}
}

// pythonEnumLiteralType renders a field's inline `of:` enum values as a
// `typing.Literal["a", "b", ...]` - the closest python has to a TS string
// union / kotlin sealed value set. Returns "" when there are no values to
// render (the caller falls back to plain `str`).
func pythonEnumLiteralType(field *core.EmiField) string {
	if len(field.OfType) == 0 {
		return ""
	}

	keys := make([]string, 0, len(field.OfType))
	for _, item := range field.OfType {
		if item == nil {
			continue
		}
		b, _ := json.Marshal(item.Key)
		keys = append(keys, string(b))
	}
	if len(keys) == 0 {
		return ""
	}

	return fmt.Sprintf("Literal[%v]", strings.Join(keys, ", "))
}

// pythonEnumBaseTypeHint resolves an `enum`/`enum?` field, which can mean two
// different things (mirrors lib/js/js-data-types.go#TsComputedField's "enum"
// case): a `target:` reference to a module-level `enums:` entry (an actual
// generated `class X(str, Enum)`), or inline `of:` values with no target (a
// `Literal[...]`). Neither can ever be safely default-constructed on its own
// (we don't know which member is "the" default), so both are always
// optional - see the `isSingleExternalReference` handling this feeds into.
func pythonEnumBaseTypeHint(field *core.EmiField) (typeHint string, isExternalRef bool) {
	if field.Target != "" {
		return field.Target, true
	}
	if lit := pythonEnumLiteralType(field); lit != "" {
		return lit, false
	}
	return "str", false
}

// pythonFieldTypeOnNestedClasses computes the fully-resolved python type hint
// for a field, including the Optional[...] wrapper for nullable emi types.
// nestedClassName is the PascalCase name already computed for this field, in
// case it turns out to be an object/array of objects (see flattenNestedName).
func pythonFieldTypeOnNestedClasses(field *core.EmiField, nestedClassName string) string {
	if field == nil {
		return "Any"
	}

	if field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable {
		base, isExternalRef := pythonEnumBaseTypeHint(field)
		if core.IsNullable(string(field.Type)) || isExternalRef {
			return fmt.Sprintf("Optional[%v]", base)
		}
		return base
	}

	base := extractPrimitive(string(field.Type))
	if base == "" {
		base = pythonDataStructureType(field, nestedClassName)
	}
	if base == "" {
		base = "Any"
	}

	// A single `one`/`class` relation (unlike `collection`, a *list* of
	// relations, or `array`/`object`, always locally/eagerly constructible)
	// is never eagerly default-constructed - see PythonSafeDefaultValue -
	// even when the schema marks it required, so its type-hint always has to
	// say Optional too, or every dto referencing one fails a type-checker on
	// its own default.
	isSingleRelation := field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass
	if core.IsNullable(string(field.Type)) || isSingleRelation {
		return fmt.Sprintf("Optional[%v]", base)
	}

	return base
}
