package kotlin

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

func goComputedField(field *core.EmiField) string {
	switch field.Type {

	case "string", "text", "html", "enum":
		return "String"
	case "string?", "text?", "html?", "enum?":
		return "String?"
	case "one":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "array":
		return field.PublicName()
	case "collection":
		if field.Module != "" {
			return field.Module + "." + field.Target
		}
		return field.Target
	case "slice":
		return fmt.Sprintf("List<%v>", core.ToUpper(field.Primitive))
	case "int64":
		return "Long"
	case "int32", "int":
		return "Int"
	case "float64":
		return "Double"
	case "float32":
		return "Float"
	case "bool":
		return "Boolean"
	case "int64?":
		return "Long?"
	case "int32?", "int?":
		return "Int?"
	case "float64?":
		return "Double?"
	case "float32?":
		return "Float?"
	case "bool?":
		return "Boolean?"
	case "object", "embed":
		return field.PublicName()
	case "json":
		return "JSON"
	default:
		return "Any"
	}
}

func goFieldTypeOnNestedClasses(field *core.EmiField, parentChain string) string {
	if field == nil {
		return ""
	}
	prefix := core.ToUpper(parentChain) + core.ToUpper(field.Name)
	switch field.Type {
	case core.FieldTypeObject:
		return fmt.Sprintf(" %v", prefix)
	case core.FieldTypeArray:
		return fmt.Sprintf("List<%v>", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("List<%v>", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("List<%v>", prefix)
	default:
		return goComputedField(field)
	}
}
