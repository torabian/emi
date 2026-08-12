package dart

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type dartPathParam struct {
	Original string
	Type     string
}

var dartPathParamTmpl = template.Must(template.New("dartpathparams").Parse(`
class {{ .TypeName }} {
{{- range .Params }}
  {{ .Type }} {{ .Original }};
{{- end }}

  {{ .TypeName }}({
{{- range .Params }}
    required this.{{ .Original }},
{{- end }}
  });

  String apply(String templateUrl) {
    var url = templateUrl;
{{- range .Params }}
    url = url.replaceAll(':{{ .Original }}', {{ .Original }}.toString());
{{- end }}
    return url;
  }
}
`))

// DartActionPathParams extracts `/:id`-style placeholders out of an action's
// url and renders a small class to carry them, plus an `apply(url)` method
// that substitutes them back into the template url. Returns (nil, nil) when
// the url has no placeholders at all.
func DartActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]dartPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		if t == "" {
			t = "String"
		}
		params = append(params, dartPathParam{Original: p.Original, Type: t})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	var buf bytes.Buffer
	if err := dartPathParamTmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  core.ToSnakeCase(typeName),
		SuggestedExtension: ".dart",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}, nil
}
