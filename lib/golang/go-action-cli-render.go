package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoActionCliRender renders the cli-specific portion of an action: the IsCli() helper
// method on the action's Request type, plus a full urfave/cli v3 binding
// ({Action}CliFlags/{Action}CliHandler/{Action}Cli) that wires body, path parameters,
// query parameters and headers from CLI flags into a {Action}Request - the same job
// GoActionGinRender and GoActionHttpRender do for their own transports.
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

{{ if .realms.QueryParamsCli }}
	{{ b2s .realms.QueryParamsCli.ActualScript }}
{{ end }}

{{ if .realms.HeadersCli }}
	{{ b2s .realms.HeadersCli.ActualScript }}
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

// {{ .realms.ActionName }}CliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the {{ .realms.ActionName }} action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func {{ .realms.ActionName }}CliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   ` + "`" + `Raw request header as "Key: Value", repeatable` + "`" + `,
		},
	}

	{{ if .action.HasRequestFields }}
	flags = append(flags, emigo.CastEmiFlagToUrfave(Get{{ .realms.RequestClassName }}CliFlags(""))...)
	{{ end }}
	{{ if .realms.PathParameterCli }}
	flags = append(flags, emigo.CastEmiFlagToUrfave(Get{{ .realms.ActionName }}PathParameterCliFlags(""))...)
	{{ end }}
	{{ if .realms.QueryParamsCli }}
	flags = append(flags, emigo.CastEmiFlagToUrfave(Get{{ .realms.ActionName }}QueryCliFlags(""))...)
	{{ end }}
	{{ if .realms.HeadersCli }}
	flags = append(flags, emigo.CastEmiFlagToUrfave(Get{{ .realms.ActionName }}HeaderCliFlags(""))...)
	{{ end }}

	return flags
}

// {{ .realms.ActionName }}CliHandler builds a full *cli.Command for the
// {{ .realms.ActionName }} action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a {{ .realms.ActionName }}Request the same way
// {{ .realms.ActionName }}Handler (Gin) and {{ .realms.ActionName }}HttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func {{ .realms.ActionName }}CliHandler(
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) *cli.Command {
	meta := {{ .realms.ActionName }}Meta()

	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: {{ .realms.ActionName }}CliFlags(),
	}

	{{ if .realms.CliShort }}
	cmd.Aliases = []string{meta.CliShort}
	{{ end }}

	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := {{ .realms.ActionName }}Request{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			{{ if .action.HasRequestFields }}
			Body: Cast{{ .realms.RequestClassName }}FromCli(c),
			{{ end }}
			{{ if .realms.PathParameterCli }}
			Params: {{ .realms.ActionName }}PathParameterFromCli(c),
			{{ end }}
		}

		{{ if .realms.QueryParamsCli }}
		req.QueryParams = {{ .realms.ActionName }}QueryFromCli(c).Values()
		{{ end }}

		{{ if .realms.HeadersCli }}
		for k, v := range {{ .realms.ActionName }}HeadersFromCli(c) {
			req.Headers[k] = v
		}
		{{ end }}

		return emigo.HandleActionInCli(handler(req))
	}

	return cmd
}

// {{ .realms.ActionName }}Cli is a high-level convenience wrapper around
// {{ .realms.ActionName }}CliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way {{ .realms.ActionName }}Gin
// registers a route on a Gin engine.
func {{ .realms.ActionName }}Cli(
	app *cli.Command,
	handler func(c {{ .realms.ActionName }}Request) (*{{ .realms.ActionName }}Response, error),
) {
	app.Commands = append(app.Commands, {{ .realms.ActionName }}CliHandler(handler))
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

	f := GetCommonFlags(ctx)

	esx := core.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  realms.ActionName + "Cli",
		SuggestedExtension: ".go",
		CodeChunkDependensies: []core.CodeChunkDependency{
			{Location: "reflect"},
			{Location: "context"},
			{Location: "net/url"},
			{Location: "github.com/urfave/cli/v3"},
			{Location: f.Emigo},
		},
	}

	if realms.PathParameterCli != nil && realms.SplitCli {
		esx.CodeChunkDependensies = append(esx.CodeChunkDependensies, realms.PathParameterCli.CodeChunkDependensies...)
	}
	if realms.QueryParamsCli != nil && realms.SplitCli {
		esx.CodeChunkDependensies = append(esx.CodeChunkDependensies, realms.QueryParamsCli.CodeChunkDependensies...)
	}
	if realms.HeadersCli != nil && realms.SplitCli {
		esx.CodeChunkDependensies = append(esx.CodeChunkDependensies, realms.HeadersCli.CodeChunkDependensies...)
	}

	return &esx, nil
}
