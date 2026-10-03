package core

import "testing"

func intentNames(m *Emi) map[string]string {
	out := map[string]string{}
	for _, it := range m.Intents {
		out[it.Name] = it.From
	}
	return out
}

// A compiler backend registers PreprocessEntityActions/entity dto hooks itself; mirror just
// the part these tests need.
func runEntityIntentPreprocessing(t *testing.T, m *Emi) {
	t.Helper()
	for _, e := range m.Entities {
		m.Dto = append(m.Dto, *BuildEntityDto(e), *BuildEntityOptionalDto(e))
	}
	m.preprocessEntityActions()
	if err := m.preprocessIntents(); err != nil {
		t.Fatal(err)
	}
}

func TestEntityIntentExposesEveryGeneratedAction(t *testing.T) {
	m, err := ReadEmiFromString(`
entities:
  - name: sellableInstrument
    intent: true
    fields:
      - {name: title, type: string}
  - name: category
    fields:
      - {name: title, type: string}
`)
	if err != nil {
		t.Fatal(err)
	}
	runEntityIntentPreprocessing(t, m)

	got := intentNames(m)
	want := map[string]string{
		"createSellableInstrument":         "sellableInstrumentCreate",
		"updateSellableInstrument":         "sellableInstrumentUpdate",
		"getSellableInstrument":            "sellableInstrumentGet",
		"listSellableInstruments":          "sellableInstrumentBrowse",
		"previewDeleteSellableInstruments": "sellableInstrumentAwareDeletePreview",
		"deleteSellableInstruments":        "sellableInstrumentAwareDelete",
	}
	if len(got) != len(want) {
		t.Fatalf("intents = %v, want exactly %v (category has no intent: true)", got, want)
	}
	for name, from := range want {
		if got[name] != from {
			t.Errorf("intent %q from %q, want %q", name, got[name], from)
		}
	}
}

func TestEntityIntentYieldsToHandDeclaredIntent(t *testing.T) {
	m, err := ReadEmiFromString(`
entities:
  - name: book
    intent: true
    fields:
      - {name: title, type: string}
`)
	if err != nil {
		t.Fatal(err)
	}
	m.Intents = append(m.Intents, &EmiIntent{Name: "listBooks", From: "bookBrowse", Title: "Mine"})
	runEntityIntentPreprocessing(t, m)
	count := 0
	for _, it := range m.Intents {
		if it.Name == "listBooks" {
			count++
			if it.Title != "Mine" {
				t.Errorf("hand-declared intent was replaced: title %q", it.Title)
			}
		}
	}
	if count != 1 {
		t.Fatalf("listBooks declared %d times, want 1", count)
	}
}

func TestEntityPlural(t *testing.T) {
	for in, want := range map[string]string{"book": "books", "category": "categories", "day": "days", "box": "boxes", "address": "addresses"} {
		if got := entityPlural(in); got != want {
			t.Errorf("entityPlural(%q) = %q, want %q", in, got, want)
		}
	}
}
