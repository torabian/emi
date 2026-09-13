package golang

import (
	"bytes"
	"maps"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

// intentBodyType resolves the Go type name for one side (In/Out) of an intent, and any
// extra struct chunk(s) that need rendering alongside it: a Dto reference already names
// an existing Go type (EmiActionBody.Dto is conventionally the dto's *class* name
// already, e.g. "UserDto" - see how preprocessEntityActions/BuildEntityDto populate
// it), while inline Fields need a fresh struct generated for them, the same way
// GoModuleFull renders a plain dto's fields via GoCommonStructGenerator - complexes is
// threaded through unchanged so a `type: complex` field resolves to its real
// import-qualified type here exactly as it would on a dto or action body.
//
// cliChunk is only non-nil when the "split-cli" tag is set: GoCommonStructGenerator
// then keeps the CLI flag/cast helpers out of mainChunk (each carrying its own
// `//go:build !wasm` guard - see GoCommonStructGeneratorCli), so unlike mainChunk it
// can never be folded into the shared Intents.go body; the caller must emit it as its
// own file, same as the dto loop does for actionRendered.CliHelpers.
func intentBodyType(
	body *core.EmiActionBody,
	className string,
	ctx core.MicroGenContext,
	emigoImportPath string,
	complexes []RecognizedComplex,
) (typeName string, mainChunk *core.CodeChunkCompiled, cliChunk *core.CodeChunkCompiled, deps []core.CodeChunkDependency, err error) {

	if body == nil {
		return "", nil, nil, nil, nil
	}

	if body.Dto != "" {
		return body.Dto, nil, nil, nil, nil
	}

	if body.Primitive != "" {
		return goPrimitiveType(body.Primitive), nil, nil, nil, nil
	}

	if len(body.Fields) == 0 {
		return "", nil, nil, nil, nil
	}

	rendered, err := GoCommonStructGenerator(body.Fields, ctx, GoCommonStructContext{
		RootClassName:       className,
		EmiLocation:         emigoImportPath,
		RecognizedComplexes: complexes,
	})
	if err != nil {
		return "", nil, nil, nil, err
	}

	if rendered.MainClass != nil {
		deps = rendered.MainClass.CodeChunkDependensies
	}

	return className, rendered.MainClass, rendered.CliHelpers, deps, nil
}

// goPrimitiveType maps an EmiActionBody.Primitive value to its Go type - the same set
// EmiActionBody.Primitive's own jsonschema enum allows.
func goPrimitiveType(primitive string) string {
	switch primitive {
	case "bytes":
		return "[]byte"
	default:
		return primitive
	}
}

// GoIntentsGenerate renders every module-level intent's tool signature. The first
// returned chunk is always "Intents.go": an "<Name>IntentArgs"/"<Name>IntentResult"
// struct pair per intent that declares inline fields (a Dto/primitive reference reuses
// the existing type instead - see intentBodyType), plus an "<Name>IntentMeta()"
// accessor carrying the tool name/title/description/annotations - enough for a server
// (e.g. fireback's McpCli.go) to build a concrete MCP tool registration from without
// redeclaring any of this by hand. Any further chunks are per-type CLI helper files
// (only present when the "split-cli" tag is set - see intentBodyType), each already
// carrying its own `//go:build !wasm` guard and needing its own file for that guard to
// apply correctly. Wiring these into an actual MCP server/SDK is deliberately left to
// the caller; this only generates the typed shapes and metadata.
func GoIntentsGenerate(
	intents []*core.EmiIntent,
	ctx core.MicroGenContext,
	emigoImportPath string,
	complexes []RecognizedComplex,
) ([]*core.CodeChunkCompiled, error) {

	if len(intents) == 0 {
		return nil, nil
	}

	type renderedIntent struct {
		// GoName is the UpperCamelCase identifier prefix ("SearchInvoices" ->
		// SearchInvoicesIntentMeta/Args/Result) - always Go-identifier-safe.
		GoName string
		// ToolName is the literal name reported to the model/MCP client, exactly as
		// declared in the intent's own Name (e.g. "searchInvoices").
		ToolName        string
		ArgsType        string
		ResultType      string
		Title           string
		Description     string
		ReadOnlyHint    bool
		DestructiveHint bool
		IdempotentHint  bool
		OpenWorldHint   bool
		Permission      string
	}

	var (
		structChunks []*core.CodeChunkCompiled
		cliChunks    []*core.CodeChunkCompiled
		deps         []core.CodeChunkDependency
		rendered     []renderedIntent
	)

	for _, it := range intents {
		if it == nil || it.Name == "" {
			continue
		}

		goName := it.Upper()

		argsType, argsChunk, argsCliChunk, argsDeps, err := intentBodyType(it.In, goName+"IntentArgs", ctx, emigoImportPath, complexes)
		if err != nil {
			return nil, err
		}
		resultType, resultChunk, resultCliChunk, resultDeps, err := intentBodyType(it.Out, goName+"IntentResult", ctx, emigoImportPath, complexes)
		if err != nil {
			return nil, err
		}

		if argsChunk != nil {
			structChunks = append(structChunks, argsChunk)
		}
		if resultChunk != nil {
			structChunks = append(structChunks, resultChunk)
		}
		if argsCliChunk != nil {
			cliChunks = append(cliChunks, argsCliChunk)
		}
		if resultCliChunk != nil {
			cliChunks = append(cliChunks, resultCliChunk)
		}
		deps = append(deps, argsDeps...)
		deps = append(deps, resultDeps...)

		r := renderedIntent{
			GoName:      goName,
			ToolName:    it.Name,
			ArgsType:    argsType,
			ResultType:  resultType,
			Title:       it.Title,
			Description: it.Description,
			Permission:  it.Permission,
		}
		if it.Annotations != nil {
			r.ReadOnlyHint = it.Annotations.ReadOnlyHint
			r.DestructiveHint = it.Annotations.DestructiveHint
			r.IdempotentHint = it.Annotations.IdempotentHint
			r.OpenWorldHint = it.Annotations.OpenWorldHint
		}
		rendered = append(rendered, r)
	}

	const tmpl = `/**
* Model-facing tool signatures, generated from this module's intents list. See
* EmiIntent in lib/core for the source definitions - each one below either declared its
* own in/out inline, or was derived from (or auto-created by "intent: true" on) an
* existing action.
*/

{{ range .intents }}
// {{ .GoName }}IntentMeta describes the "{{ .ToolName }}" MCP tool intent - name,
// title, description and MCP behavior-hint annotations - so a server can build a
// concrete tool registration from it without redeclaring any of this by hand.
{{ if .ArgsType }}// Input shape: {{ .ArgsType }}.{{ end }}
{{ if .ResultType }}// Output shape: {{ .ResultType }}.{{ end }}
func {{ .GoName }}IntentMeta() struct {
	Name            string
	Title           string
	Description     string
	ReadOnlyHint    bool
	DestructiveHint bool
	IdempotentHint  bool
	OpenWorldHint   bool
	Permission      string
} {
	return struct {
		Name            string
		Title           string
		Description     string
		ReadOnlyHint    bool
		DestructiveHint bool
		IdempotentHint  bool
		OpenWorldHint   bool
		Permission      string
	}{
		Name:            "{{ .ToolName }}",
		Title:           "{{ escape .Title }}",
		Description:     "{{ escape .Description }}",
		ReadOnlyHint:    {{ .ReadOnlyHint }},
		DestructiveHint: {{ .DestructiveHint }},
		IdempotentHint:  {{ .IdempotentHint }},
		OpenWorldHint:   {{ .OpenWorldHint }},
		Permission:      "{{ escape .Permission }}",
	}
}
{{ end }}

// IntentNames lists every intent's tool name this module contributes, for a server to
// enumerate/register (e.g. iterating this to look up each "<GoName>IntentMeta").
func IntentNames() []string {
	return []string{
		{{ range .intents }}"{{ escape .ToolName }}",
		{{ end }}
	}
}
`

	funcs := template.FuncMap{}
	maps.Copy(funcs, core.CommonMap)

	t := template.Must(template.New("intents").Funcs(funcs).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{"intents": rendered}); err != nil {
		return nil, err
	}

	main := &core.CodeChunkCompiled{
		SuggestedFileName:     "Intents",
		SuggestedExtension:    ".go",
		CodeChunkDependensies: deps,
	}

	for _, chunk := range structChunks {
		if chunk == nil {
			continue
		}
		main.ActualScript = append(main.ActualScript, chunk.ActualScript...)
		main.ActualScript = append(main.ActualScript, '\n')
		main.CodeChunkDependensies = append(main.CodeChunkDependensies, chunk.CodeChunkDependensies...)
	}

	main.ActualScript = append(main.ActualScript, buf.Bytes()...)

	files := []*core.CodeChunkCompiled{main}
	for _, chunk := range cliChunks {
		if chunk == nil {
			continue
		}
		files = append(files, chunk)
	}

	return files, nil
}
