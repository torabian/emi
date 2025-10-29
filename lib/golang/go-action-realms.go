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

func GoActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (*goActionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{}
	realms := goActionRealms{
		ActionName: action.GetName(),
	}

	return &realms, deps, nil
}
