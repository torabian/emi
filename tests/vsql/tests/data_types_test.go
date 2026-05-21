package tests

import (
	"strings"
	"testing"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/vsql_test/dto"
	"github.com/torabian/emi/vsql_test/sqlfiles"
	"github.com/torabian/emi/vsql_test/vsql"
)

// render is a small wrapper so each subtest stays focused on the assertion.
func renderDataTypes(t *testing.T, d dto.DataTypesDto) string {
	t.Helper()
	out, err := vsql.Render(sqlfiles.Files, "insert_data_types.sql", d)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	return out
}

// TestDataTypes exercises every emi data type the compiler can produce.
// Each subtest sets a focused slice of fields and asserts on the parts of
// the generated SQL that prove the type was rendered as expected.
func TestDataTypes(t *testing.T) {

	t.Run("Primitives", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			StringField:  "hello",
			BoolField:    true,
			IntField:     42,
			Int32Field:   -7,
			Int64Field:   1_000_000_000_000,
			Float32Field: 3.14,
			Float64Field: 2.718281828,
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out, "string_field", "'hello'", "TRUE",
			"int_field", "42", "1000000000000", "3.14", "2.718281828")
	})

	t.Run("NullablesSet_emitsValue", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			StringFieldNullable:  emigo.NullableOf("set"),
			BoolFieldNullable:    emigo.NullableOf(false),
			IntFieldNullable:     emigo.NullableOf(9),
			Int64FieldNullable:   emigo.NullableOf[int64](99),
			Float64FieldNullable: emigo.NullableOf(1.5),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out,
			"string_field_nullable", "'set'",
			"bool_field_nullable", "FALSE",
			"int_field_nullable", "9",
			"int64_field_nullable", "99",
			"float64_field_nullable", "1.5")
	})

	t.Run("NullablesUnset_omitsColumn", func(t *testing.T) {
		// Empty DTO — every Nullable is unset, so none of the *_nullable
		// columns should appear at all.
		out := renderDataTypes(t, dto.DataTypesDto{})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustNotContain(t, out,
			"string_field_nullable",
			"bool_field_nullable",
			"int_field_nullable",
			"int32_field_nullable",
			"int64_field_nullable",
			"float32_field_nullable",
			"float64_field_nullable",
			"enum_field_nullable",
			"object_field_nullable",
			"array_field_nullable",
			"slice_field_nullable",
			"map_field_nullable",
			"one_ref_nullable",
			"collection_ref_nullable")
	})

	t.Run("Enum_rendersAsQuotedString", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			EnumField:         "Alpha",
			EnumFieldNullable: emigo.NullableOf("Delta"),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out, "enum_field", "'Alpha'", "'Delta'")
	})

	t.Run("Object_flattensWithPrefix", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			ObjectField: dto.DataTypesDtoObjectField{
				NestedString: "inside",
				NestedInt:    11,
			},
			ObjectFieldNullable: emigo.NullableOf(dto.DataTypesDtoObjectFieldNullable{
				OptionalString: "set",
			}),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out,
			"object_field_nested_string", "'inside'",
			"object_field_nested_int", "11",
			"object_field_nullable_optional_string", "'set'")
		// The parent name itself must NOT appear as a column — only the
		// flattened children should.
		mustNotContain(t, out, "object_field,", "object_field )", "object_field=")
	})

	t.Run("Slice_jsonEncodedAsJsonb", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			SliceField:         []string{"a", "b", "c"},
			SliceFieldNullable: emigo.NullableOf([]int{1, 2, 3}),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out,
			`slice_field`, `'["a","b","c"]'::jsonb`,
			`slice_field_nullable`, `'[1,2,3]'::jsonb`)
	})

	t.Run("AnyAndMap_jsonEncoded", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			AnyField: map[string]any{
				"kind": "blob",
				"n":    7,
			},
			MapField: map[any]any{
				"a": 1,
			},
			MapFieldNullable: emigo.NullableOf(map[any]any{
				"k": "v",
			}),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		// AnyField holds a map → JSON-encoded.
		mustContain(t, out, "any_field", `'{`, "::jsonb")
		mustContain(t, out, "map_field", "map_field_nullable")
	})

	t.Run("Complex_usesSQLValuer", func(t *testing.T) {
		out := renderDataTypes(t, dto.DataTypesDto{
			ComplexField: dto.GeoPoint{Lat: 52.2, Lon: 21.0, Valid: true},
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out, "complex_field", "ST_GeomFromText('POINT(21 52.2)')")
	})

	t.Run("ComplexInvalid_omits", func(t *testing.T) {
		// SQLValuer returning include=false drops the column.
		out := renderDataTypes(t, dto.DataTypesDto{
			ComplexField: dto.GeoPoint{Valid: false},
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustNotContain(t, out, "complex_field")
	})

	t.Run("OneRef_flattens", func(t *testing.T) {
		// `one` of CreateUserDtoTags → renderer flattens to one_ref_key,
		// one_ref_value (CreateUserDtoTags is a plain struct).
		out := renderDataTypes(t, dto.DataTypesDto{
			OneRef: dto.CreateUserDtoTags{Key: "role", Value: "admin"},
			OneRefNullable: emigo.NullableOf(dto.CreateUserDtoTags{
				Key: "tier", Value: "gold",
			}),
		})
		t.Logf("\n--- generated SQL ---\n%s", out)
		mustContain(t, out,
			"one_ref_key", "'role'",
			"one_ref_value", "'admin'",
			"one_ref_nullable_key", "'tier'",
			"one_ref_nullable_value", "'gold'")
	})
}

// --- assertion helpers ---------------------------------------------------

func mustContain(t *testing.T, haystack string, needles ...string) {
	t.Helper()
	for _, n := range needles {
		if !strings.Contains(haystack, n) {
			t.Errorf("expected SQL to contain %q\n--- got ---\n%s", n, haystack)
		}
	}
}

func mustNotContain(t *testing.T, haystack string, needles ...string) {
	t.Helper()
	for _, n := range needles {
		if strings.Contains(haystack, n) {
			t.Errorf("expected SQL to NOT contain %q\n--- got ---\n%s", n, haystack)
		}
	}
}
