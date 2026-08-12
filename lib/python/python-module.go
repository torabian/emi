package python

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
	pythonInclude "github.com/torabian/emi/lib/python/python-include"
)

// AsFullDocument wraps a compiled chunk into a standalone, importable python
// module: `from __future__ import annotations` (so forward references between
// flattened nested classes always resolve, regardless of declaration order),
// followed by the deduplicated imports, followed by the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	imports := CombinePythonImports(*x)

	var sb strings.Builder
	sb.WriteString("from __future__ import annotations\n\n")
	if imports != "" {
		sb.WriteString(imports)
		sb.WriteString("\n\n")
	}
	sb.Write(x.ActualScript)

	return string(core.EscapeLines([]byte(sb.String())))
}

// CombinePythonImports groups every CodeChunkDependency by its Location,
// merges/dedupes/sorts the imported Objects, and renders one `from X import
// A, B` statement per location. A Location starting with "." is treated as
// package-relative (generated dto/action/runtime files); anything else is a
// plain absolute import (stdlib or third-party, e.g. "typing", "httpx").
func CombinePythonImports(chunk core.CodeChunkCompiled) string {
	objectsByLocation := map[string]map[string]struct{}{}
	plainLocations := map[string]struct{}{}

	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		if len(dep.Objects) == 0 {
			plainLocations[dep.Location] = struct{}{}
			continue
		}
		if objectsByLocation[dep.Location] == nil {
			objectsByLocation[dep.Location] = map[string]struct{}{}
		}
		for _, obj := range dep.Objects {
			objectsByLocation[dep.Location][obj] = struct{}{}
		}
	}

	locations := make([]string, 0, len(objectsByLocation))
	for loc := range objectsByLocation {
		locations = append(locations, loc)
	}
	sort.Strings(locations)

	lines := []string{}
	for _, loc := range locations {
		objSet := objectsByLocation[loc]
		objs := make([]string, 0, len(objSet))
		for o := range objSet {
			objs = append(objs, o)
		}
		sort.Strings(objs)
		lines = append(lines, fmt.Sprintf("from %v import %v", loc, strings.Join(objs, ", ")))
	}

	plain := make([]string, 0, len(plainLocations))
	for loc := range plainLocations {
		plain = append(plain, loc)
	}
	sort.Strings(plain)
	for _, loc := range plain {
		lines = append(lines, fmt.Sprintf("import %v", loc))
	}

	return strings.Join(lines, "\n")
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/python target, mirroring lib/js/js-module.go#DiscoverComplexes.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "python" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type PythonModuleGenerationFlags struct {
	Actions  *string
	Remotes  *string
	Dtos     *string
	Entities *string
}

func (x PythonModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readPyModuleFlags(ctx core.MicroGenContext) PythonModuleGenerationFlags {
	config := PythonModuleGenerationFlags{}
	if v, ok := ctx.Flags["actions"]; ok && v != "" {
		config.Actions = &v
	}
	if v, ok := ctx.Flags["remotes"]; ok && v != "" {
		config.Remotes = &v
	}
	if v, ok := ctx.Flags["dtos"]; ok && v != "" {
		config.Dtos = &v
	}
	if v, ok := ctx.Flags["entities"]; ok && v != "" {
		config.Entities = &v
	}
	return config
}

// PythonModuleFullVirtualFiles compiles an entire Emi module into a python
// package: one file per dto/enum/action/remote, plus the embedded runtime
// (fetchx/serialization) and a requirements.txt, mirroring
// lib/js/js-module.go#JsModuleFullVirtualFiles and
// lib/kotlin/kotlin-public-api.go#KotlinModuleFull.
func PythonModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readPyModuleFlags(ctx)
	files := []core.VirtualFile{}
	exportedSymbols := []string{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}

		chunk, err := PythonCommonObjectGenerator(dto.Fields, ctx, PyCommonObjectContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
		exportedSymbols = append(exportedSymbols, dto.GetClassName())
	}

	for _, enum := range module.Enums {
		chunk, err := PythonStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
		exportedSymbols = append(exportedSymbols, enum.GetName())
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}

		chunk, err := PythonActionRender(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}

		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, remote := range module.Remotes {
		if len(remoteFilter) > 0 && !slices.Contains(remoteFilter, remote.Name) {
			continue
		}

		chunk, err := PythonActionRender(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}

		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	skipSdk := ctx.HasTag(NoSdk)
	if !skipSdk {
		files = append(files, core.FsEmbedToVirtualFile(&pythonInclude.Content, "")...)
	}

	files = append(files, core.VirtualFile{
		Name:         "__init__",
		Extension:    ".py",
		ActualScript: "\"\"\"Generated by the Emi compiler - do not edit by hand.\"\"\"\n",
	})

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "requirements",
			Extension:    ".txt",
			ActualScript: "httpx>=0.27\n",
		})
	}

	return files, nil
}
