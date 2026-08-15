package js

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/formgen"
)

// kitchenSinkDto exercises every WidgetKind the react form writer supports in
// a single dto: leaf primitives, an enum, a slice, a map, a nullable object
// group, an array-of-objects group, and one/collection relations.
func kitchenSinkDto() core.EmiDto {
	return core.EmiDto{
		Name: "profile",
		Fields: []*core.EmiField{
			{Name: "firstName", Type: core.FieldTypeString, Description: "Given name"},
			{Name: "age", Type: core.FieldTypeIntNullable},
			{Name: "active", Type: core.FieldTypeBool},
			{Name: "role", Type: core.FieldTypeEnum, OfType: []*core.EmiEnumInline{
				{Key: "admin", Description: "Administrator"},
				{Key: "member"},
			}},
			{Name: "tags", Type: core.FieldTypeSlice, Primitive: "string"},
			{Name: "settings", Type: core.FieldTypeMap, MapKeyOf: "string", MapPairOf: "string"},
			{Name: "address", Type: core.FieldTypeObjectNullable, Fields: []*core.EmiField{
				{Name: "street", Type: core.FieldTypeString},
				{Name: "city", Type: core.FieldTypeString},
			}},
			{Name: "phones", Type: core.FieldTypeArray, Fields: []*core.EmiField{
				{Name: "number", Type: core.FieldTypeString},
			}},
			{Name: "manager", Type: core.FieldTypeOne, Target: "employee"},
			{Name: "reports", Type: core.FieldTypeCollection, Target: "employee", Module: "hr"},
		},
	}
}

func TestReactFormGenerator_KitchenSink(t *testing.T) {
	toolset := DefaultReactFormToolset("", "")

	file, err := ReactFormGenerator(kitchenSinkDto(), toolset)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if file.Extension != ".tsx" {
		t.Errorf("expected .tsx extension, got %q", file.Extension)
	}
	if file.Name != "ProfileDtoForm" {
		t.Errorf("expected file name ProfileDtoForm, got %q", file.Name)
	}

	script := file.ActualScript
	t.Log(script)

	checks := []string{
		// root component + controlled props
		"export function ProfileDtoForm({ value, onChange }: ProfileDtoFormProps)",
		"export default ProfileDtoForm;",

		// primitives
		`<TextField label="First Name"`,
		`help="Given name"`,
		`<NumberField label="Age"`,
		`<CheckboxField label="Active"`,

		// enum
		`<SelectField label="Role"`,
		`{ value: "admin", label: "Administrator" }`,
		`{ value: "member", label: "member" }`,

		// slice / map
		`<ListField label="Tags"`,
		`primitive="string"`,
		`<MapField label="Settings"`,
		`keyType="string"`,

		// nested object group gets its own function + type, not inlined
		"function ProfileDtoAddressForm({ value, onChange }",
		"export type ProfileDtoAddressValues = {",
		"<ProfileDtoAddressForm value={value.address ?? {}}",

		// nested array group gets item form + array wrapper, not inlined
		"function ProfileDtoPhonesItemForm({ value, onChange }",
		"function ProfileDtoPhonesArrayForm({ value, onChange }",
		"<ProfileDtoPhonesArrayForm value={value.phones ?? []}",

		// one/collection relations import a picker instead of inlining fields
		`import { EmployeePicker } from "./relations/employee";`,
		`import { EmployeePicker } from "./relations/hr/employee";`,
		"<EmployeePicker value={value.manager}",
		"<EmployeePicker multiple value={value.reports ?? []}",

		// primitive widgets import from the configured components module
		`from "./form-controls";`,
	}

	for _, want := range checks {
		if !strings.Contains(script, want) {
			t.Errorf("expected generated form to contain %q, got:\n%s", want, script)
		}
	}

	// Required fields (non-nullable) get a `required` prop; nullable ones don't.
	if !strings.Contains(script, `<TextField label="First Name" value={value.firstName} onChange={(next) => onChange({ ...value, firstName: next })} required help="Given name" />`) &&
		!strings.Contains(script, `required`) {
		t.Errorf("expected required prop on non-nullable firstName field")
	}
}

func TestReactFormGenerator_RejectsEmptyName(t *testing.T) {
	_, err := ReactFormGenerator(core.EmiDto{}, DefaultReactFormToolset("", ""))
	if err == nil {
		t.Fatalf("expected error for dto with no name")
	}
}

func TestReactFormFileAction_RejectsNonDto(t *testing.T) {
	_, err := ReactFormFileAction.Run(core.MicroGenContext{Content: "emi: module\nname: x\n"})
	if err == nil {
		t.Fatalf("expected error when content is not a dto")
	}
}

func TestReactFormFileAction_EndToEnd(t *testing.T) {
	content := `
emi: dto
name: singleDto
fields:
  - name: min
    type: int?
  - name: max
    type: int
  - name: nullableObject
    type: object?
    fields:
      - name: firstName
        type: string
  - name: staticArray
    type: array
    fields:
      - name: item12
        type: string
`
	files, err := ReactFormFileAction.Run(core.MicroGenContext{
		Content: content,
		Flags:   map[string]string{"form-components": "@ui/form", "form-relations": "@ui/relations"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("expected exactly 1 file (single file per dto), got %d", len(files))
	}

	f := files[0]
	if f.Name != "SingleDtoDtoForm" || f.Extension != ".tsx" {
		t.Errorf("unexpected file identity: %+v", core.VirtualFile{Name: f.Name, Extension: f.Extension})
	}
	if !strings.Contains(f.ActualScript, `from "@ui/form";`) {
		t.Errorf("expected overridden --form-components import, got:\n%s", f.ActualScript)
	}
	if !strings.Contains(f.ActualScript, "function SingleDtoDtoNullableObjectForm") {
		t.Errorf("expected nested object sub-form, got:\n%s", f.ActualScript)
	}
	if !strings.Contains(f.ActualScript, "function SingleDtoDtoStaticArrayArrayForm") {
		t.Errorf("expected nested array sub-form, got:\n%s", f.ActualScript)
	}
}

// Guard against a WidgetKind formgen knows about but a given Toolset (like
// DefaultReactFormToolset) forgot to configure: the writer must still emit
// *something* editable (falls back to the toolset's JSON widget) instead of
// silently dropping the field.
func TestReactFormBuilder_UnconfiguredWidgetFallsBackToJson(t *testing.T) {
	toolset := formgen.NewToolset() // nothing registered at all
	toolset.RelationComponent = func(target string) string { return target }
	toolset.RelationImportFrom = func(target, module string) string { return target }

	b := newReactFormBuilder(toolset)
	usage := b.renderLeaf(&formgen.FieldPlan{Name: "x", Widget: formgen.WidgetText})

	if !strings.Contains(usage, "<JsonField") {
		t.Errorf("expected fallback to JsonField, got %q", usage)
	}
}
