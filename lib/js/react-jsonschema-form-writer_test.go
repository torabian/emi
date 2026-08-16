package js

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestReactJsonSchemaFormGenerator_KitchenSink(t *testing.T) {
	files, err := ReactJsonSchemaFormGenerator(kitchenSinkDto())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files (schema, wrapper), got %d", len(files))
	}

	byName := map[string]core.VirtualFile{}
	for _, f := range files {
		byName[f.Name+f.Extension] = f
	}

	schemaFile, ok := byName["ProfileDto.schema.json"]
	if !ok {
		t.Fatalf("expected ProfileDto.schema.json, got %+v", byName)
	}
	wrapperFile, ok := byName["ProfileDtoForm.tsx"]
	if !ok {
		t.Fatalf("expected ProfileDtoForm.tsx, got %+v", byName)
	}
	if _, ok := byName["ProfileDto.uiSchema.json"]; ok {
		t.Fatalf("did not expect a uiSchema.json file to be generated, got %+v", byName)
	}

	// --- schema.json ---
	var schema map[string]any
	if err := json.Unmarshal([]byte(schemaFile.ActualScript), &schema); err != nil {
		t.Fatalf("schema.json is not valid JSON: %v\n%s", err, schemaFile.ActualScript)
	}
	props := schema["properties"].(map[string]any)

	if props["firstName"].(map[string]any)["type"] != "string" {
		t.Errorf("expected firstName: string, got %v", props["firstName"])
	}
	if _, hasType := props["manager"].(map[string]any)["type"]; hasType {
		t.Errorf("expected one-relation manager to be unconstrained, got %v", props["manager"])
	}
	if props["address"].(map[string]any)["type"] != "object" {
		t.Errorf("expected nested object address inlined as type object, got %v", props["address"])
	}
	if props["phones"].(map[string]any)["type"] != "array" {
		t.Errorf("expected nested array phones inlined as type array, got %v", props["phones"])
	}

	// --- wrapper.tsx ---
	wrapper := wrapperFile.ActualScript
	checks := []string{
		`import Form from "@rjsf/core";`,
		`import schema from "./ProfileDto.schema.json";`,
		"export function ProfileDtoForm({ value, onChange }: ProfileDtoFormProps)",
		"export default ProfileDtoForm;",
	}
	for _, want := range checks {
		if !strings.Contains(wrapper, want) {
			t.Errorf("expected wrapper to contain %q, got:\n%s", want, wrapper)
		}
	}
	if strings.Contains(wrapper, "uiSchema") {
		t.Errorf("did not expect wrapper to reference a uiSchema, got:\n%s", wrapper)
	}
}

func TestReactJsonSchemaFormFileAction_EndToEnd(t *testing.T) {
	content := `
emi: dto
name: singleDto
fields:
  - name: min
    type: int?
  - name: nullableObject
    type: object?
    fields:
      - name: firstName
        type: string
`
	files, err := ReactJsonSchemaFormFileAction.Run(core.MicroGenContext{Content: content})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}
}

func TestReactJsonSchemaFormFileAction_RejectsNonDto(t *testing.T) {
	_, err := ReactJsonSchemaFormFileAction.Run(core.MicroGenContext{Content: "emi: module\nname: x\n"})
	if err == nil {
		t.Fatalf("expected error when content is not a dto")
	}
}
