package core

// Used to adjust the features generated for each entity.
type EmiEntityFeatures struct {

	// Adds a CLI task to automatically generate mocks.
	// Disable this if it is not relevant for the feature
	// or if validations are required, making a custom mock necessary.
	// IMPORTANT: This is a pointer because it is enabled by default.
	Mock *bool `yaml:"mock,omitempty" json:"mock,omitempty" jsonschema:"Adds a CLI task to automatically generate mocks Disable this if it is not relevant for the feature or if validations are required making a custom mock necessary IMPORTANT This is a pointer because it is enabled by default"`

	// Enables embedded mock files for an entity.
	// Disables the 'msync' and 'mlist' commands.
	MSync *bool `yaml:"msync,omitempty" json:"msync,omitempty" jsonschema:"Enables embedded mock files for an entity disabling the 'msync' and 'mlist' commands"`
}

// Checks if codegen needs to print a default mocking tools for the entity
func (x EmiEntityFeatures) HasMockAction() bool {

	if x.Mock != nil && !*x.Mock {
		return false
	}

	return true
}

// Checks id Emi definition enabled or disabled the feature
func (x EmiEntityFeatures) HasMsyncActions() bool {

	if x.MSync != nil && !*x.MSync {
		return false
	}

	return true
}
