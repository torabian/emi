// file: react-query-use-options.go
// A type generated for type script, which holds all information
// which we can modify for an action

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactQueryOptionsType struct {
	ActionQueryOptionsName string
	ActionName             string
	HasPathParameters      bool
	JsActionRealms         jsActionRealms
	CreatorFnType          string
}

func ReactQueryOptionsTypeFunction(rqoptions reactQueryOptionsType, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	className := fmt.Sprintf("%vQueryOptions", core.ToUpper(rqoptions.ActionName))

	var responseClass *core.GeneratedScriptToken
	if rqoptions.JsActionRealms.ResponseClass != nil {
		responseClass = findTokenByName(rqoptions.JsActionRealms.ResponseClass.Tokens, TOKEN_ROOT_CLASS)
	}

	const tmpl = `
export type {{ .className }} = Omit<
	UseQueryOptions<
		unknown,
		unknown,
		{{ if .responseClass }}
			{{ .responseClass.Value }},
		{{ end }}
		unknown[]
	>,
	"queryKey"
> &
	{{ .rqoptions.ActionQueryOptionsName }}

{{ if .rqoptions.CreatorFnType }}
& Partial<{
	{{ .rqoptions.CreatorFnType }}
}>
{{ end }}
& {
	onMessage?: (ev: MessageEvent) => void;
	overrideUrl?: string;
	headers?: Headers;
	ctx?: FetchxContext;
}

	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"rqoptions":     rqoptions,
		"ctx":           ctx,
		"responseClass": responseClass,
		"jsRealms":      rqoptions.JsActionRealms,
		"className":     className,
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

	if isTypeScript {
		res.CodeChunkDependensies = []core.CodeChunkDependency{
			{
				Objects:  []string{"type UseQueryOptions"},
				Location: "@tanstack/react-query",
			},
		}
	}

	return res, nil
}
