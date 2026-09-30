package core

import (
	"strings"
	"testing"
)

const interfaceFixture = `
templates:
  fields:
    - {name: purchasable, context: dto, type: 'one?', target: PurchasableDto, module: ctb, provider: example.com/ctb}
    - {name: purchasable, context: entity, type: 'one?', target: PurchasableEntity, module: ctb, provider: example.com/ctb}
interfaces:
  - name: sellable
    description: Can be sold.
    fields:
      - {name: product, use: purchasable}
      - {name: sku, type: string}
  - name: auditable
    fields:
      - {name: createdBy, type: 'string?'}
`

func TestInterfaceFieldsAreIncluded(t *testing.T) {
	m, err := ReadEmiFromString(interfaceFixture + `
dtos:
  - name: Score
    implements: [sellable, auditable]
    fields:
      - {name: title, type: string}
      - {name: sku, type: string, description: own docs}
`)
	if err != nil {
		t.Fatal(err)
	}
	names := []string{}
	for _, f := range m.Dto[0].Fields {
		names = append(names, f.Name)
	}
	// interface fields first, in implements order; the dto's own declaration of sku is kept.
	if strings.Join(names, ",") != "product,createdBy,title,sku" {
		t.Fatalf("fields: %v", names)
	}
	if p := m.Dto[0].Fields[0]; p.Target != "PurchasableDto" || p.Use != "" {
		t.Fatalf("interface template not resolved: %+v", p)
	}
	if m.Dto[0].Fields[3].Description != "own docs" {
		t.Fatalf("own field replaced: %+v", m.Dto[0].Fields[3])
	}

	// Preprocessing again adds nothing.
	if err := m.Preprocess(); err != nil || len(m.Dto[0].Fields) != 4 {
		t.Fatalf("second pass: %v %d", err, len(m.Dto[0].Fields))
	}
}

func TestInterfaceFieldsAreCopiedPerImplementer(t *testing.T) {
	m, err := ReadEmiFromString(interfaceFixture + `
dtos:
  - {name: A, implements: [auditable]}
  - {name: B, implements: [auditable]}
`)
	if err != nil {
		t.Fatal(err)
	}
	m.Dto[0].Fields[0].Name = "changed"
	if m.Dto[1].Fields[0].Name != "createdBy" || m.Interfaces[1].Fields[0].Name != "createdBy" {
		t.Fatal("implementers share the interface's field")
	}
}

func TestInterfaceEntityContext(t *testing.T) {
	// A template-built field only fits the context it resolved for.
	_, err := ReadEmiFromString(interfaceFixture + `
entities:
  - {name: score, implements: [sellable]}
`)
	if err == nil || !strings.Contains(err.Error(), "context") {
		t.Fatalf("want context error, got %v", err)
	}

	// Plain fields fit any owner; and an entity-context interface resolves the entity variant.
	m, err := ReadEmiFromString(interfaceFixture + `
entities:
  - {name: score, implements: [auditable]}
`)
	if err != nil || m.Entities[0].Fields[0].Name != "createdBy" {
		t.Fatalf("entity + plain interface: %v", err)
	}
	m, err = ReadEmiFromString(`
templates:
  fields:
    - {name: purchasable, context: entity, type: 'one?', target: PurchasableEntity}
interfaces:
  - name: sellableEntity
    context: entity
    fields: [{name: product, use: purchasable}]
entities:
  - {name: score, implements: [sellableEntity]}
`)
	if err != nil || m.Entities[0].Fields[0].Target != "PurchasableEntity" {
		t.Fatalf("entity-context interface: %v", err)
	}
}

func TestInterfaceErrors(t *testing.T) {
	cases := map[string]struct{ yaml, want string }{
		"unknown":         {"dtos:\n  - {name: A, implements: [nope]}\n", "no interfaces"},
		"unknownWith":     {interfaceFixture + "dtos:\n  - {name: A, implements: [nope]}\n", "unknown interface"},
		"typeConflict":    {interfaceFixture + "dtos:\n  - name: A\n    implements: [sellable]\n    fields: [{name: sku, type: int64}]\n", "requires"},
		"mapWithChildren": {"interfaces:\n  - name: x\n    fields: [{name: m, type: map, fields: [{name: a, type: string}]}]\n", "can't be shared"},
		"duplicateChild":  {"interfaces:\n  - name: x\n    fields: [{name: o, type: object, fields: [{name: a, type: string}, {name: a, type: string}]}]\n", "child"},
		"duplicateName":   {"interfaces:\n  - {name: x}\n  - {name: x}\n", "declared twice"},
		"duplicateFld":    {"interfaces:\n  - name: x\n    fields: [{name: a, type: string}, {name: a, type: string}]\n", "twice"},
		"badContext":      {"interfaces:\n  - {name: x, context: nope}\n", "context must be"},
	}
	for name, c := range cases {
		if _, err := ReadEmiFromString(c.yaml); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: want error containing %q, got %v", name, c.want, err)
		}
	}
}

func TestInterfacesMergeThroughExtends(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "base.emi.yml", "interfaces:\n  - {name: auditable, fields: [{name: createdBy, type: 'string?'}]}\n")
	p := writeEmi(t, dir, "m.emi.yml", "extends: [{from: base.emi.yml}]\ndtos:\n  - {name: A, implements: [auditable]}\n")
	m, err := ReadEmiFromFile(p)
	if err != nil || len(m.Interfaces) != 1 || m.Dto[0].Fields[0].Name != "createdBy" {
		t.Fatalf("%v %+v", err, m.Interfaces)
	}
}

func TestInterfaceObjectAndArrayFields(t *testing.T) {
	m, err := ReadEmiFromString(`
templates:
  fields:
    - {name: label, type: string, description: A label}
interfaces:
  - name: titlable
    fields:
      - name: content
        type: object
        fields:
          - {name: text, use: label}
          - name: parts
            type: array
            fields: [{name: kind, type: string}]
      - name: notes
        type: 'array?'
        fields: [{name: note, type: string}]
dtos:
  - {name: Cloth, implements: [titlable]}
  - name: Shoe
    implements: [titlable]
    fields:
      - name: content
        type: object
        fields:
          - {name: text, type: string}
          - name: parts
            type: array
            fields: [{name: kind, type: string}]
`)
	if err != nil {
		t.Fatal(err)
	}
	cloth := m.Dto[0].Fields
	if len(cloth) != 2 || cloth[0].Name != "content" || len(cloth[0].Fields) != 2 || cloth[0].Fields[0].Description != "A label" {
		t.Fatalf("cloth: %+v", cloth[0])
	}
	// children are copied: changing one implementer's nested field leaves the interface alone
	cloth[0].Fields[0].Name = "changed"
	if m.Interfaces[0].Fields[0].Fields[0].Name != "text" {
		t.Fatal("nested fields are shared with the interface")
	}
	// Shoe declared the same shape itself: kept as its own, nothing duplicated.
	if names := []string{m.Dto[1].Fields[0].Name, m.Dto[1].Fields[1].Name}; len(m.Dto[1].Fields) != 2 || names[0] != "notes" && names[1] != "notes" {
		t.Fatalf("shoe fields: %+v", m.Dto[1].Fields)
	}
}

func TestInterfaceObjectShapeMismatch(t *testing.T) {
	_, err := ReadEmiFromString(`
interfaces:
  - name: titlable
    fields: [{name: content, type: object, fields: [{name: text, type: string}]}]
dtos:
  - name: Shoe
    implements: [titlable]
    fields: [{name: content, type: object, fields: [{name: text, type: int64}]}]
`)
	if err == nil || !strings.Contains(err.Error(), "requires") {
		t.Fatalf("want shape mismatch, got %v", err)
	}
}

const entityInterfaceFixture = `
interfaces:
  - name: auditable
    fields: [{name: createdBy, type: string}]
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

// The plain dto derived from an entity implements the interfaces it can honestly satisfy
// (same types), and only those: an interface with a relation stays on the entity, and the
// optional dto - every field nullable - never implements any.
func TestEntityDerivedDtoImplementsOnlyWhatItSatisfies(t *testing.T) {
	m, err := ReadEmiFromString(entityInterfaceFixture)
	if err != nil {
		t.Fatal(err)
	}
	if fields := m.Entities[1].Fields; len(fields) != 3 || fields[0].Name != "createdBy" || fields[1].Name != "tag" {
		t.Fatalf("entity fields: %+v", fields)
	}

	m.preprocessEntityDtos()
	m.preprocessEntityOptionalDtos()

	byName := map[string]EmiDto{}
	for _, d := range m.Dto {
		byName[d.Name] = d
	}
	if got := byName["invoice"].Implements; len(got) != 1 || got[0] != "auditable" {
		t.Fatalf("derived dto implements %v, want [auditable]", got)
	}
	if got := byName[entityOptionalDtoName(m.Entities[1])].Implements; len(got) != 0 {
		t.Fatalf("optional dto must not implement anything, got %v", got)
	}
	if got := byName["tag"].Implements; len(got) != 0 {
		t.Fatalf("an entity with no interfaces derives a dto with none, got %v", got)
	}

	// preprocessing again still accepts the derived dto (it is checked like any dto)
	if err := m.Preprocess(); err != nil {
		t.Fatalf("second pass: %v", err)
	}
}

func TestEntityCannotImplementInterfaceWithInlineNestedFields(t *testing.T) {
	_, err := ReadEmiFromString(`
interfaces:
  - name: titlable
    fields: [{name: content, type: object, fields: [{name: text, type: string}]}]
entities:
  - {name: invoice, implements: [titlable]}
`)
	if err == nil || !strings.Contains(err.Error(), "child tables") {
		t.Fatalf("want entity nested-field error, got %v", err)
	}
}
