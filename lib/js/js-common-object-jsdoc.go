// Renders the common object (dtos, entities, action bodies, ...) as JSDoc
// `@typedef` blocks - a plain JavaScript file gets no static type information at
// all otherwise (the generated class itself only has private `#fields` behind
// getters/setters, see js-common-object-class.go), so editors relying on the
// TypeScript-powered JS language service (VS Code's default "JS/TS" support,
// including plain .js files) have nothing to offer intellisense from unless
// --tags typescript is set.
//
// This is a self-contained, separate section of the generated file - its own
// `/** @typedef ... */` comment blocks, added alongside (not instead of) the
// runtime class - and is skipped entirely in TypeScript mode, where a real
// `export type` already exists (see js-common-object-types.go's
// TsCommonObjectGenerator) and a second, JSDoc-flavored description of the same
// shape would just be redundant. See the isTypeScript/isJsDoc wiring in
// JsCommonObjectGenerator, js-common-object.go.
//
// Deliberately mirrors js-common-object-types.go's recursion almost field-for-
// field (same treeLocation/childName convention, and directly reuses
// TsComputedField for the type string of each field) - TypeScript's own JSDoc
// support parses `{...}` type expressions with its regular TS type parser, so
// the exact same type strings the real `.ts` types use are valid here too, and
// a nested `@typedef {Object} WidgetDtoType.NestedType` is understood as a
// namespaced member the same way `tsFieldType`'s `WidgetDtoType.NestedType`
// property-type reference already expects.
package js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

type jsdocRenderedField struct {
	Line string
}

// oneLine collapses a (possibly multi-line, user-authored) description down to
// a single line - a raw newline inside a `* ...` jsdoc comment line would
// otherwise break out of the comment's `*`-per-line convention.
func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	return strings.TrimSpace(s)
}

func jsdocRenderField(field *core.EmiField, parentChain string) jsdocRenderedField {
	typeStr := tsFieldType(field, parentChain)

	name := field.PrivateName()
	if IsNullable(string(field.Type)) {
		name = "[" + name + "]"
	}

	line := fmt.Sprintf("@property {%v} %v", typeStr, name)
	if desc := oneLine(field.Description); desc != "" {
		line += " - " + desc
	}

	return jsdocRenderedField{Line: line}
}

func jsdocRenderFields(fields []*core.EmiField, parentChain string) []jsdocRenderedField {
	out := []jsdocRenderedField{}
	for _, field := range fields {
		if field == nil {
			continue
		}
		out = append(out, jsdocRenderField(field, parentChain))
	}
	return out
}

type jsdocRenderedType struct {
	// TypeName is already the full, dotted `@typedef` name (e.g.
	// "WidgetDtoType" at the root, "WidgetDtoType.NestedType" nested) - matches
	// treeLocation exactly, so every reference tsFieldType computed against
	// that same treeLocation resolves to a typedef this file actually declares.
	TypeName string
	Fields   []jsdocRenderedField
	Doc      string
}

// jsdocRenderTypes walks the field tree the same way tsRenderTypes does, but
// returns a flat list (children first) instead of a tree of `.SubTypes` -
// JSDoc `@typedef`s have no real block-scoping construct to nest them inside
// the way a TS `namespace` does, so nesting here is conveyed purely through the
// dotted name.
func jsdocRenderTypes(fields []*core.EmiField, typeName string, treeLocation string) []jsdocRenderedType {
	current := jsdocRenderedType{
		TypeName: treeLocation,
		Fields:   jsdocRenderFields(fields, treeLocation),
		Doc:      fmt.Sprintf("The base type definition for %v", core.ToLower(typeName)),
	}

	out := []jsdocRenderedType{}
	for _, field := range fields {
		if field == nil {
			continue
		}
		if field.Type == core.FieldTypeObject || field.Type == core.FieldTypeArray ||
			field.Type == core.FieldTypeObjectNullable || field.Type == core.FieldTypeArrayNullable {
			childName := core.ToUpper(field.Name) + "Type"
			out = append(out, jsdocRenderTypes(field.Fields, childName, treeLocation+"."+childName)...)
		}
	}

	return append(out, current)
}

type JsDocCommonObjectContext struct {
	RootTypeName string
}

var jsdocTypesTmpl = template.Must(template.New("jsdoctypes").Parse(`
{{ range . }}
/**
 * {{ .Doc }}
 * @typedef {Object} {{ .TypeName }}
{{- range .Fields }}
 * {{ .Line }}
{{- end }}
 */
{{ end }}
`))

// JsDocCommonObjectGenerator generates one `@typedef` per object/array shape in
// the field tree (the root, plus any nested object/array) - the JSDoc
// equivalent of TsCommonObjectGenerator.
func JsDocCommonObjectGenerator(fields []*core.EmiField, jsdctx JsDocCommonObjectContext) (*core.CodeChunkCompiled, error) {
	res := &core.CodeChunkCompiled{}

	rootTypeName := core.ToUpper(jsdctx.RootTypeName) + "Type"
	renderedTypes := jsdocRenderTypes(fields, jsdctx.RootTypeName, rootTypeName)

	// TOKEN_OBJ_CLASS/TOKEN_OBJ_TYPE (the same string - see js-tokens.go) must
	// keep resolving to the canonical/class-equivalent name here, exactly like
	// TsCommonObjectGenerator's own (bare, un-suffixed) token does - callers
	// like fetchctx.ResponseClass (js-action-main-class.go) use it to build a
	// `new X(...)` call, which the `Type`-suffixed typedef name was never a
	// valid identifier for. The real typedef name is only ever exposed via the
	// dedicated TOKEN_TYPEDEF_NAME below.
	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_OBJ_TYPE, Value: core.ToUpper(jsdctx.RootTypeName)})
	res.Tokens = append(res.Tokens, core.GeneratedScriptToken{Name: TOKEN_TYPEDEF_NAME, Value: rootTypeName})

	var buf bytes.Buffer
	if err := jsdocTypesTmpl.Execute(&buf, renderedTypes); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	return res, nil
}
