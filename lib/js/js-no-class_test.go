package js

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func widgetGetAction() *core.EmiAction {
	return &core.EmiAction{
		Name:   "getWidget",
		Url:    "/widget/:id",
		Method: "get",
		Out: &core.EmiActionBody{
			Fields: []*core.EmiField{
				{Name: "title", Type: core.FieldTypeString},
			},
		},
	}
}

// TestNoClassTag_DtoBecomesTypeOnly covers the core ask: --tags no-class skips
// the generated class body (private fields/getters/setters/toJSON/clone/...)
// entirely, in both dialects, leaving only the type declaration that already
// exists today (a real `export type` in TypeScript, a JSDoc `@typedef` in
// plain JS - see js-common-object-types.go/js-common-object-jsdoc.go).
func TestNoClassTag_DtoBecomesTypeOnly(t *testing.T) {
	fields := []*core.EmiField{{Name: "title", Type: core.FieldTypeString}}

	t.Run("plain JS", func(t *testing.T) {
		chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{Tags: "no-class"}, JsCommonObjectContext{RootClassName: "WidgetDto"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if strings.Contains(script, "export class WidgetDto") {
			t.Errorf("expected no class body with --tags no-class, got:\n%v", script)
		}
		if !strings.Contains(script, "@typedef {Object} WidgetDtoType") {
			t.Errorf("expected the JSDoc typedef to still be present, got:\n%v", script)
		}
	})

	t.Run("typescript", func(t *testing.T) {
		chunk, err := JsCommonObjectGenerator(fields, core.MicroGenContext{Tags: "typescript,no-class"}, JsCommonObjectContext{RootClassName: "WidgetDto"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(chunk.ActualScript)
		if strings.Contains(script, "export class WidgetDto") {
			t.Errorf("expected no class body with --tags no-class, got:\n%v", script)
		}
		if !strings.Contains(script, "export type WidgetDtoType") {
			t.Errorf("expected the real TS type to still be present, got:\n%v", script)
		}
	})
}

// TestNoClassTag_DefaultPathUnaffected is the other, equally important half of
// this feature: not passing --tags no-class must produce byte-for-byte the
// same class-generating behavior as before this tag existed - every place this
// tag touches (JsCommonObjectGenerator, getCreatorFnInfo,
// getCommonFetchArguments, JsActionFetchAndMetaData) does so through a
// dedicated, always-false-by-default branch, never by altering the existing
// code path.
func TestNoClassTag_DefaultPathUnaffected(t *testing.T) {
	action := widgetGetAction()

	for _, tags := range []string{"", "typescript"} {
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: tags}, nil)
		if err != nil {
			t.Fatalf("tags=%q: unexpected error: %v", tags, err)
		}

		if realms.Fetchctx.ResponseClass != "GetWidgetActionRes" {
			t.Errorf("tags=%q: expected ResponseClass to resolve to the class name \"GetWidgetActionRes\", got %q", tags, realms.Fetchctx.ResponseClass)
		}

		script := string(realms.FetchMetaClass.ActualScript)
		if !strings.Contains(script, "new GetWidgetActionRes(item)") {
			t.Errorf("tags=%q: expected the response to still be instantiated via `new GetWidgetActionRes(item)`, got:\n%v", tags, script)
		}
		if !strings.Contains(string(realms.ResponseClass.ActualScript), "export class GetWidgetActionRes") {
			t.Errorf("tags=%q: expected the response dto to still be a real class, got:\n%v", tags, string(realms.ResponseClass.ActualScript))
		}
	}
}

// TestNoClassTag_ActionFetchDoesNotInstantiate covers the second half of the
// ask: under --tags no-class, the generated Fetch/Fetch$ methods must not
// construct anything - just hand back the parsed response, typed (a real
// TS generic, or a JSDoc @returns in plain JS) but never `new`-ed.
func TestNoClassTag_ActionFetchDoesNotInstantiate(t *testing.T) {
	action := widgetGetAction()

	t.Run("plain JS", func(t *testing.T) {
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: "no-class"}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if realms.Fetchctx.ResponseClass != "GetWidgetActionResType" {
			t.Errorf("expected ResponseClass to resolve to the typedef name, got %q", realms.Fetchctx.ResponseClass)
		}

		script := string(realms.FetchMetaClass.ActualScript)
		if strings.Contains(script, "creatorFn") {
			t.Errorf("did not expect any creatorFn/instantiation machinery, got:\n%v", script)
		}
		if !strings.Contains(script, "(item) => item") {
			t.Errorf("expected the response to be handed back untouched via an identity function, got:\n%v", script)
		}
		if !strings.Contains(script, "@returns") || !strings.Contains(script, "GetWidgetActionResType") {
			t.Errorf("expected a JSDoc @returns annotation pointing at the typedef, got:\n%v", script)
		}
	})

	t.Run("typescript", func(t *testing.T) {
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: "typescript,no-class"}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if realms.Fetchctx.ResponseClass != "GetWidgetActionResType" {
			t.Errorf("expected ResponseClass to resolve to the typedef name, got %q", realms.Fetchctx.ResponseClass)
		}

		script := string(realms.FetchMetaClass.ActualScript)
		if strings.Contains(script, "creatorFn") {
			t.Errorf("did not expect any creatorFn/instantiation machinery, got:\n%v", script)
		}
		if !strings.Contains(script, "fetchx<GetWidgetActionResType, unknown, unknown>") {
			t.Errorf("expected the response type to flow into the fetchx<> generic, got:\n%v", script)
		}
	})
}

// TestNoClassTag_GeneratedJsIsValidSyntax runs the no-class output (request +
// response + a nested object/array field, to exercise more of the field-plan)
// through `node --check` - a much stronger guarantee than any Go-side string
// assertion could give.
func TestNoClassTag_GeneratedJsIsValidSyntax(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not available")
	}

	action := &core.EmiAction{
		Name:   "getWidget",
		Url:    "/widget/:id",
		Method: "get",
		In: &core.EmiActionBody{
			Fields: []*core.EmiField{{Name: "note", Type: core.FieldTypeStringNullable}},
		},
		Out: &core.EmiActionBody{
			Fields: []*core.EmiField{
				{Name: "title", Type: core.FieldTypeString},
				{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
				{
					Name: "address",
					Type: core.FieldTypeObject,
					Fields: []*core.EmiField{
						{Name: "city", Type: core.FieldTypeString},
					},
				},
			},
		},
	}

	realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: "no-class"}, nil)
	if err != nil {
		t.Fatal(err)
	}

	script := string(realms.RequestClass.ActualScript) + "\n" +
		string(realms.ResponseClass.ActualScript) + "\n" +
		string(realms.FetchMetaClass.ActualScript) + "\n"

	f, err := os.CreateTemp("", "noclass-*.mjs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if _, err := f.WriteString(script); err != nil {
		t.Fatal(err)
	}
	f.Close()

	out, err := exec.Command("node", "--check", f.Name()).CombinedOutput()
	if err != nil {
		t.Fatalf("generated no-class JS failed node --check:\n%v\n\n--- script ---\n%v", string(out), script)
	}
}
