package python

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// PythonStandaloneEnum renders a module-level `enums:` entry as a python
// `str, Enum` class, mirroring lib/js/js-enum.go and the kotlin equivalent.
func PythonStandaloneEnum(enum core.EmiEnum, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "enum", Objects: []string{"Enum"}},
		},
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: enum.GetName()},
			{Name: TOKEN_ROOT_CLASS, Value: enum.GetName()},
		},
	}

	const tmpl = `
class {{ .enum.GetName }}(str, Enum):
{{ if not .enum.Fields }}    pass
{{ end -}}
{{- range .enum.Fields }}
    {{ .GetKey }} = "{{ .Key }}"
{{- if .Description }}  # {{ .Description }}{{ end }}
{{ end }}
`

	t := template.Must(template.New("pyenum").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"enum": enum}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(enum.GetName())
	res.SuggestedExtension = ".py"

	return res, nil
}
