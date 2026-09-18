package core

// EmiVsql defines a hand-written SQL query alongside the typed parameter
// shape it expects. Emi generates a parameters DTO and a Prepare helper so
// callers pass values through a struct instead of positional args.
type EmiVsql struct {
	// Description of what the query does. Used in generated code documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of what the query does. Used in generated code documentation."`

	// Name of the vsql query in camel case. Used as the file and identifier base.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the vsql query in camel case. Used as the file and identifier base."`

	// Params is the typed parameter list for the query, same shape as dto fields.
	// The quick, inline way to declare input - most queries only need this.
	Params []*EmiField `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Typed parameter list for the query, same shape as dto fields."`

	// In is an alternative, action-shaped way to declare input, alongside
	// (not instead of) Params - reach for it when Params' plain field list
	// isn't enough: In.Dto points at an existing top-level or template DTO
	// instead of repeating its fields, and In.Headers carries typed headers.
	// Both routes describe the same thing (the query's input), so use
	// whichever fits a given query; declare both and they're combined. This
	// is also the shape EmiAction.In uses, on purpose - a vsql exposed
	// publicly the way an action is would reuse In as its request body
	// rather than needing a second, differently-shaped declaration.
	//
	// At preprocessing time, In.Dto/In.Fields are resolved into plain fields
	// and merged into Params (a name already in Params wins over the same
	// name from In); In itself is left in place afterwards since Headers
	// (and the Dto reference, for anything that wants to know it) aren't
	// something Params can express.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Alternative, action-shaped way to declare input alongside Params - supports referencing an existing Dto by name and typed headers. Resolved into Params during preprocessing; useful when this query is meant to be exposed the way an EmiAction is."`

	// Captures pulls fields from existing DTOs into Params at preprocessing
	// time. By the time the generator runs, the resolved fields have been
	// flattened into Params and Captures is cleared. See EmiCapture.
	Captures []*EmiCapture `yaml:"captures,omitempty" json:"captures,omitempty" jsonschema:"description=Capture fields from existing DTOs into Params. Resolved during preprocessing."`

	// Query is the raw SQL (or templated SQL) string the query will execute,
	// compiled into the generated Go file as a string constant. Mutually
	// exclusive with QueryName - set exactly one.
	Query string `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Raw SQL string the query will execute. Mutually exclusive with queryName."`

	// QueryName, instead of Query, names a file to read the query text from
	// at runtime via an fs.FS the caller supplies - typically a real .sql
	// file embedded with go:embed, so the SQL lives in its own file (no YAML
	// string-escaping, real editor/syntax-highlighting support) rather than
	// inline in this module's yaml. When set, the generated Prepare<Name>Vsql
	// takes an fs.FS as its first argument and reads QueryName from it
	// instead of returning a compiled-in constant. Mutually exclusive with
	// Query - set exactly one.
	QueryName string `yaml:"queryName,omitempty" json:"queryName,omitempty" jsonschema:"description=File name to read this query's SQL from at runtime via a caller-supplied fs.FS, instead of compiling Query in as a constant. Mutually exclusive with query."`

	// Columns declares the columns a SELECT-shaped query can optionally
	// include. Emi generates a picker struct (one bool-ish field per column,
	// defaulted from Selected) plus a Cols() helper that renders the
	// currently-selected columns as a SQL projection, so the query template
	// doesn't have to repeat the column list. See EmiColumn.
	Columns []*EmiColumn `yaml:"columns,omitempty" json:"columns,omitempty" jsonschema:"description=Columns this query can optionally select. Emi generates a picker struct plus a Cols() projection helper from this list."`

	// Predict, when true and Columns is empty, populates Columns
	// automatically during preprocessing by parsing Query/QueryName as a
	// plain SELECT statement and detecting its projected columns (name,
	// type, optionality) - the column-detection mechanism Query Predict
	// (lib/querypredict) has used since before vsql existed, ported into
	// lib/sqlpredict so vsql can use it too and Query Predict can
	// eventually be retired without losing the capability. A no-op if
	// Columns is already non-empty - hand-written columns always win.
	//
	// Only a genuine SELECT can be predicted (see
	// sqlpredict.DetectSelectColumns) - the underlying parser has no notion
	// of RETURNING, so an INSERT/UPDATE/DELETE ... RETURNING vsql (the
	// shape most examples in this repo use) needs Columns declared by hand
	// instead. A SELECT list containing "*" leaves Columns empty rather than
	// guessing - what "*" expands to depends on the live schema, which Emi
	// never inspects - so such a vsql simply gets no Columns/no row DTO, the
	// same as if Predict had never been set.
	Predict bool `yaml:"predict,omitempty" json:"predict,omitempty" jsonschema:"description=When true and columns is empty, detects columns automatically by parsing query/queryName as a plain SELECT during preprocessing, instead of declaring columns by hand. Only supports a genuine SELECT statement; a select list containing '*' leaves columns empty rather than guessing."`

	// Filters declares which fields a caller is allowed to reference inside a
	// JSON-Logic filter expression (https://jsonlogic.com) run against this
	// query's rows - e.g. {"==": [{"var": "role"}, "Admin"]}. Deliberately a
	// plain []*EmiField, not []*EmiColumn: a filter allow-list is a value
	// shape (name + type), not a selection picker, and it's compiled through
	// the exact same common struct/class generator as Params or a dto's own
	// fields (GoCommonStructGenerator in Go, JsCommonObjectGenerator in JS) -
	// same reason In reuses EmiActionBody rather than inventing a fourth
	// shape. Emi never executes or transpiles the filter itself (that stays
	// the caller's job, e.g. via jsonlogic2sql) - Filters only exists to
	// generate an allow-list of legal field names/types, so a filter can be
	// validated (Go) or type-checked as far as TypeScript's string-literal
	// unions allow (JS/TS) before it ever reaches SQL.
	//
	// When left empty and Columns is non-empty, Filters defaults to
	// Columns' own fields during preprocessing - the row surface a SELECT
	// already exposes is the natural default set of things worth filtering
	// on. A query with no Columns (nothing to default from) needs Filters
	// declared explicitly if it wants filtering at all.
	Filters []*EmiField `yaml:"filters,omitempty" json:"filters,omitempty" jsonschema:"description=Fields a caller may reference inside a JSON-Logic filter expression against this query. Defaults to Columns' fields when left empty and Columns is set; compiled through the same struct/class generator as Params, not the Columns picker."`
}

func (x EmiVsql) GetClassName() string {
	return ToUpper(x.Name) + "Vsql"
}

func (x EmiVsql) GetParamsClassName() string {
	return ToUpper(x.Name) + "VsqlParams"
}

func (x EmiVsql) GetColumnsClassName() string {
	return ToUpper(x.Name) + "VsqlColumns"
}

// GetRowClassName is the response/row DTO generated from Columns - one field
// per declared column, typed exactly as it would be as a dto field.
func (x EmiVsql) GetRowClassName() string {
	return ToUpper(x.Name) + "VsqlRow"
}

// GetFiltersClassName is the allow-list DTO generated from Filters - one
// field per filterable field, typed exactly as it would be as a dto field.
func (x EmiVsql) GetFiltersClassName() string {
	return ToUpper(x.Name) + "VsqlFilters"
}
