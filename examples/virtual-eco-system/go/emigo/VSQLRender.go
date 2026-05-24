// Package vsql renders Go-template SQL files against typed parameters.
//
// The renderer is reflection-driven and recognizes two contracts:
//
//   - SQLValuer (in this package): a type implementing it controls its own
//     SQL literal and may opt the column out entirely.
//   - Nullable[T]: any struct exposing IsSet() bool and Get() (*T, bool) —
//     including emigo.Nullable[T] — is duck-typed as "optional": unset means
//     the column is omitted, set-to-nil means SQL NULL.
//
// Nested structs and slices that fall through both contracts are serialized
// as JSON and quoted as '…'::jsonb literals, which works for Postgres
// jsonb columns.
package emigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

// SQLValuer lets a type fully control how it renders in the value position of
// an INSERT/UPDATE, and whether the column should be included at all.
// Returning include=false omits the field entirely from the generated SQL.
type SQLValuer interface {
	SQLValue() (value string, include bool)
}

// Fields is the rendered projection of a struct's columns.
// Columns/Values are the INSERT-side projections; Pairs is the UPDATE-side
// "col = val" projection. All three skip fields whose value reported
// include=false (unset Nullable[T], opted-out SQLValuer).
type Fields struct {
	Columns string
	Values  string
	Pairs   string
}

// VSQLRender parses query as a Go template with the SQL helpers registered
// (sqlFields, sqlFieldsExcept, sql, deref) and executes it against params.
func VSQLRender(query string, params any) (string, error) {
	return renderTemplate("vsql", query, params)
}

// VSQLRenderFS reads the named file from fsys and renders it the same way as
// [VSQLRender]. Use this when queries are embedded via embed.FS or loaded
// from another fs.FS.
func VSQLRenderFS(fsys fs.FS, name string, params any) (string, error) {
	raw, err := fs.ReadFile(fsys, name)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", name, err)
	}
	return renderTemplate(name, string(raw), params)
}

func renderTemplate(name, body string, params any) (string, error) {
	t, err := template.New(name).Funcs(FuncMap()).Parse(body)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", name, err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, params); err != nil {
		return "", fmt.Errorf("execute %s: %w", name, err)
	}
	return buf.String(), nil
}

// FuncMap returns the template helpers used by Render. Exposed so callers can
// register them on their own *template.Template instances if they don't want
// to use Render directly.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"sqlFields":       sqlFieldsFunc,
		"sqlFieldsExcept": sqlFieldsExceptFunc,
		"sql":             sqlValueFunc,
		"deref":           derefFunc,
	}
}

func sqlFieldsFunc(in any) (Fields, error) {
	return collectFields(in, nil)
}

// sqlFieldsExceptFunc is like sqlFieldsFunc but omits the named columns.
// Useful when a field will be expanded into a separate table via {{ range }}
// rather than serialized inline.
func sqlFieldsExceptFunc(in any, exclude ...string) (Fields, error) {
	skip := make(map[string]struct{}, len(exclude))
	for _, e := range exclude {
		skip[e] = struct{}{}
	}
	return collectFields(in, skip)
}

func collectFields(in any, skip map[string]struct{}) (Fields, error) {
	v := reflect.ValueOf(in)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return Fields{}, fmt.Errorf("sqlFields: nil pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return Fields{}, fmt.Errorf("sqlFields: expected struct, got %s", v.Kind())
	}
	cols, vals, pairs, err := walkFields(v, "", skip)
	if err != nil {
		return Fields{}, err
	}
	return Fields{
		Columns: strings.Join(cols, ", "),
		Values:  strings.Join(vals, ", "),
		Pairs:   strings.Join(pairs, ", "),
	}, nil
}

// walkFields walks v's exported fields, prepending `prefix_` to each column
// name. When it encounters a nested struct (or set Nullable[Struct]) that is
// NOT a SQLValuer, it recurses and inlines the inner fields under the parent
// column name — so a Preferences{Theme: "dark"} field becomes a
// `preferences_theme` column rather than a JSON blob. SQLValuer types keep
// their single-column custom rendering.
func walkFields(v reflect.Value, prefix string, skip map[string]struct{}) ([]string, []string, []string, error) {
	var cols, vals, pairs []string
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		colName := f.Tag.Get("db")
		if colName == "-" {
			continue
		}
		if colName == "" {
			colName = toSnakeCase(f.Name)
		}
		fullName := colName
		if prefix != "" {
			fullName = prefix + "_" + colName
		}
		if _, omit := skip[fullName]; omit {
			continue
		}
		if _, omit := skip[f.Name]; omit {
			continue
		}
		// Also match the json tag name (camelCase, the form most templates
		// will naturally use — e.g. sqlFieldsExcept . "arrayField").
		if jt := f.Tag.Get("json"); jt != "" {
			jsonName := strings.SplitN(jt, ",", 2)[0]
			if _, omit := skip[jsonName]; omit {
				continue
			}
		}

		fieldVal := v.Field(i)

		flatten, inner, omit := shouldFlatten(fieldVal)
		if omit {
			continue
		}
		if flatten {
			subCols, subVals, subPairs, err := walkFields(inner, fullName, skip)
			if err != nil {
				return nil, nil, nil, fmt.Errorf("field %s: %w", f.Name, err)
			}
			cols = append(cols, subCols...)
			vals = append(vals, subVals...)
			pairs = append(pairs, subPairs...)
			continue
		}

		val, include, err := renderValue(fieldVal)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("field %s: %w", f.Name, err)
		}
		if !include {
			continue
		}
		cols = append(cols, fullName)
		vals = append(vals, val)
		pairs = append(pairs, fullName+" = "+val)
	}
	return cols, vals, pairs, nil
}

// shouldFlatten decides whether to inline a sub-struct's fields into the
// parent column list (prefix_child names). Returns:
//   - flatten: true if the value is a struct (or set Nullable[Struct]) that
//     should be walked recursively
//   - inner: the struct value to walk
//   - omit: true if the field should be skipped entirely (unset
//     Nullable[Struct] — every prefixed column is omitted, matching the
//     unset-scalar PATCH semantics)
//
// SQLValuer types always lose to their custom rendering — they stay as a
// single column, never flattened. Slices/arrays/maps are not flattened
// here; the caller handles them via {{ range }} or sqlFieldsExcept.
func shouldFlatten(v reflect.Value) (flatten bool, inner reflect.Value, omit bool) {
	if v.CanInterface() {
		if _, ok := v.Interface().(SQLValuer); ok {
			return false, reflect.Value{}, false
		}
	}
	if v.CanAddr() {
		if _, ok := v.Addr().Interface().(SQLValuer); ok {
			return false, reflect.Value{}, false
		}
	}
	if isSet, innerVal, isNullable := unwrapNullable(v); isNullable {
		if !isSet {
			return false, reflect.Value{}, true
		}
		if !innerVal.IsValid() {
			return false, reflect.Value{}, false
		}
		if innerVal.Kind() == reflect.Struct {
			return true, innerVal, false
		}
		return false, reflect.Value{}, false
	}
	if v.Kind() == reflect.Struct {
		return true, v, false
	}
	return false, reflect.Value{}, false
}

// sqlValueFunc renders a single value as a SQL literal. Useful for ad-hoc
// interpolation in templates: {{ sql .SomeValue }}.
func sqlValueFunc(in any) (string, error) {
	val, _, err := renderValue(reflect.ValueOf(in))
	return val, err
}

// derefFunc unwraps a Nullable[T] so the inner T can be ranged over or
// otherwise traversed by the template. Returns nil for an unset Nullable
// (templates' {{ range nil }} produces zero iterations), or passes the value
// through unchanged if it isn't a Nullable.
func derefFunc(in any) any {
	rv := reflect.ValueOf(in)
	if isSet, inner, ok := unwrapNullable(rv); ok {
		if !isSet || !inner.IsValid() {
			return nil
		}
		return inner.Interface()
	}
	return in
}

// renderValue dispatches by priority:
//  1. Implements SQLValuer  -> full custom control (incl. opt-out)
//  2. Looks like Nullable[T] -> conditional (unset = omit, null-set = NULL)
//  3. Primitive / struct / slice -> default literal or JSON encoding
func renderValue(v reflect.Value) (string, bool, error) {
	if !v.IsValid() {
		return "NULL", true, nil
	}

	if v.CanInterface() {
		if sv, ok := v.Interface().(SQLValuer); ok {
			value, include := sv.SQLValue()
			return value, include, nil
		}
	}
	if v.CanAddr() {
		if sv, ok := v.Addr().Interface().(SQLValuer); ok {
			value, include := sv.SQLValue()
			return value, include, nil
		}
	}

	// Emi emits interface{} for `any`, `map`, `map?`, `one?`, `collection?`.
	// Unwrap so the rest of the dispatch sees the real dynamic type.
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "NULL", true, nil
		}
		return renderValue(v.Elem())
	}

	if isSet, inner, ok := unwrapNullable(v); ok {
		if !isSet {
			return "", false, nil
		}
		if !inner.IsValid() {
			return "NULL", true, nil
		}
		return renderValue(inner)
	}

	return renderPrimitive(v)
}

// unwrapNullable duck-types emigo.Nullable[T] by looking for IsSet() bool and
// Get() (*T, bool). This avoids a hard dependency on the specific generic
// instantiation and lets any compatible "Nullable-like" type plug in.
func unwrapNullable(v reflect.Value) (isSet bool, inner reflect.Value, ok bool) {
	if v.Kind() != reflect.Struct {
		return false, reflect.Value{}, false
	}
	isSetM := v.MethodByName("IsSet")
	getM := v.MethodByName("Get")
	if !isSetM.IsValid() || !getM.IsValid() {
		return false, reflect.Value{}, false
	}
	if isSetM.Type().NumIn() != 0 || isSetM.Type().NumOut() != 1 ||
		isSetM.Type().Out(0).Kind() != reflect.Bool {
		return false, reflect.Value{}, false
	}
	gt := getM.Type()
	if gt.NumIn() != 0 || gt.NumOut() != 2 ||
		gt.Out(0).Kind() != reflect.Ptr || gt.Out(1).Kind() != reflect.Bool {
		return false, reflect.Value{}, false
	}

	if !isSetM.Call(nil)[0].Bool() {
		return false, reflect.Value{}, true
	}
	ptr := getM.Call(nil)[0]
	if ptr.IsNil() {
		return true, reflect.Value{}, true
	}
	return true, ptr.Elem(), true
}

func renderPrimitive(v reflect.Value) (string, bool, error) {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "NULL", true, nil
		}
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.String:
		return sqlQuote(v.String()), true, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int()), true, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint()), true, nil
	case reflect.Float32:
		// Preserve float32 precision — %g on float64 would expose the
		// extra bits introduced by the promotion (3.14 → 3.140000104…).
		return strconv.FormatFloat(v.Float(), 'g', -1, 32), true, nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64), true, nil
	case reflect.Bool:
		if v.Bool() {
			return "TRUE", true, nil
		}
		return "FALSE", true, nil
	case reflect.Slice, reflect.Array:
		if v.Kind() == reflect.Slice && v.IsNil() {
			return "NULL", true, nil
		}
		return jsonAsJsonb(v.Interface(), v.Type())
	case reflect.Map:
		if v.IsNil() {
			return "NULL", true, nil
		}
		if !v.CanInterface() {
			break
		}
		// json/encoding only allows maps with string-ish keys, so
		// map[any]any (emi's default for `map` / `any`) blows up. Fall back
		// to a stringified-key copy when the direct marshal refuses.
		if b, err := json.Marshal(v.Interface()); err == nil {
			return sqlQuote(string(b)) + "::jsonb", true, nil
		}
		copy := make(map[string]any, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			k := iter.Key()
			for k.Kind() == reflect.Interface {
				k = k.Elem()
			}
			copy[fmt.Sprint(k.Interface())] = iter.Value().Interface()
		}
		return jsonAsJsonb(copy, v.Type())
	case reflect.Struct:
		return jsonAsJsonb(v.Interface(), v.Type())
	}
	return "", false, fmt.Errorf("renderPrimitive: unsupported kind %s (type %s)", v.Kind(), v.Type())
}

func jsonAsJsonb(v any, t reflect.Type) (string, bool, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", false, fmt.Errorf("json encode %s: %w", t, err)
	}
	return sqlQuote(string(b)) + "::jsonb", true, nil
}

func sqlQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "''") + "'"
}

func toSnakeCase(s string) string {
	var b strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			b.WriteByte('_')
		}
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}
