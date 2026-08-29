package core

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func init() {
	// Registered globally (not per-action) so every compiler backend - go, kotlin,
	// swift, js, openapi, postman, md, ... - sees fully-expanded complexes without
	// needing to know EmiComplex.Include exists.
	RegisterPreprocessHook(resolveComplexIncludes)
}

// resolveComplexIncludes expands every EmiComplex entry in m.Complexes that sets
// Include into the (recursively resolved) complexes declared by the file it points
// to, in place. Entries that don't set Include are left untouched.
func resolveComplexIncludes(m *Emi) error {
	if m == nil || len(m.Complexes) == 0 {
		return nil
	}

	// baseDir anchors the module's own top-level Include paths - i.e. the directory
	// its own yaml file lives in, when known (see Emi.SourcePath). When the module
	// was parsed from a raw string rather than a real file, this falls back to the
	// process's current directory (filepath.Abs's own fallback for a relative path).
	baseDir := ""
	if m.SourcePath != "" {
		baseDir = filepath.Dir(m.SourcePath)
	}

	resolved, err := expandComplexes(m.Complexes, baseDir, nil)
	if err != nil {
		return err
	}
	m.Complexes = resolved
	return nil
}

// expandComplexes walks complexes in declaration order, replacing every entry that
// sets Include with the (recursively expanded) complexes of the file it points to.
// Relative Include paths are resolved against baseDir. visited holds the absolute
// paths already being expanded along the current include chain, used to reject a
// circular include instead of recursing forever.
func expandComplexes(complexes []EmiComplex, baseDir string, visited map[string]bool) ([]EmiComplex, error) {
	out := make([]EmiComplex, 0, len(complexes))
	for _, c := range complexes {
		if c.Include == "" {
			out = append(out, c)
			continue
		}
		included, err := loadIncludedComplexes(c.Include, baseDir, visited)
		if err != nil {
			return nil, err
		}
		out = append(out, included...)
	}
	return out, nil
}

// loadIncludedComplexes reads the `complexes` list out of the emi yaml file at
// includePath (resolved against baseDir if relative), then recursively expands any
// Include entries found there - relative to the included file's own directory, so an
// included file's includes work regardless of where it was pulled in from.
func loadIncludedComplexes(includePath string, baseDir string, visited map[string]bool) ([]EmiComplex, error) {
	resolved := includePath
	if !filepath.IsAbs(resolved) && baseDir != "" {
		resolved = filepath.Join(baseDir, resolved)
	}
	abs, err := filepath.Abs(resolved)
	if err != nil {
		return nil, fmt.Errorf("complexes: resolving include %q: %w", includePath, err)
	}

	if visited[abs] {
		return nil, fmt.Errorf("complexes: circular include detected at %q", abs)
	}
	nextVisited := make(map[string]bool, len(visited)+1)
	for k := range visited {
		nextVisited[k] = true
	}
	nextVisited[abs] = true

	content, err := os.ReadFile(abs)
	if err != nil {
		return nil, fmt.Errorf("complexes: reading include %q: %w", includePath, err)
	}

	var included struct {
		Complexes []EmiComplex `yaml:"complexes,omitempty"`
	}
	if err := yaml.Unmarshal(content, &included); err != nil {
		return nil, fmt.Errorf("complexes: parsing include %q: %w", includePath, err)
	}

	return expandComplexes(included.Complexes, filepath.Dir(abs), nextVisited)
}
