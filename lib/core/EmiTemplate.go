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
}
