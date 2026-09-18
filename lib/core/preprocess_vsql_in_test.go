package core

import "testing"

// fieldNames is a small assertion helper: the ordered list of field names on
// a resolved vsql's Params.
func fieldNames(fields []*EmiField) []string {
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = f.Name
	}
	return names
}

func assertFieldNames(t *testing.T, got []*EmiField, want ...string) {
	t.Helper()
	names := fieldNames(got)
	if len(names) != len(want) {
		t.Fatalf("got fields %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("got fields %v, want %v", names, want)
		}
	}
}

// TestVsqlInFieldsMergeIntoParams covers the plain-Fields half of In: it's
// just another source of fields, additive with Params, not a replacement.
func TestVsqlInFieldsMergeIntoParams(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "search",
				Params: []*EmiField{
					{Name: "limit", Type: FieldTypeInt},
				},
				In: &EmiActionBody{
					Fields: []*EmiField{
						{Name: "query", Type: FieldTypeString},
					},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	assertFieldNames(t, m.Vsqls[0].Params, "limit", "query")

	// In itself survives preprocessing (its Headers/Dto name aren't
	// expressible as plain fields, so future action-exposure code still
	// needs it).
	if m.Vsqls[0].In == nil {
		t.Fatalf("expected In to remain set after preprocessing")
	}
}

// TestVsqlInDtoResolvesAgainstTopLevelDto covers In.Dto referencing an
// existing top-level Dto instead of repeating its fields inline.
func TestVsqlInDtoResolvesAgainstTopLevelDto(t *testing.T) {
	m := &Emi{
		Dto: []EmiDto{
			{
				Name: "userFilter",
				Fields: []*EmiField{
					{Name: "email", Type: FieldTypeString},
					{Name: "active", Type: FieldTypeBool},
				},
			},
		},
		Vsqls: []EmiVsql{
			{
				Name: "listUsers",
				In:   &EmiActionBody{Dto: "userFilter"},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	assertFieldNames(t, m.Vsqls[0].Params, "email", "active")
}

// TestVsqlInDtoResolvesAgainstTemplateDto covers the Templates.Dtos lookup
// path specifically (distinct map from top-level Dto).
func TestVsqlInDtoResolvesAgainstTemplateDto(t *testing.T) {
	m := &Emi{
		Templates: &EmiTemplate{
			Dtos: []EmiDto{
				{
					Name:   "pageFilter",
					Fields: []*EmiField{{Name: "page", Type: FieldTypeInt}},
				},
			},
		},
		Vsqls: []EmiVsql{
			{
				Name: "listUsers",
				In:   &EmiActionBody{Dto: "pageFilter"},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	assertFieldNames(t, m.Vsqls[0].Params, "page")
}

// TestVsqlInUnknownDtoErrors makes sure a typo'd In.Dto fails loudly instead
// of silently producing a query with no params.
func TestVsqlInUnknownDtoErrors(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{Name: "listUsers", In: &EmiActionBody{Dto: "doesNotExist"}},
		},
	}

	if err := m.Preprocess(); err == nil {
		t.Fatalf("expected an error for an unresolvable In.Dto")
	}
}

// TestVsqlParamsWinsOverInOnNameCollision: Params and In describing the same
// field name is not an error - Params, the more specific/local declaration,
// wins, and the field appears exactly once.
func TestVsqlParamsWinsOverInOnNameCollision(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "search",
				Params: []*EmiField{
					{Name: "limit", Type: FieldTypeInt64}, // deliberately differs from In's below
				},
				In: &EmiActionBody{
					Fields: []*EmiField{
						{Name: "limit", Type: FieldTypeInt32},
						{Name: "offset", Type: FieldTypeInt32},
					},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess: %v", err)
	}

	assertFieldNames(t, m.Vsqls[0].Params, "limit", "offset")
	if m.Vsqls[0].Params[0].Type != FieldTypeInt64 {
		t.Fatalf("expected Params' own \"limit\" (int64) to win over In's (int32), got %v", m.Vsqls[0].Params[0].Type)
	}
}

// TestVsqlInIsIdempotentAcrossPreprocessPasses guards the merge against
// duplicating fields if Preprocess ever runs twice on the same module (it's
// documented as safe to do so).
func TestVsqlInIsIdempotentAcrossPreprocessPasses(t *testing.T) {
	m := &Emi{
		Vsqls: []EmiVsql{
			{
				Name: "search",
				In: &EmiActionBody{
					Fields: []*EmiField{{Name: "query", Type: FieldTypeString}},
				},
			},
		},
	}

	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (1st): %v", err)
	}
	if err := m.Preprocess(); err != nil {
		t.Fatalf("Preprocess (2nd): %v", err)
	}

	assertFieldNames(t, m.Vsqls[0].Params, "query")
}
