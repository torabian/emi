package golang

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoVsqlCompile generates a single Go file for an EmiVsql definition. The
// file contains:
//
//   - A params DTO struct (produced via GoCommonStructGenerator using the
//     vsql.Params fields, same machinery as regular EmiDto generation).
//   - A const holding the raw SQL query string.
//   - A const holding the vsql query name.
//   - A Prepare<Name>Vsql helper that returns (query string, params interface{}).
//
// The Prepare helper exposes a uniform shape consumers can pass straight to
// db.Prepare/Exec/Query call sites without binding to a specific driver.
func GoVsqlCompile(vsql core.EmiVsql, ctx core.MicroGenContext, complexes []RecognizedComplex, emiLocation string) (*core.CodeChunkCompiled, error) {

	paramsClass := vsql.GetParamsClassName()

	paramsChunk, err := GoCommonStructGenerator(vsql.Params, ctx, GoCommonStructContext{
		RootClassName:       paramsClass,
		RecognizedComplexes: complexes,
		EmiLocation:         emiLocation,
	})
	if err != nil {
		return nil, err
	}

	const tmpl = `
{{ .paramsBody }}

// {{ .queryConstName }} is the name of the vsql query, useful for logging or routing.
const {{ .queryConstName }} = "{{ .name }}"

// {{ .sqlConstName }} is the raw SQL string for the {{ .name }} vsql query.
const {{ .sqlConstName }} = ` + "`{{ .query }}`" + `

// Prepare{{ .className }} returns the query string and params for the {{ .name }} vsql query,
// ready to be passed to a SQL driver of your choice.
func Prepare{{ .className }}(params {{ .paramsClass }}) (query string, args interface{}) {
	return {{ .sqlConstName }}, params
}
`

	t := template.Must(template.New("go_vsql").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"paramsBody":     string(paramsChunk.MainClass.ActualScript),
		"className":      vsql.GetClassName(),
		"paramsClass":    paramsClass,
		"queryConstName": vsql.GetClassName() + "Name",
		"sqlConstName":   vsql.GetClassName() + "Query",
		"name":           vsql.Name,
		"query":          vsql.Query,
	}); err != nil {
		return nil, err
	}

	res := &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		CodeChunkDependensies: paramsChunk.MainClass.CodeChunkDependensies,
		Tokens:                paramsChunk.MainClass.Tokens,
		SuggestedFileName:     vsql.GetClassName() + ".go",
		SuggestedExtension:    paramsChunk.MainClass.SuggestedExtension,
	}

	return res, nil
}

// GoVsqlsGenerate iterates module.Vsqls and produces a VirtualFile per query.
func GoVsqlsGenerate(module *core.Emi, ctx core.MicroGenContext, complexes []RecognizedComplex, emiLocation, packageName string) ([]core.VirtualFile, error) {
	files := []core.VirtualFile{}
	for _, v := range module.Vsqls {
		if v.Name == "" {
			continue
		}
		chunk, err := GoVsqlCompile(v, ctx, complexes, emiLocation)
		if err != nil {
			return nil, fmt.Errorf("vsql %s: %w", v.Name, err)
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk, packageName),
		})
	}
	return files, nil
}
