// A type generated for type script, which holds all information
// which we can modify for an action

package js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactUseQueryOptions struct {
	ActionQueryOptionsName string
	ActionName             string
	NewUrlFunctionName     string
	MetaDataClassName      string
}

// generates a static function, to developers prefer to make calls via axios
func ReactUseQueryOptionsFunction(useQueryOptions reactUseQueryOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	className := fmt.Sprintf("use%v", core.ToUpper(useQueryOptions.ActionName))
	const tmpl = `
		
export const {{ .className }} = (options: {{ .useQueryOptions.ActionQueryOptionsName }}) => {
	return useQuery({
		...options,
		queryKey: [
			{{ .useQueryOptions.NewUrlFunctionName }} (options.params, options.qs)
		],
		queryFn: () =>
		{{ .useQueryOptions.MetaDataClassName }}.Fetch(options.params, options.qs, {
			headers: options.headers,
		}),
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
