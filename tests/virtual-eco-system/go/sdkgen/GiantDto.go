package unknownpackage
  // The base class definition for giantDto
type GiantDto struct {
		FirstName string `json:"firstName" yaml:"firstName"`
		FirstNameNullable emigo.Nullable[string] `json:"firstNameNullable" yaml:"firstNameNullable"`
		Array []GiantDtoArray `json:"array" yaml:"array"`
		ArrayNullable emi.Nullable[[]GiantDtoArrayNullable] `json:"arrayNullable" yaml:"arrayNullable"`
}
  // The base class definition for array
type GiantDtoArray struct {
		SubItem1 int `json:"subItem1" yaml:"subItem1"`
}
  // The base class definition for arrayNullable
type GiantDtoArrayNullable struct {
		SubItemNullable1 emigo.Nullable[int] `json:"subItemNullable1" yaml:"subItemNullable1"`
}