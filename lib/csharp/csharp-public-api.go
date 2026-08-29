package csharp

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	csharpInclude "github.com/torabian/emi/lib/csharp/csharp-include"
)

// GetCSharpPublicActions exposes every C#-target action (text and file
// based) - class/header generation, full module compilation, and the sdk
// (runtime-only) dump - the same shape the sibling generators expose, so it
// plugs into cmd/emi-wasm and lib/gorunner the same way.
func GetCSharpPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "csharp:dto",
				Description:      "Generates a C# class out of a dto signature (name, fields)",
				WasmFunctionName: "csharpGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CSharpCommonObjectGenerator(emiDto.Fields, ctx, CSharpCommonObjectContext{RootClassName: emiDto.GetClassName()})
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "csharp:headers",
				Description:      "Generates a C# class out of typed headers, usable for both request and response",
				WasmFunctionName: "csharpGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CSharpHeaderClass("Headers", headers, ctx)
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "csharp:action",
				Description:      "Creates a complete C# client method for a single action, including path/query/request/response shapes",
				WasmFunctionName: "csharpGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CSharpActionRender(&action, ctx, nil)
				if err != nil {
					return "", err
				}
				if chunk == nil {
					return "", errors.New("action does not have a name")
				}
				return AsFullDocument(chunk), nil
			},
		},
	}

	fileActions := []core.ActionFile{
		CSharpPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "csharp:sdk",
				Description:      "Writes only the C# runtime (Fetchx request/streaming helpers) to disk",
				WasmFunctionName: "csharpGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&csharpInclude.Content, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "csharp:module",
				Description:      "Compiles the entire C# module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "csharpGenModule",
				Flags: []core.FlagDef{
					{Name: "actions", Usage: "Actions to be generated - separate the action names using comma (,)", Type: core.FlagString},
					{Name: "remotes", Usage: "Remotes to be generated - separate the remote names using comma (,)", Type: core.FlagString},
					{Name: "dtos", Usage: "Dtos to be generated - separate the dto names using comma (,)", Type: core.FlagString},
				},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				emiModule, err := core.StringToEmiWithPath(ctx.Content, ctx.Path)
				if err != nil {
					return nil, err
				}
				return CSharpModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// CSharpPrimaryAction detects the emi content type (dto or module) the same
// way the sibling generators' primary actions do, and dispatches to the
// right sub-compiler.
var CSharpPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "csharp",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "csharpGen",
		Flags:            []core.FlagDef{},
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}

		if type_ == "module" {
			emiModule, err := core.StringToEmiWithPath(ctx.Content, ctx.Path)
			if err != nil {
				return nil, err
			}
			return CSharpModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			chunk, err := CSharpCommonObjectGenerator(emiDto.Fields, ctx, CSharpCommonObjectContext{RootClassName: emiDto.GetClassName()})
			if err != nil {
				return nil, err
			}
			return []core.VirtualFile{{
				Name:         chunk.SuggestedFileName,
				Extension:    chunk.SuggestedExtension,
				ActualScript: AsFullDocument(chunk),
			}}, nil
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}
