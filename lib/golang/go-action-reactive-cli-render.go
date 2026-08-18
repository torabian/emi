package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// GoActionReactiveCliRender renders the cli-specific portion of a reactive action: a
// urfave v3 command that pipes stdin/stdout through the same
// factory func(session {Action}Session) (chan []byte, error) developers already write
// for {Action}Gin and {Action}ReactiveHandlerWasm - business logic is shared verbatim
// across transports.
//
// Each line read from stdin becomes one {Action}ReadChan frame; each []byte the
// factory's returned channel produces is written straight to stdout, which is what
// makes this work as one leg of a Linux pipe (`producer | emi-cli the-action | consumer`).
//
// The caller decides whether to append this to the main action file or emit it as its
// own file (controlled by the "split-cli" tag).
func GoActionReactiveCliRender(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	realms goActionReactiveRealms,
) (*core.CodeChunkCompiled, error) {

	const tmpl = `
{{ if .realms.PathParameterCli }}
	{{ b2s .realms.PathParameterCli.ActualScript }}
{{ end }}

{{ if .realms.QueryParamsCli }}
	{{ b2s .realms.QueryParamsCli.ActualScript }}
{{ end }}

// {{ .realms.ActionName }}CliFlags returns the query-parameter flags the
// {{ .realms.ActionName }} action can bind from urfave v3.
func {{ .realms.ActionName }}CliFlags() []cli.Flag {
	flags := []cli.Flag{}
	{{ if .realms.QueryParamsCli }}
	flags = append(flags, emigo.CastEmiFlagToUrfave(Get{{ .realms.ActionName }}QueryCliFlags(""))...)
	{{ end }}
	return flags
}

// {{ .realms.ActionName }}CliReactiveHandler builds a full *cli.Command for the
// {{ .realms.ActionName }} reactive action: stdin becomes the read side (one frame per
// line), and whatever the factory's returned channel produces is written straight to
// stdout - so the generated command composes as one leg of a Linux pipe
// (` + "`" + `producer | app the-action | consumer` + "`" + `). Piping ends the command the same way an
// EOF on a socket would: stdin closing (or scanner error) ends the read loop and the
// command returns.
func {{ .realms.ActionName }}CliReactiveHandler(
	factory func(session {{ .realms.ActionName }}Session) (chan []byte, error),
) *cli.Command {
	meta := {{ .realms.ActionName }}Meta()

	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: {{ .realms.ActionName }}CliFlags(),
	}

	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		done := make(chan bool)
		read := make(chan {{ .realms.ActionName }}ReadChan)

		session := {{ .realms.ActionName }}Session{
			Ctx:  c,
			Done: done,
			Read: read,
		}
		{{ if .realms.QueryParamsCli }}
		session.QueryParams = {{ .realms.ActionName }}QueryFromCli(c)
		{{ end }}

		out, err := factory(session)
		if err != nil {
			return err
		}

		go func() {
			for {
				select {
				case <-done:
					return
				case data, more := <-out:
					os.Stdout.Write(data)
					if !more {
						return
					}
				}
			}
		}()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := make([]byte, len(scanner.Bytes()))
			copy(line, scanner.Bytes())
			read <- {{ .realms.ActionName }}ReadChan{Data: line}
		}

		return scanner.Err()
	}

	return cmd
}

// {{ .realms.ActionName }}CliReactive is a high-level convenience wrapper around
// {{ .realms.ActionName }}CliReactiveHandler. It registers the generated command as a
// subcommand of an existing urfave v3 *cli.Command, the same way {{ .realms.ActionName }}Gin
// registers a route on a Gin engine.
func {{ .realms.ActionName }}CliReactive(
	app *cli.Command,
	factory func(session {{ .realms.ActionName }}Session) (chan []byte, error),
) {
	app.Commands = append(app.Commands, {{ .realms.ActionName }}CliReactiveHandler(factory))
}
`

	t := template.Must(template.New("reactive_cli_action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action": action,
		"realms": realms,
	}); err != nil {
		return nil, err
	}

	deps := []core.CodeChunkDependency{
		{Location: "bufio"},
		{Location: "context"},
		{Location: "os"},
		{Location: "github.com/urfave/cli/v3"},
	}

	if realms.PathParameterCli != nil {
		deps = append(deps, realms.PathParameterCli.CodeChunkDependensies...)
	}
	if realms.QueryParamsCli != nil {
		deps = append(deps, realms.QueryParamsCli.CodeChunkDependensies...)
	}

	return &core.CodeChunkCompiled{
		ActualScript:          buf.Bytes(),
		SuggestedFileName:     realms.ActionName + "Cli",
		SuggestedExtension:    ".go",
		CodeChunkDependensies: deps,
	}, nil
}
