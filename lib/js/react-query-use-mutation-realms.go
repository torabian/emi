package js

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

type reactQueryUseMutationHookRealms struct {
	UseMutationFunction *core.CodeChunkCompiled
	UseMutationOptions  *core.CodeChunkCompiled
}

// useMutation realms generator (for POST/PUT/DELETE etc.)
func ReactQueryUseMutationRealms(
	ctx core.MicroGenContext,
	actionRealms jsActionRealms,
) (*reactQueryUseMutationHookRealms, []core.CodeChunkDependency, error) {

	realms := reactQueryUseMutationHookRealms{}
	deps := []core.CodeChunkDependency{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	if isTypeScript {
		rmoptions := reactMutationOptionsType{
			ActionName:                actionRealms.ActionName,
			ActionMutationOptionsName: findTokenByName(actionRealms.OptionsType.Tokens, TOKEN_ROOT_CLASS).Value,
		}

		if actionRealms.PathParameter != nil {
			rmoptions.HasPathParameters = true
		}

		reactMutationOptions, err := ReactMutationOptionsTypeFunction(rmoptions, ctx)
		if err != nil {
			return nil, nil, err
		}
		deps = append(deps, reactMutationOptions.CodeChunkDependenies...)
		realms.UseMutationOptions = reactMutationOptions
	}

	useMutationOptions := reactUseMutationOptions{
		ActionName:        actionRealms.ActionName,
		MetaDataClassName: findTokenByName(actionRealms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS).Value,
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
	deps = append(deps, useMutationFunction.CodeChunkDependenies...)
	realms.UseMutationFunction = useMutationFunction

	return &realms, deps, nil
}
