package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity2Dto
type Entity2Dto struct {
	Label2 string `json:"label2" yaml:"label2"`
}

func (x *Entity2Dto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2DtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label2",
			Type: "string",
		},
	}
}
func CastEntity2DtoFromCli(c emigo.CliCastable) Entity2Dto {
	data := Entity2Dto{}
	if c.IsSet("label2") {
		data.Label2 = c.String("label2")
	}
	return data
}
