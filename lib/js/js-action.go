package js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type jsActionRealms struct {
	UseQueryFunction    *core.CodeChunkCompiled
	ReactQueryOptions   *core.CodeChunkCompiled
	PathParameter       *core.CodeChunkCompiled
	OptionsType         *core.CodeChunkCompiled
	RequestClass        *core.CodeChunkCompiled
	ResponseClass       *core.CodeChunkCompiled
	QueryStringClass    *core.CodeChunkCompiled
	RequestHeadersClass *core.CodeChunkCompiled
}

func JsActionClass(action *core.Module3Action, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	actionRealms := jsActionRealms{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	isReact := strings.Contains(ctx.Tags, GEN_REACT_COMPATIBILITY)

	res := &core.CodeChunkCompiled{}

	// Header is the http headers, extending the Headers class from standard javascript
	header, err := JsActionHeaderClass(action, ctx)
	if err != nil {
		return nil, err
	}

	if header != nil {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, header.CodeChunkDependenies...)
		actionRealms.RequestHeadersClass = header
	}

	// Header is the http headers, extending the Headers class from standard javascript
	pathParameter, err := JsActionPathParams(action)
	if err != nil {
		return nil, err
	}

	if pathParameter != nil {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, pathParameter.CodeChunkDependenies...)
		actionRealms.PathParameter = pathParameter
	}

	// Query strings for the request builder
	qs, err := JsActionQsClass(action, ctx)
	if err != nil {
		return nil, err
	}

	if qs != nil {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, qs.CodeChunkDependenies...)
		actionRealms.QueryStringClass = qs
	}

	// Options type, the type which defines how many different things can go
	// into this request.
	optionsctx := jsActionOptionsContext{
		ActionName: action.Name,
	}
	if qs != nil {
		token := findTokenByName(qs.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.QsClassName = token.Value
		}
	}

	if pathParameter != nil {
		token := findTokenByName(pathParameter.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.ParamsTypeName = token.Value
		}
	}

	if header != nil {
		token := findTokenByName(header.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.RequestHeadersClassName = token.Value
		}
	}

	if isTypeScript {
		optionsType, err := TsActionOptionsTypeHelper(optionsctx, ctx)
		if err != nil {
			return nil, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, optionsType.CodeChunkDependenies...)
		actionRealms.OptionsType = optionsType

		if isReact {

			// React Query options
			rqoptions := reactQueryOptionsType{
				ActionName:             action.Name,
				ActionQueryOptionsName: findTokenByName(optionsType.Tokens, TOKEN_ROOT_CLASS).Value,
			}
			reactQueryOptions, err := ReactQueryOptionsTypeFunction(rqoptions, ctx)
			if err != nil {
				return nil, err
			}
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, reactQueryOptions.CodeChunkDependenies...)
			actionRealms.ReactQueryOptions = reactQueryOptions
		}
	}

	// Action request (in)
	if action.In != nil && len(action.In.Fields) > 0 {
		fields, err := JsCommonObjectGenerator(action.In.Fields, ctx, JsCommonObjectContext{
			RootClassName: action.Name + "Req",
		})

		if err != nil {
			return nil, err
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fields.CodeChunkDependenies...)
		actionRealms.RequestClass = fields
	}

	// Action response (out)
	if action.Out != nil && len(action.Out.Fields) > 0 {
		fields, err := JsCommonObjectGenerator(action.Out.Fields, ctx, JsCommonObjectContext{
			RootClassName: action.Name + "Res",
		})

		if err != nil {
			return nil, err
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fields.CodeChunkDependenies...)
		actionRealms.ResponseClass = fields
	}

	fetch, err := JsActionFetchAndMetaData(action, actionRealms, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, fetch.CodeChunkDependenies...)

	if isReact {
		// React Query use-query
		useQueryOptions := reactUseQueryOptions{
			ActionName:         action.Name,
			NewUrlFunctionName: findTokenByName(fetch.Tokens, TOKEN_NEW_URL_FN).Value,
			MetaDataClassName:  findTokenByName(fetch.Tokens, TOKEN_ROOT_CLASS).Value,
		}

		if actionRealms.ReactQueryOptions != nil {
			useQueryOptions.ActionQueryOptionsName = findTokenByName(actionRealms.ReactQueryOptions.Tokens, TOKEN_ROOT_CLASS).Value
		}

		useQueryFunction, err := ReactUseQueryOptionsFunction(useQueryOptions, ctx)
		if err != nil {
			return nil, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, useQueryFunction.CodeChunkDependenies...)
		actionRealms.UseQueryFunction = useQueryFunction
	}
	const tmpl = `/**
* Action to communicate with the action {{ .action.Name }}
*/

{{ if .realms.OptionsType }}
	{{ b2s .realms.OptionsType.ActualScript }}
{{ end }}

{{ if .realms.ReactQueryOptions }}
	{{ b2s .realms.ReactQueryOptions.ActualScript }}
{{ end }}

{{ if .realms.UseQueryFunction }}
	{{ b2s .realms.UseQueryFunction.ActualScript }}
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

{{ if .headerCompiledClass }}
{{ .headerCompiledClass }}
{{ end }}

{{ if .qsCompiledClass }}
{{ .qsCompiledClass }}
{{ end }}
 
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":              action,
		"headerCompiledClass": string(header.ActualScript),
		"qsCompiledClass":     string(qs.ActualScript),
		"shouldExport":        true,
		"nestjsDecorator":     nestJsDecorator,
		"fetch":               string(fetch.ActualScript),
		"realms":              actionRealms,
		"className":           fmt.Sprintf("%vAction", action.Upper()),
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToUpper(action.Name) + "Action"
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	return res, nil
}
