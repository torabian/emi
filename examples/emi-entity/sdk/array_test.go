package external

import (
	"reflect"
	"testing"
)

// Covers the "array" / "array?" scenario on an entity: a has-many relation where the
// child rows are the auto-generated nested struct (Entity1EntityItems /
// Entity1EntityItems2), linked back to the parent via a generic LinkerId column.
//
// ApplyEntityGormTags converts an entity's own array/array? fields into _list/_list?
// (see go-entity-gorm.go): the primary field itself becomes a real, gorm-native
// []*ChildStruct/[]ChildStruct has-many, tagged directly with the foreignKey relation -
// no gorm:"-", no hidden {field}Row shadow sibling needed, since a plain slice is
// exactly the shape gorm's own reflection-based schema builder recognizes. DTOs
// (Entity1Dto/Entity1OptionalDto) never see this conversion - they're built from
// entity.Fields before ApplyEntityGormTags ever mutates it, so they keep the portable,
// Operation-wrapped array/array? shape Create/Update's own request bodies still need.

func TestArrayField_IsARealGormHasMany(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items field")
	}

	want := "foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Items gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf([]*Entity1EntityItems{}) {
		t.Fatalf("Items field type = %v, want []*Entity1EntityItems (array -> _list, a pointer slice)", field.Type)
	}
	if got := field.Tag.Get("json"); got != "items" {
		t.Fatalf("Items json tag = %q, want %q (still the public DTO shape, just gorm-native now)", got, "items")
	}
}

func TestArrayNullableField_IsARealGormHasMany(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items2")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items2 field")
	}

	want := "foreignKey:LinkerId;references:Id;constraint:OnDelete:CASCADE"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Items2 gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf([]Entity1EntityItems2{}) {
		t.Fatalf("Items2 field type = %v, want []Entity1EntityItems2 (array? -> _list?, a value slice)", field.Type)
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
