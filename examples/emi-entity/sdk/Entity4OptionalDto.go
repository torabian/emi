package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity4OptionalDto
type Entity4OptionalDto struct {
	Note emigo.Nullable[string] `json:"note" yaml:"note"`
}

func (x *Entity4OptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity4OptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "note",
			Type: "string?",
		},
	}
}
func CastEntity4OptionalDtoFromCli(c emigo.CliCastable) Entity4OptionalDto {
	data := Entity4OptionalDto{}
	if c.IsSet("note") {
		emigo.ParseNullable(c.String("note"), &data.Note)
	}
	return data
}
