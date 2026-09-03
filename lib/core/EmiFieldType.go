package core

type FieldType string

const (

	// IMPORTANT ---- before editing.
	// Wen adding a new type here, remember to add it into func GetEmiFieldTypeCatalog
	// so the consumers know about fields.

	// Array referes to Array<object>, when the field is an object, which can contain
	// another object in it. It's different from
	FieldTypeArray FieldType = "array"

	// List is like array, but golang-only for now: it renders as a plain []*ChildStruct
	// (no emigo.Array[T]/Operation wrapper), so gorm's own reflection-based schema
	// builder recognizes it as a real has-many association directly - no gorm:"-", no
	// hidden shadow field. Use array/array? for a dto's request/response shape (where
	// Operation - replace vs append - matters); use _list/_list? for the persisted
	// side of an entity's own struct, where there's nothing to apply an operation
	// *against* other than the database itself.
	FieldTypeList  FieldType = "_list"
	FieldTypeSlice FieldType = "slice"
	FieldTypeOne   FieldType = "one"

	// Class is like one, but golang-only for now: it renders as a plain *Target (no
	// emigo.One[T]/Operation wrapper), so gorm's own reflection-based schema builder
	// recognizes it as a real belongs-to association directly - no gorm:"-", no hidden
	// {field}Row shadow field (the {field}Id FK column sibling still exists, same as
	// one/one? - a belongs-to always needs a real scalar FK column of its own). Use
	// one/one? for a dto's request/response shape (where Operation - select vs inline
	// value - matters); use class/class? for the persisted side of an entity's own
	// struct, where there's nothing to apply an operation *against* other than the
	// database itself.
	FieldTypeClass      FieldType = "class"
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
	FieldTypeMap        FieldType = "map"

	FieldTypeArrayNullable      FieldType = "array?"
	FieldTypeListNullable       FieldType = "_list?"
	FieldTypeSliceNullable      FieldType = "slice?"
	FieldTypeOneNullable        FieldType = "one?"
	FieldTypeClassNullable      FieldType = "class?"
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
	FieldTypeMapNullable        FieldType = "map?"

	// Non-nullable fields, which doesn't matter will go here.
	FieldTypeAny     FieldType = "any"
	FieldTypeComplex FieldType = "complex"

	// FieldTypeComplexNullable ("complex?") exists only so the JS/TS generator can
	// tell "always instantiate the complex class" (complex) apart from "behaves like
	// every other nullable field - undefined/null allowed on the setter and the
	// definition" (complex?). A complex value has no wire-level nullability of its
	// own (see core.nullableFieldType's doc comment - it's returned unchanged for
	// FieldTypeComplex), so every non-JS backend treats complex? exactly like
	// complex: same generated code either way. See IsComplexFieldType.
	FieldTypeComplexNullable FieldType = "complex?"
)

// IsComplexFieldType reports whether t is the complex field type or its
// JS-nullable variant complex?. Every generator except lib/js treats the two
// identically (complex has no nullability of its own outside JS), so callers that
// only care "is this a complex field" should use this instead of comparing directly
// against FieldTypeComplex.
func IsComplexFieldType(t FieldType) bool {
	return t == FieldTypeComplex || t == FieldTypeComplexNullable
}

// Expose some information about available types in the codebase.
type FieldSupportCatalog struct {
	DtoFieldTypes         []FieldType
	DtoNullableFieldTypes []FieldType
}

func GetEmiFieldTypeCatalog() FieldSupportCatalog {

	return FieldSupportCatalog{
		DtoFieldTypes: []FieldType{
			FieldTypeArray,
			FieldTypeList,
			FieldTypeSlice,
			FieldTypeOne,
			FieldTypeClass,
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
			FieldTypeMap,
		},
		DtoNullableFieldTypes: []FieldType{
			FieldTypeArrayNullable,
			FieldTypeListNullable,
			FieldTypeSliceNullable,
			FieldTypeOneNullable,
			FieldTypeClassNullable,
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
			FieldTypeMapNullable,
		},
	}
}
