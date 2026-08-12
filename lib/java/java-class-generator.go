package java

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in the sibling generators: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real type
// instead of falling back to `Object`.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type javaRenderedClass struct {
	ClassName   string
	Doc         string
	Fields      []javaFieldPlan
	InlineEnums []*javaRenderedEnum
}

// JavaCommonObjectContext configures a single class-generation pass - used
// for dtos, request/response bodies, and any inline object/array field.
type JavaCommonObjectContext struct {
	RootClassName       string
	RecognizedComplexes []RecognizedComplex
}

// renderClasses walks the field tree, flattening every object/array-of-object
// field into its own companion class, and every inline `enum: of:` field
// into its own companion enum - both named by concatenating the parent
// chain, exactly like the sibling generators. Unlike those, every one of
// them ends up `public` here: Java requires the file name to match its one
// public type, so (see JavaCommonObjectGenerator) each entry in the
// returned slice is rendered into its *own* file rather than bundled
// together - there's no package-private "just this file" visibility to fall
// back to if a consumer outside the `emisdk` package needs to construct or
// reference one directly, which they very much do (nested dto shapes, list
// element types, inline enum values...).
func renderClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []javaRenderedClass {
	fieldPlans := make([]javaFieldPlan, 0, len(fields))
	inlineEnums := []*javaRenderedEnum{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := javaResolveField(f, prefixName, prefixName, complexes)
		fieldPlans = append(fieldPlans, plan)
		if plan.InlineEnum != nil {
			inlineEnums = append(inlineEnums, plan.InlineEnum)
		}
	}

	current := javaRenderedClass{
		ClassName:   prefixName,
		Doc:         className,
		Fields:      fieldPlans,
		InlineEnums: inlineEnums,
	}

	out := []javaRenderedClass{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			out = append(out, renderClasses(f.Fields, core.ToUpper(f.Name), childPrefix, complexes)...)
		}
	}

	out = append(out, current)
	return out
}

var javaClassTmpl = template.Must(template.New("javaclass").Parse(`
/** {{ .Doc }} */
public class {{ .ClassName }} {
{{- range .Fields }}
    {{ .Attribute }}
    {{ .Decl }}
{{- end }}
}
`))

func renderJavaClass(c javaRenderedClass) (string, error) {
	var buf bytes.Buffer
	if err := javaClassTmpl.Execute(&buf, c); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// JavaCommonObjectGenerator generates one compiled chunk *per* type (the
// root, plus any flattened nested object/array-of-object, plus any inline
// enum) out of a field tree - each one is its own public top-level Java
// type and must land in its own file. Used for dtos, action
// request/response bodies, and headers-adjacent inline shapes.
func JavaCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, jctx JavaCommonObjectContext) ([]*core.CodeChunkCompiled, error) {
	var complexDeps []core.CodeChunkDependency
	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, jctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		complexDeps = append(complexDeps, core.CodeChunkDependency{Location: location})
	}

	classes := renderClasses(fields, jctx.RootClassName, jctx.RootClassName, jctx.RecognizedComplexes)

	chunks := make([]*core.CodeChunkCompiled, 0, len(classes))
	for _, c := range classes {
		script, err := renderJavaClass(c)
		if err != nil {
			return nil, err
		}
		chunk := &core.CodeChunkCompiled{
			ActualScript:          []byte(script),
			SuggestedFileName:     c.ClassName,
			SuggestedExtension:    ".java",
			CodeChunkDependensies: complexDeps,
		}
		if c.ClassName == jctx.RootClassName {
			chunk.Tokens = append(chunk.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: jctx.RootClassName})
		}
		chunks = append(chunks, chunk)

		for _, e := range c.InlineEnums {
			enumScript := renderJavaEnumDecl(&javaRenderedEnum{Name: e.Name, IsPublic: true, Cases: e.Cases})
			chunks = append(chunks, &core.CodeChunkCompiled{
				ActualScript:       []byte(enumScript),
				SuggestedFileName:  e.Name,
				SuggestedExtension: ".java",
			})
		}
	}

	return chunks, nil
}

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

func findComplexLocation(complexName string, complexes []RecognizedComplex) string {
	for _, item := range complexes {
		if item.Symbol == complexName {
			return item.ImportLocation
		}
	}
	return ""
}

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_ORIGINAL_NAME = core.TOKEN_ORIGINAL_NAME
