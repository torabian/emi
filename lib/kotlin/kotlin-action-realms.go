package kotlin

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
	deps := []core.CodeChunkDependency{
		{
			Location: "okhttp3.*",
		},
		{
			Location: "okhttp3.MediaType.Companion.toMediaType",
		},
		{
			Location: "okhttp3.RequestBody.Companion.toRequestBody",
		},
		{
			Location: "okhttp3.HttpUrl.Companion.toHttpUrl",
		},
		{
			Location: "kotlinx.coroutines.Dispatchers",
		},
		{
			Location: "kotlinx.coroutines.withContext",
		},
		{
			Location: "emikot.ClientContext",
		},
		{
			Location: "kotlinx.serialization.*",
		},
		{
			Location: "kotlinx.serialization.json.*",
		},
	}

	realms := actionRealms{
		ActionName: core.ToUpper(core.NormaliseKey(action.GetName())),
	}

	// Header is the http headers, extending the Headers class from standard javascript
	pathParameter, err := KotlinActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}

	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	if action.HasRequestFields() {
		fields, err := KotlinCommonStructGenerator(action.GetRequestFields(), ctx, commonClassContext{
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
		// For the java/kotlin it's not needed?
		//deps = append(deps, realms.RequestClass.CodeChunkDependensies...)
	}

	// Action response (out)
	if action.HasResponseFields() {
		outClassName := realms.ActionName + "Res"
		fields, err := KotlinCommonStructGenerator(action.GetResponseFields(), ctx, commonClassContext{
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
		// For the java/kotlin it's not needed?
		// deps = append(deps, realms.ResponseClass.CodeChunkDependensies...)
	}

	return realms, deps, nil
}
