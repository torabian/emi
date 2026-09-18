package external

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"net/http"
	"net/url"
	"strings"
)

// The base class definition for insertUsersVsqlParams
type InsertUsersVsqlParams struct {
	Users emigo.Array[InsertUsersVsqlParamsUsers] `json:"users" yaml:"users"`
}

// The base class definition for users
type InsertUsersVsqlParamsUsers struct {
	Email     string `json:"email" yaml:"email"`
	FirstName string `json:"firstName" yaml:"firstName"`
	LastName  string `json:"lastName" yaml:"lastName"`
}

func (x *InsertUsersVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "users",
			Type: "array",
		},
	}
}
func CastInsertUsersVsqlParamsFromCli(c emigo.CliCastable) InsertUsersVsqlParams {
	data := InsertUsersVsqlParams{}
	if c.IsSet("users") {
		data.Users = emigo.CapturePossibleArray(CastInsertUsersVsqlParamsUsersFromCli, "users", c)
	}
	return data
}
func GetInsertUsersVsqlParamsUsersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "email",
			Type: "string",
		},
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "last-name",
			Type: "string",
		},
	}
}
func CastInsertUsersVsqlParamsUsersFromCli(c emigo.CliCastable) InsertUsersVsqlParamsUsers {
	data := InsertUsersVsqlParamsUsers{}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	return data
}

// InsertUsersVsqlFilters is the allow-list DTO for the insertUsers vsql query's filter -
// see EmiVsql.Filters. Emi never executes or transpiles a filter itself
// (e.g. via jsonlogic2sql); this only exists to describe, and let you
// validate against, the fields a caller-supplied JSON-Logic expression is
// allowed to reference.
// The base class definition for insertUsersVsqlFilters
type InsertUsersVsqlFilters struct {
	// Primary key. Selected by default - stays a plain int64.
	Id int64 `json:"id" yaml:"id"`
	// Off by default - becomes emigo.Nullable[int64] in the row DTO.
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
	// Selected by default - stays a plain string.
	Email     string                 `json:"email" yaml:"email"`
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	// Selected by default - stays a plain bool.
	IsVerified  bool                 `json:"isVerified" yaml:"isVerified"`
	IsSuspended emigo.Nullable[bool] `json:"isSuspended" yaml:"isSuspended"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.Nullable[float64] `json:"rating" yaml:"rating"`
	// Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
	Status string `json:"status" yaml:"status"`
	// Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
	// Selected by default - stays the plain nested struct.
	Profile InsertUsersVsqlFiltersProfile `json:"profile" yaml:"profile"`
	// Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
	Preferences emigo.Nullable[InsertUsersVsqlFiltersPreferences] `json:"preferences" yaml:"preferences"`
	// Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
	Money Money `json:"money" yaml:"money"`
}

// The base class definition for profile
type InsertUsersVsqlFiltersProfile struct {
	Bio string `json:"bio" yaml:"bio"`
}

// The base class definition for preferences
type InsertUsersVsqlFiltersPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func (x *InsertUsersVsqlFilters) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersVsqlFiltersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Primary key. Selected by default - stays a plain int64.",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Off by default - becomes emigo.Nullable[int64] in the row DTO.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - stays a plain string.",
		},
		{
			Name: prefix + "first-name",
			Type: "string?",
		},
		{
			Name:        prefix + "is-verified",
			Type:        "bool",
			Description: "Selected by default - stays a plain bool.",
		},
		{
			Name: prefix + "is-suspended",
			Type: "bool?",
		},
		{
			Name:        prefix + "rating",
			Type:        "float64?",
			Description: "Off by default - becomes emigo.Nullable[float64].",
		},
		{
			Name:        prefix + "status",
			Type:        "enum",
			Description: "Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "Off by default - becomes emigo.Nullable[string], same as any other nullable enum.",
		},
		{
			Name:        prefix + "profile",
			Type:        "object",
			Children:    GetInsertUsersVsqlFiltersProfileCliFlags("profile-"),
			Description: "Selected by default - stays the plain nested struct.",
		},
		{
			Name:        prefix + "preferences",
			Type:        "object?",
			Description: "Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so \"not fetched\" is distinguishable from \"fetched, all fields empty\".",
		},
		{
			Name:        prefix + "money",
			Type:        "complex",
			Description: "Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).",
		},
	}
}
func CastInsertUsersVsqlFiltersFromCli(c emigo.CliCastable) InsertUsersVsqlFilters {
	data := InsertUsersVsqlFilters{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("is-verified") {
		data.IsVerified = bool(c.Bool("is-verified"))
	}
	if c.IsSet("is-suspended") {
		emigo.ParseNullable(c.String("is-suspended"), &data.IsSuspended)
	}
	if c.IsSet("rating") {
		emigo.ParseNullable(c.String("rating"), &data.Rating)
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	if c.IsSet("profile") {
		data.Profile = CastInsertUsersVsqlFiltersProfileFromCli(c)
	}
	if c.IsSet("preferences") {
		emigo.ParseNullable(c.String("preferences"), &data.Preferences)
	}
	if c.IsSet("money") {
		if u, ok := any(&data.Money).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("money")))
		}
	}
	return data
}
func GetInsertUsersVsqlFiltersProfileCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "bio",
			Type: "string",
		},
	}
}
func CastInsertUsersVsqlFiltersProfileFromCli(c emigo.CliCastable) InsertUsersVsqlFiltersProfile {
	data := InsertUsersVsqlFiltersProfile{}
	if c.IsSet("bio") {
		data.Bio = c.String("bio")
	}
	return data
}
func GetInsertUsersVsqlFiltersPreferencesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "theme",
			Type: "string",
		},
		{
			Name: prefix + "locale",
			Type: "string",
		},
	}
}
func CastInsertUsersVsqlFiltersPreferencesFromCli(c emigo.CliCastable) InsertUsersVsqlFiltersPreferences {
	data := InsertUsersVsqlFiltersPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// InsertUsersVsqlFilterFields lists the top-level field names a JSON-Logic filter against the
// insertUsers vsql query may legally reference (each {"var": ...} leaf). Does not
// enumerate a nested object field's own sub-fields - only InsertUsersVsqlFilters's own
// top-level entries.
var InsertUsersVsqlFilterFields = []string{
	"id",
	"balanceCents",
	"email",
	"firstName",
	"isVerified",
	"isSuspended",
	"rating",
	"status",
	"role",
	"profile",
	"preferences",
	"money",
}

// InsertUsersVsqlFilterFieldAllowed reports whether name is one of InsertUsersVsqlFilterFields's allowed
// top-level filter fields - a cheap check worth running against every
// {"var": ...} leaf found while walking a caller-supplied filter, before
// it ever reaches SQL.
func InsertUsersVsqlFilterFieldAllowed(name string) bool {
	for _, f := range InsertUsersVsqlFilterFields {
		if f == name {
			return true
		}
	}
	return false
}

// InsertUsersVsqlRow is the response/row DTO for the insertUsers vsql query - one
// field per entry in Columns, typed exactly as it would be as a dto field.
// It exists whether or not a given field was actually selected; scanning a
// query built with fewer columns selected just leaves the rest zero-valued.
// The base class definition for insertUsersVsqlRow
type InsertUsersVsqlRow struct {
	// Primary key. Selected by default - stays a plain int64.
	Id int64 `json:"id" yaml:"id"`
	// Off by default - becomes emigo.Nullable[int64] in the row DTO.
	BalanceCents emigo.Nullable[int64] `json:"balanceCents" yaml:"balanceCents"`
	// Selected by default - stays a plain string.
	Email     string                 `json:"email" yaml:"email"`
	FirstName emigo.Nullable[string] `json:"firstName" yaml:"firstName"`
	// Selected by default - stays a plain bool.
	IsVerified  bool                 `json:"isVerified" yaml:"isVerified"`
	IsSuspended emigo.Nullable[bool] `json:"isSuspended" yaml:"isSuspended"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.Nullable[float64] `json:"rating" yaml:"rating"`
	// Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
	Status string `json:"status" yaml:"status"`
	// Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
	// Selected by default - stays the plain nested struct.
	Profile InsertUsersVsqlRowProfile `json:"profile" yaml:"profile"`
	// Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
	Preferences emigo.Nullable[InsertUsersVsqlRowPreferences] `json:"preferences" yaml:"preferences"`
	// Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
	Money Money `json:"money" yaml:"money"`
}

// The base class definition for profile
type InsertUsersVsqlRowProfile struct {
	Bio string `json:"bio" yaml:"bio"`
}

// The base class definition for preferences
type InsertUsersVsqlRowPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func (x *InsertUsersVsqlRow) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetInsertUsersVsqlRowCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int64",
			Description: "Primary key. Selected by default - stays a plain int64.",
		},
		{
			Name:        prefix + "balance-cents",
			Type:        "int64?",
			Description: "Off by default - becomes emigo.Nullable[int64] in the row DTO.",
		},
		{
			Name:        prefix + "email",
			Type:        "string",
			Description: "Selected by default - stays a plain string.",
		},
		{
			Name: prefix + "first-name",
			Type: "string?",
		},
		{
			Name:        prefix + "is-verified",
			Type:        "bool",
			Description: "Selected by default - stays a plain bool.",
		},
		{
			Name: prefix + "is-suspended",
			Type: "bool?",
		},
		{
			Name:        prefix + "rating",
			Type:        "float64?",
			Description: "Off by default - becomes emigo.Nullable[float64].",
		},
		{
			Name:        prefix + "status",
			Type:        "enum",
			Description: "Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "Off by default - becomes emigo.Nullable[string], same as any other nullable enum.",
		},
		{
			Name:        prefix + "profile",
			Type:        "object",
			Children:    GetInsertUsersVsqlRowProfileCliFlags("profile-"),
			Description: "Selected by default - stays the plain nested struct.",
		},
		{
			Name:        prefix + "preferences",
			Type:        "object?",
			Description: "Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so \"not fetched\" is distinguishable from \"fetched, all fields empty\".",
		},
		{
			Name:        prefix + "money",
			Type:        "complex",
			Description: "Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).",
		},
	}
}
func CastInsertUsersVsqlRowFromCli(c emigo.CliCastable) InsertUsersVsqlRow {
	data := InsertUsersVsqlRow{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("balance-cents") {
		emigo.ParseNullable(c.String("balance-cents"), &data.BalanceCents)
	}
	if c.IsSet("email") {
		data.Email = c.String("email")
	}
	if c.IsSet("first-name") {
		emigo.ParseNullable(c.String("first-name"), &data.FirstName)
	}
	if c.IsSet("is-verified") {
		data.IsVerified = bool(c.Bool("is-verified"))
	}
	if c.IsSet("is-suspended") {
		emigo.ParseNullable(c.String("is-suspended"), &data.IsSuspended)
	}
	if c.IsSet("rating") {
		emigo.ParseNullable(c.String("rating"), &data.Rating)
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	if c.IsSet("profile") {
		data.Profile = CastInsertUsersVsqlRowProfileFromCli(c)
	}
	if c.IsSet("preferences") {
		emigo.ParseNullable(c.String("preferences"), &data.Preferences)
	}
	if c.IsSet("money") {
		if u, ok := any(&data.Money).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("money")))
		}
	}
	return data
}
func GetInsertUsersVsqlRowProfileCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "bio",
			Type: "string",
		},
	}
}
func CastInsertUsersVsqlRowProfileFromCli(c emigo.CliCastable) InsertUsersVsqlRowProfile {
	data := InsertUsersVsqlRowProfile{}
	if c.IsSet("bio") {
		data.Bio = c.String("bio")
	}
	return data
}
func GetInsertUsersVsqlRowPreferencesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "theme",
			Type: "string",
		},
		{
			Name: prefix + "locale",
			Type: "string",
		},
	}
}
func CastInsertUsersVsqlRowPreferencesFromCli(c emigo.CliCastable) InsertUsersVsqlRowPreferences {
	data := InsertUsersVsqlRowPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// InsertUsersVsqlColumns is the column picker for the insertUsers vsql query. Toggle Selected on
// each entry (or start from NewInsertUsersVsqlColumns()), then pass it to PrepareInsertUsersVsql -
// the query template reads {{ if .Columns.<Name>.Selected }} and {{ .Columns.Cols }}.
type InsertUsersVsqlColumns struct {
	// Primary key. Selected by default - stays a plain int64.
	Id emigo.ColumnState `json:"id" yaml:"id"`
	// Off by default - becomes emigo.Nullable[int64] in the row DTO.
	BalanceCents emigo.ColumnState `json:"balanceCents" yaml:"balanceCents"`
	// Selected by default - stays a plain string.
	Email     emigo.ColumnState `json:"email" yaml:"email"`
	FirstName emigo.ColumnState `json:"firstName" yaml:"firstName"`
	// Selected by default - stays a plain bool.
	IsVerified  emigo.ColumnState `json:"isVerified" yaml:"isVerified"`
	IsSuspended emigo.ColumnState `json:"isSuspended" yaml:"isSuspended"`
	// Off by default - becomes emigo.Nullable[float64].
	Rating emigo.ColumnState `json:"rating" yaml:"rating"`
	// Selected by default - enums render as plain string. The real column is named account_status, not status - `column:` overrides the default snake_cased-Name projection for exactly this kind of mismatch.
	Status emigo.ColumnState `json:"status" yaml:"status"`
	// Off by default - becomes emigo.Nullable[string], same as any other nullable enum.
	Role emigo.ColumnState `json:"role" yaml:"role"`
	// Selected by default - stays the plain nested struct.
	Profile emigo.ColumnState `json:"profile" yaml:"profile"`
	// Off by default - becomes emigo.Nullable[InsertUsersVsqlRowPreferences], not just a plain nested struct with zero values, so "not fetched" is distinguishable from "fetched, all fields empty".
	Preferences emigo.ColumnState `json:"preferences" yaml:"preferences"`
	// Demonstrates a complex column spanning two physical SQL columns under one projection/selection toggle - see sdkgen/money.go. `complex` has no wire-level nullability of its own (unlike a primitive or object), so being off by default only flips its type to `complex?`, a no-op for the Go backend - Money is responsible for its own optionality (see Money.SQLValue).
	Money emigo.ColumnState `json:"money" yaml:"money"`
}

// NewInsertUsersVsqlColumns returns the column picker seeded with each column's
// default Selected state, as declared on the insertUsers vsql query.
func NewInsertUsersVsqlColumns() InsertUsersVsqlColumns {
	return InsertUsersVsqlColumns{
		Id:           emigo.ColumnState{Selected: true},
		BalanceCents: emigo.ColumnState{Selected: false},
		Email:        emigo.ColumnState{Selected: true},
		FirstName:    emigo.ColumnState{Selected: false},
		IsVerified:   emigo.ColumnState{Selected: true},
		IsSuspended:  emigo.ColumnState{Selected: false},
		Rating:       emigo.ColumnState{Selected: false},
		Status:       emigo.ColumnState{Selected: true},
		Role:         emigo.ColumnState{Selected: false},
		Profile:      emigo.ColumnState{Selected: true},
		Preferences:  emigo.ColumnState{Selected: false},
		Money:        emigo.ColumnState{Selected: false},
	}
}

// Cols renders the currently-selected columns as a comma-separated SQL
// projection, in declaration order, so the query template doesn't have to
// repeat the column list.
func (c InsertUsersVsqlColumns) Cols() string {
	parts := []string{}
	if c.Id.Selected {
		parts = append(parts, "id")
	}
	if c.BalanceCents.Selected {
		parts = append(parts, "balance_cents")
	}
	if c.Email.Selected {
		parts = append(parts, "email")
	}
	if c.FirstName.Selected {
		parts = append(parts, "first_name")
	}
	if c.IsVerified.Selected {
		parts = append(parts, "is_verified")
	}
	if c.IsSuspended.Selected {
		parts = append(parts, "is_suspended")
	}
	if c.Rating.Selected {
		parts = append(parts, "rating")
	}
	if c.Status.Selected {
		parts = append(parts, "account_status")
	}
	if c.Role.Selected {
		parts = append(parts, "role")
	}
	if c.Profile.Selected {
		parts = append(parts, "profile")
	}
	if c.Preferences.Selected {
		parts = append(parts, "preferences")
	}
	if c.Money.Selected {
		parts = append(parts, "amount_cents, currency")
	}
	return strings.Join(parts, ", ")
}

// InsertUsersVsqlColumnsQuery lets a caller build InsertUsersVsqlColumns straight from a raw
// query string (e.g. "id=true&email=false") or an *http.Request, instead
// of constructing InsertUsersVsqlColumns by hand - see Columns().
type InsertUsersVsqlColumnsQuery struct {
	values       url.Values
	mapped       map[string]interface{}
	Id           bool `json:"id"`
	BalanceCents bool `json:"balanceCents"`
	Email        bool `json:"email"`
	FirstName    bool `json:"firstName"`
	IsVerified   bool `json:"isVerified"`
	IsSuspended  bool `json:"isSuspended"`
	Rating       bool `json:"rating"`
	Status       bool `json:"status"`
	Role         bool `json:"role"`
	Profile      bool `json:"profile"`
	Preferences  bool `json:"preferences"`
	Money        bool `json:"money"`
}

// InsertUsersVsqlColumnsQueryFromString parses rawQuery (a URL query string, no leading "?") into a
// InsertUsersVsqlColumnsQuery. Keys it doesn't recognize are ignored; see Columns() for how
// recognized ones fold into the actual picker.
func InsertUsersVsqlColumnsQueryFromString(rawQuery string) InsertUsersVsqlColumnsQuery {
	q := InsertUsersVsqlColumnsQuery{}
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

// InsertUsersVsqlColumnsQueryFromHttp is InsertUsersVsqlColumnsQueryFromString applied to r.URL.RawQuery.
func InsertUsersVsqlColumnsQueryFromHttp(r *http.Request) InsertUsersVsqlColumnsQuery {
	return InsertUsersVsqlColumnsQueryFromString(r.URL.RawQuery)
}

// Columns folds the parsed query string into a InsertUsersVsqlColumns, seeded from
// NewInsertUsersVsqlColumns() so a column the query string never mentioned keeps its
// own declared default rather than becoming false.
func (q InsertUsersVsqlColumnsQuery) Columns() InsertUsersVsqlColumns {
	cols := NewInsertUsersVsqlColumns()
	if q.values.Has("id") {
		cols.Id.Selected = q.Id
	}
	if q.values.Has("balanceCents") {
		cols.BalanceCents.Selected = q.BalanceCents
	}
	if q.values.Has("email") {
		cols.Email.Selected = q.Email
	}
	if q.values.Has("firstName") {
		cols.FirstName.Selected = q.FirstName
	}
	if q.values.Has("isVerified") {
		cols.IsVerified.Selected = q.IsVerified
	}
	if q.values.Has("isSuspended") {
		cols.IsSuspended.Selected = q.IsSuspended
	}
	if q.values.Has("rating") {
		cols.Rating.Selected = q.Rating
	}
	if q.values.Has("status") {
		cols.Status.Selected = q.Status
	}
	if q.values.Has("role") {
		cols.Role.Selected = q.Role
	}
	if q.values.Has("profile") {
		cols.Profile.Selected = q.Profile
	}
	if q.values.Has("preferences") {
		cols.Preferences.Selected = q.Preferences
	}
	if q.values.Has("money") {
		cols.Money.Selected = q.Money
	}
	return cols
}

// InsertUsersVsqlData is the combined query template data: the query params
// plus the caller's column selection in one struct, so a query template sees
// {{ .Name }} (promoted from InsertUsersVsqlParams) alongside {{ .Columns.Cols }}.
type InsertUsersVsqlData struct {
	InsertUsersVsqlParams
	Columns InsertUsersVsqlColumns `json:"columns" yaml:"columns"`
}

// InsertUsersVsqlName is the name of the vsql query, useful for logging or routing.
const InsertUsersVsqlName = "insertUsers"

// InsertUsersVsqlQuery is the raw SQL string for the insertUsers vsql query.
const InsertUsersVsqlQuery = `INSERT INTO users (email, first_name, last_name)
VALUES
{{- range $i, $u := .Users.Items }}{{ if $i }},{{ end }}
  ({{ sql $u.Email }}, {{ sql $u.FirstName }}, {{ sql $u.LastName }})
{{- end }}
RETURNING {{ .Columns.Cols }};
`

// PrepareInsertUsersVsql returns the query string and params for the insertUsers vsql query,
// ready to be passed to a SQL driver of your choice.
// columns controls which optional columns end up in the query - use
// NewInsertUsersVsqlColumns() to start from the defaults declared in the vsql definition.
// Scan results into InsertUsersVsqlRow.
func PrepareInsertUsersVsql(params InsertUsersVsqlParams, columns InsertUsersVsqlColumns) (query string, args interface{}) {
	return InsertUsersVsqlQuery, InsertUsersVsqlData{InsertUsersVsqlParams: params, Columns: columns}
}
