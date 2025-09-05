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

func CreateWebSocketStaticHelper(fetchctx fetchStaticFunctionContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	requestType := "unknown"
	responseType := "unknown"
	requestHeaderType := "unknown"
	MessageFactoryClass := "undefined"

	if fetchctx.ResponseClass != "" {
		responseType = fetchctx.ResponseClass
		MessageFactoryClass = fetchctx.ResponseClass
	}

	if fetchctx.RequestClass != "" {
		requestType = fetchctx.RequestClass
	}
	if fetchctx.RequestHeadersClass != "" {
		requestHeaderType = fetchctx.RequestHeadersClass
	}

	queryParams := core.ExtractPlaceholdersInUrl(fetchctx.EndpointUrl)
	claims := []core.JsFnArgument{
		{
			Key: "fetch.qs",
			Ts:  "qs?: " + fetchctx.QueryStringClass,
			Js:  "qs",
		},
		{
			Key: "fetch.overrideUrl",
			Ts:  "overrideUrl?: string",
			Js:  "overrideUrl",
		},
		{
			Key: "generic",
			Ts:  fmt.Sprintf("<%v, %v>", requestType, responseType),
			Js:  "",
		},
		{
			Key: "query.params",
			Js:  "params",
			Ts:  "params: " + fetchctx.PathParameterTypeName,
		},
		{
			Key: "request.header",
			Js:  "Headers",
			Ts:  requestHeaderType,
		},
		{
			Key: "MessageFactoryClass",
			Js:  MessageFactoryClass,
			Ts:  MessageFactoryClass,
		},
	}

	claimsRendered := core.ClaimRender(claims, ctx)

	const tmpl = `
	static Create = (
		|@fetch.overrideUrl|,
		|@fetch.qs|,
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}
	) => {
		const url = overrideUrl ?? {{  .fetchctx.UrlCreatorFunction -}}(
			{{ if .hasQueryParams }}
			params,
			{{ end }}
			qs
		)
			
		return new WebSocketX|@generic|(
			url,
			undefined,
			{
				MessageFactoryClass: |@MessageFactoryClass|,
			}
		);

	}
	`

	t := template.Must(template.New("axioshelper").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"claims":         claimsRendered,
		"fetchctx":       fetchctx,
		"hasQueryParams": len(queryParams) > 0,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		CodeChunkDependenies: []core.CodeChunkDependency{
			{
				Objects:  []string{"WebSocketX"},
				Location: INTERNAL_SDK_JS_LOCATION,
			},
		},
	}

	return res, nil
}
