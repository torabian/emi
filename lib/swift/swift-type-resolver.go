package swift

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

func extractPrimitive(field *core.EmiField) string {
	switch field.Type {

	case "string", "string?":
		return "String"
	case "int64", "int64?":
		return "Long"
	case "int32", "int", "int32?", "int?":
		return "Int"
	case "float64", "float64?":
		return "Double"
	case "float32", "float32?":
		return "Float"
	case "bool", "bool?":
		return "Boolean"
	default:
		return ""
	}
}

func swiftDataStructureType(field *core.EmiField) string {

	// Now let's check data structure types.
	switch field.Type {
	case core.FieldTypeOne:
		if field.Module != "" {
			return field.Module + field.Target
		}
		return field.Target
	case core.FieldTypeArray:
		return field.PublicName()
	case core.FieldTypeCollection:
		if field.Module != "" {
			return field.Module + field.Target
		}
		return fmt.Sprintf("[%s]", field.Target)
	case core.FieldTypeSlice:
		return fmt.Sprintf("[%v]", core.ToUpper(field.Primitive))

	case core.FieldTypeObject:
		return field.PublicName()
	default:
		return ""
	}
}

func goComputedField(field *core.EmiField) string {

	// Let's resolve the primitive type first.
	primitiveValue := extractPrimitive(field)
	if primitiveValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("%v?", primitiveValue)
		}

		return primitiveValue
	}

	// Let's try to compute the advanced fields, such as array, collection, references.
	structureFieldValue := swiftDataStructureType(field)
	if structureFieldValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("%v?", structureFieldValue)
		}

		return structureFieldValue
	}

	return "Any"
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
		return fmt.Sprintf("[%v]", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("%v?", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("[%v]?", prefix)
	default:
		return goComputedField(field)
	}
}
