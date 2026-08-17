package kotlin

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func KotlinActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	if action.GetMethod() == "reactive" {
		return KotlinActionRenderReactive(action, ctx, complexes)
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

	// responseType is the type compute() actually decodes the response body into:
	// the bare response class, wrapped in the envelope class (e.g.
	// "GResponse<ProductDto>") when the action declares one (action.Out.Envelope,
	// e.g. every entity Create/Update/Get - see preprocess-entity-actions.go). "" when
	// the action has no response shape at all, in which case payload stays untyped
	// (Any?) exactly like before this template was rewritten.
	responseType := realms.ResponseTypeName
	if realms.EnvelopeClass != "" && responseType != "" {
		responseType = fmt.Sprintf("%v<%v>", realms.EnvelopeClass, responseType)
	}

	const tmpl = `/**
 * Action to communicate with the action {{ .realms.ActionName }}
 */

data class {{ .realms.ActionName }}Meta(
    val name: String = "{{ .realms.ActionName }}",
    val url: String = "{{ .safeUrl }}",
    val method: String = "{{ .action.Method }}"
)


/*data class {{ .realms.ActionName }}Request(val call: io.ktor.server.application.ApplicationCall)*/

data class {{ .realms.ActionName }}Response(
    val statusCode: Int = 200,
    val headers: Map<String, String> = emptyMap(),
    val rawBody: String? = null,
    val payload: {{ if .responseType }}{{ .responseType }}?{{ else }}Any?{{ end }} = null
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


    // compute() actually serializes {{ if .requestType }}the typed {{ .requestType }} body{{ else }}no body (this action has none){{ end }}
    // via kotlinx.serialization before sending, and decodes the raw response body
    // {{ if .responseType }}into {{ .responseType }}{{ else }}(left untyped - this action has no response shape){{ end }} once the call returns -
    // both ends are real, not stubbed.
    suspend fun compute(
		{{ if .realms.PathParameter }}
		path: {{ .realms.ActionName}}PathParameter,
		{{ end }}
		query: Map<String, String> = emptyMap(),
		headers: Map<String, String> = emptyMap()
		{{ if .requestType }}, body: {{ .requestType }}? = null{{ end }}
	): {{ .realms.ActionName }}Response =
        withContext(Dispatchers.IO) {
            val meta = {{ .realms.ActionName }}Meta()

            var baseUrl = context?.baseUrl ?: ""
            var url = buildUrl(baseUrl, meta.url, query)

			{{ if .realms.PathParameter }}
            	url = {{ .realms.ActionName }}PathParameterApply(path, url)
			{{ end }}

            {{ if .requestType }}
            val body0 = body?.let { Json.encodeToString(it).toRequestBody(jsonType) }
            {{ else }}
            val body0: RequestBody? = null
            {{ end }}

            val requestBuilder = Request.Builder()
                .url(url)
                .method(meta.method, body0)
                .addHeader("Accept", "application/json")
            headers.forEach { (k, v) -> requestBuilder.addHeader(k, v) }

            client.newCall(requestBuilder.build()).execute().use { resp ->
                val rawBody = resp.body?.string()
                {{ if .responseType }}
                val parsedPayload: {{ .responseType }}? = rawBody?.let {
                    if (it.isEmpty()) null else Json.decodeFromString<{{ .responseType }}>(it)
                }
                {{ end }}

                {{ .realms.ActionName }}Response(
                    statusCode = resp.code,
                    headers = resp.headers.toMap(),
                    rawBody = rawBody,
                    payload = {{ if .responseType }}parsedPayload{{ else }}null{{ end }}
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
		"action":       action,
		"safeUrl":      core.RemoveTypeAnnotations(action.GetUrl()),
		"realms":       realms,
		"requestType":  realms.RequestTypeName,
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
