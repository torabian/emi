package js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type jsActionRealms struct {
	UseQueryFunction    *core.CodeChunkCompiled
	ReactQueryOptions   *core.CodeChunkCompiled
	PathParameter       *core.CodeChunkCompiled
	OptionsType         *core.CodeChunkCompiled
	FetchMetaClass      *core.CodeChunkCompiled
	RequestClass        *core.CodeChunkCompiled
	ResponseClass       *core.CodeChunkCompiled
	QueryStringClass    *core.CodeChunkCompiled
	RequestHeadersClass *core.CodeChunkCompiled
}

func JsActionClassRealms(
	action *core.Module3Action,
	ctx core.MicroGenContext,
) (*jsActionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{}
	actionRealms := jsActionRealms{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	isReact := strings.Contains(ctx.Tags, GEN_REACT_COMPATIBILITY)

	// Header is the http headers, extending the Headers class from standard javascript
	header, err := JsActionHeaderClass(action, ctx)
	if err != nil {
		return nil, nil, err
	}

	if header != nil {
		deps = append(deps, header.CodeChunkDependenies...)
		actionRealms.RequestHeadersClass = header
	}

	// Header is the http headers, extending the Headers class from standard javascript
	pathParameter, err := JsActionPathParams(action)
	if err != nil {
		return nil, nil, err
	}

	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependenies...)
		actionRealms.PathParameter = pathParameter
	}

	// Query strings for the request builder
	qs, err := JsActionQsClass(action, ctx)
	if err != nil {
		return nil, nil, err
	}

	if qs != nil {
		deps = append(deps, qs.CodeChunkDependenies...)
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
			return nil, nil, err
		}
		deps = append(deps, optionsType.CodeChunkDependenies...)
		actionRealms.OptionsType = optionsType

		if isReact {

			// React Query options
			rqoptions := reactQueryOptionsType{
				ActionName:             action.Name,
				ActionQueryOptionsName: findTokenByName(optionsType.Tokens, TOKEN_ROOT_CLASS).Value,
			}
			reactQueryOptions, err := ReactQueryOptionsTypeFunction(rqoptions, ctx)
			if err != nil {
				return nil, nil, err
			}
			deps = append(deps, reactQueryOptions.CodeChunkDependenies...)
			actionRealms.ReactQueryOptions = reactQueryOptions
		}
	}

	// Action request (in)
	if action.In != nil && len(action.In.Fields) > 0 {
		fields, err := JsCommonObjectGenerator(action.In.Fields, ctx, JsCommonObjectContext{
			RootClassName: action.Name + "Req",
		})

		if err != nil {
			return nil, nil, err
		}

		deps = append(deps, fields.CodeChunkDependenies...)
		actionRealms.RequestClass = fields
	}

	// Action response (out)
	if action.Out != nil && len(action.Out.Fields) > 0 {

		outClassName := action.Name + "Res"
		fields, err := JsCommonObjectGenerator(action.Out.Fields, ctx, JsCommonObjectContext{
			RootClassName: outClassName,
		})

		if err != nil {
			return nil, nil, err
		}

		deps = append(deps, fields.CodeChunkDependenies...)
		actionRealms.ResponseClass = fields

		mitem := fmt.Sprintf(`
export abstract class %vFactory {
	abstract create(data: unknown): %v;
}
		`, core.ToUpper(outClassName), core.ToUpper(outClassName))
		fields.ActualScript = append(fields.ActualScript, []byte(mitem)...)
	}

	fetch, err := JsActionFetchAndMetaData(action, actionRealms, ctx)
	if err != nil {
		return nil, nil, err
	}
	deps = append(deps, fetch.CodeChunkDependenies...)
	actionRealms.FetchMetaClass = fetch

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
			return nil, nil, err
		}
		deps = append(deps, useQueryFunction.CodeChunkDependenies...)
		actionRealms.UseQueryFunction = useQueryFunction
	}

	return &actionRealms, deps, nil
}
