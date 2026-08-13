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

	// --tags no-class (js-compiler-tags.go): RequestClass/ResponseClass above
	// still carry a name here, but it's the standalone type/typedef name (see
	// TOKEN_TYPEDEF_NAME, js-tokens.go) rather than a real class - there's
	// nothing to `new` it into. getCreatorFnInfo/getCommonFetchArguments use
	// this to skip instantiation entirely and just hand back the parsed
	// response as-is, typed (TS generics / a JSDoc @returns) but never
	// constructed.
	NoClass bool
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

type CreateFn struct {
	DefinitionStatement string
	CreatorStatement    string
	ArgStatement        string

	// When a return type is multiple different dtos, how can we know what is each instance?
	// It's impossible. Then developer needs to be forced to provide tht.
	IsAmbigousCreator bool
}

func getCreatorFnInfo(fetchctx fetchStaticFunctionContext, isTypescript bool) *CreateFn {
	isAmbigousCreator := false
	// There is nothing to create instance, return early.
	if fetchctx.ResponseClass == "" {
		return nil
	}
	// --tags no-class: ResponseClass is a type/typedef name here, not a real
	// class - there's nothing to construct at all (see the NoClass doc comment
	// on fetchStaticFunctionContext above).
	if fetchctx.NoClass {
		return nil
	}

	// This is a way to detect if multiple things are returning. Might be not the best solution
	// but works perfectly fine.
	isAmbigousCreator = strings.Contains(fetchctx.ResponseClass, "|")

	// universal instantiator :D
	instantiator := fmt.Sprintf("(item) => new %v(item)", fetchctx.ResponseClass)

	// In case of ambigous creator, we do not have the instantiator, and developer needs to provide it.
	if isAmbigousCreator {
		instantiator = fmt.Sprintf("() => { throw 'Return type can be different types, you need to pass a function to determine, which class needs to be created: %v' }", fetchctx.ResponseClass)
	}

	statement := fmt.Sprintf("creatorFn %v", instantiator)
	definition := fmt.Sprintf("(item: unknown) => %v", fetchctx.ResponseClass)

	if isTypescript {
		statement = fmt.Sprintf("creatorFn: %v = %v", definition, instantiator)
	}

	return &CreateFn{
		ArgStatement:        statement,
		DefinitionStatement: definition,
		CreatorStatement:    instantiator,
		IsAmbigousCreator:   isAmbigousCreator,
	}

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
			Key: "fetch.ctx",
			Ts:  "ctx?: FetchxContext | null",
			Js:  "ctx",
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
			// --tags no-class: nothing to construct - creatorFn is never declared
			// in scope at all in this mode (getCreatorFnInfo returns nil), so
			// referencing it here (like the default identity-fallback ternary
			// below does) would be a ReferenceError - hand the parsed response
			// back completely untouched instead.
			responseCls := "(item) => (creatorFn ? creatorFn(item) : item)"
			if fetchctx.NoClass {
				responseCls = "(item) => item"
			}
			claims = append(claims, core.JsFnArgument{
				Key: "response.cls",
				Ts:  responseCls,
				Js:  responseCls,
			})
		} else {

			statementTs := `(data) => {
					const resp = new %v<%v>();
					if (creatorFn) {
						resp.setCreator(creatorFn);
					}
					resp.inject(data);

					return resp;

			}`
			// --tags no-class: same reasoning as the non-envelope branch above -
			// creatorFn is never in scope, so this can't reference it either.
			// The envelope class itself is still constructed (it's a fixed SDK
			// utility, not a generated dto class - out of this tag's scope), just
			// never handed a creator to wrap its payload with.
			if fetchctx.NoClass {
				statementTs = `(data) => {
						const resp = new %v<%v>();
						resp.inject(data);

						return resp;

				}`
			}
			seqTs := fmt.Sprintf(statementTs, fetchctx.ResponseEnvelopeClass, fetchctx.ResponseClass)

			statementJs := `(data) => {
					const resp = new %v();
					if (creatorFn) {
						resp.setCreator(creatorFn);
					}
					resp.inject(data);

					return resp;

			}`
			if fetchctx.NoClass {
				statementJs = `(data) => {
						const resp = new %v();
						resp.inject(data);

						return resp;

				}`
			}
			seqJs := fmt.Sprintf(statementJs, fetchctx.ResponseEnvelopeClass)

			claims = append(claims, core.JsFnArgument{
				Key: "response.cls",
				Ts:  seqTs,
				Js:  seqJs,
			})

		}
	}

	return claims
}

func FetchStaticHelper(fetchctx fetchStaticFunctionContext, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	isTypeScript := ctx.HasTag(Typescript)
	queryParams := core.ExtractPlaceholdersInUrl(fetchctx.EndpointUrl)
	claims := []core.JsFnArgument{}

	claims = append(claims, getCommonFetchArguments(fetchctx)...)

	claimsRendered := ClaimRender(claims, ctx)

	const tmpl = `
  
	static Fetch$ = async (
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}
		|@fetch.qs|,
		|@fetch.ctx|,
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
			},
			ctx
		)
	}

	{{ if .jsdocReturnsComment }}
	{{ .jsdocReturnsComment }}
	{{ end }}
	static Fetch = async (
		{{ if .hasQueryParams }}
			|@query.params|,
		{{ end }}
		|@fetch.init|,

		{
			{{ if .creatorFn }}
			creatorFn,
			{{ end }}
			qs,
			ctx,
			onMessage,
			overrideUrl
		} {{ if .isTypeScript }}
			: {
			{{ if .creatorFn }}
				creatorFn?: ({{ .creatorFn.DefinitionStatement }}) | undefined,
			{{ end }}
			|@fetch.qs|,
			|@fetch.ctx|,
			|@message.callback|,
			|@fetch.overrideUrl|,		
		} 
			{{ end }} = {
		 	{{ if .creatorFn }}
				creatorFn: {{ .creatorFn.CreatorStatement }},
			{{ end }}
		}
	) => {
	 	{{ if .creatorFn }}
		creatorFn = creatorFn || ({{ .creatorFn.CreatorStatement }})
		{{ end }}
		const res = await {{ .fetchctx.NativeFetchFunction }}(
			{{ if .hasQueryParams }}
			params,
			{{ end }}
			qs,
			ctx,
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

	creatorFn := getCreatorFnInfo(fetchctx, isTypeScript)

	// --tags no-class, plain JS only: TypeScript mode already gets the response
	// type for free via the |@fetch.generic| generic parameter above (which
	// resolves to the same typedef name once fetchctx.ResponseClass is
	// redirected to it - see JsActionFetchAndMetaData) - JS has no generics
	// syntax, so this is the JSDoc equivalent, the only way that same
	// information reaches a plain-JS caller's editor.
	jsdocReturnsComment := ""
	if !isTypeScript && fetchctx.NoClass && fetchctx.ResponseClass != "" {
		jsdocReturnsComment = NewJsDoc("\t").
			Add(fmt.Sprintf("Resolves with `{ response, done }` - `response.result` is typed as {@link %v}.", fetchctx.ResponseClass)).
			Add("@returns {Promise<any>}").
			String()
	}

	t := template.Must(template.New("fetchstatichelper").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"claims":              claimsRendered,
		"creatorFn":           creatorFn,
		"fetchctx":            fetchctx,
		"isTypeScript":        isTypeScript,
		"hasQueryParams":      len(queryParams) > 0,
		"jsdocReturnsComment": jsdocReturnsComment,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Objects:  []string{"fetchx"},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
			},
		},
		Tokens: []core.GeneratedScriptToken{},
	}

	if isTypeScript {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{"type FetchxContext"},
			Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
		})
	}

	if creatorFn != nil {
		res.Tokens = append(res.Tokens, core.GeneratedScriptToken{
			Name:  TOKEN_CREATOR_FN,
			Value: creatorFn.ArgStatement,
		})
	}

	res.CodeChunkDependensies = append(res.CodeChunkDependensies, []core.CodeChunkDependency{
		{
			Objects:  []string{"handleFetchResponse"},
			Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
		},
	}...)

	if isTypeScript {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, []core.CodeChunkDependency{
			{
				Objects:  []string{"type TypedRequestInit"},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
			},
		}...)
	}

	return res, nil
}

// On final stage of compiling, this varialble will be replaced with context
// sdk location on the disk
var INTERNAL_SDK_JS_LOCATION string = "{sdk}/common"
var INTERNAL_SDK_REACT_LOCATION string = "{sdk}/react"
var INTERNAL_SDK_ENVELOPES_LOCATION string = "{sdk}/envelopes/index"
