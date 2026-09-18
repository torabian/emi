// Exercises the insertUsers vsql query generated from
// batch_insert_users.emi.yml (run `make generate` first, or just
// `make test`). Two things are under test:
//
//   - the mechanics a caller actually touches: an array param (a batch of
//     users to insert), the column picker controlling the RETURNING
//     projection, and the generated InsertUsersVsqlRow DTO a real driver
//     would scan results into;
//   - the type matrix itself - one selected/unselected pair per emi type
//     declared on the query's columns, proving every one is selectable,
//     shows up in Cols() under its own name, is present on the row DTO with
//     the right Go type, and - the point of nullifyUnselectedColumns - comes
//     out nullable exactly when (and only when) it defaults to unselected.
package vsqlcolumns_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"text/template"

	"github.com/torabian/emi/emigo"
	external "github.com/torabian/emi/examples/vsql-columns/sdkgen"
)

// render stands in for whatever a real project already uses (see
// examples/vsql/vsql.Render for a fuller version). sql quotes a value as a
// SQL string literal - the only helper this particular query template calls.
func render(query string, data any) (string, error) {
	t, err := template.New("q").Funcs(template.FuncMap{
		"sql": func(v any) string {
			return "'" + strings.ReplaceAll(fmt.Sprint(v), "'", "''") + "'"
		},
	}).Parse(query)
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func batchOfOne() external.InsertUsersVsqlParams {
	return external.InsertUsersVsqlParams{
		Users: emigo.ArrayReplace([]external.InsertUsersVsqlParamsUsers{
			{Email: "ali@example.com", FirstName: "Ali", LastName: "Torabi"},
		}),
	}
}

func TestInsertUsersDefaultColumns(t *testing.T) {
	params := external.InsertUsersVsqlParams{
		Users: emigo.ArrayReplace([]external.InsertUsersVsqlParamsUsers{
			{Email: "ali@example.com", FirstName: "Ali", LastName: "Torabi"},
			{Email: "jo@example.com", FirstName: "Jo", LastName: "Doe"},
		}),
	}
	columns := external.NewInsertUsersVsqlColumns() // whatever the yaml's `selected: true` columns are

	query, args := external.PrepareInsertUsersVsql(params, columns)
	sql, err := render(query, args)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL (defaults) ---\n%s", sql)

	if !strings.Contains(sql, "'ali@example.com'") || !strings.Contains(sql, "'jo@example.com'") {
		t.Fatalf("expected both batch rows to be rendered, got:\n%s", sql)
	}
	if !strings.HasSuffix(strings.TrimSpace(sql), "RETURNING id, email, is_verified, account_status, profile;") {
		t.Fatalf("expected the default-selected columns only, got:\n%s", sql)
	}
}

func TestInsertUsersEveryColumnSelected(t *testing.T) {
	columns := external.NewInsertUsersVsqlColumns()
	columns.BalanceCents.Selected = true
	columns.FirstName.Selected = true
	columns.IsSuspended.Selected = true
	columns.Rating.Selected = true
	columns.Role.Selected = true
	columns.Preferences.Selected = true
	columns.Money.Selected = true

	query, args := external.PrepareInsertUsersVsql(batchOfOne(), columns)
	sql, err := render(query, args)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL (everything selected) ---\n%s", sql)

	want := "RETURNING id, balance_cents, email, first_name, is_verified, is_suspended, rating, account_status, role, profile, preferences, amount_cents, currency;"
	if !strings.HasSuffix(strings.TrimSpace(sql), want) {
		t.Fatalf("expected every column (money as its two physical columns) in RETURNING, got:\n%s", sql)
	}
}

func TestInsertUsersColumnsAreIndependentlyToggleable(t *testing.T) {
	// Flip only two of the off-by-default columns; everything else should
	// stay exactly at its declared default. Guards against Cols()/the picker
	// accidentally coupling columns to each other.
	columns := external.NewInsertUsersVsqlColumns()
	columns.Rating.Selected = true
	columns.Role.Selected = true

	query, args := external.PrepareInsertUsersVsql(batchOfOne(), columns)
	sql, err := render(query, args)
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	want := "RETURNING id, email, is_verified, rating, account_status, role, profile;"
	if !strings.HasSuffix(strings.TrimSpace(sql), want) {
		t.Fatalf("expected exactly the defaults plus rating+role, got:\n%s", sql)
	}
}

// TestColumnTypesAndNullability is the type-matrix test: for every column
// declared in batch_insert_users.emi.yml, assert its InsertUsersVsqlRow
// field exists with the expected Go type, and that nullability tracks
// `selected` - not the type as originally declared in the yaml (every one of
// these was written as a plain, non-`?` type there).
func TestColumnTypesAndNullability(t *testing.T) {
	row := reflect.TypeOf(external.InsertUsersVsqlRow{})

	cases := []struct {
		field        string
		selected     bool
		wantKindName string // Type.String() of the expected field type
	}{
		{"Id", true, "int64"},
		{"BalanceCents", false, "emigo.Nullable[int64]"},
		{"Email", true, "string"},
		{"FirstName", false, "emigo.Nullable[string]"},
		{"IsVerified", true, "bool"},
		{"IsSuspended", false, "emigo.Nullable[bool]"},
		{"Rating", false, "emigo.Nullable[float64]"},
		{"Status", true, "string"},                 // enum, selected -> plain string
		{"Role", false, "emigo.Nullable[string]"},   // enum, unselected -> nullable string
		{"Profile", true, "external.InsertUsersVsqlRowProfile"},
		// reflect names a generic's type argument by its full import path, not
		// the local package alias - unlike a bare struct field's Type.String().
		{"Preferences", false, "emigo.Nullable[github.com/torabian/emi/examples/vsql-columns/sdkgen.InsertUsersVsqlRowPreferences]"},
		{"Money", false, "external.Money"}, // complex: no wrapper change, by design
	}

	for _, c := range cases {
		t.Run(c.field, func(t *testing.T) {
			f, ok := row.FieldByName(c.field)
			if !ok {
				t.Fatalf("InsertUsersVsqlRow has no field %q - column dropped from codegen?", c.field)
			}
			if got := f.Type.String(); got != c.wantKindName {
				t.Fatalf("field %s: got type %s, want %s", c.field, got, c.wantKindName)
			}

			isNullableWrapper := strings.HasPrefix(f.Type.String(), "emigo.Nullable[")
			if c.selected && isNullableWrapper {
				t.Fatalf("field %s defaults to selected but is wrapped in Nullable - should be the plain type", c.field)
			}
			if !c.selected && c.field != "Money" && !isNullableWrapper {
				t.Fatalf("field %s defaults to unselected but is NOT nullable - nullifyUnselectedColumns should have forced it", c.field)
			}
		})
	}
}

// TestUnselectedFieldIsAbsentNotZero is the behavioral half of the
// nullability test: a field that was never selected must be distinguishable
// from one that was fetched and happened to be zero-valued.
func TestUnselectedFieldIsAbsentNotZero(t *testing.T) {
	var row external.InsertUsersVsqlRow // as if nothing had been scanned into it yet

	if row.BalanceCents.IsSet() {
		t.Fatalf("expected BalanceCents to start unset")
	}

	// Simulate a driver scanning a real (zero) value for a column that WAS
	// selected this time.
	row.BalanceCents = emigo.NullableOf(int64(0))
	if !row.BalanceCents.IsSet() {
		t.Fatalf("expected BalanceCents to report set once a value (even zero) is assigned")
	}
	if v, ok := row.BalanceCents.Get(); !ok || v == nil || *v != 0 {
		t.Fatalf("expected the scanned zero value to round-trip, got %v (ok=%v)", v, ok)
	}
}

func TestObjectColumnRowScan(t *testing.T) {
	columns := external.NewInsertUsersVsqlColumns()
	columns.Preferences.Selected = true

	query, args := external.PrepareInsertUsersVsql(batchOfOne(), columns)
	sql, err := render(query, args)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if !strings.Contains(sql, "preferences") {
		t.Fatalf("expected preferences to join the RETURNING projection, got:\n%s", sql)
	}

	// The row DTO's Preferences field is a real nested struct (wrapped in
	// Nullable, since it's off by default), not a string or map - type
	// safety on a jsonb column, not just an opaque blob.
	row := external.InsertUsersVsqlRow{
		Preferences: emigo.NullableOf(external.InsertUsersVsqlRowPreferences{Theme: "dark", Locale: "en-US"}),
	}
	v, ok := row.Preferences.Get()
	if !ok || v == nil || v.Theme != "dark" {
		t.Fatalf("expected the nested Preferences struct to round-trip its fields")
	}
}

// TestComplexColumnSpansMultipleSqlColumns is the Money scenario: a single
// EmiColumn (one Selected toggle, one entry in the picker) can still project
// more than one physical SQL column, because Column is a free-form fragment -
// Cols() splices it in verbatim.
func TestComplexColumnSpansMultipleSqlColumns(t *testing.T) {
	columns := external.NewInsertUsersVsqlColumns()
	columns.Money.Selected = true

	got := columns.Cols()
	if !strings.Contains(got, "amount_cents, currency") {
		t.Fatalf("expected the money column to contribute both of its physical columns, got: %q", got)
	}

	// The Go-side field is still one field of the complex type - Money
	// itself is what would map those two scanned columns back together.
	row := external.InsertUsersVsqlRow{Money: external.Money{AmountCents: 1999, Currency: "USD"}}
	if sqlLit, include := row.Money.SQLValue(); !include || sqlLit == "" {
		t.Fatalf("expected Money.SQLValue to render a non-empty literal for a set amount")
	}
	if _, include := (external.Money{}).SQLValue(); include {
		t.Fatalf("expected a zero-value Money (no currency) to opt itself out, like Nullable[T] does for primitives")
	}
}
