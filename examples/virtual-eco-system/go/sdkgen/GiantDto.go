package external

import "encoding/json"
import emigo "test/emi/emigo"

func GetGiantDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "first-name",
			Type: "string",
		},
		{
			Name: prefix + "first-name-nullable",
			Type: "string?",
		},
		{
			Name: prefix + "array",
			Type: "array",
		},
		{
			Name: prefix + "array-nullable",
			Type: "array?",
		},
		{
			Name: prefix + "boolean-field",
			Type: "bool",
		},
		{
			Name: prefix + "boolean-field-nullable",
			Type: "bool?",
		},
		{
			Name: prefix + "collection-items",
			Type: "collection",
		},
		{
			Name: prefix + "collection-items-nullable",
			Type: "collection?",
		},
		{
			Name: prefix + "single-ref-nullable",
			Type: "one?",
		},
		{
			Name: prefix + "enumeration",
			Type: "enum",
		},
		{
			Name: prefix + "enumeration-nullable",
			Type: "enum?",
		},
		{
			Name: prefix + "floating-point32",
			Type: "float32",
		},
		{
			Name: prefix + "floating-point32-nullable",
			Type: "float32?",
		},
		{
			Name: prefix + "floating-point64",
			Type: "float64",
		},
		{
			Name: prefix + "floating-point64-nullable",
			Type: "float64?",
		},
		{
			Name: prefix + "integer-value",
			Type: "int",
		},
		{
			Name: prefix + "integer32-value-nullable",
			Type: "int32?",
		},
		{
			Name: prefix + "integer32-value",
			Type: "int32",
		},
		{
			Name: prefix + "integer64-value-nullable",
			Type: "int64?",
		},
		{
			Name: prefix + "integer64-value",
			Type: "int64",
		},
		{
			Name: prefix + "map-value",
			Type: "map",
		},
		{
			Name: prefix + "map-value-nullable",
			Type: "map?",
		},
		{
			Name: prefix + "slice-value",
			Type: "slice",
		},
		{
			Name: prefix + "slice-value-nullable",
			Type: "slice?",
		},
		{
			Name:     prefix + "object-value",
			Type:     "object",
			Children: GetGiantDtoObjectValueCliFlags("object-value-"),
		},
	}
}
func CastGiantDtoFromCli(c emigo.CliCastable) GiantDto {
	data := GiantDto{}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("first-name-nullable") {
		emigo.ParseNullable(c.String("first-name-nullable"), &data.FirstNameNullable)
	}
	if c.IsSet("array") {
		data.Array = emigo.CapturePossibleArray(CastGiantDtoArrayFromCli, "array", c)
	}
	if c.IsSet("array-nullable") {
		emigo.ParseNullable(c.String("array-nullable"), &data.ArrayNullable)
	}
	if c.IsSet("boolean-field") {
		data.BooleanField = bool(c.Bool("boolean-field"))
	}
	if c.IsSet("boolean-field-nullable") {
		emigo.ParseNullable(c.String("boolean-field-nullable"), &data.BooleanFieldNullable)
	}
	if c.IsSet("collection-items-nullable") {
		emigo.ParseNullable(c.String("collection-items-nullable"), &data.CollectionItemsNullable)
	}
	if c.IsSet("single-ref-nullable") {
		emigo.ParseNullable(c.String("single-ref-nullable"), &data.SingleRefNullable)
	}
	if c.IsSet("enumeration-nullable") {
		emigo.ParseNullable(c.String("enumeration-nullable"), &data.EnumerationNullable)
	}
	if c.IsSet("floating-point32") {
		data.FloatingPoint32 = float32(c.Float64("floating-point32"))
	}
	if c.IsSet("floating-point32-nullable") {
		emigo.ParseNullable(c.String("floating-point32-nullable"), &data.FloatingPoint32Nullable)
	}
	if c.IsSet("floating-point64") {
		data.FloatingPoint64 = float64(c.Float64("floating-point64"))
	}
	if c.IsSet("floating-point64-nullable") {
		emigo.ParseNullable(c.String("floating-point64-nullable"), &data.FloatingPoint64Nullable)
	}
	if c.IsSet("integer-value") {
		data.IntegerValue = int(c.Int64("integer-value"))
	}
	if c.IsSet("integer32-value-nullable") {
		emigo.ParseNullable(c.String("integer32-value-nullable"), &data.Integer32ValueNullable)
	}
	if c.IsSet("integer32-value") {
		data.Integer32Value = int32(c.Int64("integer32-value"))
	}
	if c.IsSet("integer64-value-nullable") {
		emigo.ParseNullable(c.String("integer64-value-nullable"), &data.Integer64ValueNullable)
	}
	if c.IsSet("integer64-value") {
		data.Integer64Value = int64(c.Int64("integer64-value"))
	}
	if c.IsSet("map-value-nullable") {
		emigo.ParseNullable(c.String("map-value-nullable"), &data.MapValueNullable)
	}
	if c.IsSet("slice-value") {
		emigo.InflatePossibleSlice(c.String("slice-value"), &data.SliceValue)
	}
	if c.IsSet("slice-value-nullable") {
		emigo.ParseNullable(c.String("slice-value-nullable"), &data.SliceValueNullable)
	}
	if c.IsSet("object-value") {
		data.ObjectValue = CastGiantDtoObjectValueFromCli(c)
	}
	return data
}

// The base class definition for giantDto
type GiantDto struct {
	FirstName               string                                  `json:"firstName" yaml:"firstName"`
	FirstNameNullable       emigo.Nullable[string]                  `json:"firstNameNullable" yaml:"firstNameNullable"`
	Array                   []GiantDtoArray                         `json:"array" yaml:"array"`
	ArrayNullable           emigo.Nullable[[]GiantDtoArrayNullable] `json:"arrayNullable" yaml:"arrayNullable"`
	BooleanField            bool                                    `json:"booleanField" yaml:"booleanField"`
	BooleanFieldNullable    emigo.Nullable[bool]                    `json:"booleanFieldNullable" yaml:"booleanFieldNullable"`
	CollectionItems         []GiantDto                              `json:"collectionItems" yaml:"collectionItems"`
	CollectionItemsNullable emigo.Nullable[[]GiantDto]              `json:"collectionItemsNullable" yaml:"collectionItemsNullable"`
	SingleRefNullable       emigo.Nullable[GiantDto]                `json:"singleRefNullable" yaml:"singleRefNullable"`
	Enumeration             string                                  `json:"enumeration" yaml:"enumeration"`
	EnumerationNullable     emigo.Nullable[string]                  `json:"enumerationNullable" yaml:"enumerationNullable"`
	FloatingPoint32         float32                                 `json:"floatingPoint32" yaml:"floatingPoint32"`
	FloatingPoint32Nullable emigo.Nullable[float32]                 `json:"floatingPoint32Nullable" yaml:"floatingPoint32Nullable"`
	FloatingPoint64         float64                                 `json:"floatingPoint64" yaml:"floatingPoint64"`
	FloatingPoint64Nullable emigo.Nullable[float64]                 `json:"floatingPoint64Nullable" yaml:"floatingPoint64Nullable"`
	IntegerValue            int                                     `json:"integerValue" yaml:"integerValue"`
	Integer32ValueNullable  emigo.Nullable[int32]                   `json:"integer32ValueNullable" yaml:"integer32ValueNullable"`
	Integer32Value          int32                                   `json:"integer32Value" yaml:"integer32Value"`
	Integer64ValueNullable  emigo.Nullable[int64]                   `json:"integer64ValueNullable" yaml:"integer64ValueNullable"`
	Integer64Value          int64                                   `json:"integer64Value" yaml:"integer64Value"`
	MapValue                map[any]any                             `json:"mapValue" yaml:"mapValue"`
	MapValueNullable        emigo.Nullable[map[any]any]             `json:"mapValueNullable" yaml:"mapValueNullable"`
	SliceValue              []string                                `json:"sliceValue" yaml:"sliceValue"`
	SliceValueNullable      emigo.Nullable[[]string]                `json:"sliceValueNullable" yaml:"sliceValueNullable"`
	ObjectValue             GiantDtoObjectValue                     `json:"objectValue" yaml:"objectValue"`
}

func GetGiantDtoArrayCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sub-item1",
			Type: "int",
		},
	}
}
func CastGiantDtoArrayFromCli(c emigo.CliCastable) GiantDtoArray {
	data := GiantDtoArray{}
	if c.IsSet("sub-item1") {
		data.SubItem1 = int(c.Int64("sub-item1"))
	}
	return data
}

// The base class definition for array
type GiantDtoArray struct {
	SubItem1 int `json:"subItem1" yaml:"subItem1"`
}

func GetGiantDtoArrayNullableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sub-item-nullable1",
			Type: "int?",
		},
	}
}
func CastGiantDtoArrayNullableFromCli(c emigo.CliCastable) GiantDtoArrayNullable {
	data := GiantDtoArrayNullable{}
	if c.IsSet("sub-item-nullable1") {
		emigo.ParseNullable(c.String("sub-item-nullable1"), &data.SubItemNullable1)
	}
	return data
}

// The base class definition for arrayNullable
type GiantDtoArrayNullable struct {
	SubItemNullable1 emigo.Nullable[int] `json:"subItemNullable1" yaml:"subItemNullable1"`
}

func GetGiantDtoObjectValueCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "inner-object",
			Type: "object?",
		},
	}
}
func CastGiantDtoObjectValueFromCli(c emigo.CliCastable) GiantDtoObjectValue {
	data := GiantDtoObjectValue{}
	if c.IsSet("inner-object") {
		emigo.ParseNullable(c.String("inner-object"), &data.InnerObject)
	}
	return data
}

// The base class definition for objectValue
type GiantDtoObjectValue struct {
	InnerObject emigo.Nullable[GiantDtoObjectValueInnerObject] `json:"innerObject" yaml:"innerObject"`
}

func GetGiantDtoObjectValueInnerObjectCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "inner-obj-text",
			Type: "string",
		},
	}
}
func CastGiantDtoObjectValueInnerObjectFromCli(c emigo.CliCastable) GiantDtoObjectValueInnerObject {
	data := GiantDtoObjectValueInnerObject{}
	if c.IsSet("inner-obj-text") {
		data.InnerObjText = c.String("inner-obj-text")
	}
	return data
}

// The base class definition for innerObject
type GiantDtoObjectValueInnerObject struct {
	InnerObjText string `json:"innerObjText" yaml:"innerObjText"`
}

func (x *GiantDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
