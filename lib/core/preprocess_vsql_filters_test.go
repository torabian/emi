package core

import "testing"

func TestVsqlFiltersDefaultFromColumns(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "selectUserById",
				Columns: []*EmiColumn{
					{EmiField: EmiField{Name: "id", Type: FieldTypeInt64}, Selected: true},
					{EmiField: EmiField{Name: "email", Type: FieldTypeString}, Selected: true},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	filters := m.Vsqls[0].Filters
	if len(filters) != 2 || filters[0].Name != "id" || filters[1].Name != "email" {
		t.Fatalf("expected Filters to default to Columns' fields, got %+v", filters)
	}
}

// TestVsqlFiltersMirrorNullifiedColumnTypes: Filters is built AFTER
// nullifyUnselectedColumns has already touched Columns, so an unselected
// column's Filters counterpart carries the same nullable type the row DTO
// actually has - not the plain type as originally declared.
func TestVsqlFiltersMirrorNullifiedColumnTypes(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "selectUserById",
				Columns: []*EmiColumn{
					{EmiField: EmiField{Name: "balanceCents", Type: FieldTypeInt64}, Selected: false},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	filters := m.Vsqls[0].Filters
	if len(filters) != 1 || filters[0].Type != "int64?" {
		t.Fatalf("expected Filters' type to mirror the nullified column type, got %+v", filters)
	}
}

func TestVsqlFiltersExplicitWinsOverColumnsDefault(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "selectUserById",
				Columns: []*EmiColumn{
					{EmiField: EmiField{Name: "id", Type: FieldTypeInt64}, Selected: true},
				},
				Filters: []*EmiField{
					{Name: "role", Type: FieldTypeString},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	filters := m.Vsqls[0].Filters
	if len(filters) != 1 || filters[0].Name != "role" {
		t.Fatalf("expected the explicit Filters list to win over Columns, got %+v", filters)
	}
}

func TestVsqlFiltersEmptyWithoutColumns(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "insertUser", Params: []*EmiField{{Name: "email", Type: FieldTypeString}}},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	if len(m.Vsqls[0].Filters) != 0 {
		t.Fatalf("expected no default Filters when there are no Columns to default from, got %+v", m.Vsqls[0].Filters)
	}
}

// TestVsqlFiltersIsIdempotentAcrossPreprocessPasses guards against a second
// Preprocess() pass re-defaulting Filters from Columns and duplicating it
// (Preprocess is documented as safe to call more than once).
func TestVsqlFiltersIsIdempotentAcrossPreprocessPasses(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "selectUserById",
				Columns: []*EmiColumn{
					{EmiField: EmiField{Name: "id", Type: FieldTypeInt64}, Selected: true},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (1st): %v", err)
	}
	first := m.Vsqls[0].Filters
	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (2nd): %v", err)
	}
	if len(m.Vsqls[0].Filters) != len(first) {
		t.Fatalf("expected the second pass to leave Filters untouched, got %+v", m.Vsqls[0].Filters)
	}
}
