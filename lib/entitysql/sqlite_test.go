package entitysql

import (
	"database/sql"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/torabian/emi/lib/core"
)

func openSQLite(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared&_foreign_keys=on")
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	t.Cleanup(func() { db.Close() })
	return db
}

func hasColumn(t *testing.T, db *sql.DB, table, column string) bool {
	t.Helper()
	rows, err := db.Query(`SELECT name FROM pragma_table_info(?)`, table)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		_ = rows.Scan(&name)
		if name == column {
			return true
		}
	}
	return false
}

// run is what a runner of the statements does: everything as it is, except a statement which adds
// a column, which waits for a look at the table.
func run(t *testing.T, db *sql.DB, statements []Statement) (added int) {
	t.Helper()
	for _, s := range statements {
		if s.AddsColumn() && hasColumn(t, db, s.Table, s.Column) {
			continue
		}
		if _, err := db.Exec(s.SQL); err != nil {
			t.Fatalf("%v\n%s", err, s.SQL)
		}
		if s.AddsColumn() {
			added++
		}
	}
	return added
}

func sqliteOf(t *testing.T, yml string) *Result {
	t.Helper()
	m, err := core.StringToEmi(yml)
	if err != nil {
		t.Fatal(err)
	}
	r, err := Generate(m.Entities, Options{Dialect: SQLite})
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func TestSQLiteStatementsRunAndRunAgain(t *testing.T) {
	db := openSQLite(t)
	m, _ := core.StringToEmi(sample)
	r, err := Generate(m.Entities, Options{Dialect: SQLite})
	if err != nil {
		t.Fatal(err)
	}
	if r.Dialect != SQLite {
		t.Fatalf("dialect %q", r.Dialect)
	}

	run(t, db, r.Statements)
	// a second run finds every table and every column there, and changes nothing
	if added := run(t, db, r.Statements); added != 0 {
		t.Errorf("the second run added %d columns", added)
	}

	for _, table := range []string{"order_item_entities", "order_item_entity_lines", "order_item_entity_lines_notes", "orderItem_tags", "orderItem_related"} {
		var n int
		if err := db.QueryRow(`SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?`, table).Scan(&n); err != nil || n != 1 {
			t.Errorf("table %s was not made (err %v)", table, err)
		}
	}
}

// The checks and the foreign keys are in the tables, and sqlite holds them to it.
func TestSQLiteEnforcesTheEnumAndTheForeignKeys(t *testing.T) {
	db := openSQLite(t)
	run(t, db, sqliteOf(t, sample).Statements)

	mustExec := func(q string, args ...any) {
		t.Helper()
		if _, err := db.Exec(q, args...); err != nil {
			t.Fatalf("%v\n%s", err, q)
		}
	}
	mustExec(`INSERT INTO categories_custom (title) VALUES ('c')`)

	// the enum is a check on the column
	if _, err := db.Exec(`INSERT INTO order_item_entities (status, category_id) VALUES ('bogus', 1)`); err == nil {
		t.Error("a value outside the enum must be refused")
	}
	mustExec(`INSERT INTO order_item_entities (status, category_id) VALUES ('open', 1)`)

	// a non nullable relation is a real foreign key, the nullable one (owner) is not
	if _, err := db.Exec(`INSERT INTO order_item_entities (status, category_id) VALUES ('open', 999)`); err == nil {
		t.Error("a relation to a row which does not exist must be refused")
	}
	mustExec(`INSERT INTO order_item_entities (status, category_id, owner_id) VALUES ('open', 1, 12345)`)

	// an array's rows follow their parent, with ON DELETE CASCADE
	mustExec(`INSERT INTO order_item_entity_lines (linker_id, sku) VALUES (1, 'x')`)
	if _, err := db.Exec(`INSERT INTO order_item_entity_lines (linker_id, sku) VALUES (777, 'y')`); err == nil {
		t.Error("an element of an array without its row must be refused")
	}
	mustExec(`DELETE FROM order_item_entities WHERE id = 1`)
	var left int
	_ = db.QueryRow(`SELECT count(*) FROM order_item_entity_lines`).Scan(&left)
	if left != 0 {
		t.Errorf("%d array rows left after their parent was deleted", left)
	}
}

// A schema which grows between two runs: the table which already exists is extended, with the
// columns which need more than a type (a foreign key, a check, a default) as well, and the rows
// which were there stay.
func TestSQLiteExtendsTablesWhichExist(t *testing.T) {
	db := openSQLite(t)
	v1 := `name: x
entities:
  - name: category
    fields:
      - name: title
        type: string
  - name: product
    fields:
      - name: title
        type: string
`
	v2 := `name: x
entities:
  - name: category
    fields:
      - name: title
        type: string
  - name: product
    fields:
      - name: title
        type: string
      - name: status
        type: enum
        of:
          - k: draft
          - k: live
      - name: category
        type: one
        target: CategoryEntity
      - name: price
        type: float64
        default: 1.5
      - name: active
        type: bool
        default: true
      - name: variants
        type: array
        fields:
          - name: sku
            type: string
`
	run(t, db, sqliteOf(t, v1).Statements)
	if _, err := db.Exec(`INSERT INTO product_entities (title) VALUES ('old')`); err != nil {
		t.Fatal(err)
	}

	r2 := sqliteOf(t, v2)
	if added := run(t, db, r2.Statements); added != 4 {
		t.Errorf("added %d columns, want status, category_id, price and active", added)
	}
	if added := run(t, db, r2.Statements); added != 0 {
		t.Errorf("running again added %d columns", added)
	}

	// the row from before is there, with the defaults of what was added
	var title string
	var price float64
	var active bool // a column declared boolean is read back as one
	if err := db.QueryRow(`SELECT title, price, active FROM product_entities`).Scan(&title, &price, &active); err != nil {
		t.Fatal(err)
	}
	if title != "old" || price != 1.5 || !active {
		t.Errorf("row after the extension: %q %v %v", title, price, active)
	}

	// the added columns are enforced as the created ones are
	if _, err := db.Exec(`INSERT INTO product_entities (title, status) VALUES ('x', 'bogus')`); err == nil {
		t.Error("the check of a column which was added later must hold")
	}
	if _, err := db.Exec(`INSERT INTO product_entities (title, category_id) VALUES ('x', 424242)`); err == nil {
		t.Error("the foreign key of a column which was added later must hold")
	}
	if !hasColumn(t, db, "product_entity_variants", "sku") {
		t.Error("the table of the new array was not made")
	}
}

// The script is what the statements say, and what a database which has none of it takes: the
// tables are made with all their columns, and the later columns are comments, because sqlite
// could not run them twice.
func TestSQLiteScriptRunsTwice(t *testing.T) {
	db := openSQLite(t)
	m, _ := core.StringToEmi(sample)
	script, err := EntitiesToSQL(m.Entities, Options{Dialect: SQLite})
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if _, err := db.Exec(script); err != nil {
			t.Fatalf("run %d: %v\n%s", i+1, err, script)
		}
	}
	if !strings.Contains(script, "-- ALTER TABLE") {
		t.Errorf("the column statements must be there, as comments:\n%s", script)
	}
	for _, line := range strings.Split(script, "\n") {
		if strings.HasPrefix(line, "ALTER TABLE") {
			t.Errorf("a live ALTER in the script would fail the second run: %s", line)
		}
	}
	if strings.Contains(script, "DO $$") || strings.Contains(script, "bigserial") || strings.Contains(script, "::") {
		t.Errorf("postgres syntax in the sqlite script:\n%s", script)
	}
}

func TestSQLiteLayoutUsesSQLiteTypes(t *testing.T) {
	m, _ := core.StringToEmi(dynamicSample)
	r, err := Generate(m.Entities, Options{Dialect: SQLite, ChildTablesFromParent: true})
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{}
	for _, c := range r.Layouts[0].Columns {
		got[strings.Join(c.Path, ".")] = c.SqlType + " " + string(c.Kind)
	}
	for path, want := range map[string]string{
		"last name": "text string", "age": "integer int", "score": "real float", "agreed": "boolean bool",
		"born": "date date", "seen": "datetime time", "tags": "text json", "address.geo.lat": "real float",
	} {
		if got[path] != want {
			t.Errorf("%q: got %q, want %q", path, got[path], want)
		}
	}
}

func TestDialectFromTags(t *testing.T) {
	for tags, want := range map[string]Dialect{"": Postgres, "postgres": Postgres, "sqlite": SQLite, "typescript, sqlite": SQLite, "x,y": Postgres} {
		got, err := DialectFromTags(tags)
		if err != nil || got != want {
			t.Errorf("tags %q: got %q, %v, want %q", tags, got, err, want)
		}
	}
	if _, err := DialectFromTags("sqlite,postgres"); err == nil {
		t.Error("asking for two databases must be an error")
	}
	if _, err := Generate(nil, Options{Dialect: "oracle"}); err == nil {
		t.Error("an unknown database must be an error")
	}
}

// The default is, and stays, postgres.
func TestDefaultDialectIsPostgres(t *testing.T) {
	m, _ := core.StringToEmi(sample)
	def, err := Generate(m.Entities, Options{})
	if err != nil {
		t.Fatal(err)
	}
	pg, _ := Generate(m.Entities, Options{Dialect: Postgres})
	if def.Script != pg.Script || def.Dialect != Postgres || !strings.Contains(def.Script, "bigserial") {
		t.Error("with no dialect the sql is postgres")
	}
}
