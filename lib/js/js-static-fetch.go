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

type fetchStaticFunctionContext struct {
	EndpointUrl string

	ActionMethod     string
	RequestClass     string
	ResponseClass    string
	QueryStringClass string

	// Those requests with /:id/:item in them, will generate
	// a custom type, and will be available here.
	PathParameterTypeName string

	RequestHeadersClass string

	// The variable which will be used as default url
	DefaultUrlVariable string

	UrlCreatorFunction string

	UrlMethod string

	// by default, we assume it's a classic http call, since there might be non-standard http methods.
	IsClassicHttpCall bool

	// If it's a server side event endpoint
	IsSSE bool

	// For certain types, we need to make res.json() cast in fetch, if it's returning a dto,
	// or entity, or has fields. For text, html, or others, it does not require and makes no sense,
	// therefor needs to be casted res.text() from fetch perspective
	CastToJson bool
}

func GenerateTSParams(placeholders []string) string {
	if len(placeholders) == 0 {
		return "{}"
	}

	var builder strings.Builder
	builder.WriteString("{\n")
	for _, p := range placeholders {
		builder.WriteString(fmt.Sprintf("  %s: string;\n", p))
	}
	builder.WriteString("}")

	return builder.String()
}

func getCommonFetchArguments(fetchctx fetchStaticFunctionContext) []core.JsFnArgument {
	responseType := "unknown"
	requestHeaderType := "unknown"
	requestType := "unknown"

	if fetchctx.ResponseClass != "" {
		responseType = fetchctx.ResponseClass
	}
	if fetchctx.RequestHeadersClass != "" {
		requestHeaderType = fetchctx.RequestHeadersClass
	}

	claims := []core.JsFnArgument{
		{
			Key: "fetch.init",
			Ts:  fmt.Sprintf("init?: TypedRequestInit<%v, %v>", responseType, requestHeaderType),
			Js:  "init",
		},
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
			Key: "fetch.generic",
			Ts:  fmt.Sprintf("<%v, %v, %v>", responseType, requestType, requestHeaderType),
			Js:  "",
		},
		{
			Key: "query.params",
			Js:  "params",
			Ts:  "params: " + fetchctx.PathParameterTypeName,
		},
		{
			Key: "message.callback",
			Js:  "onMessage",
			Ts:  "onMessage?: (ev: MessageEvent) => void",
		},
	}

	return claims
}

// generates a static function, to developers prefer to make calls via axios
func FetchStaticHelper(fetchctx fetchStaticFunctionContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	queryParams := core.ExtractPlaceholdersInUrl(fetchctx.EndpointUrl)
	claims := []core.JsFnArgument{}

	claims = append(claims, getCommonFetchArguments(fetchctx)...)

	claimsRendered := core.ClaimRender(claims, ctx)

	const tmpl = `
	static Fetch = async (
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}

		{{ if .fetchctx.IsSSE }}
			|@message.callback|,
		{{ end }}

		|@fetch.qs|,
		|@fetch.init|,
		|@fetch.overrideUrl|
	) => {
		const res = await fetchx|@fetch.generic|(
			overrideUrl ?? {{  .fetchctx.UrlCreatorFunction -}}(
				{{ if .hasQueryParams }}
				params,
				{{ end }}
				qs
			),
			{
				method: {{  .fetchctx.UrlMethod -}},
				...(init || {})
			}
		)

		{{ if .fetchctx.IsClassicHttpCall }}
			{{ if .fetchctx.CastToJson }}
			const result = await res.json();
			{{ else }}
			const result = await res.text();
			{{ end }}

			{{ if .fetchctx.ResponseClass }}
				res.result = new {{ .fetchctx.ResponseClass }} (result);
			{{ else }}
				res.result = result;
			{{ end }}

			return res;
		{{ end }}

		{{ if .fetchctx.IsSSE }}
			return SSEFetch(res, onMessage, init?.signal || undefined);
		{{ end }}
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
				Objects:  []string{"fetchx"},
				Location: INTERNAL_SDK_JS_LOCATION,
			},
		},
	}

	if fetchctx.IsSSE {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, []core.CodeChunkDependency{
			{
				Objects:  []string{"SSEFetch"},
				Location: INTERNAL_SDK_JS_LOCATION,
			},
		}...)

	}

	if isTypeScript {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, []core.CodeChunkDependency{
			{
				Objects:  []string{"type TypedRequestInit"},
				Location: INTERNAL_SDK_JS_LOCATION,
			},
		}...)
	}

	return res, nil
}

// On final stage of compiling, this varialble will be replaced with context
// sdk location on the disk
var INTERNAL_SDK_JS_LOCATION string = "./sdk/js"
var INTERNAL_SDK_REACT_LOCATION string = "./sdk/react"
