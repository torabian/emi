package core

// EmiColumn describes one column a vsql SELECT-shaped query can optionally
// include. It embeds EmiField, so a column has the exact same name/type/
// description vocabulary (including nested Fields, Complex, Target, enums,
// ...) as any other emi field. That lets Emi run a vsql's declared columns
// straight through GoCommonStructGenerator - the same function used for
// vsql.Params and every EmiDto - to produce a typed row/response struct for
// the query, instead of the caller hand-rolling one.
//
// On top of the embedded field, a column adds only what a plain field
// doesn't already cover: the SQL fragment to emit when selected, and its
// default inclusion state. The embedded field's own Name is the column's
// identifier - it becomes both the picker struct's field name (PascalCased)
// and the row struct's field name.
type EmiColumn struct {
	// Inlined (not nested under a "field" key) - yaml.v2 needs the explicit
	// ",inline" tag for an embedded struct's own keys (name, type, ...) to be
	// read straight off the column's YAML mapping; encoding/json already
	// flattens an anonymous field like this without needing to be told.
	EmiField `yaml:",inline" json:",inline"`

	// Column is the SQL fragment emitted when this column is selected -
	// normally just the real column name, but it can be any expression (e.g.
	// "u.email" or "count(*) as total"), or even more than one physical
	// column (e.g. "amount_cents, currency" for a Money-shaped complex
	// field - see EmiColumn's own doc comment). Defaults to the field's
	// Name, snake_cased, when left empty (so "firstName" -> "first_name"
	// without needing to repeat it); set it explicitly whenever the SQL-side
	// name differs from the field name for any other reason.
	Column string `yaml:"column,omitempty" json:"column,omitempty" jsonschema:"description=SQL fragment (column name or expression) emitted when this column is selected. Defaults to the snake_cased field name when empty; set it to override the column name used in the query."`

	// Selected is the column's default inclusion state. The generated
	// constructor seeds the picker struct from this; a caller flips
	// individual columns afterwards before preparing the query.
	Selected bool `yaml:"selected,omitempty" json:"selected,omitempty" jsonschema:"description=Whether this column is included by default, before the caller overrides anything."`
}

// GetColumn returns the SQL fragment to project for this column when
// selected: Column verbatim if set, otherwise the field Name snake_cased.
func (c EmiColumn) GetColumn() string {
	if c.Column != "" {
		return c.Column
	}
	return ToSnakeCase(c.Name)
}
