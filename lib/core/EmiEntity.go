package core

// Represents Entities in Emi. An entity in Emi is a table in database, with addition general
// features such as permissions, actions, security, and common actions which might be created or extra
// queries based on the type
type EmiEntity struct {

	// Type of the emi content
	Emi string `jsonschema:"description=Type of the emi content.;enum=entity" json:"emi" yaml:"emi"`

	// The replica configuration for the entity
	Replicas *EmiEntityReplicas `json:"replicas,omitempty" yaml:"replicas,omitempty" jsonschema:"The replica configuration for the entity"`

	// Notifications are end-user messages, such as push notification, socket notification, and could be sent to user via different channels
	Notifications []*EmiNotification `yaml:"notifications,omitempty" json:"notifications,omitempty" jsonschema:"description=Notifications are end-user messages, such as push notification, socket notification, and could be sent to user via different channels"`

	// Events are internal changes that can be triggered by different sources
	Events []*EmiEvent `yaml:"events,omitempty" json:"events,omitempty" jsonschema:"description=Events are internal changes that can be triggered by different sources"`

	// Modify the actions configuration, add headers, params and more to default generated actions and code
	Rpc EmiEntityActionConfig `yaml:"rpc,omitempty" json:"rpc,omitempty" jsonschema:"Modify the actions configuration, add headers, params and more to default generated actions and code."`

	// Rewrites the default permission generated value, for example if you want to regroup them somehow else.
	PremissionsRewrite *EmiEntityPermissionRewrite `yaml:"permRewrite,omitempty" json:"permRewrite,omitempty" jsonschema:"description=Rewrites the default permission generated value, for example if you want to regroup them somehow else."`

	// Extra permissions that an entity might need. You can add extra permissions that you will need in your
	// business logic related to entity in itself, to make it easier become as a group and document
	// later
	Permissions []EmiPermission `yaml:"permissions,omitempty" json:"permissions,omitempty" jsonschema:"description=Extra permissions that an entity might need. You can add extra permissions that you will need in your business logic related to entity in itself to make it easier become as a group and document later"`

	// Actions or extra actions (on top of default actions which automatically is generated) these are
	// the same actions that you can define for a module, but defining them on entity level make it easier
	// to relate them and group them. Also permission might be added automatically (need to clearify)
	Actions []*EmiAction `yaml:"actions,omitempty" json:"actions,omitempty" jsonschema:"description=Actions or extra actions (on top of default actions which automatically is generated) these are the same actions that you can define for a module but defining them on entity level make it easier to relate them and group them. Also permission might be added automatically (need to clearify)"`

	// The entity name is crucial as it determines database table names and is used by Emi's Go and code generation tools; note that changing an entity name does not delete previously created entities requiring manual file deletion and only camelCase naming is supported.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=The entity name is crucial as it determines database table names and is used by Emi's Go and code generation tools; note that changing an entity name does not delete previously created entities requiring manual file deletion and only camelCase naming is supported."`

	// You can make sure there is only one record of the entity per user or workspace using this option.
	// for example, if you want only one credit card per workspace, you can set distinctBy: workspace
	// and it will do the job
	DistinctBy string `yaml:"distinctBy,omitempty" json:"distinctBy,omitempty" jsonschema:"enum=workspace,enum=user,description=You can ensure there is only one record of the entity per user or workspace using this option for example if you want only one credit card per workspace set distinctBy: workspace and it will do the job"`

	// Customize the features generated for entity, less common  changes goes to this object
	Features EmiEntityFeatures `yaml:"features,omitempty" json:"features,omitempty" jsonschema:"description=Customize the features generated for entity, less common  changes goes to this object"`

	// Changes the default table name based on project prefix and entity name useful for times that you want to connect project to an existing database
	Table string `yaml:"table,omitempty" json:"table,omitempty" jsonschema:"description=Changes the default table name based on project prefix and entity name useful for times that you want to connect project to an existing database"`

	// Use fields allows you to customize the entity default generated fields.
	UseFields *EmiDataFields `yaml:"useFields,omitempty" json:"useFields,omitempty" jsonschema:"description=Use fields allows you to customize the entity default generated fields."`

	// Manages the entity models
	SecurityModel *EntitySecurityModel `yaml:"security,omitempty" json:"security,omitempty" jsonschema:"description=Manages the entity models"`

	// Adds a golang code to the geenrated code in very top location of the file after imports and before any code.
	PrependScript string `yaml:"prependScript,omitempty" json:"prependScript,omitempty" jsonschema:"description=Adds a golang code to the geenrated code in very top location of the file after imports and before any code."`

	// Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale.
	Messages EmiMessage `yaml:"messages,omitempty" json:"messages,omitempty" jsonschema:"description=Messages are translatable strings which will be used as errors and other types of messages and become automatically picked via user locale."`

	// Adds a extra code before the create action in the entity. This is pure golang code.
	// Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity.
	PrependCreateScript string `yaml:"prependCreateScript,omitempty" json:"prependCreateScript,omitempty" jsonschema:"description=Adds a extra code before the create action in the entity. This is pure golang code. Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity."`

	// Adds a extra code before the update action in the entity. This is pure golang code.
	// Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity.
	PrependUpdateScript string `yaml:"prependUpdateScript,omitempty" json:"prependUpdateScript,omitempty" jsonschema:"description=Adds a extra code before the update action in the entity. This is pure golang code. Use it with caution such meta codes make module unreadable overtime. You can add script on non-dyno file of the entity."`

	// Access is a method of limiting which type offunctionality will be created for the entity. For example access read will remove all create functionality from code and public API.
	Access string `yaml:"access,omitempty" json:"access,omitempty" jsonschema:"description=Access is a method of limiting which type offunctionality will be created for the entity. For example access read will remove all create functionality from code and public API."`

	// For entities, if the query scope is public the query action will become automatically public and without authentication
	QueryScope string `yaml:"queryScope,omitempty" json:"queryScope,omitempty" jsonschema:"enum=public,enum=specific,description=For entities, if the query scope is public the query action will become automatically public and without authentication"`

	// A list of extra queries that Emi can generate for the the entity. Emi might offer some extra queries to be generated so they will be listed here.
	Queries []string `yaml:"queries,omitempty" json:"queries,omitempty" jsonschema:"description=A list of extra queries that Emi can generate for the the entity. Emi might offer some extra queries to be generated so they will be listed here."`

	// Override the some default Emi generated fields gorm configuration.
	GormMap GormOverrideMap `yaml:"gormMap,omitempty" json:"gormMap,omitempty" jsonschema:"description=Override the some default Emi generated fields gorm configuration."`

	// Define the fields that this entity will have both in golang and database columns.
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Define the fields that this entity will have both in golang and database columns."`

	// The name of the entity which will appear in CLI. By default the name of the entity will be used with dashes.
	CliName string `yaml:"cliName,omitempty" json:"cliName,omitempty" jsonschema:"description=The name of the entity which will appear in CLI. By default the name of the entity will be used with dashes."`

	// The alternative shortcut in the CLI. By default it's empty and only the entity name or CliName.
	CliShort string `yaml:"cliShort,omitempty" json:"cliShort,omitempty" jsonschema:"description=The alternative shortcut in the CLI. By default it's empty and only the entity name or CliName."`

	// Description about the purpose of the entity. It will be used in CLI and codegen documentation.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description about the purpose of the entity. It will be used in CLI and codegen documentation."`

	// CTE is a common recursive feature of an entity; enabling it generates SQL for recursive parent-child CTE queries and makes it available in Golang.
	Cte bool `yaml:"cte,omitempty" json:"cte,omitempty" jsonschema:"description=CTE is a common recursive feature of an entity; enabling it generates SQL for recursive parent-child CTE queries and makes it available in Golang."`

	// The name of the golang function which will recieve entity pointer to make some modification
	// upon query, get or other details.
	PostFormatter string `yaml:"postFormatter,omitempty" json:"postFormatter,omitempty" jsonschema:"description=The name of the golang function which will recieve entity pointer to make some modification upon query, get or other details."`

	// Internal metadata for code generation.
	RootModule *Emi `yaml:"-" json:"-" jsonschema:"-"`
}

func (x *EmiEntity) DataFields() EmiDataFields {
	data := EmiDataFields{}

	if x.UseFields == nil {
		data = EmiDataFields{
			Essentials:       true,
			PrimaryId:        true,
			NumericTimestamp: false,
			DateTimestamp:    true,
		}

		return data
	} else {
		data = *x.UseFields
	}

	return data
}

func (x EmiEntity) HasClickHouse() bool {
	return x.Replicas != nil && x.Replicas.Clickhouse != nil
}

func (x EmiEntity) GetClassName() string {
	return ToUpper(x.Name) + "Entity"
}
