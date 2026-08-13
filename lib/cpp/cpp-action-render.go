package cpp

import (
	"reflect"

	"github.com/torabian/emi/lib/core"
)

// CppActionRender renders the complete file for one action or remote, in
// whichever dialect is requested - dispatches to the generic or unreal renderer
// (cpp-action-render-generic.go / cpp-action-render-unreal.go), each of which
// further splits on classic vs `method: reactive`.
func CppActionRender(
	action core.EmiRpcAction,
	dialect Dialect,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, deps, err := cppGetActionRealms(action, dialect, ctx, complexes)
	if err != nil {
		return nil, err
	}

	if dialect == DialectUnreal {
		return cppUnrealActionRender(realms, deps)
	}
	return cppGenericActionRender(realms, deps)
}
