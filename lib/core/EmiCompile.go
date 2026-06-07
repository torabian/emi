package core

// The emi module, can contain compiling options and flags into itself,
// so there is no need to pass options via flags.
type EmiCompile struct {
	Compiler string `yaml:"compiler,omitempty" json:"compiler,omitempty" jsonschema:"enum=go,enum=js,enum=kotlin,enum=swift,description=The compiler which will be picked."`

	// The flags to be passed to the compiler
	Flags map[string]string `yaml:"flags,omitempty" json:"flags,omitempty" jsonschema:"description=The flags to be passed to the compiler"`

	// The flags to be passed to the compiler
	Output string `yaml:"output,omitempty" json:"output,omitempty" jsonschema:"description=Output directory of the generated content"`

	// Compiler tags
	Tags []string `yaml:"tags,omitempty" json:"tags,omitempty" jsonschema:"description=Compiler tags"`

	// This section is commented at the moment. In airplane I have not enough mental capacity to implement
	// this one.
	// Custom selection of content while building the target, to include/exclude features such as actions and dtos
	// Selection EmiCompileIncludeMap `yaml:"selection,omitempty" json:"selection,omitempty" jsonschema:"description=Custom selection of content while building the target, to include/exclude features such as actions and dtos"`
}

// A mechanism to define include system
type EmiCompileIncludeMap struct {
	Actions EmiCompileIncludeCondition `yaml:"actions,omitempty" json:"actions,omitempty"`
	Dtos    EmiCompileIncludeCondition `yaml:"dtos,omitempty" json:"dtos,omitempty"`
}

type EmiCompileIncludeCondition struct {
	// Includes defines the list of actions or patterns to include.
	// Supports wildcard matching.
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty" jsonschema:"title=Includes,description=Included actions or patterns (supports wildcards)"`

	// Excludes defines the list of conditions or patterns to exclude,
	// overriding Includes when matched.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty" jsonschema:"title=Excludes,description=Excluded conditions or patterns"`
}
