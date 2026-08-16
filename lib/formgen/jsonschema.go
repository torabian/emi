package formgen

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// JSONSchema is a minimal draft-07-ish JSON Schema node - just the subset
// react-jsonschema-form (and JSON Schema form renderers in general) actually
// use. It is entirely renderer-agnostic: nothing here is RJSF-specific, so
// any other JSON-Schema-driven form library can consume it the same way
// lib/js's react-jsonschema-form-*.go does.
//
// One relation/collection field and any/complex field both resolve to an
// empty schema ({}) on purpose - this compiler was never given the target
// entity's fields (relations) or has no fixed shape to begin with (any),
// so it deliberately leaves the value unconstrained rather than guessing.
// Pair it with a custom field/widget on the renderer side (see
// lib/js/react-jsonschema-form-toolset.go) to give it real behaviour.
type JSONSchema struct {
	Type                 string
	Title                string
	Description          string
	Default              any
	Properties           []SchemaProperty
	Required             []string
	Items                *JSONSchema
	AdditionalProperties *JSONSchema
	OneOf                []JSONSchemaConst
}

// SchemaProperty is one named entry of a `type: object` schema's
// `properties`. Kept as an ordered slice (not a map) because property order
// is exactly the dto's field order, and Go map iteration order is random -
// marshaling straight through a map would make every regeneration reorder
// fields for no reason. See MarshalJSON.
type SchemaProperty struct {
	Key    string
	Schema *JSONSchema
}

// JSONSchemaConst is one option of a `oneOf` enum: `{ "const": ..., "title": ... }`,
// RJSF's supported way to pair an enum value with a human label (the older
// sibling `enumNames` array is RJSF-only and not part of JSON Schema itself).
type JSONSchemaConst struct {
	Const string `json:"const"`
	Title string `json:"title,omitempty"`
}

// MarshalJSON writes the schema's fields in a fixed, stable order (including
// Properties, in dto field order) so regenerating from an unchanged dto
// always produces byte-identical output.
func (s *JSONSchema) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer
	buf.WriteByte('{')
	first := true

	writeRaw := func(key string, raw []byte) {
		if !first {
			buf.WriteByte(',')
		}
		first = false
		keyBytes, _ := json.Marshal(key)
		buf.Write(keyBytes)
		buf.WriteByte(':')
		buf.Write(raw)
	}

	write := func(key string, value any) error {
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		writeRaw(key, valueBytes)
		return nil
	}

	if s.Type != "" {
		if err := write("type", s.Type); err != nil {
			return nil, err
		}
	}
	if s.Title != "" {
		if err := write("title", s.Title); err != nil {
			return nil, err
		}
	}
	if s.Description != "" {
		if err := write("description", s.Description); err != nil {
			return nil, err
		}
	}
	if s.Default != nil {
		if err := write("default", s.Default); err != nil {
			return nil, err
		}
	}
	if len(s.Properties) > 0 {
		var pbuf bytes.Buffer
		pbuf.WriteByte('{')
		for i, p := range s.Properties {
			if i > 0 {
				pbuf.WriteByte(',')
			}
			keyBytes, _ := json.Marshal(p.Key)
			pbuf.Write(keyBytes)
			pbuf.WriteByte(':')
			valueBytes, err := json.Marshal(p.Schema)
			if err != nil {
				return nil, err
			}
			pbuf.Write(valueBytes)
		}
		pbuf.WriteByte('}')
		writeRaw("properties", pbuf.Bytes())
	}
	if len(s.Required) > 0 {
		if err := write("required", s.Required); err != nil {
			return nil, err
		}
	}
	if s.Items != nil {
		if err := write("items", s.Items); err != nil {
			return nil, err
		}
	}
	if s.AdditionalProperties != nil {
		if err := write("additionalProperties", s.AdditionalProperties); err != nil {
			return nil, err
		}
	}
	if len(s.OneOf) > 0 {
		if err := write("oneOf", s.OneOf); err != nil {
			return nil, err
		}
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// BuildJSONSchema converts a dto's fields into a root `type: object` schema.
// title/description are typically the dto's class name and EmiDto.Description.
func BuildJSONSchema(title, description string, fields []*core.EmiField) *JSONSchema {
	plan := BuildFormPlan(title, fields)
	schema := objectSchema(plan.Fields)
	schema.Title = title
	schema.Description = description
	return schema
}

func objectSchema(fields []*FieldPlan) *JSONSchema {
	s := &JSONSchema{Type: "object"}
	for _, f := range fields {
		s.Properties = append(s.Properties, SchemaProperty{Key: f.Name, Schema: fieldSchema(f)})
		if !f.Nullable {
			s.Required = append(s.Required, f.Name)
		}
	}
	return s
}

func fieldSchema(f *FieldPlan) *JSONSchema {
	var s *JSONSchema

	switch f.Widget {
	case WidgetText:
		s = &JSONSchema{Type: "string"}
	case WidgetNumber:
		s = &JSONSchema{Type: numberJSONType(f.Field)}
	case WidgetCheckbox:
		s = &JSONSchema{Type: "boolean"}
	case WidgetSelect:
		s = enumSchema(f)
	case WidgetPrimitiveList:
		s = &JSONSchema{Type: "array", Items: primitiveSchema(f.Primitive)}
	case WidgetMap:
		s = &JSONSchema{Type: "object", AdditionalProperties: primitiveSchema(f.MapPairOf)}
	case WidgetObjectGroup:
		s = objectSchema(f.Children)
	case WidgetArrayGroup:
		s = &JSONSchema{Type: "array", Items: objectSchema(f.Children)}
	case WidgetOneRelation:
		// Left unconstrained on purpose - see the JSONSchema doc comment.
		s = &JSONSchema{}
	case WidgetCollectionRelation:
		s = &JSONSchema{Type: "array", Items: &JSONSchema{}}
	default: // WidgetJson and anything this package hasn't been taught yet.
		s = &JSONSchema{}
	}

	s.Title = HumanizeLabel(f.Name)
	if f.Field != nil {
		s.Description = f.Field.Description
	}
	return s
}

func numberJSONType(field *core.EmiField) string {
	if field != nil && strings.Contains(string(field.Type), "int") {
		return "integer"
	}
	return "number"
}

func primitiveSchema(primitive string) *JSONSchema {
	switch primitive {
	case "string", "bytes":
		return &JSONSchema{Type: "string"}
	case "bool":
		return &JSONSchema{Type: "boolean"}
	case "int", "int32", "int64":
		return &JSONSchema{Type: "integer"}
	case "float32", "float64":
		return &JSONSchema{Type: "number"}
	default:
		return &JSONSchema{}
	}
}

func enumSchema(f *FieldPlan) *JSONSchema {
	s := &JSONSchema{Type: "string"}
	for _, o := range f.EnumOptions {
		s.OneOf = append(s.OneOf, JSONSchemaConst{Const: o.Value, Title: o.Label})
	}
	return s
}
