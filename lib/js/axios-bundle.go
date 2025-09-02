package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type axiosBundleActionItem struct {
	FunctionSignature      string
	ResponseClassGenerator string
	AxiosRequestGeneric    string
	HttpCallerParams       string
	UrlMethodConst         string
	UrlGeneratorFn         string
}

// Creates a axios object out of the actions, by extending the Axios class
func AxiosBundleClass(module *core.Emi, actionsRendered []*core.CodeChunkCompiled, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}
	const tmpl = `/**
* Axios bundle service
*/
import type {
  AxiosResponseHeaders,
  AxiosRequestConfig as BaseAxiosRequestConfig,
  RawAxiosResponseHeaders,
} from "axios";

export interface TypedAxiosRequestConfig<D = unknown, P = unknown, H = unknown>
  extends Omit<BaseAxiosRequestConfig, "data" | "params" | "headers"> {
  data?: D;
  params?: P;
  headers?: H;
}

export interface TypedAxiosResponse<
  T = unknown,
  D = unknown,
  H = RawAxiosResponseHeaders | AxiosResponseHeaders
> extends Omit<AxiosResponse<T, D>, "headers"> {
  headers: H;
}

export class {{ .className }} extends Axios {
  static create(config?: AxiosRequestConfig) {
    return new {{ .className }}(config);
  }

  {{ range .items }}
		{{ .FunctionSignature }} {
		 	const url = {{ .UrlGeneratorFn }}(
				{{ .HttpCallerParams }}
			)
			return this.request<{{ .AxiosRequestGeneric }}>(
				{
					url,
					{{ if .UrlMethodConst }}
					method: {{ .UrlMethodConst }},
					{{ end }}
					...(config || {})
				} as never
			)
				
			{{ if .ResponseClassGenerator }}
				.then(res => {
					res.data = new {{ .ResponseClassGenerator }}(res);

					return res;
				})
			{{ end }}
		}
  {{ end }}
}

`

	axiosBundleActionItems := []axiosBundleActionItem{}

	for _, action := range actionsRendered {
		realms := action.Realms.(*jsActionRealms)
		fetchMetaClassName := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_ROOT_CLASS)
		urlGenerator := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_NEW_URL_FN)
		urlMethod := findTokenByName(realms.FetchMetaClass.Tokens, TOKEN_URL_METHOD)

		// Axios basically doesn't support the emi reactive or sse.
		if realms.HttpMethod == EMI_METHOD_REACTIVE || realms.HttpMethod == EMI_METHOD_SSE {
			continue
		}
		// Import the path paramter type into the class
		if realms.PathParameter != nil {
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					"type " + findTokenByName(realms.PathParameter.Tokens, TOKEN_ROOT_CLASS).Value,
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

		requestHeaderType := "unknown"
		// Import if there is a request header
		if realms.RequestHeadersClass != nil {
			requestHeaderType = findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS).Value
			res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
				Objects: []string{
					requestHeaderType,
				},
				Location: importPath(action.SuggestedFileName),
			})
		}

		responseClassType := "unknown"
		if realms.ResponseClass != nil {
			responseClass := findTokenByName(realms.ResponseClass.Tokens, TOKEN_ROOT_CLASS)

			if responseClass != nil {
				responseClassType = responseClass.Value
				res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
					Objects: []string{
						responseClassType,
					},
					Location: importPath(action.SuggestedFileName),
				})
			}

		}

		// Create a class out of the response
		responseClassGenerator := ""
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, core.CodeChunkDependency{
			Objects: []string{
				fetchMetaClassName.Value,
			},
			Location: importPath(action.SuggestedFileName),
		})

		funcSignature := fmt.Sprintf(
			"%v(%v)",
			findTokenByName(action.Tokens, TOKEN_ORIGINAL_NAME).Value,
			axiosBundleItemGetArguments(realms),
		)

		axiosRequestGeneric := fmt.Sprintf(
			`%v, TypedAxiosResponse<%v, unknown, %v>`,
			responseClassType,
			responseClassType,
			requestHeaderType,
		)

		axiosBundleActionItems = append(axiosBundleActionItems, axiosBundleActionItem{
			FunctionSignature:      funcSignature,
			UrlGeneratorFn:         urlGenerator.Value,
			UrlMethodConst:         urlMethod.Value,
			ResponseClassGenerator: responseClassGenerator,
			HttpCallerParams:       axiosHttpCallerParams(realms),
			AxiosRequestGeneric:    axiosRequestGeneric,
		})

	}

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	className := fmt.Sprintf("%vAxiosClient", core.ToUpper(module.Name))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"className": className,
		"items":     axiosBundleActionItems,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = className
	res.SuggestedExtension = ".ts"
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, getAxiosBundleDependencies()...)

	return res, nil
}

func getAxiosBundleDependencies() []core.CodeChunkDependency {
	deps := []core.CodeChunkDependency{
		{
			Objects: []string{
				"Axios",
				"type AxiosRequestConfig", "type AxiosResponse",
			},
			Location: "axios",
		},
	}
	return deps
}

func axiosBundleItemGetArguments(realms *jsActionRealms) string {
	claims := []string{}

	requestHeaderType := "unknown"
	if realms.RequestHeadersClass != nil {
		requestHeaderType = findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS).Value

	}

	requestClass := "unknown"
	if realms.RequestClass != nil {
		requestClass = findTokenByName(realms.RequestClass.Tokens, TOKEN_ROOT_CLASS).Value
	}

	queryParams := "unknown"
	if realms.QueryStringClass != nil {
		queryParams = findTokenByName(realms.QueryStringClass.Tokens, TOKEN_ROOT_CLASS).Value
	}

	if realms.PathParameter != nil {
		claims = append(
			claims,
			fmt.Sprintf("params: %v", findTokenByName(realms.PathParameter.Tokens, TOKEN_ROOT_CLASS).Value),
		)
	}

	claims = append(
		claims,
		fmt.Sprintf(`config?: TypedAxiosRequestConfig<
			%v,
			%v,
			%v
		>`, requestClass, queryParams, requestHeaderType),
	)

	return strings.Join(claims, ", ")
}

func axiosHttpCallerParams(realms *jsActionRealms) string {
	claims := []string{}

	if realms.PathParameter != nil {
		claims = append(claims, "params")
	}

	return strings.Join(claims, ", ")
}

func importPath(fileName string) string {
	return "./" + fileName
}
