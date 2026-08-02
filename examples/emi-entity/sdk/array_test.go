package external

import (
	"reflect"
	"testing"
)

// Covers the "array" / "array?" scenario: a has-many relation where the child rows are
// the auto-generated nested struct (Entity1EntityItems / Entity1EntityItems2), linked
// back to the parent via a generic LinkerId column.
//
// The DTO field itself (Items, emigo.Array[T]) can't be migrated by gorm directly - it's
// a PATCH-payload wrapper with no Scan/Value and isn't the []*T shape gorm recognizes
// as an association - so it's tagged gorm:"-", and a real gorm-shaped sibling field
// (ItemsRow []*Entity1EntityItems) carries the actual has-many relation tag instead.

func TestArrayField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Items gorm tag = %q, want %q", got, "-")
	}

	if field.Type.String() != "emigo.Array[github.com/torabian/emi/examples/emi-entity/sdk.Entity1EntityItems]" {
		t.Fatalf("Items field type = %v, want emigo.Array[Entity1EntityItems]", field.Type)
	}
}

func TestArrayNullableField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items2")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items2 field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Items2 gorm tag = %q, want %q", got, "-")
	}
}

func TestArrayField_RowSiblingCarriesHasManyTag(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("ItemsRow")
	if !ok {
		t.Fatal("expected Entity1Entity to have an ItemsRow field")
	}

	want := "foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("ItemsRow gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf([]*Entity1EntityItems{}) {
		t.Fatalf("ItemsRow field type = %v, want []*Entity1EntityItems", field.Type)
	}
	if got := field.Tag.Get("json"); got != "-" {
		t.Fatalf("ItemsRow json tag = %q, want %q (gorm-only, not part of the public DTO shape)", got, "-")
	}
}

func TestArrayField_ChildStructHasLinkerIdAndCliHelpers(t *testing.T) {
	items := Entity1EntityItems{Item2: "hello", LinkerId: 42}
	if items.Item2 != "hello" || items.LinkerId != 42 {
		t.Fatalf("unexpected child struct value: %+v", items)
	}

	field, ok := reflect.TypeOf(Entity1EntityItems{}).FieldByName("LinkerId")
	if !ok {
		t.Fatal("expected Entity1EntityItems to have a LinkerId field")
	}
	if field.Type.Kind() != reflect.Int64 {
		t.Fatalf("LinkerId kind = %v, want int64 (references the parent's real primary key, not its UniqueId)", field.Type.Kind())
	}
	if got := field.Tag.Get("gorm"); got != "index" {
		t.Fatalf("LinkerId gorm tag = %q, want %q", got, "index")
	}

	// item2, id, uniqueId, linkerId
	flags := GetEntity1EntityItemsCliFlags("")
	if len(flags) != 4 {
		t.Fatalf("expected 4 cli flags (item2, id, uniqueId, linkerId), got %+v", flags)
	}
}
