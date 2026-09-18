package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"io/fs"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for selectUsersPagedVsqlParams
type SelectUsersPagedVsqlParams struct {
	// Max rows per page.
	Limit int `json:"limit" yaml:"limit"`
	// Rows to skip before the first one on this page.
	Offset int `json:"offset" yaml:"offset"`
}

func (x *SelectUsersPagedVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUsersPagedVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "limit",
			Type:        "int",
			Description: "Max rows per page.",
		},
		{
			Name:        prefix + "offset",
			Type:        "int",
			Description: "Rows to skip before the first one on this page.",
		},
	}
}
func CastSelectUsersPagedVsqlParamsFromCli(c emigo.CliCastable) SelectUsersPagedVsqlParams {
	data := SelectUsersPagedVsqlParams{}
	if c.IsSet("limit") {
		data.Limit = int(c.Int64("limit"))
	}
	if c.IsSet("offset") {
		data.Offset = int(c.Int64("offset"))
	}
	return data
}

// SelectUsersPagedVsqlFilters is the allow-list DTO for the selectUsersPaged vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for selectUsersPagedVsqlFilters
type SelectUsersPagedVsqlFilters struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default.
	Email string `json:"email" yaml:"email"`
	// Off by default. Column name overridden to user_role.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
}

func (x *SelectUsersPagedVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUsersPagedVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "Off by default. Column name overridden to user_role.",
		},
	}
}
func CastSelectUsersPagedVsqlFiltersFromCli(c emigo.CliCastable) SelectUsersPagedVsqlFilters {
	data := SelectUsersPagedVsqlFilters{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	return data
}

// SelectUsersPagedVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// selectUsersPaged vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only SelectUsersPagedVsqlFilters's own
// top-level entries.
var SelectUsersPagedVsqlFilterFields = []string{
	"id",
	"email",
	"role",
}

// SelectUsersPagedVsqlFilterFieldAllowed reports whether name is one of SelectUsersPagedVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func SelectUsersPagedVsqlFilterFieldAllowed(name string) bool {
	for _, f := range SelectUsersPagedVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// SelectUsersPagedVsqlRow is the response/row DTO for the selectUsersPaged vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for selectUsersPagedVsqlRow
type SelectUsersPagedVsqlRow struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default.
	Email string `json:"email" yaml:"email"`
	// Off by default. Column name overridden to user_role.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
}

func (x *SelectUsersPagedVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUsersPagedVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "Off by default. Column name overridden to user_role.",
		},
	}
}
func CastSelectUsersPagedVsqlRowFromCli(c emigo.CliCastable) SelectUsersPagedVsqlRow {
	data := SelectUsersPagedVsqlRow{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	return data
}

// SelectUsersPagedVsqlColumns is the column picker for the selectUsersPaged vsql query. Toggle Selected on
// each entry (or start from NewSelectUsersPagedVsqlColumns()), then pass it to PrepareSelectUsersPagedVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type SelectUsersPagedVsqlColumns struct {
	// Selected by default.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Selected by default.
	Email emigo.ColumnState `json:"email" yaml:"email"`
	// Off by default. Column name overridden to user_role.
	Role emigo.ColumnState `json:"role" yaml:"role"`
}

// NewSelectUsersPagedVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the selectUsersPaged vsql query.
func NewSelectUsersPagedVsqlColumns() SelectUsersPagedVsqlColumns {
	return SelectUsersPagedVsqlColumns{
		Id:    emigo.ColumnState{Selected: true},
		Email: emigo.ColumnState{Selected: true},
		Role:  emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c SelectUsersPagedVsqlColumns) Cols() string {
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
	return strings.Join(parts, ", ")
}

// SelectUsersPagedVsqlColumnsQuery lets a caller build SelectUsersPagedVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing SelectUsersPagedVsqlColumns by hand - see Columns().
type SelectUsersPagedVsqlColumnsQuery struct {
	values url.Values
	mapped map[string]interface{}
	Id     bool `json:"id"`
	Email  bool `json:"email"`
	Role   bool `json:"role"`
}

// SelectUsersPagedVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// SelectUsersPagedVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func SelectUsersPagedVsqlColumnsQueryFromString(rawQuery string) SelectUsersPagedVsqlColumnsQuery {
	q := SelectUsersPagedVsqlColumnsQuery{}
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

// SelectUsersPagedVsqlColumnsQueryFromHttp is SelectUsersPagedVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func SelectUsersPagedVsqlColumnsQueryFromHttp(r *http.Request) SelectUsersPagedVsqlColumnsQuery {
	return SelectUsersPagedVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a SelectUsersPagedVsqlColumns, seeded from
// NewSelectUsersPagedVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q SelectUsersPagedVsqlColumnsQuery) Columns() SelectUsersPagedVsqlColumns {
	cols := NewSelectUsersPagedVsqlColumns()
	if q.values.Has("id") {
		cols.Id.Selected = q.Id
	}
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	if q.values.Has("role") {
		cols.Role.Selected = q.Role
	}
	return cols
}

// SelectUsersPagedVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from SelectUsersPagedVsqlParams) alongside {{ .Columns.Cols }}.
type SelectUsersPagedVsqlData struct {
	SelectUsersPagedVsqlParams
	Columns SelectUsersPagedVsqlColumns `json:"columns" yaml:"columns"`
}

// SelectUsersPagedVsqlName is the name of the vsql query, useful for logging or routing.
const SelectUsersPagedVsqlName = "selectUsersPaged"

// SelectUsersPagedVsqlQueryFile is the file name PrepareSelectUsersPagedVsql reads this vsql's SQL from, via a
// caller-supplied fs.FS, instead of a compiled-in constant.
const SelectUsersPagedVsqlQueryFile = "select_users_paged.sql"

// PrepareSelectUsersPagedVsql returns the query string and params for the selectUsersPaged vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewSelectUsersPagedVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into SelectUsersPagedVsqlRow.
// fsys is read for select_users_paged.sql (see SelectUsersPagedVsqlQueryFile) instead of a compiled-in
// constant - typically an embed.FS the caller built with go:embed over its
// own .sql files.
func PrepareSelectUsersPagedVsql(fsys fs.FS, params SelectUsersPagedVsqlParams, columns SelectUsersPagedVsqlColumns) (query string, args interface{}, err error) {
	b, err := fs.ReadFile(fsys, SelectUsersPagedVsqlQueryFile)
	if err != nil {
		return "", nil, err
	}
	return string(b), SelectUsersPagedVsqlData{SelectUsersPagedVsqlParams: params, Columns: columns}, nil
}
