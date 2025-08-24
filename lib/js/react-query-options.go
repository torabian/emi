// A type generated for type script, which holds all information
// which we can modify for an action

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactQueryOptionsType struct {
	ActionQueryOptionsName string
	ActionName             string
}

// generates a static function, to developers prefer to make calls via axios
func ReactQueryOptionsTypeFunction(rqoptions reactQueryOptionsType, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	className := fmt.Sprintf("%vActionQueryOptions", core.ToUpper(rqoptions.ActionName))
	const tmpl = `
export type {{ .className }} = Omit<
	UseQueryOptions<unknown, unknown, unknown, unknown[]>,
	"queryKey"
> &
	{{ .rqoptions.ActionQueryOptionsName }};

	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"rqoptions": rqoptions,
		"ctx":       ctx,
		"className": className,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
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
				Objects:  []string{"type UseQueryOptions"},
				Location: "@tanstack/react-query",
			},
		}
	}

	return res, nil
}
