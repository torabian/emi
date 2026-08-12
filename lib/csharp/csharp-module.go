package csharp

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
	csharpInclude "github.com/torabian/emi/lib/csharp/csharp-include"
)

// csharpNamespace is the single namespace every generated type lives in.
// Since every generated file compiles into the same project/assembly, C#
// needs no per-file import wiring between them at all (unlike lib/dart or
// lib/python) - any type is reachable from any other file for free.
const csharpNamespace = "EmiSdk"

// commonUsings is emitted at the top of every generated file. Unlike Python
// or Dart, an unused `using` in C# is (at most) a mild analyzer suggestion,
// never a compile error, so a single fixed set is simpler and just as
// correct as computing a precise one per file.
var commonUsings = []string{
	"System",
	"System.Collections.Generic",
	"System.Net.Http",
	"System.Text",
	"System.Text.Json",
	"System.Text.Json.Serialization",
	"System.Threading",
	"System.Threading.Tasks",
}

// AsFullDocument wraps a compiled chunk into a standalone .cs file: the
// common usings, the file-scoped namespace declaration, plus any
// module-specific `using` (e.g. a resolved complex type's namespace), then
// the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	usings := make([]string, len(commonUsings))
	copy(usings, commonUsings)

	seen := map[string]struct{}{}
	for _, u := range usings {
		seen[u] = struct{}{}
	}
	for _, dep := range x.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		if _, ok := seen[dep.Location]; ok {
			continue
		}
		seen[dep.Location] = struct{}{}
		usings = append(usings, dep.Location)
	}
	sort.Strings(usings)

	var sb strings.Builder
	for _, u := range usings {
		fmt.Fprintf(&sb, "using %v;\n", u)
	}
	fmt.Fprintf(&sb, "\nnamespace %v;\n\n", csharpNamespace)
	sb.Write(x.ActualScript)

	return string(core.EscapeLines([]byte(sb.String())))
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/csharp target, mirroring the same helper in the sibling
// generators. ImportLocation here is a C# namespace (`using X;`), not a
// file path.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "csharp" || complex.Compiler == "cs" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type CSharpModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x CSharpModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readCSharpModuleFlags(ctx core.MicroGenContext) CSharpModuleGenerationFlags {
	config := CSharpModuleGenerationFlags{}
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

// CSharpModuleFullVirtualFiles compiles an entire Emi module into a C#
// project: one file per dto/enum/action/remote, plus the embedded runtime
// (Fetchx) and a .csproj, mirroring the equivalent full-module compilers in
// the sibling generators.
func CSharpModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readCSharpModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunk, err := CSharpCommonObjectGenerator(dto.Fields, ctx, CSharpCommonObjectContext{
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
	}

	for _, enum := range module.Enums {
		chunk, err := CSharpStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}
		chunk, err := CSharpActionRender(action, ctx, complexes)
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
		chunk, err := CSharpActionRender(remote, ctx, complexes)
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
		files = append(files, core.FsEmbedToVirtualFile(&csharpInclude.Content, "")...)
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "EmiSdk",
			Extension:    ".csproj",
			ActualScript: csharpCsproj,
		})
	}

	return files, nil
}

const csharpCsproj = `<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <!-- Generated by the Emi compiler - do not edit by hand. -->
    <TargetFramework>net8.0</TargetFramework>
    <Nullable>enable</Nullable>
    <ImplicitUsings>disable</ImplicitUsings>
    <RootNamespace>EmiSdk</RootNamespace>
  </PropertyGroup>

</Project>
`
