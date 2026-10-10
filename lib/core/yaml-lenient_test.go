package core

import "testing"

func TestUnmarshalLenientQuotesAColonInProse(t *testing.T) {
	var out struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Items       []struct {
			Title string `yaml:"title"`
		} `yaml:"items"`
	}
	src := `name: formbuilder
description: A form named by another answer: its uniqueId, and the name it's shown with
items:
  - title: Note: this one too
`
	if err := unmarshalLenient([]byte(src), &out); err != nil {
		t.Fatal(err)
	}
	if out.Description != "A form named by another answer: its uniqueId, and the name it's shown with" {
		t.Errorf("description: %q", out.Description)
	}
	if len(out.Items) != 1 || out.Items[0].Title != "Note: this one too" {
		t.Errorf("items: %+v", out.Items)
	}
}

func TestUnmarshalLenientLeavesBlockScalarsAndValidFilesAlone(t *testing.T) {
	var out struct {
		Description string `yaml:"description"`
		Other       string `yaml:"other"`
		Broken      string `yaml:"broken"`
	}
	// the block scalar has a "key: value" looking line, and the file has a colon problem elsewhere
	src := `description: >-
  First line.
  Remark: kept as written.
other: 'already: quoted'
broken: this one: needs quoting
`
	if err := unmarshalLenient([]byte(src), &out); err != nil {
		t.Fatal(err)
	}
	if out.Description != "First line. Remark: kept as written." {
		t.Errorf("a block scalar must not be touched, got %q", out.Description)
	}
	if out.Other != "already: quoted" || out.Broken != "this one: needs quoting" {
		t.Errorf("got %q and %q", out.Other, out.Broken)
	}
}

func TestUnmarshalLenientReportsTheRealError(t *testing.T) {
	var out map[string]any
	// not a colon problem: an unterminated flow sequence
	if err := unmarshalLenient([]byte("a: [1, 2\nb: x: y\n"), &out); err == nil {
		t.Fatal("a file which is still invalid after quoting must fail")
	}
}
