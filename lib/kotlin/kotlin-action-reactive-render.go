package kotlin

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// KotlinActionRenderReactive renders a "method: reactive" action as a typed WebSocket
// client (see EmiWebSocketX, lib/kotlin/kotlin-include/emiwebsocketx.kt), mirroring the
// split lib/swift/swift-action-reactive-render.go already has between a classic
// request/response action and a reactive one - Kotlin had neither branch at all before
// this, every action rendered as if it were classic HTTP regardless of method.
//
// Unlike a browser (a relative "/path" WebSocket resolves against the page's own
// origin, swapping http/https for ws/wss automatically) there's no implicit origin in a
// JVM app: the ws(s) url is derived explicitly from ClientContext.baseUrl (the exact
// same per-object context every classic action's *Client already reads) by swapping its
// scheme, same reasoning as Swift's EmiClientConfig-based derivation.
func KotlinActionRenderReactive(
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

	deps = append(deps, core.CodeChunkDependency{Location: "emikot.EmiWebSocketX"})

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: core.TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
		},
	}

	// A bare qs-only channel (no in/out declared) falls back to JsonElement - the same
	// "no fixed shape" wrapper "type: any" fields resolve to elsewhere in this
	// compiler, and, unlike a plain String/Any, still has a real KSerializer
	// (JsonElement.serializer()) to hand EmiWebSocketX.
	requestType := realms.RequestTypeName
	if requestType == "" {
		requestType = "JsonElement"
	}
	responseType := realms.ResponseTypeName
	if responseType == "" {
		responseType = "JsonElement"
	}

	const tmpl = `/**
 * Reactive (WebSocket) action {{ .realms.ActionName }}
 */

data class {{ .realms.ActionName }}Meta(
    val name: String = "{{ .realms.ActionName }}",
    val url: String = "{{ .safeUrl }}"
)

{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

typealias {{ .realms.ActionName }}Socket = EmiWebSocketX<{{ .requestType }}, {{ .responseType }}>

object {{ .realms.ActionName }} {

	public var context: ClientContext? = null
	private val client = OkHttpClient()

	private fun webSocketBaseUrl(): String {
		val base = context?.baseUrl ?: ""
		return when {
			base.startsWith("https://") -> "wss://" + base.removePrefix("https://")
			base.startsWith("http://") -> "ws://" + base.removePrefix("http://")
			else -> base
		}
	}

	// Builds an unconnected socket - call .connect() on the result to actually open it,
	// same two-step shape (Create then connect) as Swift's EmiWebSocketX.
	fun Create(
		{{ if .realms.PathParameter }}
		path: {{ .realms.ActionName }}PathParameter,
		{{ end }}
		query: Map<String, String> = emptyMap()
	): {{ .realms.ActionName }}Socket {
		val meta = {{ .realms.ActionName }}Meta()

		val baseUrl = webSocketBaseUrl().toHttpUrl()
		val urlBuilder = baseUrl.newBuilder().encodedPath(meta.url)
		query.forEach { (k, v) -> urlBuilder.addQueryParameter(k, v) }
		var url = urlBuilder.build().toString()

		{{ if .realms.PathParameter }}
		url = {{ .realms.ActionName }}PathParameterApply(path, url)
		{{ end }}

		return {{ .realms.ActionName }}Socket(
			client = client,
			url = url,
			sendSerializer = {{ .requestType }}.serializer(),
			receiveSerializer = {{ .responseType }}.serializer(),
		)
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
	res.SuggestedExtension = ".kt"
	res.CodeChunkDependensies = deps

	return res, nil
}
