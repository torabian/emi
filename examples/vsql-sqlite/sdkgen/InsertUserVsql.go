package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for insertUserVsqlParams
type InsertUserVsqlParams struct {
	// Required - every user needs one.
	Email string `json:"email" yaml:"email"`
	// Required display name part.
	FirstName string `json:"firstName" yaml:"firstName"`
	// Optional - some users only give a first name. Nullable (`?`) so the query can tell "not provided" apart from "provided as empty string" via {{ if .LastName.IsSet }}.
	LastName emigo.Nullable[string] `json:"lastName" yaml:"lastName"`
	// Required account role, Admin or Member and nothing else - "of:" is what makes that a compile-time-checked Go type instead of a bare string.
	Role string `json:"role" yaml:"role"`
	// Whether the account starts enabled.
	IsActive bool `json:"isActive" yaml:"isActive"`
}

func (x *InsertUserVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUserVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Required - every user needs one.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string",
			Description: "Required display name part.",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string?",
			Description: "Optional - some users only give a first name. Nullable (`?`) so the query can tell \"not provided\" apart from \"provided as empty string\" via {{ if .LastName.IsSet }}.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum",
			Description: "Required account role, Admin or Member and nothing else - \"of:\" is what makes that a compile-time-checked Go type instead of a bare string.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool",
			Description: "Whether the account starts enabled.",
		},
	}
}
func CastInsertUserVsqlParamsFromCli(c emigo.CliCastable) InsertUserVsqlParams {
	data := InsertUserVsqlParams{}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
	}
	if c.IsSet("role") {
		data.Role = c.String("role")
	}
	if c.IsSet("is-active") {
		data.IsActive = bool(c.Bool("is-active"))
	}
	return data
}

// InsertUserVsqlFilters is the allow-list DTO for the insertUser vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for insertUserVsqlFilters
type InsertUserVsqlFilters struct {
	// Primary key of the new row. Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Echoes the email that was just inserted. Selected by default.
	Email string `json:"email" yaml:"email"`
	// Selected by default. The real column is user_role, not role - `column:` overrides the default snake_cased-Name projection for exactly this kind of field/column name mismatch.
	Role string `json:"role" yaml:"role"`
	// Off by default - most callers don't need it echoed back. Because it defaults unselected, Emi's preprocessing forces this field to a nullable type in the generated row DTO (emigo.Nullable[bool]) even though it's declared as a plain bool here - see nullifyUnselectedColumns.
	IsActive emigo.Nullable[bool] `json:"isActive" yaml:"isActive"`
}

func (x *InsertUserVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUserVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Primary key of the new row. Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Echoes the email that was just inserted. Selected by default.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum",
			Description: "Selected by default. The real column is user_role, not role - `column:` overrides the default snake_cased-Name projection for exactly this kind of field/column name mismatch.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool?",
			Description: "Off by default - most callers don't need it echoed back. Because it defaults unselected, Emi's preprocessing forces this field to a nullable type in the generated row DTO (emigo.Nullable[bool]) even though it's declared as a plain bool here - see nullifyUnselectedColumns.",
		},
	}
}
func CastInsertUserVsqlFiltersFromCli(c emigo.CliCastable) InsertUserVsqlFilters {
	data := InsertUserVsqlFilters{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("role") {
		data.Role = c.String("role")
	}
	if c.IsSet("is-active") {
		emigo.ParseNullable(c.String("is-active"), &data.IsActive)
	}
	return data
}

// InsertUserVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// insertUser vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only InsertUserVsqlFilters's own
// top-level entries.
var InsertUserVsqlFilterFields = []string{
	"id",
	"email",
	"role",
	"isActive",
}

// InsertUserVsqlFilterFieldAllowed reports whether name is one of InsertUserVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func InsertUserVsqlFilterFieldAllowed(name string) bool {
	for _, f := range InsertUserVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// InsertUserVsqlRow is the response/row DTO for the insertUser vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for insertUserVsqlRow
type InsertUserVsqlRow struct {
	// Primary key of the new row. Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Echoes the email that was just inserted. Selected by default.
	Email string `json:"email" yaml:"email"`
	// Selected by default. The real column is user_role, not role - `column:` overrides the default snake_cased-Name projection for exactly this kind of field/column name mismatch.
	Role string `json:"role" yaml:"role"`
	// Off by default - most callers don't need it echoed back. Because it defaults unselected, Emi's preprocessing forces this field to a nullable type in the generated row DTO (emigo.Nullable[bool]) even though it's declared as a plain bool here - see nullifyUnselectedColumns.
	IsActive emigo.Nullable[bool] `json:"isActive" yaml:"isActive"`
}

func (x *InsertUserVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUserVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Primary key of the new row. Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Echoes the email that was just inserted. Selected by default.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum",
			Description: "Selected by default. The real column is user_role, not role - `column:` overrides the default snake_cased-Name projection for exactly this kind of field/column name mismatch.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool?",
			Description: "Off by default - most callers don't need it echoed back. Because it defaults unselected, Emi's preprocessing forces this field to a nullable type in the generated row DTO (emigo.Nullable[bool]) even though it's declared as a plain bool here - see nullifyUnselectedColumns.",
		},
	}
}
func CastInsertUserVsqlRowFromCli(c emigo.CliCastable) InsertUserVsqlRow {
	data := InsertUserVsqlRow{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("role") {
		data.Role = c.String("role")
	}
	if c.IsSet("is-active") {
		emigo.ParseNullable(c.String("is-active"), &data.IsActive)
	}
	return data
}

// InsertUserVsqlColumns is the column picker for the insertUser vsql query. Toggle Selected on
// each entry (or start from NewInsertUserVsqlColumns()), then pass it to PrepareInsertUserVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type InsertUserVsqlColumns struct {
	// Primary key of the new row. Selected by default.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Echoes the email that was just inserted. Selected by default.
	Email emigo.ColumnState `json:"email" yaml:"email"`
	// Selected by default. The real column is user_role, not role - `column:` overrides the default snake_cased-Name projection for exactly this kind of field/column name mismatch.
	Role emigo.ColumnState `json:"role" yaml:"role"`
	// Off by default - most callers don't need it echoed back. Because it defaults unselected, Emi's preprocessing forces this field to a nullable type in the generated row DTO (emigo.Nullable[bool]) even though it's declared as a plain bool here - see nullifyUnselectedColumns.
	IsActive emigo.ColumnState `json:"isActive" yaml:"isActive"`
}

// NewInsertUserVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the insertUser vsql query.
func NewInsertUserVsqlColumns() InsertUserVsqlColumns {
	return InsertUserVsqlColumns{
		Id:       emigo.ColumnState{Selected: true},
		Email:    emigo.ColumnState{Selected: true},
		Role:     emigo.ColumnState{Selected: true},
		IsActive: emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c InsertUserVsqlColumns) Cols() string {
	parts := []string{}
	if c.Id.Selected {
		parts = append(parts, "id")
	}
	if c.Email.Selected {
		parts = append(parts, "email")
	}
	if c.Role.Selected {
		parts = append(parts, "user_role")
	}
	if c.IsActive.Selected {
		parts = append(parts, "is_active")
	}
	return strings.Join(parts, ", ")
}

// InsertUserVsqlColumnsQuery lets a caller build InsertUserVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing InsertUserVsqlColumns by hand - see Columns().
type InsertUserVsqlColumnsQuery struct {
	values   url.Values
	mapped   map[string]interface{}
	Id       bool `json:"id"`
	Email    bool `json:"email"`
	Role     bool `json:"role"`
	IsActive bool `json:"isActive"`
}

// InsertUserVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// InsertUserVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func InsertUserVsqlColumnsQueryFromString(rawQuery string) InsertUserVsqlColumnsQuery {
	q := InsertUserVsqlColumnsQuery{}
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

// InsertUserVsqlColumnsQueryFromHttp is InsertUserVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func InsertUserVsqlColumnsQueryFromHttp(r *http.Request) InsertUserVsqlColumnsQuery {
	return InsertUserVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a InsertUserVsqlColumns, seeded from
// NewInsertUserVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q InsertUserVsqlColumnsQuery) Columns() InsertUserVsqlColumns {
	cols := NewInsertUserVsqlColumns()
	if q.values.Has("id") {
		cols.Id.Selected = q.Id
	}
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	if q.values.Has("role") {
		cols.Role.Selected = q.Role
	}
	if q.values.Has("isActive") {
		cols.IsActive.Selected = q.IsActive
	}
	return cols
}

// InsertUserVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from InsertUserVsqlParams) alongside {{ .Columns.Cols }}.
type InsertUserVsqlData struct {
	InsertUserVsqlParams
	Columns InsertUserVsqlColumns `json:"columns" yaml:"columns"`
}

// InsertUserVsqlName is the name of the vsql query, useful for logging or routing.
const InsertUserVsqlName = "insertUser"

// InsertUserVsqlQuery is the raw SQL string for the insertUser vsql query.
const InsertUserVsqlQuery = `INSERT INTO users (email, first_name, last_name, user_role, is_active)
VALUES (
  {{ sql .Email }},
  {{ sql .FirstName }},
  {{ if .LastName.IsSet }}{{ sql .LastName }}{{ else }}NULL{{ end }},
  {{ sql .Role }},
  {{ sql .IsActive }}
)
RETURNING {{ .Columns.Cols }};
`

// PrepareInsertUserVsql returns the query string and params for the insertUser vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewInsertUserVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into InsertUserVsqlRow.
func PrepareInsertUserVsql(params InsertUserVsqlParams, columns InsertUserVsqlColumns) (query string, args interface{}) {
	return InsertUserVsqlQuery, InsertUserVsqlData{InsertUserVsqlParams: params, Columns: columns}
}
