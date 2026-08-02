package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity3Dto
type Entity3Dto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Message  string                 `json:"message" yaml:"message"`
}

func (x *Entity3Dto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity3DtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "message",
			Type: "string",
		},
	}
}
func CastEntity3DtoFromCli(c emigo.CliCastable) Entity3Dto {
	data := Entity3Dto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	return data
}
