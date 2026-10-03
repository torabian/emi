package kotlin

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// KotlinParamsDtoAliasesGenerate declares the params type of every
// event/permission that uses `params.dto` (see core.Emi.ParamsDtoAliases) as a
// typealias of the referenced dto. Returns (nil, nil) when no params reference a dto.
func KotlinParamsDtoAliasesGenerate(module *core.Emi) (*core.CodeChunkCompiled, error) {
	aliases := module.ParamsDtoAliases()
	if len(aliases) == 0 {
		return nil, nil
	}

	var script strings.Builder
	for _, alias := range aliases {
		if strings.Contains(alias.Dto, "|") {
			return nil, fmt.Errorf("params of %s: dto %q must be a single dto", alias.Class, alias.Dto)
		}
		_, className := core.ParseDtoPath(strings.TrimSpace(alias.Dto))
		fmt.Fprintf(&script, "typealias %s = %s\n", alias.Class, className)
	}

	return &core.CodeChunkCompiled{
		SuggestedFileName:  "ParamsAliases",
		SuggestedExtension: ".kt",
		ActualScript:       []byte(script.String()),
	}, nil
}
