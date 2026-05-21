package core

// EmiVsql defines a hand-written SQL query alongside the typed parameter
// shape it expects. Emi generates a parameters DTO and a Prepare helper so
// callers pass values through a struct instead of positional args.
type EmiVsql struct {
	// Description of what the query does. Used in generated code documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of what the query does. Used in generated code documentation."`

	// Name of the vsql query in camel case. Used as the file and identifier base.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the vsql query in camel case. Used as the file and identifier base."`

	// Params is the typed parameter list for the query, same shape as dto fields.
	Params []*EmiField `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Typed parameter list for the query, same shape as dto fields."`

	// Captures pulls fields from existing DTOs into Params at preprocessing
	// time. By the time the generator runs, the resolved fields have been
	// flattened into Params and Captures is cleared. See EmiCapture.
	Captures []*EmiCapture `yaml:"captures,omitempty" json:"captures,omitempty" jsonschema:"description=Capture fields from existing DTOs into Params. Resolved during preprocessing."`

	// Query is the raw SQL (or templated SQL) string the query will execute.
	Query string `yaml:"query,omitempty" json:"query,omitempty" jsonschema:"description=Raw SQL string the query will execute."`
}

func (x EmiVsql) GetClassName() string {
	return ToUpper(x.Name) + "Vsql"
}

func (x EmiVsql) GetParamsClassName() string {
	return ToUpper(x.Name) + "VsqlParams"
}
