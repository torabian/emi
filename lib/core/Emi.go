package core

import "fmt"

// Emi struct represents the entire file tree
type Emi struct {

	// Custom imports appened by some macros
	ActionsCustomImport []string `jsonschema:"-" json:"-" yaml:"-"`

	// Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty" jsonschema:"description=Represents where is the location of the module in app tree. Similar to PHP namespacing sytem it be used to explicitly as export path of the actions for client frameworks"`

	// Description of module and it's purpose. Used in code gen and creating documents.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of module and it's purpose. Used in code gen and creating documents."`

	// Version of the module. Helpful for different code generation phases but it's not necessary.
	Version string `yaml:"version,omitempty" json:"version,omitempty" jsonschema:"description=Version of the module. Helpful for different code generation phases but it's not necessary."`

	// Magic property for Emi EmiEmi.yml file. It's gonna be true only in a single file internally in Emi
	MetaWorkspace bool `yaml:"meta-workspace,omitempty" json:"meta-workspace,omitempty" jsonschema:"description=Magic property for Emi EmiEmi.yml file. It's gonna be true only in a single file internally in Emi"`

	// Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the module. Needs to be lower camel case and Module.go and Module.dyno.go will be generated based on this name."`

	// List of entities that module contains. Entities are basically tables in database with their mapping on golang and general actions generated for them
	Entities []EmiEntity `yaml:"entities,omitempty" json:"entities,omitempty" jsonschema:"description=List of entities that module contains. Entities are basically tables in database with their mapping on golang and general actions generated for them"`

	// Tasks are actions which are triggered by a queue message or a cron job.
	Tasks []*EmiTask `yaml:"tasks,omitempty" json:"tasks,omitempty" jsonschema:"description=Tasks are actions which are triggered by a queue message or a cron job."`

	// Dtos are basically golang structs with some additional functionality which can be used for request/response actions
	Dto []EmiDto `yaml:"dtos,omitempty" json:"dtos,omitempty" jsonschema:"description=Dtos are basically golang structs with some additional functionality which can be used for request/response actions"`

	// Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic
	Actions []*EmiAction `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions are similar to controllers in other frameworks. They are custom functionality available via CLI or Http requests and developer need to implement their logic"`

	// Macros are extra definition or templates which will modify the module and able to add extra fields or tables before the codegen occures.
	Macros []EmiMacro `yaml:"macros,omitempty" json:"macros,omitempty" jsonschema:"description=Macros are extra definition or templates which will modify the module and able to add extra fields or tables before the codegen occures."`

	// Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here.
	Remotes []*EmiRemote `yaml:"remotes,omitempty" json:"remotes,omitempty" jsonschema:"description=Remotes are definition of external services which could be contacted via http and Emi developer can make them typesafe by defining them here."`

	// Notifications are end-user messages, such as push notification, socket notification, and could be sent to user via different channels
	Notifications []*EmiNotification `yaml:"notifications,omitempty" json:"notifications,omitempty" jsonschema:"description=Notifications are end-user messages, such as push notification, socket notification, and could be sent to user via different channels"`

	// Events are internal changes that can be triggered by different sources
	Events []*EmiEvent `yaml:"events,omitempty" json:"events,omitempty" jsonschema:"description=Events are internal changes that can be triggered by different sources"`

	// Queries are set of SQL queries that developer writes and Emi generates tools for fetching them from database to golang code.
	Queries []*EmiQuery `yaml:"queries,omitempty" json:"queries,omitempty" jsonschema:"description=Queries are set of SQL queries that developer writes and Emi generates tools for fetching them from database to golang code."`

	// An interesting way of defining env variables
	Config []*EmiConfigField `yaml:"config,omitempty" json:"config,omitempty" jsonschema:"description=An interesting way of defining env variables"`

	// Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale.
	Messages EmiMessage `yaml:"messages,omitempty" json:"messages,omitempty" jsonschema:"description=Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale."`
}

func (x *Emi) ActionsAsList() []string {
	items := []string{}
	for index, action := range x.Actions {
		items = append(items, fmt.Sprintf("%v", index)+" >>> "+action.Name+"("+action.Url+")")
	}

	return items
}
