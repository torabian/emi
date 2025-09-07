package core

// Complex is a custom data type which developer can define, and would be used in field types.
// You need to define the custom ones here, so compiler knows from where has to import them
// based on different targets.
type EmiComplex struct {

	// Name of the complex class or struct, which will be instantiated.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description= Name of the complex class or struct, which will be instantiated."`

	// Location of the definition, for different languages to import
	Location string `yaml:"location,omitempty" json:"location,omitempty" jsonschema:"description= Location of the definition, for different languages to import"`

	// Different compilers can import the definition from different locations, for example 'js'
	Compiler string `yaml:"compiler,omitempty" json:"compiler,omitempty" jsonschema:"description= Different compilers can import the definition from different locations, for example 'js'"`
}
