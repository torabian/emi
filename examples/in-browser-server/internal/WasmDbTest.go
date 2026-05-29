//go:build js && wasm

package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"syscall/js"
)

func DBTest() {
	queryFunc := js.Global().Get("queryDatabase")

	if queryFunc.IsUndefined() {
		println("queryDatabase is not defined")
		return
	}

	var conn Database = NewWasmDatabase(queryFunc)
	// Or if you wanted an original connection
	// conn, err := pgx.Connect(context.Background(), "dsn")
	// if err != nil {
	// 	panic(err)
	// }

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
	}
	fmt.Printf("[wasm-mock] dsn (ignored): %s\n", dsn)

	ctx := context.Background()

	if _, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS dashboard_wasm_test (
			id    SERIAL PRIMARY KEY,
			label TEXT NOT NULL
		)
	`); err != nil {
		log.Fatalf("create table: %v", err)
	}

	if _, err := conn.Exec(ctx,
		`INSERT INTO dashboard_wasm_test (label) VALUES ($1)`,
		"hello from dashboard-wasm",
	); err != nil {
		log.Fatalf("insert: %v", err)
	}

	rows, err := conn.Query(ctx, `SELECT id, label FROM dashboard_wasm_test ORDER BY id`)
	if err != nil {
		log.Fatalf("query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var label string
		if err := rows.Scan(&id, &label); err != nil {
			log.Fatalf("scan: %v", err)
		}
		fmt.Printf("row: id=%d label=%q\n", id, label)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("rows: %v", err)
	}
}
