package cpp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// cppUnrealReactiveActionRender renders a `method: reactive` action as a typed
// FEmiWebSocketX factory (see cpp-include/unreal/EmiWebSocketX.h), mirroring
// cppGenericReactiveActionRender for the unreal dialect.
func cppUnrealReactiveActionRender(realms cppActionRealms, res *core.CodeChunkCompiled) (*core.CodeChunkCompiled, error) {
	const tmpl = `
{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}
{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}
#include "JsonObjectConverter.h"

// Reactive (WebSocket) action {{ .realms.ActionName }} ({{ .realms.Url }})
//
// Usage:
//   TSharedRef<FEmiWebSocketX> Socket = F{{ .realms.ActionName }}Socket::Create({{ if .realms.PathParameter }}Path{{ end }});
//   Socket->OnMessage.BindLambda([](const FString& RawJson) {
{{ if .responseType }}//       {{ .responseType }} Msg;
//       FJsonObjectConverter::JsonObjectStringToUStruct(RawJson, &Msg, 0, 0);
{{ end }}//   });
class F{{ .realms.ActionName }}Socket {
public:
    static TSharedRef<FEmiWebSocketX> Create(
        {{ if .realms.PathParameter }}const F{{ .realms.ActionName }}PathParameters& Path,
        {{ end }}FEmiClientConfig& Config = FEmiClientConfig::Default()
    ) {
        FString Url = EmiBuildUrl(ToWebSocketBaseUrl(Config.BaseUrl), TEXT({{ .realms.Url | cppString }}));
{{ if .realms.PathParameter }}        Url = Path.Apply(Url);
{{ end }}
        TSharedRef<FEmiWebSocketX> Socket = MakeShared<FEmiWebSocketX>();
        Socket->Connect(Url, Config.DefaultHeaders);
        return Socket;
    }

private:
    static FString ToWebSocketBaseUrl(const FString& Base) {
        if (Base.StartsWith(TEXT("https://"))) return TEXT("wss://") + Base.RightChop(8);
        if (Base.StartsWith(TEXT("http://"))) return TEXT("ws://") + Base.RightChop(7);
        return Base;
    }
};
`

	funcs := template.FuncMap{"cppString": escapeDoubleQuoted}
	t := template.Must(template.New("cppunrealreactive").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))

	responseType := realms.ResponseClassName
	if responseType == "" && realms.ResponseClass != nil {
		responseType = ueStructName(realms.ActionName + "Response")
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"realms": realms, "responseType": responseType}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = "F" + realms.ActionName
	res.SuggestedExtension = ".h"
	res.Realms = realms

	return res, nil
}
