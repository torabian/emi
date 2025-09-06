package core

// Contains extra configuration for each rpc which will be generated on Emi
// this is an attempt to reduce the need for custom actions, by giving you some control over
// the code might be generated.
type EmiEntityActionConfig struct {
	Query EmiActionConfig `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Modify the query rpc code."`
}
