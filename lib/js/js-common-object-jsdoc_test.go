package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

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
