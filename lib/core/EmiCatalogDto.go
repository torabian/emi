package core

// In some cases, you can set emi: dto, and the rest of fields will become EmiDto.
// This file is only used to export the json schema
type EmiCatalogDto struct {

	// Custom imports appened by some macros
	Emi string `jsonschema:"description=Type of the emi content.;enum=dto" json:"emi" yaml:"emi"`

	EmiDto
}
