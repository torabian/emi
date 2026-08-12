package dart

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type dartEnumCase struct {
	Identifier  string
	Value       string
	Description string
}

type dartRenderedEnum struct {
	Name  string
	Cases []dartEnumCase
}

var dartEnumTmplFuncs = template.FuncMap{"escapeDartString": escapeDartString}

var dartEnumTmpl = template.Must(template.New("dartenum").Funcs(core.CommonMap).Funcs(dartEnumTmplFuncs).Parse(`
enum {{ .Name }} {
{{- range .Cases }}
  {{ .Identifier }}({{ .Value | escapeDartString }}){{ if .Description }} // {{ .Description }}{{ end }},
{{- end }}
  ;

  final String value;
  const {{ .Name }}(this.value);

  static {{ .Name }} fromValue(String value) =>
      {{ .Name }}.values.firstWhere((e) => e.value == value, orElse: () => {{ .Name }}.values.first);

  static {{ .Name }}? fromValueOrNull(String? value) => value == null ? null : fromValue(value);
}
`))

func renderDartEnumDecl(e *dartRenderedEnum) string {
	var buf bytes.Buffer
	if err := dartEnumTmpl.Execute(&buf, e); err != nil {
		return ""
	}
	return buf.String()
}

func escapeDartString(s string) string {
	return escapeDoubleQuoted(s)
}

// dartRenderEnumFromInline builds the companion enum for an inline `of:`
// enum field (no `target:`), named after the flattened prefix (see
// dartEnumBaseType).
func dartRenderEnumFromInline(name string, ofType []*core.EmiEnumInline) *dartRenderedEnum {
	cases := make([]dartEnumCase, 0, len(ofType))
	for _, item := range ofType {
		if item == nil {
			continue
		}
		cases = append(cases, dartEnumCase{
			Identifier:  core.ToLower(core.NormaliseKey(item.Key)),
			Value:       item.Key,
			Description: item.Description,
		})
	}
	return &dartRenderedEnum{Name: name, Cases: cases}
}

// DartStandaloneEnum renders a module-level `enums:` entry.
func DartStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	rendered := &dartRenderedEnum{Name: enum.GetName()}
	for _, item := range enum.Fields {
		rendered.Cases = append(rendered.Cases, dartEnumCase{
			Identifier:  core.ToLower(core.NormaliseKey(item.Key)),
			Value:       item.Key,
			Description: item.Description,
		})
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
		ActualScript:       []byte(renderDartEnumDecl(rendered)),
		SuggestedFileName:  core.ToSnakeCase(enum.GetName()),
		SuggestedExtension: ".dart",
	}
	return res, nil
}
