package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type StructGenerationParticles struct {
	MainClass  *core.CodeChunkCompiled
	CliHelpers *core.CodeChunkCompiled
}

// Generates a class with getters and setters.
func GoCommonStructGenerator(fields []*core.EmiField, ctx core.MicroGenContext, goctx GoCommonStructContext) (*StructGenerationParticles, error) {

	v := PrepareStruct(fields, ctx, goctx)

	const tmpl = `

{{ define "printClass" }}

{{ $item := index . 0 }}

{{ $item.GoDoc }}
{{ $item.Signature  }} {
	{{ range $item.Fields }}
		{{ .PrivateField }}
	{{ end }}
}

{{ range $item.SubClasses }}
	{{ template "printClass" (arr . ) }}
{{ end }}


{{ end }}
{{ range .renderedClasses }}
	{{ template "printClass" (arr . ) }}
{{ end }}
func (x *{{ .rootClass }}) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

  
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	// By default cli and gin is enabled. But we can disable them
	EnabledCli := !ctx.HasTag(SkipCli)
	EnabledGin := !ctx.HasTag(SkipGin)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses": v.RenderedClasses,
		"rootClass":       core.ToUpper(goctx.RootClassName),
		"EnabledCli":      EnabledCli,
		"EnabledGin":      EnabledGin,
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	v.Res.ActualScript = []byte(result)
	v.Res.SuggestedFileName = goctx.RootClassName + ".go"

	cliContent, err := GoCommonStructGeneratorCli(fields, ctx, goctx)
	if err != nil {
		return nil, err
	}

	k := &StructGenerationParticles{
		MainClass: v.Res,
	}

	f := GetCommonFlags(ctx)

	includeEmiGo := DetectIfEmiGoIsUsed2(fields)

	if includeEmiGo {
		k.MainClass.CodeChunkDependensies = append(k.MainClass.CodeChunkDependensies, core.CodeChunkDependency{
			Location: f.Emigo,
		})
	}

	splitCli := ctx.HasTag(SplitCli)
	if splitCli {
		k.CliHelpers = cliContent
	} else {
		v.Res.ActualScript = append(v.Res.ActualScript, cliContent.ActualScript...)
	}

	return k, nil
}

// Generates a class with getters and setters.
func GoCommonStructGeneratorCli(fields []*core.EmiField, ctx core.MicroGenContext, goctx GoCommonStructContext) (*core.CodeChunkCompiled, error) {

	v := PrepareStruct(fields, ctx, goctx)

	const tmpl = `
{{ if .SplitCli }}
//go:build !wasm
{{ end }}

{{ define "printClass" }}

{{ $item := index . 0 }}
{{ $enabledCli := index . 1 }}

{{ if $enabledCli }}
func Get{{ $item.FullClassName }}CliFlags(prefix string) []emigo.CliFlag {

	return []emigo.CliFlag{
		{{ range $item.Fields }}
		{
			Name: prefix + "{{ .CliName }}",
			Type: "{{ .Type }}",
			{{ if eq .Type "object" }}
			Children: Get{{ $item.FullClassName }}{{upper .Name}}CliFlags("{{ .CliName }}-"),
			{{ end }}
			{{ if .Description }}
			Description: {{escapeBackTick .Description}},
			{{ end }}
		},
		{{ end }}
	}
}


{{ $item.CliCastFunction }} {
	data := {{ $item.FullClassName }}{}

	{{ range $item.Fields }}
		{{ if .CliCaptureStatement }}
			{{ .CliCaptureStatement}}
		{{ end }}
	{{ end }}

	return data
}
{{ end }}

{{ range $item.SubClasses }}
	{{ template "printClass" (arr . $enabledCli) }}
{{ end }}


{{ end }}
{{ range .renderedClasses }}
	{{ template "printClass" (arr . $.EnabledCli) }}
{{ end }}
  
`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	// By default cli and gin is enabled. But we can disable them
	EnabledCli := !ctx.HasTag(SkipCli)
	EnabledGin := !ctx.HasTag(SkipGin)

	SplitCli := ctx.HasTag(SplitCli)

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"renderedClasses":    v.RenderedClasses,
		"emiRuntimeLocation": v.EmiLocation,
		"rootClass":          core.ToUpper(goctx.RootClassName),
		"EnabledCli":         EnabledCli,
		"EnabledGin":         EnabledGin,
		"SplitCli":           SplitCli,
	}); err != nil {
		return nil, err
	}

	result := buf.String()

	j := &core.CodeChunkCompiled{}

	f := GetCommonFlags(ctx)

	// In cli generation, always the emigo is used, if there is any kind of fields.
	if len(fields) > 0 {
		j.CodeChunkDependensies = append(j.CodeChunkDependensies, core.CodeChunkDependency{
			Location: f.Emigo,
		})
	}

	j.ActualScript = []byte(result)
	j.SuggestedFileName = goctx.RootClassName + "Cli"
	j.SuggestedExtension = ".go"

	return j, nil
}
