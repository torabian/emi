package external

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetYokeMessageDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "pitch",
			Type: "float64",
		},
		{
			Name: prefix + "role",
			Type: "float64",
		},
	}
}
func CastYokeMessageDtoFromCli(c emigo.CliCastable) YokeMessageDto {
	data := YokeMessageDto{}
	if c.IsSet("pitch") {
		data.Pitch = float64(c.Float64("pitch"))
	}
	if c.IsSet("role") {
		data.Role = float64(c.Float64("role"))
	}
	return data
}

// The base class definition for yokeMessageDto
type YokeMessageDto struct {
	Pitch float64 `json:"pitch" yaml:"pitch"`
	Role  float64 `json:"role" yaml:"role"`
}

func (x *YokeMessageDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
