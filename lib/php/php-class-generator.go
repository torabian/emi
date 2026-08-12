package php

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in the sibling generators: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real type
// instead of falling back to `mixed`.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type phpRenderedClass struct {
	ClassName   string
	Doc         string
	Fields      []phpFieldPlan
	InlineEnums []*phpRenderedEnum
}

// PhpCommonObjectContext configures a single class-generation pass - used
// for dtos, request/response bodies, and any inline object/array field.
type PhpCommonObjectContext struct {
	RootClassName       string
	RecognizedComplexes []RecognizedComplex
}

// renderClasses walks the field tree, flattening every object/array-of-object
// field into its own top-level class, and every inline `enum: of:` field
// into its own top-level backed enum - both named by concatenating the
// parent chain, exactly like the sibling generators.
func renderClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []phpRenderedClass {
	fieldPlans := make([]phpFieldPlan, 0, len(fields))
	inlineEnums := []*phpRenderedEnum{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := phpResolveField(f, prefixName, prefixName, complexes)
		fieldPlans = append(fieldPlans, plan)
		if plan.InlineEnum != nil {
			inlineEnums = append(inlineEnums, plan.InlineEnum)
		}
	}

	current := phpRenderedClass{
		ClassName:   prefixName,
		Doc:         className,
		Fields:      fieldPlans,
		InlineEnums: inlineEnums,
	}

	out := []phpRenderedClass{}
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

	// Children first, so a plain top-to-bottom read never hits a forward
	// reference (PHP doesn't actually require this to parse - class
	// declaration order never matters - it's purely for readability,
	// mirroring the other generators).
	out = append(out, current)
	return out
}

var phpClassTmplFuncs = template.FuncMap{
	"renderEnum": renderPhpEnumDecl,
	"phpKey":     escapeDoubleQuoted,
}

var phpClassTmpl = template.Must(template.New("phpclass").Funcs(phpClassTmplFuncs).Parse(`
{{ range .classes }}
/** {{ .Doc }} */
class {{ .ClassName }}
{
{{- range .Fields }}
{{ if .Doc }}    {{ .Doc }}
{{ end -}}
    {{ .Decl }}
{{- end }}
{{ if .HasArrayItemTypes }}
    /** @var array<string, class-string> element class per array-typed property, for Hydrator::fromArray() */
    public static array $arrayItemTypes = [
{{- range .Fields }}
{{- if .ArrayItemClass }}
        {{ .Name | phpKey }} => {{ .ArrayItemClass }}::class,
{{- end }}
{{- end }}
    ];
{{ end }}
{{ if .HasCtorInit }}
    public function __construct()
    {
{{- range .Fields }}
{{- if .CtorInit }}
        {{ .CtorInit }}
{{- end }}
{{- end }}
    }
{{ end }}
}
{{ range .InlineEnums }}
{{ renderEnum . }}
{{ end }}
{{ end }}
`))

func renderPhpClasses(classes []phpRenderedClass) (string, error) {
	type classView struct {
		phpRenderedClass
		HasCtorInit       bool
		HasArrayItemTypes bool
	}

	views := make([]classView, 0, len(classes))
	for _, c := range classes {
		hasCtorInit, hasArrayItemTypes := false, false
		for _, f := range c.Fields {
			if f.CtorInit != "" {
				hasCtorInit = true
			}
			if f.ArrayItemClass != "" {
				hasArrayItemTypes = true
			}
		}
		views = append(views, classView{phpRenderedClass: c, HasCtorInit: hasCtorInit, HasArrayItemTypes: hasArrayItemTypes})
	}

	var buf bytes.Buffer
	if err := phpClassTmpl.Execute(&buf, core.H{"classes": views}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// PhpCommonObjectGenerator generates one or more class definitions (the
// root, plus any flattened nested object/array-of-object, plus any inline
// enum) out of a field tree. Used for dtos, action request/response bodies,
// and headers-adjacent inline shapes.
func PhpCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, pctx PhpCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, pctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}

	classes := renderClasses(fields, pctx.RootClassName, pctx.RootClassName, pctx.RecognizedComplexes)
	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: pctx.RootClassName})
	}

	script, err := renderPhpClasses(classes)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = pctx.RootClassName
	res.SuggestedExtension = ".php"
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
