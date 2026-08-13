package js

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestNoDefinitionTag_SkipsStaticDefinitionBlock covers --tags no-definition:
// the `static Definition = {...}` JSON dump (the whole action re-serialized as
// a literal, for introspection/tooling) is dropped, while the default path -
// no tag at all - keeps generating it exactly as before.
func TestNoDefinitionTag_SkipsStaticDefinitionBlock(t *testing.T) {
	action := widgetGetAction()

	t.Run("tag set", func(t *testing.T) {
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{Tags: "no-definition"}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(realms.FetchMetaClass.ActualScript)
		if strings.Contains(script, "static Definition") {
			t.Errorf("expected no static Definition block with --tags no-definition, got:\n%v", script)
		}
	})

	t.Run("default path unaffected", func(t *testing.T) {
		realms, _, err := JsActionManifestRealms(action, core.MicroGenContext{}, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		script := string(realms.FetchMetaClass.ActualScript)
		if !strings.Contains(script, "static Definition = {") {
			t.Errorf("expected static Definition to still be generated without the tag, got:\n%v", script)
		}
		if !strings.Contains(script, `"name": "getWidget"`) {
			t.Errorf("expected the definition JSON to still contain the action's own data, got:\n%v", script)
		}
	})
}

// TestNoDefinitionTag_GeneratedJsIsValidSyntax runs the no-definition (combined
// with no-class, to also exercise a slimmed-down dto) output through
// `node --check`.
func TestNoDefinitionTag_GeneratedJsIsValidSyntax(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not available")
	}

	realms, _, err := JsActionManifestRealms(widgetGetAction(), core.MicroGenContext{Tags: "no-definition,no-class"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	script := string(realms.ResponseClass.ActualScript) + "\n" + string(realms.FetchMetaClass.ActualScript)

	f, err := os.CreateTemp("", "nodefinition-*.mjs")
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
		t.Fatalf("generated no-definition JS failed node --check:\n%v\n\n--- script ---\n%v", string(out), script)
	}
}
