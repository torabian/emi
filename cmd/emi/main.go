package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/golang"
	"github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"
	"github.com/torabian/emi/lib/querypredict"

	"github.com/urfave/cli"
)

func main() {

	commands := []cli.Command{
		DirCommand,
		GenerateCommand,
	}
	commands = append(commands,
		cliCommandFromTextActions(js.GetJsPublicActions().TextActions)...)
	commands = append(commands,
		cliCommandFromFileActions(js.GetJsPublicActions().FileActions)...)

	commands = append(commands,
		cliCommandFromTextActions(querypredict.GetQPPublicActions().TextActions)...)
	commands = append(commands,
		cliCommandFromFileActions(querypredict.GetQPPublicActions().FileActions)...)

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
		m[flag.Name] = c.String(flag.Name)
	}

	res, _ := json.Marshal(m)

	ctx.Flags = string(res)
	ctx.Content = string(content)

	return ctx, nil
}

type MicroGenContext struct {
	Tags     string // Tags/features to enable or disable
	Output   string // Output file or directory
	Flags    string
	Content  string
	Document querypredict.QueryDocument
}

func createCliContextQp(c *cli.Context) (MicroGenContext, error) {
	ctx := MicroGenContext{
		Tags:   c.String("tags"),
		Output: c.String("output"),
	}

	content, _ := os.ReadFile(c.String("path"))
	var m map[string]string = map[string]string{}

	document, err := querypredict.LoadQueriesFromString(string(content))

	if err != nil {
		fmt.Printf("expected no error, got %v", err)
		return ctx, err
	}

	ctx.Document = *document

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

var GenerateCommand = cli.Command{
	Name:        "qp:gen",
	Description: "Generate the query predict golang files from a query predict yaml definition",
	Usage:       "Generate the query predict golang files from a query predict yaml definition",
	Flags:       commonFlags,
	Action: func(c *cli.Context) error {
		ctx, err := createCliContextQp(c)
		if err != nil {
			return err
		}

		files, err := querypredict.ProcessQueryPredicts(ctx.Document)
		if err != nil {
			return err
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

var DirCommand = cli.Command{
	Name:        "qp:dir",
	Description: "Searches for .sql files in given directory, considering maximum depth, and would generate querypredict golang files in output ",
	Usage:       "Searches for .sql files in given directory, considering maximum depth, and would generate querypredict golang files in output ",
	Flags:       commonFlags,
	Action: func(c *cli.Context) error {

		files := []core.VirtualFile{}
		doc := querypredict.QueryDocument{}

		err2 := ReadSQLFiles(DiskFS{Root: c.String("path")}, ".", 1, func(filePath string, data []byte) error {

			doc.Queries = append(doc.Queries, querypredict.QuerySpec{

				Name:  strings.ReplaceAll(path.Base(filePath), ".sql", ""),
				Query: string(data),
			})

			return nil
		})

		if err2 != nil {
			return err2
		}

		files, err := querypredict.ProcessQueryPredicts(doc)
		if err != nil {
			return err
		}

		for _, file := range files {

			filePath := path.Join(c.String("output"), file.Location, file.Name+file.Extension)
			os.MkdirAll(path.Dir(filePath), os.ModePerm)

			if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
				return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
			}
		}

		for _, file := range files {

			filePath := path.Join(c.String("output"), file.Location, file.Name+file.Extension)
			os.MkdirAll(path.Dir(filePath), os.ModePerm)

			if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
				return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
			}
		}

		return nil
	},
}

// FS defines minimal interface for reading files and walking directories
type FS interface {
	ReadDir(dirname string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

// DiskFS implements FS for the OS file system
type DiskFS struct {
	Root string
}

func (d DiskFS) ReadDir(dirname string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(d.Root, dirname))
}

func (d DiskFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(d.Root, name))
}

// EmbedFS implements FS for embed.FS
type EmbedFS struct {
	FS   embed.FS
	Root string
}

func (e EmbedFS) ReadDir(dirname string) ([]fs.DirEntry, error) {
	return e.FS.ReadDir(filepath.Join(e.Root, dirname))
}

func (e EmbedFS) ReadFile(name string) ([]byte, error) {
	return e.FS.ReadFile(filepath.Join(e.Root, name))
}

// ReadSQLFiles walks the FS and reads all .sql files up to maxDepth
func ReadSQLFiles(fsys FS, root string, maxDepth int, reader func(path string, data []byte) error) error {
	var walk func(path string, depth int) error
	walk = func(path string, depth int) error {
		if maxDepth >= 0 && depth > maxDepth {
			return nil
		}
		entries, err := fsys.ReadDir(path)
		if err != nil {
			return err
		}
		for _, e := range entries {
			p := filepath.Join(path, e.Name())
			if e.IsDir() {
				if err := walk(p, depth+1); err != nil {
					return err
				}
			} else if strings.HasSuffix(e.Name(), ".sql") {
				data, err := fsys.ReadFile(p)
				if err != nil {
					return err
				}
				if err := reader(p, data); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return walk(root, 0)
}
