// PHP 8+ has full runtime reflection (unlike lib/dart/lib/java), so like
// lib/python this only needs to compute a plain typed property declaration
// (+ a PHPDoc `@var` for anything `array`-typed, since PHP has no generics)
// - the generic Hydrator runtime handles (de)serialization for every class
// uniformly. The one PHP-specific wrinkle: property default values must be
// constant expressions, so any non-constant default (a nested object) is
// assigned in the constructor instead of the property declaration itself.
package php

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type phpFieldPlan struct {
	Name           string // field.Name as-is - already valid lowerCamelCase PHP convention
	TypeHint       string
	Doc            string // PHPDoc line, e.g. `/** @var WidgetDtoTags[] */`, or ""
	Decl           string // e.g. `public string $title = '';`
	CtorInit       string // e.g. `$this->meta = new WidgetDtoMeta();`, or "" if none needed
	ArrayItemClass string // for Hydrator::arrayItemTypes() - "" when not applicable
	InlineEnum     *phpRenderedEnum
}

func phpFieldNestedClassName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

// phpResolveField computes everything the class generator needs to render
// one field: its type, PHPDoc, declaration (with a constant default when
// possible), and constructor-assigned default otherwise.
func phpResolveField(field *core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) phpFieldPlan {
	nestedClassName := phpFieldNestedClassName(field, prefixName)
	typeHint := phpFieldType(field, nestedClassName, prefixName)
	nullable := core.IsNullable(string(field.Type))

	plan := phpFieldPlan{Name: field.Name, TypeHint: typeHint}

	// An explicit `default:` in the schema always wins, whatever the type.
	if field.Default != nil {
		if lit, ok := phpLiteral(field.Default); ok {
			plan.Decl = fmt.Sprintf("public %v $%v = %v;", typeHint, field.Name, lit)
			return plan
		}
	}

	if field.Complex != "" {
		symbol, resolved := resolveComplexSymbol(field, complexes)
		if resolved {
			plan.TypeHint = "?" + symbol
			plan.Decl = fmt.Sprintf("public ?%v $%v = null;", symbol, field.Name)
		} else {
			plan.TypeHint = "mixed"
			plan.Decl = fmt.Sprintf("public mixed $%v = null;", field.Name)
		}
		return plan
	}

	switch {
	case field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable:
		base := phpEnumBaseType(field, prefixName)
		if base == "string" {
			if nullable {
				plan.Decl = fmt.Sprintf("public ?string $%v = null;", field.Name)
			} else {
				plan.Decl = fmt.Sprintf("public string $%v = '';", field.Name)
			}
			return plan
		}
		if field.Target == "" {
			plan.InlineEnum = phpRenderEnumFromInline(base, field.OfType)
		}
		if nullable {
			plan.Decl = fmt.Sprintf("public ?%v $%v = null;", base, field.Name)
		} else {
			// `X::cases()[0]` isn't a constant expression - assigned in the
			// constructor instead. A backed enum is always non-empty and
			// enumerable, so (unlike an arbitrary object) this is always
			// safe to eagerly evaluate.
			plan.Decl = fmt.Sprintf("public %v $%v;", base, field.Name)
			plan.CtorInit = fmt.Sprintf("$this->%v = %v::cases()[0];", field.Name, base)
		}
		return plan

	case field.Type == core.FieldTypeOne || field.Type == core.FieldTypeClass ||
		field.Type == core.FieldTypeOneNullable || field.Type == core.FieldTypeClassNullable:
		// A single relation is never eagerly default-constructed - always
		// plain `null`.
		t := typeHint
		if field.Target == "" {
			t = "mixed"
			plan.TypeHint = "mixed"
			plan.Decl = fmt.Sprintf("public mixed $%v = null;", field.Name)
			return plan
		}
		plan.Decl = fmt.Sprintf("public %v $%v = null;", t, field.Name)
		return plan

	case field.Type == core.FieldTypeCollection || field.Type == core.FieldTypeCollectionNullable ||
		field.Type == core.FieldTypeArray || field.Type == core.FieldTypeArrayNullable ||
		field.Type == core.FieldTypeList || field.Type == core.FieldTypeListNullable ||
		field.Type == core.FieldTypeSlice || field.Type == core.FieldTypeSliceNullable:
		elemType := phpArrayElementType(field, nestedClassName)
		plan.Doc = fmt.Sprintf("/** @var %v[] */", elemType)
		plan.ArrayItemClass = phpArrayItemClass(field, nestedClassName)
		if nullable {
			plan.Decl = fmt.Sprintf("public ?array $%v = null;", field.Name)
		} else {
			plan.Decl = fmt.Sprintf("public array $%v = [];", field.Name)
		}
		return plan

	case field.Type == core.FieldTypeMap || field.Type == core.FieldTypeMapNullable:
		plan.Doc = fmt.Sprintf("/** @var array<%v> */", phpArrayElementType(field, nestedClassName))
		if nullable {
			plan.Decl = fmt.Sprintf("public ?array $%v = null;", field.Name)
		} else {
			plan.Decl = fmt.Sprintf("public array $%v = [];", field.Name)
		}
		return plan

	case field.Type == core.FieldTypeObject || field.Type == core.FieldTypeObjectNullable:
		selfReferencing := nestedClassName == selfClassName
		if nullable || selfReferencing {
			plan.TypeHint = "?" + nestedClassName
			plan.Decl = fmt.Sprintf("public ?%v $%v = null;", nestedClassName, field.Name)
			return plan
		}
		// `new X()` isn't a constant expression - assigned in the
		// constructor instead.
		plan.Decl = fmt.Sprintf("public %v $%v;", nestedClassName, field.Name)
		plan.CtorInit = fmt.Sprintf("$this->%v = new %v();", field.Name, nestedClassName)
		return plan

	default:
		// Primitive (string/int/float/bool/mixed).
		if typeHint == "mixed" {
			plan.Decl = fmt.Sprintf("public mixed $%v = null;", field.Name)
			return plan
		}
		if nullable {
			// typeHint is already `?`-prefixed for a nullable primitive
			// (see phpFieldType) - use it as-is instead of prefixing again.
			plan.Decl = fmt.Sprintf("public %v $%v = null;", typeHint, field.Name)
			return plan
		}
		zero := phpZeroValue(typeHint)
		plan.Decl = fmt.Sprintf("public %v $%v = %v;", typeHint, field.Name, zero)
		return plan
	}
}

func phpZeroValue(phpType string) string {
	switch phpType {
	case "string":
		return "''"
	case "int":
		return "0"
	case "float":
		return "0.0"
	case "bool":
		return "false"
	default:
		return "null"
	}
}

// phpLiteral renders an arbitrary go value (coming from EmiField.Default) as
// a PHP literal expression, when it's a simple, unambiguous scalar -
// returns ok=false for anything else, so the caller falls back to the
// type-based default instead of emitting something invalid.
func phpLiteral(v any) (string, bool) {
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
// resolvable via RecognizedComplexes - anything else falls back to `mixed`.
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
