package core

// Some emi files can be only dto, only entity, only action, or othey types.
// For this purpose, first we convert them into EmiCatalog,
// and then based on the type, we will use the sub compiler.
type EmiCatalog struct {
	Emi string `jsonschema:"description=Type of the emi content.;enum=dto" json:"emi" yaml:"emi"`
}
