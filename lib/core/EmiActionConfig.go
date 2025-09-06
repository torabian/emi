package core

type EmiActionConfig struct {
	Qs      []EmiField `yaml:"qs,omitempty" json:"qs,omitempty" jsonschema:"description=Typesafe query strings."`
	Headers []EmiField `yaml:"headers,omitempty" json:"headers,omitempty" jsonschema:"description=Typesafe headers."`
}
