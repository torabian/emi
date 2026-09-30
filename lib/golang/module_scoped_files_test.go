package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func moduleFileNames(t *testing.T, name, tags string) map[string]bool {
	t.Helper()
	m, err := core.StringToEmi(`
name: ` + name + `
permissions:
  - {name: thing, key: thing, title: {en: Thing}, description: {en: Thing}}
events:
  - {key: thingHappened, name: {en: Happened}, description: {en: Happened}}
`)
	if err != nil {
		t.Fatal(err)
	}
	files, err := GoModuleFull(&m, core.MicroGenContext{Tags: tags, Flags: map[string]string{"pkg": "x"}})
	if err != nil {
		t.Fatal(err)
	}
	names := map[string]bool{}
	for _, f := range files {
		names[f.Name+f.Extension] = true
	}
	return names
}

func TestModuleScopedFiles(t *testing.T) {
	plain := moduleFileNames(t, "clicktobuy", "")
	if !plain["Permissions.go"] || !plain["Events.go"] {
		t.Fatalf("default names changed: %v", plain)
	}

	scoped := moduleFileNames(t, "clicktobuy", string(core.ModuleScopedFiles))
	if !scoped["ClicktobuyPermissions.go"] || !scoped["ClicktobuyEvents.go"] || scoped["Permissions.go"] || scoped["Events.go"] {
		t.Fatalf("scoped names: %v", scoped)
	}

	// Two modules generated into one folder must not produce the same file names.
	other := moduleFileNames(t, "score", string(core.ModuleScopedFiles))
	for name := range scoped {
		if strings.HasSuffix(name, "Permissions.go") && other[name] {
			t.Fatalf("collision on %s", name)
		}
	}

	// A module without a name keeps the plain names instead of producing "Permissions" prefixed by nothing odd.
	if n := moduleFileNames(t, "\"\"", string(core.ModuleScopedFiles)); !n["Permissions.go"] {
		t.Fatalf("nameless module: %v", n)
	}
}
