package external
import (
"encoding"
"encoding/json"
"github.com/torabian/emi/examples/fullstack/sdk/complexes"
)
	import emigo "github.com/torabian/emi/examples/fullstack/emigo"
func GetCommonVectorComputeDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "initial-vector1",
			Type: "slice",
		},
		{
			Name: prefix + "value",
			Type: "string?",
		},
		{
			Name: prefix + "valuex",
			Type: "string",
		},
		{
			Name: prefix + "initial-vector2",
			Type: "slice",
		},
		{
			Name: prefix + "field-type-array",
			Type: "array",
		},
		{
			Name: prefix + "field-type-slice",
			Type: "slice",
		},
		{
			Name: prefix + "field-int",
			Type: "int",
		},
		{
			Name: prefix + "field-int-nullable",
			Type: "int?",
		},
		{
			Name: prefix + "complex-money",
			Type: "complex",
		},
	}
}
func CastCommonVectorComputeDtoFromCli(c emigo.CliCastable) CommonVectorComputeDto {
	data := CommonVectorComputeDto{}
			if c.IsSet("initial-vector1") { 
 emigo.InflatePossibleSlice(c.String("initial-vector1"), &data.InitialVector1) 
}
			if c.IsSet("value") { 
 emigo.ParseNullable(c.String("value"), &data.Value) 
}
			if c.IsSet("valuex") { 
 data.Valuex = c.String("valuex") 
 }
			if c.IsSet("initial-vector2") { 
 emigo.InflatePossibleSlice(c.String("initial-vector2"), &data.InitialVector2) 
}
			if c.IsSet("field-type-array") { 
 data.FieldTypeArray = emigo.CapturePossibleArray(CastCommonVectorComputeDtoFieldTypeArrayFromCli, "field-type-array", c) 
 }
			if c.IsSet("field-type-slice") { 
 emigo.InflatePossibleSlice(c.String("field-type-slice"), &data.FieldTypeSlice) 
}
			if c.IsSet("field-int") { 
 data.FieldInt = int(c.Int64("field-int")) 
 }
			if c.IsSet("field-int-nullable") { 
 emigo.ParseNullable(c.String("field-int-nullable"), &data.FieldIntNullable) 
}
			if c.IsSet("complex-money") { 
 if u, ok := any(&data.ComplexMoney).(encoding.TextUnmarshaler); ok { u.UnmarshalText([]byte(c.String("complex-money"))) } 
}
	return data
}
  // The base class definition for commonVectorComputeDto
type CommonVectorComputeDto struct {
		InitialVector1 []int `yaml:"initialVector1" json:"initialVector1"`
		Value emigo.Nullable[string] `json:"value" yaml:"value"`
		Valuex string `yaml:"valuex" json:"valuex"`
		InitialVector2 []int `json:"initialVector2" yaml:"initialVector2"`
		FieldTypeArray []CommonVectorComputeDtoFieldTypeArray `yaml:"fieldTypeArray" json:"fieldTypeArray"`
		FieldTypeSlice []string `yaml:"fieldTypeSlice" json:"fieldTypeSlice"`
		FieldInt int `json:"fieldInt" yaml:"fieldInt"`
		FieldIntNullable emigo.Nullable[int] `json:"fieldIntNullable" yaml:"fieldIntNullable"`
		ComplexMoney complexes.Money `json:"complexMoney" yaml:"complexMoney"`
}
func GetCommonVectorComputeDtoFieldTypeArrayCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "array-field1",
			Type: "string",
		},
		{
			Name: prefix + "array-field2",
			Type: "array",
		},
	}
}
func CastCommonVectorComputeDtoFieldTypeArrayFromCli(c emigo.CliCastable) CommonVectorComputeDtoFieldTypeArray {
	data := CommonVectorComputeDtoFieldTypeArray{}
			if c.IsSet("array-field1") { 
 data.ArrayField1 = c.String("array-field1") 
 }
			if c.IsSet("array-field2") { 
 data.ArrayField2 = emigo.CapturePossibleArray(CastCommonVectorComputeDtoFieldTypeArrayArrayField2FromCli, "array-field2", c) 
 }
	return data
}
  // The base class definition for fieldTypeArray
type CommonVectorComputeDtoFieldTypeArray struct {
		ArrayField1 string `json:"arrayField1" yaml:"arrayField1"`
		ArrayField2 []CommonVectorComputeDtoFieldTypeArrayArrayField2 `json:"arrayField2" yaml:"arrayField2"`
}
func GetCommonVectorComputeDtoFieldTypeArrayArrayField2CliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "last-item",
			Type: "string",
		},
	}
}
func CastCommonVectorComputeDtoFieldTypeArrayArrayField2FromCli(c emigo.CliCastable) CommonVectorComputeDtoFieldTypeArrayArrayField2 {
	data := CommonVectorComputeDtoFieldTypeArrayArrayField2{}
			if c.IsSet("last-item") { 
 data.LastItem = c.String("last-item") 
 }
	return data
}
  // The base class definition for arrayField2
type CommonVectorComputeDtoFieldTypeArrayArrayField2 struct {
		LastItem string `json:"lastItem" yaml:"lastItem"`
}
func (x *CommonVectorComputeDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}