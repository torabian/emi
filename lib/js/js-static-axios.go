// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// generates a static function, to developers prefer to make calls via axios
// axios context does not exists, uses the fetch native data
func AxiosStaticHelper(fetchctx fetchStaticFunctionContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	claims := []core.JsFnArgument{
		{
			Key: "axios.clientInstance",
			Ts:  "clientInstance: AxiosInstance",
			Js:  "clientInstance",
		},
		{
			Key: "axios.config",
			Ts:  "config: AxiosRequestConfig<unknown>",
			Js:  "config",
		},
		{
			Key: "axios.request.generic",
			Js:  "",
			Ts:  "<unknown, AxiosResponse<unknown>, unknown>",
		},
		{
			Key: "axios.returnType",
			Js:  "=",
			Ts: `: (
		clientInstance: AxiosInstance,
		config: AxiosRequestConfig<unknown>,
	)  => Promise<AxiosResponse<unknown>> = (
		clientInstance,
		config,
	) => `,
		},
		{
			Key: "axios.request.generic",
			Js:  "",
			Ts:  "<unknown, AxiosResponse<unknown>, unknown>",
		},
	}

	claimsRendered := core.ClaimRender(claims, ctx)

	const tmpl = `

	static Axios |@axios.returnType|
		clientInstance
		.request|@axios.request.generic|(
			{
				method: {{  .fetchctx.UrlMethod -}},
				...(config || {})
			}
		)

		{{ if and .fetchctx.CastToJson .fetchctx.ResponseClass }}
		.then((res) => {
			return {
			...res,

			
			// if there is a output class, create instance out of it.
			data: new {{ .fetchctx.ResponseClass }}(res.data),
			};
		});
		{{ end }}
	`

	t := template.Must(template.New("axioshelper").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"claims":   claimsRendered,
		"fetchctx": fetchctx,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res := &core.CodeChunkCompiled{}
	res.ActualScript = []byte(templateResult)

	// Axios import is only needed when it's typescript in this case
	if isTypeScript {
		res.CodeChunkDependenies = []core.CodeChunkDependency{
			{
				Objects:  []string{"AxiosInstance", "AxiosRequestConfig", "AxiosResponse"},
				Location: "axios",
			},
		}
	}

	return res, nil
}
