package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestJsDocCommonObjectGenerator_DoesNotHijackTheClassNameToken is a regression
// test: TOKEN_OBJ_CLASS and TOKEN_OBJ_TYPE are the *same* string constant by
// design (a TS class name doubles as its own type), and findTokenByName returns
// the first match - so if this generator's own token (appended before the real
// class token, see JsCommonObjectGenerator) ever carries the `Type`-suffixed
// typedef name instead of the bare/canonical one, every downstream consumer of
// TOKEN_OBJ_CLASS (most notably fetchctx.ResponseClass in
// js-action-main-class.go, which builds a `new X(...)` instantiation) would
// silently resolve to that nonexistent, typedef-only identifier instead of the
// real class - a runtime ReferenceError in generated plain-JS action code, not
// a compile-time error anywhere in this Go package.
func TestJsDocCommonObjectGenerator_DoesNotHijackTheClassNameToken(t *testing.T) {
	chunk, err := JsDocCommonObjectGenerator(
		[]*core.EmiField{{Name: "title", Type: core.FieldTypeString}},
		JsDocCommonObjectContext{RootTypeName: "WidgetDto"},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	classToken := core.FindTokenByName(chunk.Tokens, TOKEN_OBJ_CLASS)
	if classToken == nil || classToken.Value != "WidgetDto" {
		t.Errorf("expected TOKEN_OBJ_CLASS to resolve to the bare class name \"WidgetDto\", got %+v", classToken)
	}

	typedefToken := core.FindTokenByName(chunk.Tokens, TOKEN_TYPEDEF_NAME)
	if typedefToken == nil || typedefToken.Value != "WidgetDtoType" {
		t.Errorf("expected TOKEN_TYPEDEF_NAME to resolve to \"WidgetDtoType\", got %+v", typedefToken)
	}
}

// TestJsDocCommonObjectGenerator_PlainJsGetsTypedefs covers the whole point of
// js-common-object-jsdoc.go: a plain (non-TypeScript) generated dto gets a real
// `@typedef` describing its shape - required/nullable fields, primitives,
// slices, and nested object/array fields referencing their own nested,
// dot-namespaced typedef - so an editor's JS language service has something to
// offer intellisense from without needing --tags typescript.
func TestJsDocCommonObjectGenerator_PlainJsGetsTypedefs(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "title", Type: core.FieldTypeString, Description: "The widget's title"},
		{Name: "nickname", Type: core.FieldTypeStringNullable},
		{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
		{
			Name: "address",
			Type: core.FieldTypeObject,
			Fields: []*core.EmiField{
				{Name: "city", Type: core.FieldTypeString},
			},
		},
	}

	chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{}, JsCommonObjectContext{RootClassName: "WidgetDto"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	for _, want := range []string{
		"@typedef {Object} WidgetDtoType",
		"@property {string} title - The widget's title",
		// a nullable field is bracketed, the standard JSDoc "optional" marker.
		"@property {string} [nickname]",
		"@property {string[]} tags",
		// a nested object field references its own, dot-namespaced typedef.
		"@property {WidgetDtoType.AddressType} address",
		"@typedef {Object} WidgetDtoType.AddressType",
		"@property {string} city",
	} {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated script to contain %q, got:\n%v", want, script)
		}
	}

	if !strings.Contains(script, "export class WidgetDto") {
		t.Error("expected the runtime class to still be generated alongside the typedefs")
	}
}

// TestJsDocCommonObjectGenerator_SkippedInTypeScript covers the other half of
// the same feature: TypeScript mode already has a real `export type` (see
// TsCommonObjectGenerator) - a JSDoc typedef describing the exact same shape
// would be redundant, so it must never be emitted there.
func TestJsDocCommonObjectGenerator_SkippedInTypeScript(t *testing.T) {
	fields := []*core.EmiField{{Name: "title", Type: core.FieldTypeString}}

	chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{Tags: "typescript"}, JsCommonObjectContext{RootClassName: "WidgetDto"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if strings.Contains(script, "@typedef") {
		t.Errorf("did not expect a JSDoc @typedef in TypeScript output, got:\n%v", script)
	}
	if !strings.Contains(script, "export type WidgetDtoType") {
		t.Errorf("expected the real TS type instead, got:\n%v", script)
	}
}

// TestJsDocCommonObjectGenerator_NoJsDocTagOptsOut covers the --tags no-jsdoc
// escape hatch for callers who don't want the extra section even in plain JS.
func TestJsDocCommonObjectGenerator_NoJsDocTagOptsOut(t *testing.T) {
	fields := []*core.EmiField{{Name: "title", Type: core.FieldTypeString}}

	chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{Tags: "no-jsdoc"}, JsCommonObjectContext{RootClassName: "WidgetDto"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	script := string(chunk.ActualScript)

	if strings.Contains(script, "@typedef") {
		t.Errorf("expected --tags no-jsdoc to skip the typedef section, got:\n%v", script)
	}
	if !strings.Contains(script, "export class WidgetDto") {
		t.Error("expected the runtime class to still be generated")
	}
}
