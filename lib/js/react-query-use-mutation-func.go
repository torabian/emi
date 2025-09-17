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
	HasCreatorFunction        bool
	RequestClass              string
}

// generates the hook function use<ActionName>Mutation
func ReactUseMutationFunction(useMutationOptions reactUseMutationOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	tsValue := "options?: " + useMutationOptions.ActionMutationOptionsName

	if useMutationOptions.HasPathParameters {
		tsValue = "options: " + useMutationOptions.ActionMutationOptionsName
	}

	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  tsValue,
			Js:  "options",
		},
	}
	className := fmt.Sprintf("use%v", core.ToUpper(useMutationOptions.ActionName))

	const tmpl = `
export const {{ .className }} = (
	|@options.argument|
) => {
 	const [isCompleted, setCompleteState] = useState(false);

	const mutationResult =  useMutation({
		mutationFn: (body: {{ .useMutationOptions.RequestClass }}) =>
			{{ .useMutationOptions.MetaDataClassName }}.Fetch(
				{{ if .useMutationOptions.HasPathParameters }}
				options.params,
				{{ end }}
				{{ if .useMutationOptions.HasCreatorFunction }}
				options?.creatorFn,
				{{ end }}
				options?.qs,
				{
					body,
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});

				return x.response;
			}),
		...(options || {}),

	});

	return {
		...mutationResult,
		isCompleted
	}
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
			{
				Objects:  []string{"useState"},
				Location: "react",
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
