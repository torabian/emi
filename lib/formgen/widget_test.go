package formgen

import (
	"testing"

	"github.com/torabian/emi/lib/core"
)

// TestAllCatalogFieldTypesResolve is the enforcement mentioned in widget.go's
// doc comment: every field type core itself claims to support
// (core.GetEmiFieldTypeCatalog) must resolve to a real widget here. If this
// test fails after adding a new EmiField type, ResolveWidgetKind is missing a
// case for it.
func TestAllCatalogFieldTypesResolve(t *testing.T) {
	catalog := core.GetEmiFieldTypeCatalog()

	all := append([]core.FieldType{}, catalog.DtoFieldTypes...)
	all = append(all, catalog.DtoNullableFieldTypes...)

	for _, ft := range all {
		field := &core.EmiField{Name: "x", Type: ft}
		got := ResolveWidgetKind(field)
		if got == WidgetUnknown {
			t.Errorf("field type %q resolves to WidgetUnknown - add a case in ResolveWidgetKind", ft)
		}
	}
}

func TestResolveWidgetKindComplexOverridesType(t *testing.T) {
	field := &core.EmiField{Name: "x", Type: core.FieldTypeString, Complex: "geo.Point"}
	if got := ResolveWidgetKind(field); got != WidgetJson {
		t.Errorf("expected complex field to resolve to WidgetJson, got %q", got)
	}
}

func TestResolveWidgetKindNil(t *testing.T) {
	if got := ResolveWidgetKind(nil); got != WidgetUnknown {
		t.Errorf("expected nil field to resolve to WidgetUnknown, got %q", got)
	}
}
