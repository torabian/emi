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
	CliName          string
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
			CliName:          FieldToCliName(item.Original),
		})
	}

	if len(placeholders) > 0 {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Location: "github.com/urfave/cli/v3",
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

func Get{{ .TypeName }}CliFlags(prefix string) []emigo.CliFlag {

	return []emigo.CliFlag{
		{{ range .Params }}
		{
			Name: prefix + "pp-{{ .CliName }}",
			Type: "{{ .GolangType }}",
			Required: true,
		},
		{{ end }}
	}
}


// Converts a placeholder url, and applies the parameters to it.
func {{ .TypeName }}Apply(params {{ .TypeName }}, templateUrl string) string {
	{{- range .Params }}
		templateUrl = strings.ReplaceAll(templateUrl, ":{{ .PlaceHolderValue }}", fmt.Sprintf("%v", params.{{.GolangFieldName}}))
	{{- end }}

	return templateUrl
}


// Extracts the path parameter from a gin request context
func {{ .TypeName }}FromGin(g *gin.Context) {{ .TypeName }} {
	return {{ .TypeName }}FromFn(func (key string) string {
		return g.Param(key) 
	})
}

// Extracts the path parameter from a urfave v3 cli.
func {{ .TypeName }}FromCli(c *cli.Command) {{ .TypeName }} {
	return {{ .TypeName }}FromFn(func (key string) string {

		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key) 
	})
}

// General purpose to extract the value and cast based on type.
func {{ .TypeName }}FromFn(fn func(key string) string) {{ .TypeName }} {
	res := {{ .TypeName }}{}
	{{- range .Params }}
		{{- $v := printf "fn(\"%s\")" .PlaceHolderValue }}
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

	// It uses fmt.Sprintf, to handle all type of data.
	if len(placeholders) > 0 {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, []core.CodeChunkDependency{
			{Location: "fmt"},
			{Location: "strings"},
		}...)
	}

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
