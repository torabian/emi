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
	GolangType       string
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

		golangType := goPrimitiveDetect(item.Type)

		// If the type is not string, str conv is necessary to happen
		if golangType != "string" {
			res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
				Location: "strconv",
			})
		}

		placeholders = append(placeholders, GoPathParamer{
			PlaceHolderValue: item.Original,
			GolangFieldName:  core.ToUpper(item.Original),
			GolangType:       golangType,
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
	{{ .GolangFieldName }} {{ .GolangType }}
{{- end }}
}

// Converts a placeholder url, and applies the parameters to it.
func {{ .TypeName }}Apply(params {{ .TypeName }}, templateUrl string) string {
	{{- range .Params }}
		templateUrl = strings.ReplaceAll(templateUrl, "{{ .PlaceHolderValue }}", fmt.Sprintf("%v", params.{{.GolangFieldName}}))
	{{- end }}

	return templateUrl
}

// Creates the parameters from the gin
// Creates the parameters from the gin
func {{ .TypeName }}FromGin(g *gin.Context) {{ .TypeName }} {
	res := {{ .TypeName }}{}
	{{- range .Params }}
		{{- $v := printf "g.Param(\"%s\")" .PlaceHolderValue }}
		{{- if or (eq .GolangType "int") (eq .GolangType "int8") (eq .GolangType "int16") (eq .GolangType "int32") (eq .GolangType "int64") }}
			if v := {{ $v }}; v != "" {
				{{- if eq .GolangType "int" }}
					res.{{.GolangFieldName}}, _ = strconv.Atoi(v)
				{{- else if eq .GolangType "int8" }}
					t, _ := strconv.ParseInt(v, 10, 8)
					res.{{.GolangFieldName}} = int8(t)
				{{- else if eq .GolangType "int16" }}
					t, _ := strconv.ParseInt(v, 10, 16)
					res.{{.GolangFieldName}} = int16(t)
				{{- else if eq .GolangType "int32" }}
					t, _ := strconv.ParseInt(v, 10, 32)
					res.{{.GolangFieldName}} = int32(t)
				{{- else if eq .GolangType "int64" }}
					t, _ := strconv.ParseInt(v, 10, 64)
					res.{{.GolangFieldName}} = t
				{{- end }}
			}
		{{- else if or (eq .GolangType "float32") (eq .GolangType "float64") }}
			if v := {{ $v }}; v != "" {
				f, _ := strconv.ParseFloat(v, {{ if eq .GolangType "float32" }}32{{ else }}64{{ end }})
				{{- if eq .GolangType "float32" }}
					res.{{.GolangFieldName}} = float32(f)
				{{- else }}
					res.{{.GolangFieldName}} = f
				{{- end }}
			}
		{{- else if eq .GolangType "bool" }}
			if v := {{ $v }}; v != "" {
				res.{{.GolangFieldName}}, _ = strconv.ParseBool(v)
			}
		{{- else }}
			res.{{.GolangFieldName}} = {{ $v }}
		{{- end }}
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
