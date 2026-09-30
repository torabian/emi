package core

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

const (
	fieldContextEntity = "entity"
	fieldContextDto    = "dto"
)

// resolveFieldUses replaces every `use: <name>` on a field with the properties of the
// field template of that name (Emi.Templates.Fields). Properties the field sets itself
// win over the template's; maps (tags) are merged with the field's own keys winning.
//
// Which template is picked depends on where the field sits: inside an entity it prefers
// a template with context "entity", anywhere else (dtos, action bodies) one with context
// "dto"; a template without a context matches both. An unknown name is an error.
//
// Runs right after extends are resolved, so templates can come from an extended file.
func (m *Emi) resolveFieldUses() error {
	var templates []EmiFieldTemplate
	if m.Templates != nil {
		templates = m.Templates.Fields
	}

	// The templates themselves are definitions, not usages - keep them out of the walk.
	var saved []EmiFieldTemplate
	if m.Templates != nil {
		saved = m.Templates.Fields
		m.Templates.Fields = nil
		defer func() { m.Templates.Fields = saved }()
	}

	return walkFieldUses(reflect.ValueOf(m), fieldContextDto, templates)
}

func walkFieldUses(v reflect.Value, ctx string, templates []EmiFieldTemplate) error {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return nil
		}
		return walkFieldUses(v.Elem(), ctx, templates)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := walkFieldUses(v.Index(i), ctx, templates); err != nil {
				return err
			}
		}
	case reflect.Struct:
		switch v.Type() {
		case reflect.TypeOf(Module3Entity{}):
			ctx = fieldContextEntity
		case reflect.TypeOf(EmiAction{}):
			// Entity-scoped actions still describe request/response bodies, not columns.
			ctx = fieldContextDto
		case reflect.TypeOf(EmiField{}):
			field := v.Addr().Interface().(*EmiField)
			if field.Use != "" {
				if err := applyFieldTemplate(field, ctx, templates); err != nil {
					return err
				}
			}
		}
		for i := 0; i < v.NumField(); i++ {
			if !v.Type().Field(i).IsExported() {
				continue
			}
			if err := walkFieldUses(v.Field(i), ctx, templates); err != nil {
				return err
			}
		}
	}
	return nil
}

func applyFieldTemplate(field *EmiField, ctx string, templates []EmiFieldTemplate) error {
	var chosen *EmiFieldTemplate
	for i := range templates {
		t := &templates[i]
		if t.Name != field.Use {
			continue
		}
		if t.Context == ctx {
			chosen = t
			break
		}
		if t.Context == "" && chosen == nil {
			chosen = t
		}
	}
	if chosen == nil {
		return fmt.Errorf("field %q: unknown field template %q for %s context (declare it under templates.fields, possibly in an extended file)", field.Name, field.Use, ctx)
	}

	// Round-tripping through yaml gives every use its own deep copy of the template, so
	// fields built from it never share slices or maps with it or with each other.
	raw, err := yaml.Marshal(chosen.EmiField)
	if err != nil {
		return err
	}
	var base EmiField
	if err := yaml.Unmarshal(raw, &base); err != nil {
		return err
	}

	dst := reflect.ValueOf(field).Elem()
	src := reflect.ValueOf(&base).Elem()
	for i := 0; i < dst.NumField(); i++ {
		switch dst.Type().Field(i).Name {
		case "Name", "Use", "Description", "Descriptions", "Label", "Labels":
			// Each plain string travels with its locale map, see below.
			continue
		}
		d, s := dst.Field(i), src.Field(i)
		switch {
		case d.IsZero():
			d.Set(s)
		case d.Kind() == reflect.Map && !s.IsZero():
			for _, key := range s.MapKeys() {
				if !d.MapIndex(key).IsValid() {
					d.SetMapIndex(key, s.MapIndex(key))
				}
			}
		}
	}
	// A field's own description (plain or localized) always wins over the template's.
	if field.Description == "" && len(field.Descriptions) == 0 {
		field.Description, field.Descriptions = base.Description, base.Descriptions
	}
	if field.Label == "" && len(field.Labels) == 0 {
		field.Label, field.Labels = base.Label, base.Labels
	}
	field.Use = ""
	return nil
}
