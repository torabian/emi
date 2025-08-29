package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactMutationOptionsType struct {
	ActionMutationOptionsName string
	ActionName                string
	HasPathParameters         bool
}

// generates a TS type for mutation options
func ReactMutationOptionsTypeFunction(rmoptions reactMutationOptionsType, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	className := fmt.Sprintf("%vActionMutationOptions", core.ToUpper(rmoptions.ActionName))

	const tmpl = `
export type {{ .className }} = Omit<
	UseMutationOptions<unknown, unknown, unknown, unknown>,
	"mutationFn"
> &
	{{ .rmoptions.ActionMutationOptionsName }};
	`

	t := template.Must(template.New("reactmutationoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"rmoptions": rmoptions,
		"ctx":       ctx,
		"className": className,
	}); err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		ActualScript: buf.Bytes(),
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
	}

	if isTypeScript {
		res.CodeChunkDependenies = []core.CodeChunkDependency{
			{
				Objects:  []string{"type UseMutationOptions"},
				Location: "@tanstack/react-query",
			},
		}
	}

	return res, nil
}
