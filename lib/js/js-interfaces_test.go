package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func jsFilesOf(t *testing.T, tags, yaml string) map[string]string {
	t.Helper()
	m, err := core.StringToEmi(yaml)
	if err != nil {
		t.Fatal(err)
	}
	files, err := JsModuleFullVirtualFiles(&m, core.MicroGenContext{Tags: tags})
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, f := range files {
		out[f.Name+f.Extension] = f.ActualScript
	}
	return out
}

const jsInterfaceYAML = `
interfaces:
  - name: sellable
    description: Something that can be sold.
    fields:
      - {name: sku, type: string, description: Stock keeping unit}
      - {name: price, type: 'int64?'}
  - name: audited
    mutable: true
    fields: [{name: createdBy, type: string}]
dtos:
  - {name: Score, implements: [sellable, audited], fields: [{name: bpm, type: int}]}
  - {name: Book, implements: [sellable]}
`

func TestJsInterfaceGeneratedAndImplemented(t *testing.T) {
	files := jsFilesOf(t, "typescript,no-package,no-sdk", jsInterfaceYAML)

	sellable := files["Sellable.ts"]
	// the property types are the class's own (the setter's parameter type), or tsc rejects `implements`
	for _, want := range []string{"export interface Sellable", "readonly sku: string;", "readonly price: number | null | undefined;", "/** Stock keeping unit */"} {
		if !strings.Contains(sellable, want) {
			t.Errorf("Sellable.ts missing %q:\n%s", want, sellable)
		}
	}
	if audited := files["Audited.ts"]; !strings.Contains(audited, "  createdBy: string;") || strings.Contains(audited, "readonly") {
		t.Errorf("mutable interface should not be readonly:\n%s", audited)
	}

	score := files["ScoreDto.ts"]
	if !strings.Contains(score, "export class ScoreDto implements Sellable, Audited") {
		t.Errorf("ScoreDto.ts missing implements clause")
	}
	for _, want := range []string{`import { type Sellable } from "./Sellable"`, `import { type Audited } from "./Audited"`} {
		if !strings.Contains(score, want) {
			t.Errorf("ScoreDto.ts missing import %q", want)
		}
	}
	if !strings.Contains(files["BookDto.ts"], "export class BookDto implements Sellable {") {
		t.Errorf("BookDto.ts missing implements clause")
	}
}

func TestJsInterfaceSkippedForPlainJavaScript(t *testing.T) {
	files := jsFilesOf(t, "no-package,no-sdk", jsInterfaceYAML)
	if _, generated := files["Sellable.ts"]; generated {
		t.Fatal("plain JavaScript has no interfaces")
	}
	for name := range files {
		if strings.HasPrefix(name, "Sellable") || strings.HasPrefix(name, "Audited") {
			t.Fatalf("unexpected interface file %s", name)
		}
	}
	// the dto still gets the interface's fields
	if score := files["ScoreDto.js"]; !strings.Contains(score, "sku") || strings.Contains(score, "implements") {
		t.Fatalf("ScoreDto.js:\n%s", score)
	}
}

func TestJsInterfaceReferencedFromAnotherModule(t *testing.T) {
	files := jsFilesOf(t, "typescript,no-package,no-sdk", `
interfaces:
  - name: purchasable
    jsProvider: '@/modules/clicktobuy/sdk/Purchasable'
    fields: [{name: sku, type: string}]
dtos:
  - {name: Score, implements: [purchasable]}
`)
	if _, generated := files["Purchasable.ts"]; generated {
		t.Fatal("a referenced interface must not be generated again")
	}
	score := files["ScoreDto.ts"]
	for _, want := range []string{`import { type Purchasable } from "@/modules/clicktobuy/sdk/Purchasable"`, "implements Purchasable"} {
		if !strings.Contains(score, want) {
			t.Errorf("ScoreDto.ts missing %q:\n%s", want, score[:min(len(score), 600)])
		}
	}
}

func TestJsInterfaceNoInterfacesTag(t *testing.T) {
	files := jsFilesOf(t, "typescript,no-package,no-sdk,no-interfaces", jsInterfaceYAML)
	if _, generated := files["Sellable.ts"]; generated {
		t.Fatal("no-interfaces still generated an interface")
	}
	score := files["ScoreDto.ts"]
	if strings.Contains(score, "implements") || strings.Contains(score, "./Sellable") {
		t.Fatalf("no-interfaces still added implements/import")
	}
	if !strings.Contains(score, "get sku") {
		t.Fatal("the interface's fields must still be included")
	}
}

func TestJsInterfaceObjectAndArrayFields(t *testing.T) {
	files := jsFilesOf(t, "typescript,no-package,no-sdk", `
interfaces:
  - name: titlable
    fields:
      - {name: title, type: string}
      - name: content
        type: object
        fields:
          - {name: text, type: string}
          - name: parts
            type: array
            fields: [{name: kind, type: string}]
      - name: notes
        type: 'array?'
        fields: [{name: note, type: string}]
dtos:
  - {name: Cloth, implements: [titlable]}
`)
	src := files["Titlable.ts"]
	// nested interfaces are owned by the interface; property types mirror the class's
	// setter parameter (so `implements` holds) with the nested interface in place of the class
	for _, want := range []string{
		"export interface Titlable {",
		"readonly content: TitlableContent;",
		"readonly notes: MArray<TitlableNotes> | TitlableNotes[] | null | undefined;",
		"export interface TitlableContent {",
		"readonly parts: MArray<TitlableContentParts> | TitlableContentParts[];",
		"export interface TitlableContentParts {",
		"export interface TitlableNotes {",
		`import { type MArray } from`,
	} {
		if !strings.Contains(src, want) {
			t.Errorf("Titlable.ts missing %q:\n%s", want, src)
		}
	}
	if !strings.Contains(files["ClothDto.ts"], "export class ClothDto implements Titlable {") {
		t.Errorf("ClothDto.ts missing implements clause")
	}
}

func TestJsEntityDerivedDtoImplementsInterface(t *testing.T) {
	// The entity dtos are derived by a preprocess hook the compiler binary registers (it
	// links the Go backend); register the same ones here, as the binary effectively does.
	core.RegisterPreprocessHook(core.PreprocessEntityDtos)
	core.RegisterPreprocessHook(core.PreprocessEntityOptionalDtos)

	files := jsFilesOf(t, "typescript,no-package,no-sdk", `
interfaces:
  - name: auditable
    fields: [{name: createdBy, type: string}]
entities:
  - {name: invoice, implements: [auditable], fields: [{name: number, type: string}]}
`)
	if !strings.Contains(files["InvoiceDto.ts"], "export class InvoiceDto implements Auditable {") {
		t.Errorf("the dto derived from the entity should implement the interface")
	}
	if strings.Contains(files["InvoiceOptionalDto.ts"], "implements") {
		t.Errorf("optional dto must not implement interfaces")
	}
}
