//go:build js && wasm

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"syscall/js"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/examples/in-browser-server/internal"
	"github.com/torabian/emi/examples/in-browser-server/internal/defs"
)

func main() {
	fmt.Println("Wasm main started")

	// Create the database connection first. queryDatabase is injected by
	// browser/database-bridge.js and is backed by an in-browser pglite
	// instance. Everything below talks to it through the Database interface.
	queryFunc := js.Global().Get("queryDatabase")
	if queryFunc.IsUndefined() {
		log.Fatal("queryDatabase is not defined — is database-bridge.js loaded?")
	}
	var conn internal.Database = internal.NewWasmDatabase(queryFunc)

	// User CRUD module shares the same connection. Make sure its table exists
	// before any request can land.
	users := &internal.UserModule{DB: conn}
	if err := users.EnsureSchema(context.Background()); err != nil {
		log.Fatalf("ensure users schema: %v", err)
	}

	// Build a standard net/http router and register the same typed handlers the
	// real server uses. The implementations are shared verbatim across both
	// transports.
	mux := http.NewServeMux()
	defs.SubstringActionHttp(mux, internal.SubstringAction)
	defs.CreateUserActionHttp(mux, users.Create)
	defs.ListUsersActionHttp(mux, users.List)
	defs.DeleteUserActionHttp(mux, users.Delete)

	emigo.LiftWasmServer(mux, nil)

	// Keep the Go runtime alive so the exposed callback stays callable.
	select {}
}
