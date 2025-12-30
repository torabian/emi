package golang

import (
	"encoding/json"
	"fmt"

	"github.com/torabian/emi/lib/core"
)

type goActionReactiveRealms struct {
	ActionName    string
	PackageName   string
	PathParameter *core.CodeChunkCompiled
	QueryParams   *core.CodeChunkCompiled
	SafeUrl       string
}

var DEFAULT_GO_PACKAGE = "external"

func GoActionReactiveRealms(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,

) (goActionReactiveRealms, []core.CodeChunkDependency, error) {

	type Flags struct {
		Emigo       string `json:"emigo,omitempty"`
		PackageName string `json:"pkg,omitempty"`
	}
	var f Flags = Flags{
		Emigo:       "github.com/torabian/emi/emigo",
		PackageName: DEFAULT_GO_PACKAGE,
	}
	if err := json.Unmarshal([]byte(ctx.Flags), &f); err != nil {
		fmt.Println("Flags provided are not correct:", ctx.Flags)
	}
	deps := []core.CodeChunkDependency{
		{
			Location: "github.com/gin-gonic/gin",
		},
		{
			Location: "fmt",
		},
		{
			Location: "net/http",
		},
		{
			Location: f.Emigo,
		},
	}

	if err := json.Unmarshal([]byte(ctx.Flags), &f); err != nil {
		fmt.Println("Flags provided are not correct:", ctx.Flags)
	}

	realms := goActionReactiveRealms{
		ActionName:  core.ToUpper(core.NormaliseKey(action.GetName())),
		PackageName: f.PackageName,
		SafeUrl:     core.RemoveTypeAnnotations(action.GetUrl()),
	}

	pathParameter, err := GoActionPathParams(action)
	if err != nil {
		return realms, nil, err
	}
	if pathParameter != nil {
		deps = append(deps, pathParameter.CodeChunkDependensies...)
		realms.PathParameter = pathParameter
	}

	queryParams, err := GoActionQueryParams(action)
	if err != nil {
		return realms, nil, err
	}
	if queryParams != nil {
		deps = append(deps, queryParams.CodeChunkDependensies...)
		realms.QueryParams = queryParams
	}

	return realms, deps, nil
}
