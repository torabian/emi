package kotlin

import (
	"reflect"

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

	// RequestTypeName/ResponseTypeName are the bare Kotlin class names backing
	// RequestClass/ResponseClass (read off their TOKEN_OBJ_CLASS token - see
	// castDtoNameToCodeChunk/KotlinCommonStructGenerator), "" when the action has no
	// in/out shape. EnvelopeClass is action.GetResponseEnvelopeClass() (e.g.
	// "GResponse", see lib/core/preprocess-entity-actions.go) - both are resolved once
	// here so kotlin-action-render.go and kotlin-action-reactive-render.go don't each
	// re-derive them.
	RequestTypeName  string
	ResponseTypeName string
	EnvelopeClass    string
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

	realms := actionRealms{}

	if action == nil || reflect.ValueOf(action).IsNil() {
		return realms, []core.CodeChunkDependency{}, nil
	}

	realms.ActionName = core.ToUpper(core.NormaliseKey(action.GetName()))

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

	if realms.RequestClass != nil {
		if token := core.FindTokenByName(realms.RequestClass.Tokens, TOKEN_OBJ_CLASS); token != nil {
			realms.RequestTypeName = token.Value
		}
	}
	if realms.ResponseClass != nil {
		if token := core.FindTokenByName(realms.ResponseClass.Tokens, TOKEN_OBJ_CLASS); token != nil {
			realms.ResponseTypeName = token.Value
		}
	}

	realms.EnvelopeClass = action.GetResponseEnvelopeClass()
	if realms.EnvelopeClass != "" {
		// Every envelope class this compiler ships (currently just GResponse - see
		// lib/kotlin/kotlin-include/gresponse.kt) lives in the same emikot runtime
		// package as MaybeField/Maybe/EmiWebSocketX.
		deps = append(deps, core.CodeChunkDependency{Location: "emikot." + realms.EnvelopeClass})
	}

	return realms, deps, nil
}
