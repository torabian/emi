package external

import (
	"reflect"
	"testing"

	"github.com/torabian/emi/emigo"
)

// Covers Entity1EntityUpdateInput: every field - even ones that are plain, always-
// present scalars on Entity1Entity itself - must be optional here, so a caller can tell
// gorm "this field wasn't touched" vs "explicitly set to the zero value." It carries no
// gorm tags of its own: it's not a model of its own table, just an input read back by
// the (future) Update function to build a partial change-set against Entity1Entity's
// own table.

func TestUpdateInput_PlainScalarBecomesNullable(t *testing.T) {
	// Title/IsActive/ViewCount/etc. are plain string/bool/int on Entity1Entity, but
	// must be Nullable here.
	typ := reflect.TypeOf(Entity1EntityUpdateInput{})

	cases := []struct {
		field    string
		wantType reflect.Type
	}{
		{"Title", reflect.TypeOf(emigo.Nullable[string]{})},
		{"IsActive", reflect.TypeOf(emigo.Nullable[bool]{})},
		{"ViewCount", reflect.TypeOf(emigo.Nullable[int]{})},
		{"Ratio64", reflect.TypeOf(emigo.Nullable[float64]{})},
		{"Status", reflect.TypeOf(emigo.Nullable[string]{})},
		{"Metadata", reflect.TypeOf(emigo.Nullable[map[string]string]{})},
		{"Labels", reflect.TypeOf(emigo.Nullable[[]string]{})},
	}

	for _, c := range cases {
		field, ok := typ.FieldByName(c.field)
		if !ok {
			t.Fatalf("expected Entity1EntityUpdateInput to have a %s field", c.field)
		}
		if field.Type != c.wantType {
			t.Fatalf("%s type = %v, want %v", c.field, field.Type, c.wantType)
		}
		if tag := field.Tag.Get("gorm"); tag != "" {
			t.Fatalf("%s should have no gorm tag of its own, got %q", c.field, tag)
		}
	}
}

func TestUpdateInput_AlreadyOptionalFieldsStayOptional(t *testing.T) {
	typ := reflect.TypeOf(Entity1EntityUpdateInput{})

	// Subtitle is already string? on the entity; should still be Nullable[string], not
	// double-wrapped.
	field, ok := typ.FieldByName("Subtitle")
	if !ok {
		t.Fatal("expected Entity1EntityUpdateInput to have a Subtitle field")
	}
	if field.Type != reflect.TypeOf(emigo.Nullable[string]{}) {
		t.Fatalf("Subtitle type = %v, want emigo.Nullable[string]", field.Type)
	}
}

func TestUpdateInput_RelationFieldsBecomeNullableVariant(t *testing.T) {
	typ := reflect.TypeOf(Entity1EntityUpdateInput{})

	cases := []struct {
		field    string
		wantType reflect.Type
	}{
		{"Items3", reflect.TypeOf(emigo.CollectionNullable[Entity2Entity]{})}, // collection -> collection?
		{"Owner", reflect.TypeOf(emigo.OneNullable[Entity2Entity]{})},         // one -> one?
	}

	for _, c := range cases {
		field, ok := typ.FieldByName(c.field)
		if !ok {
			t.Fatalf("expected Entity1EntityUpdateInput to have a %s field", c.field)
		}
		if field.Type != c.wantType {
			t.Fatalf("%s type = %v, want %v", c.field, field.Type, c.wantType)
		}
	}

	// Items/Items2 (array) reuse the *real* Entity1EntityItems/Entity1EntityItems2
	// child row types directly - not a separate, "doubly optional" clone - since
	// ReconcileHasMany persists each item with a plain Save(), which writes every
	// field; an item type with its own sub-fields individually optional would silently
	// null out anything the caller left unset on that one item.
	itemsField, ok := typ.FieldByName("Items")
	if !ok {
		t.Fatal("expected Entity1EntityUpdateInput to have an Items field")
	}
	if itemsField.Type != reflect.TypeOf(emigo.ArrayNullable[Entity1EntityItems]{}) {
		t.Fatalf("Items type = %v, want emigo.ArrayNullable[Entity1EntityItems]", itemsField.Type)
	}
}

func TestUpdateInput_ComplexBecomesPointer(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1EntityUpdateInput{}).FieldByName("Complex1")
	if !ok {
		t.Fatal("expected Entity1EntityUpdateInput to have a Complex1 field")
	}
	if field.Type != reflect.TypeOf(&Money{}) {
		t.Fatalf("Complex1 type = %v, want *Money (nil means untouched)", field.Type)
	}
}

func TestUpdateInput_IsSetDistinguishesUntouchedFromZeroValue(t *testing.T) {
	var input Entity1EntityUpdateInput

	if input.Title.IsSet() {
		t.Fatal("expected a zero-value Entity1EntityUpdateInput to have Title unset")
	}

	input.Title.Set(emigo.Ptr(""))
	if !input.Title.IsSet() {
		t.Fatal("expected Title to be set after Set(Ptr(\"\")), even though the value is the empty string")
	}
	if input.Title.OrDefault("untouched") != "" {
		t.Fatalf("expected the explicitly-set empty string to survive, got %q", input.Title.OrDefault("untouched"))
	}
}
