package external

import (
	"reflect"
	"testing"
)

// Covers the "collection" / "collection?" scenario: a many-to-many relation to another
// entity (Entity2Entity), backed by a join table named after this entity and the field.
//
// The DTO field itself (Items3, emigo.Collection[T]) is tagged gorm:"-" for the same
// reason array's DTO field is: it's a PATCH-payload wrapper, not a gorm-recognizable
// association shape. A real []*Entity2Entity sibling field (Items3Row) carries the
// actual many2many relation tag.

func TestCollectionField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items3")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items3 field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Items3 gorm tag = %q, want %q", got, "-")
	}
}

func TestCollectionNullableField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items4")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items4 field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Items4 gorm tag = %q, want %q", got, "-")
	}
}

func TestCollectionField_RowSiblingCarriesMany2ManyTag(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items3Row")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items3Row field")
	}

	want := "many2many:entity1_items3;foreignKey:Id;references:Id"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Items3Row gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf([]*Entity2Entity{}) {
		t.Fatalf("Items3Row field type = %v, want []*Entity2Entity", field.Type)
	}
	if got := field.Tag.Get("json"); got != "-" {
		t.Fatalf("Items3Row json tag = %q, want %q", got, "-")
	}
}

func TestCollectionNullableField_RowSiblingCarriesMany2ManyTag(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Items4Row")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Items4Row field")
	}

	want := "many2many:entity1_items4;foreignKey:Id;references:Id"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Items4Row gorm tag = %q, want %q", got, want)
	}
}

func TestCollectionField_TargetEntityAndCliHelpers(t *testing.T) {
	target := Entity2Entity{Label: "tag-a"}
	if target.Label != "tag-a" {
		t.Fatalf("unexpected target entity value: %+v", target)
	}

	flags := GetEntity2EntityCliFlags("")
	if len(flags) != 3 {
		t.Fatalf("expected 3 cli flags (id, uniqueId, label), got %+v", flags)
	}
}
