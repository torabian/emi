// Renders a full, standalone Unreal-dialect C++ header for a single
// action/remote: any path-parameter/query/request/response/header classes it
// needs, plus a ready-to-call static client method - async (a completion
// callback, the idiomatic Unreal shape for anything I/O-bound) for
// `method: get/post/...`, or a typed FEmiWebSocketX factory for
// `method: reactive` (see cpp-action-render-reactive-unreal.go).
package cpp

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func cppUnrealActionRender(realms cppActionRealms, deps []core.CodeChunkDependency) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.ActionName},
		},
		CodeChunkDependensies: deps,
	}

	if realms.IsReactive {
		return cppUnrealReactiveActionRender(realms, res)
	}

	hasBody := realms.RequestClassName != "" || realms.RequestClass != nil
	requestType := realms.RequestClassName
	if requestType == "" && realms.RequestClass != nil {
		requestType = ueStructName(realms.ActionName + "Request")
	}

	responseType := realms.ResponseClassName
	if responseType == "" && realms.ResponseClass != nil {
		responseType = ueStructName(realms.ActionName + "Response")
	}
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "TSharedPtr<FJsonValue>" && !realms.ResponseIsAmbiguous

	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("const F%vPathParameters& Path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("const %v& Body", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("const F%vQueryParams& Query", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("const %vRequestHeaders& Headers", realms.ActionName))
	}

	completionArg := "bool bSucceeded, int32 StatusCode, const FString& RawBody"
	if returnsTyped {
		completionArg = fmt.Sprintf("bool bSucceeded, int32 StatusCode, const %v& Result", responseType)
	}
	params = append(params, fmt.Sprintf("TFunction<void(%v)> OnComplete", completionArg))
	params = append(params, "FEmiClientConfig& Config = FEmiClientConfig::Default()")

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
#include "JsonObjectConverter.h"

// Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }}). Async - OnComplete
// fires on the game thread once the request finishes (see EmiHttpRequest).
class F{{ .realms.ActionName }}Client {
public:
    static void {{ .realms.FunctionName }}({{ .paramList }}) {
        FString Url = EmiBuildUrl(Config.BaseUrl, TEXT({{ .realms.Url | cppString }}));
{{ if .realms.PathParameter }}        Url = Path.Apply(Url);
{{ end }}
{{ if .realms.QueryClass }}        Url = Query.Apply(Url);
{{ end }}
        FString JsonBody;
{{ if .hasBody }}        FJsonObjectConverter::UStructToJsonObjectString(Body, JsonBody);
{{ end }}
        TMap<FString, FString> RequestHeaders;
{{ if .realms.RequestHeadersClass }}        RequestHeaders = Headers.ToHeaders();
{{ end }}
        EmiHttpRequest(TEXT({{ .method | cppString }}), Url, JsonBody, RequestHeaders, Config,
            [OnComplete](bool bSucceeded, int32 StatusCode, const FString& ResponseBody) {
                if (!bSucceeded || StatusCode >= 400) {
{{ if .returnsTyped }}                    OnComplete(false, StatusCode, {{ .responseType }}());
{{ else }}                    OnComplete(false, StatusCode, FString());
{{ end }}                    return;
                }
{{ if .returnsTyped }}                {{ .responseType }} Result;
                FJsonObjectConverter::JsonObjectStringToUStruct(ResponseBody, &Result, 0, 0);
                OnComplete(true, StatusCode, Result);
{{ else }}                OnComplete(true, StatusCode, ResponseBody);
{{ end }}
            });
    }
};
{{ end }}
`

	funcs := template.FuncMap{"cppString": escapeDoubleQuoted}
	t := template.Must(template.New("cppunrealaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"paramList":    strings.Join(params, ", "),
		"method":       strings.ToUpper(httpMethod),
		"hasBody":      hasBody,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = "F" + realms.ActionName
	res.SuggestedExtension = ".h"
	res.Realms = realms

	return res, nil
}
