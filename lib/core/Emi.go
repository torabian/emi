package core

import "fmt"

type Emi struct {

	// Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty" jsonschema:"description=Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks"`

	// Compiler self containment
	Targets []EmiCompile `yaml:"targets,omitempty" json:"targets,omitempty" jsonschema:"description=Compiler self containment"`

	// Description of module and it's purpose. Used in code gen and creating documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of module and it's purpose. Used in code gen and creating documents."`

	// Version of the module. Helpful for different code generation phases but it's not necessary.
	Version string `yaml:"version,omitempty" json:"version,omitempty" jsonschema:"description=Version of the module. Helpful for different code generation phases but it's not necessary."`

	// Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name."`

	// Module level enums which can be used across the other parts of generated code.
	Enums []EmiEnum `yaml:"enums,omitempty" json:"enums,omitempty" jsonschema:"description=Module level enums which can be used across the other parts of generated code."`

	// Dtos are basically golang structs with some additional functionality which can be used for request/response actions
	Dto []EmiDto `yaml:"dtos,omitempty" json:"dtos,omitempty" jsonschema:"description=Dtos are basically golang structs with some additional functionality which can be used for request/response actions"`

	// Entities are database-backed structures, with fields becoming both golang struct fields and database columns.
	Entities []*Module3Entity `yaml:"entities,omitempty" json:"entities,omitempty" jsonschema:"description=Entities are database-backed structures with fields becoming both golang struct fields and database columns."`

	// Complex custom data types definitions and location
	Complexes []EmiComplex `yaml:"complexes,omitempty" json:"complexes,omitempty" jsonschema:"description=Complex custom data types definitions and location"`

	// Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic
	Actions []*EmiAction `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic"`

	// Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here.
	Remotes []*EmiRemote `yaml:"remotes,omitempty" json:"remotes,omitempty" jsonschema:"description=Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here."`

	// Server side configuration manager, with limited data types, good option for casting .env files.
	Config []EmiConfig `yaml:"config,omitempty" json:"config,omitempty" jsonschema:"description=Server side configuration manager, with limited data types, good option for casting .env files."`

	// Manifests are a way to create actions bundle, to implement them easier, or wrap them into a single module.
	Manifests []EmiManifest `yaml:"manifests,omitempty" json:"manifests,omitempty" jsonschema:"description=Manifests are a way to create actions bundle, to implement them easier, or wrap them into a single module."`

	// Vsqls are typed virtual SQL queries. Each entry pairs a raw SQL string with a typed parameter struct that Emi generates.
	Vsqls []EmiVsql `yaml:"vsqls,omitempty" json:"vsqls,omitempty" jsonschema:"description=Typed virtual SQL queries. Each entry pairs a raw SQL string with a typed parameter struct that Emi generates."`

	// Templates is a scratch area of reusable shape definitions (dtos, actions)
	// that are NOT compiled into output files. They exist only to be referenced
	// by other parts of the module — most notably as capture sources.
	Templates *EmiTemplate `yaml:"templates,omitempty" json:"templates,omitempty" jsonschema:"description=Reusable shape definitions (dtos, actions) that are not compiled. Available as capture sources."`

	// Permissions define the tree of access control permissions this module contributes.
	// Each node's FullKey is resolved during preprocessing to parent.FullKey + "." + Key
	// when not set explicitly, then compiled into a usable Permissions constant/tree by
	// each target backend (see lib/golang and lib/js).
	Permissions []*EmiPermission `yaml:"permissions,omitempty" json:"permissions,omitempty" jsonschema:"description=Tree of access control permissions this module contributes. Compiled into a usable Permissions constant/tree per target language."`

	// SourcePath is the absolute path of the yaml file this module was read from, if
	// known (set by ReadEmiFromFile/StringToEmiWithPath - the plain content-only
	// loaders leave it empty). It is not part of the module's own definition, so it's
	// never (un)marshalled - it only exists to let preprocessing resolve relative
	// paths (e.g. EmiComplex.Include) against the file that declared them.
	SourcePath string `yaml:"-" json:"-"`
}

func (x *Emi) ActionsAsList() []string {
	items := []string{}
	for index, action := range x.Actions {
		items = append(items, fmt.Sprintf("%v", index)+" >>> "+action.Name+"("+action.Url+")")
	}

	return items
}

func (x *Emi) PublicName() string {
	return ToUpper(x.Name)
}
