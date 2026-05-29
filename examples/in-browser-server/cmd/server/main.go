//go:build !wasm

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/examples/in-browser-server/internal"
	"github.com/torabian/emi/examples/in-browser-server/internal/defs"
)

func main() {

	g := gin.Default()

	// Using implementation from internal package for gin
	defs.SubstringActionGin(g, internal.SubstringAction)

	g.Run(":9123")
}
