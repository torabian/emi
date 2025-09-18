// file: react-query-use-query-func.go
// A type generated for type script, which holds all information
// which we can modify for an action

package js

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type reactQueryCommonFnOptions struct {
	RequestClass      string
	MetaDataClassName string
	HasPathParameters bool
}

// Creates the queryFn or mutationFn function only which can be placed inside multiple functions
func reactQueryCommonFnFunction(options reactQueryCommonFnOptions) (string, error) {
	const tmpl = `
		
	const globalCtx = useFetchxContext(); 
	const ctx = options?.ctx ?? globalCtx ?? undefined;

	const [isCompleted, setCompleteState] = useState(false);
	const [response, setResponse] = useState<TypedResponse<unknown>>();

	const fn = (
		{{ if .hookOptions.RequestClass }}
			body: {{ .hookOptions.RequestClass }},
		{{ end }}
	) =>
		{
			setCompleteState(false);
			{{ .hookOptions.MetaDataClassName }}.Fetch(
				{{ if .hookOptions.HasPathParameters }}
					options.params,
				{{ end }}
				options?.creatorFn,
				options?.qs,
				{
					{{ if .hookOptions.RequestClass }}
						body,
					{{ end }}
					headers: options?.headers,
				},
				options?.onMessage,
				options?.overrideUrl,
				ctx,
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
		return "", err
	}

	return buf.String(), nil
}
