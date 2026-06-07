package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoActionCliRender renders the cli-specific portion of an action: currently
// the IsCli() helper method on the action's Request type.
//
// The caller decides whether to append this to the main action file or
// emit it as its own file (controlled by the "split-cli" tag).
func GoActionCliRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms goActionRealms,
) (*core.CodeChunkCompiled, error) {

	const tmpl = `
{{ if .realms.SplitCli }}
//go:build !wasm
{{ end }}

{{ if .realms.PathParameterCli }}
	{{ b2s .realms.PathParameterCli.ActualScript }}
{{ end }}


func (x {{ .realms.ActionName }}Request) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}

	v := reflect.ValueOf(x.CliCtx)

	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}

	return true
}
`

	t := template.Must(template.New("action_cli").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	esx := core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  realms.ActionName + "Cli",
		SuggestedExtension: ".go",
		CodeChunkDependensies: []core.CodeChunkDependency{
			{
				Location: "reflect",
			},
		},
	}

	if realms.PathParameterCli != nil && realms.SplitCli {
		esx.CodeChunkDependensies = append(esx.CodeChunkDependensies, realms.PathParameterCli.CodeChunkDependensies...)
	}

	return &esx, nil
}
