package core

// Represents a dto in an application. Can be used for variety of reasons,
// request response of an action, or even internally. Emi generates bunch of
// helpers for each dto, so it might make sense to define them in Emi instead
// of pure struct in golang.
type EmiDto struct {
	// Description about the purpose of the dto. It will be used in CLI and codegen documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the purpose of the dto. It will be used in CLI and codegen documentation."`

	// Name of the dto, in camel case, the rest of the code related to this dto is being generated based on this
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the dto in camel case the rest of the code related to this dto is being generated based on this"`

	// List of fields and body definitions of the dto
	// Implements lists the interfaces (Emi.Interfaces) this dto satisfies. Their fields
	// are included in the dto, and each target language generates the interface plus
	// the accessors that make the dto satisfy it - see EmiInterface.
	Implements []string `yaml:"implements,omitempty" json:"implements,omitempty" jsonschema:"description=Names of interfaces (see interfaces) this dto implements. Their fields are included and the generated code satisfies the interface."`

	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=List of fields and body definitions of the dto"`
}

func (x EmiDto) GetClassName() string {
	return ToUpper(x.Name) + "Dto"
}
