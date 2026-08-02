//go:build !wasm

package defs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x SubstringActionRequest) IsCli() bool {
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

// SubstringActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the SubstringAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func SubstringActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetSubstringActionReqCliFlags(""))...)
	return flags
}

// SubstringActionCliHandler builds a full *cli.Command for the
// SubstringAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a SubstringActionRequest the same way
// SubstringActionHandler (Gin) and SubstringActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func SubstringActionCliHandler(
	handler func(c SubstringActionRequest) (*SubstringActionResponse, error),
) *cli.Command {
	meta := SubstringActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: SubstringActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := SubstringActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastSubstringActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// SubstringActionCli is a high-level convenience wrapper around
// SubstringActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way SubstringActionGin
// registers a route on a Gin engine.
func SubstringActionCli(
	app *cli.Command,
	handler func(c SubstringActionRequest) (*SubstringActionResponse, error),
) {
	app.Commands = append(app.Commands, SubstringActionCliHandler(handler))
}
