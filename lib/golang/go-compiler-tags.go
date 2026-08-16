package golang

import "github.com/torabian/emi/lib/core"

// List of all compiler tags golang supports. Add them here before using them.
const SkipWasmGin core.CTag = "skip-wasm-gin"
const SkipCli core.CTag = "skip-cli"
const SplitCli core.CTag = "split-cli"
const SkipGin core.CTag = "skip-gin"
const SkipHttp core.CTag = "skip-http"
const SplitHttp core.CTag = "split-http"
const SplitGin core.CTag = "split-gin"
const NoClient core.CTag = "no-client"

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: SkipWasmGin, Description: "Wrap the Gin route file in a `//go:build !wasm` constraint, excluding it from wasm builds"},
	{Tag: SkipCli, Description: "Skip generating CLI (urfave-style) route/flag helpers"},
	{Tag: SplitCli, Description: "Emit CLI helpers into their own *Cli.go file instead of appending them to the main file"},
	{Tag: SkipGin, Description: "Skip generating Gin route helpers"},
	{Tag: SkipHttp, Description: "Skip generating the plain net/http (framework-free) route helpers"},
	{Tag: SplitHttp, Description: "Emit net/http route helpers into their own file instead of appending them to the main file"},
	{Tag: SplitGin, Description: "Emit Gin route helpers into their own *Gin.go file instead of appending them to the main file"},
	{Tag: NoClient, Description: "Skip generating the Go HTTP client (Call/ClientBuildRequest/... helpers) for consuming the action"},
}
