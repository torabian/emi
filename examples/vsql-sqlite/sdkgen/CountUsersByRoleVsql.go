package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for countUsersByRoleVsqlParams
type CountUsersByRoleVsqlParams struct {
	// When set, only rows with this role are counted; when unset, every row is counted.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
}

func (x *CountUsersByRoleVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCountUsersByRoleVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "role",
			Type:        "enum?",
			Description: "When set, only rows with this role are counted; when unset, every row is counted.",
		},
	}
}
func CastCountUsersByRoleVsqlParamsFromCli(c emigo.CliCastable) CountUsersByRoleVsqlParams {
	data := CountUsersByRoleVsqlParams{}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	return data
}

// CountUsersByRoleVsqlName is the name of the vsql query, useful for logging or routing.
const CountUsersByRoleVsqlName = "countUsersByRole"

// CountUsersByRoleVsqlQuery is the raw SQL string for the countUsersByRole vsql query.
const CountUsersByRoleVsqlQuery = `SELECT count(*) as total FROM users WHERE 1 = 1
{{ if .Role.IsSet }} AND user_role = {{ sql .Role }}{{ end }};
`

// PrepareCountUsersByRoleVsql returns the query string and params for the countUsersByRole vsql query,
// ready to be passed to a SQL driver of your choice.
func PrepareCountUsersByRoleVsql(params CountUsersByRoleVsqlParams) (query string, args interface{}) {
	return CountUsersByRoleVsqlQuery, params
}
