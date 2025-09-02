// Renders the common object, such as entities, dtos.

package js

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type JsCommonObjectContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string
}

// This function can be used in different locations of the code generation,
// creates dtos, entities for actions or other specs.
func JsCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{}

	tsTypes, tsTypeError := TsCommonObjectGenerator(fields, ctx, TsCommonObjectContext{
		RootTypeName: jsctx.RootClassName,
	})
	if tsTypeError != nil {
		return nil, tsTypeError
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, tsTypes.CodeChunkDependenies...)

	tsClass, tsClassError := JsCommonObjectClassGenerator(
		fields,
		ctx,
		JsCommonObjectContext{RootClassName: jsctx.RootClassName},
	)
	if tsClassError != nil {
		return nil, tsClassError
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, tsClass.CodeChunkDependenies...)

	const tmpl = `
{{ .tsClass }}
{{ .tsInterface }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"shouldExport": true,
		"tsInterface":  string(tsTypes.ActualScript),
		"tsClass":      string(tsClass.ActualScript),
		"fields":       fields,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = jsctx.RootClassName
	res.SuggestedExtension = ".js"

	if isTypeScript {
		res.SuggestedExtension = ".ts"
	}

	return res, nil
}
