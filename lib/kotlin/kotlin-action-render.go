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

	public var context: ClientContext? = null
    private val client = OkHttpClient()
    private val jsonType = "application/json".toMediaType()

    fun buildUrl(base: String, path: String, query: Map<String, String>): String {
        val baseUrl = base.toHttpUrl()   // parses full URL like "http://asdasda/"

        val urlBuilder = baseUrl
            .newBuilder()
            .encodedPath(path)

        query.forEach { (k, v) ->
            urlBuilder.addQueryParameter(k, v)
        }

        return urlBuilder.build().toString()
    }


    suspend fun compute(
		{{ if .realms.PathParameter }}
		path: {{ .realms.ActionName}}PathParameter,
		{{ end }}
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap(),
		body: String? = null
	): {{ .realms.ActionName }}Response =
        withContext(Dispatchers.IO) {
            val meta = {{ .realms.ActionName }}Meta()

            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)

			{{ if .realms.PathParameter }}
            	url = {{ .realms.ActionName }}PathParameterApply(path, url)
			{{ end }}

            println(  url)

            val body0 = body?.toRequestBody(jsonType)

            val request = Request.Builder()
                .url(url)
                .method(meta.method, body0)
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
