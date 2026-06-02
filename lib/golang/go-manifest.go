package golang

import (
	"bytes"
	"fmt"
	"path"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type manifestRender struct {
	ActionName      string
	HasAlias        bool
	HasRequestFlags bool
}

func GoManifest(manifest core.EmiManifest, module *core.Emi, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	const tmpl = `	
/**
* Go manifest
*/


{{ if .actions }}
import (
	{{$.f.PackageName}} "{{ .mm }}"
)
{{ end }}

{{ if .goGin }}
func {{ upper .manifest.Name }}GinServerSetup(x *gin.Engine) {
	{{ range .actions }}
		 {{$.location}} {{ .ActionName }}Gin(x, {{ .ActionName }})
	{{ end }}
}
{{ end }}

{{ if .goClient }}
func {{ upper .manifest.Name }}ClientCliBundle(client emigo.APIClient) []*cli.Command {
	return []*cli.Command{
		{{ range .actions }}
		Get{{ .ActionName }}ClientCmd(client),
		{{ end }}
	}
}
{{ end }}

{{ if .goCli }}
func {{ upper .manifest.Name }}CliManifest() []*cli.Command {

	return []*cli.Command{
		{{ range .actions }}
			Get{{ .ActionName }}Cmd(),
		{{ end }}
	}
}
{{ end }}


{{ if .goCli }}
	{{ range .actions }}

	func Get{{ .ActionName }}Cmd() *cli.Command {
		return &cli.Command{
			Name:    {{$.location}} {{ .ActionName }}Meta().CliName,
			Usage:   {{$.location}} {{ .ActionName }}Meta().Description,

			{{ if .HasAlias }}
			Aliases: []string{
				{{$.location}} {{ .ActionName }}Meta().CliShort,
			},
			{{ end }}

			{{ if .HasRequestFlags }}
			Flags:   emigo.CastEmiFlagToUrfave({{$.location}} Get{{ .ActionName }}ReqCliFlags("")),
			{{ end }}
			Action: func(ctx context.Context, c *cli.Command) error {
				return emigo.HandleActionInCli({{ .ActionName }}(
					{{$.location}} {{ .ActionName }}Request{
						CliCtx: c,

						{{ if .HasRequestFlags }}
						Body:   {{$.location}} Cast{{ .ActionName }}ReqFromCli(c),
						{{ end }}
					},
				))
			},
		}
	}

	{{ end }}
{{ end }}

{{ if .goClient }}
	{{ range .actions }}
	func Get{{ .ActionName }}ClientCmd(client emigo.APIClient) *cli.Command {
		return &cli.Command{
			Name:    {{$.location}}{{ .ActionName }}Meta().CliName,
			Usage:   {{$.location}}{{ .ActionName }}Meta().Description,

			{{ if .HasAlias }}
			Aliases: []string{
				{{$.location}}{{ .ActionName }}Meta().CliShort,
			},
			{{ end }}

			{{ if .HasRequestFlags }}
			Flags:   emigo.CastEmiFlagToUrfave({{$.location}}Get{{ .ActionName }}ReqCliFlags("")),
			{{ end }}
			Action: func(ctx context.Context, c *cli.Command) error {
				return emigo.HandleActionInCli({{$.location}}{{ .ActionName }}Call(
					{{$.location}}{{ .ActionName }}Request{
						CliCtx: c,

						{{ if .HasRequestFlags }}
						Body:   {{$.location}}Cast{{ .ActionName }}ReqFromCli(c),
						{{ end }}
					},
					&client,
				))
			},
		}
	}
	{{ end }}
{{ end }}

`

	f := GetCommonFlags(ctx)

	rendered := []manifestRender{}

	for _, action := range module.Actions {

		act, _ := GoActionRender(action, ctx, []RecognizedComplex{})
		if len(act) == 0 {
			continue
		}

		actionName := core.FindTokenByName(
			act[0].Tokens,
			core.TOKEN_ORIGINAL_NAME,
		).Value

		if !shouldRenderAction(
			action.Name,
			manifest.Includes,
			manifest.Excludes,
		) {
			continue
		}

		rendered = append(rendered, manifestRender{
			ActionName:      actionName,
			HasAlias:        action.GetCliShort() != "",
			HasRequestFlags: action.HasRequestFields(),
		})
	}

	t := template.Must(template.New("go_manifest").Funcs(core.CommonMap).Parse(tmpl))
	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  "PACKAGE_NAME",
				Value: manifest.Package,
			},
		},
		CodeChunkDependensies: []core.CodeChunkDependency{},
	}

	mm := path.Join(manifest.ModPackageName, ctx.Output)

	goClient := slices.Contains(manifest.Types, "go-client")
	goCli := slices.Contains(manifest.Types, "go-cli")
	goGin := slices.Contains(manifest.Types, "go-gin")

	if goClient || goCli {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: "github.com/urfave/cli/v3",
			},
		)

		if len(rendered) > 0 {
			res.CodeChunkDependensies = append(
				res.CodeChunkDependensies,
				core.CodeChunkDependency{
					Location: "context",
				},
				core.CodeChunkDependency{
					Location: f.Emigo,
				},
			)
		}
	}

	if goGin {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: "github.com/gin-gonic/gin",
			},
		)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"actions":  rendered,
		"location": fmt.Sprintf("%v.", f.PackageName),
		"manifest": manifest,
		"goClient": goClient,
		"goCli":    goCli,
		"goGin":    goGin,
		"mm":       mm,
		"f":        f,
		"ctx":      ctx,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedExtension = ".go"
	res.SuggestedFileName = path.Join(
		manifest.Dist,
		fmt.Sprintf("%vManifest", core.ToUpper(manifest.Name)),
	)

	return res, nil
}

func matchAction(name string, patterns []string) bool {
	for _, p := range patterns {

		// convenience wildcard
		if p == "*" {
			return true
		}

		if name == p {
			return true
		}

		// convert shell-like wildcard
		// User* -> User.*
		p = strings.ReplaceAll(p, "*", ".*")

		ok, err := regexp.MatchString("^"+p+"$", name)
		if err == nil && ok {
			return true
		}
	}

	return false
}

func shouldRenderAction(name string, includes, excludes []string) bool {

	// include mode
	if len(includes) > 0 && !matchAction(name, includes) {
		return false
	}

	// exclude mode
	if matchAction(name, excludes) {
		return false
	}

	return true
}
