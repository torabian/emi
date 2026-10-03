package js

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// JsParamsDtoAliasesGenerate declares the params type of every event/permission that
// uses `params.dto` (see core.Emi.ParamsDtoAliases): the referenced dto is imported
// and re-exported under the params class name (`export { PostDto as
// PostPublishPermissionParams }`), which works for both the class and, in TypeScript,
// its type. Returns (nil, nil) when no params reference a dto.
func JsParamsDtoAliasesGenerate(module *core.Emi, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	aliases := module.ParamsDtoAliases()
	if len(aliases) == 0 {
		return nil, nil
	}

	var script strings.Builder
	var deps []core.CodeChunkDependency
	imported := map[string]bool{}

	for _, alias := range aliases {
		if strings.Contains(alias.Dto, "|") {
			return nil, fmt.Errorf("params of %s: dto %q must be a single dto", alias.Class, alias.Dto)
		}
		directory, className := parseDtoPath(strings.TrimSpace(alias.Dto))

		// Imported under a private name so it can never clash with a generated class.
		local := "_" + className
		if !imported[className] {
			imported[className] = true
			deps = append(deps, core.CodeChunkDependency{
				Objects:  []string{className + " as " + local},
				Location: directory,
			})
		}
		fmt.Fprintf(&script, "export { %s as %s };\n", local, alias.Class)
	}

	res := &core.CodeChunkCompiled{
		SuggestedFileName:     "ParamsAliases",
		SuggestedExtension:    ".js",
		ActualScript:          []byte(script.String()),
		CodeChunkDependensies: deps,
	}
	if ctx.HasTag(Typescript) {
		res.SuggestedExtension = ".ts"
	}
	return res, nil
}
