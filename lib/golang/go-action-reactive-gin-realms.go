package golang

import (
	"github.com/torabian/emi/lib/core"
)

type goActionReactiveGinRealms struct {
	ActionName       string
	PackageName      string
	PathParameter    *core.CodeChunkCompiled
	PathParameterGin *core.CodeChunkCompiled
	QueryParams      *core.CodeChunkCompiled
	SafeUrl          string
	SkipGinWasm      bool
}

var DEFAULT_GO_PACKAGE = "external"

func GoActionReactiveGinRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (goActionReactiveGinRealms, []core.CodeChunkDependency, error) {

	f := GetCommonFlags(ctx)

	deps := []core.CodeChunkDependency{
		{
			Location: "github.com/gin-gonic/gin",
		},
		{
			Location: "net/http",
		},
		{
			Location: "net/url",
		},
	}

	realms := goActionReactiveGinRealms{
		ActionName:  core.ToUpper(core.NormaliseKey(action.GetName())),
		PackageName: f.PackageName,
		SafeUrl:     core.RemoveTypeAnnotations(action.GetUrl()),
	}

	pathParameterGin, err := GoActionPathParamsGin(action, ctx)
	if err != nil {
		return realms, nil, err
	}
	if pathParameterGin != nil {
		realms.PathParameterGin = pathParameterGin
	}

	realms.SkipGinWasm = ctx.HasTag(SkipWasmGin)

	return realms, deps, nil
}
