package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderVsqlDebugScalarParam(t *testing.T) {
	vsql := &EmiVsql{
		Name:  "getUser",
		Query: "SELECT * FROM users WHERE id = {{ .Id }} AND name = {{ sql .Name }}",
		Params: []*EmiField{
			{Name: "id", Type: FieldTypeInt64},
			{Name: "name", Type: FieldTypeString},
		},
	}

	got, err := RenderVsqlDebug(vsql, "", `{"id": 42, "name": "O'Brien"}`, "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	want := `SELECT * FROM users WHERE id = 42 AND name = 'O''Brien'`
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRenderVsqlDebugArrayParam(t *testing.T) {
	vsql := &EmiVsql{
		Name: "insertUsers",
		Query: "INSERT INTO users (email) VALUES" +
			"{{ range $i, $u := .Users.Items }}{{ if $i }},{{ end }}({{ sql $u.Email }}){{ end }};",
		Params: []*EmiField{
			{Name: "users", Type: FieldTypeArray, Fields: []*EmiField{
				{Name: "email", Type: FieldTypeString},
			}},
		},
	}

	got, err := RenderVsqlDebug(vsql, "", `{"users": [{"email": "a@x.com"}, {"email": "b@x.com"}]}`, "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	want := "INSERT INTO users (email) VALUES('a@x.com'),('b@x.com');"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRenderVsqlDebugColumnsDefaultAndOverride(t *testing.T) {
	vsql := &EmiVsql{
		Name:  "selectUsers",
		Query: "SELECT {{ .Columns.Cols }} FROM users;",
		Columns: []*EmiColumn{
			{EmiField: EmiField{Name: "id", Type: FieldTypeInt64}, Selected: true},
			{EmiField: EmiField{Name: "email", Type: FieldTypeString}, Selected: false},
		},
	}

	got, err := RenderVsqlDebug(vsql, "", "", "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug (defaults): %v", err)
	}
	if got != "SELECT id FROM users;" {
		t.Fatalf("expected default selection, got %q", got)
	}

	got, err = RenderVsqlDebug(vsql, "", "", "", "email")
	if err != nil {
		t.Fatalf("RenderVsqlDebug (--select): %v", err)
	}
	if got != "SELECT email FROM users;" {
		t.Fatalf("expected --select to override defaults entirely, got %q", got)
	}
}

func TestRenderVsqlDebugColumnSelectedFlag(t *testing.T) {
	vsql := &EmiVsql{
		Name:  "selectUsers",
		Query: "{{ if .Columns.Email.Selected }}has email{{ else }}no email{{ end }}",
		Columns: []*EmiColumn{
			{EmiField: EmiField{Name: "email", Type: FieldTypeString}, Selected: false},
		},
	}

	got, err := RenderVsqlDebug(vsql, "", "", "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	if got != "no email" {
		t.Fatalf("got %q", got)
	}

	got, err = RenderVsqlDebug(vsql, "", "", "", "email")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	if got != "has email" {
		t.Fatalf("got %q", got)
	}
}

func TestRenderVsqlDebugInAndParamsMerge(t *testing.T) {
	vsql := &EmiVsql{
		Name:  "search",
		Query: "{{ .Limit }}/{{ .Offset }}",
		Params: []*EmiField{
			{Name: "limit", Type: FieldTypeInt},
			{Name: "offset", Type: FieldTypeInt},
		},
	}

	// --params wins the "limit" collision; --in fills in "offset".
	got, err := RenderVsqlDebug(vsql, "", `{"limit": 10}`, `{"limit": 999, "offset": 20}`, "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	if got != "10/20" {
		t.Fatalf("got %q, want 10/20", got)
	}
}

func TestRenderVsqlDebugObjectParam(t *testing.T) {
	vsql := &EmiVsql{
		Name:  "upsertPrefs",
		Query: "{{ .Preferences.Theme }}",
		Params: []*EmiField{
			{Name: "preferences", Type: FieldTypeObject, Fields: []*EmiField{
				{Name: "theme", Type: FieldTypeString},
			}},
		},
	}

	got, err := RenderVsqlDebug(vsql, "", `{"preferences": {"theme": "dark"}}`, "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	if got != "dark" {
		t.Fatalf("got %q", got)
	}
}

func TestRenderVsqlDebugQueryNameReadsFromDisk(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "search.sql"), []byte("SELECT {{ .Id }}"), 0644); err != nil {
		t.Fatalf("writing fixture: %v", err)
	}

	vsql := &EmiVsql{
		Name:      "search",
		QueryName: "search.sql",
		Params:    []*EmiField{{Name: "id", Type: FieldTypeInt64}},
	}

	got, err := RenderVsqlDebug(vsql, dir, `{"id": 5}`, "", "")
	if err != nil {
		t.Fatalf("RenderVsqlDebug: %v", err)
	}
	if got != "SELECT 5" {
		t.Fatalf("got %q", got)
	}
}

func TestRenderVsqlDebugMissingQueryErrors(t *testing.T) {
	vsql := &EmiVsql{Name: "empty"}
	if _, err := RenderVsqlDebug(vsql, "", "", "", ""); err == nil {
		t.Fatalf("expected an error for a vsql with neither query nor queryName")
	}
}

func TestRenderVsqlDebugInvalidParamsJSONErrors(t *testing.T) {
	vsql := &EmiVsql{Name: "x", Query: "select 1"}
	_, err := RenderVsqlDebug(vsql, "", "{not json", "", "")
	if err == nil {
		t.Fatalf("expected an error for invalid --params JSON")
	}
	if !strings.Contains(err.Error(), "--params") {
		t.Fatalf("expected the error to name --params, got: %v", err)
	}
}
