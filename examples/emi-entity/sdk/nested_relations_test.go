package external

import (
	"reflect"
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

// Covers relations (array/one) nested inside object/object? containers, at any depth -
// e.g. nestedContainer.nestedInner.nestedItems (array) and .nestedOwner (one). Two bugs
// had to be fixed together for this to work correctly:
//
//  1. ApplyEntityGormTags computed a nested array's child struct type name using a flat
//     "{Entity}{Field}" formula, ignoring how deeply the field was actually nested -
//     producing either an undefined type reference (compile error) or, worse, a
//     reference to some unrelated top-level type of the same field name (silently
//     wrong, still compiled). Fixed by threading the same accumulating class-name
//     prefix GoCommonStructGenerator's own nested-struct naming uses.
//  2. object? (nullable object) fields render as emigo.Nullable[T] on the entity
//     itself, and gorm's "embedded" tag only ever sees the *wrapper's* own fields
//     (unexported), never T's - so a relation (or even a plain scalar) declared inside
//     an object? container was invisible to gorm entirely, regardless of nesting.
//     Fixed by giving object? fields the same "gorm:-" DTO field + hidden gorm-visible
//     {field}Row *T sibling every other unsupported relation shape gets.
func TestNestedRelations_ArrayInsidePlainObjectContainer(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	nestedContainer, ok := typ.FieldByName("NestedContainer")
	if !ok {
		t.Fatal("expected Entity1Entity to have a NestedContainer field")
	}
	if got := nestedContainer.Tag.Get("gorm"); got != "embedded" {
		t.Fatalf("NestedContainer gorm tag = %q, want %q (plain object embeds directly)", got, "embedded")
	}

	// NestedItemsRow/NestedOwnerId/NestedOwnerRow are real Go fields nested inside
	// NestedContainer.NestedInner - not Go-anonymous-embedded, so reflect.FieldByName
	// on Entity1Entity itself won't find them; gorm's own "embedded" tag handles that
	// flattening for its own schema, independently of Go's language-level promotion.
	nestedInner := nestedContainer.Type.FieldByIndex([]int{0}) // NestedInner is the only field
	if nestedInner.Name != "NestedInner" {
		t.Fatalf("expected NestedContainer's first field to be NestedInner, got %s", nestedInner.Name)
	}

	nestedItemsRow, ok := nestedInner.Type.FieldByName("NestedItemsRow")
	if !ok {
		t.Fatal("expected NestedContainer.NestedInner to have a NestedItemsRow field (hidden has-many sibling)")
	}
	want := "foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE"
	if got := nestedItemsRow.Tag.Get("gorm"); got != want {
		t.Fatalf("NestedItemsRow gorm tag = %q, want %q", got, want)
	}
	if nestedItemsRow.Type != reflect.TypeOf([]*Entity1EntityNestedContainerNestedInnerNestedItems{}) {
		t.Fatalf("NestedItemsRow type = %v, want []*Entity1EntityNestedContainerNestedInnerNestedItems (the *real*, fully-prefixed nested struct)", nestedItemsRow.Type)
	}

	nestedOwnerId, ok := nestedInner.Type.FieldByName("NestedOwnerId")
	if !ok {
		t.Fatal("expected NestedContainer.NestedInner to have a NestedOwnerId field")
	}
	if nestedOwnerId.Type.Kind() != reflect.Int64 {
		t.Fatalf("NestedOwnerId kind = %v, want int64", nestedOwnerId.Type.Kind())
	}
	nestedOwnerRow, ok := nestedInner.Type.FieldByName("NestedOwnerRow")
	if !ok {
		t.Fatal("expected NestedContainer.NestedInner to have a NestedOwnerRow field")
	}
	if got := nestedOwnerRow.Tag.Get("gorm"); got != "foreignKey:NestedOwnerId;references:Id" {
		t.Fatalf("NestedOwnerRow gorm tag = %q", got)
	}
}

func TestNestedRelations_ArrayInsideObjectNullableContainer(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	nestedContainerOpt, ok := typ.FieldByName("NestedContainerOpt")
	if !ok {
		t.Fatal("expected Entity1Entity to have a NestedContainerOpt field")
	}
	if got := nestedContainerOpt.Tag.Get("gorm"); got != "-" {
		t.Fatalf("NestedContainerOpt gorm tag = %q, want %q (object? can't embed directly)", got, "-")
	}

	rowField, ok := typ.FieldByName("NestedContainerOptRow")
	if !ok {
		t.Fatal("expected Entity1Entity to have a NestedContainerOptRow field (hidden gorm-visible sibling)")
	}
	if got := rowField.Tag.Get("gorm"); got != "embedded" {
		t.Fatalf("NestedContainerOptRow gorm tag = %q, want %q", got, "embedded")
	}
	if rowField.Type != reflect.TypeOf(&Entity1EntityNestedContainerOpt{}) {
		t.Fatalf("NestedContainerOptRow type = %v, want *Entity1EntityNestedContainerOpt", rowField.Type)
	}
}

// Content2 (object?) itself has no relation inside it - just a plain scalar (item2) -
// which used to be entirely invisible to gorm the same way a nested relation would be.
func TestNestedRelations_ObjectNullablePlainScalarIsGormVisible(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	content2, ok := typ.FieldByName("Content2")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Content2 field")
	}
	if got := content2.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Content2 gorm tag = %q, want %q", got, "-")
	}

	content2Row, ok := typ.FieldByName("Content2Row")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Content2Row field")
	}
	if got := content2Row.Tag.Get("gorm"); got != "embedded" {
		t.Fatalf("Content2Row gorm tag = %q, want %q", got, "embedded")
	}
	if content2Row.Type != reflect.TypeOf(&Entity1EntityContent2{}) {
		t.Fatalf("Content2Row type = %v, want *Entity1EntityContent2", content2Row.Type)
	}
}

// End-to-end confirmation via gorm's own schema parser (no live database needed): every
// nested relation resolves to a real, correctly-typed association, and the previously-
// invisible object?-nested field/relation are now part of the parsed schema.
func TestNestedRelations_GormSchemaRecognizesEverything(t *testing.T) {
	s, err := schema.Parse(&Entity1Entity{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		t.Fatalf("schema.Parse error: %v", err)
	}

	wantRelations := map[string]schema.RelationshipType{
		"NestedItemsRow":    schema.HasMany,
		"NestedOwnerRow":    schema.BelongsTo,
		"NestedItemsOptRow": schema.HasMany,
	}
	for name, wantType := range wantRelations {
		rel, ok := s.Relationships.Relations[name]
		if !ok {
			t.Fatalf("expected a %q relationship, gorm didn't recognize one", name)
		}
		if rel.Type != wantType {
			t.Fatalf("%s relationship type = %v, want %v", name, rel.Type, wantType)
		}
	}

	foundItem2 := false
	for _, f := range s.Fields {
		if f.Name == "Item2" {
			foundItem2 = true
		}
	}
	if !foundItem2 {
		t.Fatal("expected Content2's Item2 field to be part of the parsed schema")
	}
}
