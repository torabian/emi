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

type jsActionRealms struct {
	RequestClass        *core.CodeChunkCompiled
	ResponseClass       *core.CodeChunkCompiled
	QueryStringClass    *core.CodeChunkCompiled
	RequestHeadersClass *core.CodeChunkCompiled
}

func JsActionClass(action *core.Module3Action, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {
	actionRealms := jsActionRealms{}

	res := &core.CodeChunkCompiled{}

	// Header is the http headers, extending the Headers class from standard javascript
	header, err := JsActionHeaderClass(action, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, header.CodeChunkDependenies...)
	actionRealms.RequestHeadersClass = header

	// Query strings for the request builder
	qs, err := JsActionQsClass(action, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, qs.CodeChunkDependenies...)
	actionRealms.QueryStringClass = qs

	// Action request (in)
	if action.In != nil && len(action.In.Fields) > 0 {
		fields, err := JsCommonObjectGenerator(action.In.Fields, ctx, JsCommonObjectContext{
			RootClassName: action.Name + "Req",
		})

		if err != nil {
			return nil, err
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fields.CodeChunkDependenies...)
		actionRealms.RequestClass = fields
	}

	// Action response (out)
	if action.Out != nil && len(action.Out.Fields) > 0 {
		fields, err := JsCommonObjectGenerator(action.Out.Fields, ctx, JsCommonObjectContext{
			RootClassName: action.Name + "Res",
		})

		if err != nil {
			return nil, err
		}

		res.CodeChunkDependenies = append(res.CodeChunkDependenies, fields.CodeChunkDependenies...)
		actionRealms.ResponseClass = fields
	}

	fetch, err := JsActionFetchAndMetaData(action, actionRealms, ctx)
	if err != nil {
		return nil, err
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, fetch.CodeChunkDependenies...)

	const tmpl = `/**
* Action to communicate with the action {{ .action.Name }}
*/

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
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":              action,
		"headerCompiledClass": string(header.ActualScript),
		"qsCompiledClass":     string(qs.ActualScript),
		"shouldExport":        true,
		"nestjsDecorator":     nestJsDecorator,
		"fetch":               string(fetch.ActualScript),
		"realms":              actionRealms,
		"className":           fmt.Sprintf("%vAction", action.Upper()),
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = core.ToUpper(action.Name) + "Action"
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	return res, nil
}
