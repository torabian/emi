package core

type EmiQueryField struct {

	// Name of the field in camel case. Will be upper case automatically when necessary
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the field in camel case. Will be upper case automatically when necessary"`

	// Description about the field for developers and generated documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the field for developers and generated documents."`

	// Type of the field based on Emi types.
	Type FieldType `yaml:"type,omitempty" json:"type,omitempty" jsonschema:"enum=array,enum=slice,enum=object,enum=string,enum=int,enum=bool,enum=int32,enum=int64,enum=float32,enum=float64,description=Type of the field based on Emi types."`

	// Primitive type in golang when type: slice is set
	Primitive string `yaml:"primitive,omitempty" json:"primitive,omitempty" jsonschema:"description=Primitive type in golang when type: slice is set"`

	// For types such as array or object children fields can be defined and will separate struct with name prefixed to parent
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=For types such as array or object children fields can be defined and will separate struct with name prefixed to parent"`
}
