package js

// Combines multiple parts of an EmiAction definition into a single file and generates
// the webrequestX based class for communication

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func JsActionClass(action *core.EmiAction, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	isReact := strings.Contains(ctx.Tags, GEN_REACT_COMPATIBILITY)

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ORIGINAL_NAME,
				Value: action.Name,
			},
		},
	}

	actionRealms, jsDependencies, err := JsActionClassRealms(action, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, jsDependencies...)

	var reactQuery *reactQueryHookRealms
	if isReact {
		reactQueryRealms, reactQuerydeps, err := ReactQueryHooksBasedOnActionRealms(action, ctx, *actionRealms)
		if err != nil {
			return nil, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, reactQuerydeps...)
		reactQuery = reactQueryRealms

		if action.MethodUpper() == "SSE" {
			if useSSEHook, err := ReactUseSSEHook(reactUseSSEOptions{
				ActionName:        fmt.Sprintf("%v", actionRealms.ActionName),
				MetaDataClassName: findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS).Value,
				Fetchctx:          actionRealms.Fetchctx,
			}, ctx); err != nil {
				return nil, err
			} else {
				reactQuery.UseSSE = useSSEHook
				res.CodeChunkDependenies = append(res.CodeChunkDependenies, useSSEHook.CodeChunkDependenies...)
			}
		}

		if action.MethodUpper() == "REACTIVE" {
			if useSSEHook, err := ReactWebsocketQueryHook(reactWebsocketOptions{
				ActionName:        fmt.Sprintf("%v", actionRealms.ActionName),
				MetaDataClassName: findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS).Value,
				Fetchctx:          actionRealms.Fetchctx,
			}, ctx); err != nil {
				return nil, err
			} else {
				reactQuery.UseSSE = useSSEHook
				res.CodeChunkDependenies = append(res.CodeChunkDependenies, useSSEHook.CodeChunkDependenies...)
			}
		}

	}

	const tmpl = `/**
* Action to communicate with the action {{ .action.Name }}
*/

{{ if .realms.OptionsType }}
	{{ b2s .realms.OptionsType.ActualScript }}
{{ end }}

{{ if .reactQuery }}
	{{ if .reactQuery.UseQuery }}
		{{ if .reactQuery.UseQuery.ReactQueryOptions }}
			{{ b2s .reactQuery.UseQuery.ReactQueryOptions.ActualScript }}
		{{ end }}

		{{ if .reactQuery.UseQuery.UseQueryFunction }}
			{{ b2s .reactQuery.UseQuery.UseQueryFunction.ActualScript }}
		{{ end }}
	{{ end }}
	
	{{ if .reactQuery.UseMutation }}
		{{ if .reactQuery.UseMutation.UseMutationOptions }}
			{{ b2s .reactQuery.UseMutation.UseMutationOptions.ActualScript }}
		{{ end }}

		{{ if .reactQuery.UseMutation.UseMutationFunction }}
			{{ b2s .reactQuery.UseMutation.UseMutationFunction.ActualScript }}
		{{ end }}
	{{ end }}

	{{ if .reactQuery.UseSSE }}
		{{ b2s .reactQuery.UseSSE.ActualScript }}
	{{ end }}
{{ end }}


{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
 

{{ if .fetch }}
	{{ .fetch }}
{{ end }}

{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}

{{ if .reqHeaderCompiledClass }}
{{ .reqHeaderCompiledClass }}
{{ end }}

{{ if .resHeaderCompiledClass }}
{{ .resHeaderCompiledClass }}
{{ end }}

{{ if .qsCompiledClass }}
{{ .qsCompiledClass }}
{{ end }}

`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":                 action,
		"reqHeaderCompiledClass": string(actionRealms.RequestHeadersClass.ActualScript),
		"resHeaderCompiledClass": string(actionRealms.ResponseHeadersClass.ActualScript),
		"qsCompiledClass":        string(actionRealms.QueryStringClass.ActualScript),
		"shouldExport":           true,
		"reactQuery":             reactQuery,
		"nestjsDecorator":        nestJsDecorator,
		"fetch":                  string(actionRealms.FetchMetaClass.ActualScript),
		"realms":                 actionRealms,
		"className":              fmt.Sprintf("%vAction", action.Upper()),
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToUpper(action.Name) + "Action"
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	res.Realms = actionRealms

	return res, nil
}
