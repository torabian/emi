package core

// EmiManifest defines a manifest configuration used to describe
// included types and actions along with exclusion rules.
type EmiManifest struct {

	// Name is the unique identifier of the manifest.
	Name string `json:"name,omitempty" yaml:"name,omitempty" jsonschema:"title=Name,description=Unique name of the manifest,required"`

	// Package name used in some programming languages with package
	Package string `json:"package,omitempty" yaml:"package,omitempty" jsonschema:"title=Name,description=Package name used in some programming languages with package"`

	// Mod package name such as github.com/org/package
	ModPackageName string `json:"mod,omitempty" yaml:"mod,omitempty" jsonschema:"title=Name,description=Mod package name such as github.com/org/package"`

	// Types represents the list of supported runtime or build targets.
	// Allowed values: go-client, go-gin, go-cli, go-wasm.
	Types []string `json:"types,omitempty" yaml:"types,omitempty" jsonschema:"title=Types,description=List of supported types,enum=go-clientenum=go-cli,enum=go-gin,enum=go-wasm,uniqueItems=true"`

	// Includes defines the list of actions or patterns to include.
	// Supports wildcard matching.
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty" jsonschema:"title=Includes,description=Included actions or patterns (supports wildcards)"`

	// Excludes defines the list of conditions or patterns to exclude,
	// overriding Includes when matched.
	Excludes []string `json:"excludes,omitempty" yaml:"excludes,omitempty" jsonschema:"title=Excludes,description=Excluded conditions or patterns"`

	// The location that manifest will be written
	Dist string `json:"dist,omitempty" yaml:"dist,omitempty" jsonschema:"title=Dist,description=The location that manifest will be written"`
}
