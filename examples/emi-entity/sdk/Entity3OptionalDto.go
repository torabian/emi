package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for entity3OptionalDto
type Entity3OptionalDto struct {
	UniqueId emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	Message  emigo.Nullable[string] `json:"message" yaml:"message"`
}

func (x *Entity3OptionalDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity3OptionalDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "unique-id",
			Type: "string?",
		},
		{
			Name: prefix + "message",
			Type: "string?",
		},
	}
}
func CastEntity3OptionalDtoFromCli(c emigo.CliCastable) Entity3OptionalDto {
	data := Entity3OptionalDto{}
	if c.IsSet("unique-id") {
		emigo.ParseNullable(c.String("unique-id"), &data.UniqueId)
	}
	if c.IsSet("message") {
		emigo.ParseNullable(c.String("message"), &data.Message)
	}
	return data
}
