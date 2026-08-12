package php

import (
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"
	phpInclude "github.com/torabian/emi/lib/php/php-include"
)

// AsFullDocument wraps a compiled chunk into a standalone .php file: the
// open tag, strict_types declaration, the `EmiSdk` namespace, then the
// actual body. Every cross-reference (sibling dto/action classes, the
// EmiSdk\Runtime helpers) is written fully-qualified inline in the
// generated code itself (see python-module.go's/dart's import-combining for
// contrast) - PHP needs no `use` statements for any of it, since dto/action
// classes all share the one `EmiSdk` namespace and the runtime is always
// referenced with its own leading `\`.
func AsFullDocument(x *core.CodeChunkCompiled) string {
	var sb strings.Builder
	sb.WriteString("<?php\n\ndeclare(strict_types=1);\n\nnamespace EmiSdk;\n\n")
	sb.Write(x.ActualScript)

	return string(core.EscapeLines([]byte(sb.String())))
}

// DiscoverComplexes finds every module-level complex type compiled for the
// general/php target, mirroring the same helper in the sibling generators.
// ImportLocation here is a fully-qualified PHP class name to reference
// inline (no `use` statement wiring is generated for it - see
// AsFullDocument), not a file path.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "php" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}
	return items
}

type PhpModuleGenerationFlags struct {
	Actions *string
	Remotes *string
	Dtos    *string
}

func (x PhpModuleGenerationFlags) filterList(raw *string) []string {
	if raw == nil || *raw == "" {
		return nil
	}
	return strings.Split(*raw, ",")
}

func readPhpModuleFlags(ctx core.MicroGenContext) PhpModuleGenerationFlags {
	config := PhpModuleGenerationFlags{}
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

// PhpModuleFullVirtualFiles compiles an entire Emi module into a PHP
// package: one file per dto/enum/action/remote (under src/, matching the
// `EmiSdk` PSR-4 namespace root), plus the embedded runtime
// (EmiSdk\Runtime\Fetchx/Hydrator, under src/Runtime/) and a composer.json,
// mirroring the equivalent full-module compilers in the sibling generators.
func PhpModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	complexes := DiscoverComplexes(module)
	config := readPhpModuleFlags(ctx)
	files := []core.VirtualFile{}

	dtoFilter := config.filterList(config.Dtos)
	actionFilter := config.filterList(config.Actions)
	remoteFilter := config.filterList(config.Remotes)

	const srcDir = "src"

	for _, dto := range module.Dto {
		if len(dtoFilter) > 0 && !slices.Contains(dtoFilter, dto.Name) {
			continue
		}
		chunk, err := PhpCommonObjectGenerator(dto.Fields, ctx, PhpCommonObjectContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, enum := range module.Enums {
		chunk, err := PhpStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, action := range module.Actions {
		if len(actionFilter) > 0 && !slices.Contains(actionFilter, action.Name) {
			continue
		}
		chunk, err := PhpActionRender(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	for _, remote := range module.Remotes {
		if len(remoteFilter) > 0 && !slices.Contains(remoteFilter, remote.Name) {
			continue
		}
		chunk, err := PhpActionRender(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		if chunk == nil {
			continue
		}
		files = append(files, core.VirtualFile{
			Name:         chunk.SuggestedFileName,
			Location:     srcDir,
			Extension:    chunk.SuggestedExtension,
			ActualScript: AsFullDocument(chunk),
		})
	}

	skipSdk := ctx.HasTag(NoSdk)
	if !skipSdk {
		files = append(files, core.FsEmbedToVirtualFile(&phpInclude.Content, srcDir)...)
	}

	skipPackage := ctx.HasTag(NoPackage)
	if !skipPackage {
		files = append(files, core.VirtualFile{
			Name:         "composer",
			Extension:    ".json",
			ActualScript: phpComposerJSON,
		})
	}

	return files, nil
}

const phpComposerJSON = `{
    "name": "emisdk/emi-sdk",
    "description": "Generated by the Emi compiler - do not edit by hand.",
    "type": "library",
    "require": {
        "php": ">=8.1",
        "ext-curl": "*",
        "ext-json": "*"
    },
    "autoload": {
        "psr-4": {
            "EmiSdk\\": "src/"
        }
    }
}
`
