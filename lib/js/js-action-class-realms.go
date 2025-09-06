package js

// Combines multiple parts of an EmiAction definition into a single file and generates
// the webrequestX based class for communication

import (
	"fmt"
	"path/filepath"
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

func JsActionManifestRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
) (*jsActionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{}
	actionRealms := jsActionRealms{
		ActionName: action.GetName(),
		HttpMethod: action.MethodUpper(),
	}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	if action.HasRequestHeaders() {
		reqheaderctx := jsHeaderClassContext{
			ClassName: fmt.Sprintf("%vReqHeaders", core.ToUpper(action.GetName())),
		}

		reqheaderctx.Columns = action.GetRequestHeaders()

		requestHeader, err := JsHeaderClass(reqheaderctx, ctx)
		if err != nil {
			return nil, nil, err
		}

		if requestHeader != nil {
			deps = append(deps, requestHeader.CodeChunkDependenies...)
			actionRealms.RequestHeadersClass = requestHeader
		}

	}

	if action.HasResponseHeaders() {

		resheaderctx := jsHeaderClassContext{
			ClassName: fmt.Sprintf("%vResHeaders", core.ToUpper(action.GetName())),
		}

		if action.HasResponseHeaders() {
			resheaderctx.Columns = action.GetResponseHeaders()
		}
		responseHeader, err := JsHeaderClass(resheaderctx, ctx)
		if err != nil {
			return nil, nil, err
		}

		if responseHeader != nil {
			deps = append(deps, responseHeader.CodeChunkDependenies...)
			actionRealms.ResponseHeadersClass = responseHeader
		}
	}

	if isTypeScript {

		// Header is the http headers, extending the Headers class from standard javascript
		pathParameter, err := JsActionPathParams(action)
		if err != nil {
			return nil, nil, err
		}

		if pathParameter != nil {
			deps = append(deps, pathParameter.CodeChunkDependenies...)
			actionRealms.PathParameter = pathParameter
		}
	}

	// Options type, the type which defines how many different things can go
	// into this request.
	optionsctx := jsActionOptionsContext{
		ActionName:  action.GetName(),
		QsClassName: "URLSearchParams",
	}

	if len(action.GetQuery()) > 0 {

		qs, err := JsActionQsClass(action, ctx)
		if err != nil {
			return nil, nil, err
		}

		if qs != nil {
			deps = append(deps, qs.CodeChunkDependenies...)
			actionRealms.QueryStringClass = qs
			token := findTokenByName(qs.Tokens, TOKEN_ROOT_CLASS)
			if token != nil {
				optionsctx.QsClassName = token.Value
			}
		}
	}

	if actionRealms.PathParameter != nil {
		token := findTokenByName(actionRealms.PathParameter.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.ParamsTypeName = token.Value
		}
	}

	if actionRealms.RequestHeadersClass != nil {
		token := findTokenByName(actionRealms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS)
		if token != nil {
			optionsctx.RequestHeadersClassName = token.Value
		}
	}

	if actionRealms.ResponseHeadersClass != nil {
		token := findTokenByName(actionRealms.ResponseHeadersClass.Tokens, TOKEN_ROOT_CLASS)
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
	if action.HasRequestFields() {
		fields, err := JsCommonObjectGenerator(action.GetRequestFields(), ctx, JsCommonObjectContext{
			RootClassName: action.GetName() + "Req",
		})

		if err != nil {
			return nil, nil, err
		}

		deps = append(deps, fields.CodeChunkDependenies...)
		actionRealms.RequestClass = fields
	} else if action.HasRequestDto() {
		actionRealms.RequestClass = castDtoNameToCodeChunk(action.GetRequestDto())
		deps = append(deps, actionRealms.RequestClass.CodeChunkDependenies...)
	}

	// Action response (out)
	if action.HasResponseFields() {
		outClassName := action.GetName() + "Res"
		fields, err := JsCommonObjectGenerator(action.GetResponseFields(), ctx, JsCommonObjectContext{
			RootClassName: outClassName,
		})
		if err != nil {
			return nil, nil, err
		}
		deps = append(deps, fields.CodeChunkDependenies...)
		actionRealms.ResponseClass = fields
	} else if action.HasResponseDto() {
		actionRealms.ResponseClass = castDtoNameToCodeChunk(action.GetResponseDto())
		deps = append(deps, actionRealms.ResponseClass.CodeChunkDependenies...)
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

func parseDtoPath(input string) (path string, className string) {
	if input == "" {
		return "./", ""
	}

	// normalize slashes
	input = filepath.ToSlash(input)

	dir := filepath.Dir(input)
	base := filepath.Base(input)

	if dir == "." {
		path = "./" + base
	} else {
		path = filepath.ToSlash(filepath.Join(dir, base))
	}

	className = base
	return
}

func castDtoNameToCodeChunk(dtoName string) *core.CodeChunkCompiled {
	directory, className := parseDtoPath(dtoName)

	return &core.CodeChunkCompiled{
		ActualScript: []byte(""),
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_OBJ_CLASS,
				Value: className,
			},
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
		CodeChunkDependenies: []core.CodeChunkDependency{
			{
				Objects: []string{className},

				// @todo This doesn't handle external files unfortunately

				Location: directory,
			},
		},
	}
}
