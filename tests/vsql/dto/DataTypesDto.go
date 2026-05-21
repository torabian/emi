package dto

import (
	"encoding"
	"encoding/json"
)
import emigo "github.com/torabian/emi/emigo"

func GetDataTypesDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "string-field",
			Type: "string",
		},
		{
			Name: prefix + "string-field-nullable",
			Type: "string?",
		},
		{
			Name: prefix + "bool-field",
			Type: "bool",
		},
		{
			Name: prefix + "bool-field-nullable",
			Type: "bool?",
		},
		{
			Name: prefix + "int-field",
			Type: "int",
		},
		{
			Name: prefix + "int-field-nullable",
			Type: "int?",
		},
		{
			Name: prefix + "int32-field",
			Type: "int32",
		},
		{
			Name: prefix + "int32-field-nullable",
			Type: "int32?",
		},
		{
			Name: prefix + "int64-field",
			Type: "int64",
		},
		{
			Name: prefix + "int64-field-nullable",
			Type: "int64?",
		},
		{
			Name: prefix + "float32-field",
			Type: "float32",
		},
		{
			Name: prefix + "float32-field-nullable",
			Type: "float32?",
		},
		{
			Name: prefix + "float64-field",
			Type: "float64",
		},
		{
			Name: prefix + "float64-field-nullable",
			Type: "float64?",
		},
		{
			Name: prefix + "enum-field",
			Type: "enum",
		},
		{
			Name: prefix + "enum-field-nullable",
			Type: "enum?",
		},
		{
			Name:     prefix + "object-field",
			Type:     "object",
			Children: GetDataTypesDtoObjectFieldCliFlags("object-field-"),
		},
		{
			Name: prefix + "object-field-nullable",
			Type: "object?",
		},
		{
			Name: prefix + "array-field",
			Type: "array",
		},
		{
			Name: prefix + "array-field-nullable",
			Type: "array?",
		},
		{
			Name: prefix + "slice-field",
			Type: "slice",
		},
		{
			Name: prefix + "slice-field-nullable",
			Type: "slice?",
		},
		{
			Name: prefix + "any-field",
			Type: "any",
		},
		{
			Name: prefix + "map-field",
			Type: "map",
		},
		{
			Name: prefix + "map-field-nullable",
			Type: "map?",
		},
		{
			Name: prefix + "complex-field",
			Type: "complex",
		},
		{
			Name: prefix + "one-ref",
			Type: "one",
		},
		{
			Name: prefix + "one-ref-nullable",
			Type: "one?",
		},
		{
			Name: prefix + "collection-ref",
			Type: "collection",
		},
		{
			Name: prefix + "collection-ref-nullable",
			Type: "collection?",
		},
	}
}
func CastDataTypesDtoFromCli(c emigo.CliCastable) DataTypesDto {
	data := DataTypesDto{}
	if c.IsSet("string-field") {
		data.StringField = c.String("string-field")
	}
	if c.IsSet("string-field-nullable") {
		emigo.ParseNullable(c.String("string-field-nullable"), &data.StringFieldNullable)
	}
	if c.IsSet("bool-field") {
		data.BoolField = bool(c.Bool("bool-field"))
	}
	if c.IsSet("bool-field-nullable") {
		emigo.ParseNullable(c.String("bool-field-nullable"), &data.BoolFieldNullable)
	}
	if c.IsSet("int-field") {
		data.IntField = int(c.Int64("int-field"))
	}
	if c.IsSet("int-field-nullable") {
		emigo.ParseNullable(c.String("int-field-nullable"), &data.IntFieldNullable)
	}
	if c.IsSet("int32-field") {
		data.Int32Field = int32(c.Int64("int32-field"))
	}
	if c.IsSet("int32-field-nullable") {
		emigo.ParseNullable(c.String("int32-field-nullable"), &data.Int32FieldNullable)
	}
	if c.IsSet("int64-field") {
		data.Int64Field = int64(c.Int64("int64-field"))
	}
	if c.IsSet("int64-field-nullable") {
		emigo.ParseNullable(c.String("int64-field-nullable"), &data.Int64FieldNullable)
	}
	if c.IsSet("float32-field") {
		data.Float32Field = float32(c.Float64("float32-field"))
	}
	if c.IsSet("float32-field-nullable") {
		emigo.ParseNullable(c.String("float32-field-nullable"), &data.Float32FieldNullable)
	}
	if c.IsSet("float64-field") {
		data.Float64Field = float64(c.Float64("float64-field"))
	}
	if c.IsSet("float64-field-nullable") {
		emigo.ParseNullable(c.String("float64-field-nullable"), &data.Float64FieldNullable)
	}
	if c.IsSet("enum-field-nullable") {
		emigo.ParseNullable(c.String("enum-field-nullable"), &data.EnumFieldNullable)
	}
	if c.IsSet("object-field") {
		data.ObjectField = CastDataTypesDtoObjectFieldFromCli(c)
	}
	if c.IsSet("object-field-nullable") {
		emigo.ParseNullable(c.String("object-field-nullable"), &data.ObjectFieldNullable)
	}
	if c.IsSet("array-field") {
		data.ArrayField = emigo.CapturePossibleArray(CastDataTypesDtoArrayFieldFromCli, "array-field", c)
	}
	if c.IsSet("array-field-nullable") {
		emigo.ParseNullable(c.String("array-field-nullable"), &data.ArrayFieldNullable)
	}
	if c.IsSet("slice-field") {
		emigo.InflatePossibleSlice(c.String("slice-field"), &data.SliceField)
	}
	if c.IsSet("slice-field-nullable") {
		emigo.ParseNullable(c.String("slice-field-nullable"), &data.SliceFieldNullable)
	}
	if c.IsSet("map-field-nullable") {
		emigo.ParseNullable(c.String("map-field-nullable"), &data.MapFieldNullable)
	}
	if c.IsSet("complex-field") {
		if u, ok := any(&data.ComplexField).(encoding.TextUnmarshaler); ok {
			u.UnmarshalText([]byte(c.String("complex-field")))
		}
	}
	if c.IsSet("one-ref-nullable") {
		emigo.ParseNullable(c.String("one-ref-nullable"), &data.OneRefNullable)
	}
	if c.IsSet("collection-ref-nullable") {
		emigo.ParseNullable(c.String("collection-ref-nullable"), &data.CollectionRefNullable)
	}
	return data
}

// The base class definition for dataTypesDto
type DataTypesDto struct {
	StringField           string                                           `json:"stringField" yaml:"stringField"`
	StringFieldNullable   emigo.Nullable[string]                           `json:"stringFieldNullable" yaml:"stringFieldNullable"`
	BoolField             bool                                             `json:"boolField" yaml:"boolField"`
	BoolFieldNullable     emigo.Nullable[bool]                             `json:"boolFieldNullable" yaml:"boolFieldNullable"`
	IntField              int                                              `json:"intField" yaml:"intField"`
	IntFieldNullable      emigo.Nullable[int]                              `json:"intFieldNullable" yaml:"intFieldNullable"`
	Int32Field            int32                                            `json:"int32Field" yaml:"int32Field"`
	Int32FieldNullable    emigo.Nullable[int32]                            `json:"int32FieldNullable" yaml:"int32FieldNullable"`
	Int64Field            int64                                            `json:"int64Field" yaml:"int64Field"`
	Int64FieldNullable    emigo.Nullable[int64]                            `json:"int64FieldNullable" yaml:"int64FieldNullable"`
	Float32Field          float32                                          `json:"float32Field" yaml:"float32Field"`
	Float32FieldNullable  emigo.Nullable[float32]                          `json:"float32FieldNullable" yaml:"float32FieldNullable"`
	Float64Field          float64                                          `json:"float64Field" yaml:"float64Field"`
	Float64FieldNullable  emigo.Nullable[float64]                          `json:"float64FieldNullable" yaml:"float64FieldNullable"`
	EnumField             string                                           `json:"enumField" yaml:"enumField"`
	EnumFieldNullable     emigo.Nullable[string]                           `json:"enumFieldNullable" yaml:"enumFieldNullable"`
	ObjectField           DataTypesDtoObjectField                          `json:"objectField" yaml:"objectField"`
	ObjectFieldNullable   emigo.Nullable[DataTypesDtoObjectFieldNullable]  `json:"objectFieldNullable" yaml:"objectFieldNullable"`
	ArrayField            []DataTypesDtoArrayField                         `json:"arrayField" yaml:"arrayField"`
	ArrayFieldNullable    emigo.Nullable[[]DataTypesDtoArrayFieldNullable] `json:"arrayFieldNullable" yaml:"arrayFieldNullable"`
	SliceField            []string                                         `json:"sliceField" yaml:"sliceField"`
	SliceFieldNullable    emigo.Nullable[[]int]                            `json:"sliceFieldNullable" yaml:"sliceFieldNullable"`
	AnyField              interface{}                                      `json:"anyField" yaml:"anyField"`
	MapField              map[any]any                                      `json:"mapField" yaml:"mapField"`
	MapFieldNullable      emigo.Nullable[map[any]any]                      `json:"mapFieldNullable" yaml:"mapFieldNullable"`
	ComplexField          GeoPoint                                         `json:"complexField" yaml:"complexField"`
	OneRef                CreateUserDtoTags                                `json:"oneRef" yaml:"oneRef"`
	OneRefNullable        emigo.Nullable[CreateUserDtoTags]                `json:"oneRefNullable" yaml:"oneRefNullable"`
	CollectionRef         []CreateUserDtoTags                              `json:"collectionRef" yaml:"collectionRef"`
	CollectionRefNullable emigo.Nullable[[]CreateUserDtoTags]              `json:"collectionRefNullable" yaml:"collectionRefNullable"`
}

func GetDataTypesDtoObjectFieldCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "nested-string",
			Type: "string",
		},
		{
			Name: prefix + "nested-int",
			Type: "int",
		},
	}
}
func CastDataTypesDtoObjectFieldFromCli(c emigo.CliCastable) DataTypesDtoObjectField {
	data := DataTypesDtoObjectField{}
	if c.IsSet("nested-string") {
		data.NestedString = c.String("nested-string")
	}
	if c.IsSet("nested-int") {
		data.NestedInt = int(c.Int64("nested-int"))
	}
	return data
}

// The base class definition for objectField
type DataTypesDtoObjectField struct {
	NestedString string `json:"nestedString" yaml:"nestedString"`
	NestedInt    int    `json:"nestedInt" yaml:"nestedInt"`
}

func GetDataTypesDtoObjectFieldNullableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "optional-string",
			Type: "string",
		},
	}
}
func CastDataTypesDtoObjectFieldNullableFromCli(c emigo.CliCastable) DataTypesDtoObjectFieldNullable {
	data := DataTypesDtoObjectFieldNullable{}
	if c.IsSet("optional-string") {
		data.OptionalString = c.String("optional-string")
	}
	return data
}

// The base class definition for objectFieldNullable
type DataTypesDtoObjectFieldNullable struct {
	OptionalString string `json:"optionalString" yaml:"optionalString"`
}

func GetDataTypesDtoArrayFieldCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "key",
			Type: "string",
		},
		{
			Name: prefix + "value",
			Type: "int",
		},
	}
}
func CastDataTypesDtoArrayFieldFromCli(c emigo.CliCastable) DataTypesDtoArrayField {
	data := DataTypesDtoArrayField{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	if c.IsSet("value") {
		data.Value = int(c.Int64("value"))
	}
	return data
}

// The base class definition for arrayField
type DataTypesDtoArrayField struct {
	Key   string `json:"key" yaml:"key"`
	Value int    `json:"value" yaml:"value"`
}

func GetDataTypesDtoArrayFieldNullableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "key",
			Type: "string",
		},
	}
}
func CastDataTypesDtoArrayFieldNullableFromCli(c emigo.CliCastable) DataTypesDtoArrayFieldNullable {
	data := DataTypesDtoArrayFieldNullable{}
	if c.IsSet("key") {
		data.Key = c.String("key")
	}
	return data
}

// The base class definition for arrayFieldNullable
type DataTypesDtoArrayFieldNullable struct {
	Key string `json:"key" yaml:"key"`
}

func (x *DataTypesDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
