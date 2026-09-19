package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"
	"strings"
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

// intentBodyJSONSchema resolves one side (In/Out) of an intent into a JSON Schema
// object - a dto reference is looked up by class name in dtoByClassName and its own
// Fields walked, inline Fields are walked directly; either way it goes through
// core.EmiFieldsToObjectSchema, the same builder a query/path-param merge (see
// intentInputSchema) folds its own fields into. A nil/empty body (no dto, no inline
// fields, no primitive) yields nil - the caller treats "no schema fragment" as "add
// nothing", not as an empty object, so an action with no body at all (e.g. a plain
// browse whose only real arguments are query params) doesn't get a stray
// "type":"object" with zero properties as its own schema when it's the only source.
func intentBodyJSONSchema(body *core.EmiActionBody, dtoByClassName map[string]core.EmiDto) map[string]any {
	if body == nil {
		return nil
	}
	if body.Dto != "" {
		dto, ok := dtoByClassName[body.Dto]
		if !ok {
			return nil
		}
		return core.EmiFieldsToObjectSchema(dto.Fields)
	}
	if body.Primitive != "" {
		t, ok := map[string]string{
			"string": "string", "bool": "boolean", "int": "integer", "int32": "integer",
			"int64": "integer", "float32": "number", "float64": "number", "bytes": "string",
		}[body.Primitive]
		if !ok {
			return nil
		}
		return map[string]any{"type": t}
	}
	if len(body.Fields) > 0 {
		return core.EmiFieldsToObjectSchema(body.Fields)
	}
	return nil
}

// mergeObjectSchemas folds extra's own properties/required into base (an
// object-shaped schema from EmiFieldsToObjectSchema/intentBodyJSONSchema) in place -
// used to combine an intent's body schema with its resolved action's query/path-param
// fields into the one flat argument list a tool call actually presents to a model
// (which has no notion of "this argument came from the URL path" vs "the request
// body" - see intentInputSchema). base is created fresh (an empty object schema) when
// nil, so this also doubles as the "start a new schema" case for an intent with query/
// path params but no body at all (e.g. userBrowse's own filter/sort/pagination
// fields).
func mergeObjectSchemas(base map[string]any, extra map[string]any) map[string]any {
	if extra == nil {
		if base == nil {
			return nil
		}
		return base
	}
	if base == nil {
		base = map[string]any{"type": "object", "properties": map[string]any{}}
	}
	baseProps, _ := base["properties"].(map[string]any)
	if baseProps == nil {
		baseProps = map[string]any{}
		base["properties"] = baseProps
	}
	if extraProps, ok := extra["properties"].(map[string]any); ok {
		for k, v := range extraProps {
			baseProps[k] = v
		}
	}
	if extraRequired, ok := extra["required"].([]string); ok && len(extraRequired) > 0 {
		existing, _ := base["required"].([]string)
		base["required"] = append(existing, extraRequired...)
	}
	return base
}

// headerFieldsToObjectSchema converts a []*EmiHeader list into an object schema via
// core.EmiHeaderToEmiField + core.EmiFieldsToObjectSchema, or nil when there are none
// - the shared step intentInputSchema/intentOutputSchema both merge in.
func headerFieldsToObjectSchema(headers []core.EmiHeader) map[string]any {
	if len(headers) == 0 {
		return nil
	}
	var fields []*core.EmiField
	for i := range headers {
		if f := core.EmiHeaderToEmiField(&headers[i]); f != nil {
			fields = append(fields, f)
		}
	}
	if len(fields) == 0 {
		return nil
	}
	return core.EmiFieldsToObjectSchema(fields)
}

// intentInputSchema builds the complete JSON Schema a tool call actually needs for an
// intent's arguments: its own In (dto or inline fields) merged with its resolved
// source action's query parameters, URL path parameters, and typed request headers,
// if any (see EmiIntent.GetResolvedFrom, core.EmiQueryFieldToEmiField,
// core.EmiPathParamsToFields, core.EmiHeaderToEmiField). A model-facing tool call is a
// single flat argument object regardless of which of these four a given field happened
// to come from on the real HTTP route, so all four are merged into one schema here
// rather than left for a caller to reassemble. Returns "" (no schema at all) only when
// none of the four contributed anything - e.g. a plain action with no body, no query
// params, no path params, and no typed headers.
func intentInputSchema(it *core.EmiIntent, dtoByClassName map[string]core.EmiDto) (string, error) {
	schema := intentBodyJSONSchema(it.In, dtoByClassName)

	if src := it.GetResolvedFrom(); src != nil {
		var queryFields []*core.EmiField
		for _, q := range src.GetQuery() {
			if f := core.EmiQueryFieldToEmiField(q); f != nil {
				queryFields = append(queryFields, f)
			}
		}
		if len(queryFields) > 0 {
			schema = mergeObjectSchemas(schema, core.EmiFieldsToObjectSchema(queryFields))
		}

		if pathFields := core.EmiPathParamsToFields(src.GetUrl()); len(pathFields) > 0 {
			schema = mergeObjectSchemas(schema, core.EmiFieldsToObjectSchema(pathFields))
		}

		if headerSchema := headerFieldsToObjectSchema(src.GetRequestHeaders()); headerSchema != nil {
			schema = mergeObjectSchemas(schema, headerSchema)
		}
	}

	if schema == nil {
		return "", nil
	}
	b, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// intentOutputSchema builds an intent's response schema from Out plus its resolved
// source action's typed response headers, if any (see core.EmiHeaderToEmiField) -
// query/path params never describe a response, so unlike intentInputSchema those two
// are never merged in here.
func intentOutputSchema(it *core.EmiIntent, dtoByClassName map[string]core.EmiDto) (string, error) {
	schema := intentBodyJSONSchema(it.Out, dtoByClassName)

	if src := it.GetResolvedFrom(); src != nil {
		if headerSchema := headerFieldsToObjectSchema(src.GetResponseHeaders()); headerSchema != nil {
			schema = mergeObjectSchemas(schema, headerSchema)
		}
	}

	if schema == nil {
		return "", nil
	}
	b, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// bodyGoTypeName returns the Go type name of src's own <ActionName>Request.Body
// field - "" when the action has no body at all (Body then stays a bare
// interface{}, e.g. a browse-shaped action - see the generated Request struct's own
// template in go-action-render.go). Mirrors that exact same template's own
// RequestClassName derivation (go-action-realms.go): a dto reference names the
// dto's own class, inline fields get a generated "<ActionName>Req" struct - this
// has to match precisely, since renderIntentToolRegistration builds a real
// <ActionName>Request{Body: ...} literal against whichever type this returns.
func bodyGoTypeName(src *core.EmiAction) string {
	if src == nil || src.In == nil {
		return ""
	}
	if src.In.Dto != "" {
		return src.In.Dto
	}
	if len(src.In.Fields) > 0 {
		return src.GetName() + "Req"
	}
	return ""
}

// renderIntentToolRegistration builds a "Register<Name>IntentTool" function for one
// intent - the tool-calling equivalent of a generated "<Action>Gin"/"<Action>Cli"
// function: it takes the real handler function (the same one an HTTP route would
// use, e.g. abac.UserBrowseAction) and, from it, builds and registers a complete
// emigo.Tool + emigo.ToolHandlerFn pair on a project's own *emigo.ToolRegistry, with
// no hand-written JSON-unmarshal-into-the-right-DTO step or switch-on-tool-name
// dispatch left for a project to write (see AiTools.go's own history: this
// generator exists specifically to eliminate that file's hand-maintained
// boilerplate, one intent at a time, as this compiler grows to cover more of the
// same shapes *ActionGin already does).
//
// Only generated when it was derived via From from a real EmiAction (GetResolvedFrom)
// AND the intent didn't hand-override In/Out away from that action's own (same
// pointer identity check both places) - a hand-customized in/out on a from:-derived
// intent has no guaranteed correspondence to the actual handler function's real
// Request/Response shape, so auto-wiring it would silently produce a tool whose
// declared schema doesn't match what the handler actually expects/returns; ok is
// false in that case (and for a standalone intent with no From at all), meaning the
// caller should fall back to IntentMeta()-only and leave wiring to hand-written code,
// exactly as before this generator existed.
func renderIntentToolRegistration(goName string, it *core.EmiIntent, ctx core.MicroGenContext) (code string, deps []core.CodeChunkDependency, ok bool, err error) {
	src := it.GetResolvedFrom()
	if src == nil {
		return "", nil, false, nil
	}
	if it.In != src.In || it.Out != src.Out {
		return "", nil, false, nil
	}

	actionName := src.GetName()
	bodyType := bodyGoTypeName(src)

	pathRealms, err := GoActionPathParamsRealms(src, ctx)
	if err != nil {
		return "", nil, false, err
	}

	var b strings.Builder

	fmt.Fprintf(&b, "\n// Register%sIntentTool registers the %q intent as a tool on reg, dispatching\n", goName, it.Name)
	fmt.Fprintf(&b, "// to handler exactly as %sGin dispatches to it for its own HTTP route - built\n", actionName)
	fmt.Fprintf(&b, "// from the same %sIntentMeta() this reuses. A malformed call or a handler error is\n", goName)
	fmt.Fprintf(&b, "// reported back as a failed tool result rather than propagated, so one bad call\n")
	fmt.Fprintf(&b, "// can't take down whatever loop is iterating tool calls.\n")
	fmt.Fprintf(&b, "func Register%sIntentTool(reg *emigo.ToolRegistry, handler func(%sRequest) (*%sResponse, error)) {\n", goName, actionName, actionName)
	fmt.Fprintf(&b, "\tmeta := %sIntentMeta()\n", goName)
	fmt.Fprintf(&b, "\treg.Register(emigo.Tool{\n")
	fmt.Fprintf(&b, "\t\tName:         meta.Name,\n")
	fmt.Fprintf(&b, "\t\tDescription:  meta.Description,\n")
	fmt.Fprintf(&b, "\t\tInputSchema:  []byte(meta.InputSchema),\n")
	fmt.Fprintf(&b, "\t\tOutputSchema: []byte(meta.OutputSchema),\n")
	fmt.Fprintf(&b, "\t}, func(ctx context.Context, req emigo.EmiRequestContexts, call emigo.ToolCall) (string, bool) {\n")
	fmt.Fprintf(&b, "\t\targs := map[string]interface{}{}\n")
	fmt.Fprintf(&b, "\t\tif len(call.Input) > 0 {\n")
	fmt.Fprintf(&b, "\t\t\tif err := json.Unmarshal(call.Input, &args); err != nil {\n")
	fmt.Fprintf(&b, "\t\t\t\treturn \"invalid arguments: \" + err.Error(), true\n")
	fmt.Fprintf(&b, "\t\t\t}\n")
	fmt.Fprintf(&b, "\t\t}\n\n")

	fmt.Fprintf(&b, "\t\tqueryParams := url.Values{}\n")
	for _, q := range src.GetQuery() {
		if q == nil || q.Name == "" {
			continue
		}
		fmt.Fprintf(&b, "\t\tif v, ok := args[%q]; ok {\n", q.Name)
		fmt.Fprintf(&b, "\t\t\tqueryParams.Set(%q, fmt.Sprintf(\"%%v\", v))\n", q.Name)
		fmt.Fprintf(&b, "\t\t}\n")
		fmt.Fprintf(&b, "\t\tdelete(args, %q)\n", q.Name)
	}
	b.WriteString("\n")

	hasPathParams := pathRealms != nil && len(pathRealms.Params) > 0
	if hasPathParams {
		fmt.Fprintf(&b, "\t\tvar pathParams %s\n", pathRealms.TypeName)
		for _, p := range pathRealms.Params {
			fmt.Fprintf(&b, "\t\tif v, ok := args[%q]; ok {\n", p.PlaceHolderValue)
			switch p.GolangType {
			case "string":
				fmt.Fprintf(&b, "\t\t\tpathParams.%s = fmt.Sprintf(\"%%v\", v)\n", p.GolangFieldName)
			case "bool":
				fmt.Fprintf(&b, "\t\t\tparsed, parseErr := strconv.ParseBool(fmt.Sprintf(\"%%v\", v))\n")
				fmt.Fprintf(&b, "\t\t\tif parseErr != nil {\n")
				fmt.Fprintf(&b, "\t\t\t\treturn \"invalid \\\"%s\\\" argument: \" + parseErr.Error(), true\n", p.PlaceHolderValue)
				fmt.Fprintf(&b, "\t\t\t}\n")
				fmt.Fprintf(&b, "\t\t\tpathParams.%s = parsed\n", p.GolangFieldName)
			case "float32", "float64":
				fmt.Fprintf(&b, "\t\t\tparsed, parseErr := strconv.ParseFloat(fmt.Sprintf(\"%%v\", v), 64)\n")
				fmt.Fprintf(&b, "\t\t\tif parseErr != nil {\n")
				fmt.Fprintf(&b, "\t\t\t\treturn \"invalid \\\"%s\\\" argument: \" + parseErr.Error(), true\n", p.PlaceHolderValue)
				fmt.Fprintf(&b, "\t\t\t}\n")
				fmt.Fprintf(&b, "\t\t\tpathParams.%s = %s(parsed)\n", p.GolangFieldName, p.GolangType)
			default:
				// int, int32, int64, uint, uint32, uint64, uint8, uint16, ...
				fmt.Fprintf(&b, "\t\t\tparsed, parseErr := strconv.ParseInt(fmt.Sprintf(\"%%v\", v), 10, 64)\n")
				fmt.Fprintf(&b, "\t\t\tif parseErr != nil {\n")
				fmt.Fprintf(&b, "\t\t\t\treturn \"invalid \\\"%s\\\" argument: \" + parseErr.Error(), true\n", p.PlaceHolderValue)
				fmt.Fprintf(&b, "\t\t\t}\n")
				fmt.Fprintf(&b, "\t\t\tpathParams.%s = %s(parsed)\n", p.GolangFieldName, p.GolangType)
			}
			fmt.Fprintf(&b, "\t\t}\n")
			fmt.Fprintf(&b, "\t\tdelete(args, %q)\n", p.PlaceHolderValue)
		}
		b.WriteString("\n")
	}

	if bodyType != "" {
		fmt.Fprintf(&b, "\t\tbodyJSON, err := json.Marshal(args)\n")
		fmt.Fprintf(&b, "\t\tif err != nil {\n")
		fmt.Fprintf(&b, "\t\t\treturn \"invalid arguments: \" + err.Error(), true\n")
		fmt.Fprintf(&b, "\t\t}\n")
		fmt.Fprintf(&b, "\t\tvar body %s\n", bodyType)
		fmt.Fprintf(&b, "\t\tif err := json.Unmarshal(bodyJSON, &body); err != nil {\n")
		fmt.Fprintf(&b, "\t\t\treturn \"invalid arguments: \" + err.Error(), true\n")
		fmt.Fprintf(&b, "\t\t}\n\n")
	}

	fmt.Fprintf(&b, "\t\tresp, err := handler(%sRequest{\n", actionName)
	fmt.Fprintf(&b, "\t\t\tGinCtx:      req.GetGinCtx(),\n")
	fmt.Fprintf(&b, "\t\t\tQueryParams: queryParams,\n")
	if hasPathParams {
		fmt.Fprintf(&b, "\t\t\tParams:      pathParams,\n")
	}
	if bodyType != "" {
		fmt.Fprintf(&b, "\t\t\tBody:        body,\n")
	}
	fmt.Fprintf(&b, "\t\t})\n")
	fmt.Fprintf(&b, "\t\tif err != nil {\n")
	fmt.Fprintf(&b, "\t\t\treturn err.Error(), true\n")
	fmt.Fprintf(&b, "\t\t}\n")
	fmt.Fprintf(&b, "\t\tb, err := json.Marshal(resp.Payload)\n")
	fmt.Fprintf(&b, "\t\tif err != nil {\n")
	fmt.Fprintf(&b, "\t\t\treturn err.Error(), true\n")
	fmt.Fprintf(&b, "\t\t}\n")
	fmt.Fprintf(&b, "\t\treturn string(b), false\n")
	fmt.Fprintf(&b, "\t})\n")
	fmt.Fprintf(&b, "}\n")

	deps = []core.CodeChunkDependency{
		{Location: "context"},
		{Location: "encoding/json"},
		{Location: "fmt"},
		{Location: "net/url"},
		{Location: "github.com/torabian/emi/emigo"},
	}
	if hasPathParams {
		deps = append(deps, pathRealms.Dependencies...)
	}

	return b.String(), deps, true, nil
}

// GoIntentsGenerate renders every module-level intent's tool signature. The first
// returned chunk is always "Intents.go": an "<Name>IntentArgs"/"<Name>IntentResult"
// struct pair per intent that declares inline fields (a Dto/primitive reference reuses
// the existing type instead - see intentBodyType), plus an "<Name>IntentMeta()"
// accessor carrying the tool name/title/description/annotations *and* a ready-to-use
// InputSchema/OutputSchema JSON Schema string (see intentInputSchema/
// intentOutputSchema) - enough for a server (e.g. fireback's own in-process tool-use
// wiring, or a real MCP server) to build a concrete tool registration from without
// redeclaring any of this by hand, whether the intent's In/Out was a dto reference,
// inline fields, or (for input) a resolved action's own query/path params. Any further
// chunks are per-type CLI helper files (only present when the "split-cli" tag is set -
// see intentBodyType), each already carrying its own `//go:build !wasm` guard and
// needing its own file for that guard to apply correctly. Wiring these into an actual
// MCP server/SDK is deliberately left to the caller; this only generates the typed
// shapes and metadata.
func GoIntentsGenerate(
	module *core.Emi,
	ctx core.MicroGenContext,
	emigoImportPath string,
	complexes []RecognizedComplex,
) ([]*core.CodeChunkCompiled, error) {

	if module == nil || len(module.Intents) == 0 {
		return nil, nil
	}
	intents := module.Intents

	dtoByClassName := make(map[string]core.EmiDto, len(module.Dto))
	for i := range module.Dto {
		dtoByClassName[module.Dto[i].GetClassName()] = module.Dto[i]
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
		// InputSchema/OutputSchema are JSON Schema object literals (or "" when
		// there's nothing to describe) - see intentInputSchema/intentOutputSchema.
		// Rendered as Go raw string literals (backtick-quoted): JSON's own quoting
		// never needs a literal backtick, so this needs no escaping the way a
		// double-quoted Go string embedding arbitrary JSON would.
		InputSchema  string
		OutputSchema string
	}

	var (
		structChunks      []*core.CodeChunkCompiled
		cliChunks         []*core.CodeChunkCompiled
		deps              []core.CodeChunkDependency
		rendered          []renderedIntent
		toolRegistrations []string
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
		inputSchema, err := intentInputSchema(it, dtoByClassName)
		if err != nil {
			return nil, err
		}
		outputSchema, err := intentOutputSchema(it, dtoByClassName)
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
			GoName:       goName,
			ToolName:     it.Name,
			ArgsType:     argsType,
			ResultType:   resultType,
			Title:        it.Title,
			Description:  it.Description,
			Permission:   it.Permission,
			InputSchema:  inputSchema,
			OutputSchema: outputSchema,
		}
		if it.Annotations != nil {
			r.ReadOnlyHint = it.Annotations.ReadOnlyHint
			r.DestructiveHint = it.Annotations.DestructiveHint
			r.IdempotentHint = it.Annotations.IdempotentHint
			r.OpenWorldHint = it.Annotations.OpenWorldHint
		}
		rendered = append(rendered, r)

		toolCode, toolDeps, ok, err := renderIntentToolRegistration(goName, it, ctx)
		if err != nil {
			return nil, err
		}
		if ok {
			toolRegistrations = append(toolRegistrations, toolCode)
			deps = append(deps, toolDeps...)
		}
	}

	const tmpl = `/**
* Model-facing tool signatures, generated from this module's intents list. See
* EmiIntent in lib/core for the source definitions - each one below either declared its
* own in/out inline, or was derived from (or auto-created by "intent: true" on) an
* existing action. InputSchema/OutputSchema also fold in that source action's own query
* parameters and URL path parameters, when it has any - a tool call is a single flat
* argument object regardless of which part of the real HTTP route a field came from.
*/

{{ range .intents }}
// {{ .GoName }}IntentMeta describes the "{{ .ToolName }}" MCP tool intent - name,
// title, description, behavior-hint annotations, and a ready-to-use JSON Schema for
// both the tool's arguments (InputSchema) and its response (OutputSchema, when the
// source has one) - so a server can build a concrete tool registration from it without
// redeclaring any of this by hand.
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
	InputSchema     string
	OutputSchema    string
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
		InputSchema     string
		OutputSchema    string
	}{
		Name:            "{{ .ToolName }}",
		Title:           "{{ escape .Title }}",
		Description:     "{{ escape .Description }}",
		ReadOnlyHint:    {{ .ReadOnlyHint }},
		DestructiveHint: {{ .DestructiveHint }},
		IdempotentHint:  {{ .IdempotentHint }},
		OpenWorldHint:   {{ .OpenWorldHint }},
		Permission:      "{{ escape .Permission }}",
		InputSchema:     ` + "`{{ .InputSchema }}`" + `,
		OutputSchema:    ` + "`{{ .OutputSchema }}`" + `,
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

	for _, code := range toolRegistrations {
		main.ActualScript = append(main.ActualScript, '\n')
		main.ActualScript = append(main.ActualScript, []byte(code)...)
	}

	files := []*core.CodeChunkCompiled{main}
	for _, chunk := range cliChunks {
		if chunk == nil {
			continue
		}
		files = append(files, chunk)
	}

	return files, nil
}
