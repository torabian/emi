package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for insertUsersBatchVsqlParams
type InsertUsersBatchVsqlParams struct {
	// One entry per user row to insert.
	Users emigo.Array[InsertUsersBatchVsqlParamsUsers] `json:"users" yaml:"users"`
}

// The base class definition for users
type InsertUsersBatchVsqlParamsUsers struct {
	// Required per entry, same as insertUser's.
	Email string `json:"email" yaml:"email"`
	// Required per entry.
	FirstName string `json:"firstName" yaml:"firstName"`
	// Required here (unlike insertUser) - kept non-nullable on purpose to show both variants exist side by side across these ten queries.
	LastName string `json:"lastName" yaml:"lastName"`
	// Opening balance, in integer cents (no float rounding surprises).
	MoneyAmountCents int64 `json:"moneyAmountCents" yaml:"moneyAmountCents"`
	// ISO 4217 currency code for moneyAmountCents, e.g. "USD".
	MoneyCurrency string `json:"moneyCurrency" yaml:"moneyCurrency"`
}

func (x *InsertUsersBatchVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersBatchVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "users",
			Type:        "array",
			Description: "One entry per user row to insert.",
		},
	}
}
func CastInsertUsersBatchVsqlParamsFromCli(c emigo.CliCastable) InsertUsersBatchVsqlParams {
	data := InsertUsersBatchVsqlParams{}
	if c.IsSet("users") {
		data.Users = emigo.CapturePossibleArray(CastInsertUsersBatchVsqlParamsUsersFromCli, "users", c)
	}
	return data
}
func GetInsertUsersBatchVsqlParamsUsersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Required per entry, same as insertUser's.",
		},
		{
			Name:        prefix + "first-name",
			Type:        "string",
			Description: "Required per entry.",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string",
			Description: "Required here (unlike insertUser) - kept non-nullable on purpose to show both variants exist side by side across these ten queries.",
		},
		{
			Name:        prefix + "money-amount-cents",
			Type:        "int64",
			Description: "Opening balance, in integer cents (no float rounding surprises).",
		},
		{
			Name:        prefix + "money-currency",
			Type:        "string",
			Description: "ISO 4217 currency code for moneyAmountCents, e.g. \"USD\".",
		},
	}
}
func CastInsertUsersBatchVsqlParamsUsersFromCli(c emigo.CliCastable) InsertUsersBatchVsqlParamsUsers {
	data := InsertUsersBatchVsqlParamsUsers{}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("money-amount-cents") {
		data.MoneyAmountCents = int64(c.Int64("money-amount-cents"))
	}
	if c.IsSet("money-currency") {
		data.MoneyCurrency = c.String("money-currency")
	}
	return data
}

// InsertUsersBatchVsqlFilters is the allow-list DTO for the insertUsersBatch vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for insertUsersBatchVsqlFilters
type InsertUsersBatchVsqlFilters struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default.
	Email string `json:"email" yaml:"email"`
	// Off by default. Money is a consumer-owned Go type (see sdkgen/money.go) backed by two physical columns (money_amount_cents, money_currency) - `column:` lists both under this one Selected toggle, and Cols() splices them in verbatim. `complex` has no wire-level nullability of its own, so being off by default only flips the type to `complex?`, a no-op for Go - Money.SQLValue is what decides how an absent amount would render, not this field's nullability.
	Money Money `json:"money" yaml:"money"`
}

func (x *InsertUsersBatchVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersBatchVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "money",
			Type:        "complex",
			Description: "Off by default. Money is a consumer-owned Go type (see sdkgen/money.go) backed by two physical columns (money_amount_cents, money_currency) - `column:` lists both under this one Selected toggle, and Cols() splices them in verbatim. `complex` has no wire-level nullability of its own, so being off by default only flips the type to `complex?`, a no-op for Go - Money.SQLValue is what decides how an absent amount would render, not this field's nullability.",
		},
	}
}
func CastInsertUsersBatchVsqlFiltersFromCli(c emigo.CliCastable) InsertUsersBatchVsqlFilters {
	data := InsertUsersBatchVsqlFilters{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("money") {
		if u, ok := any(&data.Money).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("money")))
		}
	}
	return data
}

// InsertUsersBatchVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// insertUsersBatch vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only InsertUsersBatchVsqlFilters's own
// top-level entries.
var InsertUsersBatchVsqlFilterFields = []string{
	"id",
	"email",
	"money",
}

// InsertUsersBatchVsqlFilterFieldAllowed reports whether name is one of InsertUsersBatchVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func InsertUsersBatchVsqlFilterFieldAllowed(name string) bool {
	for _, f := range InsertUsersBatchVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// InsertUsersBatchVsqlRow is the response/row DTO for the insertUsersBatch vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for insertUsersBatchVsqlRow
type InsertUsersBatchVsqlRow struct {
	// Selected by default.
	Id int64 `json:"id" yaml:"id"`
	// Selected by default.
	Email string `json:"email" yaml:"email"`
	// Off by default. Money is a consumer-owned Go type (see sdkgen/money.go) backed by two physical columns (money_amount_cents, money_currency) - `column:` lists both under this one Selected toggle, and Cols() splices them in verbatim. `complex` has no wire-level nullability of its own, so being off by default only flips the type to `complex?`, a no-op for Go - Money.SQLValue is what decides how an absent amount would render, not this field's nullability.
	Money Money `json:"money" yaml:"money"`
}

func (x *InsertUsersBatchVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersBatchVsqlRowCliFlags(prefix string) []emigo.CliFlag {
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
			Name:        prefix + "money",
			Type:        "complex",
			Description: "Off by default. Money is a consumer-owned Go type (see sdkgen/money.go) backed by two physical columns (money_amount_cents, money_currency) - `column:` lists both under this one Selected toggle, and Cols() splices them in verbatim. `complex` has no wire-level nullability of its own, so being off by default only flips the type to `complex?`, a no-op for Go - Money.SQLValue is what decides how an absent amount would render, not this field's nullability.",
		},
	}
}
func CastInsertUsersBatchVsqlRowFromCli(c emigo.CliCastable) InsertUsersBatchVsqlRow {
	data := InsertUsersBatchVsqlRow{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("money") {
		if u, ok := any(&data.Money).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("money")))
		}
	}
	return data
}

// InsertUsersBatchVsqlColumns is the column picker for the insertUsersBatch vsql query. Toggle Selected on
// each entry (or start from NewInsertUsersBatchVsqlColumns()), then pass it to PrepareInsertUsersBatchVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type InsertUsersBatchVsqlColumns struct {
	// Selected by default.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Selected by default.
	Email emigo.ColumnState `json:"email" yaml:"email"`
	// Off by default. Money is a consumer-owned Go type (see sdkgen/money.go) backed by two physical columns (money_amount_cents, money_currency) - `column:` lists both under this one Selected toggle, and Cols() splices them in verbatim. `complex` has no wire-level nullability of its own, so being off by default only flips the type to `complex?`, a no-op for Go - Money.SQLValue is what decides how an absent amount would render, not this field's nullability.
	Money emigo.ColumnState `json:"money" yaml:"money"`
}

// NewInsertUsersBatchVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the insertUsersBatch vsql query.
func NewInsertUsersBatchVsqlColumns() InsertUsersBatchVsqlColumns {
	return InsertUsersBatchVsqlColumns{
		Id:    emigo.ColumnState{Selected: true},
		Email: emigo.ColumnState{Selected: true},
		Money: emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c InsertUsersBatchVsqlColumns) Cols() string {
	parts := []string{}
	if c.Id.Selected {
		parts = append(parts, "id")
	}
	if c.Email.Selected {
		parts = append(parts, "email")
	}
	if c.Money.Selected {
		parts = append(parts, "money_amount_cents, money_currency")
	}
	return strings.Join(parts, ", ")
}

// InsertUsersBatchVsqlColumnsQuery lets a caller build InsertUsersBatchVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing InsertUsersBatchVsqlColumns by hand - see Columns().
type InsertUsersBatchVsqlColumnsQuery struct {
	values url.Values
	mapped map[string]interface{}
	Id     bool `json:"id"`
	Email  bool `json:"email"`
	Money  bool `json:"money"`
}

// InsertUsersBatchVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// InsertUsersBatchVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func InsertUsersBatchVsqlColumnsQueryFromString(rawQuery string) InsertUsersBatchVsqlColumnsQuery {
	q := InsertUsersBatchVsqlColumnsQuery{}
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

// InsertUsersBatchVsqlColumnsQueryFromHttp is InsertUsersBatchVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func InsertUsersBatchVsqlColumnsQueryFromHttp(r *http.Request) InsertUsersBatchVsqlColumnsQuery {
	return InsertUsersBatchVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a InsertUsersBatchVsqlColumns, seeded from
// NewInsertUsersBatchVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q InsertUsersBatchVsqlColumnsQuery) Columns() InsertUsersBatchVsqlColumns {
	cols := NewInsertUsersBatchVsqlColumns()
	if q.values.Has("id") {
		cols.Id.Selected = q.Id
	}
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	if q.values.Has("money") {
		cols.Money.Selected = q.Money
	}
	return cols
}

// InsertUsersBatchVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from InsertUsersBatchVsqlParams) alongside {{ .Columns.Cols }}.
type InsertUsersBatchVsqlData struct {
	InsertUsersBatchVsqlParams
	Columns InsertUsersBatchVsqlColumns `json:"columns" yaml:"columns"`
}

// InsertUsersBatchVsqlName is the name of the vsql query, useful for logging or routing.
const InsertUsersBatchVsqlName = "insertUsersBatch"

// InsertUsersBatchVsqlQuery is the raw SQL string for the insertUsersBatch vsql query.
const InsertUsersBatchVsqlQuery = `INSERT INTO users (email, first_name, last_name, money_amount_cents, money_currency)
VALUES
{{- range $i, $u := .Users.Items }}{{ if $i }},{{ end }}
  ({{ sql $u.Email }}, {{ sql $u.FirstName }}, {{ sql $u.LastName }}, {{ sql $u.MoneyAmountCents }}, {{ sql $u.MoneyCurrency }})
{{- end }}
RETURNING {{ .Columns.Cols }};
`

// PrepareInsertUsersBatchVsql returns the query string and params for the insertUsersBatch vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewInsertUsersBatchVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into InsertUsersBatchVsqlRow.
func PrepareInsertUsersBatchVsql(params InsertUsersBatchVsqlParams, columns InsertUsersBatchVsqlColumns) (query string, args interface{}) {
	return InsertUsersBatchVsqlQuery, InsertUsersBatchVsqlData{InsertUsersBatchVsqlParams: params, Columns: columns}
}
