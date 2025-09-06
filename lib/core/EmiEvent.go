package core

// Events are definitions of a low level occurence across the application,
// for example an entity created - a user logged in - etc.
type EmiEvent struct {

	// Name of the event which will be generated in golang and used as key to trigger or subscribe
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the event which will be generated in golang and used as key to trigger or subscribe"`

	// Description of the event (developer visible only)
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the event (developer visible only)"`

	// Payload of the event
	Payload *EmiActionBody `yaml:"payload,omitempty" json:"payload,omitempty" jsonschema:"description=Payload of the event"`

	// Security model of the event, which determines who can see it
	SecurityModel *SecurityModel `yaml:"security,omitempty" json:"security,omitempty" jsonschema:"description=Security model of the event, which determines who can see it"`

	// Mechanism to trigger a cache refresh on clients
	CacheKey string `yaml:"cacheKey,omitempty" json:"cacheKey,omitempty" jsonschema:"description=Mechanism to trigger a cache refresh on clients"`
}
