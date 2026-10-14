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
	ActionName string
	IsReactive bool
	// CliGroup is the kebab-case entity an entity-style action belongs to ("role" for
	// RoleBrowseAction), empty for actions that stand on their own. See cliGroupOf.
	CliGroup string
}

// crudSuffixes are the generated entity action verbs. Every entity has the same ones, so
// their command names (create, get, browse, ...) collide when many entities are bundled
// flat - ClientCliGroups nests them under the entity instead. Longest first.
var crudSuffixes = []string{"AwareDeletePreview", "AwareDelete", "Browse", "Create", "Update", "Get"}

// cliGroupOf returns the entity group of an action Go name such as "RoleBrowseAction"
// ("role"), or "" when the name is not an entity verb.
func cliGroupOf(actionName string) string {
	name := strings.TrimSuffix(actionName, "Action")
	for _, suffix := range crudSuffixes {
		if strings.HasSuffix(name, suffix) && len(name) > len(suffix) {
			return kebabCase(strings.TrimSuffix(name, suffix))
		}
	}
	return ""
}

func kebabCase(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('-')
			}
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}

// mcpManifestRender is one module intent a go-mcp manifest bundles: GoName names the
// Register<GoName>IntentTool function GoIntentsGenerate already generated for it (see
// go-intent-render.go), and ActionName is the resolved source action's own Go name
// (e.g. "RoleUpdateAction") - the exact handler function name
// Register<GoName>IntentTool's own signature expects, assumed already implemented in
// this same package, the same convention {{Name}}CliManifest/{{Name}}GinServerSetup
// already assume for {{.ActionName}} itself.
type mcpManifestRender struct {
	GoName     string
	ActionName string
}

func GoManifest(manifest core.EmiManifest, module *core.Emi, ctx core.MicroGenContext) (*core.CodeChunkCompiled, error) {

	const tmpl = `{{ if .cliOnly }}//go:build !wasm{{ end }}
/**
* Go manifest
*/


{{ if and .actions (not .sameLocation) }}
import (
	{{$.f.PackageName}} "{{ .mm }}"
)
{{ end }}

{{ if .goGin }}
func {{ upper .manifest.Name }}GinServerSetup(x *gin.RouterGroup) {
	{{ range .actions }}
		 {{$.location}} {{ .ActionName }}Gin(x, {{ .ActionName }})
	{{ end }}
}
{{ end }}

{{ if .goClient }}
// {{ upper .manifest.Name }}ClientCliBundle wraps each non-reactive action's
// {Action}CliHandler around a call to {Action}Call against the given API client, the
// same flags a server-side command would bind.
func {{ upper .manifest.Name }}ClientCliBundle(client emigo.APIClient) []*cli.Command {
	return []*cli.Command{
		{{ range .actions }}
		{{ if not .IsReactive }}
		{{$.location}} {{ .ActionName }}CliHandler(func(c {{$.location}} {{ .ActionName }}Request) (*{{$.location}} {{ .ActionName }}Response, error) {
			return {{$.location}} {{ .ActionName }}Call(c, &client)
		}),
		{{ end }}
		{{ end }}
	}
}

// {{ upper .manifest.Name }}ClientCliGroups is {{ upper .manifest.Name }}ClientCliBundle with the entity
// actions nested under their entity (role get, role browse, ...), since the same verbs
// exist for every entity and cannot sit side by side in one flat command list.
func {{ upper .manifest.Name }}ClientCliGroups(client emigo.APIClient) []*cli.Command {
	var out []*cli.Command
	groups := map[string]*cli.Command{}
	add := func(group string, cmd *cli.Command) {
		if group == "" {
			out = append(out, cmd)
			return
		}
		g, ok := groups[group]
		if !ok {
			g = &cli.Command{Name: group, Usage: "Actions of " + group}
			groups[group] = g
			out = append(out, g)
		}
		g.Commands = append(g.Commands, cmd)
	}
	{{ range .actions }}
	{{ if not .IsReactive }}
	add("{{ .CliGroup }}", {{$.location}} {{ .ActionName }}CliHandler(func(c {{$.location}} {{ .ActionName }}Request) (*{{$.location}} {{ .ActionName }}Response, error) {
		return {{$.location}} {{ .ActionName }}Call(c, &client)
	}))
	{{ end }}
	{{ end }}
	return out
}
{{ end }}

{{ if .goCli }}
// {{ upper .manifest.Name }}CliManifest bundles every action's *cli.Command: reactive
// actions go through {Action}CliReactiveHandler (stdin/stdout piping), everything else
// through {Action}CliHandler (body/path/query/header binding). Each developer implements
// {Action}(req) or {Action}(session) themselves (see the action file for the expected
// signature).
func {{ upper .manifest.Name }}CliManifest() []*cli.Command {
	return []*cli.Command{
		{{ range .actions }}
		{{ if .IsReactive }}
		{{$.location}} {{ .ActionName }}CliReactiveHandler({{ .ActionName }}),
		{{ else }}
		{{$.location}} {{ .ActionName }}CliHandler( {{ .ActionName }}),
		{{ end }}
		{{ end }}
	}
}
{{ end }}

{{ if .goMcp }}
// {{ upper .manifest.Name }}McpToolManifest registers this module's own intents (see
// Abac.emi.yml-style intents: blocks) onto reg as real emigo.Tool/emigo.ToolHandlerFn
// pairs, combining every Register<Name>IntentTool GoIntentsGenerate already generated
// for them (see Intents.go) - the same "just wire it all up" role
// {{ upper .manifest.Name }}CliManifest plays for CLI commands. Each intent's own
// resolved action handler (e.g. {{$.location}}RoleUpdateAction) is assumed already
// implemented in this package, same convention as every other bundle above. An
// intent with no From (no resolved action, so no Register...IntentTool was generated
// for it at all) is skipped here the same way GoIntentsGenerate itself skips it.
func {{ upper .manifest.Name }}McpToolManifest(reg *emigo.ToolRegistry) {
	{{ range .mcpIntents }}
	{{$.location}}Register{{ .GoName }}IntentTool(reg, {{ .ActionName }})
	{{ end }}
}
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
			ActionName: actionName,
			IsReactive: action.GetMethod() == "reactive",
			CliGroup:   cliGroupOf(actionName),
		})
	}

	mcpRendered := []mcpManifestRender{}

	for _, it := range module.Intents {
		if it == nil || it.Name == "" {
			continue
		}

		// No From means no resolved action, which means GoIntentsGenerate never
		// generated a Register<Name>IntentTool for it in the first place (see its
		// own doc comment) - nothing here to bundle.
		src := it.GetResolvedFrom()
		if src == nil {
			continue
		}

		if !shouldRenderAction(
			it.Name,
			manifest.Includes,
			manifest.Excludes,
		) {
			continue
		}

		mcpRendered = append(mcpRendered, mcpManifestRender{
			GoName:     it.Upper(),
			ActionName: src.GetName(),
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

	// When no dist is set, the manifest file lands directly alongside the actions it
	// wraps (same directory, same package) - importing that package from itself would
	// be a compile error, and the qualifier would be redundant, so both are dropped and
	// every reference below stays unqualified.
	sameLocation := manifest.Dist == ""

	location := fmt.Sprintf("%v.", f.PackageName)
	if sameLocation {
		location = ""
	}

	goClient := slices.Contains(manifest.Types, "go-client")
	goCli := slices.Contains(manifest.Types, "go-cli")
	goGin := slices.Contains(manifest.Types, "go-gin")
	goMcp := slices.Contains(manifest.Types, "go-mcp")

	if goClient || goCli {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: "github.com/urfave/cli/v3",
			},
		)
	}

	// emigo.APIClient is only referenced by the go-client bundle/command functions.
	if goClient && len(rendered) > 0 {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: f.Emigo,
			},
		)
	}

	if goGin {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: "github.com/gin-gonic/gin",
			},
		)
	}

	// emigo.ToolRegistry is only referenced by the go-mcp bundle function's own
	// signature - added even when mcpRendered ends up empty (an intent-less module
	// with go-mcp declared still gets a valid, if empty, McpToolManifest function).
	if goMcp {
		res.CodeChunkDependensies = append(
			res.CodeChunkDependensies,
			core.CodeChunkDependency{
				Location: f.Emigo,
			},
		)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"actions":      rendered,
		"mcpIntents":   mcpRendered,
		"location":     location,
		"sameLocation": sameLocation,
		"manifest":     manifest,
		"goClient":     goClient,
		"goCli":        goCli,
		"goGin":        goGin,
		"goMcp":        goMcp,
		// urfave/cli (and the generated {Action}CliHandler functions) are not built for wasm,
		// so a manifest with nothing but cli/client bundles is excluded there as a whole.
		"cliOnly": (goClient || goCli) && !goGin && !goMcp,
		"mm":      mm,
		"f":       f,
		"ctx":     ctx,
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

func MatchNameWithPattern(name string, patterns []string) bool {
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
	if len(includes) > 0 && !MatchNameWithPattern(name, includes) {
		return false
	}

	// exclude mode
	if MatchNameWithPattern(name, excludes) {
		return false
	}

	return true
}
