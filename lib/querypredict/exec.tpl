package artifacts

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
    "regexp"
)

const {{ .FuncName }}SQL = `{{ .SQL }}`

type {{ .FuncName }}Context struct {
    Filter string
    Having string
    Restriction string
    Params      map[string]interface{}

	// Native sql placeholder values such as where id = ?
	Placeholders []any
}

{{ template "prepareSql" . }}

func {{ .FuncName }}(db *sql.DB, ctx {{ .FuncName }}Context,{{ if .Params }}, {{ .Params }}{{ end }}) (sql.Result, error) {
    script, err := {{ .FuncName }}PrepreSql(ctx)
    if err != nil {
		return nil, err
	}


    log.Default().Println(script)

	res, err := db.Exec(script, ctx.Placeholders...)
	if err != nil {
		return res, fmt.Errorf("exec {{ .FuncName }} failed: %w", err)
	}

	return res, nil
}
