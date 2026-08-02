//go:build !wasm

package defs

import (
	"context"
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"net/url"
	"reflect"
)

func (x ListUsersActionRequest) IsCli() bool {
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

// ListUsersActionCliFlags returns every flag (request body, path parameters,
// query parameters and typed headers) the ListUsersAction action can bind from
// urfave v3, plus a generic repeatable --header/-H flag for anything not covered by a
// typed header.
func ListUsersActionCliFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "header",
			Aliases: []string{"H"},
			Usage:   `Raw request header as "Key: Value", repeatable`,
		},
	}
	return flags
}

// ListUsersActionCliHandler builds a full *cli.Command for the
// ListUsersAction action: it wires body, path parameters, query parameters and
// headers from urfave v3 CLI flags into a ListUsersActionRequest the same way
// ListUsersActionHandler (Gin) and ListUsersActionHttpHandler (net/http)
// do from their own transports, then prints the JSON response (or returns the error) so
// urfave reports the right exit code.
func ListUsersActionCliHandler(
	handler func(c ListUsersActionRequest) (*ListUsersActionResponse, error),
) *cli.Command {
	meta := ListUsersActionMeta()
	cmd := &cli.Command{
		Name:  meta.CliName,
		Usage: meta.Description,
		Flags: ListUsersActionCliFlags(),
	}
	cmd.Action = func(ctx context.Context, c *cli.Command) error {
		req := ListUsersActionRequest{
			CliCtx:      c,
			QueryParams: url.Values{},
			Headers:     emigo.ParseCliHeaders(c.StringSlice("header")),
		}
		return emigo.HandleActionInCli(handler(req))
	}
	return cmd
}

// ListUsersActionCli is a high-level convenience wrapper around
// ListUsersActionCliHandler. It registers the generated command as a subcommand
// of an existing urfave v3 *cli.Command, the same way ListUsersActionGin
// registers a route on a Gin engine.
func ListUsersActionCli(
	app *cli.Command,
	handler func(c ListUsersActionRequest) (*ListUsersActionResponse, error),
) {
	app.Commands = append(app.Commands, ListUsersActionCliHandler(handler))
}
