package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity4Dto
type Entity4Dto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Note     string                 `json:"note" yaml:"note"`
}

func (x *Entity4Dto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity4DtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "note",
			Type: "string",
		},
	}
}
func CastEntity4DtoFromCli(c emigo.CliCastable) Entity4Dto {
	data := Entity4Dto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("note") {
		data.Note = c.String("note")
	}
	return data
}
