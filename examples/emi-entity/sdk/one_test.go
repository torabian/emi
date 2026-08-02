package external

import (
	"reflect"
	"testing"
)

// Covers the "one" / "one?" scenario: a one-to-one relation to another entity
// (Entity2Entity), backed by a foreign key column on this entity named after the field.
//
// ApplyEntityGormTags converts an entity's own one/one? fields into class/class? (see
// go-entity-gorm.go): the primary field itself becomes a real, gorm-native belongs-to
// (*Entity2Entity/Entity2Entity), tagged directly with the foreignKey relation - no
// gorm:"-", no hidden {field}Row shadow sibling needed, since a plain struct/pointer is
// exactly the shape gorm's own reflection-based schema builder recognizes. DTOs
// (Entity1Dto/Entity1OptionalDto) never see this conversion - they're built from
// entity.Fields before ApplyEntityGormTags ever mutates it, so they keep the portable,
// Operation-wrapped one/one? shape Create/Update's own request bodies still need.

func TestOneField_IsARealGormBelongsTo(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	field, ok := typ.FieldByName("Owner")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Owner field")
	}
	want := "foreignKey:OwnerId;references:Id"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Owner gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf(&Entity2Entity{}) {
		t.Fatalf("Owner field type = %v, want *Entity2Entity (one -> class, a pointer)", field.Type)
	}
	if got := field.Tag.Get("json"); got != "owner" {
		t.Fatalf("Owner json tag = %q, want %q (still the public DTO shape, just gorm-native now)", got, "owner")
	}

	idField, ok := typ.FieldByName("OwnerId")
	if !ok {
		t.Fatal("expected Entity1Entity to have an OwnerId field")
	}
	if idField.Type.Kind() != reflect.Int64 {
		t.Fatalf("OwnerId kind = %v, want int64 (references the target's real primary key, not its UniqueId)", idField.Type.Kind())
	}
	if got := idField.Tag.Get("json"); got != "-" {
		t.Fatalf("OwnerId json tag = %q, want %q", got, "-")
	}
}

func TestOneNullableField_IsARealGormBelongsTo(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	field, ok := typ.FieldByName("Manager")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Manager field")
	}
	want := "foreignKey:ManagerId;references:Id"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("Manager gorm tag = %q, want %q", got, want)
	}
	if field.Type != reflect.TypeOf(Entity2Entity{}) {
		t.Fatalf("Manager field type = %v, want Entity2Entity (one? -> class?, a value)", field.Type)
	}
}

func TestOneField_CliHelpers(t *testing.T) {
	flags := GetEntity1EntityCliFlags("")

	var ownerFlag, managerFlag *string
	for _, f := range flags {
		f := f
		switch f.Name {
		case "owner":
			ownerFlag = &f.Type
		case "manager":
			managerFlag = &f.Type
		}
	}

	if ownerFlag == nil || *ownerFlag != "class" {
		t.Fatalf("expected owner cli flag with type 'class' (one -> class on an entity), got %v", ownerFlag)
	}
	if managerFlag == nil || *managerFlag != "class?" {
		t.Fatalf("expected manager cli flag with type 'class?' (one? -> class? on an entity), got %v", managerFlag)
	}
}
