package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity2UpdateDto
type Entity2UpdateDto struct {
	Label2 emigo.Nullable[string] `json:"label2" yaml:"label2"`
}

func (x *Entity2UpdateDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2UpdateDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label2",
			Type: "string?",
		},
	}
}
func CastEntity2UpdateDtoFromCli(c emigo.CliCastable) Entity2UpdateDto {
	data := Entity2UpdateDto{}
	if c.IsSet("label2") {
		emigo.ParseNullable(c.String("label2"), &data.Label2)
	}
	return data
}
