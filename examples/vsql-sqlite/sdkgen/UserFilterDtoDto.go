package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for userFilterDtoDto
type UserFilterDtoDto struct {
	// Exact email match. Unset (not merely empty) means no email filter is applied at all.
	Email emigo.Nullable[string] `json:"email" yaml:"email"`
	// Exact role match. Unset means no role filter.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
}

func (x *UserFilterDtoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserFilterDtoDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "email",
			Type:        "string?",
			Description: "Exact email match. Unset (not merely empty) means no email filter is applied at all.",
		},
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "Exact role match. Unset means no role filter.",
		},
	}
}
func CastUserFilterDtoDtoFromCli(c emigo.CliCastable) UserFilterDtoDto {
	data := UserFilterDtoDto{}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	return data
}
