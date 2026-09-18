package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for listUsersSortedVsqlParams
type ListUsersSortedVsqlParams struct {
	// Max rows to return.
	Limit int `json:"limit" yaml:"limit"`
	// Rows to skip before the first one returned.
	Offset int `json:"offset" yaml:"offset"`
	// Column to sort by. Deliberately an enum of the exact lowercase column names (not a display label like "Id") - `of:` constrains it to two known-safe values, so the query can splice {{ .OrderBy }} straight in as a raw SQL identifier (never quoted like a value) without risking injection through this field.
	OrderBy string `json:"orderBy" yaml:"orderBy"`
}

func (x *ListUsersSortedVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetListUsersSortedVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "limit",
			Type:        "int",
			Description: "Max rows to return.",
		},
		{
			Name:        prefix + "offset",
			Type:        "int",
			Description: "Rows to skip before the first one returned.",
		},
		{
			Name:        prefix + "order-by",
			Type:        "enum",
			Description: "Column to sort by. Deliberately an enum of the exact lowercase column names (not a display label like \"Id\") - `of:` constrains it to two known-safe values, so the query can splice {{ .OrderBy }} straight in as a raw SQL identifier (never quoted like a value) without risking injection through this field.",
		},
	}
}
func CastListUsersSortedVsqlParamsFromCli(c emigo.CliCastable) ListUsersSortedVsqlParams {
	data := ListUsersSortedVsqlParams{}
	if c.IsSet("limit") {
		data.Limit = int(c.Int64("limit"))
	}
	if c.IsSet("offset") {
		data.Offset = int(c.Int64("offset"))
	}
	if c.IsSet("order-by") {
		data.OrderBy = c.String("order-by")
	}
	return data
}

// ListUsersSortedVsqlName is the name of the vsql query, useful for logging or routing.
const ListUsersSortedVsqlName = "listUsersSorted"

// ListUsersSortedVsqlQuery is the raw SQL string for the listUsersSorted vsql query.
const ListUsersSortedVsqlQuery = `SELECT id, email, first_name, last_name FROM users
ORDER BY {{ .OrderBy }}
LIMIT {{ .Limit }} OFFSET {{ .Offset }};
`

// PrepareListUsersSortedVsql returns the query string and params for the listUsersSorted vsql query,
// ready to be passed to a SQL driver of your choice.
func PrepareListUsersSortedVsql(params ListUsersSortedVsqlParams) (query string, args interface{}) {
	return ListUsersSortedVsqlQuery, params
}
