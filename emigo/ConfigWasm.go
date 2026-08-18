//go:build wasm

package emigo

import "github.com/kelseyhightower/envconfig"

// HandleEnvVars is wasm's counterpart to Config.go's own (!wasm) HandleEnvVars -
// same signature and the same envconfig.MustProcess call underneath, so any
// package's generated LoadConfiguration() (see lib/golang/go-config.go's
// GoConfigGenerate) works unmodified regardless of build target.
//
// There is no .env/.env.<ENV> file to read inside a browser sandbox (no
// filesystem, and godotenv itself doesn't build for js/wasm), so this skips
// straight to envconfig.MustProcess, which reads whatever env vars are
// already set via os.Setenv - Go's js/wasm runtime keeps a real in-memory env
// for that (os.Getenv/Setenv both work under wasm, they just start empty).
// The host page is expected to call os.Setenv (or its own bridge into it,
// e.g. reading a JS object handed to the wasm module before boot) for
// whatever keys it wants to seed before the wasm module's LoadConfiguration()
// runs - see cmd/fireback-wasm/main.go in the fireback repo for the pattern.
func HandleEnvVars(spec interface{}) {
	envconfig.MustProcess("", spec)
}
