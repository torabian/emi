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

func JsActionFetchAndMetaData(action core.EmiRpcAction, realms jsActionRealms, ctx core.MicroGenContext) (*core.CodeChunkCompiled, fetchStaticFunctionContext, error) {

	className := action.GetName()
	fetchctx := fetchStaticFunctionContext{
		DefaultUrlVariable:  fmt.Sprintf("%v.URL", className),
		UrlCreatorFunction:  fmt.Sprintf("%v.NewUrl", className),
		NativeFetchFunction: fmt.Sprintf("%v.Fetch$", className),
		UrlMethod:           fmt.Sprintf("%v.Method", className),
		EndpointUrl:         action.GetUrl(),
		QueryStringClass:    "URLSearchParams",
		// By default, use the classic http call, which covers files, json, sse, etc...
		IsClassicHttpCall: true,
		ActionMethod:      action.GetMethod(),
	}

	if action.MethodUpper() == EMI_METHOD_REACTIVE {
		fetchctx.IsClassicHttpCall = false
	}

	if realms.RequestHeadersClass != nil {
		requestHeadersClassToken := findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS)
		if requestHeadersClassToken != nil {
			fetchctx.RequestHeadersClass = requestHeadersClassToken.Value
		}
	}

	if realms.RequestClass != nil {
		requestClassToken := findTokenByName(realms.RequestClass.Tokens, TOKEN_OBJ_CLASS)
		if requestClassToken != nil {
			fetchctx.RequestClass = requestClassToken.Value
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
				Objects:  []string{"buildUrl", "withPrefix", "isPlausibleObject"},
				Location: INTERNAL_SDK_JS_LOCATION,
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
 

  static Method = '{{ .fetchctx.ActionMethod }}';
  

  {{ if .fetchStaticFunction }}
  	{{ b2s .fetchStaticFunction.ActualScript }}
  {{ end }}

  {{ if .axiosStaticFunction }}
  	{{ b2s .axiosStaticFunction.ActualScript }}
  {{ end }}

  {{ if .websocketCreateFunction }}
  	{{ b2s .websocketCreateFunction.ActualScript }}
  {{ end }}
  
}
`

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_NEW_URL_FN,
		Value: fetchctx.UrlCreatorFunction,
	})

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_URL_METHOD,
		Value: fetchctx.UrlMethod,
	})

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_ACTUAL_METHOD,
		Value: fetchctx.ActionMethod,
	})

	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
		Name:  TOKEN_ROOT_CLASS,
		Value: className,
	})

	if len(core.ExtractPlaceholdersInUrl(action.GetUrl())) > 0 {
		fetchctx.PathParameterTypeName = fmt.Sprintf("%vPathParameter", className)
	}

	if realms.ResponseClass != nil {
		responseClassToken := findTokenByName(realms.ResponseClass.Tokens, TOKEN_OBJ_CLASS)
		if responseClassToken != nil {

			// Not sure about this yet. Primitives also can be a class, right?
			// therefor they might not need to cast to json, but still you need to create a class out of them.
			fetchctx.CastToJson = true
			fetchctx.ResponseClass = responseClassToken.Value
		}
	}

	var axiosStaticFunction *core.CodeChunkCompiled

	if isAxiosSupported && fetchctx.IsClassicHttpCall {
		// Add the axios helper function to the response depencencies,
		fn, err := AxiosStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, fetchctx, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fn.CodeChunkDependenies...)
		axiosStaticFunction = fn
	}

	var fetchStaticFunction *core.CodeChunkCompiled

	if strings.ToUpper(fetchctx.ActionMethod) != EMI_METHOD_REACTIVE {
		// add the native fetch function to the axios
		fn, err := FetchStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, fetchctx, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fn.CodeChunkDependenies...)
		fetchStaticFunction = fn
	}

	var websocketCreateFunction *core.CodeChunkCompiled
	if strings.ToUpper(fetchctx.ActionMethod) == EMI_METHOD_REACTIVE {

		// add the native fetch function to the axios
		fn, err := CreateWebSocketStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, fetchctx, err
		}
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fn.CodeChunkDependenies...)
		websocketCreateFunction = fn
	}

	t := template.Must(template.New("qsclass").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":                  action,
		"axiosStaticFunction":     axiosStaticFunction,
		"fetchStaticFunction":     fetchStaticFunction,
		"websocketCreateFunction": websocketCreateFunction,
		"shouldExport":            true,
		"realms":                  realms,
		"fetchctx":                fetchctx,
		"className":               className,
		"queryParams":             core.ExtractPlaceholdersInUrl(action.GetUrl()),
	}); err != nil {
		return nil, fetchctx, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res.ActualScript = []byte(templateResult)

	return res, fetchctx, nil
}

var EMI_METHOD_REACTIVE = "REACTIVE"
var EMI_METHOD_SSE = "SSE"
