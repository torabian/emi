package external

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
)

// The base class definition for createSchemaVsqlParams
type CreateSchemaVsqlParams struct {
}

func (x *CreateSchemaVsqlParams) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetCreateSchemaVsqlParamsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastCreateSchemaVsqlParamsFromCli(c emigo.CliCastable) CreateSchemaVsqlParams {
	data := CreateSchemaVsqlParams{}
	return data
}

// CreateSchemaVsqlName is the name of the vsql query, useful for logging or routing.
const CreateSchemaVsqlName = "createSchema"

// CreateSchemaVsqlQuery is the raw SQL string for the createSchema vsql query.
const CreateSchemaVsqlQuery = `CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  email TEXT NOT NULL,
  first_name TEXT,
  last_name TEXT,
  user_role TEXT,
  is_active INTEGER,
  balance_cents INTEGER,
  rating REAL,
  street TEXT,
  city TEXT,
  money_amount_cents INTEGER,
  money_currency TEXT
);
`

// PrepareCreateSchemaVsql returns the query string and params for the createSchema vsql query,
// ready to be passed to a SQL driver of your choice.
func PrepareCreateSchemaVsql(params CreateSchemaVsqlParams) (query string, args interface{}) {
	return CreateSchemaVsqlQuery, params
}
