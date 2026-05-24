package dto

import (
	"encoding"
	"encoding/json"
)
import emigo "github.com/torabian/emi/emigo"

func GetCreateUserDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
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
			Children: GetCreateUserDtoPreferencesCliFlags("preferences-"),
		},
		{
			Name: prefix + "tags",
			Type: "array",
		},
		{
			Name:        prefix + "address",
			Type:        "object",
			Children:    GetCreateUserDtoAddressCliFlags("address-"),
			Description: "Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.",
		},
	}
}
func CastCreateUserDtoFromCli(c emigo.CliCastable) CreateUserDto {
	data := CreateUserDto{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
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
		data.Preferences = CastCreateUserDtoPreferencesFromCli(c)
	}
	if c.IsSet("tags") {
		data.Tags = emigo.CapturePossibleArray(CastCreateUserDtoTagsFromCli, "tags", c)
	}
	if c.IsSet("address") {
		data.Address = CastCreateUserDtoAddressFromCli(c)
	}
	return data
}

// The base class definition for createUserDto
type CreateUserDto struct {
	Id          int64                          `json:"id" yaml:"id"`
	Name        string                         `json:"name" yaml:"name"`
	Email       emigo.Nullable[string]         `json:"email" yaml:"email"`
	Age         int                            `json:"age" yaml:"age"`
	Preferences CreateUserDtoPreferences       `json:"preferences" yaml:"preferences"`
	Tags        emigo.Array[CreateUserDtoTags] `json:"tags" yaml:"tags"`
	// Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
	Address CreateUserDtoAddress `json:"address" yaml:"address"`
}

func GetCreateUserDtoPreferencesCliFlags(prefix string) []emigo.CliFlag {
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
func CastCreateUserDtoPreferencesFromCli(c emigo.CliCastable) CreateUserDtoPreferences {
	data := CreateUserDtoPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// The base class definition for preferences
type CreateUserDtoPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func GetCreateUserDtoTagsCliFlags(prefix string) []emigo.CliFlag {
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
func CastCreateUserDtoTagsFromCli(c emigo.CliCastable) CreateUserDtoTags {
	data := CreateUserDtoTags{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for tags
type CreateUserDtoTags struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

func GetCreateUserDtoAddressCliFlags(prefix string) []emigo.CliFlag {
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
		{
			Name: prefix + "location",
			Type: "complex",
		},
	}
}
func CastCreateUserDtoAddressFromCli(c emigo.CliCastable) CreateUserDtoAddress {
	data := CreateUserDtoAddress{}
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
	if c.IsSet("location") {
		if u, ok := any(&data.Location).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("location")))
		}
	}
	return data
}

// The base class definition for address
type CreateUserDtoAddress struct {
	Id       int64                  `json:"id" yaml:"id"`
	UserId   int64                  `json:"userId" yaml:"userId"`
	Street   string                 `json:"street" yaml:"street"`
	City     string                 `json:"city" yaml:"city"`
	Postcode emigo.Nullable[string] `json:"postcode" yaml:"postcode"`
	Location GeoPoint               `json:"location" yaml:"location"`
}

func (x *CreateUserDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
