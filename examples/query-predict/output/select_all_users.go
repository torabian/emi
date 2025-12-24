package xxx

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
    "regexp"
)

const SelectAllUsersSQL = `select 
    name,
    email,
    id 
from users
where filter();`

type SelectAllUsersRow struct {
	Name string
	Email string
	Id string
}

type SelectAllUsersContext struct {
    Filter string
    Having string
    Restriction string
    Params      map[string]interface{}
    Placeholders      []any
}


func SelectAllUsersPrepreSql(ctx SelectAllUsersContext) (string, error) {
    replaceUseVal := func(sql string, values map[string]interface{}) string {
        re := regexp.MustCompile(`useval\(\s*['"]([^'"]+)['"]\s*\)`)
        return re.ReplaceAllStringFunc(sql, func(match string) string {
            if m := re.FindStringSubmatch(match); len(m) > 1 {
                if val, ok := values[m[1]]; ok {
                    // escape the value safely
                    switch v := val.(type) {
                    case string:
                        safe := strings.ReplaceAll(v, `'`, `''`)
                        return "'" + safe + "'"
                    default:
                        return fmt.Sprintf("%v", v)
                    }
                }
            }
            return match
        })
    }

    script := replaceUseVal(SelectAllUsersSQL, ctx.Params)
    filter := "1"
	if ctx.Filter != "" {
		filter = ctx.Filter
	}
	script = strings.ReplaceAll(script, "filter()", "(" +filter+ ")")

    restriction := "1"
    if ctx.Restriction != "" {
        restriction = ctx.Restriction
    }
    script = strings.ReplaceAll(script, "restriction()", "(" +restriction + ")")
    
    having := ""
    if ctx.Having != "" {
        having = ctx.Having
    }
    script = strings.ReplaceAll(script, "having()", having)

    return script, nil
}


func SelectAllUsers(db *sql.DB, ctx SelectAllUsersContext,) ([]SelectAllUsersRow, error) {
    script, err := SelectAllUsersPrepreSql(ctx)
    if err != nil {
		return nil, err
	}

    log.Default().Println(script)

	rows, err := db.Query(script, ctx.Placeholders...)
	if err != nil {
		return nil, fmt.Errorf("query SelectAllUsers failed: %w", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []SelectAllUsersRow
	for rows.Next() {
		var r SelectAllUsersRow

		scanArgs := make([]interface{}, len(cols))
		for i, col := range cols {
			switch col {
			case "name":
				scanArgs[i] = &r.Name
			case "email":
				scanArgs[i] = &r.Email
			case "id":
				scanArgs[i] = &r.Id
			default:
				var discard interface{}
				scanArgs[i] = &discard
			}
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}
