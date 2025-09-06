package core

// Permission is an access key to limit the usages of a feature.
type EmiPermission struct {
	// Name of the permission which will be used in golang and external ui
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the permission which will be used in golang and external ui"`

	// Key of the permission, separated with dots such as root.feature.action
	Key string `yaml:"key,omitempty" json:"key,omitempty" jsonschema:"description=Key of the permission, separated with dots such as root.feature.action"`

	// Description of the permission for developers or users. Not translated at this moment.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the permission for developers or users. Not translated at this moment."`
}
