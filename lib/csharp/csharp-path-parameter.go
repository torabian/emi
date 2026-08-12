package csharp

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type csPathParam struct {
	WireName string // the raw `:name` placeholder in the url template
	Property string // PascalCase C# property name
	Type     string
}

var csPathParamTmpl = template.Must(template.New("cspathparams").Parse(`
public class {{ .TypeName }}
{
{{- range .Params }}
    public required {{ .Type }} {{ .Property }} { get; set; }
{{- end }}

    public string Apply(string templateUrl)
    {
        var url = templateUrl;
{{- range .Params }}
        url = url.Replace(":{{ .WireName }}", {{ .Property }}.ToString());
{{- end }}
        return url;
    }
}
`))

// CSharpActionPathParams extracts `/:id`-style placeholders out of an
// action's url and renders a small class to carry them, plus an Apply(url)
// method that substitutes them back into the template url. Returns (nil,
// nil) when the url has no placeholders at all.
func CSharpActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil
	}

	params := make([]csPathParam, 0, len(placeholders))
	for _, p := range placeholders {
		t := extractPrimitive(p.Type)
		if t == "" {
			t = "string"
		}
		params = append(params, csPathParam{WireName: p.Original, Property: core.NormaliseKey(p.Original), Type: t})
	}

	className := core.ToUpper(core.NormaliseKey(action.GetName()))
	typeName := fmt.Sprintf("%vPathParameters", className)

	var buf bytes.Buffer
	if err := csPathParamTmpl.Execute(&buf, core.H{"TypeName": typeName, "Params": params}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  typeName,
		SuggestedExtension: ".cs",
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ROOT_CLASS, Value: typeName},
		},
	}, nil
}
