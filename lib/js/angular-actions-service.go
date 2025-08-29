package js

// In Angular, on top of the existing pure javascript for each action,
// We need to create a service which would combine

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// represents each function in the service in Angular, which will be callable to make a
// action (http/socket/etc), we first compute that information
type angularServiceActionItem struct {
	FunctionSignature string

	FetchMetaClassName string

	ResponseClassGenerator string

	AngularGenericReturnType string

	HttpCallerParams string

	UrlGeneratorFn string

	// the function on angular http client service will be called, such this.http.post,get
	AngularHttpMethodFunction string
}

func angularRequestTypes(realms *jsActionRealms) string {
	claims := []string{}
	if realms.QueryStringClass != nil {
		claims = append(
			claims,
			fmt.Sprintf("params?: %v", findTokenByName(realms.QueryStringClass.Tokens, TOKEN_ROOT_CLASS).Value),
		)
	}

	if realms.RequestHeadersClass != nil {
		claims = append(
			claims,
			fmt.Sprintf("headers?: %v", findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS).Value),
		)
	}

	return strings.Join(claims, ",")
}

func angularServiceGetArguments(realms *jsActionRealms) string {
	claims := []string{}

	if realms.PathParameter != nil {
		claims = append(
			claims,
			fmt.Sprintf("params: %v", findTokenByName(realms.PathParameter.Tokens, TOKEN_ROOT_CLASS).Value),
		)
	}

	claims = append(
		claims,
		fmt.Sprintf(`options?: Parameters<HttpClient['get']>[1] & { %v }`, angularRequestTypes(realms)),
	)

	return strings.Join(claims, ", ")
}

func angularHttpCallerParams(realms *jsActionRealms) string {
	claims := []string{}

	if realms.PathParameter != nil {
		claims = append(claims, "params")
	}

	return strings.Join(claims, ", ")
}

// Somehow import from generated files
func importPath(fileName string) string {
	return "./" + fileName
}

// Combines multiple actions
func AngularActionsClass(module *core.Module3, actionsRendered []*core.CodeChunkCompiled, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}
	const tmpl = `/**
* Angular service for actions
*/

@Injectable({ providedIn: 'root' })
export class {{ .className }} {

	constructor(
		private http: HttpClient,

		{{ range .angularServiceActionItems }}
		@Optional()
		@Inject({{ .ResponseClassGenerator }}Factory)
		private {{ .ResponseClassGenerator }}Factory: {{ .ResponseClassGenerator }}Factory
		{{ end }}
	) {}

	{{ range .angularServiceActionItems }}
		{{ .FunctionSignature }} {
			return this.http.{{ .AngularHttpMethodFunction }}<{{.AngularGenericReturnType}}>(
				{{ .UrlGeneratorFn }}(
					{{ .HttpCallerParams }}
				),
				options
			)
			{{ if .ResponseClassGenerator }}
			// Create instance of {{ .ResponseClassGenerator }}
			{{ end }}
		}
	{{end }}
}

`

	angularServiceActionItems := []angularServiceActionItem{}
	// compute the actions as functions to be placed inside the service
	for _, action := range actionsRendered {
		realms := action.Realms.(*jsActionRealms)
		fetchMetaClassName := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS)
		urlGenerator := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_NEW_URL_FN)
		method := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_ACTUAL_METHOD)

		// Import the path paramter type into the class
		if realms.PathParameter != nil {
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					findTokenByName(realms.PathParameter.Tokens, TOKEN_ROOT_CLASS).Value,
				},
				Location: importPath(action.SuggestedFileName),
			})
		}

		// Import if there is a querystring class
		if realms.QueryStringClass != nil {
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					findTokenByName(realms.QueryStringClass.Tokens, TOKEN_ROOT_CLASS).Value,
				},
				Location: importPath(action.SuggestedFileName),
			})
		}

		// Import if there is a request header
		if realms.RequestHeadersClass != nil {
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS).Value,
				},
				Location: importPath(action.SuggestedFileName),
			})
		}

		// Create a class out of the response
		responseClassGenerator := ""

		angularGenericReturnType := "unknown"
		if realms.ResponseClass != nil {
			angularGenericReturnType = findTokenByName(realms.ResponseClass.Tokens, TOKEN_ROOT_CLASS).Value
			responseClassGenerator = angularGenericReturnType

			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					angularGenericReturnType,
				},
				Location: importPath(action.SuggestedFileName),
			})

			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					fmt.Sprintf("%vFactory", responseClassGenerator),
				},
				Location: importPath(action.SuggestedFileName),
			})

			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					"map",
				},
				Location: "rxjs",
			})
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					"Optional",
					"Inject",
				},
				Location: "@angular/core",
			})

			// Here maybe if it's a type, we need to skip it.
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
			Objects: []string{
				fetchMetaClassName.Value,
			},
			Location: importPath(action.SuggestedFileName),
		})

		funcSignature := fmt.Sprintf(
			"%v(%v)",
			findTokenByName(action.Tokens, TOKEN_ORIGINAL_NAME).Value,
			angularServiceGetArguments(realms),
		)

		angularServiceActionItems = append(angularServiceActionItems, angularServiceActionItem{
			FunctionSignature:         funcSignature,
			UrlGeneratorFn:            urlGenerator.Value,
			AngularGenericReturnType:  angularGenericReturnType,
			ResponseClassGenerator:    responseClassGenerator,
			HttpCallerParams:          angularHttpCallerParams(realms),
			FetchMetaClassName:        fetchMetaClassName.Value,
			AngularHttpMethodFunction: method.Value,
		})

	}

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	className := fmt.Sprintf("%vActionsService", core.ToUpper(module.Name))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"angularServiceActionItems": angularServiceActionItems,
		"className":                 className,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = className
	res.SuggestedExtension = ".ts"
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, getAngularServiceImports()...)

	return res, nil
}
