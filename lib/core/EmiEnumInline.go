package core

// Used in EmiField as the definition of enum items
type EmiEnumInline struct {
	// Enum key which will be used in golang generation and validation
	Key string `yaml:"k,omitempty" json:"k,omitempty" jsonschema:"description=Enum key which will be used in golang generation and validation"`

	// Description of the enum for developers. It's not translated or meant to be shown to end users.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the enum for developers. It's not translated or meant to be shown to end users."`
}
