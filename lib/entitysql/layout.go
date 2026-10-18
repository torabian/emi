package entitysql

import (
	"strings"

	"github.com/torabian/emi/lib/core"
)

// Layout is what the sql of an entity stores, in the terms of the json the entity is made
// from. It is produced together with the sql (see PostgresStatements), so the statements
// that create a table and the code that reads and writes its rows can not disagree on which
// column holds which value.
//
// A field is found under the json key of the field: its `json` tag when it has one, else its
// name. A field tagged `json:"-"` still gets its column (a table may need columns the json
// never carries, like the workspace of a row) but is not in the layout.
type Layout struct {
	// Table is the table of the rows.
	Table string
	// Columns are the scalar values, in the order the fields are declared.
	Columns []LayoutColumn
	// Children are the arrays of objects: each element is a row of a table of its own,
	// which holds the id of the row it belongs to in linker_id.
	Children []LayoutChild
}

// LayoutColumn is one column, and the json path to its value inside the row's object: an
// object field is flattened into the table, so a column can sit under several keys (for
// address { street } the path is ["address", "street"] and the column address_street).
type LayoutColumn struct {
	Path    []string
	Column  string
	SqlType string
	Kind    Kind
}

// LayoutChild is an array of objects: the json path to the array, and the table of its elements.
type LayoutChild struct {
	Path   []string
	Layout *Layout
}

// Kind tells how a column's value is carried in json.
type Kind string

const (
	KindString Kind = "string"
	KindInt    Kind = "int"
	KindFloat  Kind = "float"
	KindBool   Kind = "bool"
	// KindJSON is any json value, kept as a jsonb document: maps, lists of plain values, any.
	KindJSON Kind = "json"
	// KindTime is an RFC 3339 string in json.
	KindTime Kind = "time"
	// KindDate is a YYYY-MM-DD string in json.
	KindDate Kind = "date"
)

// fieldKey is the json key of a field, and whether the json carries it at all.
func fieldKey(f *core.EmiField) (key string, visible bool) {
	if tag := f.Tags["json"]; tag != "" {
		name := strings.Split(tag, ",")[0]
		if name == "-" {
			return "", false
		}
		if name != "" {
			return name, true
		}
	}
	return f.Name, true
}
