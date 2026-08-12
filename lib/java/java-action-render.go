// Renders a full set of standalone Java files for a single action/remote:
// any path-parameter/query/request/response/header classes it needs (each
// its own public type, since Java allows only one public type per file),
// plus a ready-to-call public `<Action>Client` class with a static method
// (an `Iterator<T>` for `method: reactive`) that performs the actual HTTP
// call through the runtime Fetchx helpers.
package java

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// JavaActionRender renders the complete set of files for one action or
// remote - the last entry is always the `<Action>Client` file itself.
func JavaActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) ([]*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, err := javaGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}
	if !realms.HasUrl {
		// CLI-only action - nothing to generate for a client target, other
		// than the (already independently useful) request/response shapes.
		files := append([]*core.CodeChunkCompiled{}, realms.RequestClasses...)
		files = append(files, realms.ResponseClasses...)
		return files, nil
	}

	files := []*core.CodeChunkCompiled{}
	if realms.PathParameter != nil {
		files = append(files, realms.PathParameter)
	}
	if realms.QueryClass != nil {
		files = append(files, realms.QueryClass)
	}
	files = append(files, realms.RequestClasses...)
	files = append(files, realms.ResponseClasses...)
	if realms.RequestHeadersClass != nil {
		files = append(files, realms.RequestHeadersClass)
	}
	if realms.ResponseHeadersClass != nil {
		files = append(files, realms.ResponseHeadersClass)
	}

	hasBody := realms.RequestClassName != "" || len(realms.RequestClasses) > 0
	requestType := realms.RequestClassName
	if requestType == "" && len(realms.RequestClasses) > 0 {
		requestType = realms.ActionName + "Request"
	}

	responseType := realms.ResponseClassName
	if responseType == "" && len(realms.ResponseClasses) > 0 {
		responseType = realms.ActionName + "Response"
	}
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "Object" && !realms.ResponseIsAmbiguous

	// --- build the method parameter list -----------------------------------------------
	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("%vPathParameters path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("%v body", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("%vQueryParams query", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("%vRequestHeaders headers", realms.ActionName))
	} else {
		params = append(params, "java.util.Map<String, String> headers")
	}
	params = append(params, "FetchxContext ctx")

	returnHint := "String"
	if returnsTyped {
		returnHint = responseType
	}

	headersArg := "headers == null ? null : headers.toHttpHeaders()"
	if realms.RequestHeadersClass == nil {
		headersArg = "headers"
	}

	jsonBodyArg := "null"
	if hasBody {
		jsonBodyArg = "body"
	}

	queryArg := "null"
	if realms.QueryClass != nil {
		queryArg = "query == null ? null : query.toQueryParams()"
	}

	httpMethod := realms.HttpMethod
	if httpMethod == "" {
		if hasBody {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	const tmpl = `/** Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }}) */
public class {{ .realms.ActionName }}Client {
{{ if .realms.IsReactive }}
    public static java.util.Iterator<{{ .returnHint }}> {{ .realms.FunctionName }}(
        {{ .paramList }}) {
        var url = (ctx != null ? ctx : Fetchx.DEFAULT_CONTEXT).buildUrl({{ .realms.Url | javaString }});
{{ if .realms.PathParameter }}        var resolvedUrl = path.apply(url);
{{ else }}        var resolvedUrl = url;
{{ end }}
        var events = Fetchx.streamEvents(
            {{ .method | javaString }},
            resolvedUrl,
            {{ .queryArg }},
            {{ .jsonBodyArg }},
            {{ .headersArg }},
            ctx);

        return new java.util.Iterator<{{ .returnHint }}>() {
            @Override
            public boolean hasNext() {
                return events.hasNext();
            }

            @Override
            public {{ .returnHint }} next() {
                var raw = events.next();
                return {{ if .returnsTyped }}Fetchx.deserialize(raw, {{ .responseType }}.class){{ else }}raw{{ end }};
            }
        };
    }
{{ else }}
    public static {{ .returnHint }} {{ .realms.FunctionName }}(
        {{ .paramList }}) {
        var url = (ctx != null ? ctx : Fetchx.DEFAULT_CONTEXT).buildUrl({{ .realms.Url | javaString }});
{{ if .realms.PathParameter }}        var resolvedUrl = path.apply(url);
{{ else }}        var resolvedUrl = url;
{{ end }}
        var response = Fetchx.request(
            {{ .method | javaString }},
            resolvedUrl,
            {{ .queryArg }},
            {{ .jsonBodyArg }},
            {{ .headersArg }},
            ctx);
{{ if .returnsTyped }}        return Fetchx.deserialize(response, {{ .responseType }}.class);
{{ else }}        return response;
{{ end }}
    }
{{ end }}
}
`

	funcs := template.FuncMap{"javaString": escapeDoubleQuoted}
	t := template.Must(template.New("javaaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"paramList":    strings.Join(params, ", "),
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

	clientChunk := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.ActionName + "Client"},
		},
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  realms.ActionName + "Client",
		SuggestedExtension: ".java",
		Realms:             realms,
	}
	files = append(files, clientChunk)

	return files, nil
}
