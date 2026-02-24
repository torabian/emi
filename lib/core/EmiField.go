package core

import "strings"

type EmiField struct {

	// Name of the field in camel case. Will be upper case automatically when necessary
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the field in camel case. Will be upper case automatically when necessary"`

	// The field name apperance on the cli tools, such as --field-name. If empty, computed automatically
	CliName string `yaml:"cli,omitempty" json:"cli,omitempty" jsonschema:"description=The field name apperance on the cli tools, such as --field-name. If empty, computed automatically"`

	// Recommended field will be asked upon an interactive cli operation.
	Recommended bool `yaml:"recommended,omitempty" json:"recommended,omitempty" jsonschema:"description=Recommended field will be asked upon an interactive cli operation."`

	// Description about the field for developers and generated documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the field for developers and generated documents."`

	// Type of the field based on Emi types.
	Type FieldType `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=array,enum=map?,enum=map,enum=slice,enum=one,enum=collection,enum=object,enum=enum,enum=string,enum=bool,enum=int,enum=int32,enum=int64,enum=float32,enum=float64,enum=array?,enum=slice?,enum=one?,enum=collection?,enum=object?,enum=enum?,enum=string?,enum=bool?,enum=int?,enum=int32?,enum=int64?,enum=float32?,enum=float64?,enum=any,enum=complex,description=Type of the field based on Emi types."`

	// Type of the field based on Emi types.
	MapKey string `yaml:"mapKey,omitempty" json:"mapKey,omitempty" jsonschema:"enum=string,enum=int,enum=any,description=When the field type is map or map? you can modify the key type."`

	// Primitive type in golang when type: slice is set
	Primitive string `yaml:"primitive,omitempty" json:"primitive,omitempty" jsonschema:"description=Primitive type in golang when type: slice is set"`

	// The entity in golang which will be operated on in case of type: one or type: collection
	Target string `yaml:"target,omitempty" json:"target,omitempty" jsonschema:"description=The entity in golang which will be operated on in case of type: one or type: collection"`

	// Default value of the field which will be added to the meta tags
	Default interface{} `yaml:"default,omitempty" json:"default,omitempty" jsonschema:"description=Default value of the field which will be added to the meta tags"`

	// When using one or collection types you need to set the module name here to import that in generated go file.
	Module string `yaml:"module,omitempty" json:"module,omitempty" jsonschema:"description=When using one or collection types you need to set the module name here to import tha"`

	// The go project module of the important target for one or collection fields if its from external library
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty" jsonschema:"description=The go project module of the important target for one or collection fields if its from exte"`

	// List of enum values in case of enum type for the field. Check EmiEnumInline for more details how to define them.
	OfType []*EmiEnumInline `yaml:"of,omitempty" json:"of,omitempty" jsonschema:"description=List of enum values in case of enum type for the field. Check EmiEnumInline for more d"`

	// In case of complex data type, it would be using this as class or struct name
	Complex string `yaml:"complex,omitempty" json:"complex,omitempty" jsonschema:"description=In case of complex data type, it would be using this as class or struct name"`

	// On the json type this field will generate necessary code to cast it into different dtos
	Matches []*EmiFieldMatch `yaml:"matches,omitempty" json:"matches,omitempty" jsonschema:"description=On the json type this field will generate necessary code to cast it into different dtos"`

	// For types such as array or object children fields can be defined and will separate struct with name prefixed to parent
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=For types such as array or object children fields can be defined and will separate struct with name prefixed to parent"`

	// Tags are key/pair values which might be useful in some languages - such as golang. There you can specifiy extra tags for language.
	Tags map[string]string `yaml:"tags,omitempty" json:"tags,omitempty" jsonschema:"description=Tags are key/pair values which might be useful in some languages - such as golang. There you can specifiy extra tags for language."`
}

func (x *EmiField) PublicName() string {
	return ToUpper(x.Name)
}
func (x *EmiField) PrivateName() string {
	return x.Name
}

func (x *EmiField) GetType() FieldType {
	return x.Type
}

func (x *EmiField) GetModule() string {
	return x.Module
}

func (x *EmiField) GetTarget() string {
	return x.Target
}

func (x *EmiField) GetPrimitive() string {
	return x.Primitive
}

func (x *EmiField) GetName() string {
	return x.Name
}

func (x *EmiField) GetComplex() string {
	return x.Complex
}

func (x *EmiField) GetCliName() string {
	if x.CliName != "" {
		return x.CliName
	}
	return strings.ReplaceAll(ToSnakeCase((x.Name)), "_", "-")
}

func (x *EmiField) GetDescription() string {
	return x.Description
}

func (x *EmiField) GetTags() map[string]string {
	return x.Tags
}
