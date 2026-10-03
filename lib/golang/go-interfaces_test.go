package golang

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func goFilesOf(t *testing.T, yaml string) map[string]string {
	t.Helper()
	m, err := core.StringToEmi(yaml)
	if err != nil {
		t.Fatal(err)
	}
	files, err := GoModuleFull(&m, core.MicroGenContext{Flags: map[string]string{"pkg": "defs"}})
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]string{}
	for _, f := range files {
		out[f.Name+f.Extension] = f.ActualScript
	}
	return out
}

func TestGoInterfaceGeneratedAndImplemented(t *testing.T) {
	files := goFilesOf(t, `
name: shop
interfaces:
  - name: sellable
    description: Something that can be sold.
    fields:
      - {name: sku, type: string, description: Stock keeping unit}
      - {name: price, type: 'int64?'}
dtos:
  - {name: Score, implements: [sellable], fields: [{name: bpm, type: int}]}
  - {name: Book, implements: [sellable]}
`)
	iface := files["Sellable.go"]
	for _, want := range []string{"type Sellable interface", "GetSku() string", "GetPrice() emigo.Nullable[int64]", "// Stock keeping unit"} {
		if !strings.Contains(iface, want) {
			t.Errorf("Sellable.go missing %q:\n%s", want, iface)
		}
	}
	if strings.Contains(iface, "SetSku") {
		t.Errorf("getters-only interface has setters:\n%s", iface)
	}
	if strings.Contains(iface, "encoding/json") {
		t.Errorf("interface imports encoding/json it never uses:\n%s", iface)
	}
	for _, dto := range []string{"ScoreDto.go", "BookDto.go"} {
		src := files[dto]
		name := strings.TrimSuffix(dto, ".go")
		for _, want := range []string{"func (x *" + name + ") GetSku() string", "var _ Sellable = (*" + name + ")(nil)"} {
			if !strings.Contains(src, want) {
				t.Errorf("%s missing %q", dto, want)
			}
		}
	}
}

func TestGoInterfaceMutableAddsSetters(t *testing.T) {
	files := goFilesOf(t, `
interfaces:
  - {name: sellable, mutable: true, fields: [{name: sku, type: string}]}
dtos:
  - {name: Score, implements: [sellable]}
`)
	if !strings.Contains(files["Sellable.go"], "SetSku(v string)") || !strings.Contains(files["ScoreDto.go"], "func (x *ScoreDto) SetSku(v string)") {
		t.Fatalf("setters missing:\n%s\n%s", files["Sellable.go"], files["ScoreDto.go"])
	}
}

// A module implementing an interface declared elsewhere imports it: no second type.
func TestGoInterfaceReferencedFromAnotherPackage(t *testing.T) {
	files := goFilesOf(t, `
interfaces:
  - name: purchasable
    module: ctbdefs
    provider: example.com/clicktobuy/defs
    fields: [{name: sku, type: string}]
dtos:
  - {name: Score, implements: [purchasable]}
`)
	if _, generated := files["Purchasable.go"]; generated {
		t.Fatal("a referenced interface must not be generated again")
	}
	src := files["ScoreDto.go"]
	for _, want := range []string{`ctbdefs "example.com/clicktobuy/defs"`, "var _ ctbdefs.Purchasable = (*ScoreDto)(nil)", "GetSku() string"} {
		if !strings.Contains(src, want) {
			t.Errorf("ScoreDto.go missing %q:\n%s", want, src)
		}
	}
}

func TestGoInterfaceReferenceNeedsModuleAlias(t *testing.T) {
	m, err := core.StringToEmi(`
interfaces:
  - {name: purchasable, provider: example.com/x, fields: [{name: sku, type: string}]}
dtos:
  - {name: Score, implements: [purchasable]}
`)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := GoModuleFull(&m, core.MicroGenContext{Flags: map[string]string{"pkg": "defs"}}); err == nil || !strings.Contains(err.Error(), "module") {
		t.Fatalf("want missing-module error, got %v", err)
	}
}

const goNestedInterfaceYAML = `
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
  - {name: Shoe, implements: [titlable]}
`

// The interface declares the nested structs once; every implementer's own nested types are
// aliases of them, so the accessor returns one type for all implementers.
func TestGoInterfaceObjectAndArrayFields(t *testing.T) {
	files := goFilesOf(t, goNestedInterfaceYAML)

	iface := files["Titlable.go"]
	for _, want := range []string{
		"GetContent() TitlableContent",
		"GetNotes() emigo.ArrayNullable[TitlableNotes]",
		"type TitlableContent struct",
		"Parts emigo.Array[TitlableContentParts]",
		"type TitlableContentParts struct",
		"type TitlableNotes struct",
	} {
		if !strings.Contains(iface, want) {
			t.Errorf("Titlable.go missing %q:\n%s", want, iface)
		}
	}

	for _, dto := range []string{"Cloth", "Shoe"} {
		src := files[dto+"Dto.go"]
		for _, want := range []string{
			"type " + dto + "DtoContent = TitlableContent",
			"type " + dto + "DtoContentParts = TitlableContentParts",
			"type " + dto + "DtoNotes = TitlableNotes",
			"GetContent() " + dto + "DtoContent",
		} {
			if !strings.Contains(src, want) {
				t.Errorf("%sDto.go missing %q", dto, want)
			}
		}
		// aliased: no second struct definition of the same shape
		if strings.Contains(src, "type "+dto+"DtoContent struct") {
			t.Errorf("%sDto.go still declares its own Content struct", dto)
		}
	}
}

// A dto's own nested field that is not from an interface keeps its struct.
func TestGoInterfaceLeavesOtherNestedFieldsAlone(t *testing.T) {
	files := goFilesOf(t, `
interfaces:
  - name: titlable
    fields:
      - name: content
        type: object
        fields: [{name: text, type: string}]
dtos:
  - name: Cloth
    implements: [titlable]
    fields:
      - name: extra
        type: object
        fields: [{name: a, type: string}]
`)
	src := files["ClothDto.go"]
	if !strings.Contains(src, "type ClothDtoExtra struct") || !strings.Contains(src, "type ClothDtoContent = TitlableContent") {
		t.Fatalf("ClothDto.go:\n%s", src)
	}
}

func TestGoInterfaceNestedTypesReferencedFromAnotherPackage(t *testing.T) {
	files := goFilesOf(t, `
interfaces:
  - name: titlable
    module: shared
    provider: example.com/shared/defs
    fields:
      - name: content
        type: object
        fields: [{name: text, type: string}]
dtos:
  - {name: Cloth, implements: [titlable]}
`)
	if _, generated := files["Titlable.go"]; generated {
		t.Fatal("a referenced interface must not be generated again")
	}
	src := files["ClothDto.go"]
	for _, want := range []string{"type ClothDtoContent = shared.TitlableContent", `shared "example.com/shared/defs"`, "var _ shared.Titlable = (*ClothDto)(nil)"} {
		if !strings.Contains(src, want) {
			t.Errorf("ClothDto.go missing %q:\n%s", want, src)
		}
	}
}

const goEntityInterfaceYAML = `
interfaces:
  - name: auditable
    fields:
      - {name: createdBy, type: string}
      - {name: note, type: 'string?'}
  - name: tagged
    context: entity
    fields: [{name: tag, type: 'one?', target: TagEntity}]
entities:
  - name: tag
    fields: [{name: label, type: string}]
  - name: invoice
    implements: [auditable, tagged]
    fields: [{name: number, type: string}]
`

// An entity satisfies its interfaces from its own final struct types, and the dto derived
// from it satisfies the ones whose types it shares.
func TestGoEntityImplementsInterfaces(t *testing.T) {
	files := goFilesOf(t, goEntityInterfaceYAML)

	entity := files["InvoiceEntity.go"]
	for _, want := range []string{
		"func (x *InvoiceEntity) GetCreatedBy() string",
		"func (x *InvoiceEntity) GetNote() emigo.Nullable[string]",
		"var _ Auditable = (*InvoiceEntity)(nil)",
		// one? is a plain struct in the entity's own storage types, and the interface is typed the same way
		"func (x *InvoiceEntity) GetTag() TagEntity",
		"var _ Tagged = (*InvoiceEntity)(nil)",
	} {
		if !strings.Contains(entity, want) {
			t.Errorf("InvoiceEntity.go missing %q", want)
		}
	}
	if tagged := files["Tagged.go"]; !strings.Contains(tagged, "GetTag() TagEntity") || strings.Contains(tagged, "emigo") {
		t.Errorf("Tagged.go should be typed like the entity and import nothing it doesn't use:\n%s", tagged)
	}

	dto := files["InvoiceDto.go"]
	if !strings.Contains(dto, "var _ Auditable = (*InvoiceDto)(nil)") {
		t.Errorf("derived dto should satisfy Auditable")
	}
	if strings.Contains(dto, "Tagged") {
		t.Errorf("derived dto must not claim an interface whose types it doesn't share")
	}
	if strings.Contains(files["InvoiceOptionalDto.go"], "var _ ") {
		t.Errorf("optional dto must not implement interfaces")
	}
}
