package kotlin

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type KotlinPathParamer struct {
	PlaceHolderValue string
	GolangFieldName  string
}

func KotlinActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	placeholders0 := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders0) == 0 {
		return nil, nil // nothing to generate
	}

	placeholders := []KotlinPathParamer{}
	for _, item := range placeholders0 {
		placeholders = append(placeholders, KotlinPathParamer{
			PlaceHolderValue: item,
			GolangFieldName:  core.ToUpper(item),
		})
	}

	className := action.GetName()
	typeName := fmt.Sprintf("%vPathParameter", className)

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: typeName})
	const tmpl = `/**
 * Path parameters for {{ .ClassName }}
 */
data class {{ .TypeName }} (
{{- range .Params }}
	var {{ .GolangFieldName }}: String,
{{- end }}
)

// Converts a placeholder url, and applies the parameters to it.
fun {{ .TypeName }}Apply(params: {{ .TypeName }}, templateUrl: String): String {
	var url = templateUrl
	{{- range .Params }}
		url = url.replace("{{ .PlaceHolderValue }}", params.{{.GolangFieldName}})
	{{- end }}

	return url
}

`

	t := template.Must(template.New("pathParams").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]any{
		"ClassName": className,
		"TypeName":  typeName,
		"Params":    placeholders,
	}); err != nil {
		return nil, err
	}
	res.ActualScript = buf.Bytes()

	return res, nil
}
