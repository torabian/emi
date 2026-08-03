package swift

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// swiftFindTokenByName mirrors js-action-main-class.go's findTokenByName - swift has no
// shared copy of its own yet, and this is small enough not to warrant exporting one
// across packages for a single caller.
func swiftFindTokenByName(tokens []core.GeneratedScriptToken, name string) *core.GeneratedScriptToken {
	for _, item := range tokens {
		if item.Name == name {
			return &item
		}
	}
	return nil
}

// swiftReactiveMessageTypes resolves the Send/Receive type names EmiWebSocketX gets
// parameterized with, from the exact same realms.RequestClass/ResponseClass GetActionRealms
// already builds for classic actions (see swift-action-realms.go) - a reactive action's
// in/out fields work identically to a classic one's, only the transport differs. An
// action with no in/out declared (qs-only, e.g. a bare subscribe channel) falls back to
// EmiAnyCodable - the same "no fixed shape" wrapper a "type: any" field uses.
func swiftReactiveMessageTypes(realms actionRealms) (requestType string, responseType string) {
	requestType = "EmiAnyCodable"
	responseType = "EmiAnyCodable"

	if realms.RequestClass != nil {
		if token := swiftFindTokenByName(realms.RequestClass.Tokens, TOKEN_OBJ_CLASS); token != nil {
			requestType = token.Value
		}
	}
	if realms.ResponseClass != nil {
		if token := swiftFindTokenByName(realms.ResponseClass.Tokens, TOKEN_OBJ_CLASS); token != nil {
			responseType = token.Value
		}
	}

	return requestType, responseType
}

// SwiftActionRenderReactive renders a "method: reactive" action as a typed WebSocket
// client (see EmiWebSocketX, swift-websocket-runtime.go), mirroring the split Go
// (go-action-render.go) and JS (js-static-websocket-create.go) already have between a
// classic request/response action and a reactive one - Swift had neither branch at all
// before this, every action rendered as if it were classic HTTP regardless of method.
//
// Unlike JS (a browser WebSocket resolves a relative "/path" url against the current
// page's origin automatically, swapping http/https for ws/wss) there is no implicit
// origin in a native Swift app: the ws(s) url is derived explicitly from
// EmiClientConfig.baseUrl (see swift-client-config.go) by swapping its scheme.
func SwiftActionRenderReactive(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, deps, err := GetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	requestType, responseType := swiftReactiveMessageTypes(realms)

	const tmpl = `import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

/**
 Reactive (WebSocket) action {{ .realms.ActionName }}
 */
struct {{ .realms.ActionName }}Meta {
    let name: String = "{{ .realms.ActionName }}"
    let url: String = "{{ .safeUrl }}"
}

{{ if .realms.PathParameter }}
{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

typealias {{ .realms.ActionName }}Socket = EmiWebSocketX<{{ .requestType }}, {{ .responseType }}>

enum {{ .realms.ActionName }} {
    private static func webSocketBaseUrl() -> String {
        let base = EmiClientConfig.baseUrl
        if base.hasPrefix("https://") {
            return "wss://" + base.dropFirst("https://".count)
        }
        if base.hasPrefix("http://") {
            return "ws://" + base.dropFirst("http://".count)
        }
        return base
    }

    // Opens a new connection. Nil only if EmiClientConfig.baseUrl (plus path/query) is
    // not a valid URL - call .connect() on the result to actually open the socket.
    static func Create(
        {{ if .realms.PathParameter }}
        path: {{ .realms.ActionName }}PathParameter,
        {{ end }}
        query: [String: String] = [:]
    ) -> {{ .realms.ActionName }}Socket? {
        let meta = {{ .realms.ActionName }}Meta()

        guard var components = URLComponents(string: webSocketBaseUrl()) else {
            return nil
        }

        {{ if .realms.PathParameter }}
        components.path = {{ .realms.ActionName }}PathParameterApply(path, meta.url)
        {{ else }}
        components.path = meta.url
        {{ end }}

        if !query.isEmpty {
            components.queryItems = query.map {
                URLQueryItem(name: $0.key, value: $0.value)
            }
        }

        guard let url = components.url else {
            return nil
        }

        return {{ .realms.ActionName }}Socket(url: url)
    }
}

{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}
`

	t := template.Must(template.New("actionreactive").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"safeUrl":      core.RemoveTypeAnnotations(action.GetUrl()),
		"realms":       realms,
		"requestType":  requestType,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".swift"
	res.CodeChunkDependensies = deps

	return res, nil
}
