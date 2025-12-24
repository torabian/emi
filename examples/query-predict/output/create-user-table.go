package xxx

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
    "regexp"
)

const CreateUserTableSQL = `CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type CreateUserTableContext struct {
    Filter string
    Having string
    Restriction string
    Params      map[string]interface{}

	// Native sql placeholder values such as where id = ?
	Placeholders []any
}


func CreateUserTablePrepreSql(ctx CreateUserTableContext) (string, error) {
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

    script := replaceUseVal(CreateUserTableSQL, ctx.Params)
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


func CreateUserTable(db *sql.DB, ctx CreateUserTableContext,) (sql.Result, error) {
    script, err := CreateUserTablePrepreSql(ctx)
    if err != nil {
		return nil, err
	}


    log.Default().Println(script)

	res, err := db.Exec(script, ctx.Placeholders...)
	if err != nil {
		return res, fmt.Errorf("exec CreateUserTable failed: %w", err)
	}

	return res, nil
}
