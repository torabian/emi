package defs

import (
	"encoding"
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"math/big"
)

// The base class definition for addIntentArgs
type AddIntentArgs struct {
	// First operand.
	A   emigo.Nullable[float64] `json:"a" yaml:"a"`
	Ccc big.Int                 `json:"ccc" yaml:"ccc"`
	// Second operand.
	B float64 `json:"b" yaml:"b"`
}

func (x *AddIntentArgs) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetAddIntentArgsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "a",
			Type:        "float64?",
			Description: "First operand.",
		},
		{
			Name: prefix + "ccc",
			Type: "complex",
		},
		{
			Name:        prefix + "b",
			Type:        "float64",
			Description: "Second operand.",
		},
	}
}
func CastAddIntentArgsFromCli(c emigo.CliCastable) AddIntentArgs {
	data := AddIntentArgs{}
	if c.IsSet("a") {
		emigo.ParseNullable(c.String("a"), &data.A)
	}
	if c.IsSet("ccc") {
		if u, ok := any(&data.Ccc).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("ccc")))
		}
	}
	if c.IsSet("b") {
		data.B = float64(c.Float64("b"))
	}
	return data
}

// The base class definition for addIntentResult
type AddIntentResult struct {
	Sum float64 `json:"sum" yaml:"sum"`
}

func (x *AddIntentResult) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetAddIntentResultCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sum",
			Type: "float64",
		},
	}
}
func CastAddIntentResultFromCli(c emigo.CliCastable) AddIntentResult {
	data := AddIntentResult{}
	if c.IsSet("sum") {
		data.Sum = float64(c.Float64("sum"))
	}
	return data
}

// The base class definition for isPrimeIntentArgs
type IsPrimeIntentArgs struct {
	// The integer to test.
	N int64 `json:"n" yaml:"n"`
}

func (x *IsPrimeIntentArgs) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetIsPrimeIntentArgsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "n",
			Type:        "int64",
			Description: "The integer to test.",
		},
	}
}
func CastIsPrimeIntentArgsFromCli(c emigo.CliCastable) IsPrimeIntentArgs {
	data := IsPrimeIntentArgs{}
	if c.IsSet("n") {
		data.N = int64(c.Int64("n"))
	}
	return data
}

// The base class definition for isPrimeIntentResult
type IsPrimeIntentResult struct {
	Prime bool `json:"prime" yaml:"prime"`
}

func (x *IsPrimeIntentResult) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetIsPrimeIntentResultCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "prime",
			Type: "bool",
		},
	}
}
func CastIsPrimeIntentResultFromCli(c emigo.CliCastable) IsPrimeIntentResult {
	data := IsPrimeIntentResult{}
	if c.IsSet("prime") {
		data.Prime = bool(c.Bool("prime"))
	}
	return data
}

/**
* Model-facing tool signatures, generated from this module's intents list. See
* EmiIntent in lib/core for the source definitions - each one below either declared its
* own in/out inline, or was derived from (or auto-created by "intent: true" on) an
* existing action.
 */
// AddIntentMeta describes the "add" MCP tool intent - name,
// title, description and MCP behavior-hint annotations - so a server can build a
// concrete tool registration from it without redeclaring any of this by hand.
// Input shape: AddIntentArgs.
// Output shape: AddIntentResult.
func AddIntentMeta() struct {
	Name            string
	Title           string
	Description     string
	ReadOnlyHint    bool
	DestructiveHint bool
	IdempotentHint  bool
	OpenWorldHint   bool
	Permission      string
} {
	return struct {
		Name            string
		Title           string
		Description     string
		ReadOnlyHint    bool
		DestructiveHint bool
		IdempotentHint  bool
		OpenWorldHint   bool
		Permission      string
	}{
		Name:            "add",
		Title:           "Add two numbers",
		Description:     "Adds two numbers together and returns their sum.",
		ReadOnlyHint:    true,
		DestructiveHint: false,
		IdempotentHint:  true,
		OpenWorldHint:   false,
		Permission:      "",
	}
}

// IsPrimeIntentMeta describes the "isPrime" MCP tool intent - name,
// title, description and MCP behavior-hint annotations - so a server can build a
// concrete tool registration from it without redeclaring any of this by hand.
// Input shape: IsPrimeIntentArgs.
// Output shape: IsPrimeIntentResult.
func IsPrimeIntentMeta() struct {
	Name            string
	Title           string
	Description     string
	ReadOnlyHint    bool
	DestructiveHint bool
	IdempotentHint  bool
	OpenWorldHint   bool
	Permission      string
} {
	return struct {
		Name            string
		Title           string
		Description     string
		ReadOnlyHint    bool
		DestructiveHint bool
		IdempotentHint  bool
		OpenWorldHint   bool
		Permission      string
	}{
		Name:            "isPrime",
		Title:           "Check if a number is prime",
		Description:     "Reports whether a given integer is a prime number.",
		ReadOnlyHint:    true,
		DestructiveHint: false,
		IdempotentHint:  true,
		OpenWorldHint:   false,
		Permission:      "",
	}
}

// IntentNames lists every intent's tool name this module contributes, for a server to
// enumerate/register (e.g. iterating this to look up each "<GoName>IntentMeta").
func IntentNames() []string {
	return []string{
		"add",
		"isPrime",
	}
}
