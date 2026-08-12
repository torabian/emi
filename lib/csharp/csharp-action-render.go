// Renders a full, standalone C# file for a single action/remote: any
// path-parameter/query/request/response/header classes it needs, plus a
// ready-to-call static client method (a `Task<T>`, or an
// `IAsyncEnumerable<T>` for `method: reactive`) that performs the actual
// HTTP call through the runtime Fetchx helpers.
package csharp

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// CSharpActionRender renders the complete file for one action or remote.
func CSharpActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, deps, err := csGetActionRealms(action, ctx, complexes)
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

	hasBody := realms.RequestClassName != "" || realms.RequestClass != nil
	requestType := realms.RequestClassName
	if requestType == "" && realms.RequestClass != nil {
		requestType = realms.ActionName + "Request"
	}

	responseType := realms.ResponseClassName
	if responseType == "" && realms.ResponseClass != nil {
		responseType = realms.ActionName + "Response"
	}
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "object" && !realms.ResponseIsAmbiguous

	// --- build the method parameter list -----------------------------------------------
	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("%vPathParameters path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("%v body", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("%vQueryParams? query = null", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("%vRequestHeaders? headers = null", realms.ActionName))
	} else {
		params = append(params, "Dictionary<string, string>? headers = null")
	}
	params = append(params, "FetchxContext? ctx = null")

	returnHint := "string?"
	if returnsTyped {
		returnHint = responseType + "?"
	}

	headersArg := "headers?.ToHttpHeaders()"
	if realms.RequestHeadersClass == nil {
		headersArg = "headers"
	}

	jsonBodyArg := "null"
	if hasBody {
		jsonBodyArg = "body"
	}

	queryArg := "null"
	if realms.QueryClass != nil {
		queryArg = "query?.ToQueryParams()"
	}

	httpMethod := realms.HttpMethod
	if httpMethod == "" {
		if hasBody {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	const tmpl = `/// <summary>Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }})</summary>

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
public static class {{ .realms.ActionName }}Client
{
{{ if .realms.IsReactive }}
    public static async IAsyncEnumerable<{{ .returnHint }}> {{ .realms.FunctionName }}(
{{- range .params }}
        {{ . }},
{{- end }}
        [System.Runtime.CompilerServices.EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        var url = (ctx ?? Fetchx.DefaultContext).BuildUrl({{ .realms.Url | csString }});
{{ if .realms.PathParameter }}        var resolvedUrl = path.Apply(url);
{{ else }}        var resolvedUrl = url;
{{ end }}
        await foreach (var evt in Fetchx.StreamEvents(
            {{ .method | csString }},
            resolvedUrl,
            {{ .queryArg }},
            {{ .jsonBodyArg }},
            {{ .headersArg }},
            ctx))
        {
            yield return {{ if .returnsTyped }}JsonSerializer.Deserialize<{{ .responseType }}>(evt){{ else }}evt{{ end }};
        }
    }
{{ else }}
    public static async Task<{{ .returnHint }}> {{ .realms.FunctionName }}(
        {{ .paramList }})
    {
        var url = (ctx ?? Fetchx.DefaultContext).BuildUrl({{ .realms.Url | csString }});
{{ if .realms.PathParameter }}        var resolvedUrl = path.Apply(url);
{{ else }}        var resolvedUrl = url;
{{ end }}
        var response = await Fetchx.Request(
            {{ .method | csString }},
            resolvedUrl,
            {{ .queryArg }},
            {{ .jsonBodyArg }},
            {{ .headersArg }},
            ctx);
{{ if .returnsTyped }}        if (response == null) return default;
        return JsonSerializer.Deserialize<{{ .responseType }}>(response);
{{ else }}        return response;
{{ end }}
    }
{{ end }}
}
{{ end }}
`

	funcs := template.FuncMap{"csString": escapeDoubleQuoted}
	t := template.Must(template.New("csaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"params":       params,
		"paramList":    strings.Join(params, ",\n        "),
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
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".cs"
	res.Realms = realms

	return res, nil
}
