package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for deleteUserByIdVsqlParams
type DeleteUserByIdVsqlParams struct {
	// Row to delete.
	Id int64 `json:"id" yaml:"id"`
}

func (x *DeleteUserByIdVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetDeleteUserByIdVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Row to delete.",
		},
	}
}
func CastDeleteUserByIdVsqlParamsFromCli(c emigo.CliCastable) DeleteUserByIdVsqlParams {
	data := DeleteUserByIdVsqlParams{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	return data
}

// DeleteUserByIdVsqlFilters is the allow-list DTO for the deleteUserById vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for deleteUserByIdVsqlFilters
type DeleteUserByIdVsqlFilters struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default - lets the caller log/confirm which address was removed without a separate SELECT first.
	Email string `json:"email" yaml:"email"`
}

func (x *DeleteUserByIdVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetDeleteUserByIdVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - lets the caller log/confirm which address was removed without a separate SELECT first.",
		},
	}
}
func CastDeleteUserByIdVsqlFiltersFromCli(c emigo.CliCastable) DeleteUserByIdVsqlFilters {
	data := DeleteUserByIdVsqlFilters{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	return data
}

// DeleteUserByIdVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// deleteUserById vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only DeleteUserByIdVsqlFilters's own
// top-level entries.
var DeleteUserByIdVsqlFilterFields = []string{
	"id",
	"email",
}

// DeleteUserByIdVsqlFilterFieldAllowed reports whether name is one of DeleteUserByIdVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func DeleteUserByIdVsqlFilterFieldAllowed(name string) bool {
	for _, f := range DeleteUserByIdVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// DeleteUserByIdVsqlRow is the response/row DTO for the deleteUserById vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for deleteUserByIdVsqlRow
type DeleteUserByIdVsqlRow struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default - lets the caller log/confirm which address was removed without a separate SELECT first.
	Email string `json:"email" yaml:"email"`
}

func (x *DeleteUserByIdVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetDeleteUserByIdVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - lets the caller log/confirm which address was removed without a separate SELECT first.",
		},
	}
}
func CastDeleteUserByIdVsqlRowFromCli(c emigo.CliCastable) DeleteUserByIdVsqlRow {
	data := DeleteUserByIdVsqlRow{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	return data
}

// DeleteUserByIdVsqlColumns is the column picker for the deleteUserById vsql query. Toggle Selected on
// each entry (or start from NewDeleteUserByIdVsqlColumns()), then pass it to PrepareDeleteUserByIdVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type DeleteUserByIdVsqlColumns struct {
	// Selected by default.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Selected by default - lets the caller log/confirm which address was removed without a separate SELECT first.
	Email emigo.ColumnState `json:"email" yaml:"email"`
}

// NewDeleteUserByIdVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the deleteUserById vsql query.
func NewDeleteUserByIdVsqlColumns() DeleteUserByIdVsqlColumns {
	return DeleteUserByIdVsqlColumns{
		Id:    emigo.ColumnState{Selected: true},
		Email: emigo.ColumnState{Selected: true},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c DeleteUserByIdVsqlColumns) Cols() string {
	parts := []string{}
	if c.Id.Selected {
		parts = append(parts, "id")
	}
	if c.Email.Selected {
		parts = append(parts, "email")
	}
	return strings.Join(parts, ", ")
}

// DeleteUserByIdVsqlColumnsQuery lets a caller build DeleteUserByIdVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing DeleteUserByIdVsqlColumns by hand - see Columns().
type DeleteUserByIdVsqlColumnsQuery struct {
	values url.Values
	mapped map[string]interface{}
	Id     bool `json:"id"`
	Email  bool `json:"email"`
}

// DeleteUserByIdVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// DeleteUserByIdVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func DeleteUserByIdVsqlColumnsQueryFromString(rawQuery string) DeleteUserByIdVsqlColumnsQuery {
	q := DeleteUserByIdVsqlColumnsQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           &q,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	q.values = values
	q.mapped = mapped
	return q
}

// DeleteUserByIdVsqlColumnsQueryFromHttp is DeleteUserByIdVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func DeleteUserByIdVsqlColumnsQueryFromHttp(r *http.Request) DeleteUserByIdVsqlColumnsQuery {
	return DeleteUserByIdVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a DeleteUserByIdVsqlColumns, seeded from
// NewDeleteUserByIdVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q DeleteUserByIdVsqlColumnsQuery) Columns() DeleteUserByIdVsqlColumns {
	cols := NewDeleteUserByIdVsqlColumns()
	if q.values.Has("id") {
		cols.Id.Selected = q.Id
	}
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	return cols
}

// DeleteUserByIdVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from DeleteUserByIdVsqlParams) alongside {{ .Columns.Cols }}.
type DeleteUserByIdVsqlData struct {
	DeleteUserByIdVsqlParams
	Columns DeleteUserByIdVsqlColumns `json:"columns" yaml:"columns"`
}

// DeleteUserByIdVsqlName is the name of the vsql query, useful for logging or routing.
const DeleteUserByIdVsqlName = "deleteUserById"

// DeleteUserByIdVsqlQuery is the raw SQL string for the deleteUserById vsql query.
const DeleteUserByIdVsqlQuery = `DELETE FROM users WHERE id = {{ .Id }} RETURNING {{ .Columns.Cols }};
`

// PrepareDeleteUserByIdVsql returns the query string and params for the deleteUserById vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewDeleteUserByIdVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into DeleteUserByIdVsqlRow.
func PrepareDeleteUserByIdVsql(params DeleteUserByIdVsqlParams, columns DeleteUserByIdVsqlColumns) (query string, args interface{}) {
	return DeleteUserByIdVsqlQuery, DeleteUserByIdVsqlData{DeleteUserByIdVsqlParams: params, Columns: columns}
}
