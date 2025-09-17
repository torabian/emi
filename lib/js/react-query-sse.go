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

type reactUseSSEOptions struct {
	ActionName        string
	MetaDataClassName string
	Fetchctx          fetchStaticFunctionContext
}

func ReactUseSSEHook(options reactUseSSEOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	className := fmt.Sprintf("use%v", core.ToUpper(options.ActionName))

	claims := getCommonFetchArguments(options.Fetchctx)

	const tmpl = `
		
export const {{ .className }} = (options: {
	|@fetch.qs|,
	|@fetch.init|,
	|@fetch.overrideUrl|
}) => {
	return useSse({{ .options.MetaDataClassName }}.Fetch, options);
};


	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"options":   options,
		"ctx":       ctx,
		"className": className,
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
				Location: INTERNAL_SDK_REACT_LOCATION + "/useSse",
				Objects:  []string{"useSse"},
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
