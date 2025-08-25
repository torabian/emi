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

func findTokenByName(realms []core.GeneratedScriptToken, name string) *core.GeneratedScriptToken {

	for _, item := range realms {
		if item.Name == name {
			return &item
		}
	}

	return nil
}

func JsActionFetchAndMetaData(action *core.Module3Action, realms jsActionRealms, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	className := fmt.Sprintf("Fetch%vAction", core.ToUpper(action.Name))
	fetchctx := fetchStaticFunctionContext{
		DefaultUrlVariable: fmt.Sprintf("%v.URL", className),
		UrlCreatorFunction: fmt.Sprintf("%v.NewUrl", className),
		EndpointUrl:        action.Url,
	}

	if realms.RequestHeadersClass != nil {
		requestHeadersClassToken := findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS)
		if requestHeadersClassToken != nil {
			fetchctx.RequestHeadersClass = requestHeadersClassToken.Value
		}
	}

	if realms.QueryStringClass != nil {
		qsClassToken := findTokenByName(realms.QueryStringClass.Tokens, TOKEN_ROOT_CLASS)
		if qsClassToken != nil {
			fetchctx.QueryStringClass = qsClassToken.Value
		}
	}

	claims := []core.JsFnArgument{
		{
			Key: "query.params",

			// This is not perfect, the classname needs to come from previous state
			Ts: fmt.Sprintf("params: %vPathParameter", className),
			Js: "params",
		},
		{
			Key: "qs",

			// This is not perfect, the classname needs to come from previous state
			Ts: fmt.Sprintf("qs?: %v", fetchctx.QueryStringClass),
			Js: "qs",
		},
	}
	claimsRendered := core.ClaimRender(claims, ctx)

	isAxiosSupported := strings.Contains(ctx.Tags, GEN_AXIOS_COMPATIBILITY)
	res := &core.CodeChunkCompiled{
		CodeChunkDependenies: []core.CodeChunkDependency{
			{
				Objects:  []string{"buildUrl", "fetchx"},
				Location: INTERNAL_SDK_LOCATION,
			},
		},
	}
	// How to do it iterte and call Compile?

	const tmpl = `/**
 * {{.className}}
 */


export class {{ .className }} {
  static URL = '{{ .action.Url }}';

  static NewUrl = (
	{{ if .queryParams }}
	|@query.params|,
	{{ end }}
	|@qs|
  ) => buildUrl(
		{{ .className }}.URL,
		
		{{ if .queryParams }}
		params,
		{{ else }}
		 undefined,
		{{ end }}
		qs
	);
 

  static Method = '{{ .action.Method }}';

  {{ .axiosStaticFunction }}
  
  {{ .fetchStaticFunction }}
}
`

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_NEW_URL_FN,
		Value: fetchctx.UrlCreatorFunction,
	})

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_ROOT_CLASS,
		Value: className,
	})

	if len(core.ExtractPlaceholdersInUrl(action.Url)) > 0 {
		fetchctx.PathParameterTypeName = fmt.Sprintf("%vPathParameter", className)
	}

	if realms.ResponseClass != nil {
		responseClassToken := findTokenByName(realms.ResponseClass.Tokens, TOKEN_ROOT_CLASS)
		if responseClassToken != nil {

			// Not sure about this yet. Primitives also can be a class, right?
			// therefor they might not need to cast to json, but still you need to create a class out of them.
			fetchctx.CastToJson = true
			fetchctx.ResponseClass = responseClassToken.Value
		}
	}

	var axiosStaticFunction = ""
	if isAxiosSupported {

		// Add the axios helper function to the response depencencies,
		fn, err := AxiosStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fn.CodeChunkDependenies...)
		axiosStaticFunction = string(fn.ActualScript)
	}

	// add the native fetch function to the axios
	fetchStaticFunction, err := FetchStaticHelper(fetchctx, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, fetchStaticFunction.CodeChunkDependenies...)

	t := template.Must(template.New("qsclass").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":              action,
		"axiosStaticFunction": axiosStaticFunction,
		"fetchStaticFunction": string(fetchStaticFunction.ActualScript),
		"shouldExport":        true,
		"realms":              realms,
		"fetchctx":            fetchctx,
		"className":           className,
		"queryParams":         core.ExtractPlaceholdersInUrl(action.Url),
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res.ActualScript = []byte(templateResult)

	return res, nil
}
