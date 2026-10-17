package core

import "strings"

// Module3ConfigField represents a configuration field, typically a variable definition,
// which is converted into a Go struct. It provides functionality to read the value from
// YAML, environment variables, and CLI flags, with the option to save it as a .env file.
// This struct offers a type-safe way to use environment variables while documenting
// them in a YAML configuration file.
type EmiConfig struct {

	// Name is the identifier for the configuration field, used both in Go code and
	// the environment file. By default, the name will be converted to uppercase with
	// underscores to reference the environment variable, unless overridden by the 'env' field.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name is the identifier for the configuration field used both in Go code and the environment file. By default the name will be converted to uppercase with underscores to reference the environment variable unless overridden by the 'env' field."`

	// Type defines the data type for the environment variable. It supports standard Go types
	// such as string, bool, int64, and others, along with custom field types.
	// Ensure that the chosen type is supported.
	Type string `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"description=Type defines the data type for the environment variable. It supports standard Go types. such as string, bool, int64, and others, along with custom field types. Ensure that the chosen type is supported."`

	// Description explains the purpose of the configuration field. It can be helpful for developers
	// and also used in CLI for interactive configuration.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description explains the purpose of the configuration field. It can be helpful for developers and also used in CLI for interactive configuration."`

	// Default specifies the default value for the configuration field if it is not defined.
	Default string `yaml:"default,omitempty" json:"default,omitempty" jsonschema:"description=Default specifies the default value for the configuration field if it is not defined."`

	// Hint is the value that will be shown in auto complete or cli context as default value, but only to help user for understanding.
	Hint string `yaml:"hint,omitempty" json:"hint,omitempty" jsonschema:"description=Hint is the value that will be shown in auto complete or cli context as default value - but only to help user for understanding."`

	// Env allows you to override the default environment variable name, which is automatically
	// generated from the Name field. Use this field if you want to manually specify the environment variable name.
	Env string `yaml:"env,omitempty" json:"env,omitempty" jsonschema:"description=Env allows you to override the default environment variable name, which is automatically generated from the Name field. Use this field if you want to manually specify the environment variable name."`

	// Min and Max bound an int config. Only the C target's kconfig output reads them (Kconfig `range`).
	Min string `yaml:"min,omitempty" json:"min,omitempty" jsonschema:"description=Lower bound of an int config. Only used by the C kconfig output (Kconfig range). Needs max as well."`
	Max string `yaml:"max,omitempty" json:"max,omitempty" jsonschema:"description=Upper bound of an int config. Only used by the C kconfig output (Kconfig range). Needs min as well."`

	// Depends makes the config visible only while the expression holds. Only the C kconfig output reads it.
	Depends string `yaml:"depends,omitempty" json:"depends,omitempty" jsonschema:"description=Kconfig expression (&&, ||, !, parentheses) which has to hold for this config to be offered. Names of other configs of the module are replaced with their symbols; anything else (like an ESP-IDF symbol) is kept as written. Only used by the C kconfig output."`

	// Fields defines child configuration fields in case the current field represents an object
	// or an array of subfields. Note that support for nested fields may be limited.
	Fields []EmiConfig `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Fields defines child configuration fields in case the current field represents an object or an array of subfields. Note that support for nested fields may be limited."`
}

func (x *EmiConfig) DashedName() string {
	return strings.ReplaceAll(ToSnakeCase(x.Name), "_", "-")
}
