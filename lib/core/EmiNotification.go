package core

// Events are definitions of a low level occurence across the application,
// for example an entity created - a user logged in - etc.
type EmiNotification struct {

	// Name of the notification which will be generated in golang and used as key to trigger or subscribe
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the notification which will be generated in golang and used as key to trigger or subscribe"`

	// Description of the event (developer visible only)
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the event (developer visible only)"`

	// Payload of the notification
	Payload *EmiActionBody `yaml:"payload,omitempty" json:"payload,omitempty" jsonschema:"description=Payload of the notification"`
}
