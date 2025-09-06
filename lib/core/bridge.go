package core

// Package core provides types and utilities for defining actions that
// can be invoked both from WASM and Go. It includes flag definitions,
// base action structures, and collections of public API actions.

// FlagType represents the type of a command/action flag.
type FlagType string

const (
	// FlagString represents a string flag.
	FlagString FlagType = "string"
	// FlagBool represents a boolean flag.
	FlagBool FlagType = "bool"
	// FlagInt represents an integer flag.
	FlagInt FlagType = "int"
)

// FlagDef defines a single flag for an action, including its name,
// description, type, whether it is required, and an optional default value.
type FlagDef struct {
	Name     string      // Flag name, used in CLI or programmatic calls.
	Usage    string      // Description for the flag.
	Required bool        // Whether the flag must be provided.
	Type     FlagType    // Type of the flag: string, bool, int.
	Default  interface{} // Optional default value if not provided.
}

// BaseAction is embedded by specific action types (ActionText, ActionFile)
// and provides common metadata for actions.
type BaseAction struct {
	Name             string    // Action name.
	WasmFunctionName string    // Name of the corresponding WASM function.
	Description      string    // Description of the action.
	Flags            []FlagDef // List of flags accepted by the action.
}

// ActionText defines an action that returns a string result.
// Run is the function that executes the action logic.
type ActionText struct {
	BaseAction
	Run func(ctx MicroGenContext) (string, error)
}

// ActionFile defines an action that returns virtual files.
// Run is the function that executes the action logic.
type ActionFile struct {
	BaseAction
	Run func(ctx MicroGenContext) ([]VirtualFile, error)
}

// PublicAPIActions collects all actions exposed in the public API,
// separating them into text-based actions and file-based actions.
type PublicAPIActions struct {
	TextActions []ActionText // All text-returning actions.
	FileActions []ActionFile // All file-returning actions.
}
