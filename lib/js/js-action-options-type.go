// A type generated for type script, which holds all information
// which we can modify for an action

package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type jsActionOptionsContext struct {
	ParamsTypeName           string
	QsClassName              string
	RequestHeadersClassName  string
	ResponseHeadersClassName string
	ActionName               string
}

func TsActionOptionsTypeHelper(optionsctx jsActionOptionsContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	className := fmt.Sprintf("%vActionOptions", core.ToUpper(optionsctx.ActionName))
	const tmpl = `
export type {{ .className }} = {
	queryKey?: unknown[];
	{{ if .optionsctx.ParamsTypeName }}
	params: {{ .optionsctx.ParamsTypeName }};
	{{ end }}
	{{ if .optionsctx.QsClassName }}
	qs?: {{ .optionsctx.QsClassName }};
	{{ end }}
	{{ if .optionsctx.RequestHeadersClassName }}
	headers?: {{ .optionsctx.RequestHeadersClassName }};
	{{ end }}
};
	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"optionsctx": optionsctx,
		"ctx":        ctx,
		"className":  className,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
	}

	return res, nil
}
