package php

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type phpPathParam struct {
	WireName string // the raw `:name` placeholder in the url template
	Field    string // camelCase PHP property name
	Type     string
}

var phpPathParamTmpl = template.Must(template.New("phppathparams").Parse(`
class {{ .TypeName }}
{
{{- range .Params }}
    public {{ .Type }} ${{ .Field }};
{{- end }}

    public function __construct({{ range $i, $p := .Params }}{{ if $i }}, {{ end }}{{ $p.Type }} ${{ $p.Field }}{{ end }})
    {
{{- range .Params }}
        $this->{{ .Field }} = ${{ .Field }};
{{- end }}
    }

    public function apply(string $templateUrl): string
    {
        $url = $templateUrl;
{{- range .Params }}
        $url = str_replace(':{{ .WireName }}', (string) $this->{{ .Field }}, $url);
{{- end }}
        return $url;
    }
}
`))

// PhpActionPathParams extracts `/:id`-style placeholders out of an action's
// url and renders a small class to carry them, plus an apply(url) method
// that substitutes them back into the template url. Returns (nil, nil) when
// the url has no placeholders at all.
func PhpActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]phpPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		if t == "" {
			t = "string"
		}
		params = append(params, phpPathParam{WireName: p.Original, Field: core.ToLower(core.NormaliseKey(p.Original)), Type: t})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	var buf bytes.Buffer
	if err := phpPathParamTmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  typeName,
		SuggestedExtension: ".php",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}, nil
}
