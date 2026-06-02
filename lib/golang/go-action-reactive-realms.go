package golang

import (
	"github.com/torabian/emi/lib/core"
)

type goActionReactiveRealms struct {
	ActionName    string
	PackageName   string
	PathParameter *core.CodeChunkCompiled
	QueryParams   *core.CodeChunkCompiled
	SafeUrl       string
}

func GoActionReactiveRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (goActionReactiveRealms, []core.CodeChunkDependency, error) {

	f := GetCommonFlags(ctx)

	deps := []core.CodeChunkDependency{
		{
			Location: f.Emigo,
		},
		{
			Location: "net/url",
		},
		{
			Location: "net/http",
		},
	}

	realms := goActionReactiveRealms{
		ActionName:  core.ToUpper(core.NormaliseKey(action.GetName())),
		PackageName: f.PackageName,
		SafeUrl:     core.RemoveTypeAnnotations(action.GetUrl()),
	}

	pathParameter, err := GoActionPathParams(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	queryParams, err := GoActionQueryParams(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if queryParams != nil {
		deps = append(deps, queryParams.CodeChunkDependensies...)
		realms.QueryParams = queryParams
	}

	return realms, deps, nil
}
