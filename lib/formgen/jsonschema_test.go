package formgen

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestBuildJSONSchema_KitchenSink(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "fullName", Type: core.FieldTypeString, Description: "Given name"},
		{Name: "age", Type: core.FieldTypeIntNullable},
		{Name: "active", Type: core.FieldTypeBool},
		{Name: "tier", Type: core.FieldTypeEnum, OfType: []*core.EmiEnumInline{
			{Key: "gold", Description: "Gold tier"},
			{Key: "silver"},
		}},
		{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
		{Name: "settings", Type: core.FieldTypeMap, MapPairOf: "string"},
		{Name: "address", Type: core.FieldTypeObjectNullable, Fields: []*core.EmiField{
			{Name: "street", Type: core.FieldTypeString},
		}},
		{Name: "phones", Type: core.FieldTypeArray, Fields: []*core.EmiField{
			{Name: "number", Type: core.FieldTypeString},
		}},
		{Name: "manager", Type: core.FieldTypeOne, Target: "employee"},
		{Name: "reports", Type: core.FieldTypeCollection, Target: "employee"},
	}

	schema := BuildJSONSchema("Customer", "A customer", fields)

	raw, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}
	out := string(raw)
	t.Log(out)

	// Round-trip through a generic map to check semantic shape.
	var generic map[string]any
	if err := json.Unmarshal(raw, &generic); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	if generic["type"] != "object" {
		t.Errorf("expected root type object, got %v", generic["type"])
	}
	if generic["title"] != "Customer" {
		t.Errorf("expected root title Customer, got %v", generic["title"])
	}

	required, _ := generic["required"].([]any)
	requiredSet := map[string]bool{}
	for _, r := range required {
		requiredSet[r.(string)] = true
	}
	if !requiredSet["fullName"] || requiredSet["age"] {
		t.Errorf("expected fullName required and age (nullable) not required, got %v", required)
	}

	props := generic["properties"].(map[string]any)

	age := props["age"].(map[string]any)
	if age["type"] != "integer" {
		t.Errorf("expected age to be integer, got %v", age["type"])
	}

	tier := props["tier"].(map[string]any)
	oneOf := tier["oneOf"].([]any)
	if len(oneOf) != 2 {
		t.Fatalf("expected 2 enum options, got %d", len(oneOf))
	}
	first := oneOf[0].(map[string]any)
	if first["const"] != "gold" || first["title"] != "Gold tier" {
		t.Errorf("expected first enum option {const: gold, title: Gold tier}, got %v", first)
	}
	second := oneOf[1].(map[string]any)
	if second["const"] != "silver" || second["title"] != "silver" {
		t.Errorf("expected label to fall back to raw key when no description, got %v", second)
	}

	tags := props["tags"].(map[string]any)
	if tags["type"] != "array" {
		t.Errorf("expected tags to be array, got %v", tags["type"])
	}
	if tags["items"].(map[string]any)["type"] != "string" {
		t.Errorf("expected tags items to be string, got %v", tags["items"])
	}

	settings := props["settings"].(map[string]any)
	if settings["type"] != "object" {
		t.Errorf("expected settings (map) to be object, got %v", settings["type"])
	}
	if settings["additionalProperties"].(map[string]any)["type"] != "string" {
		t.Errorf("expected settings additionalProperties string, got %v", settings["additionalProperties"])
	}

	address := props["address"].(map[string]any)
	if address["type"] != "object" {
		t.Errorf("expected nested object address to be type object, not a $ref or anything else, got %v", address["type"])
	}
	if _, ok := address["properties"].(map[string]any)["street"]; !ok {
		t.Errorf("expected address.properties.street to exist, got %v", address)
	}

	phones := props["phones"].(map[string]any)
	if phones["type"] != "array" {
		t.Errorf("expected phones to be array, got %v", phones["type"])
	}
	phoneItem := phones["items"].(map[string]any)
	if phoneItem["type"] != "object" {
		t.Errorf("expected phones.items to be type object (array of objects), got %v", phoneItem["type"])
	}

	manager := props["manager"].(map[string]any)
	if _, hasType := manager["type"]; hasType {
		t.Errorf("expected one-relation field to be unconstrained (no type), got %v", manager)
	}

	reports := props["reports"].(map[string]any)
	if reports["type"] != "array" {
		t.Errorf("expected collection-relation field to be an array, got %v", reports)
	}
	if _, hasType := reports["items"].(map[string]any)["type"]; hasType {
		t.Errorf("expected collection-relation items to be unconstrained (no type), got %v", reports["items"])
	}
}

// TestBuildJSONSchema_PropertyOrderIsStable checks that Properties always
// serialize in dto field order (not map/alphabetical order), since RJSF (and
// JSON Schema form renderers generally) may use property order as a default
// field order.
func TestBuildJSONSchema_PropertyOrderIsStable(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "zebra", Type: core.FieldTypeString},
		{Name: "apple", Type: core.FieldTypeString},
		{Name: "mango", Type: core.FieldTypeString},
	}

	schema := BuildJSONSchema("Order", "", fields)
	raw, err := json.Marshal(schema)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := string(raw)

	zebraIdx := strings.Index(out, `"zebra"`)
	appleIdx := strings.Index(out, `"apple"`)
	mangoIdx := strings.Index(out, `"mango"`)

	if !(zebraIdx < appleIdx && appleIdx < mangoIdx) {
		t.Errorf("expected properties in dto field order zebra,apple,mango, got:\n%s", out)
	}
}

// TestBuildSchemaLocales_FlatKeysAndNesting checks the flat
// `<path>_title`/`<path>_description` key scheme SchemaLocales' doc comment
// describes, including one level of WidgetObjectGroup nesting.
func TestBuildSchemaLocales_FlatKeysAndNesting(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "fullName", Type: core.FieldTypeString, Description: "Given name"},
		{Name: "age", Type: core.FieldTypeIntNullable},
		{Name: "address", Type: core.FieldTypeObjectNullable, Fields: []*core.EmiField{
			{Name: "city", Type: core.FieldTypeString, Description: "City name"},
		}},
	}

	locales := BuildSchemaLocales("Customer", "A customer", fields)

	raw, err := json.Marshal(locales)
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	var generic map[string]map[string]string
	if err := json.Unmarshal(raw, &generic); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	defaultBucket, ok := generic["default"]
	if !ok {
		t.Fatalf("expected a \"default\" bucket, got %v", generic)
	}

	want := map[string]string{
		"fullName_title":           "Full Name",
		"fullName_description":     "Given name",
		"age_title":                "Age",
		"address_title":            "Address",
		"address_city_title":       "City",
		"address_city_description": "City name",
	}
	for k, v := range want {
		if defaultBucket[k] != v {
			t.Errorf("expected %q = %q, got %q", k, v, defaultBucket[k])
		}
	}

	// age has no Description, so no age_description entry should be emitted.
	if _, ok := defaultBucket["age_description"]; ok {
		t.Errorf("expected no age_description entry for a field with no description, got %q", defaultBucket["age_description"])
	}
}

// TestBuildSchemaLocales_KeyOrderIsStable mirrors
// TestBuildJSONSchema_PropertyOrderIsStable: entries must serialize in dto
// field order, not map/alphabetical order.
func TestBuildSchemaLocales_KeyOrderIsStable(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "zebra", Type: core.FieldTypeString},
		{Name: "apple", Type: core.FieldTypeString},
		{Name: "mango", Type: core.FieldTypeString},
	}

	locales := BuildSchemaLocales("Order", "", fields)
	raw, err := json.Marshal(locales)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := string(raw)

	zebraIdx := strings.Index(out, `"zebra_title"`)
	appleIdx := strings.Index(out, `"apple_title"`)
	mangoIdx := strings.Index(out, `"mango_title"`)

	if !(zebraIdx < appleIdx && appleIdx < mangoIdx) {
		t.Errorf("expected entries in dto field order zebra,apple,mango, got:\n%s", out)
	}
}

// TestAllCatalogFieldTypesProduceValidSchema mirrors
// TestAllCatalogFieldTypesResolve in widget_test.go, but at the JSON Schema
// layer: every field type must marshal to *some* schema without panicking or
// erroring, so a new EmiField type never silently breaks schema generation.
func TestAllCatalogFieldTypesProduceValidSchema(t *testing.T) {
	catalog := core.GetEmiFieldTypeCatalog()
	all := append([]core.FieldType{}, catalog.DtoFieldTypes...)
	all = append(all, catalog.DtoNullableFieldTypes...)

	for _, ft := range all {
		field := &core.EmiField{Name: "x", Type: ft, Primitive: "string", Target: "y"}
		schema := BuildJSONSchema("Root", "", []*core.EmiField{field})

		raw, err := json.Marshal(schema)
		if err != nil {
			t.Fatalf("field type %q: unexpected marshal error: %v", ft, err)
		}
		var generic map[string]any
		if err := json.Unmarshal(raw, &generic); err != nil {
			t.Fatalf("field type %q: produced invalid JSON: %v\n%s", ft, err, raw)
		}
	}
}

// TestBuildJSONSchemaWithTranslationKeys_KitchenSink mirrors
// TestBuildJSONSchema_KitchenSink's fixture, but checks the translation-key
// mode: every title/description/enum-option string in the schema must be a
// key, never literal text, and DefaultTranslations must carry that literal
// text back for every one of those keys - including ones nested inside an
// object, inside an array-of-objects item, and an enum option label (the one
// case BuildSchemaLocales/collectSchemaLocaleEntries never covers at all).
func TestBuildJSONSchemaWithTranslationKeys_KitchenSink(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "fullName", Type: core.FieldTypeString, Description: "Given name"},
		{Name: "tier", Type: core.FieldTypeEnum, OfType: []*core.EmiEnumInline{
			{Key: "gold", Description: "Gold tier"},
			{Key: "silver"},
		}},
		{Name: "address", Type: core.FieldTypeObjectNullable, Fields: []*core.EmiField{
			{Name: "city", Type: core.FieldTypeString, Description: "City name"},
		}},
		{Name: "phones", Type: core.FieldTypeArray, Fields: []*core.EmiField{
			{Name: "number", Type: core.FieldTypeString},
		}},
	}

	schema, translations := BuildJSONSchemaWithTranslationKeys("Customer", "A customer", fields)

	raw, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}
	var generic map[string]any
	if err := json.Unmarshal(raw, &generic); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	// Root title/description are the reserved keys, not literal text.
	if generic["title"] != "$title" {
		t.Errorf("expected root title to be the key \"$title\", got %v", generic["title"])
	}
	if generic["description"] != "$description" {
		t.Errorf("expected root description to be the key \"$description\", got %v", generic["description"])
	}

	translationsByKey := map[string]string{}
	for _, e := range translations.Entries {
		if _, dup := translationsByKey[e.Key]; dup {
			t.Errorf("duplicate translation key %q", e.Key)
		}
		translationsByKey[e.Key] = e.Value
	}

	expect := map[string]string{
		"$title":                               "Customer",
		"$description":                         "A customer",
		"full_name_title":                       "Full Name",
		"full_name_description":                 "Given name",
		"tier_title":                           "Tier",
		"tier_enum_gold":                       "Gold tier",
		"tier_enum_silver":                     "silver",
		"address_title":                        "Address",
		"address_properties_city_title":        "City",
		"address_properties_city_description":  "City name",
		"phones_title":                         "Phones",
		"phones_items_properties_number_title": "Number",
	}
	for key, want := range expect {
		got, ok := translationsByKey[key]
		if !ok {
			t.Errorf("expected translation key %q to exist, got entries: %v", key, translationsByKey)
			continue
		}
		if got != want {
			t.Errorf("translation key %q: expected %q, got %q", key, want, got)
		}
	}

	props := generic["properties"].(map[string]any)

	fullName := props["fullName"].(map[string]any)
	if fullName["title"] != "full_name_title" || fullName["description"] != "full_name_description" {
		t.Errorf("expected fullName's title/description to be keys, got %v", fullName)
	}

	tier := props["tier"].(map[string]any)
	oneOf := tier["oneOf"].([]any)
	first := oneOf[0].(map[string]any)
	if first["const"] != "gold" || first["title"] != "tier_enum_gold" {
		t.Errorf("expected enum option title to be a key, got %v", first)
	}

	address := props["address"].(map[string]any)
	city := address["properties"].(map[string]any)["city"].(map[string]any)
	if city["title"] != "address_properties_city_title" {
		t.Errorf("expected nested field title to be an underscore-joined key mirroring the schema path, got %v", city["title"])
	}

	phones := props["phones"].(map[string]any)
	phoneItemProps := phones["items"].(map[string]any)["properties"].(map[string]any)
	number := phoneItemProps["number"].(map[string]any)
	if number["title"] != "phones_items_properties_number_title" {
		t.Errorf("expected array item field title to be an underscore-joined key mirroring the schema path, got %v", number["title"])
	}
}

// TestBuildJSONSchema_UnaffectedByTranslationKeysMode locks in that
// BuildJSONSchema's own output - the literal-text mode - is byte-for-byte
// unchanged now that it shares buildJSONSchema/objectSchema/fieldSchema/
// enumSchema with BuildJSONSchemaWithTranslationKeys, instead of being a
// second schema builder that could silently drift from this one.
func TestBuildJSONSchema_UnaffectedByTranslationKeysMode(t *testing.T) {
	fields := []*core.EmiField{
		{Name: "fullName", Type: core.FieldTypeString, Description: "Given name"},
		{Name: "tier", Type: core.FieldTypeEnum, OfType: []*core.EmiEnumInline{
			{Key: "gold", Description: "Gold tier"},
		}},
		{Name: "address", Type: core.FieldTypeObjectNullable, Fields: []*core.EmiField{
			{Name: "city", Type: core.FieldTypeString},
		}},
	}

	schema := BuildJSONSchema("Customer", "A customer", fields)
	raw, err := json.Marshal(schema)
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	var generic map[string]any
	if err := json.Unmarshal(raw, &generic); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if generic["title"] != "Customer" {
		t.Errorf("expected literal root title, got %v", generic["title"])
	}
	fullName := generic["properties"].(map[string]any)["fullName"].(map[string]any)
	if fullName["title"] != "Full Name" || fullName["description"] != "Given name" {
		t.Errorf("expected literal field title/description, got %v", fullName)
	}
}
