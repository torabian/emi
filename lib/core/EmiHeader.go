package core

// Represent a header in the general web requests
type EmiHeader struct {
	// Name of the header accessible on programming languages.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the header accessible on programming languages"`

	// Value type, which will be parsed or read. All will be cast from the string value of header
	Type string `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=string,enum=int64,enum=float64,enum=bool,description=Value type, which will be parsed or read. All will be cast from the string value of header"`

	// Description of the value and reason usually to developers, http specs and cli
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the value and reason usually to developers, http specs and cli"`
}
