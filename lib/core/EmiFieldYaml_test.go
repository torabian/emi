package core

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestFieldDescriptionStringOrMap(t *testing.T) {
	m, err := ReadEmiFromString(`
dtos:
  - name: A
    fields:
      - {name: plain, type: string, description: just text}
      - name: localized
        type: string
        description:
          pl: Po polsku
          en: In English
      - name: noEnglish
        type: string
        description: {fa: farsi, de: deutsch}
      - name: none
        type: string
      - name: group
        type: object
        description: {en: Group}
        fields:
          - {name: inner, type: string, description: {en: Inner, ru: Vnutri}}
`)
	if err != nil {
		t.Fatal(err)
	}
	f := m.Dto[0].Fields
	if f[0].Description != "just text" || f[0].Descriptions != nil {
		t.Fatalf("plain: %+v", f[0])
	}
	if f[1].Description != "In English" || f[1].Descriptions["pl"] != "Po polsku" || len(f[1].Descriptions) != 2 {
		t.Fatalf("localized: %+v", f[1])
	}
	if f[2].Description != "deutsch" { // no "en": first locale alphabetically
		t.Fatalf("noEnglish: %+v", f[2])
	}
	if f[3].Description != "" || f[3].Descriptions != nil {
		t.Fatalf("none: %+v", f[3])
	}
	if in := f[4].Fields[0]; in.Description != "Inner" || in.Descriptions["ru"] != "Vnutri" {
		t.Fatalf("nested: %+v", in)
	}
	// Everything else on the field is still parsed.
	if f[1].Type != "string" || f[1].Name != "localized" {
		t.Fatalf("other keys lost: %+v", f[1])
	}
}

func TestFieldDescriptionRejectsNonStringLocale(t *testing.T) {
	_, err := ReadEmiFromString("dtos:\n  - name: A\n    fields:\n      - {name: x, type: string, description: {en: [a]}}\n")
	if err == nil || !strings.Contains(err.Error(), "description") {
		t.Fatalf("got %v", err)
	}
}

func TestFieldDescriptionMarshalRoundTrip(t *testing.T) {
	src := "name: x\ntype: string\ndescription:\n  en: E\n  pl: P\n"
	var f EmiField
	if err := yaml.Unmarshal([]byte(src), &f); err != nil {
		t.Fatal(err)
	}
	out, err := yaml.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	var back EmiField
	if err := yaml.Unmarshal(out, &back); err != nil {
		t.Fatal(err)
	}
	if back.Description != "E" || back.Descriptions["pl"] != "P" || back.Name != "x" {
		t.Fatalf("round trip lost data:\n%s\n%+v", out, back)
	}

	plain := EmiField{Name: "p", Description: "text"}
	out, _ = yaml.Marshal(plain)
	if !strings.Contains(string(out), "description: text") {
		t.Fatalf("plain marshal: %s", out)
	}
}

// EmiColumn and EmiFieldTemplate embed EmiField; their own keys must survive it.
func TestFieldDescriptionInEmbeddingTypes(t *testing.T) {
	var col EmiColumn
	if err := yaml.Unmarshal([]byte("name: email\ncolumn: u.email\nselected: true\ndescription: {en: Mail}\n"), &col); err != nil {
		t.Fatal(err)
	}
	if col.Column != "u.email" || !col.Selected || col.Description != "Mail" || col.Descriptions["en"] != "Mail" {
		t.Fatalf("column: %+v", col)
	}
	out, _ := yaml.Marshal(col)
	var colBack EmiColumn
	if err := yaml.Unmarshal(out, &colBack); err != nil || colBack.Column != "u.email" || !colBack.Selected || colBack.Descriptions["en"] != "Mail" {
		t.Fatalf("column round trip: %v\n%s", err, out)
	}

	var tpl EmiFieldTemplate
	if err := yaml.Unmarshal([]byte("name: t\ncontext: entity\ntype: string\ndescription: {en: T, pl: TP}\n"), &tpl); err != nil {
		t.Fatal(err)
	}
	if tpl.Context != "entity" || tpl.Name != "t" || tpl.Descriptions["pl"] != "TP" {
		t.Fatalf("template: %+v", tpl)
	}
	out, _ = yaml.Marshal(tpl)
	var tplBack EmiFieldTemplate
	if err := yaml.Unmarshal(out, &tplBack); err != nil || tplBack.Context != "entity" || tplBack.Descriptions["pl"] != "TP" {
		t.Fatalf("template round trip: %v\n%s", err, out)
	}
}

func TestFieldUseKeepsLocalizedDescription(t *testing.T) {
	m, err := ReadEmiFromString(`
templates:
  fields:
    - {name: tpl, type: string, description: {en: From template, pl: Z szablonu}}
dtos:
  - name: A
    fields:
      - {name: inherits, use: tpl}
      - {name: ownText, use: tpl, description: mine}
      - {name: ownMap, use: tpl, description: {en: Mine, de: Meins}}
`)
	if err != nil {
		t.Fatal(err)
	}
	f := m.Dto[0].Fields
	if f[0].Description != "From template" || f[0].Descriptions["pl"] != "Z szablonu" {
		t.Fatalf("inherits: %+v", f[0])
	}
	if f[1].Description != "mine" || f[1].Descriptions != nil {
		t.Fatalf("ownText: %+v", f[1])
	}
	if f[2].Description != "Mine" || f[2].Descriptions["de"] != "Meins" || len(f[2].Descriptions) != 2 {
		t.Fatalf("ownMap: %+v", f[2])
	}
}
