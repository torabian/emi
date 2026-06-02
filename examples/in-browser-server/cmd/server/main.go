//go:build !wasm

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/in-browser-server/internal"
	"github.com/torabian/emi/examples/in-browser-server/internal/defs"
)

func main() {

	g := gin.Default()

	// We need to actually connect to posgres.
	// var conn internal.Database

	// // User CRUD module shares the same connection. Make sure its table exists
	// // before any request can land.
	// users := &internal.UserModule{DB: conn}
	// if err := users.EnsureSchema(context.Background()); err != nil {
	// 	log.Fatalf("ensure users schema: %v", err)
	// }

	// Using implementation from internal package for gin
	defs.SubstringActionGin(g, internal.SubstringAction)
	// defs.CreateUserActionGin(g, users.Create)
	// defs.ListUsersActionGin(g, users.List)
	// defs.DeleteUserActionGin(g, users.Delete)

	defs.ChatActionGin(g, internal.ChatHandler)

	g.Run(":9123")
}
