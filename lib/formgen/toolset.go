package formgen

// WidgetComponent describes, for one WidgetKind, which concrete UI component
// a renderer should emit and where it comes from. The shape is intentionally
// generic (component identifier + import location) so it fits a React named
// import today and a Compose @Composable function or SwiftUI View later,
// without this package knowing anything about any of them.
type WidgetComponent struct {
	// Component is the identifier the writer emits, e.g. "TextField".
	Component string

	// ImportFrom is the module/package path Component is imported from. Empty
	// means "already in scope" (e.g. a built-in), useful for renderers that
	// don't need an import statement at all.
	ImportFrom string

	// Named selects `import { Component } from ImportFrom` (true) vs a default
	// import `import Component from ImportFrom` (false). Renderers that have no
	// notion of named/default imports (Kotlin, Swift) can ignore this field.
	Named bool
}

// Toolset is a renderer's full configuration: one WidgetComponent per
// WidgetKind it supports, plus how to resolve `one`/`collection` relation
// fields, which point at another entity this compiler was never given the
// fields of. Build a Toolset with NewToolset and register widgets with
// Toolset.Set - see lib/js/react-form-toolset.go for the default React one.
type Toolset struct {
	widgets map[WidgetKind]WidgetComponent

	// RelationComponent names the picker component for a one/collection
	// field given its EmiField.Target (e.g. "customer" -> "CustomerPicker").
	RelationComponent func(target string) string

	// RelationImportFrom builds the import path for that picker component,
	// given the field's Target and Module (Module is only set when the target
	// lives in a different Emi module - see EmiField.Module).
	RelationImportFrom func(target, module string) string
}

// NewToolset returns an empty Toolset ready for Set calls.
func NewToolset() Toolset {
	return Toolset{widgets: map[WidgetKind]WidgetComponent{}}
}

// Set registers the component a renderer should use for a given WidgetKind.
func (t *Toolset) Set(kind WidgetKind, component WidgetComponent) {
	if t.widgets == nil {
		t.widgets = map[WidgetKind]WidgetComponent{}
	}
	t.widgets[kind] = component
}

// Resolve looks up the component configured for kind. ok is false when the
// renderer never registered anything for that kind - callers should fall
// back to their own default (typically the WidgetJson component) rather than
// silently drop the field.
func (t Toolset) Resolve(kind WidgetKind) (WidgetComponent, bool) {
	w, ok := t.widgets[kind]
	return w, ok
}
