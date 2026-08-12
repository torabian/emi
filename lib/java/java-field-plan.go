// Java, like C#, needs no per-class (de)serialization code: Jackson
// (com.fasterxml.jackson.databind.ObjectMapper) reads/writes public fields
// directly via reflection, matching whatever a field's own initializer left
// it at for anything missing in the payload. So this file only has to
// compute each field's declaration (type + `@JsonProperty` + safe default
// initializer) - there's no fromJson/toJson to hand-generate, unlike
// lib/dart.
package java

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type javaFieldPlan struct {
	Name       string // field.Name as-is - already valid lowerCamelCase Java convention
	TypeHint   string
	Attribute  string // `@JsonProperty("wireName")`
	Decl       string // e.g. `public String title = "";`
	InlineEnum *javaRenderedEnum
}

func javaFieldNestedClassName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

func javaJsonPropertyAttr(field *core.EmiField) string {
	b, _ := json.Marshal(field.Name)
	return fmt.Sprintf("@JsonProperty(%v)", string(b))
}

// javaResolveField computes everything the class generator needs to render
// one field: its type, `@JsonProperty` attribute, and default initializer.
func javaResolveField(field *core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) javaFieldPlan {
	nestedClassName := javaFieldNestedClassName(field, prefixName)
	typeHint := javaFieldType(field, nestedClassName, prefixName)
	nullable := core.IsNullable(string(field.Type))

	plan := javaFieldPlan{Name: field.Name, TypeHint: typeHint, Attribute: javaJsonPropertyAttr(field)}

	// An explicit `default:` in the schema always wins, whatever the type.
	if field.Default != nil {
		if lit, ok := javaLiteral(field.Default); ok {
			plan.Decl = fmt.Sprintf("public %v %v = %v;", typeHint, field.Name, lit)
			return plan
		}
	}

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			plan.TypeHint = symbol
			plan.Decl = fmt.Sprintf("public %v %v;", symbol, field.Name)
		} else {
			plan.TypeHint = "Object"
			plan.Decl = fmt.Sprintf("public Object %v;", field.Name)
		}
		return plan
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := javaEnumBaseType(field, prefixName)
		if field.Target == "" && base != "String" {
			plan.InlineEnum = javaRenderEnumFromInline(base, field.OfType)
		}
		if base == "String" {
			if nullable {
				plan.Decl = fmt.Sprintf("public String %v;", field.Name)
			} else {
				plan.Decl = fmt.Sprintf(`public String %v = "";`, field.Name)
			}
			return plan
		}
		if nullable {
			plan.Decl = fmt.Sprintf("public %v %v;", base, field.Name)
		} else {
			plan.Decl = fmt.Sprintf("public %v %v = %v.values()[0];", base, field.Name, base)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		// A single relation is never eagerly default-constructed - always
		// plain `null` (Java's implicit default for reference types).
		t := typeHint
		if field.Target == "" {
			t = "Object"
			plan.TypeHint = "Object"
		}
		plan.Decl = fmt.Sprintf("public %v %v;", t, field.Name)
		return plan

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable:
		if nullable {
			plan.Decl = fmt.Sprintf("public %v %v;", typeHint, field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v %v = new java.util.ArrayList<>();", typeHint, field.Name)
		return plan

	case field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable ||
		field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		if nullable {
			plan.Decl = fmt.Sprintf("public %v %v;", typeHint, field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v %v = new java.util.ArrayList<>();", typeHint, field.Name)
		return plan

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		if nullable {
			plan.Decl = fmt.Sprintf("public %v %v;", typeHint, field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v %v = new java.util.HashMap<>();", typeHint, field.Name)
		return plan

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedClassName == selfClassName
		if nullable || selfReferencing {
			plan.Decl = fmt.Sprintf("public %v %v;", nestedClassName, field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v %v = new %v();", nestedClassName, field.Name, nestedClassName)
		return plan

	default:
		// Primitive (String/Integer/Long/Float/Double/Boolean/Object).
		if nullable || typeHint == "Object" {
			plan.Decl = fmt.Sprintf("public %v %v;", typeHint, field.Name)
			return plan
		}
		zero, ok := javaZeroValue(typeHint)
		if !ok {
			plan.Decl = fmt.Sprintf("public %v %v;", typeHint, field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v %v = %v;", typeHint, field.Name, zero)
		return plan
	}
}

func javaZeroValue(javaType string) (string, bool) {
	switch javaType {
	case "String":
		return `""`, true
	case "Integer":
		return "0", true
	case "Long":
		return "0L", true
	case "Float":
		return "0.0f", true
	case "Double":
		return "0.0", true
	case "Boolean":
		return "false", true
	default:
		return "", false
	}
}

// javaLiteral renders an arbitrary go value (coming from EmiField.Default) as
// a Java literal expression, when it's a simple, unambiguous scalar -
// returns ok=false for anything else, so the caller falls back to the
// type-based default instead of emitting something invalid.
func javaLiteral(v any) (string, bool) {
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
		return fmt.Sprintf("%v", value), true
	default:
		return "", false
	}
}

// resolveComplexSymbol mirrors the sibling generators: a `complex:` field is
// only ever trusted as a real generated type when it's `+`-prefixed AND
// resolvable via RecognizedComplexes - anything else falls back to `Object`.
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
