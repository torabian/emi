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

type reactUseQueryOptions struct {
	ActionQueryOptionsName string
	ActionName             string
	NewUrlFunctionName     string
	RequestClass           string
	MetaDataClassName      string
	HasPathParameters      bool
	ActionRealms           jsActionRealms
}

func ReactUseQueryOptionsFunction(useQueryOptions reactUseQueryOptions, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	fn, extraDeps, err := reactQueryCommonFnFunction(reactQueryCommonFnOptions{
		RequestClass:      useQueryOptions.RequestClass,
		MetaDataClassName: useQueryOptions.MetaDataClassName,
		HasPathParameters: useQueryOptions.HasPathParameters,
	}, ctx)

	if err != nil {
		return nil, err
	}

	claims := []core.JsFnArgument{
		{
			Key: "options.argument",
			Ts:  "options: " + useQueryOptions.ActionQueryOptionsName,
			Js:  "options",
		},
		{
			Key: "fn",
			Ts:  fn,
			Js:  fn,
		},
	}

	className := fmt.Sprintf("use%v", core.ToUpper(useQueryOptions.ActionName))
	const tmpl = `
		
export const {{ .className }}Query = (
	|@options.argument|
) => {

	|@fn|

	const result = useQuery({
		queryKey: [
			{{ .hookOptions.NewUrlFunctionName }} (
			 	{{ if .hookOptions.HasPathParameters }}
				options.params,
				{{ end }}

				options?.qs
			)
		],
		queryFn: fn,
		{{ if .hookOptions.HasPathParameters }}
		// Bug fix: a "new"/create screen has no uniqueId (or other path param) yet -
		// every *EntityManager.tsx across this app builds this action's own
		// getSingleHook unconditionally (React's rules of hooks - a hook can never be
		// called conditionally), which used to mean this fired for real, every time,
		// against a URL with a literal "undefined" in place of the missing param (e.g.
		// GET /treasury/undefined) - a wasted request and, in a route path built with
		// this before uniqueId resolves, a genuinely wrong one. Defaulting enabled to
		// false whenever any path parameter is missing/empty fixes this the same way
		// for every generated get-by-id hook at once - options.enabled below still
		// wins if a caller explicitly opts back in.
		//
		// Also checks the *string* "undefined"/"null", not just the real values -
		// a caller building params off something already coerced to text before
		// this hook ever sees it (e.g. a router param read via a JS template
		// literal, or String(value) upstream) hands this a param that's
		// technically a non-empty string, but is exactly the same "nothing to
		// fetch yet" case as the real undefined/null it stringified from.
		enabled: !Object.values(options.params || {}).some(
			(v) => v === undefined || v === null || v === "" || v === "undefined" || v === "null",
		),
		{{ end }}
		...(options || {}),
	});

	return {
		...result,
		isCompleted,
		response
	}
};


	`

	t := template.Must(template.New("jsactionoptions").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"hookOptions": useQueryOptions,
		"ctx":         ctx,
		"className":   className,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()

	claimsRendered := ClaimRender(claims, ctx)
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}
	reactQueryLocation := getReactQueryInfo(ctx)

	res := &core.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Objects:  []string{"useQuery"},
				Location: reactQueryLocation.PackageName,
			},
			{
				Objects:  []string{"useState"},
				Location: "react",
			},

			{
				Objects:  []string{"useFetchxContext"},
				Location: getSdkAwareLocation(ctx, INTERNAL_SDK_REACT_LOCATION, "useFetchx"),
			},
		},
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  TOKEN_ROOT_CLASS,
				Value: className,
			},
		},
	}

	isTypeScript := ctx.HasTag(Typescript)
	if isTypeScript {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Objects:  []string{"type TypedResponse"},
			Location: getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION, "fetchx"),
		})
	}

	res.CodeChunkDependensies = append(res.CodeChunkDependensies, extraDeps...)

	return res, nil
}
