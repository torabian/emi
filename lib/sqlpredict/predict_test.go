package sqlpredict

import "testing"

func TestDetectSelectColumnsPlain(t *testing.T) {
	cols, err := DetectSelectColumns("SELECT id, email, first_name FROM users WHERE id = 1")
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	want := []struct{ name, source, typ string }{
		{"Id", "id", "string"},
		{"Email", "email", "string"},
		{"FirstName", "first_name", "string"},
	}
	if len(cols) != len(want) {
		t.Fatalf("got %d columns, want %d: %+v", len(cols), len(want), cols)
	}
	for i, w := range want {
		if cols[i].Name != w.name || cols[i].Source != w.source || cols[i].Type != w.typ {
			t.Fatalf("col %d: got %+v, want %+v", i, cols[i], w)
		}
	}
}

func TestDetectSelectColumnsWithTemplateMarkup(t *testing.T) {
	cols, err := DetectSelectColumns("SELECT id, email FROM users WHERE id = {{ .Id }} AND role = {{ sql .Role }}")
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if len(cols) != 2 || cols[0].Source != "id" || cols[1].Source != "email" {
		t.Fatalf("got %+v", cols)
	}
}

func TestDetectSelectColumnsFieldFunc(t *testing.T) {
	cols, err := DetectSelectColumns(
		"SELECT u.user_id, field(u.user_name, 'string', 'UserName') as x, field(count(o.order_id), 'int64', 'TotalOrders', true) " +
			"FROM users u LEFT JOIN orders o ON u.user_id = o.user_id GROUP BY u.user_id",
	)
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if len(cols) != 3 {
		t.Fatalf("expected 3 columns, got %+v", cols)
	}

	if cols[0].Name != "UUserId" || cols[0].Source != "u.user_id" {
		t.Fatalf("expected the table qualifier preserved in both Name and Source, got %+v", cols[0])
	}

	if cols[1].Name != "UserName" || cols[1].Type != "string" {
		t.Fatalf("expected field()'s explicit name/type to win, got %+v", cols[1])
	}
	// field() overrides the alias here (as x) - matches lib/querypredict's
	// own handleExpr ordering (field()'s result is set first, then As is
	// applied - but field()'s Name is what a caller actually wants).

	if cols[2].Name != "TotalOrders" || cols[2].Type != "int64" || !cols[2].Optional {
		t.Fatalf("expected field()'s type/name/optional to be captured, got %+v", cols[2])
	}
}

func TestDetectSelectColumnsRejectsNonSelect(t *testing.T) {
	if _, err := DetectSelectColumns("INSERT INTO users (email) VALUES ({{ sql .Email }})"); err == nil {
		t.Fatalf("expected an error for a non-SELECT statement")
	}
}

// TestDetectSelectColumnsJoinSameColumnName is the case that mattered most:
// a join where two tables share a bare column name (id). Naively dropping
// the table qualifier (as lib/querypredict has always done) makes both
// resolve to the same Go field name - a duplicate struct field that fails
// to compile - and the same ambiguous "id" text if ever reprojected. Both
// must come out distinct and separately qualified.
func TestDetectSelectColumnsJoinSameColumnName(t *testing.T) {
	cols, err := DetectSelectColumns(
		"SELECT u.id, o.id, u.email, o.total FROM users u JOIN orders o ON u.id = o.user_id",
	)
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	want := []struct{ name, source string }{
		{"UId", "u.id"},
		{"OId", "o.id"},
		{"UEmail", "u.email"},
		{"OTotal", "o.total"},
	}
	if len(cols) != len(want) {
		t.Fatalf("got %d columns, want %d: %+v", len(cols), len(want), cols)
	}
	for i, w := range want {
		if cols[i].Name != w.name || cols[i].Source != w.source {
			t.Fatalf("col %d: got %+v, want name=%q source=%q", i, cols[i], w.name, w.source)
		}
	}
}

// TestDetectSelectColumnsAliasedJoinColumn covers the query author
// disambiguating themselves with AS - the alias should win as the Go field
// name, and the projected SQL should keep both the source expression and
// the alias ("o.id AS order_id"), not collapse to the bare alias text
// (which isn't a real column of either table on its own).
func TestDetectSelectColumnsAliasedJoinColumn(t *testing.T) {
	cols, err := DetectSelectColumns(
		"SELECT u.id AS user_id, o.id AS order_id FROM users u JOIN orders o ON u.id = o.user_id",
	)
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if len(cols) != 2 {
		t.Fatalf("got %+v", cols)
	}
	if cols[0].Name != "UserId" || cols[0].Source != "u.id AS user_id" {
		t.Fatalf("col 0: got %+v", cols[0])
	}
	if cols[1].Name != "OrderId" || cols[1].Source != "o.id AS order_id" {
		t.Fatalf("col 1: got %+v", cols[1])
	}
}

// TestDetectSelectColumnsRejectsUnresolvableCollision is the backstop for
// whatever qualifier-preservation doesn't catch on its own - two distinct
// expressions that still land on the same derived Go name should fail
// loudly instead of silently generating a struct with a duplicate field.
func TestDetectSelectColumnsRejectsUnresolvableCollision(t *testing.T) {
	_, err := DetectSelectColumns("SELECT id AS dupe, email AS dupe FROM users")
	if err == nil {
		t.Fatalf("expected an error for two columns resolving to the same name")
	}
}

// TestDetectSelectColumnsStar: a wildcard select can't be resolved into
// discrete columns without inspecting the live schema, which this package
// deliberately never does - it returns no columns and no error, standing in
// for "we cannot determine this at all", not a failure.
func TestDetectSelectColumnsStar(t *testing.T) {
	cols, err := DetectSelectColumns("SELECT * FROM users")
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if cols != nil {
		t.Fatalf("expected no columns for SELECT *, got %+v", cols)
	}
}

func TestDetectSelectColumnsQualifiedStar(t *testing.T) {
	cols, err := DetectSelectColumns("SELECT u.* FROM users u")
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if cols != nil {
		t.Fatalf("expected no columns for u.*, got %+v", cols)
	}
}

// TestDetectSelectColumnsMixedStarAndNamed: even one named column alongside
// "*" gives up entirely, rather than returning a column list that's
// silently missing whatever "*" would have expanded to.
func TestDetectSelectColumnsMixedStarAndNamed(t *testing.T) {
	cols, err := DetectSelectColumns("SELECT id, * FROM users")
	if err != nil {
		t.Fatalf("DetectSelectColumns: %v", err)
	}
	if cols != nil {
		t.Fatalf("expected no columns when \"*\" is present at all, got %+v", cols)
	}
}
