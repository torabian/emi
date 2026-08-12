package dart

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
	dartInclude "github.com/torabian/emi/lib/dart/dart-include"
)

// AsFullDocument wraps a compiled chunk into a standalone dart file: the
// deduplicated imports, followed by the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	imports := CombineDartImports(*x)

	var sb strings.Builder
	if imports != "" {
		sb.WriteString(imports)
		sb.WriteString("\n\n")
	}
	sb.Write(x.ActualScript)

	return string(core.EscapeLines([]byte(sb.String())))
}

// CombineDartImports collects every CodeChunkDependency's Location (already a
// full relative-import path, e.g. "user_dto.dart" or "runtime/fetchx.dart"),
// dedupes and sorts them, and renders one `import 'x';` statement per
// location.
func CombineDartImports(chunk core.CodeChunkCompiled) string {
	seen := map[string]struct{}{}
	for _, dep := range chunk.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		seen[dep.Location] = struct{}{}
	}

	locations := make([]string, 0, len(seen))
	for loc := range seen {
		locations = append(locations, loc)
	}
	sort.Strings(locations)

	lines := make([]string, 0, len(locations))
	for _, loc := range locations {
		lines = append(lines, fmt.Sprintf("import '%v';", loc))
	}
	return strings.Join(lines, "\n")
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/dart target, mirroring the same helper in the sibling generators.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "dart" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type DartModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x DartModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readDartModuleFlags(ctx core.MicroGenContext) DartModuleGenerationFlags {
	config := DartModuleGenerationFlags{}
	if v, ok := ctx.Flags["actions"]; ok && v != "" {
		config.Actions = &v
	}
	if v, ok := ctx.Flags["remotes"]; ok && v != "" {
		config.Remotes = &v
	}
	if v, ok := ctx.Flags["dtos"]; ok && v != "" {
		config.Dtos = &v
	}
	return config
}

// DartModuleFullVirtualFiles compiles an entire Emi module into a dart
// package: one file per dto/enum/action/remote, plus the embedded runtime
// (fetchx) and a pubspec.yaml, mirroring the equivalent full-module
// compilers in the sibling generators.
func DartModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readDartModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunk, err := DartCommonObjectGenerator(dto.Fields, ctx, DartCommonObjectContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     "lib",
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, enum := range module.Enums {
		chunk, err := DartStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     "lib",
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}
		chunk, err := DartActionRender(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     "lib",
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, remote := range module.Remotes {
		if len(remoteFilter) > 0 && !slices.Contains(remoteFilter, remote.Name) {
			continue
		}
		chunk, err := DartActionRender(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     "lib",
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	skipSdk := ctx.HasTag(NoSdk)
	if !skipSdk {
		files = append(files, core.FsEmbedToVirtualFile(&dartInclude.Content, "lib")...)
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "pubspec",
			Extension:    ".yaml",
			ActualScript: dartPubspec,
		})
	}

	return files, nil
}

const dartPubspec = `name: emi_sdk
description: Generated by the Emi compiler - do not edit by hand.
environment:
  sdk: '>=3.0.0 <4.0.0'
dependencies:
  http: ^1.2.0
`
