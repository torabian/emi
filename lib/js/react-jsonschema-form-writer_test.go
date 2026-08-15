package js

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestReactJsonSchemaFormGenerator_KitchenSink(t *testing.T) {
	files, err := ReactJsonSchemaFormGenerator(kitchenSinkDto(), DefaultRjsfToolset(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 3 {
		t.Fatalf("expected 3 files (schema, uiSchema, wrapper), got %d", len(files))
	}

	byName := map[string]core.VirtualFile{}
	for _, f := range files {
		byName[f.Name+f.Extension] = f
	}

	schemaFile, ok := byName["ProfileDto.schema.json"]
	if !ok {
		t.Fatalf("expected ProfileDto.schema.json, got %+v", byName)
	}
	uiSchemaFile, ok := byName["ProfileDto.uiSchema.json"]
	if !ok {
		t.Fatalf("expected ProfileDto.uiSchema.json, got %+v", byName)
	}
	wrapperFile, ok := byName["ProfileDtoForm.tsx"]
	if !ok {
		t.Fatalf("expected ProfileDtoForm.tsx, got %+v", byName)
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

	// --- uiSchema.json ---
	var uiSchema map[string]any
	if err := json.Unmarshal([]byte(uiSchemaFile.ActualScript), &uiSchema); err != nil {
		t.Fatalf("uiSchema.json is not valid JSON: %v\n%s", err, uiSchemaFile.ActualScript)
	}

	order, _ := uiSchema["ui:order"].([]any)
	if len(order) == 0 {
		t.Fatalf("expected top-level ui:order, got %v", uiSchema)
	}
	if order[0] != "firstName" {
		t.Errorf("expected ui:order to start with firstName (dto field order), got %v", order)
	}

	managerUi := uiSchema["manager"].(map[string]any)
	if managerUi["ui:field"] != "emiRelationField" {
		t.Errorf("expected manager ui:field emiRelationField, got %v", managerUi)
	}
	managerOpts := managerUi["ui:options"].(map[string]any)
	if managerOpts["target"] != "employee" {
		t.Errorf("expected manager ui:options.target employee, got %v", managerOpts)
	}
	if _, hasMultiple := managerOpts["multiple"]; hasMultiple {
		t.Errorf("expected one-relation to not set multiple, got %v", managerOpts)
	}

	reportsUi := uiSchema["reports"].(map[string]any)
	reportsOpts := reportsUi["ui:options"].(map[string]any)
	if reportsOpts["target"] != "employee" || reportsOpts["module"] != "hr" || reportsOpts["multiple"] != true {
		t.Errorf("expected reports ui:options {target: employee, module: hr, multiple: true}, got %v", reportsOpts)
	}

	// nested object group's ui:order lives under the field key directly
	addressUi := uiSchema["address"].(map[string]any)
	addressOrder, _ := addressUi["ui:order"].([]any)
	if len(addressOrder) != 2 || addressOrder[0] != "street" {
		t.Errorf("expected address ui:order [street, city], got %v", addressOrder)
	}

	// nested array group's ui:order lives under field.items
	phonesUi := uiSchema["phones"].(map[string]any)
	phonesItemsUi := phonesUi["items"].(map[string]any)
	if _, ok := phonesItemsUi["ui:order"]; !ok {
		t.Errorf("expected phones.items.ui:order, got %v", phonesUi)
	}

	// --- wrapper.tsx ---
	wrapper := wrapperFile.ActualScript
	checks := []string{
		`import Form from "@rjsf/core";`,
		`import { emiRelationField, emiJsonField } from "./rjsf-fields";`,
		`import schema from "./ProfileDto.schema.json";`,
		`import uiSchema from "./ProfileDto.uiSchema.json";`,
		"export function ProfileDtoForm({ value, onChange }: ProfileDtoFormProps)",
		"fields={{ emiRelationField, emiJsonField }}",
		"export default ProfileDtoForm;",
	}
	for _, want := range checks {
		if !strings.Contains(wrapper, want) {
			t.Errorf("expected wrapper to contain %q, got:\n%s", want, wrapper)
		}
	}
}

func TestReactJsonSchemaFormGenerator_CustomFieldsImport(t *testing.T) {
	files, err := ReactJsonSchemaFormGenerator(kitchenSinkDto(), DefaultRjsfToolset(), "@ui/rjsf")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var wrapper core.VirtualFile
	for _, f := range files {
		if f.Extension == ".tsx" {
			wrapper = f
		}
	}
	if !strings.Contains(wrapper.ActualScript, `from "@ui/rjsf";`) {
		t.Errorf("expected overridden --rjsf-fields import, got:\n%s", wrapper.ActualScript)
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
	if len(files) != 3 {
		t.Fatalf("expected 3 files, got %d", len(files))
	}
}

func TestReactJsonSchemaFormFileAction_RejectsNonDto(t *testing.T) {
	_, err := ReactJsonSchemaFormFileAction.Run(core.MicroGenContext{Content: "emi: module\nname: x\n"})
	if err == nil {
		t.Fatalf("expected error when content is not a dto")
	}
}
