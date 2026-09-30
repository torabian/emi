package core

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// EmiExtends is one entry of Emi.Extends: a definition file to merge in, plus the
// values for the {{placeholders}} written inside it.
type EmiExtends struct {
	// From is the path of the emi definition file to extend, relative to the file that
	// declares this entry (or absolute).
	From string `yaml:"from,omitempty" json:"from,omitempty" jsonschema:"description=Path of the emi definition file to merge in. Relative paths resolve against the file declaring this entry."`

	// Params replace {{name}} placeholders in every string value of the extended file
	// (including the params it passes on to files it extends itself). A placeholder with
	// no matching param is left untouched, unless it has a fallback: {{name|default}}.
	Params map[string]string `yaml:"params,omitempty" json:"params,omitempty" jsonschema:"description=Values for {{name}} placeholders in string values of the extended file; {{name|fallback}} sets a default. {{extends.dir}} is built in - the extended file directory relative to the extending module. Unmatched placeholders are left as written."`
}

var placeholderPattern = regexp.MustCompile(`\{\{\s*([A-Za-z_][A-Za-z0-9_.-]*)\s*(?:\|([^}]*))?\}\}`)

// applyParams replaces {{key}} in every string reachable from v (v must be a pointer)
// with params[key]. Substitution happens on the parsed structure, not the raw yaml, so a
// param value can never break the document's syntax.
func applyParams(v any, params map[string]string) {
	substituteValue(reflect.ValueOf(v), params)
}

func substituteString(s string, params map[string]string) string {
	return placeholderPattern.ReplaceAllStringFunc(s, func(match string) string {
		groups := placeholderPattern.FindStringSubmatch(match)
		if val, ok := params[groups[1]]; ok {
			return val
		}
		// {{key|fallback}} - the fallback applies when the param isn't passed.
		if strings.Contains(match, "|") {
			return strings.TrimSpace(groups[2])
		}
		return match
	})
}

func substituteValue(v reflect.Value, params map[string]string) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return
		}
		if v.Kind() == reflect.Interface {
			// An interface's content isn't addressable: substitute in a copy and put it back.
			elem := reflect.New(v.Elem().Type()).Elem()
			elem.Set(v.Elem())
			substituteValue(elem, params)
			if v.CanSet() {
				v.Set(elem)
			}
			return
		}
		substituteValue(v.Elem(), params)
	case reflect.String:
		if v.CanSet() {
			v.SetString(substituteString(v.String(), params))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).IsExported() {
				substituteValue(v.Field(i), params)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			substituteValue(v.Index(i), params)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			elem := reflect.New(v.Type().Elem()).Elem()
			elem.Set(v.MapIndex(key))
			substituteValue(elem, params)
			v.SetMapIndex(key, elem)
		}
	}
}

// resolveExtends merges every file listed in m.Extends into m, then clears Extends.
//
// Order: the extended files are applied in declaration order, each overriding what the
// previous ones set, and m's own content is applied last - so the module that declares
// `extends` always wins. Extended files are resolved recursively (an extended file may
// itself extend others, relative to its own directory). A file that ends up extending
// itself, directly or through a chain, is rejected; the same file reached twice through
// different branches (a diamond) is fine.
//
// Merge rules:
//   - Scalars (name, namespace, description, version): last non-empty value wins.
//   - Lists (targets, complexes, enums, dtos, entities, actions, ...): entries are
//     matched by identity (usually name; targets: compiler + cleaned output path). A match replaces the earlier entry in place;
//     anything else is appended. Order of first appearance is preserved.
//   - templates: dtos and actions merge the same way, by name.
//
// Besides Params, {{extends.dir}} is always available - see withImplicitParams.
//
// Each entry's Params are substituted into the extended file (see EmiExtends.Params)
// before it is merged, so one shared definition can serve many modules.
//
// Paths inside extended content (e.g. a target's output) are NOT rewritten - they are
// interpreted relative to the module that finally extends them.
func (m *Emi) resolveExtends() error {
	if len(m.Extends) == 0 {
		return nil
	}

	baseDir := ""
	visited := map[string]bool{}
	if m.SourcePath != "" {
		baseDir = filepath.Dir(m.SourcePath)
		visited[m.SourcePath] = true
	}

	merged := &Emi{}
	for _, ext := range m.Extends {
		base, err := loadExtended(ext, baseDir, baseDir, visited)
		if err != nil {
			return err
		}
		mergeEmi(merged, base)
	}

	own := *m
	own.Extends = nil
	mergeEmi(merged, &own)

	merged.SourcePath = m.SourcePath
	*m = *merged
	return nil
}

// loadExtended reads the emi file at path (relative to baseDir), resolving its own
// extends chain, and returns it fully merged. chain holds the absolute paths on the
// current extends chain only, so cycles are caught without forbidding diamonds.
func loadExtended(ext EmiExtends, baseDir string, rootDir string, chain map[string]bool) (*Emi, error) {
	path := ext.From
	resolved := path
	if !filepath.IsAbs(resolved) && baseDir != "" {
		resolved = filepath.Join(baseDir, resolved)
	}
	abs, err := filepath.Abs(resolved)
	if err != nil {
		return nil, fmt.Errorf("extends: resolving %q: %w", path, err)
	}
	if chain[abs] {
		return nil, fmt.Errorf("extends: circular extends detected at %q", abs)
	}

	content, err := os.ReadFile(abs)
	if err != nil {
		return nil, fmt.Errorf("extends: reading %q: %w", path, err)
	}
	var base Emi
	if err := yaml.Unmarshal(content, &base); err != nil {
		return nil, fmt.Errorf("extends: parsing %q: %w", path, err)
	}
	applyParams(&base, withImplicitParams(ext.Params, abs, rootDir))
	base.SourcePath = abs

	next := make(map[string]bool, len(chain)+1)
	for k := range chain {
		next[k] = true
	}
	next[abs] = true

	dir := filepath.Dir(abs)
	merged := &Emi{}
	for _, ext := range base.Extends {
		parent, err := loadExtended(ext, dir, rootDir, next)
		if err != nil {
			return nil, err
		}
		mergeEmi(merged, parent)
	}
	base.Extends = nil

	// Complex `include:` entries are relative to the file that declares them, which stops
	// being true once merged into another module - expand them now.
	if base.Complexes, err = expandComplexes(base.Complexes, dir, nil); err != nil {
		return nil, err
	}

	mergeEmi(merged, &base)
	return merged, nil
}

// withImplicitParams returns params plus the placeholders emi provides itself (an
// explicit param of the same name wins):
//
//	{{extends.dir}} - the directory of the extended file, as a path relative to the
//	                  directory of the module at the root of the extends chain (the one
//	                  whose output ends up being generated). Because target outputs are
//	                  interpreted relative to that module, writing them as
//	                  "{{extends.dir}}/../../ui/sdk" in a shared file points at the same
//	                  place no matter how deep the extending module sits.
func withImplicitParams(params map[string]string, extendedFile, rootDir string) map[string]string {
	out := make(map[string]string, len(params)+1)
	dir := filepath.Dir(extendedFile)
	if rootDir != "" {
		if rel, err := filepath.Rel(rootDir, dir); err == nil {
			dir = rel
		}
	}
	out["extends.dir"] = filepath.ToSlash(dir)
	for k, v := range params {
		out[k] = v
	}
	return out
}

// mergeEmi applies src onto dst (src wins). See resolveExtends for the rules.
func mergeEmi(dst, src *Emi) {
	if src.Namespace != "" {
		dst.Namespace = src.Namespace
	}
	if src.Description != "" {
		dst.Description = src.Description
	}
	if src.Version != "" {
		dst.Version = src.Version
	}
	if src.Name != "" {
		dst.Name = src.Name
	}

	dst.Targets = mergeByKey(dst.Targets, src.Targets, func(x EmiCompile) string {
		// Compared as cleaned paths, so an output built from {{extends.dir}}/../.. matches
		// the same directory written directly in the extending module.
		return x.Compiler + "|" + filepath.Clean(x.Output)
	})
	dst.Enums = mergeByKey(dst.Enums, src.Enums, func(x EmiEnum) string { return x.Name })
	dst.Dto = mergeByKey(dst.Dto, src.Dto, func(x EmiDto) string { return x.Name })
	dst.Entities = mergeByKey(dst.Entities, src.Entities, func(x *Module3Entity) string { return x.Name })
	dst.Complexes = mergeByKey(dst.Complexes, src.Complexes, func(x EmiComplex) string {
		if x.Include != "" {
			return "include|" + x.Include
		}
		return x.Compiler + "|" + x.Namespace + "|" + x.Name
	})
	dst.Actions = mergeByKey(dst.Actions, src.Actions, func(x *EmiAction) string { return x.Name })
	dst.Remotes = mergeByKey(dst.Remotes, src.Remotes, func(x *EmiRemote) string { return x.Name })
	dst.Config = mergeByKey(dst.Config, src.Config, func(x EmiConfig) string { return x.Name })
	dst.Manifests = mergeByKey(dst.Manifests, src.Manifests, func(x EmiManifest) string { return x.Name })
	dst.Vsqls = mergeByKey(dst.Vsqls, src.Vsqls, func(x EmiVsql) string { return x.Name })
	dst.Permissions = mergeByKey(dst.Permissions, src.Permissions, func(x *EmiPermission) string { return x.Key })
	dst.Intents = mergeByKey(dst.Intents, src.Intents, func(x *EmiIntent) string { return x.Name })
	dst.Events = mergeByKey(dst.Events, src.Events, func(x *EmiEvent) string { return x.Key })

	if src.Templates != nil {
		if dst.Templates == nil {
			dst.Templates = &EmiTemplate{}
		}
		dst.Templates.Dtos = mergeByKey(dst.Templates.Dtos, src.Templates.Dtos, func(x EmiDto) string { return x.Name })
		dst.Templates.Actions = mergeByKey(dst.Templates.Actions, src.Templates.Actions, func(x *EmiAction) string { return x.Name })
		dst.Templates.Fields = mergeByKey(dst.Templates.Fields, src.Templates.Fields, func(x EmiFieldTemplate) string { return x.Name + "|" + x.Context })
	}
}

// mergeByKey returns dst with every src item applied: an item whose key already exists
// replaces that entry in place, any other is appended. Nil pointer items are kept as-is
// (validateNoNilEntries reports them later) and never matched.
func mergeByKey[T any](dst, src []T, key func(T) string) []T {
	index := make(map[string]int, len(dst))
	for i, item := range dst {
		if k, ok := safeKey(item, key); ok {
			index[k] = i
		}
	}
	for _, item := range src {
		k, ok := safeKey(item, key)
		if !ok {
			dst = append(dst, item)
			continue
		}
		if i, exists := index[k]; exists {
			dst[i] = item
			continue
		}
		index[k] = len(dst)
		dst = append(dst, item)
	}
	return dst
}

// safeKey computes an item's identity, treating nil pointers and empty keys as
// unmatchable rather than panicking or collapsing all unnamed entries into one.
func safeKey[T any](item T, key func(T) string) (k string, ok bool) {
	defer func() {
		if recover() != nil {
			k, ok = "", false
		}
	}()
	k = key(item)
	return k, k != "" && k != "|" && k != "||"
}
