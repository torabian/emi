package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for searchUsersByFilterVsqlParams
type SearchUsersByFilterVsqlParams struct {
	// Exact email match. Unset (not merely empty) means no email filter is applied at all.
	Email emigo.Nullable[string] `json:"email" yaml:"email"`
	// Exact role match. Unset means no role filter.
	Role emigo.Nullable[string] `json:"role" yaml:"role"`
}

func (x *SearchUsersByFilterVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetSearchUsersByFilterVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
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
func CastSearchUsersByFilterVsqlParamsFromCli(c emigo.CliCastable) SearchUsersByFilterVsqlParams {
	data := SearchUsersByFilterVsqlParams{}
	if c.IsSet("email") {
		emigo.ParseNullable(c.String("email"), &data.Email)
	}
	if c.IsSet("role") {
		emigo.ParseNullable(c.String("role"), &data.Role)
	}
	return data
}

// SearchUsersByFilterVsqlName is the name of the vsql query, useful for logging or routing.
const SearchUsersByFilterVsqlName = "searchUsersByFilter"

// SearchUsersByFilterVsqlQuery is the raw SQL string for the searchUsersByFilter vsql query.
const SearchUsersByFilterVsqlQuery = `SELECT id, email FROM users WHERE 1 = 1
{{ if .Email.IsSet }} AND email = {{ sql .Email }}{{ end }}
{{ if .Role.IsSet }} AND user_role = {{ sql .Role }}{{ end }};
`

// PrepareSearchUsersByFilterVsql returns the query string and params for the searchUsersByFilter vsql query,
// ready to be passed to a SQL driver of your choice.
func PrepareSearchUsersByFilterVsql(params SearchUsersByFilterVsqlParams) (query string, args interface{}) {
	return SearchUsersByFilterVsqlQuery, params
}
