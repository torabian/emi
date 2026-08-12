// Renders a full, standalone dart file for a single action/remote: any
// path-parameter/query/request/response/header classes it needs, plus a
// ready-to-call client function (a `Future<T>`, or a `Stream<T>` for
// `method: reactive`) that performs the actual HTTP call through the runtime
// fetchx helpers.
package dart

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// DartActionRender renders the complete file for one action or remote.
func DartActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, deps, err := dartGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.ActionName},
		},
		CodeChunkDependensies: deps,
	}

	if realms.HasUrl {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: "runtime/fetchx.dart"})
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
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "dynamic" && !realms.ResponseIsAmbiguous

	// --- build the named-parameter signature ------------------------------------------
	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("required %vPathParameters path,", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("required %v body,", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("%vQueryParams? query,", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("%vRequestHeaders? headers,", realms.ActionName))
	} else {
		params = append(params, "Map<String, String>? headers,")
	}
	params = append(params, "FetchxContext? ctx,")

	returnHint := "dynamic"
	if returnsTyped {
		returnHint = responseType
	}

	headersArg := "headers?.toHttpHeaders()"
	if realms.RequestHeadersClass == nil {
		headersArg = "headers"
	}

	jsonBodyArg := "null"
	if hasBody {
		jsonBodyArg = "body.toJson()"
	}

	queryArg := "null"
	if realms.QueryClass != nil {
		queryArg = "query?.toQueryParams()"
	}

	const tmpl = `/// Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }})

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
{{ if .realms.IsReactive }}
Stream<{{ .returnHint }}> {{ .realms.FunctionName }}({
{{- range .params }}
  {{ . }}
{{- end }}
}) async* {
  final url = (ctx ?? defaultContext).buildUrl({{ .realms.Url | dartString }});
{{ if .realms.PathParameter }}  final resolvedUrl = path.apply(url);
{{ else }}  final resolvedUrl = url;
{{ end }}
  await for (final event in streamEvents(
    {{ .method | dartString }},
    resolvedUrl,
    params: {{ .queryArg }},
    jsonBody: {{ .jsonBodyArg }},
    headers: {{ .headersArg }},
    ctx: ctx,
  )) {
    yield {{ if .returnsTyped }}{{ .responseType }}.fromJson(event as Map<String, dynamic>){{ else }}event{{ end }};
  }
}
{{ else }}
Future<{{ .returnHint }}> {{ .realms.FunctionName }}({
{{- range .params }}
  {{ . }}
{{- end }}
}) async {
  final url = (ctx ?? defaultContext).buildUrl({{ .realms.Url | dartString }});
{{ if .realms.PathParameter }}  final resolvedUrl = path.apply(url);
{{ else }}  final resolvedUrl = url;
{{ end }}
  final response = await request(
    {{ .method | dartString }},
    resolvedUrl,
    params: {{ .queryArg }},
    jsonBody: {{ .jsonBodyArg }},
    headers: {{ .headersArg }},
    ctx: ctx,
  );
{{ if .returnsTyped }}  return {{ .responseType }}.fromJson(response as Map<String, dynamic>);
{{ else }}  return response;
{{ end }}
}
{{ end }}
{{ end }}
`

	httpMethod := realms.HttpMethod
	if httpMethod == "" {
		if hasBody {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	funcs := template.FuncMap{"dartString": escapeDoubleQuoted}
	t := template.Must(template.New("dartaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"params":       params,
		"returnHint":   returnHint,
		"method":       strings.ToUpper(httpMethod),
		"queryArg":     queryArg,
		"jsonBodyArg":  jsonBodyArg,
		"headersArg":   headersArg,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(realms.ActionName)
	res.SuggestedExtension = ".dart"
	res.Realms = realms

	return res, nil
}
