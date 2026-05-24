package core

import (
	"fmt"
	"strings"
)

// SelfFieldsToken is the special Include entry that, when found inside an
// EmiCapture, inlines the owner's own declared fields at that position.
const SelfFieldsToken = "self.fields"

// Preprocess resolves derived/derived-like definitions on the module so the
// downstream generators receive a fully-expanded structure. Currently this:
//
//   - Flattens EmiVsql.Captures into EmiVsql.Params and clears Captures.
//
// It is safe to call more than once: a second pass is a no-op once captures
// have been consumed.
func (m *Emi) Preprocess() error {
	if m == nil {
		return nil
	}

	dtoByName := make(map[string]*EmiDto, len(m.Dto))
	for i := range m.Dto {
		dtoByName[m.Dto[i].Name] = &m.Dto[i]
	}

	templateDtoByName := map[string]*EmiDto{}
	if m.Templates != nil {
		for i := range m.Templates.Dtos {
			templateDtoByName[m.Templates.Dtos[i].Name] = &m.Templates.Dtos[i]
		}
	}

	actionByName := make(map[string]*EmiAction)
	for _, a := range m.Actions {
		if a != nil && a.Name != "" {
			actionByName[a.Name] = a
		}
	}
	if m.Templates != nil {
		for _, a := range m.Templates.Actions {
			if a != nil && a.Name != "" {
				actionByName[a.Name] = a
			}
		}
	}

	for i := range m.Vsqls {
		v := &m.Vsqls[i]
		if len(v.Captures) == 0 {
			continue
		}
		merged, err := resolveCaptures(v.Captures, v.Params, dtoByName, templateDtoByName, actionByName, fmt.Sprintf("vsql %q", v.Name))
		if err != nil {
			return err
		}
		v.Params = merged
		v.Captures = nil
	}
	return nil
}

// resolveCaptures applies a list of EmiCapture entries against the owner's
// inline fields, returning the merged field list. Duplicate field names (by
// EmiField.Name) are kept on first occurrence — later entries are skipped so
// the order remains deterministic and explicit.
func resolveCaptures(captures []*EmiCapture, selfFields []*EmiField, dtoByName, templateDtoByName map[string]*EmiDto, actionByName map[string]*EmiAction, owner string) ([]*EmiField, error) {
	var out []*EmiField
	seen := make(map[string]bool)

	add := func(f *EmiField) {
		if f == nil || seen[f.Name] {
			return
		}
		seen[f.Name] = true
		out = append(out, f)
	}

	selfConsumed := false

	for capIdx, c := range captures {
		if c == nil {
			continue
		}
		if c.Action != "" && (c.Dto != "" || c.Template != "") {
			return nil, fmt.Errorf("%s: capture #%d: action cannot be combined with dto/template", owner, capIdx)
		}

		var (
			source     []*EmiField
			sourceDesc string
		)
		switch {
		case c.Dto != "":
			dtoName := strings.TrimPrefix(c.Dto, "dto/")
			dto, ok := dtoByName[dtoName]
			if !ok {
				return nil, fmt.Errorf("%s: capture #%d: dto %q not found", owner, capIdx, dtoName)
			}
			source = dto.Fields
			sourceDesc = "dto " + dtoName
		case c.Template != "":
			tplName := strings.TrimPrefix(c.Template, "dto/")
			dto, ok := templateDtoByName[tplName]
			if !ok {
				return nil, fmt.Errorf("%s: capture #%d: template %q not found", owner, capIdx, tplName)
			}
			source = dto.Fields
			sourceDesc = "template " + tplName
		case c.Action != "":
			fields, desc, err := lookupActionFields(c.Action, actionByName)
			if err != nil {
				return nil, fmt.Errorf("%s: capture #%d: %w", owner, capIdx, err)
			}
			source = fields
			sourceDesc = desc
		}

		excludes := make(map[string]bool, len(c.Exclude))
		for _, e := range c.Exclude {
			excludes[e] = true
		}

		if len(c.Include) == 0 {
			if c.Dto == "" && c.Template == "" && c.Action == "" {
				return nil, fmt.Errorf("%s: capture #%d: must set dto, template, action, or include", owner, capIdx)
			}
			for _, f := range source {
				if excludes[f.Name] {
					continue
				}
				add(f)
			}
			continue
		}

		byName := make(map[string]*EmiField, len(source))
		for _, f := range source {
			byName[f.Name] = f
		}
		for _, inc := range c.Include {
			if inc == SelfFieldsToken {
				selfConsumed = true
				for _, sf := range selfFields {
					if excludes[sf.Name] {
						continue
					}
					add(sf)
				}
				continue
			}
			f, ok := byName[inc]
			if !ok {
				return nil, fmt.Errorf("%s: capture #%d: include field %q not found in %s", owner, capIdx, inc, sourceDesc)
			}
			if excludes[inc] {
				continue
			}
			add(f)
		}
	}

	if !selfConsumed {
		for _, sf := range selfFields {
			add(sf)
		}
	}

	return out, nil
}

// lookupActionFields resolves a capture.Action ref of the form "name",
// "name.in", or "name.out" against the merged action lookup map.
func lookupActionFields(ref string, actionByName map[string]*EmiAction) ([]*EmiField, string, error) {
	name, side, _ := strings.Cut(ref, ".")
	if side == "" {
		side = "in"
	}
	if side != "in" && side != "out" {
		return nil, "", fmt.Errorf("action %q: side must be 'in' or 'out', got %q", ref, side)
	}
	a, ok := actionByName[name]
	if !ok {
		return nil, "", fmt.Errorf("action %q not found", name)
	}
	body := a.In
	if side == "out" {
		body = a.Out
	}
	if body == nil {
		return nil, "", fmt.Errorf("action %q has no %s body", name, side)
	}
	return body.Fields, fmt.Sprintf("action %s.%s", name, side), nil
}
