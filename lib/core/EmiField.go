package core

type FieldType string

const (

	// Wen adding a new type here, remember to add it into func GetEmiFieldTypeCatalog
	// so the consumers know about fields.
	FieldTypeArray     FieldType = "array"
	FieldTypeArrayP    FieldType = "arrayP"
	FieldTypeOne       FieldType = "one"
	FieldTypeMany2Many FieldType = "many2many"
	FieldTypeObject    FieldType = "object"
	FieldTypeEmbed     FieldType = "embed"
	FieldTypeEnum      FieldType = "enum"
	FieldTypeText      FieldType = "text"
	FieldTypeString    FieldType = "string"
	FieldTypeBool      FieldType = "bool"
	FieldTypeInt       FieldType = "int"
	FieldTypeInt32     FieldType = "int32"
	FieldTypeInt64     FieldType = "int64"
	FieldTypeFloat32   FieldType = "float32"
	FieldTypeFloat64   FieldType = "float64"

	FieldTypeArrayNullable     FieldType = "array?"
	FieldTypeArrayPNullable    FieldType = "arrayP?"
	FieldTypeOneNullable       FieldType = "one?"
	FieldTypeMany2ManyNullable FieldType = "many2many?"
	FieldTypeObjectNullable    FieldType = "object?"
	FieldTypeEmbedNullable     FieldType = "embed?"
	FieldTypeEnumNullable      FieldType = "enum?"
	FieldTypeTextNullable      FieldType = "text?"
	FieldTypeStringNullable    FieldType = "string?"
	FieldTypeBoolNullable      FieldType = "bool?"
	FieldTypeIntNullable       FieldType = "int?"
	FieldTypeInt32Nullable     FieldType = "int32?"
	FieldTypeInt64Nullable     FieldType = "int64?"
	FieldTypeFloat32Nullable   FieldType = "float32?"
	FieldTypeFloat64Nullable   FieldType = "float64?"

	// Non-nullable fields, which doesn't matter will go here.
	FieldTypeAny FieldType = "any"
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
			FieldTypeArrayP,
			FieldTypeOne,
			FieldTypeMany2Many,
			FieldTypeObject,
			FieldTypeEmbed,
			FieldTypeEnum,
			FieldTypeText,
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
			FieldTypeArrayPNullable,
			FieldTypeOneNullable,
			FieldTypeMany2ManyNullable,
			FieldTypeObjectNullable,
			FieldTypeEmbedNullable,
			FieldTypeEnumNullable,
			FieldTypeTextNullable,
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

type EmiField struct {

	// Name of the field in camel case. Will be upper case automatically when necessary
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the field in camel case. Will be upper case automatically when necessary"`

	// Recommended field will be asked upon an interactive cli operation.
	Recommended bool `yaml:"recommended,omitempty" json:"recommended,omitempty" jsonschema:"description=Recommended field will be asked upon an interactive cli operation."`

	// Description about the field for developers and generated documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the field for developers and generated documents."`

	// Type of the field based on Emi types.
	Type FieldType `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=object?,enum=many2many?,enum=one?,enum=array?,enum=string?,enum=int?,enum=any,enum=complex,enum=float64?,enum=money?,enum=xfile?,enum=float32?,enum=bool?,enum=int32?,enum=int64?,enum=int,enum=datetime,enum=json,enum=embed,enum=datenano,enum=html,enum=text,enum=date,enum=daterange,enum=many2many,enum=arrayP,enum=enum,enum=bool,enum=one,enum=int64,enum=float64,enum=duration?,enum=object,enum=array,enum=string,description=Type of the field based on Emi types."`

	// Primitive type in golang when type: arrayP is set
	Primitive string `yaml:"primitive,omitempty" json:"primitive,omitempty" jsonschema:"description=Primitive type in golang when type: arrayP is set"`

	// The entity in golang which will be operated on in case of type: one or type: many2many
	Target string `yaml:"target,omitempty" json:"target,omitempty" jsonschema:"description=The entity in golang which will be operated on in case of type: one or type: many2many"`

	// The meta tag for validate library which will be checked on different operations
	Validate string `yaml:"validate,omitempty" json:"validate,omitempty" jsonschema:"description=The meta tag for validate library which will be checked on different operations"`

	// For the html and text fields there will be a automatic excerpt generated.
	ExcerptSize int `yaml:"excerptSize,omitempty" json:"excerptSize,omitempty" jsonschema:"description=For the html and text fields there will be a automatic excerpt generated."`

	// Default value of the field which will be added to the meta tags
	Default interface{} `yaml:"default,omitempty" json:"default,omitempty" jsonschema:"description=Default value of the field which will be added to the meta tags"`

	// If true adds the field into polyglot table for translations. Only works with the first level fields.
	Translate bool `yaml:"translate,omitempty" json:"translate,omitempty" jsonschema:"description=If true adds the field into polyglot table for translations. Only works with the first leve"`

	// It would skip the sanitization for html field types allowing store anything as html
	Unsafe bool `yaml:"unsafe,omitempty" json:"unsafe,omitempty" jsonschema:"description=It would skip the sanitization for html field types allowing store anything as htm"`

	// Allow create is a useful option to set true if the type one or many2many could be allowed to create entities on target table.
	AllowCreate bool `yaml:"allowCreate,omitempty" json:"allowCreate,omitempty" jsonschema:"description=Allow create is a useful option to set true if the type one or many2many could be allowed to crea"`

	// When using one or many2many types you need to set the module name here to import that in generated go file.
	Module string `yaml:"module,omitempty" json:"module,omitempty" jsonschema:"description=When using one or many2many types you need to set the module name here to import tha"`

	// The go project module of the important target for one or many2many fields if its from external library
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty" jsonschema:"description=The go project module of the important target for one or many2many fields if its from exte"`

	// The json tag of the generated field. Defaults to the name but can be overwritten with this field
	Json string `yaml:"json,omitempty" json:"json,omitempty" jsonschema:"description=The json tag of the generated field. Defaults to the name but can be overwritten"`

	// The yaml tag of the generated field. Defaults to the name but can be overwritten with this field
	Yaml string `yaml:"yaml,omitempty" json:"yaml,omitempty" jsonschema:"description=The yaml tag of the generated field. Defaults to the name but can be overwritten"`

	// The xml tag of the generated field. Defaults to the name but can be overwritten with this field
	Xml string `yaml:"xml,omitempty" json:"xml,omitempty" jsonschema:"description=The xml tag of the generated field. Defaults to the name but can be overwritten"`

	// List of enum values in case of enum type for the field. Check EmiEnumInline for more details how to define them.
	OfType []*EmiEnumInline `yaml:"of,omitempty" json:"of,omitempty" jsonschema:"description=List of enum values in case of enum type for the field. Check EmiEnumInline for more d"`

	// When type is one there will be another field added with Id prefix. This tag will override gorm meta tag of that field
	IdFieldGorm string `yaml:"idFieldGorm,omitempty" json:"idFieldGorm,omitempty" jsonschema:"description=When type is one there will be another field added with Id prefix. This tag will override gorm meta"`

	// In case of complex data type, it would be using this as class or struct name
	Complex string `yaml:"complex,omitempty" json:"complex,omitempty" jsonschema:"description=In case of complex data type, it would be using this as class or struct name"`

	// On the json type this field will generate necessary code to cast it into different dtos
	Matches []*EmiFieldMatch `yaml:"matches,omitempty" json:"matches,omitempty" jsonschema:"description=On the json type this field will generate necessary code to cast it into different dtos"`

	// Override the gorm meta tag generated for golang, to add custom types or anything else.
	Gorm string `yaml:"gorm,omitempty" json:"gorm,omitempty" jsonschema:"description=Override the gorm meta tag generated for golang, to add custom types or anything else."`

	// Override the Emi default fields gorm tags for extra constraint or other configuration.
	GormMap GormOverrideMap `yaml:"gormMap,omitempty" json:"gormMap,omitempty" jsonschema:"description=Used in Module code generation to customized the generated code for gorm tags on Emi Data management fields such as workspace or user id. For example, you can add extra indexes on these fields."`

	// Direct manipulation of the sql meta tag on the field.
	Sql string `yaml:"sql,omitempty" json:"sql,omitempty" jsonschema:"description=Direct manipulation of the sql meta tag on the field."`

	// This is the name of field considering how deep it is. Used internally for Emi codegen, not available on definition
	FullName string `yaml:"-,omitempty" json:"-,omitempty" jsonschema:"-"`

	// For types such as array or object children fields can be defined and will separate struct with name prefixed to parent
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=For types such as array or object children fields can be defined and will separate struct with name prefixed to parent"`

	IsVirtualObject bool `yaml:"-" json:"-" jsonschema:"-"`

	LinkedTo string `yaml:"linkedTo,omitempty" json:"linkedTo,omitempty" jsonschema:"-"`

	RootClass string `yaml:"rootClass,omitempty" json:"rootClass,omitempty" jsonschema:"-"`

	BelongingEntityName string `yaml:"-" json:"-"`
}
