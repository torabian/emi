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

	ActionMethod          string
	RequestClass          string
	ResponseClass         string
	ResponseEnvelopeClass string
	QueryStringClass      string

	// Those requests with /:id/:item in them, will generate
	// a custom type, and will be available here.
	PathParameterTypeName string

	RequestHeadersClass string

	// The variable which will be used as default url
	DefaultUrlVariable string

	UrlCreatorFunction  string
	NativeFetchFunction string

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

func GetClassOrUnknown(value string) string {
	if value != "" {
		return value
	}

	return "unknown"
}

func getCommonFetchArguments(fetchctx fetchStaticFunctionContext) []core.JsFnArgument {
	responseType := GetClassOrUnknown(fetchctx.ResponseClass)
	requestHeaderType := GetClassOrUnknown(fetchctx.RequestHeadersClass)
	requestType := GetClassOrUnknown(fetchctx.RequestClass)

	// Wrap the class into the envelope type
	if fetchctx.ResponseEnvelopeClass != "" {
		responseType = fmt.Sprintf("%v<%v>", fetchctx.ResponseEnvelopeClass, responseType)
	}

	claims := []core.JsFnArgument{
		{
			Key: "fetch.init",
			Ts:  fmt.Sprintf("init?: TypedRequestInit<%v, %v>", requestType, requestHeaderType),
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

	if fetchctx.ResponseClass != "" {
		// If there is no envelope, passing the class with constructor is enough
		if fetchctx.ResponseEnvelopeClass == "" {
			claims = append(claims, core.JsFnArgument{
				Key: "response.cls",
				Ts:  fetchctx.ResponseClass,
				Js:  fetchctx.ResponseClass,
			})
		} else {

			statement := `(data) => { const envelope = new %v<%v>(data); envelope.updatePayload(new %v(envelope.getPayload())); return envelope;}`
			seq := fmt.Sprintf(statement, fetchctx.ResponseEnvelopeClass, fetchctx.ResponseClass, fetchctx.ResponseClass)

			// seq := fmt.Sprintf("(data) => new %v(data).updatePayload(new %v())", fetchctx.ResponseEnvelopeClass, fetchctx.ResponseClass)
			claims = append(claims, core.JsFnArgument{
				Key: "response.cls",
				Ts:  seq,
				Js:  seq,
			})

		}
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
  
	static Fetch$ = async (
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}
		|@fetch.qs|,
		|@fetch.init|,
		|@fetch.overrideUrl|,
	) => {
		return fetchx|@fetch.generic|(
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
	}

	static Fetch = async (
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}
		|@fetch.qs|,
		|@fetch.init|,
		|@message.callback|,
		|@fetch.overrideUrl|,
	) => {
	 

		const res = await {{ .fetchctx.NativeFetchFunction }}(
			{{ if .hasQueryParams }}
			params,
			{{ end }}
			qs,
			init,
			overrideUrl,
		);

		{{ if .fetchctx.IsClassicHttpCall }}

			return handleFetchResponse(
				res, 
				{{ if .fetchctx.ResponseClass }}
				|@response.cls|,
				{{ else }}
				undefined,
				{{ end }}
				onMessage,
				init?.signal,
			);
		{{ else }}
		 	return res
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
				Location: INTERNAL_SDK_JS_LOCATION + "/fetchx",
			},
		},
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, []core.CodeChunkDependency{
		{
			Objects:  []string{"handleFetchResponse"},
			Location: INTERNAL_SDK_JS_LOCATION + "/fetchx",
		},
	}...)

	if isTypeScript {
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, []core.CodeChunkDependency{
			{
				Objects:  []string{"type TypedRequestInit"},
				Location: INTERNAL_SDK_JS_LOCATION + "/fetchx",
			},
		}...)
	}

	return res, nil
}

// On final stage of compiling, this varialble will be replaced with context
// sdk location on the disk
var INTERNAL_SDK_JS_LOCATION string = "./sdk/common"
var INTERNAL_SDK_REACT_LOCATION string = "./sdk/react"
var INTERNAL_SDK_ENVELOPES_LOCATION string = "./sdk/envelopes/index"
