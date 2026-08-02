//go:build !wasm

package defs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x DeleteUserActionRequest) IsCli() bool {
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

// DeleteUserActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the DeleteUserAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func DeleteUserActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	flags = append(flags, emigo.CastEmiFlagToUrfave(GetDeleteUserActionReqCliFlags(""))...)
	return flags
}

// DeleteUserActionCliHandler builds a full *cli.Command for the
// DeleteUserAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a DeleteUserActionRequest the same way
// DeleteUserActionHandler (Gin) and DeleteUserActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func DeleteUserActionCliHandler(
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) *cli.Command {
	meta := DeleteUserActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: DeleteUserActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := DeleteUserActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
			Body:        CastDeleteUserActionReqFromCli(c),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// DeleteUserActionCli is a high-level convenience wrapper around
// DeleteUserActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way DeleteUserActionGin
// registers a route on a Gin engine.
func DeleteUserActionCli(
	app *cli.Command,
	handler func(c DeleteUserActionRequest) (*DeleteUserActionResponse, error),
) {
	app.Commands = append(app.Commands, DeleteUserActionCliHandler(handler))
}
