package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity2EntityUpdateInput
type Entity2EntityUpdateInput struct {
	Label emigo.Nullable[string] `json:"label" yaml:"label"`
}

func (x *Entity2EntityUpdateInput) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2EntityUpdateInputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "label",
			Type: "string?",
		},
	}
}
func CastEntity2EntityUpdateInputFromCli(c emigo.CliCastable) Entity2EntityUpdateInput {
	data := Entity2EntityUpdateInput{}
	if c.IsSet("label") {
		emigo.ParseNullable(c.String("label"), &data.Label)
	}
	return data
}
