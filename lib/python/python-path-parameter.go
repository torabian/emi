package python

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type pyPathParam struct {
	Original string
	Type     string
}

// PythonActionPathParams extracts `/:id`-style placeholders out of an
// action's url and renders a small dataclass to carry them, plus an
// `apply_url` helper that substitutes them back into the template url.
// Returns (nil, nil) when the url has no placeholders at all.
func PythonActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]pyPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		if t == "" {
			t = "str"
		}
		params = append(params, pyPathParam{Original: p.Original, Type: t})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "dataclasses", Objects: []string{"dataclass"}},
		},
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}

	const tmpl = `
@dataclass
class {{ .TypeName }}:
{{- range .Params }}
    {{ .Original }}: {{ .Type }}
{{- end }}

    def apply(self, template_url: str) -> str:
        url = template_url
{{- range .Params }}
        url = url.replace(":{{ .Original }}", str(self.{{ .Original }}))
{{- end }}
        return url
`

	t := template.Must(template.New("pypathparams").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(typeName)
	res.SuggestedExtension = ".py"

	return res, nil
}
