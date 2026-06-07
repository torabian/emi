//go:build !wasm

package external

import (
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
)

func GetStreamYokeActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-path222",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func StreamYokeActionPathParameterFromCli(c *cli.Command) StreamYokeActionPathParameter {
	return StreamYokeActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
