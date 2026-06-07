package golang

import (
	"github.com/torabian/emi/lib/core"
)

type goActionReactiveRealms struct {
	ActionName       string
	PackageName      string
	PathParameter    *core.CodeChunkCompiled
	PathParameterCli *core.CodeChunkCompiled
	PathParameterGin *core.CodeChunkCompiled
	QueryParams      *core.CodeChunkCompiled
	SafeUrl          string
	EnabledCli       bool
	SplitCli         bool
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
	realms.EnabledCli = !ctx.HasTag(SkipCli)
	realms.SplitCli = ctx.HasTag(SplitCli)

	pathParameter, err := GoActionPathParams(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	if realms.EnabledCli {
		pathParameterCli, err := GoActionPathParamsCli(action, ctx)
		if err != nil {
			return realms, nil, err
		}
		if pathParameterCli != nil {
			realms.PathParameterCli = pathParameterCli
		}
	}

	// Let's say gin is enabled :)
	pathParameterGin, err := GoActionPathParamsGin(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if pathParameterGin != nil {
		realms.PathParameterGin = pathParameterGin
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
