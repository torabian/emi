// Renders a full, standalone python module for a single action/remote: any
// path-parameter/query/request/response/header classes it needs, plus a
// ready-to-call client function (or async generator, for `method: reactive`)
// that performs the actual HTTP call through the runtime fetchx helpers.
//
// Mirrors lib/kotlin/kotlin-action-render.go and lib/js/js-action-manifest.go,
// scoped down to what a plain http client needs (no react-query style hooks).
package python

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// PythonActionRender renders the complete file for one action or remote.
func PythonActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, deps, err := pyGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	isAsync := ctx.HasTag(Async)

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.ActionName},
		},
		CodeChunkDependensies: deps,
	}

	requestFn, streamFn := "request", "stream_events"
	if isAsync {
		requestFn, streamFn = "request_async", "stream_events_async"
	}

	hasBody := realms.RequestClassName != "" || realms.RequestClass != nil
	requestType := realms.RequestClassName
	if requestType == "" && realms.RequestClass != nil {
		requestType = realms.ActionName + "Request"
	}

	responseType := realms.ResponseClassName
	if responseType == "" && realms.ResponseClass != nil {
		responseType = realms.ActionName + "Response"
	}
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "Any" && !realms.ResponseIsAmbiguous

	if realms.HasUrl {
		runtimeObjects := []string{"FetchxContext", "DEFAULT_CONTEXT"}
		if realms.IsReactive {
			runtimeObjects = append(runtimeObjects, streamFn)
		} else {
			runtimeObjects = append(runtimeObjects, requestFn)
		}
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
			Location: ".runtime.fetchx",
			Objects:  runtimeObjects,
		})

		serializationObjects := []string{}
		if hasBody {
			serializationObjects = append(serializationObjects, "to_dict")
		}
		if returnsTyped {
			serializationObjects = append(serializationObjects, "from_dict")
		}
		if len(serializationObjects) > 0 {
			res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
				Location: ".runtime.serialization",
				Objects:  serializationObjects,
			})
		}

		if realms.IsReactive {
			iterKind := "Iterator"
			if isAsync {
				iterKind = "AsyncIterator"
			}
			res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{
				Location: "typing",
				Objects:  []string{iterKind},
			})
		}
	}

	// A missing `method:` defaults to GET, unless the action carries a
	// request body, in which case POST is the more sensible guess -
	// mirrors what most REST frameworks infer when it's left unset.
	httpMethod := realms.HttpMethod
	if httpMethod == "" {
		if hasBody {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	// --- build the function signature -------------------------------------------------
	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("path: %vPathParameters", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("body: %v", requestType))
	}
	params = append(params, "*")
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("query: Optional[%vQueryParams] = None", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("headers: Optional[%vRequestHeaders] = None", realms.ActionName))
	} else {
		params = append(params, "headers: Optional[Dict[str, str]] = None")
	}
	params = append(params, "ctx: Optional[FetchxContext] = None")

	def := "def"
	if isAsync {
		def = "async def"
	}

	returnHint := "Any"
	if returnsTyped {
		returnHint = responseType
	}
	if realms.IsReactive {
		iterKind := "Iterator"
		if isAsync {
			iterKind = "AsyncIterator"
		}
		returnHint = fmt.Sprintf("%v[%v]", iterKind, returnHint)
	}

	headersArg := "headers.to_http_headers() if headers is not None else None"
	if realms.RequestHeadersClass == nil {
		headersArg = "headers"
	}

	jsonBodyArg := "None"
	if hasBody {
		jsonBodyArg = "to_dict(body)"
	}

	queryArg := "None"
	if realms.QueryClass != nil {
		queryArg = "query.to_query_params() if query is not None else None"
	}

	await := ""
	if isAsync {
		await = "await "
	}

	const tmpl = `"""
Action to communicate with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }})
"""

{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
{{ if .realms.QueryClass }}
{{ b2s .realms.QueryClass.ActualScript }}
{{ end }}
{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}
{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}
{{ if .realms.RequestHeadersClass }}
{{ b2s .realms.RequestHeadersClass.ActualScript }}
{{ end }}
{{ if .realms.ResponseHeadersClass }}
{{ b2s .realms.ResponseHeadersClass.ActualScript }}
{{ end }}
{{ if .realms.HasUrl }}
{{ .def }} {{ .realms.FunctionName }}(
{{- range .params }}
    {{ . }},
{{- end }}
) -> {{ .returnHint }}:
{{ if .realms.IsReactive }}
    url = (ctx or DEFAULT_CONTEXT).build_url("{{ .realms.Url }}")
{{ if .realms.PathParameter }}    url = path.apply(url)
{{ end }}
    {{ if .isAsync }}async {{ end }}for event in {{ .streamFn }}(
        "{{ .method }}",
        url,
        params={{ .queryArg }},
        json_body={{ .jsonBodyArg }},
        headers={{ .headersArg }},
        ctx=ctx,
    ):
        yield {{ if .returnsTyped }}from_dict({{ .responseType }}, event){{ else }}event{{ end }}
{{ else }}
    url = (ctx or DEFAULT_CONTEXT).build_url("{{ .realms.Url }}")
{{ if .realms.PathParameter }}    url = path.apply(url)
{{ end }}
    response = {{ .await }}{{ .requestFn }}(
        "{{ .method }}",
        url,
        params={{ .queryArg }},
        json_body={{ .jsonBodyArg }},
        headers={{ .headersArg }},
        ctx=ctx,
    )
{{ if .returnsTyped }}
    return from_dict({{ .responseType }}, response)
{{ else }}
    return response
{{ end }}
{{ end }}
{{ end }}
`

	t := template.Must(template.New("pyaction").Funcs(core.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"def":          def,
		"params":       params,
		"returnHint":   returnHint,
		"isAsync":      isAsync,
		"requestFn":    requestFn,
		"streamFn":     streamFn,
		"method":       strings.ToUpper(httpMethod),
		"queryArg":     queryArg,
		"jsonBodyArg":  jsonBodyArg,
		"headersArg":   headersArg,
		"await":        await,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.FunctionName
	res.SuggestedExtension = ".py"
	res.Realms = realms

	return res, nil
}
