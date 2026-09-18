package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for userOrderTotalsVsqlParams
type UserOrderTotalsVsqlParams struct {
}

func (x *UserOrderTotalsVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserOrderTotalsVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastUserOrderTotalsVsqlParamsFromCli(c emigo.CliCastable) UserOrderTotalsVsqlParams {
	data := UserOrderTotalsVsqlParams{}
	return data
}

// UserOrderTotalsVsqlFilters is the allow-list DTO for the userOrderTotals vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for userOrderTotalsVsqlFilters
type UserOrderTotalsVsqlFilters struct {
	UId          string `json:"uId" yaml:"uId"`
	OrderId      string `json:"orderId" yaml:"orderId"`
	UEmail       string `json:"uEmail" yaml:"uEmail"`
	OAmountCents string `json:"oAmountCents" yaml:"oAmountCents"`
}

func (x *UserOrderTotalsVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserOrderTotalsVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "u-id",
			Type: "string",
		},
		{
			Name: prefix + "order-id",
			Type: "string",
		},
		{
			Name: prefix + "u-email",
			Type: "string",
		},
		{
			Name: prefix + "o-amount-cents",
			Type: "string",
		},
	}
}
func CastUserOrderTotalsVsqlFiltersFromCli(c emigo.CliCastable) UserOrderTotalsVsqlFilters {
	data := UserOrderTotalsVsqlFilters{}
	if c.IsSet("u-id") {
		data.UId = c.String("u-id")
	}
	if c.IsSet("order-id") {
		data.OrderId = c.String("order-id")
	}
	if c.IsSet("u-email") {
		data.UEmail = c.String("u-email")
	}
	if c.IsSet("o-amount-cents") {
		data.OAmountCents = c.String("o-amount-cents")
	}
	return data
}

// UserOrderTotalsVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// userOrderTotals vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only UserOrderTotalsVsqlFilters's own
// top-level entries.
var UserOrderTotalsVsqlFilterFields = []string{
	"uId",
	"orderId",
	"uEmail",
	"oAmountCents",
}

// UserOrderTotalsVsqlFilterFieldAllowed reports whether name is one of UserOrderTotalsVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func UserOrderTotalsVsqlFilterFieldAllowed(name string) bool {
	for _, f := range UserOrderTotalsVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// UserOrderTotalsVsqlRow is the response/row DTO for the userOrderTotals vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for userOrderTotalsVsqlRow
type UserOrderTotalsVsqlRow struct {
	UId          string `json:"uId" yaml:"uId"`
	OrderId      string `json:"orderId" yaml:"orderId"`
	UEmail       string `json:"uEmail" yaml:"uEmail"`
	OAmountCents string `json:"oAmountCents" yaml:"oAmountCents"`
}

func (x *UserOrderTotalsVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserOrderTotalsVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "u-id",
			Type: "string",
		},
		{
			Name: prefix + "order-id",
			Type: "string",
		},
		{
			Name: prefix + "u-email",
			Type: "string",
		},
		{
			Name: prefix + "o-amount-cents",
			Type: "string",
		},
	}
}
func CastUserOrderTotalsVsqlRowFromCli(c emigo.CliCastable) UserOrderTotalsVsqlRow {
	data := UserOrderTotalsVsqlRow{}
	if c.IsSet("u-id") {
		data.UId = c.String("u-id")
	}
	if c.IsSet("order-id") {
		data.OrderId = c.String("order-id")
	}
	if c.IsSet("u-email") {
		data.UEmail = c.String("u-email")
	}
	if c.IsSet("o-amount-cents") {
		data.OAmountCents = c.String("o-amount-cents")
	}
	return data
}

// UserOrderTotalsVsqlColumns is the column picker for the userOrderTotals vsql query. Toggle Selected on
// each entry (or start from NewUserOrderTotalsVsqlColumns()), then pass it to PrepareUserOrderTotalsVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type UserOrderTotalsVsqlColumns struct {
	UId          emigo.ColumnState `json:"uId" yaml:"uId"`
	OrderId      emigo.ColumnState `json:"orderId" yaml:"orderId"`
	UEmail       emigo.ColumnState `json:"uEmail" yaml:"uEmail"`
	OAmountCents emigo.ColumnState `json:"oAmountCents" yaml:"oAmountCents"`
}

// NewUserOrderTotalsVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the userOrderTotals vsql query.
func NewUserOrderTotalsVsqlColumns() UserOrderTotalsVsqlColumns {
	return UserOrderTotalsVsqlColumns{
		UId:          emigo.ColumnState{Selected: true},
		OrderId:      emigo.ColumnState{Selected: true},
		UEmail:       emigo.ColumnState{Selected: true},
		OAmountCents: emigo.ColumnState{Selected: true},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c UserOrderTotalsVsqlColumns) Cols() string {
	parts := []string{}
	if c.UId.Selected {
		parts = append(parts, "u.id")
	}
	if c.OrderId.Selected {
		parts = append(parts, "o.id AS order_id")
	}
	if c.UEmail.Selected {
		parts = append(parts, "u.email")
	}
	if c.OAmountCents.Selected {
		parts = append(parts, "o.amount_cents")
	}
	return strings.Join(parts, ", ")
}

// UserOrderTotalsVsqlColumnsQuery lets a caller build UserOrderTotalsVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing UserOrderTotalsVsqlColumns by hand - see Columns().
type UserOrderTotalsVsqlColumnsQuery struct {
	values       url.Values
	mapped       map[string]interface{}
	UId          bool `json:"uId"`
	OrderId      bool `json:"orderId"`
	UEmail       bool `json:"uEmail"`
	OAmountCents bool `json:"oAmountCents"`
}

// UserOrderTotalsVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// UserOrderTotalsVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func UserOrderTotalsVsqlColumnsQueryFromString(rawQuery string) UserOrderTotalsVsqlColumnsQuery {
	q := UserOrderTotalsVsqlColumnsQuery{}
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

// UserOrderTotalsVsqlColumnsQueryFromHttp is UserOrderTotalsVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func UserOrderTotalsVsqlColumnsQueryFromHttp(r *http.Request) UserOrderTotalsVsqlColumnsQuery {
	return UserOrderTotalsVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a UserOrderTotalsVsqlColumns, seeded from
// NewUserOrderTotalsVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q UserOrderTotalsVsqlColumnsQuery) Columns() UserOrderTotalsVsqlColumns {
	cols := NewUserOrderTotalsVsqlColumns()
	if q.values.Has("uId") {
		cols.UId.Selected = q.UId
	}
	if q.values.Has("orderId") {
		cols.OrderId.Selected = q.OrderId
	}
	if q.values.Has("uEmail") {
		cols.UEmail.Selected = q.UEmail
	}
	if q.values.Has("oAmountCents") {
		cols.OAmountCents.Selected = q.OAmountCents
	}
	return cols
}

// UserOrderTotalsVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from UserOrderTotalsVsqlParams) alongside {{ .Columns.Cols }}.
type UserOrderTotalsVsqlData struct {
	UserOrderTotalsVsqlParams
	Columns UserOrderTotalsVsqlColumns `json:"columns" yaml:"columns"`
}

// UserOrderTotalsVsqlName is the name of the vsql query, useful for logging or routing.
const UserOrderTotalsVsqlName = "userOrderTotals"

// UserOrderTotalsVsqlQuery is the raw SQL string for the userOrderTotals vsql query.
const UserOrderTotalsVsqlQuery = `SELECT u.id, o.id AS order_id, u.email, o.amount_cents
FROM users u
JOIN orders o ON o.user_id = u.id;
`

// PrepareUserOrderTotalsVsql returns the query string and params for the userOrderTotals vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewUserOrderTotalsVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into UserOrderTotalsVsqlRow.
func PrepareUserOrderTotalsVsql(params UserOrderTotalsVsqlParams, columns UserOrderTotalsVsqlColumns) (query string, args interface{}) {
	return UserOrderTotalsVsqlQuery, UserOrderTotalsVsqlData{UserOrderTotalsVsqlParams: params, Columns: columns}
}
