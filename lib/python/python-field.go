package python

import (
	"encoding/json"
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// pythonLiteral renders an arbitrary go value (coming from EmiField.Default,
// which is free-form yaml/json) as a python literal expression.
func pythonLiteral(v any) string {
	switch value := v.(type) {
	case string:
		b, _ := json.Marshal(value)
		return string(b)
	case bool:
		if value {
			return "True"
		}
		return "False"
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", value)
	case nil:
		return "None"
	default:
		b, _ := json.Marshal(value)
		return string(b)
	}
}

// PythonSafeDefaultValue computes the right-hand side of a dataclass field
// assignment (everything after the `=`), mirroring the "safe default"
// philosophy used by the sibling generators (JsGetSafeFieldValue,
// KotlinSafeDefaultValue): every field always gets *some* default, so field
// ordering in the generated dataclass never has to be worried about, and a
// bare `SomeDto()` is always constructible.
//
// nestedClassName is the already-computed PascalCase name to use when the
// field itself is a locally generated nested class (object/array of object).
// selfClassName, when non-empty and equal to nestedClassName, signals a
// self-referencing field - defaulting eagerly to `SelfClass()` would recurse
// forever building defaults, so it falls back to None instead.
func PythonSafeDefaultValue(field *core.EmiField, nestedClassName string, selfClassName string) string {
	if field == nil {
		return "None"
	}

	if field.Default != nil {
		if core.IsNullable(string(field.Type)) {
			return pythonLiteral(field.Default)
		}
		return pythonLiteral(field.Default)
	}

	if core.IsNullable(string(field.Type)) {
		return "None"
	}

	switch field.Type {
	case core.FieldTypeArray, core.FieldTypeSlice, core.FieldTypeCollection, core.FieldTypeList:
		return "field(default_factory=list)"
	case core.FieldTypeMap:
		return "field(default_factory=dict)"
	case core.FieldTypeObject:
		if nestedClassName != "" && nestedClassName != selfClassName {
			return fmt.Sprintf("field(default_factory=lambda: %v())", nestedClassName)
		}
		return "None"
	case core.FieldTypeOne, core.FieldTypeClass:
		// A single relation to another (possibly foreign, possibly
		// self-referencing) class - it's never safe to eagerly build one,
		// so it always starts unset even though the field isn't nullable.
		return "None"
	case core.FieldTypeEnum:
		// A `target:` reference to a module enum is a single external
		// reference (same reasoning as one/class above) - never
		// default-constructed. Inline `of:` values default to their first
		// member (a valid `Literal[...]` value); with neither, it's a plain
		// (untyped) string.
		if field.Target != "" {
			return "None"
		}
		if len(field.OfType) > 0 && field.OfType[0] != nil {
			return pythonLiteral(field.OfType[0].Key)
		}
		return `""`
	case core.FieldTypeString:
		return `""`
	case core.FieldTypeBool:
		return "False"
	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		return "0"
	case core.FieldTypeFloat32, core.FieldTypeFloat64:
		return "0.0"
	default:
		return "None"
	}
}
