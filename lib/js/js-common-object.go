package js

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type RecognizedComplex struct {
	Symbol string

	ImportLocation string
}

type JsCommonObjectContext struct {

	// The class name which will be used to generate nested classes,
	// in case of array or object
	RootClassName string

	// List of allowed complexes type to be used on fields
	RecognizedComplexes []RecognizedComplex
}

// This function can be used in different locations of the code generation,
// creates dtos, entities for actions or other specs.
func JsCommonObjectGenerator(fields []*core.EmiField, ctx core.MicroGenContext, jsctx JsCommonObjectContext) (*core.CodeChunkCompiled, error) {
	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	res := &core.CodeChunkCompiled{
		CodeChunkDependenies: []core.CodeChunkDependency{
			{
				Objects:  []string{"isPlausibleObject"},
				Location: INTERNAL_SDK_JS_LOCATION + "/isPlausibleObject",
			},
			{
				Objects:  []string{"withPrefix"},
				Location: INTERNAL_SDK_JS_LOCATION + "/withPrefix",
			},
		},
	}
	var tsTypes *core.CodeChunkCompiled

	if isTypeScript {
		chunk, tsTypeError := TsCommonObjectGenerator(fields, ctx, TsCommonObjectContext{
			RootTypeName: jsctx.RootClassName,
		})

		if tsTypeError != nil {
			return nil, tsTypeError
		}

		res.Tokens = append(res.Tokens, chunk.Tokens...)
		res.CodeChunkDependenies = append(res.CodeChunkDependenies, chunk.CodeChunkDependenies...)
		tsTypes = chunk
	}

	tsClass, tsClassError := JsCommonObjectClassGenerator(
		fields,
		ctx,
		jsctx,
	)

	if tsClassError != nil {
		return nil, tsClassError
	}

	res.Tokens = append(res.Tokens, tsClass.Tokens...)
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, tsClass.CodeChunkDependenies...)

	const tmpl = `
{{ if .tsClass }}
	{{ b2s .tsClass.ActualScript }}
{{ end }}

{{ if .tsInterface }}
	{{ b2s .tsInterface.ActualScript }}
{{ end }}
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"shouldExport": true,
		"tsInterface":  tsTypes,
		"tsClass":      tsClass,
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
