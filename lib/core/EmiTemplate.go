package core

// EmiTemplate is a scratch area for definitions that exist solely to be
// referenced by other parts of the module (e.g. via EmiCapture). Nothing in
// Templates is fed to language generators — no Go struct, no client SDK code,
// no manifest entry is produced from these. Treat them as reusable shape
// libraries: they hold dtos and actions whose fields can be picked up by
// captures elsewhere in the module.
type EmiTemplate struct {
	// Dtos defined for reuse only. Same shape as top-level Emi.Dto entries,
	// but never compiled into output files.
	Dtos []EmiDto `yaml:"dtos,omitempty" json:"dtos,omitempty" jsonschema:"description=DTOs defined for reuse only. Never compiled into output files; available as capture sources."`

	// Actions defined for reuse only. Same shape as top-level Emi.Actions,
	// but never compiled into routes, CLI commands, or client bindings. Their
	// in/out body fields can be sourced via EmiCapture.Action.
	Actions []*EmiAction `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions defined for reuse only. Never compiled; in/out fields are available as capture sources."`

	// Fields defined for reuse only: named field shapes that any field can pull in with
	// `use: <name>`. See EmiFieldTemplate.
	Fields []EmiFieldTemplate `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Reusable field shapes. Any field can build itself from one with use: <name>."`
}

// EmiFieldTemplate is a reusable field definition. Name is the template's own name (what
// `use:` refers to), not the name of the fields built from it.
type EmiFieldTemplate struct {
	EmiField `yaml:",inline"`

	// Context restricts the template to fields declared inside an entity ("entity") or
	// anywhere else - dtos and action bodies ("dto"). Leave it empty to match both. When a
	// name has both a context-specific and a general template, the specific one wins - so
	// one name can describe a relation as PurchasableEntity in entities and as
	// PurchasableDto in dtos.
	Context string `yaml:"context,omitempty" json:"context,omitempty" jsonschema:"enum=entity,enum=dto,description=Limit the template to fields inside entities or inside dtos and action bodies. Empty matches both."`
}
