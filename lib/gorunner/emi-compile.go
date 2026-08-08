package gorunner

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/golang"
	"github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"
	"github.com/torabian/emi/lib/swift"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

// Compiles a emi file automatically, based on common compiler functions enabled into
// the command, and it would read tags, out dirs, flags, etc from the file itself.
// this makes emi self contained.
var CompileCommand = cli.Command{
	Name:  "compile",
	Usage: "Compiles a emi file, without accepting tags or flags, and reads them all from the 'targets' section of the emi definition file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "path",
			Usage:    "Translation yaml or json entry point",
			Required: true,
		},
	},
	Action: func(ctx0 context.Context, c *cli.Command) error {
		ctx, err := createCliContext(c, []core.FlagDef{})
		if err != nil {
			return err
		}

		m, err := core.ReadEmiFromString(ctx.Content)
		if err != nil {
			return err
		}

		outdir := ctx.Output
		for _, target := range m.Targets {
			ctx.Flags = target.Flags
			ctx.Tags = strings.Join(target.Tags, ",")
			files := []core.VirtualFile{}

			// "preprocessor" isn't a real per-language core.ActionFile - it doesn't
			// generate code, it just writes back out what core.Emi.Preprocess()
			// expanded the module into (synthesized entity update dtos/actions, etc. -
			// see lib/core/preprocess*.go), as plain yaml. Useful for inspecting what
			// every other compiler actually sees once preprocessing has run.
			if target.Compiler == "preprocessor" {
				preprocessed, err := core.ReadEmiFromString(ctx.Content)
				if err != nil {
					return err
				}
				preprocessed.StripActionUrlTypeAnnotations()

				yamlBytes, err := yaml.Marshal(preprocessed)
				if err != nil {
					return err
				}

				files = []core.VirtualFile{
					{
						Name:         "preprocessed",
						Extension:    ".yml",
						ActualScript: string(yamlBytes),
					},
				}
			} else {
				var action core.ActionFile

				switch target.Compiler {
				case "go":
					action = golang.GoPrimaryAction
				case "kotlin":
					action = kotlin.KotlinPrimaryAction
				case "swift":
					action = swift.SwiftPrimaryAction
				case "js":
					action = js.JsPrimaryAction
				default:
					continue
				}

				res, err := action.Run(ctx)
				if err != nil {
					fmt.Println(err)
					return err
				}
				files = res
			}

			if target.Output != "" {
				if path.IsAbs(target.Output) {
					outdir = target.Output
				} else {
					outdir = path.Join(path.Dir(c.String("path")), target.Output)
				}

				// Only ever applies when this target set its own Output (never to a
				// directory it merely inherited from a previous target that left
				// outdir unchanged), and only removes that directory itself - a
				// sibling target writing somewhere else is never touched.
				if target.Clean {
					if err := os.RemoveAll(outdir); err != nil {
						return fmt.Errorf("error cleaning output directory %v: %w", outdir, err)
					}
				}
			}

			for _, file := range files {

				filePath := path.Join(outdir, file.Location, file.Name+file.Extension)
				os.MkdirAll(path.Dir(filePath), os.ModePerm)

				if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
					return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
				}
			}

		}

		return nil
	},
}
