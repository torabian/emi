package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func JsActionPathParams(action core.EmiRpcAction) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	placeholders := core.ExtractPlaceholdersInUrl(action.GetUrl())
	if len(placeholders) == 0 {
		return nil, nil // nothing to generate
	}

	className := action.GetName()
	typeName := fmt.Sprintf("%vPathParameter", className)

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_ROOT_CLASS, Value: typeName})
	const tmpl = `/**
 * Path parameters for {{ .ClassName }}
 */
export type {{ .TypeName }} = {
{{- range .Params }}
	{{ . }}: string | number | boolean;
{{- end }}
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
