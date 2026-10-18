//go:build wasm

package emigo

import (
	"errors"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/urfave/cli/v3"
)

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

// SaveEnvFile is wasm's counterpart to Config.go's own (!wasm) SaveEnvFile. A browser sandbox
// has no filesystem to write a .env file to, so this is a no-op: generated Config.Save methods
// keep compiling, and the config simply stays in the process environment.
func SaveEnvFile(config interface{}, filename string) error {
	return nil
}

var errConfigSetWasm = errors.New("config values are set through the environment under wasm")

// configSetArg hands the one argument a `config <name> set <value>` command was given to parse.
// Config.go's versions fall back to an interactive prompt (promptui) when there is none, which a
// browser has no terminal for, so here a missing value is an error.
func configSetArg(c *cli.Command, apply func(read string) error) error {
	if c == nil || c.Args().Len() == 0 {
		return errConfigSetWasm
	}
	return apply(c.Args().First())
}

func ConfigSetBoolean(c *cli.Command, currentValue bool, setValue func(value bool)) error {
	return configSetArg(c, func(read string) error {
		switch read {
		case "true", "1", "yes":
			setValue(true)
		case "false", "0", "no":
			setValue(false)
		default:
			return errors.New("the value for boolean needs to be true, false, 0, 1, yes, no")
		}
		return nil
	})
}

func ConfigSetString(c *cli.Command, currentValue string, setValue func(value string)) error {
	return configSetArg(c, func(read string) error {
		setValue(read)
		return nil
	})
}

func ConfigSetInt64(c *cli.Command, currentValue int64, setValue func(value int64)) error {
	return configSetArg(c, func(read string) error {
		v, err := strconv.ParseInt(read, 10, 64)
		if err != nil {
			return err
		}
		setValue(v)
		return nil
	})
}

func ConfigSetInt(c *cli.Command, currentValue int, setValue func(value int)) error {
	return configSetArg(c, func(read string) error {
		v, err := strconv.Atoi(read)
		if err != nil {
			return err
		}
		setValue(v)
		return nil
	})
}

func ConfigSetFloat64(c *cli.Command, currentValue float64, setValue func(value float64)) error {
	return configSetArg(c, func(read string) error {
		v, err := strconv.ParseFloat(read, 64)
		if err != nil {
			return err
		}
		setValue(v)
		return nil
	})
}
