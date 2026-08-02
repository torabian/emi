package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity2OptionalDto
type Entity2OptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Label2   emigo.Nullable[string] `json:"label2" yaml:"label2"`
}

func (x *Entity2OptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2OptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "label2",
			Type: "string?",
		},
	}
}
func CastEntity2OptionalDtoFromCli(c emigo.CliCastable) Entity2OptionalDto {
	data := Entity2OptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("label2") {
		emigo.ParseNullable(c.String("label2"), &data.Label2)
	}
	return data
}
