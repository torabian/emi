package dart

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in lib/js/lib/kotlin/lib/python:
// a custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real
// import instead of falling back to `dynamic`.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type dartRenderedClass struct {
	ClassName   string
	Doc         string
	Fields      []dartFieldPlan
	InlineEnums []*dartRenderedEnum
}

// DartCommonObjectContext configures a single class-generation pass - used
// for dtos, request/response bodies, and any inline object/array field.
type DartCommonObjectContext struct {
	RootClassName       string
	RecognizedComplexes []RecognizedComplex
}

// renderClasses walks the field tree, flattening every object/array-of-object
// field into its own top-level class, and every inline `enum: of:` field
// into its own top-level enum - both named by concatenating the parent chain
// (exactly like lib/kotlin and lib/python's equivalents).
func renderClasses(fields []*core.EmiField, className string, prefixName string, complexes []RecognizedComplex) []dartRenderedClass {
	doc := NewDartDoc("").Add(className).String()

	fieldPlans := make([]dartFieldPlan, 0, len(fields))
	inlineEnums := []*dartRenderedEnum{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		plan := dartResolveField(f, prefixName, prefixName, complexes)
		fieldPlans = append(fieldPlans, plan)
		if plan.InlineEnum != nil {
			inlineEnums = append(inlineEnums, plan.InlineEnum)
		}
	}

	current := dartRenderedClass{
		ClassName:   prefixName,
		Doc:         doc,
		Fields:      fieldPlans,
		InlineEnums: inlineEnums,
	}

	out := []dartRenderedClass{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		// Every field whose type can resolve to a flattened nested class
		// name must get that class generated here - even with zero
		// sub-fields - or the reference would dangle and fail to compile.
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			out = append(out, renderClasses(f.Fields, core.ToUpper(f.Name), childPrefix, complexes)...)
		}
	}

	// Children before the parent, so a plain top-to-bottom read never hits
	// a forward reference (dart doesn't need this to compile - unlike
	// python it has no lazy-annotation trick, but declaration order simply
	// doesn't matter to dartanalyzer/dart2js for top-level classes either -
	// this is purely for readability, mirroring the other generators).
	out = append(out, current)
	return out
}

var dartClassTmplFuncs = template.FuncMap{
	"dartJsonKey": escapeDoubleQuoted,
	"renderEnum":  renderDartEnumDecl,
}

var dartClassTmpl = template.Must(template.New("dartclass").Funcs(core.CommonMap).Funcs(dartClassTmplFuncs).Parse(`
{{ range .classes }}
{{ if .Doc }}{{ .Doc }}
{{ end -}}
class {{ .ClassName }} {
{{- range .Fields }}
  {{ .FieldDecl }}
{{- end }}

  {{ .ClassName }}({{ if .Fields }}{
{{- range .Fields }}
    {{ .CtorParam }}
{{- end }}
  }{{ end }}){{ if .HasInitializers }}
      : {{ .InitializerList }}{{ end }};

  factory {{ .ClassName }}.fromJson(Map<String, dynamic> json) => {{ .ClassName }}(
{{- range .Fields }}
    {{ .Name }}: {{ .FromJson }},
{{- end }}
  );

  Map<String, dynamic> toJson() => {
{{- range .Fields }}
    {{ .Name | dartJsonKey }}: {{ .ToJson }},
{{- end }}
  };
}
{{ range .InlineEnums }}
{{ renderEnum . }}
{{ end }}
{{ end }}
`))

func renderDartClasses(classes []dartRenderedClass) (string, error) {
	type classView struct {
		dartRenderedClass
		HasInitializers bool
		InitializerList string
	}

	views := make([]classView, 0, len(classes))
	for _, c := range classes {
		initializers := []string{}
		for _, f := range c.Fields {
			if f.NeedsInitializer() {
				initializers = append(initializers, fmt.Sprintf("%v = %v", f.Name, f.InitializerExpr))
			}
		}
		views = append(views, classView{
			dartRenderedClass: c,
			HasInitializers:   len(initializers) > 0,
			InitializerList:   strings.Join(initializers, ",\n        "),
		})
	}

	var buf bytes.Buffer
	if err := dartClassTmpl.Execute(&buf, core.H{"classes": views}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// DartCommonObjectGenerator generates one or more class definitions (the
// root, plus any flattened nested object/array-of-object, plus any inline
// enum) out of a field tree. Used for dtos, action request/response bodies,
// and headers-adjacent inline shapes.
func DartCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, dctx DartCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, dctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: location})
	}

	for _, target := range core.CollectTargets(fields, dctx.RootClassName) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, DtoNameToImportDependency(target)...)
	}

	classes := renderClasses(fields, dctx.RootClassName, dctx.RootClassName, dctx.RecognizedComplexes)
	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: dctx.RootClassName})
	}

	script, err := renderDartClasses(classes)
	if err != nil {
		return nil, err
	}

	res.ActualScript = []byte(script)
	res.SuggestedFileName = core.ToSnakeCase(dctx.RootClassName)
	res.SuggestedExtension = ".dart"
	return res, nil
}

// DtoNameToImportDependency turns a `target: SomeDto` / `dto: SomeDto|OtherDto`
// style reference into import dependencies pointing at the generated file
// for each name (files are always named after `ToSnakeCase(ClassName)+".dart"`).
func DtoNameToImportDependency(dtoName string) []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}
	for name := range strings.SplitSeq(dtoName, "|") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		_, className := core.ParseDtoPath(name)
		deps = append(deps, core.CodeChunkDependency{Location: core.ToSnakeCase(className) + ".dart"})
	}
	return deps
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
