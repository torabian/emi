package core

// JsFnArgument represents a JS/TS function argument signature, e.g., "name?: string | null".
type JsFnArgument struct {
	Key string // Argument name
	Ts  string // TypeScript type declaration
	Js  string // JavaScript type/representation
}

// CompileJs returns the JavaScript representation of the argument.
func (x JsFnArgument) CompileJs() string {
	return x.Js
}

// CompileTs returns the TypeScript representation of the argument.
func (x JsFnArgument) CompileTs() string {
	return x.Ts
}

// NewJsArgument creates a new JsFnArgument. Optional helper if you want to extend processing in the future.
func NewJsArgument(dto JsFnArgument) JsFnArgument {
	return dto
}
