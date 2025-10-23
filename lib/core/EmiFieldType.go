package core

type FieldType string

const (

	// IMPORTANT ---- before editing.
	// Wen adding a new type here, remember to add it into func GetEmiFieldTypeCatalog
	// so the consumers know about fields.

	// Array referes to Array<object>, when the field is an object, which can contain
	// another object in it. It's different from
	FieldTypeArray      FieldType = "array"
	FieldTypeSlice      FieldType = "slice"
	FieldTypeOne        FieldType = "one"
	FieldTypeCollection FieldType = "collection"
	FieldTypeObject     FieldType = "object"
	FieldTypeEnum       FieldType = "enum"
	FieldTypeString     FieldType = "string"
	FieldTypeBool       FieldType = "bool"
	FieldTypeInt        FieldType = "int"
	FieldTypeInt32      FieldType = "int32"
	FieldTypeInt64      FieldType = "int64"
	FieldTypeFloat32    FieldType = "float32"
	FieldTypeFloat64    FieldType = "float64"

	FieldTypeArrayNullable      FieldType = "array?"
	FieldTypeSliceNullable      FieldType = "slice?"
	FieldTypeOneNullable        FieldType = "one?"
	FieldTypeCollectionNullable FieldType = "collection?"
	FieldTypeObjectNullable     FieldType = "object?"
	FieldTypeEnumNullable       FieldType = "enum?"
	FieldTypeStringNullable     FieldType = "string?"
	FieldTypeBoolNullable       FieldType = "bool?"
	FieldTypeIntNullable        FieldType = "int?"
	FieldTypeInt32Nullable      FieldType = "int32?"
	FieldTypeInt64Nullable      FieldType = "int64?"
	FieldTypeFloat32Nullable    FieldType = "float32?"
	FieldTypeFloat64Nullable    FieldType = "float64?"

	// Non-nullable fields, which doesn't matter will go here.
	FieldTypeAny     FieldType = "any"
	FieldTypeComplex FieldType = "complex"
)

// Expose some information about available types in the codebase.
type FieldSupportCatalog struct {
	DtoFieldTypes         []FieldType
	DtoNullableFieldTypes []FieldType
}

func GetEmiFieldTypeCatalog() FieldSupportCatalog {

	return FieldSupportCatalog{
		DtoFieldTypes: []FieldType{
			FieldTypeArray,
			FieldTypeSlice,
			FieldTypeOne,
			FieldTypeCollection,
			FieldTypeObject,
			FieldTypeEnum,
			FieldTypeString,
			FieldTypeBool,
			FieldTypeInt,
			FieldTypeInt32,
			FieldTypeInt64,
			FieldTypeFloat32,
			FieldTypeFloat64,
		},
		DtoNullableFieldTypes: []FieldType{
			FieldTypeArrayNullable,
			FieldTypeSliceNullable,
			FieldTypeOneNullable,
			FieldTypeCollectionNullable,
			FieldTypeObjectNullable,
			FieldTypeEnumNullable,
			FieldTypeStringNullable,
			FieldTypeBoolNullable,
			FieldTypeIntNullable,
			FieldTypeInt32Nullable,
			FieldTypeInt64Nullable,
			FieldTypeFloat32Nullable,
			FieldTypeFloat64Nullable,
		},
	}
}
