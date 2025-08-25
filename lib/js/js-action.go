package js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func JsActionClass(action *core.Module3Action, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{}
	actionRealms, deps, err := JsActionClassRealms(action, ctx)

	res.Tokens = []core.GeneratedScriptToken{
		{
			Name:  TOKEN_ORIGINAL_NAME,
			Value: action.Name,
		},
	}
	if err != nil {
		return nil, err
	}

	const tmpl = `/**
* Action to communicate with the action {{ .action.Name }}
*/

{{ if .realms.OptionsType }}
	{{ b2s .realms.OptionsType.ActualScript }}
{{ end }}

{{ if .realms.ReactQueryOptions }}
	{{ b2s .realms.ReactQueryOptions.ActualScript }}
{{ end }}

{{ if .realms.UseQueryFunction }}
	{{ b2s .realms.UseQueryFunction.ActualScript }}
{{ end }}


{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}
 

{{ if .fetch }}
	{{ .fetch }}
{{ end }}

{{ if .realms.RequestClass }}
{{ b2s .realms.RequestClass.ActualScript }}
{{ end }}

{{ if .realms.ResponseClass }}
{{ b2s .realms.ResponseClass.ActualScript }}
{{ end }}

{{ if .headerCompiledClass }}
{{ .headerCompiledClass }}
{{ end }}

{{ if .qsCompiledClass }}
{{ .qsCompiledClass }}
{{ end }}
 
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))
	nestJsDecorator := strings.Contains(ctx.Tags, GEN_NEST_JS_COMPATIBILITY)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":              action,
		"headerCompiledClass": string(actionRealms.RequestHeadersClass.ActualScript),
		"qsCompiledClass":     string(actionRealms.QueryStringClass.ActualScript),
		"shouldExport":        true,
		"nestjsDecorator":     nestJsDecorator,
		"fetch":               string(actionRealms.FetchMetaClass.ActualScript),
		"realms":              actionRealms,
		"className":           fmt.Sprintf("%vAction", action.Upper()),
	}); err != nil {
		return nil, err
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, deps...)
	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToUpper(action.Name) + "Action"
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	res.Realms = actionRealms

	return res, nil
}
