package entitysql

import (
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/torabian/emi/lib/core"
)

// The sql is produced in phases so that it can always run top to bottom, whatever the
// order or the relations of the entities are: tables first, then columns, then indexes
// and finally the foreign keys (which need the referenced tables to exist). Every
// statement is idempotent, so the whole file can be executed any number of times:
//
//   - CREATE TABLE IF NOT EXISTS / CREATE INDEX IF NOT EXISTS
//   - ALTER TABLE ... ADD COLUMN IF NOT EXISTS, so a new field is picked up on the next
//     run, while existing columns are never dropped or altered (that stays a hand
//     written migration)
//   - foreign keys are added inside a DO block which first looks into pg_constraint
//
// The shape follows what the golang entity does (see lib/golang/go-entity-gorm.go), so
// the tables are the ones gorm would query:
//
//   - every entity table has id (bigserial primary key) and unique_id
//   - object / object?  are flattened into the same table, as <object>_<field> columns
//     (object? makes the inner columns nullable)
//   - array / array?    are a child table (<entity table singular>_<field>, plural by
//     gorm naming) with its own id, unique_id and linker_id referencing the parent row
//     with ON DELETE CASCADE. Their own fields are processed recursively
//   - one               is a <field>_id bigint column, indexed, with a foreign key
//   - one?              is the same column and index, without the foreign key constraint
//   - collection        is a many2many join table named <entity>_<field>
//   - enum              is text with a CHECK constraint on the allowed keys
//   - map, slice, any, complex are stored as jsonb
type plan struct {
	byClass map[string]*core.Module3Entity
	opts    Options

	tables  []string
	columns []string
	indexes []string
	fks     []string
	seen    map[string]bool

	// layouts is what the sql above stores, table by table, see Layout.
	layouts map[string]*Layout
	roots   []*Layout

	// defs is the same tables as data, in the order they were met: what a dialect which can
	// not be written as a list of "add this if it is missing" statements (sqlite) is rendered
	// from. Recorded for every dialect, the postgres statements above are built as they go.
	defs      map[string]*tableDef
	defOrder  []string
	indexDefs []indexDef
}

// tableDef is a table, and everything which has to be said about it when it is created.
type tableDef struct {
	name    string
	columns []*colDef
	// noID marks a join table, which has no id of its own: its primary key is pk.
	noID bool
	pk   []string
}

// colDef is a column, with the foreign key and the check which belong to it. Postgres adds those
// as constraints of their own, afterwards; sqlite can only have them in the column's own definition.
type colDef struct {
	col      column
	base     bool // part of every table from the start (unique_id), never added afterwards
	refTable string
	onDelete string
	check    string
}

type indexDef struct{ table, column string }

func (d *tableDef) column(name string) *colDef {
	for _, c := range d.columns {
		if c.col.name == name {
			return c
		}
	}
	return nil
}

func newPlan(entities []*core.Module3Entity, opts Options) *plan {
	p := &plan{
		byClass: map[string]*core.Module3Entity{},
		opts:    opts,
		seen:    map[string]bool{},
		layouts: map[string]*Layout{},
		defs:    map[string]*tableDef{},
	}
	for _, e := range entities {
		if e != nil {
			p.byClass[e.GetClassName()] = e
		}
	}
	return p
}

// Options tune how the entities become tables.
type Options struct {
	// Dialect is the database the sql is for. Empty is postgres.
	Dialect Dialect

	// ChildTablesFromParent names the table of an array field <parent table>_<field>,
	// instead of the gorm derived name (plural snake case of the struct name). The gorm name
	// is what an application using gorm structs queries; entities which have no gorm struct
	// behind them (a table made on the fly) are better off with a name which is short and
	// follows from the table it belongs to.
	ChildTablesFromParent bool
}

// EntityToPostgres renders the idempotent postgres DDL of a single entity. Relations
// to entities which are not part of the call are resolved with the default table name.
func EntityToPostgres(entity *core.Module3Entity) (string, error) {
	return EntitiesToPostgres([]*core.Module3Entity{entity})
}

// EntitiesToPostgres renders the idempotent postgres DDL of the entities together, so
// relations between them resolve to the right (possibly custom named) tables.
func EntitiesToPostgres(entities []*core.Module3Entity) (string, error) {
	return EntitiesToSQL(entities, Options{})
}

// EntitiesToSQL renders the DDL of the entities for the database in opts.Dialect (postgres when
// it is not set), as a script.
func EntitiesToSQL(entities []*core.Module3Entity, opts Options) (string, error) {
	r, err := Generate(entities, opts)
	if err != nil {
		return "", err
	}
	return r.Script, nil
}

// PostgresStatements is the same ddl as EntitiesToPostgres, as the list of statements in the
// order they have to run, so a caller can execute them one by one (a driver does not take
// several statements in one call). The layouts are what those tables hold, one per entity.
// Kept for the callers which only ever speak postgres; see Generate for any other dialect.
func PostgresStatements(entities []*core.Module3Entity, opts Options) ([]string, []*Layout, error) {
	opts.Dialect = Postgres
	r, err := Generate(entities, opts)
	if err != nil {
		return nil, nil, err
	}
	sql := make([]string, len(r.Statements))
	for i, s := range r.Statements {
		sql[i] = s.SQL
	}
	return sql, r.Layouts, nil
}

func buildPlan(entities []*core.Module3Entity, opts Options) (*plan, error) {
	p := newPlan(entities, opts)
	for _, e := range entities {
		if e == nil {
			continue
		}
		if e.Name == "" {
			return nil, fmt.Errorf("entity needs a name")
		}
		if err := p.addEntity(e); err != nil {
			return nil, fmt.Errorf("entity %s: %w", e.Name, err)
		}
	}
	return p, nil
}

// TableName is the postgres table name of the entity: entity.Table when set, otherwise
// the gorm default, which is the plural snake case of the struct name (ProductEntity ->
// product_entities).
func TableName(entity *core.Module3Entity) string {
	if entity.Table != "" {
		return entity.Table
	}
	return structTable(entity.GetClassName())
}

func structTable(structName string) string {
	return inflection.Plural(core.ToSnakeCase(structName))
}

// tableOfTarget finds the table of a relation target such as UserEntity.
func (p *plan) tableOfTarget(target string) string {
	if e, ok := p.byClass[target]; ok {
		return TableName(e)
	}
	return structTable(target)
}

type scope struct {
	table      string   // table the columns are added to
	structName string   // golang struct name; child tables are named after it
	entityName string   // used by many2many join table names
	colPrefix  string   // set while flattening objects
	nullable   bool     // inside object?, all columns are nullable
	root       bool     // root of the entity, where id / uniqueId are reserved
	path       []string // json keys from the row of the table down to the field, see Layout
}

func (p *plan) addEntity(e *core.Module3Entity) error {
	table := TableName(e)
	p.addBaseTable(table)
	p.roots = append(p.roots, p.layouts[table])
	return p.addFields(scope{
		table:      table,
		structName: e.GetClassName(),
		entityName: e.Name,
		root:       true,
	}, e.Fields)
}

func (p *plan) def(table string) *tableDef {
	d := p.defs[table]
	if d == nil {
		d = &tableDef{name: table}
		p.defs[table] = d
		p.defOrder = append(p.defOrder, table)
	}
	return d
}

func (p *plan) addBaseTable(table string) {
	p.def(table)
	if p.layouts[table] == nil {
		p.layouts[table] = &Layout{Table: table}
	}
	p.add(&p.tables, fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  \"id\" bigserial PRIMARY KEY\n);", quote(table)))
	p.addColumn(table, column{name: "unique_id", sqlType: "varchar(100)", unique: true})
}

func (p *plan) addFields(s scope, fields []*core.EmiField) error {
	for _, f := range fields {
		if f == nil || f.Name == "" {
			continue
		}
		if s.root && (f.Name == "id" || f.Name == "uniqueId") {
			continue
		}
		if err := p.addField(s, f); err != nil {
			return fmt.Errorf("field %s: %w", f.Name, err)
		}
	}
	return nil
}

func (p *plan) addField(s scope, f *core.EmiField) error {
	t := string(f.Type)
	nullable := strings.HasSuffix(t, "?")
	base := core.FieldType(strings.TrimSuffix(t, "?"))
	colName := s.colPrefix + core.ToSnakeCase(f.Name)
	key, visible := fieldKey(f)
	path := append(append([]string{}, s.path...), key)

	switch base {

	case core.FieldTypeObject:
		inner := s
		inner.root = false
		inner.colPrefix = colName + "_"
		inner.structName += core.ToUpper(f.Name)
		inner.nullable = s.nullable || nullable
		inner.path = path
		return p.addFields(inner, f.Fields)

	case core.FieldTypeArray, core.FieldTypeList:
		childStruct := s.structName + core.ToUpper(f.Name)
		child := structTable(childStruct)
		if p.opts.ChildTablesFromParent {
			child = shorten(s.table + "_" + core.ToSnakeCase(f.Name))
		}
		p.addBaseTable(child)
		p.addColumn(child, column{name: "linker_id", sqlType: "bigint"})
		p.addIndex(child, "linker_id")
		p.addForeignKey(child, "linker_id", s.table, "CASCADE")
		if visible {
			parent := p.layouts[s.table]
			parent.Children = append(parent.Children, LayoutChild{Path: path, Layout: p.layouts[child]})
		}
		return p.addFields(scope{
			table:      child,
			structName: childStruct,
			entityName: s.entityName,
			root:       true,
		}, f.Fields)

	case core.FieldTypeOne, core.FieldTypeClass:
		col := colName + "_id"
		p.addColumn(s.table, column{name: col, sqlType: "bigint"})
		p.addIndex(s.table, col)
		// Like the golang side, only the non nullable relation is a real constraint.
		if !nullable {
			p.addForeignKey(s.table, col, p.tableOfTarget(f.Target), "")
		}
		return nil

	case core.FieldTypeCollection:
		if f.Target == "" {
			return fmt.Errorf("collection needs a target")
		}
		p.addJoinTable(s, f)
		return nil
	}

	col, err := scalarColumn(f, base, nullable || s.nullable)
	if err != nil {
		return err
	}
	col.name = colName
	p.addColumn(s.table, col)
	if visible {
		layout := p.layouts[s.table]
		layout.Columns = append(layout.Columns, LayoutColumn{Path: path, Column: colName, SqlType: p.vendorType(col.sqlType), Kind: col.kind})
	}
	if base == core.FieldTypeEnum {
		p.addEnumCheck(s.table, colName, f.OfType)
	}
	return nil
}

// addJoinTable is the many2many table gorm uses for a collection field. The columns are
// named after the two structs (product_entity_id, tag_entity_id); a self reference gets
// a reference_ prefix on the target side so the two columns differ.
func (p *plan) addJoinTable(s scope, f *core.EmiField) {
	join := s.entityName + "_" + f.Name
	owner := core.ToSnakeCase(s.structName) + "_id"
	target := core.ToSnakeCase(f.Target) + "_id"
	if owner == target {
		target = "reference_" + target
	}

	p.add(&p.tables, fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (\n  %s bigint NOT NULL,\n  %s bigint NOT NULL,\n  PRIMARY KEY (%s, %s)\n);",
		quote(join), quote(owner), quote(target), quote(owner), quote(target)))
	jd := p.def(join)
	jd.noID, jd.pk = true, []string{owner, target}
	jd.columns = append(jd.columns,
		&colDef{col: column{name: owner, sqlType: "bigint", notNull: true}},
		&colDef{col: column{name: target, sqlType: "bigint", notNull: true}})
	p.addIndex(join, target)
	p.addForeignKey(join, owner, s.table, "CASCADE")
	p.addForeignKey(join, target, p.tableOfTarget(f.Target), "CASCADE")
}

type column struct {
	name       string
	sqlType    string // the postgres type, whichever dialect is rendered: see sqliteType
	kind       Kind
	notNull    bool
	unique     bool
	defaultSql string
}

func (c column) definition() string {
	s := quote(c.name) + " " + c.sqlType
	if c.notNull {
		s += " NOT NULL"
	}
	if c.defaultSql != "" {
		s += " DEFAULT " + c.defaultSql
	}
	if c.unique {
		s += " UNIQUE"
	}
	return s
}

// scalarColumn maps the field types which are a single column of their own. A non
// nullable column gets a zero value default, so it can be added to a table which
// already has rows.
// complexStorage says how a complex is stored in a column of its own. The listed ones are written
// down on purpose - a translated text, a free json document and a multi-currency price are
// documents kept whole in a jsonb column - and a complex which is not listed here is a jsonb column too.
// This is the one place to say that a complex is something else: a time is a real timestamp, a date
// a real date.
var complexStorage = map[string]struct {
	sqlType string
	kind    Kind
}{
	"PlainTime": {"timestamptz", KindTime},
	"XDateTime": {"timestamptz", KindTime},
	"XDate":     {"date", KindDate},
	"TString":   {"jsonb", KindJSON},
	"MJson":     {"jsonb", KindJSON},
	"TMoney":    {"jsonb", KindJSON},
}

func scalarColumn(f *core.EmiField, base core.FieldType, nullable bool) (column, error) {
	col := column{notNull: !nullable}

	switch base {
	case core.FieldTypeString:
		col.sqlType, col.kind, col.defaultSql = "text", KindString, "''"
	case core.FieldTypeEnum:
		// The column is checked against the enum's keys, so its zero value is the first of them:
		// an empty string would break the check of every row which is there when the column is added.
		col.sqlType, col.kind, col.defaultSql = "text", KindString, "''"
		for _, v := range f.OfType {
			if v != nil {
				col.defaultSql = sqlString(v.Key)
				break
			}
		}
	case core.FieldTypeBool:
		col.sqlType, col.kind, col.defaultSql = "boolean", KindBool, "false"
	case core.FieldTypeInt, core.FieldTypeInt32:
		col.sqlType, col.kind, col.defaultSql = "integer", KindInt, "0"
	case core.FieldTypeInt64:
		col.sqlType, col.kind, col.defaultSql = "bigint", KindInt, "0"
	case core.FieldTypeFloat32:
		col.sqlType, col.kind, col.defaultSql = "real", KindFloat, "0"
	case core.FieldTypeFloat64:
		col.sqlType, col.kind, col.defaultSql = "double precision", KindFloat, "0"
	case core.FieldTypeMap:
		col.sqlType, col.kind, col.defaultSql = "jsonb", KindJSON, "'{}'"
	case core.FieldTypeSlice:
		col.sqlType, col.kind, col.defaultSql = "jsonb", KindJSON, "'[]'"
	case core.FieldTypeAny:
		col.sqlType, col.kind, col.notNull = "jsonb", KindJSON, false
	case core.FieldTypeComplex:
		// What a complex is stored as is looked up in complexStorage; one which is not listed there is a
		// json document, which every complex can be.
		col.sqlType, col.kind, col.notNull = "jsonb", KindJSON, false
		if storage, ok := complexStorage[f.Complex]; ok {
			col.sqlType, col.kind = storage.sqlType, storage.kind
		}
	default:
		return col, fmt.Errorf("unsupported type %q", f.Type)
	}

	if !col.notNull {
		col.defaultSql = ""
	}

	if f.Default != nil {
		lit, ok := literal(f.Default)
		if !ok {
			return col, fmt.Errorf("default value %v is not a scalar", f.Default)
		}
		if col.sqlType == "jsonb" {
			lit = "'" + strings.ReplaceAll(fmt.Sprint(f.Default), "'", "''") + "'"
		}
		col.defaultSql = lit
	}

	return col, nil
}

func literal(v any) (string, bool) {
	switch x := v.(type) {
	case string:
		return "'" + strings.ReplaceAll(x, "'", "''") + "'", true
	case bool, int, int32, int64, float32, float64:
		return fmt.Sprint(x), true
	}
	return "", false
}

func (p *plan) add(list *[]string, stmt string) {
	if p.seen[stmt] {
		return
	}
	p.seen[stmt] = true
	*list = append(*list, stmt)
}

func (p *plan) addColumn(table string, c column) {
	p.add(&p.columns, fmt.Sprintf("ALTER TABLE %s ADD COLUMN IF NOT EXISTS %s;", quote(table), c.definition()))
	if d := p.def(table); d.column(c.name) == nil {
		d.columns = append(d.columns, &colDef{col: c, base: c.name == "unique_id"})
	}
}

func (p *plan) addIndex(table, col string) {
	p.indexDefs = append(p.indexDefs, indexDef{table, col})
	name := shorten("idx_" + table + "_" + col)
	p.add(&p.indexes, fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s);", quote(name), quote(table), quote(col)))
}

// addForeignKey adds the constraint only while the referenced table exists: a relation
// target can live in another module, whose tables are created by another run, and the
// next run of this file then adds the constraint.
func (p *plan) addForeignKey(table, col, refTable, onDelete string) {
	if c := p.def(table).column(col); c != nil {
		c.refTable, c.onDelete = refTable, onDelete
	}
	name := shorten("fk_" + table + "_" + col)
	alter := fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (\"id\")", quote(col), quote(refTable))
	if onDelete != "" {
		alter += " ON DELETE " + onDelete
	}
	p.addConstraint(table, name, alter, refTable)
}

func (p *plan) addEnumCheck(table, col string, values []*core.EmiEnumInline) {
	var keys []string
	for _, v := range values {
		if v != nil {
			keys = append(keys, sqlString(v.Key))
		}
	}
	if len(keys) == 0 {
		return
	}
	if c := p.def(table).column(col); c != nil {
		c.check = fmt.Sprintf("%s IN (%s)", quote(col), strings.Join(keys, ", "))
	}
	p.addConstraint(table, shorten("ck_"+table+"_"+col),
		fmt.Sprintf("CHECK (%s IN (%s))", quote(col), strings.Join(keys, ", ")), "")
}

// addConstraint adds the constraint only when pg_constraint doesn't know it yet, since
// postgres has no ADD CONSTRAINT IF NOT EXISTS. requires, when set, is a table which has
// to exist first.
func (p *plan) addConstraint(table, name, definition, requires string) {
	cond := fmt.Sprintf("NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = %s AND conrelid = %s::regclass)",
		sqlString(name), sqlString(quote(table)))
	if requires != "" {
		cond += fmt.Sprintf(" AND to_regclass(%s) IS NOT NULL", sqlString(quote(requires)))
	}
	p.add(&p.fks, fmt.Sprintf("DO $$ BEGIN\n  IF %s THEN\n    ALTER TABLE %s ADD CONSTRAINT %s %s;\n  END IF;\nEND $$;",
		cond, quote(table), quote(name), definition))
}

func (p *plan) render() string {
	var b strings.Builder
	for _, phase := range []struct {
		title string
		stmts []string
	}{
		{"Tables", p.tables},
		{"Columns", p.columns},
		{"Indexes", p.indexes},
		{"Constraints", p.fks},
	} {
		if len(phase.stmts) == 0 {
			continue
		}
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		fmt.Fprintf(&b, "-- %s\n", phase.title)
		for _, s := range phase.stmts {
			b.WriteString(s)
			b.WriteString("\n")
		}
	}
	return b.String()
}

// shorten keeps generated constraint / index names inside the 63 bytes postgres allows,
// while staying deterministic (the name is how a rerun recognises it already exists).
func shorten(name string) string {
	if len(name) <= 63 {
		return name
	}
	h := fnv.New32a()
	h.Write([]byte(name))
	return fmt.Sprintf("%s_%08x", name[:54], h.Sum32())
}

func quote(ident string) string {
	return `"` + strings.ReplaceAll(ident, `"`, `""`) + `"`
}

func sqlString(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}
