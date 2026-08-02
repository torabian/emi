package external

import (
	"reflect"
	"testing"
)

// Covers the "one" / "one?" scenario: a one-to-one relation to another entity
// (Entity2Entity), backed by a foreign key column on this entity named after the field.
//
// The DTO field itself (Owner, emigo.One[T]) is tagged gorm:"-" for the same reason
// array/collection's DTO fields are. Two real gorm-shaped siblings carry the actual
// "belongs to" relation: {field}Id (the FK column) and {field}Row (*Entity2Entity).

func TestOneField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Owner")
	if !ok {
		t.Fatal("expected Entity1Entity to have an Owner field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Owner gorm tag = %q, want %q", got, "-")
	}
}

func TestOneNullableField_DtoFieldSkipsGorm(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Manager")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Manager field")
	}

	if got := field.Tag.Get("gorm"); got != "-" {
		t.Fatalf("Manager gorm tag = %q, want %q", got, "-")
	}
}

func TestOneField_IdAndRowSiblingsCarryBelongsToTag(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

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

	rowField, ok := typ.FieldByName("OwnerRow")
	if !ok {
		t.Fatal("expected Entity1Entity to have an OwnerRow field")
	}
	want := "foreignKey:OwnerId;references:Id"
	if got := rowField.Tag.Get("gorm"); got != want {
		t.Fatalf("OwnerRow gorm tag = %q, want %q", got, want)
	}
	if rowField.Type != reflect.TypeOf(&Entity2Entity{}) {
		t.Fatalf("OwnerRow field type = %v, want *Entity2Entity", rowField.Type)
	}
}

func TestOneNullableField_IdAndRowSiblingsCarryBelongsToTag(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	rowField, ok := typ.FieldByName("ManagerRow")
	if !ok {
		t.Fatal("expected Entity1Entity to have a ManagerRow field")
	}
	want := "foreignKey:ManagerId;references:Id"
	if got := rowField.Tag.Get("gorm"); got != want {
		t.Fatalf("ManagerRow gorm tag = %q, want %q", got, want)
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

	if ownerFlag == nil || *ownerFlag != "one" {
		t.Fatalf("expected owner cli flag with type 'one', got %v", ownerFlag)
	}
	if managerFlag == nil || *managerFlag != "one?" {
		t.Fatalf("expected manager cli flag with type 'one?', got %v", managerFlag)
	}
}
