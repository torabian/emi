package core

// Task is a general function or similarly Emi Action, which has no results
// and could be run via Queue services or cronjobs
// Developer needs to implement the functionality manually, Emi only generates the func signature
// Tasks are only available internally and not exported via http or client sdks
type EmiTask struct {

	// List of triggers such as cronjobs which can make this task run automatically.
	Triggers []EmiTrigger `yaml:"triggers,omitempty" json:"triggers,omitempty" jsonschema:"description=List of triggers such as cronjobs which can make this task run automatically."`

	// Name of the task is general identifier and golang functions will be generated based on it.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the task is general identifier and golang functions will be generated based on it."`

	// Description of the task useful for developers and generated documentations.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the task useful for developers and generated documentations."`

	// Parameters that can be sent to this task. Since tasks are runnable in the golang as well
	// they can get parameters in go and cli if necessary. For cronjobs might make no sense.
	In *EmiActionBody `yaml:"in,omitempty" json:"in,omitempty" jsonschema:"description=Parameters that can be sent to this task. Since tasks are runnable in the golang as well they can get parameters in go and cli if necessary. For cronjobs might make no sense."`
}
