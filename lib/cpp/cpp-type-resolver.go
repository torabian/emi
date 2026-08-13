package cpp

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// extractPrimitiveGeneric maps a leaf emi primitive type onto its generic-dialect
// (portable ISO C++17) scalar type. Returns "" when the field isn't a primitive,
// so callers fall through to structural resolution (array, object, one, ...).
func extractPrimitiveGeneric(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "std::string"
	case "int", "int32", "int32?", "int?":
		return "int32_t"
	case "int64", "int64?":
		return "int64_t"
	case "float32", "float32?":
		return "float"
	case "float64", "float64?":
		return "double"
	case "bool", "bool?":
		return "bool"
	default:
		return ""
	}
}

// extractPrimitiveUnreal maps a leaf emi primitive type onto its Unreal Engine
// scalar type - UE's own aliases (int32/int64/FString/...), not raw std:: types,
// so a generated USTRUCT reads exactly like hand-written Unreal code.
func extractPrimitiveUnreal(fieldType string) string {
	switch fieldType {
	case "string", "string?":
		return "FString"
	case "int", "int32", "int32?", "int?":
		return "int32"
	case "int64", "int64?":
		return "int64"
	case "float32", "float32?":
		return "float"
	case "float64", "float64?":
		return "double"
	case "bool", "bool?":
		return "bool"
	default:
		return ""
	}
}

func extractPrimitive(dialect Dialect, fieldType string) string {
	if dialect == DialectUnreal {
		return extractPrimitiveUnreal(fieldType)
	}
	return extractPrimitiveGeneric(fieldType)
}

// cppNullable wraps a generic-dialect type in std::optional<T>, the uniform
// nullability wrapper this dialect uses for primitives/strings/enums/collections
// (a single `one`/`class` relation, and a self-referencing `object` field, are
// already inherently nullable via std::unique_ptr<T> - see cppGenericRelationField/
// cppGenericObjectField - so they never get double-wrapped).
func cppNullable(cppType string) string {
	return fmt.Sprintf("std::optional<%v>", cppType)
}

// Note: there's no ueOptional/TOptional<T> equivalent of cppNullable above for
// *reflected* (UPROPERTY) unreal-dialect fields - see ueFieldMaybeNullable's doc
// comment in cpp-field-plan-unreal.go for why. TOptional<T> is still used
// verbatim in the plain (non-UPROPERTY) header/query-param classes
// (cpp-headers.go, cpp-query-params.go), where no UHT/engine-version
// constraint applies.

// cppFieldNestedTypeName computes the flattened companion class/struct name for
// an object/array-of-object/inline-enum field, exactly like every sibling
// generator's own equivalent (csFieldNestedClassName, cFieldNestedTypeName, ...).
func cppFieldNestedTypeName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

// cppEnumBaseType resolves an `enum`/`enum?` field: a `target:` reference to a
// module-level `enums:` entry, or inline `of:` values (a flattened, generated
// companion enum named `<prefixName><FieldName>`).
func cppEnumBaseType(field *core.EmiField, prefixName string) string {
	if field.Target != "" {
		return field.Target
	}
	return prefixName + core.ToUpper(field.Name)
}
