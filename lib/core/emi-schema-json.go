package core

import "strings"

// jsonSchemaPrimitive maps one of Emi's own scalar field types (with any
// trailing "?" already stripped) to the JSON Schema "type" keyword a tool's
// inputSchema/outputSchema needs - EmiFieldToJSONSchema and
// EmiQueryFieldToEmiField both funnel through this one table rather than
// each declaring their own copy.
func jsonSchemaPrimitive(base string) (jsonType string, ok bool) {
	switch base {
	case string(FieldTypeString):
		return "string", true
	case string(FieldTypeBool):
		return "boolean", true
	case string(FieldTypeInt), string(FieldTypeInt32), string(FieldTypeInt64):
		return "integer", true
	case string(FieldTypeFloat32), string(FieldTypeFloat64):
		return "number", true
	case "bytes":
		return "string", true
	default:
		return "", false
	}
}

// EmiFieldRequired reports whether a field belongs in a JSON Schema's
// "required" array - every "?"-suffixed type is explicitly optional by
// convention (see e.g. preprocess-entities.go's own nullableFieldType), and
// "any" carries no meaningful default/required semantics of its own either.
func EmiFieldRequired(t FieldType) bool {
	s := string(t)
	return s != "" && t != FieldTypeAny && !strings.HasSuffix(s, "?")
}

// EmiFieldToJSONSchema converts a single field into a JSON Schema fragment,
// recursing into nested Fields for an object/array-of-object field. A branch
// that can't be reasonably represented from field metadata alone (a one/
// collection relation, an opaque complex type) degrades to a schema-less {}
// (accepts anything) rather than guessing a shape and being wrong - the
// field's own Description is still carried over either way, so the model at
// least knows what the argument is for.
func EmiFieldToJSONSchema(f *EmiField) map[string]any {
	if f == nil {
		return map[string]any{}
	}

	base := strings.TrimSuffix(string(f.Type), "?")
	schema := map[string]any{}
	if f.Description != "" {
		schema["description"] = f.Description
	}

	switch base {
	case string(FieldTypeEnum):
		schema["type"] = "string"
		if len(f.OfType) > 0 {
			keys := make([]string, 0, len(f.OfType))
			for _, o := range f.OfType {
				if o != nil {
					keys = append(keys, o.Key)
				}
			}
			schema["enum"] = keys
		}

	case string(FieldTypeArray), string(FieldTypeSlice), string(FieldTypeList):
		schema["type"] = "array"
		switch {
		case f.Primitive != "":
			if t, ok := jsonSchemaPrimitive(f.Primitive); ok {
				schema["items"] = map[string]any{"type": t}
			} else {
				schema["items"] = map[string]any{}
			}
		case len(f.Fields) > 0:
			schema["items"] = EmiFieldsToObjectSchema(f.Fields)
		default:
			schema["items"] = map[string]any{}
		}

	case string(FieldTypeObject), string(FieldTypeClass):
		for k, v := range EmiFieldsToObjectSchema(f.Fields) {
			schema[k] = v
		}

	case string(FieldTypeMap):
		schema["type"] = "object"
		if t, ok := jsonSchemaPrimitive(f.MapPairOf); ok {
			schema["additionalProperties"] = map[string]any{"type": t}
		} else {
			schema["additionalProperties"] = true
		}

	case string(FieldTypeCollection):
		// A has-many relation is represented on the wire as a list of the
		// target's own uniqueIds, never the full related objects.
		schema["type"] = "array"
		schema["items"] = map[string]any{"type": "string"}

	case string(FieldTypeOne):
		// A belongs-to/has-one relation, likewise reduced to the target's
		// own uniqueId.
		schema["type"] = "string"

	case string(FieldTypeAny), string(FieldTypeComplex):
		// No further constraint - see this function's own doc comment.

	default:
		if t, ok := jsonSchemaPrimitive(base); ok {
			schema["type"] = t
		}
	}

	return schema
}

// EmiFieldsToObjectSchema builds a JSON Schema object
// ({"type":"object","properties":{...},"required":[...]}) from a field list
// - the shared core both a dto's own Fields and an intent's inline body
// Fields go through, so there is exactly one place that decides how an Emi
// field becomes a tool argument's schema.
func EmiFieldsToObjectSchema(fields []*EmiField) map[string]any {
	properties := map[string]any{}
	var required []string
	for _, f := range fields {
		if f == nil || f.Name == "" {
			continue
		}
		properties[f.Name] = EmiFieldToJSONSchema(f)
		if EmiFieldRequired(f.Type) {
			required = append(required, f.Name)
		}
	}
	schema := map[string]any{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return schema
}

// EmiQueryFieldToEmiField adapts a query parameter into the same *EmiField
// shape EmiFieldToJSONSchema already knows how to render - EmiQueryField's
// own Type enum is a strict subset of EmiField's, so this is a plain field
// copy, not a translation. Query parameters are always marked optional
// (there's no "?"-suffix convention for EmiQueryField.Type the way body
// fields have), matching how every existing filter/sort/pagination-style
// query field in this codebase already behaves in practice.
func EmiQueryFieldToEmiField(q *EmiQueryField) *EmiField {
	if q == nil {
		return nil
	}
	t := string(q.Type)
	if !strings.HasSuffix(t, "?") {
		t += "?"
	}
	return &EmiField{
		Name:        q.Name,
		Description: q.Description,
		Type:        FieldType(t),
		Primitive:   q.Primitive,
		Fields:      q.Fields,
	}
}

// EmiHeaderToEmiField adapts a typed HTTP header (EmiActionBody.Headers) into
// the same *EmiField shape EmiFieldToJSONSchema already knows how to render -
// mirrors EmiQueryFieldToEmiField's own reasoning: headers are always
// optional (there's no required/optional convention on EmiHeader.Type the
// way body fields have "?"), and the header's own Description is prefixed
// with "HTTP header." so a model reading the merged schema can tell this
// argument travels as a request header rather than a JSON body field - the
// two often use different naming conventions (e.g. "X-Api-Key" vs
// "apiKey"), which the plain property name alone wouldn't convey.
func EmiHeaderToEmiField(h *EmiHeader) *EmiField {
	if h == nil || h.Name == "" {
		return nil
	}
	t := h.Type
	if t == "" {
		t = string(FieldTypeString)
	}
	if !strings.HasSuffix(t, "?") {
		t += "?"
	}
	description := "HTTP header."
	if h.Description != "" {
		description = "HTTP header. " + h.Description
	}
	return &EmiField{
		Name:        h.Name,
		Description: description,
		Type:        FieldType(t),
	}
}

// EmiPathParamsToFields converts a URL's own ":name type" placeholders
// (ExtractPlaceholdersInUrl) into EmiFields - always required (a path
// parameter that's genuinely optional wouldn't be a path parameter at all;
// it'd be a query one), string-typed unless the URL's own type annotation
// says otherwise (e.g. ":uniqueId string", ":page int").
func EmiPathParamsToFields(url string) []*EmiField {
	placeholders := ExtractPlaceholdersInUrl(url)
	if len(placeholders) == 0 {
		return nil
	}
	fields := make([]*EmiField, 0, len(placeholders))
	for _, p := range placeholders {
		fields = append(fields, &EmiField{
			Name: p.Original,
			Type: FieldType(p.Type),
		})
	}
	return fields
}
