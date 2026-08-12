// Renders a full, standalone PHP file for a single action/remote: any
// path-parameter/query/request/response/header classes it needs, plus a
// ready-to-call namespaced function (a plain function, or a `\Generator`
// for `method: reactive`) that performs the actual HTTP call through the
// runtime Fetchx helpers.
package php

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// PhpActionRender renders the complete file for one action or remote.
func PhpActionRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {
	if action == nil || reflect.ValueOf(action).IsNil() {
		return nil, nil
	}

	realms, err := phpGetActionRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{Name: TOKEN_ORIGINAL_NAME, Value: realms.ActionName},
			{Name: TOKEN_ROOT_CLASS, Value: realms.FunctionName},
		},
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
	returnsTyped := realms.HasResponseShape && responseType != "" && responseType != "mixed" && !realms.ResponseIsAmbiguous

	// --- build the function parameter list ---------------------------------------------
	params := []string{}
	if realms.PathParameter != nil {
		params = append(params, fmt.Sprintf("%vPathParameters $path", realms.ActionName))
	}
	if hasBody {
		params = append(params, fmt.Sprintf("%v $body", requestType))
	}
	if realms.QueryClass != nil {
		params = append(params, fmt.Sprintf("?%vQueryParams $query = null", realms.ActionName))
	}
	if realms.RequestHeadersClass != nil {
		params = append(params, fmt.Sprintf("?%vRequestHeaders $headers = null", realms.ActionName))
	} else {
		params = append(params, "?array $headers = null")
	}
	params = append(params, `?\EmiSdk\Runtime\FetchxContext $ctx = null`)

	returnHint := "mixed"
	if returnsTyped {
		returnHint = responseType
	}

	headersArg := "$headers?->toHttpHeaders()"
	if realms.RequestHeadersClass == nil {
		headersArg = "$headers"
	}

	jsonBodyArg := "null"
	if hasBody {
		jsonBodyArg = "\\EmiSdk\\Runtime\\Hydrator::toArray($body)"
	}

	queryArg := "null"
	if realms.QueryClass != nil {
		queryArg = "$query?->toQueryParams()"
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
/** Communicates with {{ .realms.ActionName }} ({{ .realms.HttpMethod }} {{ .realms.Url }}) */
{{ if .realms.IsReactive }}
function {{ .realms.FunctionName }}({{ .paramList }}): \Generator
{
    $url = ($ctx ?? \EmiSdk\Runtime\Fetchx::defaultContext())->buildUrl({{ .realms.Url | phpString }});
{{ if .realms.PathParameter }}    $resolvedUrl = $path->apply($url);
{{ else }}    $resolvedUrl = $url;
{{ end }}
    foreach (\EmiSdk\Runtime\Fetchx::streamEvents(
        {{ .method | phpString }},
        $resolvedUrl,
        {{ .queryArg }},
        {{ .jsonBodyArg }},
        {{ .headersArg }},
        $ctx
    ) as $event) {
        yield {{ if .returnsTyped }}\EmiSdk\Runtime\Hydrator::fromArray({{ .responseType | phpString }}, json_decode($event, true)){{ else }}$event{{ end }};
    }
}
{{ else }}
function {{ .realms.FunctionName }}({{ .paramList }}): {{ if .returnsTyped }}?{{ .returnHint }}{{ else }}mixed{{ end }}
{
    $url = ($ctx ?? \EmiSdk\Runtime\Fetchx::defaultContext())->buildUrl({{ .realms.Url | phpString }});
{{ if .realms.PathParameter }}    $resolvedUrl = $path->apply($url);
{{ else }}    $resolvedUrl = $url;
{{ end }}
    $response = \EmiSdk\Runtime\Fetchx::request(
        {{ .method | phpString }},
        $resolvedUrl,
        {{ .queryArg }},
        {{ .jsonBodyArg }},
        {{ .headersArg }},
        $ctx
    );
{{ if .returnsTyped }}    return \EmiSdk\Runtime\Fetchx::deserialize($response, {{ .responseType | phpString }});
{{ else }}    return $response;
{{ end }}
}
{{ end }}
{{ end }}
`

	funcs := template.FuncMap{"phpString": escapeDoubleQuoted}
	t := template.Must(template.New("phpaction").Funcs(core.CommonMap).Funcs(funcs).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"realms":       realms,
		"paramList":    strings.Join(params, ", "),
		"returnHint":   returnHint,
		"method":       strings.ToUpper(httpMethod),
		"queryArg":     queryArg,
		"jsonBodyArg":  jsonBodyArg,
		"headersArg":   headersArg,
		"returnsTyped": returnsTyped,
		"responseType": responseType,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".php"
	res.Realms = realms

	return res, nil
}
