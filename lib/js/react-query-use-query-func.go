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
	MetaDataClassName      string
	HasPathParameters      bool
}

// generates a static function, to developers prefer to make calls via axios
func ReactUseQueryOptionsFunction(useQueryOptions reactUseQueryOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  "options: " + useQueryOptions.ActionQueryOptionsName,
			Js:  "options",
		},
	}
	className := fmt.Sprintf("use%v", core.ToUpper(useQueryOptions.ActionName))
	const tmpl = `
		
export const {{ .className }} = (
	|@options.argument|
) => {
	return useQuery({
		queryKey: [
			{{ .useQueryOptions.NewUrlFunctionName }} (
			 	{{ if .useQueryOptions.HasPathParameters }}
				options.params,
				{{ end }}
				
				options.qs
			)
		],
		queryFn: () =>
		{{ .useQueryOptions.MetaDataClassName }}.Fetch(
		 	{{ if .useQueryOptions.HasPathParameters }}
				options.params,
			{{ end }}
			options.qs,
			{
				headers: options.headers,
			}
		),
		...(options || {}),
	});
};


	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"useQueryOptions": useQueryOptions,
		"ctx":             ctx,
		"className":       className,
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
				Objects:  []string{"useQuery"},
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
