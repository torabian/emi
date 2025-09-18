// file: react-query-use-query-func.go
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

type reactUseQueryOptions struct {
	ActionQueryOptionsName string
	ActionName             string
	NewUrlFunctionName     string
	RequestClass           string
	MetaDataClassName      string
	HasPathParameters      bool
	ActionRealms           jsActionRealms
}

func ReactUseQueryOptionsFunction(useQueryOptions reactUseQueryOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	fn, err := reactQueryCommonFnFunction(reactQueryCommonFnOptions{
		RequestClass:      useQueryOptions.RequestClass,
		MetaDataClassName: useQueryOptions.MetaDataClassName,
		HasPathParameters: useQueryOptions.HasPathParameters,
	})

	if err != nil {
		return nil, err
	}

	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  "options: " + useQueryOptions.ActionQueryOptionsName,
			Js:  "options",
		},
		{
			Key: "fn",
			Ts:  fn,
			Js:  fn,
		},
	}

	className := fmt.Sprintf("use%v", core.ToUpper(useQueryOptions.ActionName))
	const tmpl = `
		
export const {{ .className }}Query = (
	|@options.argument|
) => {

	|@fn|

	const result = useQuery({
		queryKey: [
			{{ .hookOptions.NewUrlFunctionName }} (
			 	{{ if .hookOptions.HasPathParameters }}
				options.params,
				{{ end }}
				
				options?.qs
			)
		],
		queryFn: fn,
		...(options || {}),
	});

	return {
		...result,
		isCompleted,
		response
	}
};


	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"hookOptions": useQueryOptions,
		"ctx":         ctx,
		"className":   className,
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
				Objects:  []string{"useQuery"},
				Location: "@tanstack/react-query",
			},
			{
				Objects:  []string{"useState"},
				Location: "react",
			},
			{
				Objects:  []string{"type TypedResponse"},
				Location: INTERNAL_SDK_JS_LOCATION + "/fetchx",
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

	return res, nil
}
