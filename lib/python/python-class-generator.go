package python

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// RecognizedComplex mirrors the same concept in lib/js and lib/kotlin: a
// custom/complex data type that's importable from somewhere, so a field
// referencing it (via `complex: "+Vector3"`) can be resolved to a real
// import instead of falling back to `Any`.
type RecognizedComplex struct {
	Symbol         string
	ImportLocation string
}

type pyRenderedClass struct {
	ClassName  string
	Doc        string
	Fields     []pyRenderedField
	SubClasses []pyRenderedClass
}

// PyCommonObjectContext configures a single class-generation pass - used for
// dtos, request/response bodies, and any inline object/array field.
type PyCommonObjectContext struct {
	// The class name used as the root, and as the flattening prefix for any
	// nested object/array-of-object field (RootClassName + FieldName).
	RootClassName string

	// Complex types recognized (and importable) in the current module.
	RecognizedComplexes []RecognizedComplex
}

// renderClasses walks the field tree, flattening every object/array-of-object
// field into its own top-level dataclass, named by concatenating the parent
// chain (exactly like lib/kotlin/kotlin-class-generator.go#renderClasses).
func renderClasses(fields []*core.EmiField, className string, prefixName string, ctx core.MicroGenContext, complexes []RecognizedComplex) []pyRenderedClass {
	doc := NewPyDoc("    ").Add(fmt.Sprintf("%v", className)).String()

	fieldsRendered := pyRenderFieldsShallow(fields, prefixName, prefixName, complexes)

	current := pyRenderedClass{
		ClassName: prefixName,
		Doc:       doc,
		Fields:    fieldsRendered,
	}

	for _, f := range fields {
		if f == nil {
			continue
		}
		// Every field whose type-hint can resolve to nestedClassName (see
		// pythonDataStructureType) must get its class generated here - even
		// with zero sub-fields (an empty `pass`-bodied dataclass) - or the
		// reference in the parent's type-hint would dangle and blow up
		// `typing.get_type_hints()` at deserialization time.
		if f.Type == core.FieldTypeObject || f.Type == core.FieldTypeObjectNullable ||
			f.Type == core.FieldTypeArray || f.Type == core.FieldTypeArrayNullable ||
			f.Type == core.FieldTypeList || f.Type == core.FieldTypeListNullable {
			childPrefix := prefixName + core.ToUpper(f.Name)
			current.SubClasses = append(current.SubClasses, renderClasses(f.Fields, core.ToUpper(f.Name), childPrefix, ctx, complexes)...)
		}
	}

	// Children first: they must be defined before the parent references them
	// textually, since (unlike TS/Kotlin) plain assignment-time evaluation of
	// dataclass field defaults such as `field(default_factory=lambda: X())`
	// only needs X to exist by the time the *module* finishes importing, but
	// keeping children-before-parent keeps the file readable top-down without
	// relying on `from __future__ import annotations` alone.
	out := []pyRenderedClass{}
	for _, sub := range current.SubClasses {
		out = append(out, sub)
	}
	current.SubClasses = nil
	out = append(out, current)

	return out
}

// PythonCommonObjectGenerator generates one or more `@dataclass` definitions
// (the root, plus any flattened nested object/array-of-object) out of a field
// tree. Used for dtos, action request/response bodies, and headers-adjacent
// inline shapes.
func PythonCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, pyctx PyCommonObjectContext) (*core.CodeChunkCompiled, error) {
	typingObjects := []string{"Optional", "List", "Dict", "Any"}
	if fieldsUseLiteralEnum(fields) {
		typingObjects = append(typingObjects, "Literal")
	}

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "dataclasses", Objects: []string{"dataclass", "field"}},
			{Location: "typing", Objects: typingObjects},
		},
	}

	for _, symbol := range collectComplexSymbols(fields) {
		location := findComplexLocation(symbol, pyctx.RecognizedComplexes)
		if location == "" {
			continue
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{symbol},
			Location: location,
		})
	}

	for _, target := range core.CollectTargets(fields, pyctx.RootClassName) {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, DtoNameToImportDependency(target)...)
	}

	classes := renderClasses(fields, pyctx.RootClassName, pyctx.RootClassName, ctx, pyctx.RecognizedComplexes)

	if len(classes) > 0 {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: pyctx.RootClassName})
	}

	const tmpl = `
{{ range .classes }}
@dataclass
class {{ .ClassName }}:
{{ if .Doc }}{{ .Doc }}
{{ end }}
{{ if not .Fields }}    pass
{{ end }}
{{- range .Fields }}
    {{ .FieldLine }}
{{ end }}

{{ end }}
`

	t := template.Must(template.New("pyclass").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"classes": classes}); err != nil {
		return nil, err
	}

	res.ActualScript = []byte(strings.TrimRight(buf.String(), "\n") + "\n")
	res.SuggestedFileName = core.ToSnakeCase(pyctx.RootClassName)
	res.SuggestedExtension = ".py"

	return res, nil
}

// DtoNameToImportDependency turns a `target: SomeDto` / `dto: SomeDto|OtherDto`
// style reference into import dependencies pointing at the generated module
// for each name (dto files are always named after `ToSnakeCase(ClassName)`).
func DtoNameToImportDependency(dtoName string) []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{}
	for name := range strings.SplitSeq(dtoName, "|") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		_, className := core.ParseDtoPath(name)
		deps = append(deps, core.CodeChunkDependency{
			Objects:  []string{className},
			Location: "." + core.ToSnakeCase(className),
		})
	}
	return deps
}

var TOKEN_ROOT_CLASS = "root.class"
var TOKEN_ORIGINAL_NAME = core.TOKEN_ORIGINAL_NAME
