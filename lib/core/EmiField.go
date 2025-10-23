package core

type EmiField struct {

	// Name of the field in camel case. Will be upper case automatically when necessary
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the field in camel case. Will be upper case automatically when necessary"`

	// Recommended field will be asked upon an interactive cli operation.
	Recommended bool `yaml:"recommended,omitempty" json:"recommended,omitempty" jsonschema:"description=Recommended field will be asked upon an interactive cli operation."`

	// Description about the field for developers and generated documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the field for developers and generated documents."`

	// Type of the field based on Emi types.
	Type FieldType `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=array,enum=slice,enum=one,enum=collection,enum=object,enum=enum,enum=string,enum=bool,enum=int,enum=int32,enum=int64,enum=float32,enum=float64,enum=array?,enum=slice?,enum=one?,enum=collection?,enum=object?,enum=enum?,enum=string?,enum=bool?,enum=int?,enum=int32?,enum=int64?,enum=float32?,enum=float64?,enum=any,enum=complex,description=Type of the field based on Emi types."`

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

	// The json tag of the generated field. Defaults to the name but can be overwritten with this field
	Json string `yaml:"json,omitempty" json:"json,omitempty" jsonschema:"description=The json tag of the generated field. Defaults to the name but can be overwritten"`

	// The yaml tag of the generated field. Defaults to the name but can be overwritten with this field
	Yaml string `yaml:"yaml,omitempty" json:"yaml,omitempty" jsonschema:"description=The yaml tag of the generated field. Defaults to the name but can be overwritten"`

	// The xml tag of the generated field. Defaults to the name but can be overwritten with this field
	Xml string `yaml:"xml,omitempty" json:"xml,omitempty" jsonschema:"description=The xml tag of the generated field. Defaults to the name but can be overwritten"`

	// List of enum values in case of enum type for the field. Check EmiEnumInline for more details how to define them.
	OfType []*EmiEnumInline `yaml:"of,omitempty" json:"of,omitempty" jsonschema:"description=List of enum values in case of enum type for the field. Check EmiEnumInline for more d"`

	// In case of complex data type, it would be using this as class or struct name
	Complex string `yaml:"complex,omitempty" json:"complex,omitempty" jsonschema:"description=In case of complex data type, it would be using this as class or struct name"`

	// On the json type this field will generate necessary code to cast it into different dtos
	Matches []*EmiFieldMatch `yaml:"matches,omitempty" json:"matches,omitempty" jsonschema:"description=On the json type this field will generate necessary code to cast it into different dtos"`

	// For types such as array or object children fields can be defined and will separate struct with name prefixed to parent
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=For types such as array or object children fields can be defined and will separate struct with name prefixed to parent"`
}

func (x *EmiField) PublicName() string {
	return ToUpper(x.Name)
}
func (x *EmiField) PrivateName() string {
	return x.Name
}
