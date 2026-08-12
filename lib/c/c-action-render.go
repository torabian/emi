// Renders a full, standalone C header for a single action/remote: any
// path-parameter/query/request/response/header structs it needs, plus a
// ready-to-call `static inline` function (returning a `FetchxEventList` for
// `method: reactive`) that performs the actual HTTP call through the
// runtime fetchx helpers.
package c

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// CActionRender renders the complete file for one action or remote.
func CActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, err := cGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.FunctionName},
		},
	}

	if realms.HasUrl {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, core.CodeChunkDependency{Location: "fetchx.h"})
	}
	if realms.RequestStruct != nil {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, realms.RequestStruct.CodeChunkDependensies...)
	} else if realms.RequestTypeName != "" {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, HeaderIncludeDependency(realms.RequestTypeName)...)
	}
	if realms.ResponseStruct != nil {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, realms.ResponseStruct.CodeChunkDependensies...)
	} else if realms.ResponseTypeName != "" {
		res.CodeChunkDependensies = append(res.CodeChunkDependensies, HeaderIncludeDependency(realms.ResponseTypeName)...)
	}

	hasBody := realms.RequestTypeName != "" || realms.RequestStruct != nil
	requestType := realms.RequestTypeName
	if requestType == "" && realms.RequestStruct != nil {
		requestType = realms.ActionName + "Request"
	}

	responseType := realms.ResponseTypeName
	if responseType == "" && realms.ResponseStruct != nil {
		responseType = realms.ActionName + "Response"
	}
	returnsTyped := realms.HasResponseShape && responseType != "" && !realms.ResponseIsAmbiguous

	params := []string{"const FetchxContext* ctx"}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("const %vPathParameters* path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("const %v* body", requestType))
	}
	if realms.QueryStruct != nil {
		params = append(params, fmt.Sprintf("const %vQueryParams* query", realms.ActionName))
	}
	if realms.RequestHeadersStruct != nil {
		params = append(params, fmt.Sprintf("const %vRequestHeaders* headers", realms.ActionName))
	}
	if !realms.IsReactive {
		params = append(params, "FetchxResponse* out_response")
	}

	returnHint := "char*"
	if returnsTyped {
		returnHint = responseType + "*"
	}

	httpMethod := realms.HttpMethod
	if httpMethod == "" {
		if hasBody {
			httpMethod = "POST"
		} else {
			httpMethod = "GET"
		}
	}

	const tmpl = `
{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
{{ if .realms.QueryStruct }}
{{ b2s .realms.QueryStruct.ActualScript }}
{{ end }}
{{ if .realms.RequestStruct }}
{{ b2s .realms.RequestStruct.ActualScript }}
{{ end }}
{{ if .realms.ResponseStruct }}
{{ b2s .realms.ResponseStruct.ActualScript }}
{{ end }}
{{ if .realms.RequestHeadersStruct }}
{{ b2s .realms.RequestHeadersStruct.ActualScript }}
{{ end }}
{{ if .realms.HasUrl }}
/* Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }}) */
{{ if .realms.IsReactive }}
static inline FetchxEventList {{ .realms.FunctionName }}({{ .paramList }}) {
    char* base = fetchx_build_url(ctx ? ctx->base_url : "", {{ .realms.Url | cString }});
{{ if .realms.PathParameter }}    char* url1 = {{ .realms.ActionName }}PathParameters_apply(path, base);
    free(base);
    base = url1;
{{ end }}
{{ if .realms.QueryStruct }}    char* url2 = {{ .realms.ActionName }}QueryParams_apply(query, base);
    free(base);
    base = url2;
{{ end }}
    struct curl_slist* header_list = NULL;
{{ if .realms.RequestHeadersStruct }}    header_list = {{ .realms.ActionName }}RequestHeaders_to_curl_slist(headers, header_list);
{{ end }}
{{ if .hasBody }}    cJSON* body_json = {{ .requestType }}_to_json(body);
    char* body_str = cJSON_PrintUnformatted(body_json);
{{ else }}    char* body_str = NULL;
{{ end }}
    FetchxEventList events = fetchx_stream_events({{ .method | cString }}, base, body_str, header_list, ctx, NULL);

    free(base);
    curl_slist_free_all(header_list);
{{ if .hasBody }}    cJSON_Delete(body_json);
    free(body_str);
{{ end }}
    return events;
}

{{ if .returnsTyped }}
/* Parses one event yielded by {{ .realms.FunctionName }}() - the caller owns the result and must {{ .responseType }}_free() it. */
static inline {{ .responseType }}* {{ .realms.FunctionName }}_parse_event(const char* json_str) {
    cJSON* root = cJSON_Parse(json_str);
    {{ .responseType }}* result = {{ .responseType }}_from_json(root);
    cJSON_Delete(root);
    return result;
}
{{ end }}

{{ else }}
static inline {{ .returnHint }} {{ .realms.FunctionName }}({{ .paramList }}) {
    char* base = fetchx_build_url(ctx ? ctx->base_url : "", {{ .realms.Url | cString }});
{{ if .realms.PathParameter }}    char* url1 = {{ .realms.ActionName }}PathParameters_apply(path, base);
    free(base);
    base = url1;
{{ end }}
{{ if .realms.QueryStruct }}    char* url2 = {{ .realms.ActionName }}QueryParams_apply(query, base);
    free(base);
    base = url2;
{{ end }}
    struct curl_slist* header_list = NULL;
{{ if .realms.RequestHeadersStruct }}    header_list = {{ .realms.ActionName }}RequestHeaders_to_curl_slist(headers, header_list);
{{ end }}
{{ if .hasBody }}    cJSON* body_json = {{ .requestType }}_to_json(body);
    char* body_str = cJSON_PrintUnformatted(body_json);
{{ else }}    char* body_str = NULL;
{{ end }}
    FetchxResponse response = fetchx_request({{ .method | cString }}, base, body_str, header_list, ctx);

    free(base);
    curl_slist_free_all(header_list);
{{ if .hasBody }}    cJSON_Delete(body_json);
    free(body_str);
{{ end }}

    {{ .returnHint }} result = NULL;
    if (response.ok && response.status_code < 400) {
{{ if .returnsTyped }}        cJSON* root = cJSON_Parse(response.body);
        result = {{ .responseType }}_from_json(root);
        cJSON_Delete(root);
{{ else }}        result = response.body ? strdup(response.body) : NULL;
{{ end }}
    }

    if (out_response) { *out_response = response; } else { FetchxResponse_free(&response); }
    return result;
}
{{ end }}
{{ end }}
`

	funcs := template.FuncMap{"cString": escapeDoubleQuoted}
	t := template.Must(template.New("caction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"paramList":    strings.Join(params, ", "),
		"returnHint":   returnHint,
		"method":       strings.ToUpper(httpMethod),
		"hasBody":      hasBody,
		"requestType":  requestType,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToSnakeCase(realms.ActionName)
	res.SuggestedExtension = ".h"
	res.Realms = realms

	return res, nil
}
