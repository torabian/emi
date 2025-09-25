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

type reactQueryCommonFnOptions struct {
	RequestClass      string
	MetaDataClassName string
	HasPathParameters bool
}

// Creates the queryFn or mutationFn function only which can be placed inside multiple functions
func reactQueryCommonFnFunction(options reactQueryCommonFnOptions, ctx core.MicroGenContext) (string, []core.CodeChunkDependency, error) {

	claims := []core.JsFnArgument{
		{
			Key: "response.state",
			Ts:  "<TypedResponse<unknown>>",
			Js:  "",
		},
	}
	claimsRendered := core.ClaimRender(claims, ctx)

	const tmpl = `
		
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;

	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState|@response.state|();

	const fn = (
		{{ if .hookOptions.RequestClass }}
			body: {{ .hookOptions.RequestClass }},
		{{ end }}
	) =>
		{
			setCompleteState(false);
			return {{ .hookOptions.MetaDataClassName }}.Fetch(
				{{ if .hookOptions.HasPathParameters }}
					options.params,
				{{ end }}
				{
					{{ if .hookOptions.RequestClass }}
						body,
					{{ end }}
					headers: options?.headers,
				},
				{
					creatorFn: options?.creatorFn,
					qs: options?.qs,
					ctx,
					onMessage: options?.onMessage,
					overrideUrl: options?.overrideUrl,
				}
			).then((x) => {
				x.done.then(() => {
					setCompleteState(true);
				});
				
				setResponse(x.response)
				return x.response.result;
			})
		}
 
	`

	t := template.Must(template.New("react-query-commmon-fn-function").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"hookOptions": options,
	}); err != nil {
		return "", []core.CodeChunkDependency{}, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	deps := []core.CodeChunkDependency{}
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	if isTypeScript {
		deps = append(deps, core.CodeChunkDependency{
			Objects:  []string{"type TypedResponse"},
			Location: INTERNAL_SDK_JS_LOCATION + "/fetchx",
		})
	}
	return templateResult, deps, nil
}
