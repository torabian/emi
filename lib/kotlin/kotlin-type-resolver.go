package kotlin

import (
	"fmt"
	"strings"

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
	// enum is a plain string on the wire, same as Go (go-common-fields.go) and JS
	// (js-common-fields.go) already treat it - there's no real Kotlin `enum class`
	// generated, so "enum"/"enum?" resolve exactly like "string"/"string?".
	case "enum", "enum?":
		return "String"
	default:
		return ""
	}
}

// kotlinPrimitiveMapType maps a map's mapKeyOf/mapPairOf primitive name (see
// EmiField.MapKeyOf/MapPairOf) to its Kotlin type. Unknown/unset -> String, matching
// extractPrimitive's own default posture rather than failing closed.
func kotlinPrimitiveMapType(primitive string) string {
	switch primitive {
	case "int":
		return "Int"
	case "any":
		return "Any"
	case "string", "":
		return "String"
	default:
		return "String"
	}
}

// kotlinDataStructureType resolves every "structured" (non-primitive) field type to its
// Kotlin type. field.Type is normalized by trimming a trailing "?" first, so a nullable
// field (e.g. "one?", "map?") resolves identically to its non-nullable twin - nullability
// itself is applied separately by goComputedField's MaybeField<...> wrap, based on
// core.IsNullable. Before this fix the switch only ever matched the non-nullable
// constant, so "one?"/"collection?"/"map?"/... all silently fell through to "Any".
func kotlinDataStructureType(field *core.EmiField) string {
	baseType := core.FieldType(strings.TrimSuffix(string(field.Type), "?"))

	switch baseType {
	case core.FieldTypeOne, core.FieldTypeCollection:
		// Bare class name in both cases: same-package (module unset) references need
		// no import at all (see CombineJavaImport), and cross-module (module set)
		// references get a real `import <module>.<Target>` line generated alongside -
		// see kotlinCollectTargetDeps - so the type itself never needs qualifying.
		if baseType == core.FieldTypeCollection {
			return fmt.Sprintf("List<%s>", field.Target)
		}
		return field.Target
	case core.FieldTypeArray:
		return field.PublicName()
	case core.FieldTypeSlice:
		return fmt.Sprintf("List<%v>", core.ToUpper(field.Primitive))
	case core.FieldTypeObject:
		return field.PublicName()
	case core.FieldTypeMap:
		keyType := kotlinPrimitiveMapType(field.MapKeyOf)
		valueType := kotlinPrimitiveMapType(field.MapPairOf)
		if field.Target != "" {
			valueType = field.Target
		}
		return fmt.Sprintf("Map<%s, %s>", keyType, valueType)
	case core.FieldTypeComplex:
		return strings.ReplaceAll(field.Complex, "+", "")
	default:
		return ""
	}
}

func goComputedField(field *core.EmiField) string {

	// Let's resolve the primitive type first.
	primitiveValue := extractPrimitive(field)
	if primitiveValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("MaybeField<%v>", primitiveValue)
		}

		return primitiveValue
	}

	// Let's try to compute the advanced fields, such as array, collection, references.
	structureFieldValue := kotlinDataStructureType(field)
	if structureFieldValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("MaybeField<%v>", structureFieldValue)
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
		return fmt.Sprintf("List<%v>", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("MaybeField<%v>", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("MaybeField<List<%v>>", prefix)
	default:
		return goComputedField(field)
	}
}
