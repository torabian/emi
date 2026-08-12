package c

import (
	"fmt"
	"slices"
	"strings"

	cInclude "github.com/torabian/emi/lib/c/c-include"
	"github.com/torabian/emi/lib/core"
)

// commonIncludes is emitted at the top of every generated header. An
// unused #include is never a compiler error in C, so a single fixed set -
// covering every standard header any of this package's templates might
// need - is simpler and just as correct as computing a precise one per
// file.
var commonIncludes = []string{
	`"cjson/cJSON.h"`,
	"<stdbool.h>",
	"<stddef.h>",
	"<stdint.h>",
	"<stdio.h>",
	"<stdlib.h>",
	"<string.h>",
}

// AsFullDocument wraps a compiled chunk into a standalone, header-guarded
// .h file: the guard, the common includes, plus any module-specific
// #include (a sibling generated header, or a resolved complex type's own
// header), then the actual body.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	guard := "EMISDK_" + strings.ToUpper(core.ToSnakeCase(x.SuggestedFileName)) + "_H"

	includes := make([]string, len(commonIncludes))
	copy(includes, commonIncludes)
	seen := map[string]struct{}{}
	for _, i := range includes {
		seen[i] = struct{}{}
	}
	for _, dep := range x.CodeChunkDependensies {
		if dep.Location == "" {
			continue
		}
		spelling := dep.Location
		if !strings.HasPrefix(spelling, "<") && !strings.HasPrefix(spelling, `"`) {
			spelling = `"` + spelling + `"`
		}
		if _, ok := seen[spelling]; ok {
			continue
		}
		seen[spelling] = struct{}{}
		includes = append(includes, spelling)
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "#ifndef %s\n#define %s\n\n", guard, guard)
	for _, i := range includes {
		fmt.Fprintf(&sb, "#include %s\n", i)
	}
	sb.WriteString("\n")
	sb.Write(x.ActualScript)
	fmt.Fprintf(&sb, "\n#endif /* %s */\n", guard)

	return string(core.EscapeLines([]byte(sb.String())))
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/C target, mirroring the same helper in the sibling generators.
// ImportLocation here is a header to #include (`"vector3.h"` or
// `<mymath/vector3.h>`), not a generated file path.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "c" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type CModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x CModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readCModuleFlags(ctx core.MicroGenContext) CModuleGenerationFlags {
	config := CModuleGenerationFlags{}
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

// CModuleFullVirtualFiles compiles an entire Emi module into a C header
// library: one header per dto/enum/action/remote, plus the embedded
// runtime (vendored cJSON + fetchx.h) and a short README on how to build,
// mirroring the equivalent full-module compilers in the sibling
// generators.
func CModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readCModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunk, err := CCommonObjectGenerator(dto.Fields, ctx, CCommonObjectContext{
			RootTypeName:        dto.GetClassName(),
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
		chunk, err := CStandaloneEnum(enum, ctx)
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
		chunk, err := CActionRender(action, ctx, complexes)
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
		chunk, err := CActionRender(remote, ctx, complexes)
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
		files = append(files, core.FsEmbedToVirtualFile(&cInclude.Content, "")...)
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "README",
			Extension:    ".md",
			ActualScript: cReadme,
		})
	}

	return files, nil
}

const cReadme = `# Generated Emi C SDK

Generated by the Emi compiler - do not edit by hand.

Every ` + "`.h`" + ` file here is header-only (` + "`static inline`" + ` throughout, like the vendored
` + "`cjson/cJSON.h`" + `) - just ` + "`#include`" + ` whichever ones you need, from as many ` + "`.c`" + ` files as you
like, and compile ` + "`cjson/cJSON.c`" + ` into your program once:

` + "```" + `sh
cc -I. your_program.c cjson/cJSON.c -lcurl -o your_program
` + "```" + `

Every generated struct ` + "`Type`" + ` gets: ` + "`Type_new()`" + `/` + "`Type_free()`" + ` (heap
alloc/dealloc, always safe to call), ` + "`Type_from_json(const cJSON*)`" + ` /
` + "`Type_to_json(const Type*)`" + ` for (de)serialization, and - since every field
always gets a safe default - casting one dto's data into a different (but
JSON-compatible) struct is just going through cJSON as the interchange
format:

` + "```" + `c
cJSON* json = SourceDto_to_json(source);
TargetDto* target = TargetDto_from_json(json);
cJSON_Delete(json);
/* ... use target ... */
TargetDto_free(target);
` + "```" + `

Action functions (` + "`create_widget_action(...)`" + ` etc.) return ` + "`NULL`" + ` on any
failure (transport-level or a non-2xx status) - pass a ` + "`FetchxResponse*`" + ` as
their last argument to inspect ` + "`.status_code`" + `/` + "`.error_message`" + ` on failure (and
` + "`FetchxResponse_free()`" + ` it when you're done either way).
`
