package kotlin

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func KotlinActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
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

	const tmpl = `/**
 * Action to communicate with the action {{ .realms.ActionName }}
 */

data class {{ .realms.ActionName }}Meta(
    val name: String = "{{ .realms.ActionName }}",
    val url: String = "{{ .action.Url }}",
    val method: String = "{{ .action.Method }}"
)

data class {{ .realms.ActionName }}Request(val call: io.ktor.server.application.ApplicationCall)

data class {{ .realms.ActionName }}Response(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)

// Registers a raw route for {{ .realms.ActionName }} using Ktor routing
fun {{ .realms.ActionName }}Raw(route: io.ktor.server.routing.Route, handler: suspend ({{ .realms.ActionName }}Request) -> {{ .realms.ActionName }}Response) {
    val meta = {{ .realms.ActionName }}Meta()
    route.route(meta.url, io.ktor.server.routing.HttpMethod.valueOf(meta.method)) {
        handle {
            val req = {{ .realms.ActionName }}Request(call)
            val resp = handler(req)
            resp.headers.forEach { (k, v) -> call.response.headers.append(k, v) }
            call.respond(resp.statusCode, resp.payload ?: "")
        }
    }
}

// Convenience wrapper for route registration
fun {{ .realms.ActionName }}(route: io.ktor.server.routing.Route, handler: suspend ({{ .realms.ActionName }}Request) -> {{ .realms.ActionName }}Response) {
    {{ .realms.ActionName }}Raw(route, handler)
}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".kt"
	res.CodeChunkDependensies = deps

	return res, nil
}
