// Covers predict:true against a join across two tables sharing a column
// name (id) - the case that turned out to be broken before lib/sqlpredict
// started preserving table qualifiers: u.id/o.id used to both collapse to
// the Go field name "Id", producing a struct with a duplicate field that
// failed to compile. See lib/sqlpredict/predict_test.go for the
// package-level coverage of the same fix; this is the same check run
// through real codegen and a real database.
package vsqlsqlite_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	external "github.com/torabian/emi/examples/vsql-sqlite/sdkgen"
)

func TestPredictAcrossJoinWithAmbiguousColumnName(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer db.Close()

	for _, stmt := range []string{
		`CREATE TABLE users (id INTEGER PRIMARY KEY, email TEXT)`,
		`CREATE TABLE orders (id INTEGER PRIMARY KEY, user_id INTEGER, amount_cents INTEGER)`,
		`INSERT INTO users (id, email) VALUES (1, 'alice@example.com')`,
		`INSERT INTO orders (id, user_id, amount_cents) VALUES (100, 1, 2500)`,
	} {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatalf("setup %q: %v", stmt, err)
		}
	}

	params := external.UserOrderTotalsVsqlParams{}
	columns := external.NewUserOrderTotalsVsqlColumns()

	// The predicted Row/Columns types themselves compiling at all (UId and
	// OrderId as two distinct fields, not one duplicated "Id") is most of
	// what this test guards - it would have failed to build otherwise.
	query, args := external.PrepareUserOrderTotalsVsql(params, columns)

	sqlText := render(t, query, args)
	if got, want := sqlText, "SELECT u.id, o.id AS order_id, u.email, o.amount_cents\nFROM users u\nJOIN orders o ON o.user_id = u.id;\n"; got != want {
		t.Fatalf("got sql:\n%s\nwant:\n%s", got, want)
	}

	var row external.UserOrderTotalsVsqlRow
	if err := db.QueryRow(sqlText).Scan(&row.UId, &row.OrderId, &row.UEmail, &row.OAmountCents); err != nil {
		t.Fatalf("scan: %v\nsql:\n%s", err, sqlText)
	}

	if row.UId != "1" || row.OrderId != "100" || row.UEmail != "alice@example.com" || row.OAmountCents != "2500" {
		t.Fatalf("unexpected row: %+v", row)
	}
}
