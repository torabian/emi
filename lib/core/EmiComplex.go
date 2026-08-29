package core

// Complex is a custom data type which developer can define, and would be used in field types.
// You need to define the custom ones here, so compiler knows from where has to import them
// based on different targets.
type EmiComplex struct {

	// Name of the complex class or struct, which will be instantiated.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description= Name of the complex class or struct, which will be instantiated."`

	// Namespace or package name for example in golang
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty" jsonschema:"description= Namespace or package name for example in golang"`

	// Location of the definition, for different languages to import
	Location string `yaml:"location,omitempty" json:"location,omitempty" jsonschema:"description= Location of the definition, for different languages to import"`

	// Different compilers can import the definition from different locations, for example 'js'
	Compiler string `yaml:"compiler,omitempty" json:"compiler,omitempty" jsonschema:"enum=go,enum=ts,description=Different compilers can import the definition from different locations, for example 'js'"`

	// Include is a relative or absolute path to another emi yaml file. When set, the
	// rest of this entry (Name/Namespace/Location/Compiler) is ignored, and this entry
	// is replaced by every complex declared in the referenced file's own `complexes`
	// list - so a module can share one set of complex definitions instead of repeating
	// them. A relative path is resolved against the directory of the file that
	// declares it (so an included file's own Include entries resolve relative to
	// itself, not to the module that first pulled it in), and is read again indefinitely
	// if the referenced file also sets Include, all the way down. A file that
	// (directly or transitively) tries to include itself again is a circular
	// dependency - resolution stops with an error instead of looping forever.
	Include string `yaml:"include,omitempty" json:"include,omitempty" jsonschema:"description=Relative or absolute path to another emi yaml file whose 'complexes' list should be included here in place of this entry. Resolved recursively; circular includes are rejected."`
}
