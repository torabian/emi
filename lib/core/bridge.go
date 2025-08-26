package core

// Here we put functions that can be used to define actions which are callable
// both from wasm, and go. this library should provide details to outside about flags,
// internal functions, and more.

type FlagType string

const (
	FlagString FlagType = "string"
	FlagBool   FlagType = "bool"
	FlagInt    FlagType = "int"
)

type FlagDef struct {
	Name     string
	Usage    string
	Required bool
	Type     FlagType
	Default  interface{}
}

// BaseAction is embedded by ActionText and ActionFile
type BaseAction struct {
	Name             string
	WasmFunctionName string
	Description      string
	Flags            []FlagDef
}

type ActionText struct {
	BaseAction
	Run func(ctx MicroGenContext) (string, error)
}

type ActionFile struct {
	BaseAction
	Run func(ctx MicroGenContext) ([]VirtualFile, error)
}

type PublicAPIActions struct {
	TextActions []ActionText
	FileActions []ActionFile
}
