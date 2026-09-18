package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for updateUserPartialVsqlParams
type UpdateUserPartialVsqlParams struct {
	// Row to update. Always applied - there's no picker toggle for the WHERE clause.
	Id int64 `json:"id" yaml:"id"`
	// New email, or unset to leave it unchanged.
	Email emigo.Nullable[string] `json:"email" yaml:"email"`
	// New first name, or unset to leave it unchanged.
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	// New last name, or unset to leave it unchanged.
	LastName emigo.Nullable[string] `json:"lastName" yaml:"lastName"`
	// New balance, or unset to leave it unchanged.
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
}

func (x *UpdateUserPartialVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUpdateUserPartialVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Row to update. Always applied - there's no picker toggle for the WHERE clause.",
		},
		{
			Name:        prefix + "email",
			Type:        "string?",
			Description: "New email, or unset to leave it unchanged.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string?",
			Description: "New first name, or unset to leave it unchanged.",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string?",
			Description: "New last name, or unset to leave it unchanged.",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "New balance, or unset to leave it unchanged.",
		},
	}
}
func CastUpdateUserPartialVsqlParamsFromCli(c emigo.CliCastable) UpdateUserPartialVsqlParams {
	data := UpdateUserPartialVsqlParams{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
	}
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	return data
}

// UpdateUserPartialVsqlFilters is the allow-list DTO for the updateUserPartial vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for updateUserPartialVsqlFilters
type UpdateUserPartialVsqlFilters struct {
	// Set true exactly when params.Email actually carries a value the caller wants applied - the query trusts this flag rather than re-checking Email.IsSet itself.
	Email emigo.Nullable[string] `json:"email" yaml:"email"`
	// Same convention as email, for firstName.
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	// Same convention as email, for lastName.
	LastName emigo.Nullable[string] `json:"lastName" yaml:"lastName"`
	// Same convention as email, for balanceCents.
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
}

func (x *UpdateUserPartialVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUpdateUserPartialVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "email",
			Type:        "string?",
			Description: "Set true exactly when params.Email actually carries a value the caller wants applied - the query trusts this flag rather than re-checking Email.IsSet itself.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string?",
			Description: "Same convention as email, for firstName.",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string?",
			Description: "Same convention as email, for lastName.",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Same convention as email, for balanceCents.",
		},
	}
}
func CastUpdateUserPartialVsqlFiltersFromCli(c emigo.CliCastable) UpdateUserPartialVsqlFilters {
	data := UpdateUserPartialVsqlFilters{}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
	}
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	return data
}

// UpdateUserPartialVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// updateUserPartial vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only UpdateUserPartialVsqlFilters's own
// top-level entries.
var UpdateUserPartialVsqlFilterFields = []string{
	"email",
	"firstName",
	"lastName",
	"balanceCents",
}

// UpdateUserPartialVsqlFilterFieldAllowed reports whether name is one of UpdateUserPartialVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func UpdateUserPartialVsqlFilterFieldAllowed(name string) bool {
	for _, f := range UpdateUserPartialVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// UpdateUserPartialVsqlRow is the response/row DTO for the updateUserPartial vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for updateUserPartialVsqlRow
type UpdateUserPartialVsqlRow struct {
	// Set true exactly when params.Email actually carries a value the caller wants applied - the query trusts this flag rather than re-checking Email.IsSet itself.
	Email emigo.Nullable[string] `json:"email" yaml:"email"`
	// Same convention as email, for firstName.
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	// Same convention as email, for lastName.
	LastName emigo.Nullable[string] `json:"lastName" yaml:"lastName"`
	// Same convention as email, for balanceCents.
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
}

func (x *UpdateUserPartialVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUpdateUserPartialVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "email",
			Type:        "string?",
			Description: "Set true exactly when params.Email actually carries a value the caller wants applied - the query trusts this flag rather than re-checking Email.IsSet itself.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string?",
			Description: "Same convention as email, for firstName.",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string?",
			Description: "Same convention as email, for lastName.",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Same convention as email, for balanceCents.",
		},
	}
}
func CastUpdateUserPartialVsqlRowFromCli(c emigo.CliCastable) UpdateUserPartialVsqlRow {
	data := UpdateUserPartialVsqlRow{}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("last-name") {
		emigo.ParseNullable(c.String("last-name"), &data.LastName)
	}
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	return data
}

// UpdateUserPartialVsqlColumns is the column picker for the updateUserPartial vsql query. Toggle Selected on
// each entry (or start from NewUpdateUserPartialVsqlColumns()), then pass it to PrepareUpdateUserPartialVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type UpdateUserPartialVsqlColumns struct {
	// Set true exactly when params.Email actually carries a value the caller wants applied - the query trusts this flag rather than re-checking Email.IsSet itself.
	Email emigo.ColumnState `json:"email" yaml:"email"`
	// Same convention as email, for firstName.
	FirstName emigo.ColumnState `json:"firstName" yaml:"firstName"`
	// Same convention as email, for lastName.
	LastName emigo.ColumnState `json:"lastName" yaml:"lastName"`
	// Same convention as email, for balanceCents.
	BalanceCents emigo.ColumnState `json:"balanceCents" yaml:"balanceCents"`
}

// NewUpdateUserPartialVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the updateUserPartial vsql query.
func NewUpdateUserPartialVsqlColumns() UpdateUserPartialVsqlColumns {
	return UpdateUserPartialVsqlColumns{
		Email:        emigo.ColumnState{Selected: false},
		FirstName:    emigo.ColumnState{Selected: false},
		LastName:     emigo.ColumnState{Selected: false},
		BalanceCents: emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c UpdateUserPartialVsqlColumns) Cols() string {
	parts := []string{}
	if c.Email.Selected {
		parts = append(parts, "email")
	}
	if c.FirstName.Selected {
		parts = append(parts, "first_name")
	}
	if c.LastName.Selected {
		parts = append(parts, "last_name")
	}
	if c.BalanceCents.Selected {
		parts = append(parts, "balance_cents")
	}
	return strings.Join(parts, ", ")
}

// UpdateUserPartialVsqlColumnsQuery lets a caller build UpdateUserPartialVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing UpdateUserPartialVsqlColumns by hand - see Columns().
type UpdateUserPartialVsqlColumnsQuery struct {
	values       url.Values
	mapped       map[string]interface{}
	Email        bool `json:"email"`
	FirstName    bool `json:"firstName"`
	LastName     bool `json:"lastName"`
	BalanceCents bool `json:"balanceCents"`
}

// UpdateUserPartialVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// UpdateUserPartialVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func UpdateUserPartialVsqlColumnsQueryFromString(rawQuery string) UpdateUserPartialVsqlColumnsQuery {
	q := UpdateUserPartialVsqlColumnsQuery{}
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

// UpdateUserPartialVsqlColumnsQueryFromHttp is UpdateUserPartialVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func UpdateUserPartialVsqlColumnsQueryFromHttp(r *http.Request) UpdateUserPartialVsqlColumnsQuery {
	return UpdateUserPartialVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a UpdateUserPartialVsqlColumns, seeded from
// NewUpdateUserPartialVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q UpdateUserPartialVsqlColumnsQuery) Columns() UpdateUserPartialVsqlColumns {
	cols := NewUpdateUserPartialVsqlColumns()
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	if q.values.Has("firstName") {
		cols.FirstName.Selected = q.FirstName
	}
	if q.values.Has("lastName") {
		cols.LastName.Selected = q.LastName
	}
	if q.values.Has("balanceCents") {
		cols.BalanceCents.Selected = q.BalanceCents
	}
	return cols
}

// UpdateUserPartialVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from UpdateUserPartialVsqlParams) alongside {{ .Columns.Cols }}.
type UpdateUserPartialVsqlData struct {
	UpdateUserPartialVsqlParams
	Columns UpdateUserPartialVsqlColumns `json:"columns" yaml:"columns"`
}

// UpdateUserPartialVsqlName is the name of the vsql query, useful for logging or routing.
const UpdateUserPartialVsqlName = "updateUserPartial"

// UpdateUserPartialVsqlQuery is the raw SQL string for the updateUserPartial vsql query.
const UpdateUserPartialVsqlQuery = `UPDATE users SET
  email = COALESCE({{ if .Columns.Email.Selected }}{{ sql .Email }}{{ else }}NULL{{ end }}, email),
  first_name = COALESCE({{ if .Columns.FirstName.Selected }}{{ sql .FirstName }}{{ else }}NULL{{ end }}, first_name),
  last_name = COALESCE({{ if .Columns.LastName.Selected }}{{ sql .LastName }}{{ else }}NULL{{ end }}, last_name),
  balance_cents = COALESCE({{ if .Columns.BalanceCents.Selected }}{{ sql .BalanceCents }}{{ else }}NULL{{ end }}, balance_cents)
WHERE id = {{ .Id }}
RETURNING id, email, first_name, last_name, balance_cents;
`

// PrepareUpdateUserPartialVsql returns the query string and params for the updateUserPartial vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewUpdateUserPartialVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into UpdateUserPartialVsqlRow.
func PrepareUpdateUserPartialVsql(params UpdateUserPartialVsqlParams, columns UpdateUserPartialVsqlColumns) (query string, args interface{}) {
	return UpdateUserPartialVsqlQuery, UpdateUserPartialVsqlData{UpdateUserPartialVsqlParams: params, Columns: columns}
}
