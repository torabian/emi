package core

// Used in EmiField as the definition of enum items
type EmiEnum struct {

	// Name of the enum, in camel case, the rest of the code related to this enum is being generated based on this
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the enum in camel case the rest of the code related to this enum is being generated based on this"`

	// Define the enum fields as an array, check EmiEnumInline struct for more details
	Fields []EmiEnumInline `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Define the enum fields as an array, check EmiEnumInline struct for more details"`
}

func (x EmiEnum) GetName() string {
	return ToUpper(x.Name)
}
