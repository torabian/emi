package core

// Used for defining the entity overall action permissions
type EntitySecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can write to this entity from Emi default functionality. Read mechanism won't be affected.
	WriteOnRoot *bool `json:"writeOnRoot,omitempty" yaml:"writeOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can write to this entity from Emi default functionality. Read mechanism won't be affected."`

	// Only users which belong to root and actively selected the root workspace can read from entity from Emi default functionality. Write mechanism is not affected.
	ReadOnRoot *bool `json:"readOnRoot,omitempty" yaml:"readOnRoot,omitempty" jsonschema:"description=Only users which belong to root and actively selected the root workspace can read from entity from Emi default functionality. Write mechanism is not affected."`

	// Resolve strategy means that the content belongs either to workspace or user. It affects the query.
	ResolveStrategy *string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty" jsonschema:"enum=workspace,enum=user, description=Resolve strategy means that the content belongs either to workspace or user. It affects the query."`
}
