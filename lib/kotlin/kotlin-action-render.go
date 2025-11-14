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


/*data class {{ .realms.ActionName }}Request(val call: io.ktor.server.application.ApplicationCall)*/

data class {{ .realms.ActionName }}Response(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val payload: Any? = null
)


{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
 
object {{ .realms.ActionName }}Client {
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()

    suspend fun compute(jsonPayload: String): {{ .realms.ActionName }}Response =
        withContext(Dispatchers.IO) {
            val meta = {{ .realms.ActionName }}Meta()
            val body = jsonPayload.toRequestBody(jsonType)

            val request = Request.Builder()
                .url(meta.url)
                .method(meta.method, body)
                .addHeader("Accept", "application/json")
                .build()

            client.newCall(request).execute().use { resp ->
                {{ .realms.ActionName }}Response(
                    statusCode = resp.code,
                    // body = resp.body?.string().orEmpty(),
                    headers = resp.headers.toMap()
                )
            }
        }
}

{{ if .realms.RequestClass }}
	{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
	{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}


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
