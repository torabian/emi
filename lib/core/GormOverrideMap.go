package core

// Used in Module code generation to customized the generated code for gorm tags on Emi
// Data management fields such as workspace or user id. For example, you can add extra indexes on these
// fields.
type GormOverrideMap struct {

	// Override the workspace id configuration for gorm instead of default config. Useful for adding extra constraints or indexes.
	WorkspaceId string `yaml:"workspaceId,omitempty" json:"workspaceId,omitempty" jsonschema:"description=Override the workspace id configuration for gorm instead of default config. Useful for adding extra constraints or indexes."`

	// Override the user id configuration for gorm instead of default config. Useful for adding extra constraints or indexes.
	UserId string `yaml:"userId,omitempty" json:"userId,omitempty" jsonschema:"description=Override the user id configuration for gorm instead of default config. Useful for adding extra constraints or indexes."`
}
