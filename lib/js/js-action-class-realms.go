package js

// Combines multiple parts of an EmiAction definition into a single file and generates
// the webrequestX based class for communication

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type jsActionRealms struct {
	ActionName           string
	HttpMethod           string
	UseQueryFunction     *core.CodeChunkCompiled
	ReactQueryOptions    *core.CodeChunkCompiled
	PathParameter        *core.CodeChunkCompiled
	OptionsType          *core.CodeChunkCompiled
	FetchMetaClass       *core.CodeChunkCompiled
	RequestClass         *core.CodeChunkCompiled
	ResponseClass        *core.CodeChunkCompiled
	QueryStringClass     *core.CodeChunkCompiled
	RequestHeadersClass  *core.CodeChunkCompiled
	ResponseHeadersClass *core.CodeChunkCompiled
	Fetchctx             fetchStaticFunctionContext
}

func JsActionClassRealms(
	action *core.EmiAction,
	ctx core.MicroGenContext,
) (*jsActionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{}
	actionRealms := jsActionRealms{
		ActionName: action.Name,
		HttpMethod: action.MethodUpper(),
	}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	reqheaderctx := jsHeaderClassContext{
		ClassName: fmt.Sprintf("%vReqHeaders", core.ToUpper(action.Name)),
	}

	if action.In != nil && len(action.In.Headers) > 0 {
		reqheaderctx.Columns = action.In.Headers
	}
	requestHeader, err := JsHeaderClass(reqheaderctx, ctx)
	if err != nil {
		return nil, nil, err
	}

	if requestHeader != nil {
		deps = append(deps, requestHeader.CodeChunkDependenies...)
		actionRealms.RequestHeadersClass = requestHeader
	}

	resheaderctx := jsHeaderClassContext{
		ClassName: fmt.Sprintf("%vResHeaders", core.ToUpper(action.Name)),
	}

	if action.Out != nil && len(action.Out.Headers) > 0 {
		resheaderctx.Columns = action.Out.Headers
	}
	responseHeader, err := JsHeaderClass(resheaderctx, ctx)
	if err != nil {
		return nil, nil, err
	}

	if responseHeader != nil {
		deps = append(deps, responseHeader.CodeChunkDependenies...)
		actionRealms.ResponseHeadersClass = responseHeader
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

	if requestHeader != nil {
		token := findTokenByName(requestHeader.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.RequestHeadersClassName = token.Value
		}
	}

	if responseHeader != nil {
		token := findTokenByName(responseHeader.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.ResponseHeadersClassName = token.Value
		}
	}

	if isTypeScript {
		optionsType, err := TsActionOptionsTypeHelper(optionsctx, ctx)
		if err != nil {
			return nil, nil, err
		}
		deps = append(deps, optionsType.CodeChunkDependenies...)
		actionRealms.OptionsType = optionsType

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

	}

	fetch, fetchctx, err := JsActionFetchAndMetaData(action, actionRealms, ctx)
	if err != nil {
		return nil, nil, err
	}
	deps = append(deps, fetch.CodeChunkDependenies...)
	actionRealms.FetchMetaClass = fetch
	actionRealms.Fetchctx = fetchctx

	return &actionRealms, deps, nil
}
