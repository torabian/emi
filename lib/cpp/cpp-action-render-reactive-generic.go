package cpp

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// cppGenericReactiveActionRender renders a `method: reactive` action as a typed
// emi::EmiWebSocketX factory (see cpp-include/generic/EmiWebSocketX.hpp),
// mirroring the split every other reactive-capable target has between a classic
// request/response action and a websocket one (lib/swift/swift-action-reactive-render.go,
// lib/golang/go-action-reactive-*.go, ...).
func cppGenericReactiveActionRender(action core.EmiRpcAction, realms cppActionRealms, res *core.CodeChunkCompiled) (*core.CodeChunkCompiled, error) {
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
// Reactive (WebSocket) action {{ .realms.ActionName }} ({{ .realms.Url }})
//
// Usage:
//   emi::EmiByteStreamPosix stream; // or EmiByteStreamArduino
{{ if .realms.QueryClass }}//   {{ .realms.ActionName }}QueryParams query; // typed, appended to the url
{{ end }}//   auto socket = {{ .realms.ActionName }}Socket::Create(stream{{ if .realms.PathParameter }}, path{{ end }}{{ if .realms.QueryClass }}, query{{ end }});
//   socket->onMessage = [](const std::string& raw) { // text frames (JSON dtos)
{{ if .responseType }}//       auto msg = {{ .responseType }}::Parse(raw);
{{ end }}//   };
{{ if .binary }}//   socket->onBinary = [](const uint8_t* data, size_t length) { /* raw bytes */ };
//   socket->SendBinary(data, length); // or emi::EmiMediaSender (EmiMediaFrame.hpp) for mic + camera
{{ end }}//   while (socket->IsOpen()) { socket->Poll(); }
class {{ .realms.ActionName }}Socket {
public:
    static std::unique_ptr<emi::EmiWebSocketX> Create(
        emi::IEmiByteStream& stream{{ if .realms.PathParameter }},
        const {{ .realms.ActionName }}PathParameters& path{{ end }}{{ if .realms.QueryClass }},
        const {{ .realms.ActionName }}QueryParams& query{{ end }},
        emi::EmiClientConfig& config = emi::EmiClientConfig::Default()
    ) {
        std::string url = emi::EmiBuildUrl(ToWebSocketBaseUrl(config.baseUrl), {{ .realms.Url | cppString }});
{{ if .realms.PathParameter }}        url = path.Apply(url);
{{ end }}{{ if .realms.QueryClass }}        url = query.Apply(url);
{{ end }}
        auto socket = std::make_unique<emi::EmiWebSocketX>(stream);
        socket->Connect(url);
        return socket;
    }

private:
    static std::string ToWebSocketBaseUrl(const std::string& base) {
        if (base.rfind("https://", 0) == 0) return "wss://" + base.substr(8);
        if (base.rfind("http://", 0) == 0) return "ws://" + base.substr(7);
        return base;
    }
};
`

	funcs := template.FuncMap{"cppString": escapeDoubleQuoted}
	t := template.Must(template.New("cppgenericreactive").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))

	responseType := realms.ResponseClassName
	if responseType == "" && realms.ResponseClass != nil {
		responseType = realms.ActionName + "Response"
	}

	binary := false
	if a, ok := action.(interface{ GetBinaryType() string }); ok {
		binary = a.GetBinaryType() != "" && a.GetBinaryType() != "text"
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"realms": realms, "responseType": responseType, "binary": binary}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".hpp"
	res.Realms = realms

	return res, nil
}
