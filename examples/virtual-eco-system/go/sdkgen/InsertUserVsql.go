package external

import "encoding/json"
import emigo "test/emi/emigo"

func GetInsertUserVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "email",
			Type: "string?",
		},
		{
			Name: prefix + "age",
			Type: "int",
		},
		{
			Name:     prefix + "preferences",
			Type:     "object",
			Children: GetInsertUserVsqlParamsPreferencesCliFlags("preferences-"),
		},
		{
			Name: prefix + "tags",
			Type: "array",
		},
		{
			Name:        prefix + "address",
			Type:        "object",
			Children:    GetInsertUserVsqlParamsAddressCliFlags("address-"),
			Description: "Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.",
		},
	}
}
func CastInsertUserVsqlParamsFromCli(c emigo.CliCastable) InsertUserVsqlParams {
	data := InsertUserVsqlParams{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("age") {
		data.Age = int(c.Int64("age"))
	}
	if c.IsSet("preferences") {
		data.Preferences = CastInsertUserVsqlParamsPreferencesFromCli(c)
	}
	if c.IsSet("tags") {
		data.Tags = emigo.CapturePossibleArray(CastInsertUserVsqlParamsTagsFromCli, "tags", c)
	}
	if c.IsSet("address") {
		data.Address = CastInsertUserVsqlParamsAddressFromCli(c)
	}
	return data
}

// The base class definition for insertUserVsqlParams
type InsertUserVsqlParams struct {
	Name        string                          `json:"name" yaml:"name"`
	Email       emigo.Nullable[string]          `json:"email" yaml:"email"`
	Age         int                             `json:"age" yaml:"age"`
	Preferences InsertUserVsqlParamsPreferences `json:"preferences" yaml:"preferences"`
	Tags        []InsertUserVsqlParamsTags      `json:"tags" yaml:"tags"`
	// Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
	Address InsertUserVsqlParamsAddress `json:"address" yaml:"address"`
}

func GetInsertUserVsqlParamsPreferencesCliFlags(prefix string) []emigo.CliFlag {
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
func CastInsertUserVsqlParamsPreferencesFromCli(c emigo.CliCastable) InsertUserVsqlParamsPreferences {
	data := InsertUserVsqlParamsPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// The base class definition for preferences
type InsertUserVsqlParamsPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func GetInsertUserVsqlParamsTagsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "key",
			Type: "string",
		},
		{
			Name: prefix + "value",
			Type: "string",
		},
	}
}
func CastInsertUserVsqlParamsTagsFromCli(c emigo.CliCastable) InsertUserVsqlParamsTags {
	data := InsertUserVsqlParamsTags{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for tags
type InsertUserVsqlParamsTags struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

func GetInsertUserVsqlParamsAddressCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "user-id",
			Type: "int64",
		},
		{
			Name: prefix + "street",
			Type: "string",
		},
		{
			Name: prefix + "city",
			Type: "string",
		},
		{
			Name: prefix + "postcode",
			Type: "string?",
		},
	}
}
func CastInsertUserVsqlParamsAddressFromCli(c emigo.CliCastable) InsertUserVsqlParamsAddress {
	data := InsertUserVsqlParamsAddress{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("user-id") {
		data.UserId = int64(c.Int64("user-id"))
	}
	if c.IsSet("street") {
		data.Street = c.String("street")
	}
	if c.IsSet("city") {
		data.City = c.String("city")
	}
	if c.IsSet("postcode") {
		emigo.ParseNullable(c.String("postcode"), &data.Postcode)
	}
	return data
}

// The base class definition for address
type InsertUserVsqlParamsAddress struct {
	Id       int64                  `json:"id" yaml:"id"`
	UserId   int64                  `json:"userId" yaml:"userId"`
	Street   string                 `json:"street" yaml:"street"`
	City     string                 `json:"city" yaml:"city"`
	Postcode emigo.Nullable[string] `json:"postcode" yaml:"postcode"`
}

func (x *InsertUserVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// InsertUserVsqlName is the name of the vsql query, useful for logging or routing.
const InsertUserVsqlName = "insertUser"

// InsertUserVsqlQuery is the raw SQL string for the insertUser vsql query.
const InsertUserVsqlQuery = `
BEGIN;
{{ $u := sqlFieldsExcept . "tags" -}}
INSERT INTO users ({{ $u.Columns }})
VALUES ({{ $u.Values }});
{{ $a := sqlFields .Address -}}
INSERT INTO addresses ({{ $a.Columns }})
VALUES ({{ $a.Values }});
{{- if .Tags }}
INSERT INTO user_tags (user_id, key, value) VALUES
{{- range $i, $t := .Tags }}{{ if $i }},{{ end }}
  ({{ sql $.Id }}, {{ sql $t.Key }}, {{ sql $t.Value }})
{{- end }};
{{- end }}
COMMIT;
`

// PrepareInsertUserVsql returns the query string and params for the insertUser vsql query,
// ready to be passed to a SQL driver of your choice.
func PrepareInsertUserVsql(params InsertUserVsqlParams) (query string, args interface{}) {
	return InsertUserVsqlQuery, params
}
