package core

type EmiDataFields struct {

	// Essential is a set of the fields which Emi uses to give userId and workspaceId
	Essentials bool `yaml:"essentials,omitempty" json:"essentials,omitempty" jsonschema:"default=true,description=Essential is a set of the fields which Emi uses to give userId and workspaceId"`

	// Adds a int primary key auto increment
	PrimaryId bool `yaml:"primaryId,omitempty" json:"primaryId,omitempty" jsonschema:"default=true,description=Adds a int primary key auto increment"`

	// adds created - updated - delete as nano seconds to the database
	NumericTimestamp bool `yaml:"numericTimestamp,omitempty" json:"numericTimestamp,omitempty" jsonschema:"default=true,description=adds created - updated - delete as nano seconds to the database"`

	// adds created - updated - deleted fields as timestamps
	DateTimestamp bool `yaml:"dateTimestamp,omitempty" json:"dateTimestamp,omitempty" jsonschema:"default=false,description=adds created - updated - deleted fields as timestamps"`
}
