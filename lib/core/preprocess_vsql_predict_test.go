package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVsqlPredictPopulatesColumnsFromPlainSelect(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name:    "listUsers",
				Predict: true,
				Query:   "SELECT id, email, first_name FROM users WHERE id = {{ .Id }}",
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	cols := m.Vsqls[0].Columns
	if len(cols) != 3 {
		t.Fatalf("expected 3 predicted columns, got %d: %+v", len(cols), cols)
	}

	want := []struct{ name, column string }{
		{"id", "id"},
		{"email", "email"},
		{"firstName", "first_name"},
	}
	for i, w := range want {
		if cols[i].Name != w.name {
			t.Fatalf("column %d: got name %q, want %q", i, cols[i].Name, w.name)
		}
		if cols[i].Column != w.column {
			t.Fatalf("column %d: got Column %q, want %q", i, cols[i].Column, w.column)
		}
		if !cols[i].Selected {
			t.Fatalf("column %d: expected Selected true (the query already selects it unconditionally)", i)
		}
		if cols[i].Type != FieldTypeString {
			t.Fatalf("column %d: expected default type string, got %q", i, cols[i].Type)
		}
	}
}

func TestVsqlPredictHonorsFieldFuncTypeAndOptional(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name:    "userOrderCounts",
				Predict: true,
				Query: "SELECT u.user_id, field(count(o.order_id), 'int64', 'totalOrders', true) " +
					"FROM users u LEFT JOIN orders o ON u.user_id = o.user_id GROUP BY u.user_id",
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	cols := m.Vsqls[0].Columns
	if len(cols) != 2 {
		t.Fatalf("expected 2 columns, got %+v", cols)
	}

	total := cols[1]
	if total.Name != "totalOrders" {
		t.Fatalf("expected field()'s explicit name, got %q", total.Name)
	}
	if total.Column != "count(o.order_id)" {
		t.Fatalf("expected the aggregate expression preserved verbatim as Column, got %q", total.Column)
	}
	if total.Type != "int64?" {
		t.Fatalf("expected field()'s optional=true to produce int64?, got %q", total.Type)
	}
}

func TestVsqlPredictSkipsWhenColumnsAlreadyDeclared(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name:    "listUsers",
				Predict: true,
				Query:   "SELECT id, email FROM users",
				Columns: []*EmiColumn{
					{EmiField: EmiField{Name: "handWritten", Type: FieldTypeString}, Selected: true},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	cols := m.Vsqls[0].Columns
	if len(cols) != 1 || cols[0].Name != "handWritten" {
		t.Fatalf("expected hand-written Columns to be left untouched, got %+v", cols)
	}
}

func TestVsqlPredictRejectsNonSelect(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "insertUser", Predict: true, Query: "INSERT INTO users (email) VALUES ({{ sql .Email }})"},
		},
	}

	if err := m.Preprocess(); err == nil {
		t.Fatalf("expected an error - predict only supports a plain SELECT")
	}
}

func TestVsqlPredictWithoutFlagLeavesColumnsEmpty(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "listUsers", Query: "SELECT id, email FROM users"},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}
	if len(m.Vsqls[0].Columns) != 0 {
		t.Fatalf("expected no columns without predict: true, got %+v", m.Vsqls[0].Columns)
	}
}

func TestVsqlPredictReadsFromQueryNameOnDisk(t *testing.T) {
	dir := t.TempDir()
	sqlPath := filepath.Join(dir, "list_users.sql")
	if err := os.WriteFile(sqlPath, []byte("SELECT id, email FROM users"), 0644); err != nil {
		t.Fatalf("writing fixture: %v", err)
	}
	modulePath := filepath.Join(dir, "module.emi.yml")

	m := &Emi{
		SourcePath: modulePath,
		Vsqls: []EmiVsql{
			{Name: "listUsers", Predict: true, QueryName: "list_users.sql"},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	cols := m.Vsqls[0].Columns
	if len(cols) != 2 || cols[0].Name != "id" || cols[1].Name != "email" {
		t.Fatalf("got %+v", cols)
	}
}

// TestVsqlPredictIsIdempotentAcrossPreprocessPasses guards against a second
// Preprocess() re-running detection and duplicating/re-deriving columns
// (Preprocess is documented as safe to call more than once).
func TestVsqlPredictIsIdempotentAcrossPreprocessPasses(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "listUsers", Predict: true, Query: "SELECT id, email FROM users"},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (1st): %v", err)
	}
	first := m.Vsqls[0].Columns
	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (2nd): %v", err)
	}
	if len(m.Vsqls[0].Columns) != len(first) {
		t.Fatalf("expected the second pass to leave Columns untouched, got %+v", m.Vsqls[0].Columns)
	}
}

// TestVsqlPredictStarLeavesColumnsEmpty covers "we cannot determine columns
// at all" - a wildcard select is a valid, non-error case: no Columns get
// populated, and no Row DTO/picker gets generated as a result, exactly as if
// predict had never been set.
func TestVsqlPredictStarLeavesColumnsEmpty(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "listUsers", Predict: true, Query: "SELECT * FROM users"},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}
	if len(m.Vsqls[0].Columns) != 0 {
		t.Fatalf("expected no columns for SELECT *, got %+v", m.Vsqls[0].Columns)
	}
}
