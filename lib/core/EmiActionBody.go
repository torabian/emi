package core

// Represents any action request or response DTO.
// Used in multiple places as the request/response signature.
type EmiActionBody struct {

	// Typesafe headers for the action
	Headers []EmiHeader `yaml:"headers,omitempty" json:"headers,omitempty" jsonschema:"description=Typesafe headers."`

	// Envelop the content, for fields, dto automatically
	Envelope string `yaml:"envelope,omitempty" json:"envelope,omitempty" jsonschema:"description=Envelop the content, for fields, dto automatically."`

	// Defines the fields directly, and DTO will be generated
	// and assigned automatically.
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"Defines the fields directly and DTO will be generated and assigned automatically"`

	// Selects the DTO existing in the module from Golang.
	// It can also be a pure Go struct, but those do not compile.
	Dto string `yaml:"dto,omitempty" json:"dto,omitempty" jsonschema:"Selects the DTO existing in the module from Golang It can also be a pure Go struct but those do not compile"`

	// Implements lists the interfaces (Emi.Interfaces) the generated type satisfies.
	// Only honored where a body declares an inline shape for a standalone type: an
	// event's params/payload and a permission's params. Their fields are included in
	// front of the body's own, exactly like a dto's `implements`. Requires `fields`.
	Implements []string `yaml:"implements,omitempty" json:"implements,omitempty" jsonschema:"description=Names of interfaces (see interfaces) the generated type implements. Their fields are included and the generated code satisfies the interface. Honored for event params/payload and permission params; requires fields."`

	// Uses a primitive type instead, such as a string or int64.
	Primitive string `yaml:"primitive,omitempty" json:"primitive,omitempty" jsonschema:"enum=string,enum=bool,enum=int,enum=int32,enum=int64,enum=float32,enum=float64,enum=bytes,description=The primitve that is allowed for the slice is limited to following values"`
}
