package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// JsActionPathParams generates a TypeScript type with the path parameters only
func JsActionPathParams(action *core.Module3Action) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	placeholders := core.ExtractPlaceholdersInUrl(action.Url)
	if len(placeholders) == 0 {
		return nil, nil // nothing to generate
	}

	className := fmt.Sprintf("Fetch%vAction", core.ToUpper(action.Name))
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
