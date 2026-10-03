package swift

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

const interfaceYAML = `
interfaces:
  - name: sellable
    description: Something that can be sold.
    fields:
      - {name: sku, type: string}
      - {name: price, type: 'int64?'}
  - name: audited
    mutable: true
    fields: [{name: createdBy, type: string}]
dtos:
  - {name: Score, implements: [sellable, audited], fields: [{name: bpm, type: int}]}
  - {name: Book, fields: [{name: title, type: string}]}
`

func interfaceFiles(t *testing.T) map[string]string {
	t.Helper()
	m, err := core.StringToEmi(interfaceYAML)
	if err != nil {
		t.Fatal(err)
	}
	files, err := SwiftFullModule(&m, core.MicroGenContext{})
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, f := range files {
		out[f.Name+f.Extension] = f.ActualScript
	}
	return out
}

func TestInterfaceGeneratedAndImplemented(t *testing.T) {
	files := interfaceFiles(t)
	for name, want := range map[string][]string{
		"Sellable.swift": []string{"protocol Sellable", "var sku: String { get }", "var price:"},
		"Audited.swift":  []string{"protocol Audited", "var createdBy: String { get set }"},
		"ScoreDto.swift": []string{"struct ScoreDto: Codable, Sellable, Audited", "var createdBy: String"},
	} {
		got, ok := files[name]
		if !ok {
			t.Fatalf("%s not generated; have %v", name, len(files))
		}
		for _, w := range want {
			if !strings.Contains(got, w) {
				t.Errorf("%s missing %q:\n%s", name, w, got)
			}
		}
	}
	if book := files["BookDto.swift"]; strings.Contains(book, "Sellable") {
		t.Errorf("BookDto must not implement anything:\n%s", book)
	}
}
