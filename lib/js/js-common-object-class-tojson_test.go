package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestToJSONResolvesNestedValuesToPlainShape covers the fix to toJSON():
// before it, a nested array/object/complex field's toJSON() entry was a bare
// `this.#field` - the raw emigo Array/One/Collection wrapper or nested dto
// instance, not the plain shape the exported *Type declares. Assigning
// `const x: FooType = new Foo(...).toJSON()` failed to compile for any dto
// with such a field (verified against a real tsc run - see
// examples/vsql-columns/ts/tojson-typesafety.ts). Every field now goes
// through a private #toPlainJSON helper that recurses into anything
// carrying its own toJSON instead of returning it as-is.
func TestToJSONResolvesNestedValuesToPlainShape(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "email", Type: core.FieldTypeString},
		{
			Name: "users",
			Type: core.FieldTypeArray,
			Fields: []*core.EmiField{
				{Name: "email", Type: core.FieldTypeString},
			},
		},
	}

	t.Run("plain JS", func(t *testing.T) {
		chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{}, JsCommonObjectContext{RootClassName: "Foo"})
		if err != nil {
			t.Fatalf("JsCommonObjectGenerator: %v", err)
		}
		src := string(chunk.ActualScript)

		if !strings.Contains(src, "email: this.#toPlainJSON(this.#email)") {
			t.Fatalf("expected a primitive field to still go through #toPlainJSON, got:\n%s", src)
		}
		if !strings.Contains(src, "users: this.#toPlainJSON(this.#users)") {
			t.Fatalf("expected the array field to go through #toPlainJSON instead of a bare this.#users, got:\n%s", src)
		}
		if !strings.Contains(src, "#toPlainJSON(value) {") {
			t.Fatalf("expected an untyped #toPlainJSON helper in plain JS output, got:\n%s", src)
		}
		if strings.Contains(src, ": unknown") || strings.Contains(src, "as any") {
			t.Fatalf("did not expect any TS-only type annotation in plain JS output, got:\n%s", src)
		}
	})

	t.Run("typescript", func(t *testing.T) {
		ctx := core.MicroGenContext{Tags: "typescript"}
		chunk, err := JsCommonObjectGenerator(fields, ctx, JsCommonObjectContext{RootClassName: "Foo"})
		if err != nil {
			t.Fatalf("JsCommonObjectGenerator: %v", err)
		}
		src := string(chunk.ActualScript)

		if !strings.Contains(src, "toJSON(): FooType {") {
			t.Fatalf("expected toJSON() to be annotated with the class's own exported Type, got:\n%s", src)
		}
		if !strings.Contains(src, "} as FooType;") {
			t.Fatalf("expected the returned object literal to be asserted to FooType (every field is #toPlainJSON's `unknown`), got:\n%s", src)
		}
		if !strings.Contains(src, "#toPlainJSON(value: unknown): unknown {") {
			t.Fatalf("expected a typed #toPlainJSON helper in TypeScript output, got:\n%s", src)
		}

		// The nested Users class gets its own correctly-scoped toJSON too -
		// namespaced under FooType, not a second bare "FooType".
		if !strings.Contains(src, "toJSON(): FooType.UsersType {") {
			t.Fatalf("expected the nested Users class's toJSON to be annotated with its own namespaced type, got:\n%s", src)
		}
	})
}
