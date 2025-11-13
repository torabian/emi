package golang

import (
	"encoding/json"

	"github.com/torabian/emi/lib/core"
)

type goActionRealms struct {
	ActionName           string
	HttpMethod           string
	PathParameter        *core.CodeChunkCompiled
	FetchMetaClass       *core.CodeChunkCompiled
	RequestClass         *core.CodeChunkCompiled
	ResponseClass        *core.CodeChunkCompiled
	QueryStringClass     *core.CodeChunkCompiled
	RequestHeadersClass  *core.CodeChunkCompiled
	ResponseHeadersClass *core.CodeChunkCompiled
	RequestClassName     string
}

var GENERATE_GO_CLINET = true

func GoActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (goActionRealms, []core.CodeChunkDependency, error) {

	type Flags struct {
		Emigo string `json:"emigo"`
	}
	var f Flags
	if err := json.Unmarshal([]byte(ctx.Flags), &f); err != nil {
		panic(err)
	}

	deps := []core.CodeChunkDependency{
		{
			Location: "net/http",
		},
		{
			Location: "github.com/gin-gonic/gin",
		},
		{
			Location: f.Emigo,
		},
	}

	if GENERATE_GO_CLINET {
		deps = append(
			deps,
			core.CodeChunkDependency{Location: "encoding/json"},
			core.CodeChunkDependency{Location: "fmt"},
			core.CodeChunkDependency{Location: "io"},
			core.CodeChunkDependency{Location: "net/url"},
		)

		if action.MethodUpper() != "GET" {
			deps = append(deps, core.CodeChunkDependency{Location: "bytes"})
		}
	}

	realms := goActionRealms{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
	}

	// Header is the http headers, extending the Headers class from standard javascript
	pathParameter, err := GoActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}

	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	if action.HasRequestFields() {
		fields, err := GoCommonStructGenerator(action.GetRequestFields(), ctx, GoCommonStructContext{
			RootClassName:       realms.ActionName + "Req",
			RecognizedComplexes: complexes,
		})

		if err != nil {
			return realms, nil, err
		}

		deps = append(deps, fields.CodeChunkDependensies...)
		realms.RequestClass = fields
	} else if action.HasRequestDto() {
		realms.RequestClass = castDtoNameToCodeChunk(action.GetRequestDto())
		// Not sure if this is needed in golang
		// deps = append(deps, realms.RequestClass.CodeChunkDependensies...)
	}

	// Action response (out)
	if action.HasResponseFields() {
		outClassName := realms.ActionName + "Res"
		fields, err := GoCommonStructGenerator(action.GetResponseFields(), ctx, GoCommonStructContext{
			RootClassName:       outClassName,
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return realms, nil, err
		}
		deps = append(deps, fields.CodeChunkDependensies...)
		realms.ResponseClass = fields
	} else if action.HasResponseDto() {
		realms.ResponseClass = castDtoNameToCodeChunk(action.GetResponseDto())
		// Not sure if this is needed in golang
		// deps = append(deps, realms.ResponseClass.CodeChunkDependensies...)
	}

	if realms.RequestClass != nil {
		if token := core.FindTokenByName(realms.RequestClass.Tokens, TOKEN_ROOT_CLASS); token != nil {
			realms.RequestClassName = token.Value
		}
	}

	return realms, deps, nil
}
