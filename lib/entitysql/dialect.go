package entitysql

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// Dialect is the database vendor the sql is made for.
type Dialect string

const (
	// Postgres is the default: every statement is idempotent on its own (IF NOT EXISTS, and a
	// check of pg_constraint for the constraints), so a script can be run any number of times.
	Postgres Dialect = "postgres"
	// SQLite can not say "add this column if it is missing", add a constraint to a table which
	// exists, or add a unique column, so its tables are made with every column, foreign key and
	// check in them, and the columns which come later are statements which say which column
	// they add (see Statement), for a runner to skip when the table has it already.
	SQLite Dialect = "sqlite"
)

// The --tags which pick the dialect, for the cli, and for the `tags:` of a target of `emi compile`.
const (
	TagPostgres core.CTag = "postgres"
	TagSQLite   core.CTag = "sqlite"
)

// CompilerTags lists the tags this package understands, for `emi tags` to display.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: TagPostgres, Description: "Generate sql for postgres (the default): idempotent statements which add what is missing"},
	{Tag: TagSQLite, Description: "Generate sql for sqlite: the tables with all their columns, and the later columns as separate statements a runner has to check first"},
}

// ParseDialect reads a dialect by its name; empty is postgres.
func ParseDialect(name string) (Dialect, error) {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "", "postgres", "postgresql", "pg":
		return Postgres, nil
	case "sqlite", "sqlite3":
		return SQLite, nil
	}
	return "", fmt.Errorf("the database %q is not supported: entity sql is made for postgres and sqlite", name)
}

// DialectFromTags is the dialect the tags (the comma separated --tags value) ask for: sqlite or
// postgres, postgres when neither is there. Asking for both is an error.
func DialectFromTags(tags string) (Dialect, error) {
	var found []Dialect
	for _, tag := range strings.Split(tags, ",") {
		switch core.CTag(strings.TrimSpace(tag)) {
		case TagPostgres:
			found = append(found, Postgres)
		case TagSQLite:
			found = append(found, SQLite)
		}
	}
	switch len(found) {
	case 0:
		return Postgres, nil
	case 1:
		return found[0], nil
	}
	return "", fmt.Errorf("the tags ask for more than one database (%s and %s)", TagPostgres, TagSQLite)
}

// Statement is one statement to run.
type Statement struct {
	SQL string
	// Table and Column are set on a statement which adds a column to a table which may be there
	// already (sqlite). The statement fails when the table has the column, and can not be told to
	// skip it, so a runner looks at the table first (pragma table_info) and runs it only if the
	// column is missing. A statement of any other kind runs as it is.
	Table  string
	Column string
}

// AddsColumn tells whether the statement needs the check described on Statement.
func (s Statement) AddsColumn() bool { return s.Column != "" }

// Result is what Generate makes out of entities.
type Result struct {
	Dialect Dialect
	// Statements run in this order, one by one.
	Statements []Statement
	// Layouts say what each entity's table holds, one per entity, see Layout.
	Layouts []*Layout
	// Script is the statements as one file, with a comment on each part.
	Script string
}

// Generate makes the tables of the entities for opts.Dialect: the statements to run, the layout of
// each table, and the same as a script to read or run by hand.
func Generate(entities []*core.Module3Entity, opts Options) (*Result, error) {
	if opts.Dialect == "" {
		opts.Dialect = Postgres
	}
	if _, err := ParseDialect(string(opts.Dialect)); err != nil {
		return nil, err
	}
	p, err := buildPlan(entities, opts)
	if err != nil {
		return nil, err
	}
	r := &Result{Dialect: opts.Dialect, Layouts: p.roots}
	switch opts.Dialect {
	case SQLite:
		r.Statements = p.sqliteStatements()
		r.Script = p.sqliteScript(r.Statements)
	default:
		for _, sql := range p.postgresStatements() {
			r.Statements = append(r.Statements, Statement{SQL: sql})
		}
		r.Script = p.render()
	}
	return r, nil
}

func (p *plan) postgresStatements() []string {
	var all []string
	all = append(all, p.tables...)
	all = append(all, p.columns...)
	all = append(all, p.indexes...)
	all = append(all, p.fks...)
	return all
}

// vendorType is the type of a column in the dialect: a column is described with its postgres
// type, whatever the dialect is.
func (p *plan) vendorType(pgType string) string {
	if p.opts.Dialect == SQLite {
		return sqliteType(pgType)
	}
	return pgType
}
