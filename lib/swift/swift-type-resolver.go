package swift

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

func extractPrimitive(field *core.EmiField) string {
	switch field.Type {

	case "string", "string?":
		return "String"
	// enum/enum? carries no native Swift enum type of its own (see EmiField's "of" list) -
	// matches Go's own dto codegen, which also represents it as a plain string (see
	// examples/emi-entity/sdk/Entity1Dto.go's Status field).
	case "enum", "enum?":
		return "String"
	case "int64", "int64?":
		return "Int64"
	case "int32", "int", "int32?", "int?":
		return "Int"
	case "float64", "float64?":
		return "Double"
	case "float32", "float32?":
		return "Float"
	case "bool", "bool?":
		return "Bool"
	// A bare "Any" cannot satisfy Codable (Swift, unlike Go's interface{}, has no
	// built-in way to encode/decode an unconstrained value) - EmiAnyCodable (see
	// swift-any-codable.go) is a small wrapper that can, emitted once per module by
	// SwiftFullModule regardless of whether any field actually needs it. any has no
	// portable "?" form (see preprocess-entities.go's cloneEntityOptionalField), so
	// there's no nullable variant to match here.
	case "any":
		return "EmiAnyCodable"
	default:
		return ""
	}
}

// swiftPrimitiveTypeName maps an EmiField primitive name (map key/value type, "string",
// "int", ...) to its Swift equivalent. Falls back to "String" for anything unrecognized
// (including "object"/"slice" map-value shapes, which - unlike Go's own map codegen in
// go-struct-generator-common.go - this doesn't attempt to resolve into a nested class
// reference) - a safe, always-Codable-compatible default rather than emitting nothing.
func swiftPrimitiveTypeName(t string) string {
	switch t {
	case "string":
		return "String"
	case "int", "int32":
		return "Int"
	case "int64":
		return "Int64"
	case "float32":
		return "Float"
	case "float64":
		return "Double"
	case "bool":
		return "Bool"
	default:
		return "String"
	}
}

func swiftDataStructureType(field *core.EmiField) string {

	// Now let's check data structure types. Every case has to match both a field's
	// plain and "?" (nullable) type - goComputedField below only appends the "?"
	// wrapper when this returns a non-empty base type, so a case missing its nullable
	// counterpart here (e.g. "one?"/"collection?" - always the case for an entity's own
	// relation fields, see preprocess-entities.go's shallowCloneField, which forces
	// every one/collection field on a portable dto nullable) silently falls through to
	// the "Any" catch-all in goComputedField instead of an optional Target.
	switch field.Type {
	case core.FieldTypeMap, core.FieldTypeMapNullable:
		return fmt.Sprintf("[%s: %s]", swiftPrimitiveTypeName(field.GetMapKeyType()), swiftPrimitiveTypeName(field.GetMapValueType()))
	case core.FieldTypeOne, core.FieldTypeOneNullable:
		if field.Module != "" {
			return field.Module + field.Target
		}
		return field.Target
	case core.FieldTypeArray, core.FieldTypeArrayNullable:
		return field.PublicName()
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		if field.Module != "" {
			return fmt.Sprintf("[%s%s]", field.Module, field.Target)
		}
		return fmt.Sprintf("[%s]", field.Target)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		return fmt.Sprintf("[%v]", core.ToUpper(field.Primitive))

	case core.FieldTypeObject, core.FieldTypeObjectNullable:
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
	// A complex field's Swift type is the complex's own name (e.g. "Money") - matches
	// lib/golang/go-struct-generator-common.go's own FieldTypeComplex case. Unlike Go,
	// there's no namespace-qualification here yet: Swift's RecognizedComplex (see
	// swift-class-generator.go) carries no Namespace field the way Go's does.
	case core.FieldTypeComplex:
		return strings.ReplaceAll(field.Complex, "+", "")
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
