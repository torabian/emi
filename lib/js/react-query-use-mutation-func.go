package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactUseMutationOptions struct {
	ActionMutationOptionsName string
	ActionName                string
	MetaDataClassName         string
	HasPathParameters         bool
}

// generates the hook function use<ActionName>Mutation
func ReactUseMutationFunction(useMutationOptions reactUseMutationOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  "options: " + useMutationOptions.ActionMutationOptionsName,
			Js:  "options",
		},
	}
	className := fmt.Sprintf("use%vMutation", core.ToUpper(useMutationOptions.ActionName))

	const tmpl = `
export const {{ .className }} = (
	|@options.argument|
) => {
	return useMutation({
		mutationFn: (vars: unknown) =>
			{{ .useMutationOptions.MetaDataClassName }}.Fetch(
				{{ if .useMutationOptions.HasPathParameters }}
				options.params,
				{{ end }}
				options.qs,
				{
					body: vars,
					headers: options.headers,
				}
			),
		...(options || {}),

	});
};
	`

	t := template.Must(template.New("jsmutationoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"useMutationOptions": useMutationOptions,
		"ctx":                ctx,
		"className":          className,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	claimsRendered := core.ClaimRender(claims, ctx)
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		CodeChunkDependenies: []core.CodeChunkDependency{
			{
				Objects:  []string{"useMutation"},
				Location: "@tanstack/react-query",
			},
		},
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
	}

	return res, nil
}
