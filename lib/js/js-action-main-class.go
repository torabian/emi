// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package js

import (
	"bytes"
	"fmt"
	"reflect"
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
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, fetchStaticFunctionContext{}, nil
	}

	className := realms.ActionName
	// --tags no-class: the request/response dto files this action references
	// never got a class body (see js-common-object.go) - resolve their names
	// via TOKEN_TYPEDEF_NAME (the standalone type/typedef identifier) instead
	// of TOKEN_OBJ_CLASS (which - by design, see js-tokens.go - still resolves
	// to the same bare name a class would have had, so it's never actually a
	// dangling reference on its own; it's just not useful to build a `new
	// X(...)` off of anymore, which is exactly what fetchctx.NoClass tells
	// getCreatorFnInfo/getCommonFetchArguments (js-static-fetch.go) not to do).
	noClass := ctx.HasTag(NoClass)
	requestResponseToken := TOKEN_OBJ_CLASS
	if noClass {
		requestResponseToken = TOKEN_TYPEDEF_NAME
	}

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
		NoClass:           noClass,
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
		requestClassToken := findTokenByName(realms.RequestClass.Tokens, requestResponseToken)
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
	claimsRendered := ClaimRender(claims, ctx)

	res := &core.CodeChunkCompiled{
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Objects:  []string{"buildUrl"},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "buildUrl"),
			},
		},
	}

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
		responseClassToken := findTokenByName(realms.ResponseClass.Tokens, requestResponseToken)
		if responseClassToken != nil {

			// Not sure about this yet. Primitives also can be a class, right?
			// therefor they might not need to cast to json, but still you need to create a class out of them.
			fetchctx.CastToJson = true
			fetchctx.ResponseClass = responseClassToken.Value
		}

		responseEnvelopeToken := findTokenByName(realms.ResponseClass.Tokens, TOKEN_RESPONSE_ENVELOPE)
		if responseEnvelopeToken != nil {
			fetchctx.ResponseEnvelopeClass = responseEnvelopeToken.Value
		}
	}

	var fetchStaticFunction *core.CodeChunkCompiled

	if strings.ToUpper(fetchctx.ActionMethod) != EMI_METHOD_REACTIVE {
		fn, err := FetchStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, fetchctx, err
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, fn.CodeChunkDependensies...)
		fetchStaticFunction = fn
	}

	var websocketCreateFunction *core.CodeChunkCompiled
	if strings.ToUpper(fetchctx.ActionMethod) == EMI_METHOD_REACTIVE {

		fn, err := CreateWebSocketStaticHelper(fetchctx, ctx)
		if err != nil {
			return nil, fetchctx, err
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, fn.CodeChunkDependensies...)
		websocketCreateFunction = fn
	}

	const tmpl = `/**
 * {{.className}}
 */


export class {{ .className }} { //
  static URL = '{{ .safeUrl }}';

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
 

  static Method = '{{ UPPER .fetchctx.ActionMethod }}';
  

  {{ if .fetchStaticFunction }}
  	{{ b2s .fetchStaticFunction.ActualScript }}
  {{ end }}

  {{ if .websocketCreateFunction }}
  	{{ b2s .websocketCreateFunction.ActualScript }}
  {{ end }}

  {{ if .definition }}
  static Definition = {{ .definition }}
  {{ end }}
  
}
`

	t := template.Must(template.New("qsclass").Funcs(core.CommonMap).Parse(tmpl))

	// --tags no-definition: the raw JSON dump of the whole action (url, method,
	// in/out shapes, ...) is handy for introspection/tooling, but it's also a
	// full second copy of everything the class above already expresses through
	// real code - skippable for a size-constrained target the same way
	// --tags no-class skips the dto class body.
	definition := action.GetDefinition()
	if ctx.HasTag(NoDefinition) {
		definition = ""
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":                  action,
		"safeUrl":                 core.RemoveTypeAnnotations(action.GetUrl()),
		"fetchStaticFunction":     fetchStaticFunction,
		"websocketCreateFunction": websocketCreateFunction,
		"shouldExport":            true,
		"definition":              definition,
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

	if fetchStaticFunction != nil {
		if token := findTokenByName(fetchStaticFunction.Tokens, TOKEN_CREATOR_FN); token != nil {
			res.Tokens = append(res.Tokens, *token)
		}
	}

	return res, fetchctx, nil
}

var EMI_METHOD_REACTIVE = "REACTIVE"
var EMI_METHOD_SSE = "SSE"
