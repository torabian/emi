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

// swiftResolveTarget resolves both the `target: self@` recursive-reference convention
// and the `target: <EntityName>Entity` convention for a one/collection field's target -
// the Swift counterpart of kotlin/kotlin-type-resolver.go's kotlinResolveTarget (see
// its own doc comment for the full rationale: only Go's own entity codegen emits a
// literal "<EntityName>Entity" type, and lib/core/preprocess-entities.go's
// entityDtoRelationTarget only rewrites a *synthesized* entity dto's own fields, not a
// hand-declared dto's - both were previously unhandled here, rendering literal,
// unresolved `self@`/`XEntity` types). rootClassName is the enclosing top-level dto's
// own name (e.g. "CapabilityInfoDto"), not the nested parentChain used for object/array
// child-class naming - a self@ field always means the outermost class, regardless of
// nesting depth.
func swiftResolveTarget(target string, rootClassName string) string {
	const selfField = "self@"
	if strings.Contains(target, selfField) {
		target = strings.ReplaceAll(target, selfField, rootClassName)
	}
	const entitySuffix = "Entity"
	if len(target) > len(entitySuffix) && strings.HasSuffix(target, entitySuffix) {
		target = strings.TrimSuffix(target, entitySuffix) + "Dto"
	}
	return target
}

func swiftDataStructureType(field *core.EmiField, rootClassName string) string {

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
		target := swiftResolveTarget(field.Target, rootClassName)
		if field.Module != "" {
			return field.Module + target
		}
		return target
	case core.FieldTypeArray, core.FieldTypeArrayNullable:
		return field.PublicName()
	case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
		target := swiftResolveTarget(field.Target, rootClassName)
		if field.Module != "" {
			return fmt.Sprintf("[%s%s]", field.Module, target)
		}
		return fmt.Sprintf("[%s]", target)
	case core.FieldTypeSlice, core.FieldTypeSliceNullable:
		return fmt.Sprintf("[%v]", core.ToUpper(field.Primitive))

	case core.FieldTypeObject, core.FieldTypeObjectNullable:
		return field.PublicName()
	default:
		return ""
	}
}

func goComputedField(field *core.EmiField, rootClassName string) string {

	// Let's resolve the primitive type first.
	primitiveValue := extractPrimitive(field)
	if primitiveValue != "" {
		if core.IsNullable(string(field.Type)) || alwaysOptionalOnTheWire(field) {
			return fmt.Sprintf("%v?", primitiveValue)
		}

		return primitiveValue
	}

	// Let's try to compute the advanced fields, such as array, collection, references.
	structureFieldValue := swiftDataStructureType(field, rootClassName)
	if structureFieldValue != "" {
		if core.IsNullable(string(field.Type)) || alwaysOptionalOnTheWire(field) {
			return fmt.Sprintf("%v?", structureFieldValue)
		}

		return structureFieldValue
	}

	return "Any"
}

// alwaysOptionalOnTheWire reports whether field's Swift type should be Optional
// regardless of whether the emi field itself was declared nullable ("?") - two cases,
// neither about the field's "required" validation status:
//
//  1. Array/slice/collection/map: Go's own zero value for these is nil, which
//     encoding/json marshals as JSON `null` (never `[]`/`{}`) whenever a hand-written
//     Implementation.go builds a response struct without explicitly initializing that
//     field. Confirmed for real: CheckClassicPassportActionRes.flags (declared plain
//     `type: string` array, no "?") came back as `"flags":null` whenever empty.
//  2. `tags: {json: "-"}`: Go's own encoding/json omits the field entirely - it's
//     never on the wire at all, e.g. PassportEntity.Password (write-only, only ever
//     used to *set* a password, never to read one back).
//
// Codable's synthesized `init(from:)` throws for a non-Optional property in either
// case (`try decode(...)` requires the key to both exist and be non-null) but resolves
// cleanly to nil for an Optional one (`try decodeIfPresent(...)`) - confirmed for
// real: ClassicSigninActionRes (nests PassportDto) failed to decode entirely with
// "keyNotFound: password" despite the sign-in having genuinely succeeded server-side.
// A `default:`-driven fallback (SwiftSafeDefaultValue) can't substitute for this: Swift
// only uses a property's default to fill in a *missing* key for genuinely Optional
// properties (via decodeIfPresent) - for a non-Optional `var x: T = default`, synthesis
// still calls plain `decode(...)` and still throws, confirmed empirically against this
// toolchain (Xcode 27 / Swift 6.4).
func alwaysOptionalOnTheWire(field *core.EmiField) bool {
	if field == nil {
		return false
	}
	if field.Tags["json"] == "-" {
		return true
	}
	switch core.FieldType(strings.TrimSuffix(string(field.Type), "?")) {
	case core.FieldTypeArray, core.FieldTypeSlice, core.FieldTypeCollection, core.FieldTypeMap:
		return true
	default:
		return false
	}
}

func goFieldTypeOnNestedClasses(field *core.EmiField, parentChain string, rootClassName string) string {
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
		if alwaysOptionalOnTheWire(field) {
			return fmt.Sprintf("%v?", prefix)
		}
		return fmt.Sprintf(" %v", prefix)
	case core.FieldTypeArray:
		// A nested-class-shaped array (e.g. WhoamiActionRes.workspaces, each item an
		// inline `object`) bypasses goComputedField/swiftDataStructureType entirely -
		// alwaysOptionalOnTheWire's array/slice/collection/map case still applies
		// (same Go nil-slice-marshals-as-null concern), so it needs the same check
		// here, not just there.
		if alwaysOptionalOnTheWire(field) {
			return fmt.Sprintf("[%v]?", prefix)
		}
		return fmt.Sprintf("[%v]", prefix)
	case core.FieldTypeObjectNullable:
		return fmt.Sprintf("%v?", prefix)
	case core.FieldTypeArrayNullable:
		return fmt.Sprintf("[%v]?", prefix)
	default:
		return goComputedField(field, rootClassName)
	}
}
