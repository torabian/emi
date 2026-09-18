package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for userSortDtoDto
type UserSortDtoDto struct {
	// Column to sort by. Deliberately an enum of the exact lowercase column names (not a display label like "Id") - `of:` constrains it to two known-safe values, so the query can splice {{ .OrderBy }} straight in as a raw SQL identifier (never quoted like a value) without risking injection through this field.
	OrderBy string `json:"orderBy" yaml:"orderBy"`
}

func (x *UserSortDtoDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetUserSortDtoDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "order-by",
			Type:        "enum",
			Description: "Column to sort by. Deliberately an enum of the exact lowercase column names (not a display label like \"Id\") - `of:` constrains it to two known-safe values, so the query can splice {{ .OrderBy }} straight in as a raw SQL identifier (never quoted like a value) without risking injection through this field.",
		},
	}
}
func CastUserSortDtoDtoFromCli(c emigo.CliCastable) UserSortDtoDto {
	data := UserSortDtoDto{}
	if c.IsSet("order-by") {
		data.OrderBy = c.String("order-by")
	}
	return data
}
