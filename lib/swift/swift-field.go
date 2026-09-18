package swift

import (
	"encoding/json"
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// SwiftSafeDefaultValue computes a field's default value as real Swift syntax -
// previously this was a straight copy-paste of Kotlin's own KotlinSafeDefaultValue
// (see git history), producing e.g. `MaybeField(Maybe.Value("90"))` or `emptyList()`,
// neither valid Swift (MaybeField doesn't exist in this package - Swift represents
// nullability with plain `T?`, not a wrapper type - and Kotlin's `emptyList()` isn't
// Swift syntax at all). Currently unused by fieldVariable.Compile() (see
// swift-common-fields.go, which leaves every generated field un-defaulted rather than
// wiring this in), so the previous bug never actually corrupted generated output - kept
// correct here regardless, both so it's ready to wire in and so it doesn't mislead
// anyone maintaining this file into thinking Kotlin syntax is intentional here.
func SwiftSafeDefaultValue(field *core.EmiField) string {
	if field == nil {
		return "nil"
	}

	if field.Default != nil {
		switch v := field.Default.(type) {
		case string:
			return fmt.Sprintf("%q", v)
		case int, int64, float64, bool:
			return fmt.Sprintf("%v", v)
		default:
			b, _ := json.Marshal(v)
			return string(b)
		}
	}

	switch field.Type {
	case core.FieldTypeAny:
		return "nil"
	case core.FieldTypeArray, core.FieldTypeSlice, core.FieldTypeCollection:
		return "[]"
	case core.FieldTypeArrayNullable, core.FieldTypeSliceNullable, core.FieldTypeCollectionNullable,
		core.FieldTypeObjectNullable, core.FieldTypeOneNullable, core.FieldTypeMapNullable,
		core.FieldTypeStringNullable, core.FieldTypeBoolNullable, core.FieldTypeIntNullable,
		core.FieldTypeInt32Nullable, core.FieldTypeInt64Nullable, core.FieldTypeFloat32Nullable,
		core.FieldTypeFloat64Nullable, core.FieldTypeEnumNullable:
		return "nil"
	case core.FieldTypeMap:
		return "[:]"
	case core.FieldTypeBool:
		return "false"
	case core.FieldTypeString, core.FieldTypeEnum:
		// field.Default is already known nil at this point (the `field.Default !=
		// nil` block above returns first when it isn't) - this used to re-check via
		// `field.Default != ""`, comparing an untyped-nil interface{} against a typed
		// string constant, which Go never considers equal, so it always took the
		// "has a default" branch and called getDefault(field) on a nil Default -
		// `json.Marshal(nil)` there produced the 4-char string `null`, rendered
		// verbatim as the Swift string literal `"null"` (a real, non-empty default
		// value, not an absence of one). Confirmed live: PassportDto.password
		// (String, tags: {json: "-"}, no explicit `default:`) got `= "null"`.
		return `""`
	case core.FieldTypeFloat32:
		return "0.0"
	case core.FieldTypeFloat64:
		return "0.0"
	case core.FieldTypeInt, core.FieldTypeInt32, core.FieldTypeInt64:
		return "0"
	default:
		return ""
	}
}
