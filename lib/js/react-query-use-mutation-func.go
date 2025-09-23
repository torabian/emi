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

	fn, extraDeps, err := reactQueryCommonFnFunction(reactQueryCommonFnOptions{
		RequestClass:      useMutationOptions.RequestClass,
		MetaDataClassName: useMutationOptions.MetaDataClassName,
		HasPathParameters: useMutationOptions.HasPathParameters,
	}, ctx)

	if err != nil {
		return nil, err
	}

	if useMutationOptions.HasPathParameters {
		tsValue = "options: " + useMutationOptions.ActionMutationOptionsName
	}

	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  tsValue,
			Js:  "options",
		},
		{
			Key: "fn",
			Ts:  fn,
			Js:  fn,
		},
	}
	className := fmt.Sprintf("use%v", core.ToUpper(useMutationOptions.ActionName))

	const tmpl = `
export const {{ .className }} = (
	|@options.argument|
) => {

	|@fn|
 
	const result =  useMutation({
		mutationFn: fn,
		...(options || {}),
	});

	return {
		...result,
		isCompleted,
		response
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
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Objects:  []string{"useMutation"},
				Location: "@tanstack/react-query",
			},
			{
				Objects:  []string{"useState"},
				Location: "react",
			},
			{
				Objects:  []string{"useFetchxContext"},
				Location: INTERNAL_SDK_REACT_LOCATION + "/useFetchx",
			},
		},
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
	}

	res.CodeChunkDependensies = append(res.CodeChunkDependensies, extraDeps...)

	return res, nil
}
