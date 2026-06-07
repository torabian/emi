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
