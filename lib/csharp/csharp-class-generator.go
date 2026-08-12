package csharp

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in the sibling generators: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real type
// instead of falling back to `object`.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type csRenderedClass struct {
	ClassName   string
	Doc         string
	Fields      []csFieldPlan
	InlineEnums []*csRenderedEnum
}

// CSharpCommonObjectContext configures a single class-generation pass - used
// for dtos, request/response bodies, and any inline object/array field.
type CSharpCommonObjectContext struct {
	RootClassName       string
	RecognizedComplexes []RecognizedComplex
}

// renderClasses walks the field tree, flattening every object/array-of-object
// field into its own top-level class, and every inline `enum: of:` field
// into its own top-level enum - both named by concatenating the parent
// chain, exactly like the sibling generators.
func renderClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []csRenderedClass {
	fieldPlans := make([]csFieldPlan, 0, len(fields))
	inlineEnums := []*csRenderedEnum{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := csResolveField(f, prefixName, prefixName, complexes)
		fieldPlans = append(fieldPlans, plan)
		if plan.InlineEnum != nil {
			inlineEnums = append(inlineEnums, plan.InlineEnum)
		}
	}

	current := csRenderedClass{
		ClassName:   prefixName,
		Doc:         className,
		Fields:      fieldPlans,
		InlineEnums: inlineEnums,
	}

	out := []csRenderedClass{}
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

var csClassTmplFuncs = template.FuncMap{"renderEnum": renderCsEnumDecl}

var csClassTmpl = template.Must(template.New("csclass").Funcs(csClassTmplFuncs).Parse(`
{{ range .classes }}
/// <summary>{{ .Doc }}</summary>
public class {{ .ClassName }}
{
{{- range .Fields }}
    {{ range .Attributes }}{{ . }} {{ end }}
    public {{ .TypeHint }} {{ .Name }} { get; set; }{{ if .Initializer }} {{ .Initializer }}{{ end }}
{{- end }}
}
{{ range .InlineEnums }}
{{ renderEnum . }}
{{ end }}
{{ end }}
`))

func renderCsClasses(classes []csRenderedClass) (string, error) {
	var buf bytes.Buffer
	if err := csClassTmpl.Execute(&buf, core.H{"classes": classes}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// CSharpCommonObjectGenerator generates one or more class definitions (the
// root, plus any flattened nested object/array-of-object, plus any inline
// enum) out of a field tree. Used for dtos, action request/response bodies,
// and headers-adjacent inline shapes.
func CSharpCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, cctx CSharpCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, cctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}

	classes := renderClasses(fields, cctx.RootClassName, cctx.RootClassName, cctx.RecognizedComplexes)
	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: cctx.RootClassName})
	}

	script, err := renderCsClasses(classes)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = cctx.RootClassName
	res.SuggestedExtension = ".cs"
	return res, nil
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
