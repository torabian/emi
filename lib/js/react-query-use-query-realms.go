// file: react-query-use-query-realms.go
package js

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

type reactQueryUseQueryHookRealms struct {
	UseMutationFunctions *core.CodeChunkCompiled
	UseQueryFunction     *core.CodeChunkCompiled
	ReactQueryOptions    *core.CodeChunkCompiled
}

// useQuery is a feature from react-query, this function generates necessary realms
// to render that function and types necessary, use this only if method is get.
func ReactQueryUseQueryRealms(
	ctx core.MicroGenContext,
	actionRealms jsActionRealms,
) (*reactQueryUseQueryHookRealms, []core.CodeChunkDependency, error) {
	creatorFn := findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_CREATOR_FN)

	realms := reactQueryUseQueryHookRealms{}
	deps := []core.CodeChunkDependency{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	if isTypeScript {
		// React Query options
		rqoptions := reactQueryOptionsType{
			ActionName:             actionRealms.ActionName,
			ActionQueryOptionsName: findTokenByName(actionRealms.OptionsType.Tokens, TOKEN_ROOT_CLASS).Value,
			JsActionRealms:         actionRealms,
		}

		if creatorFn != nil {
			if strings.Contains(creatorFn.Value, " = ") {
				rqoptions.CreatorFnType = creatorFn.Value[0:strings.Index(creatorFn.Value, " = ")]
			}
		}

		if actionRealms.PathParameter != nil {
			rqoptions.HasPathParameters = true
		}

		reactQueryOptions, err := ReactQueryOptionsTypeFunction(rqoptions, ctx)
		if err != nil {
			return nil, nil, err
		}
		deps = append(deps, reactQueryOptions.CodeChunkDependenies...)
		realms.ReactQueryOptions = reactQueryOptions
	}

	// React Query use-query
	useQueryOptions := reactUseQueryOptions{
		ActionName:         actionRealms.ActionName,
		NewUrlFunctionName: findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_NEW_URL_FN).Value,
		MetaDataClassName:  findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS).Value,
	}

	if actionRealms.PathParameter != nil {
		useQueryOptions.HasPathParameters = true
	}

	if realms.ReactQueryOptions != nil {
		useQueryOptions.ActionQueryOptionsName = findTokenByName(realms.ReactQueryOptions.Tokens, TOKEN_ROOT_CLASS).Value
	}

	useQueryFunction, err := ReactUseQueryOptionsFunction(useQueryOptions, ctx)
	if err != nil {
		return nil, nil, err
	}
	deps = append(deps, useQueryFunction.CodeChunkDependenies...)
	realms.UseQueryFunction = useQueryFunction

	return &realms, deps, nil
}
