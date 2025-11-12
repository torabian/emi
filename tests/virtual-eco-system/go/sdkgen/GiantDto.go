package unknownpackage

import "github.com/torabian/emi/emigo"

// The base class definition for giantDto
type GiantDto struct {
	FirstName         string                                `json:"firstName" yaml:"firstName"`
	FirstNameNullable emigo.Nullable[string]                `json:"firstNameNullable" yaml:"firstNameNullable"`
	Array             []GiantDtoArray                       `json:"array" yaml:"array"`
	ArrayNullable     emi.Nullable[[]GiantDtoArrayNullable] `json:"arrayNullable" yaml:"arrayNullable"`
}
