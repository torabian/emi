package core

type EmiEntityPermissionRewrite struct {
	Replace string `yaml:"replace,omitempty" json:"replace,omitempty" jsonschema:"description=The value to be replaced"`
	With    string `yaml:"with,omitempty" json:"with,omitempty" jsonschema:"description=The value to be replaced"`
}
