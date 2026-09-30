package core

// EmiInterface is a named set of fields shared by several dtos. A dto (or entity) that
// lists it under `implements` gets the fields included, and the target languages that
// support it also generate an interface type every implementing dto satisfies (Go: a Go
// interface of Get<Field>() accessors, TypeScript: an `interface`), so a function can accept
// any of them.
//
// Interface fields must be self-contained: no inline object/array children, since an
// inline child type belongs to one dto and can't be shared. Reference a declared
// target instead (with `module`/`provider`/`jsProvider` when it lives elsewhere).
type EmiInterface struct {
	// Name of the interface, lower camel case. The generated type is its PascalCase form.
	Name string `yaml:"name,omitempty" json:"name,omitempty" jsonschema:"description=Name of the interface in camel case. The generated type is the PascalCase form (purchasable -> Purchasable)."`

	// Description of what implementing this interface means. Used in generated docs.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=What implementing this interface means. Used in generated documentation."`

	// Fields every implementer gets. Same shape as dto fields, and they can build
	// themselves from a field template with `use`.
	Fields []*EmiField `yaml:"fields,omitempty" json:"fields,omitempty" jsonschema:"description=Fields every implementing dto gets. Same shape as dto fields."`

	// Context is where the fields' `use:` templates resolve: "dto" (default) or "entity".
	// A field built from a template only fits owners of that same kind, because a
	// template can resolve to a different type per context (PurchasableDto vs
	// PurchasableEntity). Plain fields fit any owner.
	Context string `yaml:"context,omitempty" json:"context,omitempty" jsonschema:"enum=dto,enum=entity,description=Where the fields' use: templates resolve. Fields built from a template only fit owners of the same kind. Defaults to dto."`

	// Mutable also generates Set<Field>() in the interface. By default it only has getters,
	// which is all most consumers need and keeps the contract easy to satisfy.
	Mutable bool `yaml:"mutable,omitempty" json:"mutable,omitempty" jsonschema:"description=Also generate setters in the interface. Getters only by default."`

	// Module and Provider mark an interface as declared in another Go package: the
	// package alias and import path, same meaning as on relation fields. When Provider is
	// set, Go does not generate the interface type here - it imports it and only asserts
	// that the implementing dtos satisfy it.
	Module   string `yaml:"module,omitempty" json:"module,omitempty" jsonschema:"description=Go package alias of the package declaring the interface (like module on a relation field)."`
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty" jsonschema:"description=Go import path of the package declaring the interface. When set the interface is referenced instead of generated."`

	// JsProvider is the import path of the generated TypeScript file declaring the
	// interface, including the file name (e.g. @/modules/clicktobuy/sdk/Purchasable). When
	// set, TypeScript imports the interface instead of generating it.
	JsProvider string `yaml:"jsProvider,omitempty" json:"jsProvider,omitempty" jsonschema:"description=Import path (including the file name) of the TypeScript file declaring the interface. When set the interface is imported instead of generated."`

	// templated records which of the fields were built from a `use:` template, for the
	// context check when an owner includes them. Not part of the definition.
	templated map[string]bool
}

// GetClassName is the generated type name: the PascalCase form of Name.
func (x *EmiInterface) GetClassName() string {
	return ToUpper(x.Name)
}

// IsGoReference reports whether Go should import this interface instead of generating it.
func (x *EmiInterface) IsGoReference() bool {
	return x.Provider != ""
}

// IsJsReference reports whether TypeScript should import this interface instead of generating it.
func (x *EmiInterface) IsJsReference() bool {
	return x.JsProvider != ""
}

// FindInterfaces returns the interfaces of m matching names, in the order of names.
// Preprocessing has already rejected unknown names.
func (m *Emi) FindInterfaces(names []string) []*EmiInterface {
	found := make([]*EmiInterface, 0, len(names))
	for _, name := range names {
		for i := range m.Interfaces {
			if m.Interfaces[i].Name == name {
				found = append(found, &m.Interfaces[i])
				break
			}
		}
	}
	return found
}
