// Renders a full, standalone generic-dialect C++ header for a single
// action/remote: any path-parameter/query/request/response/header classes it
// needs, plus a ready-to-call static client method - a blocking call returning
// std::optional<T> for `method: get/post/...`, or a typed emi::EmiWebSocketX
// factory for `method: reactive` (see cpp-action-render-reactive.go).
package cpp

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func cppGenericActionRender(realms cppActionRealms, deps []core.CodeChunkDependency) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.ActionName},
		},
		CodeChunkDependensies: deps,
	}

	if realms.IsReactive {
		return cppGenericReactiveActionRender(realms, res)
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
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "emi::EmiJson" && !realms.ResponseIsAmbiguous

	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("const %vPathParameters& path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("const %v& body", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("const %vQueryParams& query", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("const %vRequestHeaders& headers", realms.ActionName))
	}
	params = append(params, "emi::EmiHttpResponse* outResponse = nullptr")
	params = append(params, "emi::EmiClientConfig& config = emi::EmiClientConfig::Default()")

	returnType := "std::optional<std::string>"
	if returnsTyped {
		returnType = fmt.Sprintf("std::optional<%v>", responseType)
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
// Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }})
class {{ .realms.ActionName }}Client {
public:
    static {{ .returnType }} {{ .realms.FunctionName }}({{ .paramList }}) {
        std::string url = emi::EmiBuildUrl(config.baseUrl, {{ .realms.Url | cppString }});
{{ if .realms.PathParameter }}        url = path.Apply(url);
{{ end }}
{{ if .realms.QueryClass }}        url = query.Apply(url);
{{ end }}
        emi::EmiHttpRequest request;
        request.method = {{ .method | cppString }};
        request.url = url;
        request.headers = config.defaultHeaders;
{{ if .realms.RequestHeadersClass }}        {
            auto extra = headers.ToHeaders();
            request.headers.insert(request.headers.end(), extra.begin(), extra.end());
        }
{{ end }}
{{ if .hasBody }}        {
            cJSON* bodyJson = body.ToJson();
            char* bodyStr = cJSON_PrintUnformatted(bodyJson);
            request.body = bodyStr ? bodyStr : "";
            cJSON_free(bodyStr);
            cJSON_Delete(bodyJson);
        }
{{ end }}
        if (!config.transport) {
            emi::EmiHttpResponse failure;
            failure.ok = false;
            failure.error = "emi::EmiClientConfig::transport is not set";
            if (outResponse) *outResponse = failure;
            return std::nullopt;
        }

        emi::EmiHttpResponse response = config.transport->Send(request);
        if (outResponse) *outResponse = response;
        if (!response.ok || response.statusCode >= 400) {
            return std::nullopt;
        }
{{ if .returnsTyped }}        return {{ .responseType }}::Parse(response.body);
{{ else }}        return response.body;
{{ end }}
    }
};
{{ end }}
`

	funcs := template.FuncMap{"cppString": escapeDoubleQuoted}
	t := template.Must(template.New("cppgenericaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"paramList":    strings.Join(params, ", "),
		"returnType":   returnType,
		"method":       strings.ToUpper(httpMethod),
		"hasBody":      hasBody,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".hpp"
	res.Realms = realms

	return res, nil
}
