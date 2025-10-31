package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/golang"
	"github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"

	"github.com/urfave/cli"
)

func main() {

	commands := []cli.Command{}
	commands = append(commands,
		cliCommandFromTextActions(js.GetJsPublicActions().TextActions)...)
	commands = append(commands,
		cliCommandFromFileActions(js.GetJsPublicActions().FileActions)...)

	commands = append(commands,
		cliCommandFromTextActions(golang.GetGolangPublicActions().TextActions)...)
	commands = append(commands,
		cliCommandFromFileActions(golang.GetGolangPublicActions().FileActions)...)

	commands = append(commands,
		cliCommandFromTextActions(kotlin.GetKotlinPublicActions().TextActions)...)
	commands = append(commands,
		cliCommandFromFileActions(kotlin.GetKotlinPublicActions().FileActions)...)

	app := &cli.App{
		Name:     "Emi compiler",
		Usage:    "Module3 definitions code generator",
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

// Casts the core module action definitions into callable cli.Flags
func cliFlagsFromDefs(defs []core.FlagDef) []cli.Flag {
	var flags []cli.Flag

	for _, f := range defs {
		switch f.Type {
		case core.FlagString:
			flags = append(flags, cli.StringFlag{
				Name:     f.Name,
				Usage:    f.Usage,
				Required: f.Required,
				Value:    fmt.Sprintf("%v", f.Default),
			})
		case core.FlagBool:
			flags = append(flags, cli.BoolFlag{
				Name:  f.Name,
				Usage: f.Usage,
			})
		case core.FlagInt:
			flags = append(flags, cli.IntFlag{
				Name:  f.Name,
				Usage: f.Usage,
			})
		}
	}

	return flags
}

func cliCommandFromTextActions(actions []core.ActionText) []cli.Command {
	items := []cli.Command{}
	for _, item := range actions {
		items = append(items, cliCommandFromTextAction(item))
	}
	return items
}

func cliCommandFromFileActions(actions []core.ActionFile) []cli.Command {
	items := []cli.Command{}
	for _, item := range actions {
		items = append(items, cliCommandFromFileAction(item))
	}
	return items
}

var commonFlags []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "path",
		Usage: "Path of the file on the disk",
	},
	cli.StringFlag{
		Name:  "output",
		Usage: "The directory which the generated files will be rewritten to",
	},
	cli.StringFlag{
		Name:  "tags",
		Usage: "A set of string flags separated by comma (,) to add or remove compile feature. Such as 'nestjs-headers-decorator'",
	},
}

func createCliContext(c *cli.Context, flags []core.FlagDef) (core.MicroGenContext, error) {
	ctx := core.MicroGenContext{
		Tags:   c.String("tags"),
		Output: c.String("output"),
	}

	// Allow silently missing file to fail.
	content, _ := os.ReadFile(c.String("path"))
	// if err != nil {
	// 	return core.MicroGenContext{}, err
	// }

	var m map[string]string = map[string]string{}

	for _, flag := range flags {
		fmt.Println("Flag:", flag.Name, c.String(flag.Name))
		m[flag.Name] = c.String(flag.Name)
	}

	res, _ := json.Marshal(m)

	ctx.Flags = string(res)
	ctx.Content = string(content)

	return ctx, nil
}

func cliCommandFromTextAction(a core.ActionText) cli.Command {

	// The action definition might ask for some flags,
	// but in the meantime we need to add some related core.ActionText
	flags := cliFlagsFromDefs(a.BaseAction.Flags)
	flags = append(flags, commonFlags...)

	return cli.Command{
		Name:        a.Name,
		Description: a.Description,
		Usage:       a.Description,
		Flags:       flags,
		Action: func(c *cli.Context) error {
			ctx, err := createCliContext(c, a.Flags)
			if err != nil {
				return err
			}

			// Let's combine the import requirements of the chunk
			output, err := a.Run(ctx)
			if err != nil {
				return err
			}

			// If there is no output, we just write the content as json into the output
			if ctx.Output == "" {

				fmt.Println(string(output))

				return nil
			}

			if err := os.WriteFile(ctx.Output, []byte(output), 0644); err != nil {
				return fmt.Errorf("error on writing file to disk: %v, %v, %w", ctx.Output, err)
			}

			return nil
		},
	}
}

func cliCommandFromFileAction(a core.ActionFile) cli.Command {

	// The action definition might ask for some flags,
	// but in the meantime we need to add some related core.ActionText
	flags := cliFlagsFromDefs(a.Flags)
	flags = append(flags, commonFlags...)

	return cli.Command{
		Name:        a.Name,
		Description: a.Description,
		Usage:       a.Description,
		Flags:       flags,
		Action: func(c *cli.Context) error {
			ctx, err := createCliContext(c, a.Flags)
			if err != nil {
				return err
			}

			data := make(map[string]interface{})
			for _, flag := range a.Flags {
				if !c.IsSet(flag.Name) {
					continue
				}
				switch flag.Type {
				case core.FlagBool:
					data[flag.Name] = c.Bool(flag.Name)
				case core.FlagString:
					data[flag.Name] = c.String(flag.Name)
				case core.FlagInt:
					data[flag.Name] = c.Int(flag.Name)
				}
			}
			out, _ := json.MarshalIndent(data, "", "  ")
			ctx.Flags = string(out)

			// Let's combine the import requirements of the chunk
			files, err := a.Run(ctx)
			if err != nil {
				return err
			}

			if ctx.Output == "" {
				res, _ := json.MarshalIndent(files, "", "  ")
				fmt.Println(string(res))

				return nil
			}

			for _, file := range files {

				filePath := path.Join(ctx.Output, file.Location, file.Name+file.Extension)
				os.MkdirAll(path.Dir(filePath), os.ModePerm)

				if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
					return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
				}
			}

			return nil
		},
	}
}
