package java

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type javaPathParam struct {
	WireName string // the raw `:name` placeholder in the url template
	Field    string // lowerCamelCase Java field name
	Type     string
}

var javaPathParamTmpl = template.Must(template.New("javapathparams").Parse(`
public class {{ .TypeName }} {
{{- range .Params }}
    public {{ .Type }} {{ .Field }};
{{- end }}

    public String apply(String templateUrl) {
        var url = templateUrl;
{{- range .Params }}
        url = url.replace(":{{ .WireName }}", String.valueOf({{ .Field }}));
{{- end }}
        return url;
    }
}
`))

// JavaActionPathParams extracts `/:id`-style placeholders out of an action's
// url and renders a small (package-private) class to carry them, plus an
// apply(url) method that substitutes them back into the template url.
// Returns (nil, nil) when the url has no placeholders at all.
func JavaActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]javaPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		if t == "" {
			t = "String"
		}
		params = append(params, javaPathParam{WireName: p.Original, Field: core.ToLower(core.NormaliseKey(p.Original)), Type: t})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	var buf bytes.Buffer
	if err := javaPathParamTmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  typeName,
		SuggestedExtension: ".java",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}, nil
}
