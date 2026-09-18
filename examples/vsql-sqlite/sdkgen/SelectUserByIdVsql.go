package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for selectUserByIdVsqlParams
type SelectUserByIdVsqlParams struct {
	// Row to fetch. Always required - there's no picker toggle for the WHERE clause, only for what comes back.
	Id int64 `json:"id" yaml:"id"`
}

func (x *SelectUserByIdVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUserByIdVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Row to fetch. Always required - there's no picker toggle for the WHERE clause, only for what comes back.",
		},
	}
}
func CastSelectUserByIdVsqlParamsFromCli(c emigo.CliCastable) SelectUserByIdVsqlParams {
	data := SelectUserByIdVsqlParams{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	return data
}

// SelectUserByIdVsqlFilters is the allow-list DTO for the selectUserById vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for selectUserByIdVsqlFilters
type SelectUserByIdVsqlFilters struct {
	// Selected by default - stays a plain int64 in the row DTO.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default - stays a plain string.
	Email string `json:"email" yaml:"email"`
	// Selected by default - stays a plain string (enums have no separate Go type). Column name overridden to user_role.
	Role string `json:"role" yaml:"role"`
	// Off by default - becomes emigo.Nullable[bool] in the row DTO, distinguishing "not fetched" from "fetched, false".
	IsActive emigo.Nullable[bool] `json:"isActive" yaml:"isActive"`
	// Off by default - becomes emigo.Nullable[int64].
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.Nullable[float64] `json:"rating" yaml:"rating"`
	// Off by default - becomes emigo.Nullable[SelectUserByIdVsqlRowAddress], a real nested struct, not two loose strings. The `column:` here isn't a plain column name at all - it's a json_object(...) SQL expression aggregating street+city into one JSON text value, which the caller then json.Unmarshal's into that struct after scanning. Demonstrates that `column:` accepts any SQL expression, not just identifiers.
	Address emigo.Nullable[SelectUserByIdVsqlFiltersAddress] `json:"address" yaml:"address"`
}

// The base class definition for address
type SelectUserByIdVsqlFiltersAddress struct {
	// Street address line.
	Street string `json:"street" yaml:"street"`
	// City name.
	City string `json:"city" yaml:"city"`
}

func (x *SelectUserByIdVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUserByIdVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default - stays a plain int64 in the row DTO.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - stays a plain string.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum",
			Description: "Selected by default - stays a plain string (enums have no separate Go type). Column name overridden to user_role.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool?",
			Description: "Off by default - becomes emigo.Nullable[bool] in the row DTO, distinguishing \"not fetched\" from \"fetched, false\".",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Off by default - becomes emigo.Nullable[int64].",
		},
		{
			Name:        prefix + "rating",
			Type:        "float64?",
			Description: "Off by default - becomes emigo.Nullable[float64].",
		},
		{
			Name:        prefix + "address",
			Type:        "object?",
			Description: "Off by default - becomes emigo.Nullable[SelectUserByIdVsqlRowAddress], a real nested struct, not two loose strings. The `column:` here isn't a plain column name at all - it's a json_object(...) SQL expression aggregating street+city into one JSON text value, which the caller then json.Unmarshal's into that struct after scanning. Demonstrates that `column:` accepts any SQL expression, not just identifiers.",
		},
	}
}
func CastSelectUserByIdVsqlFiltersFromCli(c emigo.CliCastable) SelectUserByIdVsqlFilters {
	data := SelectUserByIdVsqlFilters{}
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
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	if c.IsSet("rating") {
		emigo.ParseNullable(c.String("rating"), &data.Rating)
	}
	if c.IsSet("address") {
		emigo.ParseNullable(c.String("address"), &data.Address)
	}
	return data
}
func GetSelectUserByIdVsqlFiltersAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "street",
			Type:        "string",
			Description: "Street address line.",
		},
		{
			Name:        prefix + "city",
			Type:        "string",
			Description: "City name.",
		},
	}
}
func CastSelectUserByIdVsqlFiltersAddressFromCli(c emigo.CliCastable) SelectUserByIdVsqlFiltersAddress {
	data := SelectUserByIdVsqlFiltersAddress{}
	if c.IsSet("street") {
		data.Street = c.String("street")
	}
	if c.IsSet("city") {
		data.City = c.String("city")
	}
	return data
}

// SelectUserByIdVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// selectUserById vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only SelectUserByIdVsqlFilters's own
// top-level entries.
var SelectUserByIdVsqlFilterFields = []string{
	"id",
	"email",
	"role",
	"isActive",
	"balanceCents",
	"rating",
	"address",
}

// SelectUserByIdVsqlFilterFieldAllowed reports whether name is one of SelectUserByIdVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func SelectUserByIdVsqlFilterFieldAllowed(name string) bool {
	for _, f := range SelectUserByIdVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// SelectUserByIdVsqlRow is the response/row DTO for the selectUserById vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for selectUserByIdVsqlRow
type SelectUserByIdVsqlRow struct {
	// Selected by default - stays a plain int64 in the row DTO.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default - stays a plain string.
	Email string `json:"email" yaml:"email"`
	// Selected by default - stays a plain string (enums have no separate Go type). Column name overridden to user_role.
	Role string `json:"role" yaml:"role"`
	// Off by default - becomes emigo.Nullable[bool] in the row DTO, distinguishing "not fetched" from "fetched, false".
	IsActive emigo.Nullable[bool] `json:"isActive" yaml:"isActive"`
	// Off by default - becomes emigo.Nullable[int64].
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.Nullable[float64] `json:"rating" yaml:"rating"`
	// Off by default - becomes emigo.Nullable[SelectUserByIdVsqlRowAddress], a real nested struct, not two loose strings. The `column:` here isn't a plain column name at all - it's a json_object(...) SQL expression aggregating street+city into one JSON text value, which the caller then json.Unmarshal's into that struct after scanning. Demonstrates that `column:` accepts any SQL expression, not just identifiers.
	Address emigo.Nullable[SelectUserByIdVsqlRowAddress] `json:"address" yaml:"address"`
}

// The base class definition for address
type SelectUserByIdVsqlRowAddress struct {
	// Street address line.
	Street string `json:"street" yaml:"street"`
	// City name.
	City string `json:"city" yaml:"city"`
}

func (x *SelectUserByIdVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSelectUserByIdVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Selected by default - stays a plain int64 in the row DTO.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - stays a plain string.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum",
			Description: "Selected by default - stays a plain string (enums have no separate Go type). Column name overridden to user_role.",
		},
		{
			Name:        prefix + "is-active",
			Type:        "bool?",
			Description: "Off by default - becomes emigo.Nullable[bool] in the row DTO, distinguishing \"not fetched\" from \"fetched, false\".",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Off by default - becomes emigo.Nullable[int64].",
		},
		{
			Name:        prefix + "rating",
			Type:        "float64?",
			Description: "Off by default - becomes emigo.Nullable[float64].",
		},
		{
			Name:        prefix + "address",
			Type:        "object?",
			Description: "Off by default - becomes emigo.Nullable[SelectUserByIdVsqlRowAddress], a real nested struct, not two loose strings. The `column:` here isn't a plain column name at all - it's a json_object(...) SQL expression aggregating street+city into one JSON text value, which the caller then json.Unmarshal's into that struct after scanning. Demonstrates that `column:` accepts any SQL expression, not just identifiers.",
		},
	}
}
func CastSelectUserByIdVsqlRowFromCli(c emigo.CliCastable) SelectUserByIdVsqlRow {
	data := SelectUserByIdVsqlRow{}
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
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	if c.IsSet("rating") {
		emigo.ParseNullable(c.String("rating"), &data.Rating)
	}
	if c.IsSet("address") {
		emigo.ParseNullable(c.String("address"), &data.Address)
	}
	return data
}
func GetSelectUserByIdVsqlRowAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "street",
			Type:        "string",
			Description: "Street address line.",
		},
		{
			Name:        prefix + "city",
			Type:        "string",
			Description: "City name.",
		},
	}
}
func CastSelectUserByIdVsqlRowAddressFromCli(c emigo.CliCastable) SelectUserByIdVsqlRowAddress {
	data := SelectUserByIdVsqlRowAddress{}
	if c.IsSet("street") {
		data.Street = c.String("street")
	}
	if c.IsSet("city") {
		data.City = c.String("city")
	}
	return data
}

// SelectUserByIdVsqlColumns is the column picker for the selectUserById vsql query. Toggle Selected on
// each entry (or start from NewSelectUserByIdVsqlColumns()), then pass it to PrepareSelectUserByIdVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type SelectUserByIdVsqlColumns struct {
	// Selected by default - stays a plain int64 in the row DTO.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Selected by default - stays a plain string.
	Email emigo.ColumnState `json:"email" yaml:"email"`
	// Selected by default - stays a plain string (enums have no separate Go type). Column name overridden to user_role.
	Role emigo.ColumnState `json:"role" yaml:"role"`
	// Off by default - becomes emigo.Nullable[bool] in the row DTO, distinguishing "not fetched" from "fetched, false".
	IsActive emigo.ColumnState `json:"isActive" yaml:"isActive"`
	// Off by default - becomes emigo.Nullable[int64].
	BalanceCents emigo.ColumnState `json:"balanceCents" yaml:"balanceCents"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.ColumnState `json:"rating" yaml:"rating"`
	// Off by default - becomes emigo.Nullable[SelectUserByIdVsqlRowAddress], a real nested struct, not two loose strings. The `column:` here isn't a plain column name at all - it's a json_object(...) SQL expression aggregating street+city into one JSON text value, which the caller then json.Unmarshal's into that struct after scanning. Demonstrates that `column:` accepts any SQL expression, not just identifiers.
	Address emigo.ColumnState `json:"address" yaml:"address"`
}

// NewSelectUserByIdVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the selectUserById vsql query.
func NewSelectUserByIdVsqlColumns() SelectUserByIdVsqlColumns {
	return SelectUserByIdVsqlColumns{
		Id:           emigo.ColumnState{Selected: true},
		Email:        emigo.ColumnState{Selected: true},
		Role:         emigo.ColumnState{Selected: true},
		IsActive:     emigo.ColumnState{Selected: false},
		BalanceCents: emigo.ColumnState{Selected: false},
		Rating:       emigo.ColumnState{Selected: false},
		Address:      emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c SelectUserByIdVsqlColumns) Cols() string {
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
	if c.BalanceCents.Selected {
		parts = append(parts, "balance_cents")
	}
	if c.Rating.Selected {
		parts = append(parts, "rating")
	}
	if c.Address.Selected {
		parts = append(parts, "json_object('street', street, 'city', city)")
	}
	return strings.Join(parts, ", ")
}

// SelectUserByIdVsqlColumnsQuery lets a caller build SelectUserByIdVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing SelectUserByIdVsqlColumns by hand - see Columns().
type SelectUserByIdVsqlColumnsQuery struct {
	values       url.Values
	mapped       map[string]interface{}
	Id           bool `json:"id"`
	Email        bool `json:"email"`
	Role         bool `json:"role"`
	IsActive     bool `json:"isActive"`
	BalanceCents bool `json:"balanceCents"`
	Rating       bool `json:"rating"`
	Address      bool `json:"address"`
}

// SelectUserByIdVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// SelectUserByIdVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func SelectUserByIdVsqlColumnsQueryFromString(rawQuery string) SelectUserByIdVsqlColumnsQuery {
	q := SelectUserByIdVsqlColumnsQuery{}
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

// SelectUserByIdVsqlColumnsQueryFromHttp is SelectUserByIdVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func SelectUserByIdVsqlColumnsQueryFromHttp(r *http.Request) SelectUserByIdVsqlColumnsQuery {
	return SelectUserByIdVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a SelectUserByIdVsqlColumns, seeded from
// NewSelectUserByIdVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q SelectUserByIdVsqlColumnsQuery) Columns() SelectUserByIdVsqlColumns {
	cols := NewSelectUserByIdVsqlColumns()
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
	if q.values.Has("balanceCents") {
		cols.BalanceCents.Selected = q.BalanceCents
	}
	if q.values.Has("rating") {
		cols.Rating.Selected = q.Rating
	}
	if q.values.Has("address") {
		cols.Address.Selected = q.Address
	}
	return cols
}

// SelectUserByIdVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from SelectUserByIdVsqlParams) alongside {{ .Columns.Cols }}.
type SelectUserByIdVsqlData struct {
	SelectUserByIdVsqlParams
	Columns SelectUserByIdVsqlColumns `json:"columns" yaml:"columns"`
}

// SelectUserByIdVsqlName is the name of the vsql query, useful for logging or routing.
const SelectUserByIdVsqlName = "selectUserById"

// SelectUserByIdVsqlQuery is the raw SQL string for the selectUserById vsql query.
const SelectUserByIdVsqlQuery = `SELECT {{ .Columns.Cols }} FROM users WHERE id = {{ .Id }};
`

// PrepareSelectUserByIdVsql returns the query string and params for the selectUserById vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewSelectUserByIdVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into SelectUserByIdVsqlRow.
func PrepareSelectUserByIdVsql(params SelectUserByIdVsqlParams, columns SelectUserByIdVsqlColumns) (query string, args interface{}) {
	return SelectUserByIdVsqlQuery, SelectUserByIdVsqlData{SelectUserByIdVsqlParams: params, Columns: columns}
}
