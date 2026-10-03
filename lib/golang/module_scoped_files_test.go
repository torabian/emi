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
intents:
  - name: findThing
    description: Finds a thing.
    in:
      fields:
        - {name: query, type: string}
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
	scoped := moduleFileNames(t, "clicktobuy", "")
	for _, name := range []string{"ClicktobuyPermissions.go", "ClicktobuyEvents.go", "ClicktobuyIntents.go"} {
		if !scoped[name] {
			t.Fatalf("missing %s: %v", name, scoped)
		}
	}
	for _, name := range []string{"Permissions.go", "Events.go", "Intents.go"} {
		if scoped[name] {
			t.Fatalf("unprefixed %s still generated: %v", name, scoped)
		}
	}

	// Two modules generated into one folder must not produce the same file names.
	other := moduleFileNames(t, "score", "")
	for name := range scoped {
		if strings.HasSuffix(name, "Permissions.go") && other[name] {
			t.Fatalf("collision on %s", name)
		}
	}

	// A module without a name keeps the plain names instead of producing "Permissions" prefixed by nothing odd.
	if n := moduleFileNames(t, "\"\"", ""); !n["Permissions.go"] {
		t.Fatalf("nameless module: %v", n)
	}
}
