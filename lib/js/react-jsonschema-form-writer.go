// react-jsonschema-form-writer.go is the react-jsonschema-form (RJSF) half
// of Emi's form compiler, an alternative to react-form-writer.go's generated
// TSX: instead of emitting one bespoke component per dto, it emits *data* -
// a JSON Schema (lib/formgen's renderer-agnostic BuildJSONSchema) plus an
// RJSF uiSchema - and a single generic `<Form schema={...} uiSchema={...} />`
// renders any dto. Nested object/array fields need no sub-components at all
// here: JSON Schema already expresses them natively (`properties`, `items`),
// and RJSF's own ObjectField/ArrayField render them recursively - the "keep
// array/object forms separate" requirement that forced dedicated functions
// in the TSX writer simply doesn't apply to a schema-driven renderer.
//
// What plain JSON Schema *can't* express is exactly what forced custom
// components in the TSX writer too: `one`/`collection` relation fields (this
// compiler was never given the target entity's fields) and `any`/`complex`
// fields (no fixed shape at all). Both resolve to an unconstrained `{}`
// schema (see lib/formgen/jsonschema.go) plus a `"ui:field"` entry in the
// uiSchema naming a custom RJSF field component - see
// react-jsonschema-form-toolset.go for how that name is chosen, and
// --rjsf-fields for where the generated wrapper imports it from. That's
// RJSF's built-in extension point: `<Form fields={{ [name]: Component }} />`
// registers any custom field, keyed by exactly the name a "ui:field" entry
// references.
package js

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/formgen"
)

// ReactJsonSchemaFormGenerator turns a single dto into three files:
//   - {Dto}.schema.json      - the JSON Schema (formgen.BuildJSONSchema)
//   - {Dto}.uiSchema.json    - field order + ui:field/ui:widget/ui:options
//   - {Dto}Form.tsx          - a thin wrapper around @rjsf/core's <Form>,
//     wired to the two files above and controlled via the same value/onChange
//     convention react-form-writer.go's components use.
func ReactJsonSchemaFormGenerator(dto core.EmiDto, toolset RjsfToolset, fieldsImport string) ([]core.VirtualFile, error) {
	if dto.Name == "" {
		return nil, errors.New("js:rjsf: dto has no name")
	}
	if fieldsImport == "" {
		fieldsImport = DefaultRjsfFieldsImport
	}

	className := dto.GetClassName()
	plan := formgen.BuildFormPlan(className, dto.Fields)

	schema := formgen.BuildJSONSchema(className, dto.Description, dto.Fields)
	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("js:rjsf: failed marshaling schema: %w", err)
	}

	uiSchema := buildUiSchema(plan.Fields, toolset)
	uiSchemaJSON, err := json.MarshalIndent(uiSchema, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("js:rjsf: failed marshaling uiSchema: %w", err)
	}

	wrapper := rjsfWrapperComponent(className, toolset, fieldsImport)

	return []core.VirtualFile{
		{Name: className + ".schema", Extension: ".json", ActualScript: string(schemaJSON) + "\n"},
		{Name: className + ".uiSchema", Extension: ".json", ActualScript: string(uiSchemaJSON) + "\n"},
		{Name: className + "Form", Extension: ".tsx", ActualScript: FormatWithPrettier(wrapper, true)},
	}, nil
}

// buildUiSchema walks a FormPlan's fields into RJSF's uiSchema shape: one
// nested object per level (mirroring the schema's own object/array nesting),
// each carrying an explicit "ui:order" - JSON Schema `properties` order
// isn't a guarantee every consumer preserves, "ui:order" is RJSF's own
// documented, authoritative way to pin it - plus a "ui:field"/"ui:widget"
// entry for whichever fields need one.
func buildUiSchema(fields []*formgen.FieldPlan, toolset RjsfToolset) map[string]any {
	node := map[string]any{}
	order := make([]string, 0, len(fields))

	for _, f := range fields {
		order = append(order, f.Name)
		if entry := uiSchemaForField(f, toolset); entry != nil {
			node[f.Name] = entry
		}
	}

	if len(order) > 0 {
		node["ui:order"] = order
	}

	return node
}

func uiSchemaForField(f *formgen.FieldPlan, toolset RjsfToolset) map[string]any {
	switch f.Widget {
	case formgen.WidgetObjectGroup:
		child := buildUiSchema(f.Children, toolset)
		if len(child) == 0 {
			return nil
		}
		return child

	case formgen.WidgetArrayGroup:
		itemUi := buildUiSchema(f.Children, toolset)
		if len(itemUi) == 0 {
			return nil
		}
		return map[string]any{"items": itemUi}

	case formgen.WidgetOneRelation, formgen.WidgetCollectionRelation:
		options := map[string]any{"target": f.Target}
		if f.Module != "" {
			options["module"] = f.Module
		}
		if f.Widget == formgen.WidgetCollectionRelation {
			options["multiple"] = true
		}
		return map[string]any{
			"ui:field":   toolset.RelationFieldName,
			"ui:options": options,
		}

	case formgen.WidgetJson:
		return map[string]any{"ui:field": toolset.JsonFieldName}

	default:
		if widget, ok := toolset.Widgets[f.Widget]; ok && widget != "" {
			return map[string]any{"ui:widget": widget}
		}
		return nil
	}
}

func rjsfWrapperComponent(className string, toolset RjsfToolset, fieldsImport string) string {
	return fmt.Sprintf(`// Code generated by Emi's form compiler (js:rjsf) from the %q dto. DO NOT EDIT.
// This file is regenerated in full every run alongside %s.schema.json and
// %s.uiSchema.json - wrap %sForm instead of editing any of the three.
//
// Assumes %s exports the custom RJSF fields %s and %s referenced by the
// generated uiSchema for one/collection relation and any/complex fields
// respectively (see --rjsf-fields to point this at a different module).

import Form from "@rjsf/core";
import validator from "@rjsf/validator-ajv8";
import { %s, %s } from %q;
import schema from "./%s.schema.json";
import uiSchema from "./%s.uiSchema.json";

export interface %sFormProps {
  value: any;
  onChange: (next: any) => void;
}

export function %sForm({ value, onChange }: %sFormProps) {
  return (
    <Form
      schema={schema as any}
      uiSchema={uiSchema as any}
      validator={validator}
      formData={value}
      fields={{ %s, %s }}
      onChange={(e) => onChange(e.formData)}
    >
      <></>
    </Form>
  );
}

export default %sForm;
`,
		className,
		className, className, className,
		fieldsImport, toolset.RelationFieldName, toolset.JsonFieldName,
		toolset.RelationFieldName, toolset.JsonFieldName, fieldsImport,
		className, className,
		className,
		className, className,
		toolset.RelationFieldName, toolset.JsonFieldName,
		className,
	)
}
