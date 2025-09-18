package js

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

type reactQueryUseMutationHookRealms struct {
	UseMutationFunction *core.CodeChunkCompiled
	UseMutationOptions  *core.CodeChunkCompiled
}

// useMutation realms generator (for GET/POST/PUT/DELETE etc.)
func ReactQueryUseMutationRealms(
	ctx core.MicroGenContext,
	actionRealms jsActionRealms,
) (*reactQueryUseMutationHookRealms, []core.CodeChunkDependency, error) {

	realms := reactQueryUseMutationHookRealms{}
	deps := []core.CodeChunkDependency{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	creatorFn := findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_CREATOR_FN)

	if isTypeScript {
		rmoptions := reactMutationOptionsType{
			ActionName:                actionRealms.ActionName,
			ActionMutationOptionsName: findTokenByName(actionRealms.OptionsType.Tokens, TOKEN_ROOT_CLASS).Value,
		}

		if creatorFn != nil {
			if strings.Contains(creatorFn.Value, " = ") {
				rmoptions.CreatorFnType = creatorFn.Value[0:strings.Index(creatorFn.Value, " = ")]
			}
		}

		if actionRealms.PathParameter != nil {
			rmoptions.HasPathParameters = true
		}

		reactMutationOptions, err := ReactMutationOptionsTypeFunction(rmoptions, ctx)
		if err != nil {
			return nil, nil, err
		}
		deps = append(deps, reactMutationOptions.CodeChunkDependensies...)
		realms.UseMutationOptions = reactMutationOptions
	}

	useMutationOptions := reactUseMutationOptions{
		ActionName:        actionRealms.ActionName,
		MetaDataClassName: findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS).Value,
		RequestClass:      "unknown",
	}

	if actionRealms.RequestClass != nil {
		useMutationOptions.RequestClass = findTokenByName(actionRealms.RequestClass.Tokens, TOKEN_ROOT_CLASS).Value
	}

	if creatorFn != nil {
		useMutationOptions.HasCreatorFunction = true
	}

	if actionRealms.PathParameter != nil {
		useMutationOptions.HasPathParameters = true
	}

	if realms.UseMutationOptions != nil {
		useMutationOptions.ActionMutationOptionsName = findTokenByName(realms.UseMutationOptions.Tokens, TOKEN_ROOT_CLASS).Value
	}

	useMutationFunction, err := ReactUseMutationFunction(useMutationOptions, ctx)
	if err != nil {
		return nil, nil, err
	}
	deps = append(deps, useMutationFunction.CodeChunkDependensies...)
	realms.UseMutationFunction = useMutationFunction

	return &realms, deps, nil
}
