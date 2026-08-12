// Renders the fields of a single generated python dataclass (dto, request,
// response, nested object...). Mirrors lib/kotlin/kotlin-common-fields.go and
// lib/js/js-common-fields.go in spirit, but python needs far less ceremony:
// no getters/setters, just an annotated `name: Type = default` field line.
package python

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type pyRenderedField struct {
	Name      string
	FieldLine string
}

// pyFieldNestedClassName computes the PascalCase name a field would use if it
// turned out to be a locally-nested object/array-of-object, following the
// same "ParentChain + FieldName" flattening convention used by kotlin/js.
func pyFieldNestedClassName(field *core.EmiField, prefixName string) string {
	return prefixName + core.ToUpper(field.Name)
}

// collectComplexSymbols finds every `+`-prefixed complex reference in the
// field tree so the caller can resolve and import them (see
// findComplexLocation / RecognizedComplexes).
func collectComplexSymbols(fields []*core.EmiField) []string {
	var result []string
	var walk func([]*core.EmiField)
	walk = func(f []*core.EmiField) {
		for _, field := range f {
			if field == nil {
				continue
			}
			if strings.Contains(field.Complex, "+") {
				result = append(result, strings.ReplaceAll(field.Complex, "+", ""))
			}
			if len(field.Fields) > 0 {
				walk(field.Fields)
			}
		}
	}
	walk(fields)
	return result
}

// fieldsUseLiteralEnum reports whether any field in the tree is an inline
// `enum`/`enum?` (no `target:`, at least one `of:` value) - those render as
// `typing.Literal[...]`, which needs its own import, unlike a `target:`
// enum reference (which just imports the generated Enum class by name).
func fieldsUseLiteralEnum(fields []*core.EmiField) bool {
	for _, field := range fields {
		if field == nil {
			continue
		}
		if (field.Type == core.FieldTypeEnum || field.Type == core.FieldTypeEnumNullable) &&
			field.Target == "" && pythonEnumLiteralType(field) != "" {
			return true
		}
		if len(field.Fields) > 0 && fieldsUseLiteralEnum(field.Fields) {
			return true
		}
	}
	return false
}

func findComplexLocation(complexName string, complexes []RecognizedComplex) string {
	for _, item := range complexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}
	return ""
}

// pyRenderField renders a single field line, e.g.:
//
//	name: Optional[str] = None  # the user's display name
//
// complexes is only used to decide whether a `complex:` reference is
// actually resolvable to an import - see resolveComplexTypeHint.
func pyRenderField(field *core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) pyRenderedField {
	nestedClassName := pyFieldNestedClassName(field, prefixName)
	typeHint := pythonFieldTypeOnNestedClasses(field, nestedClassName)
	defaultValue := PythonSafeDefaultValue(field, nestedClassName, selfClassName)

	if field.Complex != "" {
		typeHint, defaultValue = resolveComplexTypeHint(field, complexes)
	}

	line := fmt.Sprintf("%v: %v = %v", field.Name, typeHint, defaultValue)

	if field.Description != "" {
		line += fmt.Sprintf("  # %v", strings.ReplaceAll(field.Description, "\n", " "))
	}

	return pyRenderedField{
		Name:      field.Name,
		FieldLine: line,
	}
}

// resolveComplexTypeHint decides the type-hint/default for a `complex:`
// field. A symbol is only ever used bare (unquoted) as a type-hint when it's
// actually going to be imported (a `+`-prefixed complex resolvable via
// RecognizedComplexes) - anything else (a plain marker with no known import,
// or a dotted go-style qualified name like `big.Int`) would blow up
// `typing.get_type_hints()` at deserialization time with a NameError, so it
// safely falls back to `Any` instead.
func resolveComplexTypeHint(field *core.EmiField, complexes []RecognizedComplex) (typeHint string, defaultValue string) {
	symbol := strings.ReplaceAll(field.Complex, "+", "")

	// A single complex value can never be safely default-constructed (we
	// don't know its shape), so it's always optional-with-None, exactly
	// like a `one`/`class` relation.
	if strings.Contains(field.Complex, "+") && findComplexLocation(symbol, complexes) != "" {
		return fmt.Sprintf("Optional[%v]", symbol), "None"
	}

	return "Optional[Any]", "None"
}

func pyRenderFieldsShallow(fields []*core.EmiField, prefixName string, selfClassName string, complexes []RecognizedComplex) []pyRenderedField {
	out := make([]pyRenderedField, 0, len(fields))
	for _, f := range fields {
		if f != nil {
			out = append(out, pyRenderField(f, prefixName, selfClassName, complexes))
		}
	}
	return out
}
