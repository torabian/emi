package golang

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type GoQueryParameterCtx struct {
	QueryKey        string
	GolangFieldName string
	GolangType      string
	ActionName      string
}

func GoActionQueryParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {

	// If there is no query params set, we avoid it. If developer wants, manually adds it.

	// New: We devlier the query params anyway, because it can be fully dynamic as well
	// It's helping a lot.
	// if len(action.GetQuery()) == 0 {
	// 	return nil, nil
	// }

	ctx := GoQueryParameterCtx{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
	}

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Location: "net/url",
			},
		},
	}

	className := action.GetName()
	typeName := fmt.Sprintf("%vQueryParameter", className)

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: core.TOKEN_ORIGINAL_NAME, Value: typeName})

	const tmpl = `/**
 * Query parameters for {{ .ClassName }}
 */

{{ define "printthem" }}
	{{ range . }}

	 	{{ if or (eq .Type "string") (eq .Type "float64") (eq .Type "float32") (eq .Type "bool") (eq .Type "int") (eq .Type "int8") (eq .Type "int16") (eq .Type "int32") (eq .Type "int64")}}
			{{ .Name }} {{ .Type }} ` + "`json:\"{{ .Name }}\"`" + `
		{{ end }}

		{{ if (eq .Type "object") }}
			{{ .Name }} struct {
				{{ template "printthem" .Fields }}
			} ` + "`json:\"{{ .Name }}\"`" + `
		{{ end }}

		{{ if (eq .Type "array") }}
			{{ .Name }} [] struct {
				{{ template "printthem" .Fields }}
			} ` + "`json:\"{{ .Name }}\"`" + `
		{{ end }}

		{{ if (eq .Type "slice") }}
			{{ .Name }} []{{ .Primitive}} ` + "`json:\"{{ .Name }}\"`" + `
		{{ end }} 
		 
	{{ end }}
{{ end }}

// Query wrapper with private fields
type {{ .ctx.ActionName }}Query struct {
	values url.Values
	mapped map[string]interface{}

	// Typesafe fields
	{{ template "printthem" .qs}}
}


func {{ .ctx.ActionName }}QueryFromString(rawQuery string) {{ .ctx.ActionName }}Query {
	v := {{ .ctx.ActionName }}Query{}

	values, _ := url.ParseQuery(rawQuery)

	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}

	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}

	v.values = values
	v.mapped = mapped
	return v
}



func {{ .ctx.ActionName }}QueryFromGin(c *gin.Context) {{ .ctx.ActionName }}Query {
	return {{ .ctx.ActionName }}QueryFromString(c.Request.URL.RawQuery)
}

func {{ .ctx.ActionName }}QueryFromHttp(r *http.Request) {{ .ctx.ActionName }}Query {
	return {{ .ctx.ActionName }}QueryFromString(r.URL.RawQuery)
}

func (q {{ .ctx.ActionName }}Query) Values() url.Values {
	return q.values
}

func (q {{ .ctx.ActionName }}Query) Mapped() map[string]interface{} {
	return q.mapped
}

func (q *{{ .ctx.ActionName }}Query) SetValues(v url.Values) {
	q.values = v
}

func (q *{{ .ctx.ActionName }}Query) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

`

	t := template.Must(template.New("queryParams").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, map[string]any{
		"ClassName": className,
		"TypeName":  typeName,
		"ctx":       ctx,
		"qs":        action.GetQuery(),
	}); err != nil {
		return nil, err
	}
	res.ActualScript = buf.Bytes()

	return res, nil
}
