package golang

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type GoPathParamer struct {
	PlaceHolderValue string
	GolangFieldName  string
}

func GoActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Location: "strings",
			},
		},
	}

	placeholders0 := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders0) == 0 {
		return nil, nil // nothing to generate
	}

	placeholders := []GoPathParamer{}
	for _, item := range placeholders0 {
		placeholders = append(placeholders, GoPathParamer{
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
type {{ .TypeName }} struct {
{{- range .Params }}
	{{ .GolangFieldName }} string
{{- end }}
}

// Converts a placeholder url, and applies the parameters to it.
func {{ .TypeName }}Apply(params {{ .TypeName }}, templateUrl string) string {
	{{- range .Params }}
		templateUrl = strings.ReplaceAll(templateUrl, "{{ .PlaceHolderValue }}", params.{{.GolangFieldName}})
	{{- end }}

	return templateUrl
}
// Creates the parameters from the gin
func {{ .TypeName }}FromGin(g *gin.Context) {{ .TypeName }} {
	res := {{ .TypeName }}{}
	{{- range .Params }}
		res.{{.GolangFieldName}} = g.Param("{{.PlaceHolderValue}}")
	{{- end }}

	return res
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
