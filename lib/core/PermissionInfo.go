package core

type PermissionInfo struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Description string `yaml:"description,omitempty" json:"description,omitempty"`
	CompleteKey string `yaml:"completeKey,omitempty" json:"completeKey,omitempty"`
	GoVariable  string `yaml:"-" json:"-"`
}
