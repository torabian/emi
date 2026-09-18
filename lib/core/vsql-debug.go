package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// vsqlDebugTemplateFuncs is the (deliberately small) function set available
// to a query template rendered via RenderVsqlDebug: "sql" quotes a Go value
// as a SQL literal. Debug rendering works against plain map[string]any data
// (see buildVsqlDebugFields), not the real generated Go structs a compiled
// project would have - so the reflection-based sqlFields/sqlFieldsExcept
// helpers (see examples/vsql/vsql.FuncMap) aren't available here; write a
// debug query using range/if directly, the way every vsql example in this
// repo does.
func vsqlDebugTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"sql": func(v any) string {
			switch t := v.(type) {
			case nil:
				return "NULL"
			case string:
				return "'" + strings.ReplaceAll(t, "'", "''") + "'"
			case bool:
				if t {
					return "true"
				}
				return "false"
			default:
				return fmt.Sprint(t)
			}
		},
	}
}

// RenderVsqlDebug renders one vsql's query template against caller-supplied
// values, exactly the way `emi vsql:debug` does - useful directly from Go
// tests/tools too, not just the CLI.
//
// paramsJSON/inJSON are each an optional JSON object of field values, keyed
// by field Name (e.g. `{"id": 42}`); both are merged into one set of values
// before building template data, with paramsJSON winning a name collision -
// the same precedence Preprocess gives Params over In. selected is an
// optional comma-separated list of column Names to mark Selected; when
// empty, every column's own declared `selected:` default is used instead.
//
// moduleDir anchors QueryName when the vsql reads its SQL from disk instead
// of an inline Query (pass "" if the vsql has no QueryName, or if QueryName
// is already absolute).
func RenderVsqlDebug(vsql *EmiVsql, moduleDir, paramsJSON, inJSON, selected string) (string, error) {
	if vsql == nil {
		return "", fmt.Errorf("vsql is nil")
	}

	query := vsql.Query
	if vsql.QueryName != "" {
		p := vsql.QueryName
		if !filepath.IsAbs(p) && moduleDir != "" {
			p = filepath.Join(moduleDir, p)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return "", fmt.Errorf("reading queryName %q: %w", vsql.QueryName, err)
		}
		query = string(b)
	}
	if query == "" {
		return "", fmt.Errorf("vsql %q has neither query nor queryName set", vsql.Name)
	}

	values, err := mergeJSONObjects(paramsJSON, inJSON)
	if err != nil {
		return "", err
	}

	data := coerceFieldsForTemplate(vsql.Params, values)
	if len(vsql.Columns) > 0 {
		data["Columns"] = buildColumnsDebugData(vsql.Columns, selected)
	}

	t, err := template.New(vsql.Name).Funcs(vsqlDebugTemplateFuncs()).Parse(query)
	if err != nil {
		return "", fmt.Errorf("parse query template: %w", err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("render query template: %w", err)
	}
	return buf.String(), nil
}

// mergeJSONObjects parses a and b as JSON objects (either may be empty) and
// combines them into one map, with a's entries winning on a name collision.
func mergeJSONObjects(a, b string) (map[string]any, error) {
	out := map[string]any{}
	if strings.TrimSpace(b) != "" {
		var m map[string]any
		if err := json.Unmarshal([]byte(b), &m); err != nil {
			return nil, fmt.Errorf("--in: invalid JSON: %w", err)
		}
		maps.Copy(out, m)
	}
	if strings.TrimSpace(a) != "" {
		var m map[string]any
		if err := json.Unmarshal([]byte(a), &m); err != nil {
			return nil, fmt.Errorf("--params: invalid JSON: %w", err)
		}
		maps.Copy(out, m)
	}
	return out, nil
}

// coerceFieldsForTemplate builds the template-data map for a field list the
// same way the real generated code's struct would look to a query template:
// each field is keyed by its PascalCased Name (matching an exported Go
// field), with array/object fields recursively reshaped so
// {{ range .Foo.Items }} and {{ .Foo.Bar }} work the same as against a real
// generated struct. A field missing from values is simply omitted, the same
// as a zero-valued Go struct field would be for most practical template
// purposes.
func coerceFieldsForTemplate(fields []*EmiField, values map[string]any) map[string]any {
	out := map[string]any{}
	for _, f := range fields {
		if f == nil || f.Name == "" {
			continue
		}
		raw, ok := values[f.Name]
		if !ok {
			continue
		}
		out[ToUpper(f.Name)] = coerceFieldForTemplate(f, raw)
	}
	return out
}

// coerceFieldForTemplate reshapes one field's raw JSON value to match its
// declared type: an array field becomes {"Items": [...]} (mirroring
// emigo.Array[T]'s own field name), an object field becomes a nested map
// built the same way as the root. Every other type - primitives, enum, map,
// any, complex, one/collection - is passed through unchanged: debug
// rendering has no real Go type to coerce into, and every query in this
// repo's own examples only ever needs a plain value or an array-of-object at
// this level.
func coerceFieldForTemplate(f *EmiField, raw any) any {
	switch strings.TrimSuffix(string(f.Type), "?") {
	case string(FieldTypeArray):
		items, _ := raw.([]any)
		coerced := make([]any, 0, len(items))
		for _, it := range items {
			m, _ := it.(map[string]any)
			coerced = append(coerced, coerceFieldsForTemplate(f.Fields, m))
		}
		return map[string]any{"Items": coerced}
	case string(FieldTypeObject):
		m, _ := raw.(map[string]any)
		return coerceFieldsForTemplate(f.Fields, m)
	default:
		return raw
	}
}

// buildColumnsDebugData builds the {{ .Columns }} value: one
// {"Selected": bool} entry per column keyed by its PascalCased Name, plus a
// "Cols" entry standing in for the real Cols() method (a map has no
// methods, but {{ .Columns.Cols }} looks up a map key exactly the same way
// it would call a zero-argument method on a struct).
//
// selected, when non-empty, is a comma-separated list of column Names that
// are the only ones marked Selected - every other declared column becomes
// unselected, regardless of its own default. Leave it empty to use each
// column's own `selected:` default untouched.
func buildColumnsDebugData(columns []*EmiColumn, selected string) map[string]any {
	var forced map[string]bool
	if s := strings.TrimSpace(selected); s != "" {
		forced = map[string]bool{}
		for name := range strings.SplitSeq(s, ",") {
			if name = strings.TrimSpace(name); name != "" {
				forced[name] = true
			}
		}
	}

	out := map[string]any{}
	var cols []string
	for _, col := range columns {
		if col == nil || col.Name == "" {
			continue
		}
		isSelected := col.Selected
		if forced != nil {
			isSelected = forced[col.Name]
		}
		out[ToUpper(col.Name)] = map[string]any{"Selected": isSelected}
		if isSelected {
			cols = append(cols, col.GetColumn())
		}
	}
	out["Cols"] = strings.Join(cols, ", ")
	return out
}
