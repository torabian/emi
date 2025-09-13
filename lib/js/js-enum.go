package js

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func JsStandaloneEnum(
	enum core.EmiEnum,
	ctx core.MicroGenContext,
) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ORIGINAL_NAME,
				Value: enum.GetName(),
			},
		},
	}

	// Enum is easy, I do not tokenize it. Until I realises I had to do it :)))

	const tmpl = `/**
* Enum {{ .enum.GetName }}
*/


{{ if .isTypeScript }}
export enum {{.enum.GetName}} {
	{{ range .enum.Fields }} 
	 	/*
		*	{{ .Description }}
		**/
		{{ .GetKey }} = '{{.Key}}',
	{{ end }}
}

{{ else }}
export const {{.enum.GetName}} = Object.freeze({
	{{ range .enum.Fields }} 
	 	/*
		*	{{ .Description }}
		**/
		{{ .GetKey }}: '{{.Key}}',
	{{ end }}
});


{{ end }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"enum":         enum,
		"isTypeScript": isTypeScript,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = enum.GetName()
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	return res, nil
}
