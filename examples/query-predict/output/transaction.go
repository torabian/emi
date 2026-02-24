package queries

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
    "regexp"
)

const TransactionSQL = `--skip
BEGIN TRANSACTION;

INSERT INTO users(name, email) VALUES ('Alice', 'alice@example.com');
INSERT INTO users(name, email) VALUES ('Bob', 'bob@example.com');

COMMIT;`

type TransactionContext struct {
    Filter string
    Having string
    Restriction string
    Params      map[string]interface{}

	// Native sql placeholder values such as where id = ?
	Placeholders []any
}


func TransactionPrepreSql(ctx TransactionContext) (string, error) {
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

    script := replaceUseVal(TransactionSQL, ctx.Params)
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


func Transaction(db *sql.DB, ctx TransactionContext,) (sql.Result, error) {
    script, err := TransactionPrepreSql(ctx)
    if err != nil {
		return nil, err
	}

	res, err := db.Exec(script, ctx.Placeholders...)
	if err != nil {
		return res, fmt.Errorf("exec Transaction failed: %w", err)
	}

	return res, nil
}
