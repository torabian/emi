// C# is the simplest of the newer targets to generate for: System.Text.Json
// deserializes/serializes any plain POCO (public settable properties, a
// public parameterless constructor) automatically via reflection, matching
// property initializers for anything missing in the payload. So unlike
// lib/dart, there's no per-class fromJson/toJson to hand-generate here -
// just the property declaration (with its `[JsonPropertyName]` and a safe
// default initializer) computed once per field.
package csharp

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type csFieldPlan struct {
	Name        string // C# PascalCase property name
	WireName    string // the raw field.Name, used in [JsonPropertyName(...)]
	TypeHint    string
	Attributes  []string // e.g. `[JsonPropertyName("title")]`, `[JsonConverter(typeof(StatusConverter))]`
	Initializer string   // e.g. `= "";`, `= new();`, "" (none needed)
	InlineEnum  *csRenderedEnum
}

func csFieldNestedClassName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

func csJsonPropertyNameAttr(field *core.EmiField) string {
	b, _ := json.Marshal(field.Name)
	return fmt.Sprintf("[JsonPropertyName(%v)]", string(b))
}

// csResolveField computes everything the class generator needs to render one
// property: its type, attributes, and default initializer.
func csResolveField(field *core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) csFieldPlan {
	nestedClassName := csFieldNestedClassName(field, prefixName)
	typeHint := csFieldType(field, nestedClassName, prefixName)
	propertyName := core.ToUpper(field.Name)

	plan := csFieldPlan{
		Name:       propertyName,
		WireName:   field.Name,
		TypeHint:   typeHint,
		Attributes: []string{csJsonPropertyNameAttr(field)},
	}

	// An explicit `default:` in the schema always wins, whatever the type.
	if field.Default != nil {
		if lit, ok := csLiteral(field.Default, typeHint); ok {
			plan.Initializer = "= " + lit + ";"
			return plan
		}
	}

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			plan.TypeHint = csNullable(symbol)
		} else {
			plan.TypeHint = "object?"
		}
		return plan // reference type, implicitly defaults to null
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := csEnumBaseType(field, prefixName)
		if field.Target == "" && base != "string" {
			plan.InlineEnum = csRenderEnumFromInline(base, field.OfType)
		}
		if base == "string" {
			if !core.IsNullable(string(field.Type)) {
				plan.Initializer = `= "";`
			}
			return plan
		}
		// The `[JsonConverter(typeof(XConverter))]` attribute lives on the
		// enum *type* declaration itself (see csRenderEnumDecl) - System.Text.Json
		// honors it for every property of that type automatically, so
		// nothing extra is needed here. Non-nullable enum properties already
		// default to their first declared member (ordinal 0) with no
		// initializer needed - C#'s own `default` semantics for enums.
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		// A single relation is never eagerly default-constructed - always a
		// plain nullable reference, implicitly defaulting to null.
		if field.Target == "" {
			plan.TypeHint = "object?"
		}
		return plan

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		if core.IsNullable(string(field.Type)) {
			return plan
		}
		plan.Initializer = "= new();"
		return plan

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable ||
		field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable ||
		field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		if core.IsNullable(string(field.Type)) {
			return plan
		}
		plan.Initializer = "= new();"
		return plan

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedClassName == selfClassName
		if core.IsNullable(string(field.Type)) || selfReferencing {
			plan.TypeHint = csNullable(nestedClassName)
			return plan
		}
		plan.Initializer = "= new();"
		return plan

	default:
		// Primitive (string/int/float/bool/any).
		if core.IsNullable(string(field.Type)) {
			return plan
		}
		zero, ok := csZeroValue(typeHint)
		if ok {
			plan.Initializer = "= " + zero + ";"
		}
		return plan
	}
}

func csZeroValue(csType string) (string, bool) {
	switch csType {
	case "string":
		return `""`, true
	case "int", "long":
		return "0", true
	case "float":
		return "0.0f", true
	case "double":
		return "0.0", true
	case "bool":
		return "false", true
	default:
		return "", false
	}
}

// csLiteral renders an arbitrary go value (coming from EmiField.Default) as a
// C# literal expression, when it's a simple, unambiguous scalar - returns
// ok=false for anything else (nested objects/arrays), so the caller falls
// back to the type-based default instead of emitting something invalid.
func csLiteral(v any, csType string) (string, bool) {
	switch value := v.(type) {
	case string:
		b, _ := json.Marshal(value)
		return string(b), true
	case bool:
		if value {
			return "true", true
		}
		return "false", true
	case int, int32, int64, float32, float64:
		suffix := ""
		if csType == "float" || strings.HasPrefix(csType, "float?") {
			suffix = "f"
		}
		return fmt.Sprintf("%v%v", value, suffix), true
	default:
		return "", false
	}
}

// resolveComplexSymbol mirrors the sibling generators: a `complex:` field is
// only ever trusted as a real generated type when it's `+`-prefixed AND
// resolvable via RecognizedComplexes - anything else falls back to `object`.
func resolveComplexSymbol(field *core.EmiField, complexes []RecognizedComplex) (symbol string, resolved bool) {
	symbol = strings.ReplaceAll(field.Complex, "+", "")
	if !strings.Contains(field.Complex, "+") {
		return "", false
	}
	for _, item := range complexes {
		if item.Symbol == symbol {
			return symbol, item.ImportLocation != ""
		}
	}
	return "", false
}
