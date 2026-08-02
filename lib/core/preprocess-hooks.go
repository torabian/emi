package core

// PreprocessHook expands or mutates a parsed module before any compiler-specific
// codegen runs. Register one instead of adding code directly to Emi.Preprocess (in
// preprocess.go) - that keeps preprocess.go itself generic, and lets a feature (e.g.
// entities' update-dto synthesis in preprocess-entities.go) plug into every backend
// that calls Preprocess/PreprocessForAction, including ones that have no idea the
// feature exists (openapi, postman, md, ...).
type PreprocessHook func(m *Emi) error

var globalPreprocessHooks []PreprocessHook
var actionPreprocessHooks = map[string][]PreprocessHook{}

// RegisterPreprocessHook registers hook to run on every Preprocess() call, regardless
// of which compiler action ends up consuming the module. Use this for functionality
// every backend needs to see - e.g. entities' update dto, since Go, Kotlin, and OpenAPI
// all need that dto to exist, none of them aware entities are what produced it.
//
// Typically called once from an init() in the package that owns the feature.
func RegisterPreprocessHook(hook PreprocessHook) {
	if hook == nil {
		return
	}
	globalPreprocessHooks = append(globalPreprocessHooks, hook)
}

// RegisterPreprocessHookForAction registers hook to run only when preprocessing for
// the named compiler action (e.g. "go", "kotlin", "openapi" - the same names as
// core.BaseAction.Name), via PreprocessForAction. Use this when an expansion only
// makes sense for one backend, instead of every module.Preprocess() caller.
func RegisterPreprocessHookForAction(action string, hook PreprocessHook) {
	if hook == nil || action == "" {
		return
	}
	actionPreprocessHooks[action] = append(actionPreprocessHooks[action], hook)
}

func runPreprocessHooks(m *Emi, hooks []PreprocessHook) error {
	for _, hook := range hooks {
		if hook == nil {
			continue
		}
		if err := hook(m); err != nil {
			return err
		}
	}
	return nil
}
