package external

import (
	"reflect"
	"testing"

	"github.com/torabian/emi/emigo"
)

// Covers the remaining emi field types from the module3-json-schema enum that aren't
// already exercised by array_test.go/collection_test.go/one_test.go/complex_test.go:
// string, string?, bool, bool?, int, int?, int32, int32?, int64, int64?, float32,
// float32?, float64, float64?, enum, enum?, map, map?, slice, slice?, any.
// (object/object? are also covered here, alongside array/one/collection, in the earlier
// test files.)

// Plain scalars and their nullable counterparts need no gorm tag at all: the "?"
// variants render as emigo.Nullable[T], which already implements Value()/Scan() for
// every primitive T, and the non-nullable variants are native Go types gorm already
// knows how to store.
func TestScalarFields_NoGormTagNeeded(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	cases := []struct {
		field    string
		wantType reflect.Type
	}{
		{"Title", reflect.TypeOf("")},
		{"Subtitle", reflect.TypeOf(emigo.Nullable[string]{})},
		{"IsActive", reflect.TypeOf(false)},
		{"IsFeatured", reflect.TypeOf(emigo.Nullable[bool]{})},
		{"ViewCount", reflect.TypeOf(int(0))},
		{"ViewCountOpt", reflect.TypeOf(emigo.Nullable[int]{})},
		{"SmallCount", reflect.TypeOf(int32(0))},
		{"SmallCountOpt", reflect.TypeOf(emigo.Nullable[int32]{})},
		{"BigCount", reflect.TypeOf(int64(0))},
		{"BigCountOpt", reflect.TypeOf(emigo.Nullable[int64]{})},
		{"Ratio32", reflect.TypeOf(float32(0))},
		{"Ratio32Opt", reflect.TypeOf(emigo.Nullable[float32]{})},
		{"Ratio64", reflect.TypeOf(float64(0))},
		{"Ratio64Opt", reflect.TypeOf(emigo.Nullable[float64]{})},
	}

	for _, c := range cases {
		field, ok := typ.FieldByName(c.field)
		if !ok {
			t.Fatalf("expected Entity1Entity to have a %s field", c.field)
		}
		if field.Type != c.wantType {
			t.Fatalf("%s type = %v, want %v", c.field, field.Type, c.wantType)
		}
		if tag := field.Tag.Get("gorm"); tag != "" {
			t.Fatalf("%s should not need a gorm tag, got %q", c.field, tag)
		}
	}
}

// enum/enum? are explicitly called out because they're conceptually a closed set of
// values, but they still need to physically land in the database as a string column -
// which is exactly what they already reduce to (no wrapper type of their own).
func TestEnumFields_ResolveToString(t *testing.T) {
	typ := reflect.TypeOf(Entity1Entity{})

	statusField, ok := typ.FieldByName("Status")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Status field")
	}
	if statusField.Type.Kind() != reflect.String {
		t.Fatalf("Status kind = %v, want string", statusField.Type.Kind())
	}
	if tag := statusField.Tag.Get("gorm"); tag != "" {
		t.Fatalf("Status should not need a gorm tag, got %q", tag)
	}

	statusOptField, ok := typ.FieldByName("StatusOpt")
	if !ok {
		t.Fatal("expected Entity1Entity to have a StatusOpt field")
	}
	if statusOptField.Type != reflect.TypeOf(emigo.Nullable[string]{}) {
		t.Fatalf("StatusOpt type = %v, want emigo.Nullable[string]", statusOptField.Type)
	}
}

// map is a raw Go map with no Scan/Value of its own, so it needs an explicit
// serializer:json tag or gorm won't know how to persist it.
func TestMapField_NeedsJsonSerializer(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Metadata")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Metadata field")
	}
	if field.Type != reflect.TypeOf(map[string]string{}) {
		t.Fatalf("Metadata type = %v, want map[string]string", field.Type)
	}
	if got := field.Tag.Get("gorm"); got != "serializer:json" {
		t.Fatalf("Metadata gorm tag = %q, want %q", got, "serializer:json")
	}
}

// map? is wrapped in emigo.Nullable, which already round-trips arbitrary T (including
// maps) through JSON via its own Value()/Scan(); tagging it with serializer:json on top
// would make gorm bypass that Scan/Value pair and serialize the wrapper's internal
// fields instead, so it must be left untouched.
func TestMapNullableField_UsesNullableRoundtrip(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("MetadataOpt")
	if !ok {
		t.Fatal("expected Entity1Entity to have a MetadataOpt field")
	}
	if field.Type != reflect.TypeOf(emigo.Nullable[map[string]string]{}) {
		t.Fatalf("MetadataOpt type = %v, want emigo.Nullable[map[string]string]", field.Type)
	}
	if tag := field.Tag.Get("gorm"); tag != "" {
		t.Fatalf("MetadataOpt should not have a gorm tag, got %q", tag)
	}

	var n emigo.Nullable[map[string]string]
	n.Set(&map[string]string{"a": "b"})
	value, err := n.Value()
	if err != nil {
		t.Fatalf("Value() error: %v", err)
	}
	if _, ok := value.([]byte); !ok {
		t.Fatalf("expected Value() to JSON-encode the map to []byte, got %T", value)
	}
}

// slice is a raw Go slice with no Scan/Value of its own, so it needs serializer:json
// the same way map does.
func TestSliceField_NeedsJsonSerializer(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Labels")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Labels field")
	}
	if field.Type != reflect.TypeOf([]string{}) {
		t.Fatalf("Labels type = %v, want []string", field.Type)
	}
	if got := field.Tag.Get("gorm"); got != "serializer:json" {
		t.Fatalf("Labels gorm tag = %q, want %q", got, "serializer:json")
	}
}

// slice? is wrapped in emigo.Nullable for the same reason map? is - already handles its
// own JSON round-trip, so no extra tag.
func TestSliceNullableField_UsesNullableRoundtrip(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("LabelsOpt")
	if !ok {
		t.Fatal("expected Entity1Entity to have a LabelsOpt field")
	}
	if field.Type != reflect.TypeOf(emigo.Nullable[[]string]{}) {
		t.Fatalf("LabelsOpt type = %v, want emigo.Nullable[[]string]", field.Type)
	}
	if tag := field.Tag.Get("gorm"); tag != "" {
		t.Fatalf("LabelsOpt should not have a gorm tag, got %q", tag)
	}
}

// any is unwrapped interface{}, so it needs serializer:json just like map/slice.
func TestAnyField_NeedsJsonSerializer(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("Misc")
	if !ok {
		t.Fatal("expected Entity1Entity to have a Misc field")
	}
	if field.Type.Kind() != reflect.Interface {
		t.Fatalf("Misc kind = %v, want interface{}", field.Type.Kind())
	}
	if got := field.Tag.Get("gorm"); got != "serializer:json" {
		t.Fatalf("Misc gorm tag = %q, want %q", got, "serializer:json")
	}
}

// An explicit tags: { gorm: ... } on the field (rawSettings, a "map" field like
// Metadata) always wins over the computed "serializer:json" default.
func TestExplicitGormTag_Overrides(t *testing.T) {
	field, ok := reflect.TypeOf(Entity1Entity{}).FieldByName("RawSettings")
	if !ok {
		t.Fatal("expected Entity1Entity to have a RawSettings field")
	}

	want := "serializer:json;type:jsonb"
	if got := field.Tag.Get("gorm"); got != want {
		t.Fatalf("RawSettings gorm tag = %q, want %q (explicit override should win)", got, want)
	}
}
