package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/torabian/emi/lib/sqlpredict"
)

// SelfFieldsToken is the special Include entry that, when found inside an
// EmiCapture, inlines the owner's own declared fields at that position.
const SelfFieldsToken = "self.fields"

// Preprocess resolves derived/derived-like definitions on the module so the
// downstream generators receive a fully-expanded structure. Currently this:
//
//   - Resolves EmiVsql.In (Dto reference or inline Fields) and merges it into
//     EmiVsql.Params - see resolveVsqlIn. In is left in place afterwards
//     (its Headers/Dto name aren't expressible as plain fields).
//   - Flattens EmiVsql.Captures into EmiVsql.Params and clears Captures.
//   - When EmiVsql.Predict is set and Columns is empty, detects Columns by
//     parsing the query as a plain SELECT - see predictVsqlColumns.
//   - Forces every vsql column that defaults to unselected onto its nullable
//     type variant, regardless of what was declared - see
//     nullifyUnselectedColumns.
//   - When EmiVsql.Filters is empty, defaults it to Columns' own fields -
//     see defaultVsqlFilters.
//   - Runs every hook registered via RegisterPreprocessHook (see
//     preprocess-hooks.go) - e.g. entities' update-dto synthesis in
//     preprocess-entities.go plugs in this way, rather than being called here
//     directly, so this file doesn't need to know entities exist.
//
// It is safe to call more than once: a second pass is a no-op once captures
// have been consumed (and hooks are expected to be similarly idempotent -
// preprocessEntityUpdateDtos, for instance, skips dtos it already synthesized;
// nullifyUnselectedColumns is naturally idempotent since it only touches
// types that aren't already nullable; resolveVsqlIn is idempotent because its
// merge into Params keeps whichever fields are already there by name).
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
		owner := fmt.Sprintf("vsql %q", v.Name)

		if v.Predict && len(v.Columns) == 0 {
			cols, err := predictVsqlColumns(v, m.SourcePath)
			if err != nil {
				return fmt.Errorf("%s: predict: %w", owner, err)
			}
			v.Columns = cols
		}

		nullifyUnselectedColumns(v.Columns)

		if len(v.Filters) == 0 {
			v.Filters = defaultVsqlFilters(v.Columns)
		}

		if v.In != nil {
			inFields, err := resolveVsqlIn(v.In, dtoByName, templateDtoByName, owner)
			if err != nil {
				return err
			}
			v.Params = mergeFieldsKeepFirst(v.Params, inFields)
		}

		if len(v.Captures) == 0 {
			continue
		}
		merged, err := resolveCaptures(v.Captures, v.Params, dtoByName, templateDtoByName, actionByName, owner)
		if err != nil {
			return err
		}
		v.Params = merged
		v.Captures = nil
	}

	ResolvePermissionFullKeys(m.Permissions, "")
	if err := ValidatePermissionIdentifiers(m.Permissions); err != nil {
		return err
	}

	if err := ValidateEventIdentifiers(m.Events); err != nil {
		return err
	}

	if err := runPreprocessHooks(m, globalPreprocessHooks); err != nil {
		return err
	}

	// Runs after globalPreprocessHooks (not before) so that hook-synthesized actions -
	// e.g. entity actions added by PreprocessEntityActions, when a backend opts into it -
	// are already in m.Actions and reachable by an intent's From reference.
	return m.preprocessIntents()
}

// PreprocessForAction runs Preprocess, then any hooks registered specifically for
// action (e.g. "go", "kotlin") via RegisterPreprocessHookForAction. Compiler backends
// that want expansions scoped to just their own output should call this - with their
// own core.BaseAction.Name - instead of Preprocess directly.
func (m *Emi) PreprocessForAction(action string) error {
	if err := m.Preprocess(); err != nil {
		return err
	}
	return runPreprocessHooks(m, actionPreprocessHooks[action])
}

// resolveVsqlIn turns an EmiVsql.In body into a plain field list: In.Dto
// (looked up first against top-level Dtos, then template Dtos) if set,
// otherwise In.Fields verbatim. Headers, Envelope and Primitive aren't
// resolved here - they only matter once a vsql can actually be exposed the
// way an EmiAction is, which isn't wired up yet; In is left on the vsql
// afterwards precisely so that future code still has them available.
func resolveVsqlIn(body *EmiActionBody, dtoByName, templateDtoByName map[string]*EmiDto, owner string) ([]*EmiField, error) {
	if body == nil {
		return nil, nil
	}
	if body.Dto != "" {
		if dto, ok := dtoByName[body.Dto]; ok {
			return dto.Fields, nil
		}
		if dto, ok := templateDtoByName[body.Dto]; ok {
			return dto.Fields, nil
		}
		return nil, fmt.Errorf("%s: in: dto %q not found", owner, body.Dto)
	}
	return body.Fields, nil
}

// mergeFieldsKeepFirst combines two field lists into one, in order, dropping
// any field from b whose Name already appears in a - so Params and In can
// describe the same query without one silently overwriting the other's
// field, and re-running this (e.g. on a second Preprocess pass) is a no-op.
func mergeFieldsKeepFirst(a, b []*EmiField) []*EmiField {
	if len(b) == 0 {
		return a
	}
	seen := make(map[string]bool, len(a))
	for _, f := range a {
		if f != nil {
			seen[f.Name] = true
		}
	}
	out := a
	for _, f := range b {
		if f == nil || seen[f.Name] {
			continue
		}
		seen[f.Name] = true
		out = append(out, f)
	}
	return out
}

// predictVsqlColumns resolves the SQL text to parse (Query verbatim, or
// QueryName read from disk relative to sourcePath's directory) and runs it
// through sqlpredict.DetectSelectColumns, converting each detected column
// into an *EmiColumn: Column is set to the detected Source verbatim (the
// literal expression to project - never derived from Name, since a detected
// name like "TotalOrders" snake-casing to "total_orders" would silently
// drop an aggregate expression like "count(order_id)"), Type from the
// detected type (or its nullable variant when field() marked it optional),
// and Selected: true - the query, as written, already selects every one of
// these unconditionally.
func predictVsqlColumns(v *EmiVsql, sourcePath string) ([]*EmiColumn, error) {
	query := v.Query
	if v.QueryName != "" {
		p := v.QueryName
		if !filepath.IsAbs(p) && sourcePath != "" {
			p = filepath.Join(filepath.Dir(sourcePath), p)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("reading queryName %q: %w", v.QueryName, err)
		}
		query = string(b)
	}
	if query == "" {
		return nil, fmt.Errorf("neither query nor queryName is set")
	}

	detected, err := sqlpredict.DetectSelectColumns(query)
	if err != nil {
		return nil, err
	}

	cols := make([]*EmiColumn, 0, len(detected))
	for _, d := range detected {
		typ := d.Type
		if typ == "" {
			typ = string(FieldTypeString)
		}
		if d.Optional {
			typ += "?"
		}
		cols = append(cols, &EmiColumn{
			EmiField: EmiField{
				Name: ToLower(d.Name),
				Type: FieldType(typ),
			},
			Column:   d.Source,
			Selected: true,
		})
	}
	return cols, nil
}

// defaultVsqlFilters builds the default Filters list from columns: one
// *EmiField per column, copied (not aliased) so a later mutation of one list
// - nullifyUnselectedColumns already ran on columns by the time this is
// called, precisely so Filters' types mirror the row's actual scanned types
// - never reaches back into the other. Returns nil for no columns, which is
// exactly "Filters stays empty" (a query with nothing to select has nothing
// sensible to default filtering to either).
func defaultVsqlFilters(columns []*EmiColumn) []*EmiField {
	if len(columns) == 0 {
		return nil
	}
	fields := make([]*EmiField, 0, len(columns))
	for _, col := range columns {
		if col == nil {
			continue
		}
		f := col.EmiField
		fields = append(fields, &f)
	}
	return fields
}

// nullifyUnselectedColumns forces every column that defaults to unselected
// (Selected: false) onto its nullable type variant, even when the column was
// declared as a non-nullable type. A column the caller can choose to leave
// out of the query has no guarantee of being populated once a result is
// scanned into the row DTO, so the generated field must be able to represent
// "not fetched" regardless of what the column's own type would otherwise be
// - the same way an actually-optional SQL column would be modeled.
//
// FieldTypeAny is left alone: it already renders as interface{}, which is
// unconditionally nil-capable on its own (see its own doc comment in
// EmiFieldType.go). FieldTypeComplex becomes FieldTypeComplexNullable, which
// every backend but the JS/TS generator treats identically to plain complex
// (see FieldTypeComplexNullable's doc comment) - harmless for those, and the
// correct "optional" signal for JS/TS. Every other emi type has a mechanical
// "?"-suffixed nullable counterpart (see EmiFieldType.go), so appending it is
// safe generically. A column already declared nullable, or a column that is
// selected by default, is left untouched.
func nullifyUnselectedColumns(columns []*EmiColumn) {
	for _, col := range columns {
		if col == nil || col.Selected || col.Type == "" {
			continue
		}
		if col.Type == FieldTypeAny || strings.HasSuffix(string(col.Type), "?") {
			continue
		}
		col.Type = FieldType(string(col.Type) + "?")
	}
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
