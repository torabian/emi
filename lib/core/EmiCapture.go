package core

// EmiCapture pulls fields from another structure (typically an EmiDto) into
// the owner's field list during preprocessing. It lets a vsql (or any other
// field-bearing construct) reuse a DTO's shape without redeclaring each field.
//
// At preprocessing time the captures are resolved and flattened — by the time
// the generator runs, the captured fields are already present in the owner's
// Params/Fields slice as if they had been written inline.
type EmiCapture struct {
	// Description of the capture. Used in generated documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the capture. Used in generated documentation."`

	// Dto is the name of a top-level EmiDto (Emi.Dto) whose fields will be
	// captured. Accepts the bare dto name (e.g. "userDto") or the prefixed
	// form "dto/userDto" — the prefix is stripped during preprocessing.
	// If both Dto and Template are set, Dto wins.
	Dto string `yaml:"dto,omitempty" json:"dto,omitempty" jsonschema:"description=Source EmiDto name from top-level Emi.Dto. Preferred over Template if both are set."`

	// Template is the name of an EmiDto defined under Emi.Templates.Dtos.
	// Used the same way as Dto but resolves only against the template scratch
	// area. Ignored when Dto is also set.
	Template string `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"description=Source EmiDto name from Emi.Templates.Dtos. Ignored when Dto is set."`

	// Action sources fields from an action's body. Accepts "name" (defaults to
	// the in body), "name.in", or "name.out". Lookup considers Emi.Actions and
	// Emi.Templates.Actions. Mutually exclusive with Dto/Template.
	Action string `yaml:"action,omitempty" json:"action,omitempty" jsonschema:"description=Source action body fields. Format: name, name.in, or name.out. Searches top-level actions and templates."`

	// Include limits which fields are pulled from the source. Each entry is a
	// field name (matched by EmiField.Name). When empty, every field in the
	// source DTO is captured.
	//
	// The special token "self.fields" merges the owner's own inline fields at
	// that position — so you can pick fields from multiple DTOs and combine
	// them with inline fields into a single ordered list.
	Include []string `yaml:"include,omitempty" json:"include,omitempty" jsonschema:"description=Whitelist of field names to capture from the source DTO. The token 'self.fields' inlines the owner's own fields at that position."`

	// Exclude removes fields by name after Include has been applied. Useful for
	// 'everything except a few' style captures.
	Exclude []string `yaml:"exclude,omitempty" json:"exclude,omitempty" jsonschema:"description=Field names to drop from the captured set (applied after Include)."`
}
