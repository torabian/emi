package unk
	import emigo "test/emi/emigo"
  // The base class definition for giantDto
type GiantDto struct {
		FirstName string `json:"firstName" yaml:"firstName"`
		FirstNameNullable emigo.Nullable[string] `json:"firstNameNullable" yaml:"firstNameNullable"`
		Array []GiantDtoArray `json:"array" yaml:"array"`
		ArrayNullable emigo.Nullable[[]GiantDtoArrayNullable] `json:"arrayNullable" yaml:"arrayNullable"`
		BooleanField bool `json:"booleanField" yaml:"booleanField"`
		BooleanFieldNullable emigo.Nullable[bool] `json:"booleanFieldNullable" yaml:"booleanFieldNullable"`
		CollectionItems []GiantDto `json:"collectionItems" yaml:"collectionItems"`
		CollectionItemsNullable emigo.Nullable[interface{}] `json:"collectionItemsNullable" yaml:"collectionItemsNullable"`
		DateObject interface{} `json:"dateObject" yaml:"dateObject"`
		SingleRefNullable emigo.Nullable[interface{}] `json:"singleRefNullable" yaml:"singleRefNullable"`
		Enumeration string `json:"enumeration" yaml:"enumeration"`
		EnumerationNullable emigo.Nullable[string] `json:"enumerationNullable" yaml:"enumerationNullable"`
		FloatingPoint32 float32 `json:"floatingPoint32" yaml:"floatingPoint32"`
		FloatingPoint32Nullable emigo.Nullable[float32] `json:"floatingPoint32Nullable" yaml:"floatingPoint32Nullable"`
		FloatingPoint64 float64 `json:"floatingPoint64" yaml:"floatingPoint64"`
		FloatingPoint64Nullable emigo.Nullable[float64] `json:"floatingPoint64Nullable" yaml:"floatingPoint64Nullable"`
		IntegerValue int `json:"integerValue" yaml:"integerValue"`
		Integer32ValueNullable emigo.Nullable[int32] `json:"integer32ValueNullable" yaml:"integer32ValueNullable"`
		Integer32Value int32 `json:"integer32Value" yaml:"integer32Value"`
		Integer64ValueNullable emigo.Nullable[int64] `json:"integer64ValueNullable" yaml:"integer64ValueNullable"`
		Integer64Value int64 `json:"integer64Value" yaml:"integer64Value"`
		MapValue interface{} `json:"mapValue" yaml:"mapValue"`
		MapValueNullable emigo.Nullable[interface{}] `json:"mapValueNullable" yaml:"mapValueNullable"`
		SliceValue []string `json:"sliceValue" yaml:"sliceValue"`
		SliceValueNullable emigo.Nullable[emigo.Nullable[[]string]] `json:"sliceValueNullable" yaml:"sliceValueNullable"`
		ObjectValue  GiantDtoObjectValue `json:"objectValue" yaml:"objectValue"`
}
  // The base class definition for array
type GiantDtoArray struct {
		SubItem1 int `json:"subItem1" yaml:"subItem1"`
}
  // The base class definition for arrayNullable
type GiantDtoArrayNullable struct {
		SubItemNullable1 emigo.Nullable[int] `json:"subItemNullable1" yaml:"subItemNullable1"`
}
  // The base class definition for objectValue
type GiantDtoObjectValue struct {
		InnerObject emigo.Nullable[GiantDtoObjectValueInnerObject] `json:"innerObject" yaml:"innerObject"`
}
  // The base class definition for innerObject
type GiantDtoObjectValueInnerObject struct {
		InnerObjText string `json:"innerObjText" yaml:"innerObjText"`
}