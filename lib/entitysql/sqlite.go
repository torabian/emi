package entitysql

import (
	"fmt"
	"strings"
)

// sqliteType is the type a column has in sqlite, from the postgres one. What matters is the
// affinity the name gives it, and what a driver does with a declared type: boolean, date and
// datetime are read back as a bool and as times, a json document is plain text (a type which
// is not text would be turned into a number when it looks like one).
func sqliteType(pgType string) string {
	switch pgType {
	case "bigint", "integer":
		return "integer"
	case "real", "double precision":
		return "real"
	case "jsonb":
		return "text"
	case "timestamptz":
		return "datetime"
	}
	return pgType // text, boolean, date, varchar(n)
}

func sqliteDefault(c column) string {
	if c.sqlType == "boolean" {
		switch c.defaultSql {
		case "true":
			return "1"
		case "false":
			return "0"
		}
	}
	return c.defaultSql
}

// sqliteColumn is the definition of a column, with its foreign key and check in it. known are
// the tables of this generation: a foreign key to any other table is left out. Postgres adds
// such a constraint later, once the table is there; a foreign key which sqlite has in the column
// can not wait, and one to a table which does not exist makes every change of the rows which
// have it fail ("no such table").
func sqliteColumn(c *colDef, known map[string]*tableDef) string {
	s := quote(c.col.name) + " " + sqliteType(c.col.sqlType)
	if c.col.notNull {
		s += " NOT NULL"
	}
	if d := sqliteDefault(c.col); d != "" {
		s += " DEFAULT " + d
	}
	if c.col.unique {
		s += " UNIQUE"
	}
	if _, ok := known[c.refTable]; ok {
		s += " REFERENCES " + quote(c.refTable) + ` ("id")`
		if c.onDelete != "" {
			s += " ON DELETE " + c.onDelete
		}
	}
	if c.check != "" {
		s += " CHECK (" + c.check + ")"
	}
	return s
}

func sqliteCreateTable(d *tableDef, known map[string]*tableDef) string {
	var lines []string
	if !d.noID {
		lines = append(lines, `"id" integer PRIMARY KEY AUTOINCREMENT`)
	}
	for _, c := range d.columns {
		lines = append(lines, sqliteColumn(c, known))
	}
	if d.noID {
		pk := make([]string, len(d.pk))
		for i, c := range d.pk {
			pk[i] = quote(c)
		}
		lines = append(lines, "PRIMARY KEY ("+strings.Join(pk, ", ")+")")
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  %s\n);", quote(d.name), strings.Join(lines, ",\n  "))
}

// sqliteStatements are the tables, each with every column it has, then the columns again as
// statements of their own (for a table which was made before the column was added), then the
// indexes, which need the columns to be there.
func (p *plan) sqliteStatements() []Statement {
	var tables, columns, indexes []Statement
	for _, name := range p.defOrder {
		d := p.defs[name]
		tables = append(tables, Statement{SQL: sqliteCreateTable(d, p.defs)})
		if d.noID {
			continue // a join table is never extended: its columns are its key
		}
		for _, c := range d.columns {
			if c.base {
				continue // sqlite can not add a unique column, and every table is made with it
			}
			columns = append(columns, Statement{
				SQL:    fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s;", quote(name), sqliteColumn(c, p.defs)),
				Table:  name,
				Column: c.col.name,
			})
		}
	}
	seen := map[indexDef]bool{}
	for _, ix := range p.indexDefs {
		if seen[ix] {
			continue
		}
		seen[ix] = true
		indexes = append(indexes, Statement{SQL: fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s);",
			quote(shorten("idx_"+ix.table+"_"+ix.column)), quote(ix.table), quote(ix.column))})
	}
	return append(append(tables, columns...), indexes...)
}

// sqliteScript is the statements as a file. The column statements are comments in it: they fail
// on a table which has the column, and sqlite has no way to skip them, so a script which has
// them live could not be run twice. The tables above them are made with every column, which is
// all a new database needs; a runner (see Statement.AddsColumn) is what extends an existing one.
func (p *plan) sqliteScript(statements []Statement) string {
	var tables, columns, indexes []string
	for _, s := range statements {
		switch {
		case s.AddsColumn():
			columns = append(columns, "-- "+s.SQL)
		case strings.HasPrefix(s.SQL, "CREATE TABLE"):
			tables = append(tables, s.SQL)
		default:
			indexes = append(indexes, s.SQL)
		}
	}
	var b strings.Builder
	section := func(title string, lines []string, note string) {
		if len(lines) == 0 {
			return
		}
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		fmt.Fprintf(&b, "-- %s\n", title)
		if note != "" {
			b.WriteString(note)
		}
		for _, l := range lines {
			b.WriteString(l + "\n")
		}
	}
	section("Tables", tables, "")
	section("Columns", columns, "-- sqlite can not add a column only if it is missing. Every table above is made with all\n"+
		"-- its columns; these are for a table which was made before the column was added, and fail\n"+
		"-- on one which has it, so run one only after looking at `PRAGMA table_info(<table>)`.\n")
	section("Indexes", indexes, "")
	return b.String()
}
