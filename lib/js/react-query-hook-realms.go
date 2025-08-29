package js

import (
	"github.com/torabian/emi/lib/core"
)

type reactQueryHookRealms struct {

	// In case of get methods, or queries, you can use this.
	UseQuery *reactQueryUseQueryHookRealms

	UseMutation *reactQueryUseMutationHookRealms
}

func ReactQueryHooksBasedOnActionRealms(
	action *core.Module3Action,
	ctx core.MicroGenContext,
	actionRealms jsActionRealms,
) (*reactQueryHookRealms, []core.CodeChunkDependency, error) {

	realms := reactQueryHookRealms{}
	deps := []core.CodeChunkDependency{}

	if action.MethodUpper() == "GET" {
		if useQueryRealms, useQueryDeps, err := ReactQueryUseQueryRealms(ctx, actionRealms); err != nil {
			return nil, nil, err
		} else {
			realms.UseQuery = useQueryRealms
			deps = append(deps, useQueryDeps...)
		}
	}

	if action.MethodUpper() == "POST" {
		if useMutationRealms, useMutationDeps, err := ReactQueryUseMutationRealms(ctx, actionRealms); err != nil {
			return nil, nil, err
		} else {
			realms.UseMutation = useMutationRealms
			deps = append(deps, useMutationDeps...)
		}
	}

	return &realms, deps, nil
}
