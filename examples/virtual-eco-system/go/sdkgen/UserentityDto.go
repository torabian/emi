package external

import "encoding/json"
import emigo "test/emi/emigo"

func GetUserentityDtoCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetUserentityDtoPreferencesCliFlags("preferences-"),
		},
		{
			Name: prefix + "tags",
			Type: "array",
		},
		{
			Name:        prefix + "address",
			Type:        "object",
			Children:    GetUserentityDtoAddressCliFlags("address-"),
			Description: "Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.",
		},
	}
}
func CastUserentityDtoFromCli(c emigo.CliCastable) UserentityDto {
	data := UserentityDto{}
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
		data.Preferences = CastUserentityDtoPreferencesFromCli(c)
	}
	if c.IsSet("tags") {
		data.Tags = emigo.CapturePossibleArray(CastUserentityDtoTagsFromCli, "tags", c)
	}
	if c.IsSet("address") {
		data.Address = CastUserentityDtoAddressFromCli(c)
	}
	return data
}

// The base class definition for userentityDto
type UserentityDto struct {
	Name        string                   `json:"name" yaml:"name"`
	Email       emigo.Nullable[string]   `json:"email" yaml:"email"`
	Age         int                      `json:"age" yaml:"age"`
	Preferences UserentityDtoPreferences `json:"preferences" yaml:"preferences"`
	Tags        []UserentityDtoTags      `json:"tags" yaml:"tags"`
	// Insert payload for the addresses table. The `location` field is a complex type (GeoPoint) defined in the consumer's package; emi just references it by name so the renderer's SQLValuer hook can take over.
	Address UserentityDtoAddress `json:"address" yaml:"address"`
}

func GetUserentityDtoPreferencesCliFlags(prefix string) []emigo.CliFlag {
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
func CastUserentityDtoPreferencesFromCli(c emigo.CliCastable) UserentityDtoPreferences {
	data := UserentityDtoPreferences{}
	if c.IsSet("theme") {
		data.Theme = c.String("theme")
	}
	if c.IsSet("locale") {
		data.Locale = c.String("locale")
	}
	return data
}

// The base class definition for preferences
type UserentityDtoPreferences struct {
	Theme  string `json:"theme" yaml:"theme"`
	Locale string `json:"locale" yaml:"locale"`
}

func GetUserentityDtoTagsCliFlags(prefix string) []emigo.CliFlag {
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
func CastUserentityDtoTagsFromCli(c emigo.CliCastable) UserentityDtoTags {
	data := UserentityDtoTags{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("value") {
		data.Value = c.String("value")
	}
	return data
}

// The base class definition for tags
type UserentityDtoTags struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

func GetUserentityDtoAddressCliFlags(prefix string) []emigo.CliFlag {
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
func CastUserentityDtoAddressFromCli(c emigo.CliCastable) UserentityDtoAddress {
	data := UserentityDtoAddress{}
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
type UserentityDtoAddress struct {
	Id       int64                  `json:"id" yaml:"id"`
	UserId   int64                  `json:"userId" yaml:"userId"`
	Street   string                 `json:"street" yaml:"street"`
	City     string                 `json:"city" yaml:"city"`
	Postcode emigo.Nullable[string] `json:"postcode" yaml:"postcode"`
}

func (x *UserentityDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
