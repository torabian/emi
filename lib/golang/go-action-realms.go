package golang

import (
	"github.com/torabian/emi/lib/core"
)

type goActionRealms struct {
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
}

var GENERATE_GO_CLINET = true

func GoActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (goActionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{
		{
			Location: "net/http",
		},
		{
			Location: "github.com/gin-gonic/gin",
		},
	}

	if GENERATE_GO_CLINET {
		deps = append(
			deps,
			core.CodeChunkDependency{Location: "bytes"},
			core.CodeChunkDependency{Location: "encoding/json"},
			core.CodeChunkDependency{Location: "fmt"},
			core.CodeChunkDependency{Location: "io"},
			core.CodeChunkDependency{Location: "net/url"},
		)
	}

	realms := goActionRealms{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
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
		deps = append(deps, realms.RequestClass.CodeChunkDependensies...)
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
		deps = append(deps, realms.ResponseClass.CodeChunkDependensies...)
	}

	return realms, deps, nil
}
