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
func GoActionReactiveCliRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms goActionReactiveRealms,
) (*core.CodeChunkCompiled, error) {
	deps := []core.CodeChunkDependency{}

	if realms.PathParameterCli != nil {
		deps = append(deps, realms.PathParameterCli.CodeChunkDependensies...)
	}

	const tmpl = `
{{ if .realms.SplitCli }}
//go:build !wasm
{{ end }}

{{ if .realms.PathParameterCli }}
	{{ b2s .realms.PathParameterCli.ActualScript }}
{{ end }}

`

	t := template.Must(template.New("reactive_cli_action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		SuggestedFileName:     realms.ActionName + "Cli",
		SuggestedExtension:    ".go",
		CodeChunkDependensies: deps,
	}, nil
}
