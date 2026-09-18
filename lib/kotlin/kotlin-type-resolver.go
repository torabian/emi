package kotlin

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// kotlinResolveRelationTarget rewrites a one/collection field's target from another
// entity's own persisted-struct name convention ("<EntityName>Entity") to that same
// entity's portable dto class instead ("<EntityName>Dto") - mirrors
// lib/js/js-common-object-class.go's entityTargetToCodeChunk (the "Entity" suffix
// branch) and lib/core/preprocess-entities.go's entityDtoRelationTarget: only Go's own
// entity codegen ever emits a literal "<EntityName>Entity" type: every other backend
// (js, kotlin, swift, openapi) only ever sees that entity's derived dto - see those two
// functions' own doc comments for the full rationale. Unlike JS (which resolves this via
// an aliased import, `import { XDto as XEntity } from "./XDto"`, since a TS caller
// addresses the type itself) Kotlin just rewrites the type reference directly to XDto -
// a generated Kotlin dto/action field is addressed by its field name (e.g. `passport`),
// never by its type's class name, so there's nothing that needs a literal `XEntity`
// symbol to exist. Before this, a hand-declared dto's `target: XEntity` field (e.g.
// UserSessionDto.passport/user/userWorkspaces in modules/abac/Abac.emi.yml) rendered an
// unresolved `XEntity` type, since Kotlin (unlike Go) never emits one.
func kotlinResolveRelationTarget(target string) string {
	const suffix = "Entity"
	if len(target) <= len(suffix) || !strings.HasSuffix(target, suffix) {
		return target
	}
	return strings.TrimSuffix(target, suffix) + "Dto"
}

// kotlinSelfField is the documented `target: self@` convention for a one/collection
// field that recursively references its own enclosing dto (e.g.
// CapabilityInfoDto.children in modules/abac/Abac.emi.yml, a capability tree) - mirrors
// lib/js/js-common-object-class.go's SELF_FIELD/getSelfReferencingField. Kotlin had no
// handling for it at all, so it rendered as the literal, unresolved type token `self@`.
const kotlinSelfField = "self@"

// kotlinResolveTarget resolves both the self@ convention and the Entity-suffix
// convention (kotlinResolveRelationTarget) for a one/collection field's target.
// parentChain is the enclosing top-level dto/action-response class's own name (e.g.
// "CapabilityInfoDto") - see renderClasses/KotlinCommonStructGenerator's initial call,
// which seeds it from goctx.RootClassName - matching JS's
// `strings.Split(parentChain, ".")[0]` (always the outermost class, never a nested
// object's own synthesized name, even when self@ appears inside one).
func kotlinResolveTarget(target string, parentChain string) string {
	if strings.Contains(target, kotlinSelfField) {
		target = strings.ReplaceAll(target, kotlinSelfField, parentChain)
	}
	return kotlinResolveRelationTarget(target)
}

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
func kotlinDataStructureType(field *core.EmiField, parentChain string) string {
	baseType := core.FieldType(strings.TrimSuffix(string(field.Type), "?"))

	switch baseType {
	case core.FieldTypeOne, core.FieldTypeCollection:
		// Bare class name in both cases: same-package (module unset) references need
		// no import at all (see CombineJavaImport), and cross-module (module set)
		// references get a real `import <module>.<Target>` line generated alongside -
		// see kotlinCollectTargetDeps - so the type itself never needs qualifying.
		target := kotlinResolveTarget(field.Target, parentChain)
		if baseType == core.FieldTypeCollection {
			return fmt.Sprintf("List<%s>", target)
		}
		return target
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

func goComputedField(field *core.EmiField, parentChain string) string {

	// Let's resolve the primitive type first.
	primitiveValue := extractPrimitive(field)
	if primitiveValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("MaybeField<%v>", primitiveValue)
		}

		return primitiveValue
	}

	// Let's try to compute the advanced fields, such as array, collection, references.
	structureFieldValue := kotlinDataStructureType(field, parentChain)
	if structureFieldValue != "" {
		if core.IsNullable(string(field.Type)) {
			return fmt.Sprintf("MaybeField<%v>", structureFieldValue)
		}

		return structureFieldValue
	}

	return "Any"
}

// rootClassName is the enclosing top-level dto/action-response class's own name (e.g.
// "CapabilityInfoDto") - distinct from parentChain, which grows with nesting depth
// (e.g. "CapabilityInfoDtoChildren" for a nested object) and isn't what a `target:
// self@` field should resolve to regardless of how deep it's nested. See
// kotlinResolveTarget's own doc comment.
func goFieldTypeOnNestedClasses(field *core.EmiField, parentChain string, rootClassName string) string {
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
		return goComputedField(field, rootClassName)
	}
}
