package core

import "fmt"

type Emi struct {

	// Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty" jsonschema:"description=Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks"`

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

	// Complex custom data types definitions and location
	Complexes []EmiComplex `yaml:"complexes,omitempty" json:"complexes,omitempty" jsonschema:"description=Complex custom data types definitions and location"`

	// Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic
	Actions []*EmiAction `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic"`

	// Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here.
	Remotes []*EmiRemote `yaml:"remotes,omitempty" json:"remotes,omitempty" jsonschema:"description=Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here."`
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
