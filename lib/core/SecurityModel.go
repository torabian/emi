package core

// Used for actions generally
type SecurityModel struct {
	// Only users which belong to root and actively selected the root workspace can
	// write to this entity from Emi default functionality
	AllowOnRoot bool `json:"allowOnRoot,omitempty" yaml:"allowOnRoot,omitempty"`

	// Set of permissions which are required for this service.
	ActionRequires []PermissionInfo `json:"requires,omitempty" yaml:"requires,omitempty"`

	// Resolve strategy is by default on the workspace, you can change it by user
	// also. Be sure of the consequences
	ResolveStrategy string `json:"resolveStrategy,omitempty" yaml:"resolveStrategy,omitempty"`
}
