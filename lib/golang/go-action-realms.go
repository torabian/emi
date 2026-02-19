package golang

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

type goActionRealms struct {
	ActionName           string
	SafeUrl              string
	HttpMethod           string
	PathParameter        *core.CodeChunkCompiled
	FetchMetaClass       *core.CodeChunkCompiled
	RequestClass         *core.CodeChunkCompiled
	ResponseClass        *core.CodeChunkCompiled
	CliName              string
	QueryStringClass     *core.CodeChunkCompiled
	QueryParams          *core.CodeChunkCompiled
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
			Location: "net/http",
		},
		{
			Location: "bytes",
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
		CliName:    ActionToCliName(action),
		SafeUrl:    core.RemoveTypeAnnotations(action.GetUrl()),
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

func ActionToCliName(x core.EmiRpcAction) string {
	if x.GetCliName() != "" {
		return x.GetCliName()
	}

	return strings.ReplaceAll(core.ToSnakeCase((x.GetName())), "_", "-")
}
