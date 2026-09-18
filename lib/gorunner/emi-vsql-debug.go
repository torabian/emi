package gorunner

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/torabian/emi/lib/core"
	"github.com/urfave/cli/v3"
)

// VsqlDebugCommand renders one vsql query's final SQL against
// caller-supplied parameter values and column selection, and prints it -
// without a compiled project, generated Go code, or a database connection.
// Useful while writing a vsql definition: does the template even parse, does
// the projection come out right, are optional params/columns handled the
// way you expect.
var VsqlDebugCommand = cli.Command{
	Name:  "vsql:debug",
	Usage: "Renders one vsql query's final SQL against given --params/--in/--select and prints it, for debugging a vsql definition without a database.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Usage:    "Emi module yaml file declaring the vsql",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Name of the vsql query to render (EmiVsql.Name)",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "params",
			Usage: `JSON object of parameter values keyed by field name, e.g. --params '{"id": 42}'`,
		},
		&cli.StringFlag{
			Name:  "in",
			Usage: "JSON object of values for fields declared under 'in:' - merged with --params (--params wins on a name collision, same precedence Preprocess gives Params over In)",
		},
		&cli.StringFlag{
			Name:  "select",
			Usage: "Comma-separated column names to mark selected; every other declared column becomes unselected. Omit to use each column's own 'selected:' default instead",
		},
	},
	Action: func(ctx0 context.Context, c *cli.Command) error {
		path := c.String("path")

		module, err := core.ReadEmiFromFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}

		name := c.String("name")
		var vsql *core.EmiVsql
		for i := range module.Vsqls {
			if module.Vsqls[i].Name == name {
				vsql = &module.Vsqls[i]
				break
			}
		}
		if vsql == nil {
			return fmt.Errorf("no vsql named %q in %s", name, path)
		}

		abs, err := filepath.Abs(path)
		if err != nil {
			abs = path
		}

		sql, err := core.RenderVsqlDebug(vsql, filepath.Dir(abs), c.String("params"), c.String("in"), c.String("select"))
		if err != nil {
			return err
		}

		fmt.Println(sql)
		return nil
	},
}
