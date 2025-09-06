package core

// Represents a dto in an application. Can be used for variety of reasons,
// request response of an action, or even internally. Emi generates bunch of
// helpers for each dto, so it might make sense to define them in Emi instead
// of pure struct in golang.
type EmiDto struct {

	// Name of the dto, in camel case, the rest of the code related to this dto is being generated based on this
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the dto in camel case the rest of the code related to this dto is being generated based on this"`

	// List of fields and body definitions of the dto
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=List of fields and body definitions of the dto"`
}

func (x EmiDto) GetClassName() string {
	return ToUpper(x.Name) + "Dto"
}
