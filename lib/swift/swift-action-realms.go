package swift

import (
	"github.com/torabian/emi/lib/core"
)

type actionRealms struct {
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

func GetActionRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (actionRealms, []core.CodeChunkDependency, error) {
	deps := []core.CodeChunkDependency{}

	realms := actionRealms{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
	}

	// Header is the http headers, extending the Headers class from standard javascript
	pathParameter, err := SwiftActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}

	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	if action.HasRequestFields() {
		fields, err := SwiftCommonStructGenerator(action.GetRequestFields(), ctx, commonClassContext{
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
	}

	// Action response (out)
	if action.HasResponseFields() {
		outClassName := realms.ActionName + "Res"
		fields, err := SwiftCommonStructGenerator(action.GetResponseFields(), ctx, commonClassContext{
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
	}

	return realms, deps, nil
}
