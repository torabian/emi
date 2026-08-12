package java

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	javaInclude "github.com/torabian/emi/lib/java/java-include"
)

// GetJavaPublicActions exposes every Java-target action (text and file
// based) - class/header generation, full module compilation, and the sdk
// (runtime-only) dump - the same shape the sibling generators expose, so it
// plugs into cmd/emi-wasm and lib/gorunner the same way.
//
// Unlike lib/python/lib/dart/lib/csharp, `java:dto` and `java:action` are
// FileActions rather than TextActions: Java's one-public-type-per-file rule
// means a single dto (or action) can legitimately expand into several real
// files (flattened nested objects, inline enums, path/query/header
// classes...) - concatenating them into one string wouldn't even compile.
func GetJavaPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "java:headers",
				Description:      "Generates a Java class out of typed headers, usable for both request and response",
				WasmFunctionName: "javaGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := JavaHeaderClass("Headers", headers, ctx)
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
	}

	fileActions := []core.ActionFile{
		JavaPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "java:dto",
				Description:      "Generates the Java class(es) out of a dto signature (name, fields) - may be more than one file (flattened nested objects/inline enums)",
				WasmFunctionName: "javaGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return nil, err
				}
				chunks, err := JavaCommonObjectGenerator(emiDto.Fields, ctx, JavaCommonObjectContext{RootClassName: emiDto.GetClassName()})
				if err != nil {
					return nil, err
				}
				return chunksToVirtualFiles(chunks, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "java:action",
				Description:      "Creates the complete Java client file(s) for a single action, including path/query/request/response shapes",
				WasmFunctionName: "javaGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return nil, err
				}
				chunks, err := JavaActionRender(&action, ctx, nil)
				if err != nil {
					return nil, err
				}
				if len(chunks) == 0 {
					return nil, errors.New("action does not have a name")
				}
				return chunksToVirtualFiles(chunks, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "java:sdk",
				Description:      "Writes only the Java runtime (Fetchx request/streaming helpers) to disk",
				WasmFunctionName: "javaGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&javaInclude.Content, "src/main/java/emisdk"), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "java:module",
				Description:      "Compiles the entire Java module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "javaGenModule",
				Flags: []core.FlagDef{
					{Name: "actions", Usage: "Actions to be generated - separate the action names using comma (,)", Type: core.FlagString},
					{Name: "remotes", Usage: "Remotes to be generated - separate the remote names using comma (,)", Type: core.FlagString},
					{Name: "dtos", Usage: "Dtos to be generated - separate the dto names using comma (,)", Type: core.FlagString},
				},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				emiModule, err := core.StringToEmi(ctx.Content)
				if err != nil {
					return nil, err
				}
				return JavaModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// JavaPrimaryAction detects the emi content type (dto or module) the same
// way the sibling generators' primary actions do, and dispatches to the
// right sub-compiler.
var JavaPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "java",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "javaGen",
		Flags:            []core.FlagDef{},
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}

		if type_ == "module" {
			emiModule, err := core.StringToEmi(ctx.Content)
			if err != nil {
				return nil, err
			}
			return JavaModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			chunks, err := JavaCommonObjectGenerator(emiDto.Fields, ctx, JavaCommonObjectContext{RootClassName: emiDto.GetClassName()})
			if err != nil {
				return nil, err
			}
			return chunksToVirtualFiles(chunks, ""), nil
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}
