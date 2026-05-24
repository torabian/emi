package dto

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetUpdateUserDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64?",
		},
		{
			Name: prefix + "name",
			Type: "string?",
		},
		{
			Name: prefix + "email",
			Type: "string?",
		},
		{
			Name: prefix + "age",
			Type: "int?",
		},
		{
			Name: prefix + "preferences",
			Type: "object?",
		},
		{
			Name: prefix + "tags",
			Type: "array?",
		},
	}
}
func CastUpdateUserDtoFromCli(c emigo.CliCastable) UpdateUserDto {
	data := UpdateUserDto{}
	if c.IsSet("id") {
		emigo.ParseNullable(c.String("id"), &data.Id)
	}
	if c.IsSet("name") {
		emigo.ParseNullable(c.String("name"), &data.Name)
	}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("age") {
		emigo.ParseNullable(c.String("age"), &data.Age)
	}
	if c.IsSet("preferences") {
		emigo.ParseNullable(c.String("preferences"), &data.Preferences)
	}
	if c.IsSet("tags") {
		data.Tags = emigo.CapturePossibleArrayNullable(CastUpdateUserDtoTagsFromCli, "tags", c)
	}
	return data
}

// The base class definition for updateUserDto
type UpdateUserDto struct {
	Id          emigo.Nullable[int64]                    `json:"id" yaml:"id"`
	Name        emigo.Nullable[string]                   `json:"name" yaml:"name"`
	Email       emigo.Nullable[string]                   `json:"email" yaml:"email"`
	Age         emigo.Nullable[int]                      `json:"age" yaml:"age"`
	Preferences emigo.Nullable[UpdateUserDtoPreferences] `json:"preferences" yaml:"preferences"`
	Tags        emigo.ArrayNullable[UpdateUserDtoTags]   `json:"tags" yaml:"tags"`
}

func GetUpdateUserDtoPreferencesCliFlags(prefix string) []emigo.CliFlag {
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
func CastUpdateUserDtoPreferencesFromCli(c emigo.CliCastable) UpdateUserDtoPreferences {
	data := UpdateUserDtoPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// The base class definition for preferences
type UpdateUserDtoPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func GetUpdateUserDtoTagsCliFlags(prefix string) []emigo.CliFlag {
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
func CastUpdateUserDtoTagsFromCli(c emigo.CliCastable) UpdateUserDtoTags {
	data := UpdateUserDtoTags{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for tags
type UpdateUserDtoTags struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

func (x *UpdateUserDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
