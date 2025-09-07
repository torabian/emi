package js

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

func IsNumericDataType(value string) bool {
	switch value {
	case "int64?", "int32?", "int?", "float64?", "float32?", "int64", "int32", "int":
		return true
	default:
		return false
	}
}
func IsNullable(value string) bool {
	return strings.Contains(value, "?")
}

func TsComputedField(field *core.EmiField, isWorkspace bool) string {
	if field.Complex != "" {
		return field.Complex
	}
	switch field.Type {
	case "string", "text", "string?":
		return "string"
	case "one":
		return field.Target
	case "daterange":
		return "any"
	case "enum":
		items := []string{}
		for _, item := range field.OfType {
			items = append(items, "\""+item.Key+"\"")
		}
		return strings.Join(items, " | ")
	case "json":
		return TsCalcJsonField(field)
	case "many2many":
		return field.Target + "[]"
	case "int64?", "int32?", "int?", "float64?", "float32?":
		return "number"
	case "bool?":
		return "boolean"
	case "array":
		return field.PublicName() + "[]"
	case "arrayP":
		return TsPrimitive(field.Primitive) + "[]"
	case "html":
		return "string"
	case "int64", "int32", "int":
		return "number"
	case "float64", "float32", "float":
		return "number"
	case "bool":
		return "boolean"
	case "Timestamp", "datenano":
		return "string"
	case "date":
		return "Date"
	case "double":
		return "number"
	case "object", "embed":
		return field.PublicName()
	case "money?":
		return "{amount: number, currency: string, formatted?: string}"
	default:
		return "any"
	}
}

func TsPrimitive(primitive string) string {
	switch primitive {
	case "string", "text":
		return "string"
	case "string?", "text?":
		return "string"
	case "int64", "int32", "int", "float64", "float32":
		return "number"
	case "int64?", "int32?", "int?", "float64?", "float32?":
		return "number"
	case "bool":
		return "boolean"
	case "bool?":
		return "boolean"
	default:
		return "unknown"
	}
}

func TsCalcJsonField(field *core.EmiField) string {
	t := []string{}

	if len(field.Matches) > 0 {

		for _, match := range field.Matches {
			if match.Dto != nil {
				t = append(t, match.PublicName())
			}
		}

	} else {
		t = append(t, "any")
	}

	return strings.Join(t, "|")
}
