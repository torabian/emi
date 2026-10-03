package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeEmi(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestExtendsMergeOrderAndOverride(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "defs/a.emi.yml", `
description: from a
complexes:
  - {compiler: go, namespace: c, name: X, location: a-loc}
targets:
  - {compiler: go, output: ./defs, flags: {pkg: a}}
`)
	writeEmi(t, dir, "defs/b.emi.yml", `
extends: [{from: ./a.emi.yml}]
description: from b
complexes:
  - {compiler: go, namespace: c, name: X, location: b-loc}
  - {compiler: go, namespace: c, name: Y, location: b-loc}
`)
	root := writeEmi(t, dir, "m/m.emi.yml", `
name: m
extends: [{from: ../defs/b.emi.yml}]
targets:
  - {compiler: go, output: ./defs, flags: {pkg: mine}}
  - {compiler: js, output: ./sdk}
`)
	m, err := ReadEmiFromFile(root)
	if err != nil {
		t.Fatal(err)
	}
	if m.Name != "m" || m.Description != "from b" || m.Extends != nil {
		t.Fatalf("scalars/extends wrong: %+v", m.Name)
	}
	if len(m.Complexes) != 2 || m.Complexes[0].Location != "b-loc" || m.Complexes[1].Name != "Y" {
		t.Fatalf("complexes: %+v", m.Complexes)
	}
	if len(m.Targets) != 2 || m.Targets[0].Flags["pkg"] != "mine" || m.Targets[1].Compiler != "js" {
		t.Fatalf("targets: %+v", m.Targets)
	}
}

func TestExtendsCycleRejected(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "a.emi.yml", "extends: [{from: b.emi.yml}]\n")
	writeEmi(t, dir, "b.emi.yml", "extends: [{from: a.emi.yml}]\n")
	if _, err := ReadEmiFromFile(filepath.Join(dir, "a.emi.yml")); err == nil || !strings.Contains(err.Error(), "circular") {
		t.Fatalf("want circular error, got %v", err)
	}
}

func TestExtendsSelfRejected(t *testing.T) {
	dir := t.TempDir()
	p := writeEmi(t, dir, "a.emi.yml", "extends: [{from: a.emi.yml}]\n")
	if _, err := ReadEmiFromFile(p); err == nil || !strings.Contains(err.Error(), "circular") {
		t.Fatalf("want circular error, got %v", err)
	}
}

func TestExtendsDiamondAllowed(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "base.emi.yml", "enums:\n  - {name: E}\n")
	writeEmi(t, dir, "l.emi.yml", "extends: [{from: base.emi.yml}]\n")
	writeEmi(t, dir, "r.emi.yml", "extends: [{from: base.emi.yml}]\n")
	p := writeEmi(t, dir, "top.emi.yml", "extends: [{from: l.emi.yml}, {from: r.emi.yml}]\n")
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Enums) != 1 {
		t.Fatalf("enums: %+v", m.Enums)
	}
}

func TestExtendsExpandsComplexIncludeRelativeToDeclaringFile(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "shared/c.emi.yml", "complexes:\n  - {compiler: go, namespace: c, name: Z, location: l}\n")
	writeEmi(t, dir, "shared/base.emi.yml", "complexes:\n  - include: ./c.emi.yml\n")
	p := writeEmi(t, dir, "app/m.emi.yml", "extends: [{from: ../shared/base.emi.yml}]\n")
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Complexes) != 1 || m.Complexes[0].Name != "Z" {
		t.Fatalf("complexes: %+v", m.Complexes)
	}
}

func TestExtendsMissingFile(t *testing.T) {
	if _, err := ReadEmiFromString("extends: [{from: /nope/x.yml}]\n"); err == nil {
		t.Fatal("want error")
	}
}

func TestExtendsParams(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "defs/base.emi.yml", `
targets:
  - {compiler: go, output: ./defs, flags: {pkg: "{{ name }}defs"}}
  - {compiler: js, output: "../ui/{{name}}/{{missing}}"}
`)
	writeEmi(t, dir, "defs/mid.emi.yml", `
extends:
  - from: ./base.emi.yml
    params: {name: "{{who}}"}
`)
	p := writeEmi(t, dir, "m/m.emi.yml", `
extends:
  - from: ../defs/mid.emi.yml
    params: {who: "wallet: x"}
`)
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if m.Targets[0].Flags["pkg"] != "wallet: xdefs" || m.Targets[1].Output != "../ui/wallet: x/{{missing}}" {
		t.Fatalf("targets: %+v", m.Targets)
	}
}

func TestExtendsParamDefaults(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "base.emi.yml", `
targets:
  - {compiler: go, output: "{{output|./defs}}", flags: {pkg: "{{pkg}}"}}
`)
	// No params at all: fallbacks still apply, params without one stay as written.
	p := writeEmi(t, dir, "a.emi.yml", "extends: [{from: base.emi.yml}]\n")
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if m.Targets[0].Output != "./defs" || m.Targets[0].Flags["pkg"] != "{{pkg}}" {
		t.Fatalf("targets: %+v", m.Targets)
	}
	p = writeEmi(t, dir, "b.emi.yml", "extends: [{from: base.emi.yml, params: {output: ./x}}]\n")
	if m, err = ReadEmiFromFile(p); err != nil || m.Targets[0].Output != "./x" {
		t.Fatalf("%v %+v", err, m.Targets)
	}
}

// The extending module's own definitions are appended after the extended ones, and
// override any extended entry they share an identity with - for every list, not just
// complexes/targets.
func TestExtendsOwnContentAppendsAndOverrides(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "base.emi.yml", `
actions:
  - {name: shared, url: /base}
  - {name: baseOnly, url: /b}
entities:
  - {name: BaseEntity}
permissions:
  - {key: read}
config:
  - {name: PORT}
`)
	p := writeEmi(t, dir, "m.emi.yml", `
extends: [{from: base.emi.yml}]
actions:
  - {name: shared, url: /own}
  - {name: ownOnly, url: /o}
entities:
  - {name: OwnEntity}
permissions:
  - {key: write}
config:
  - {name: PORT, description: mine}
`)
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	names := []string{}
	for _, a := range m.Actions {
		names = append(names, a.Name+"="+a.Url)
	}
	if strings.Join(names, ",") != "shared=/own,baseOnly=/b,ownOnly=/o" {
		t.Fatalf("actions: %v", names)
	}
	if len(m.Entities) != 2 || m.Entities[0].Name != "BaseEntity" || m.Entities[1].Name != "OwnEntity" {
		t.Fatalf("entities: %+v", m.Entities)
	}
	if len(m.Permissions) != 2 || m.Permissions[1].Key != "write" {
		t.Fatalf("permissions: %+v", m.Permissions)
	}
	if len(m.Config) != 1 || m.Config[0].Description != "mine" {
		t.Fatalf("config: %+v", m.Config)
	}
}

// {{extends.dir}} keeps a shared file's outputs pointing at the same place no matter how
// deep the extending module is.
func TestExtendsDirParam(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "modules/defs/base.emi.yml", `
targets:
  - {compiler: js, output: "{{extends.dir}}/../../ui/sdk"}
`)
	writeEmi(t, dir, "modules/defs/mid.emi.yml", "extends: [{from: ./base.emi.yml}]\n")
	for rel, want := range map[string]string{
		"modules/shallow/m.emi.yml":      "../defs/../../ui/sdk",
		"modules/finance/deep/m.emi.yml": "../../defs/../../ui/sdk",
	} {
		up := "../defs/mid.emi.yml"
		if strings.Count(rel, "/") == 3 {
			up = "../../defs/mid.emi.yml"
		}
		p := writeEmi(t, dir, rel, "extends: [{from: "+up+"}]\n")
		m, err := ReadEmiFromFile(p)
		if err != nil {
			t.Fatal(err)
		}
		got := m.Targets[0].Output
		if got != want || filepath.Clean(filepath.Join(filepath.Dir(p), got)) != filepath.Join(dir, "ui/sdk") {
			t.Fatalf("%s: output %q", rel, got)
		}
	}
}

// A module can override an inherited target by writing the same output directory, even
// when the extended file spelled the path differently.
func TestExtendsTargetIdentityIgnoresPathSpelling(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "defs/base.emi.yml", `
targets:
  - {compiler: js, output: "{{extends.dir}}/../ui/sdk/", clean: true}
`)
	p := writeEmi(t, dir, "m/m.emi.yml", `
extends: [{from: ../defs/base.emi.yml}]
targets:
  - {compiler: js, output: ../ui/sdk}
`)
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Targets) != 1 || m.Targets[0].Clean {
		t.Fatalf("targets: %+v", m.Targets)
	}
}

func TestFieldUse(t *testing.T) {
	dir := t.TempDir()
	writeEmi(t, dir, "tpl.emi.yml", `
templates:
  fields:
    - name: purchasable
      context: dto
      type: one?
      target: PurchasableDto
      module: clicktobuydefs
      jsProvider: "{{jsProvider|}}"
      tags: {validate: "-", keep: base}
      description: from template
    - name: purchasable
      context: entity
      type: one?
      target: PurchasableEntity
      module: clicktobuydefs
    - name: anywhere
      type: string
`)
	p := writeEmi(t, dir, "m.emi.yml", `
extends:
  - from: tpl.emi.yml
    params: {jsProvider: "@/sdk/PurchasableDto"}
dtos:
  - name: ScoreDto
    fields:
      - {name: product, use: purchasable, description: mine, tags: {keep: own, extra: x}}
      - {name: plain, use: anywhere}
entities:
  - name: score
    fields:
      - {name: product, use: purchasable}
`)
	m, err := ReadEmiFromFile(p)
	if err != nil {
		t.Fatal(err)
	}
	dto := m.Dto[0].Fields[0]
	if dto.Use != "" || dto.Name != "product" || dto.Target != "PurchasableDto" || dto.JsProvider != "@/sdk/PurchasableDto" ||
		dto.Description != "mine" || dto.Tags["validate"] != "-" || dto.Tags["keep"] != "own" || dto.Tags["extra"] != "x" {
		t.Fatalf("dto field: %+v", dto)
	}
	if f := m.Dto[0].Fields[1]; f.Type != "string" || f.Name != "plain" {
		t.Fatalf("general template: %+v", f)
	}
	if e := m.Entities[0].Fields[0]; e.Target != "PurchasableEntity" || e.Tags != nil || e.Name != "product" {
		t.Fatalf("entity field: %+v", e)
	}
}

func TestFieldUseUnknownTemplate(t *testing.T) {
	_, err := ReadEmiFromString("dtos:\n  - name: A\n    fields:\n      - {name: x, use: nope}\n")
	if err == nil || !strings.Contains(err.Error(), "unknown field template") {
		t.Fatalf("got %v", err)
	}
}
