package php

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type phpEnumCase struct {
	Identifier  string
	Value       string
	Description string
}

type phpRenderedEnum struct {
	Name  string
	Cases []phpEnumCase
}

var phpEnumTmplFuncs = template.FuncMap{"phpString": escapeDoubleQuoted}

// PHP 8.1+ backed enums map directly onto a string wire value with no
// converter/annotation machinery needed at all (unlike Dart/C#/Java) -
// `Status::from('active')` and `$status->value` are exactly the (de)
// serialization primitives the Hydrator runtime needs.
var phpEnumTmpl = template.Must(template.New("phpenum").Funcs(phpEnumTmplFuncs).Parse(`
enum {{ .Name }}: string
{
{{- range .Cases }}
    case {{ .Identifier }} = {{ .Value | phpString }};{{ if .Description }} // {{ .Description }}{{ end }}
{{- end }}
}
`))

func renderPhpEnumDecl(e *phpRenderedEnum) string {
	var buf bytes.Buffer
	if err := phpEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

// phpRenderEnumFromInline builds the companion backed-enum for an inline
// `of:` enum field (no `target:`), named after the flattened prefix (see
// phpEnumBaseType).
func phpRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *phpRenderedEnum {
	cases := make([]phpEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, phpEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	return &phpRenderedEnum{Name: name, Cases: cases}
}

// PhpStandaloneEnum renders a module-level `enums:` entry.
func PhpStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := &phpRenderedEnum{Name: enum.GetName()}
	for _, item := range enum.Fields {
		rendered.Cases = append(rendered.Cases, phpEnumCase{
			Identifier:  core.NormaliseKey(item.Key),
			Value:       item.Key,
			Description: item.Description,
		})
	}

	return &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript:       []byte(renderPhpEnumDecl(rendered)),
		SuggestedFileName:  enum.GetName(),
		SuggestedExtension: ".php",
	}, nil
}
