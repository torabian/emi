package dart

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	dartInclude "github.com/torabian/emi/lib/dart/dart-include"
)

// GetDartPublicActions exposes every dart-target action (text and file
// based) - class/header generation, full module compilation, and the sdk
// (runtime-only) dump - the same shape lib/js, lib/kotlin, lib/swift and
// lib/python expose, so it plugs into cmd/emi-wasm and lib/gorunner the same
// way.
func GetDartPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "dart:dto",
				Description:      "Generates a dart class out of a dto signature (name, fields)",
				WasmFunctionName: "dartGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := DartCommonObjectGenerator(emiDto.Fields, ctx, DartCommonObjectContext{RootClassName: emiDto.GetClassName()})
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "dart:headers",
				Description:      "Generates a dart class out of typed headers, usable for both request and response",
				WasmFunctionName: "dartGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := DartHeaderClass("Headers", headers, ctx)
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "dart:action",
				Description:      "Creates a complete dart client function for a single action, including path/query/request/response shapes",
				WasmFunctionName: "dartGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := DartActionRender(&action, ctx, nil)
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
		DartPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "dart:sdk",
				Description:      "Writes only the dart runtime (fetchx request/streaming helpers) to disk",
				WasmFunctionName: "dartGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&dartInclude.Content, "lib"), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "dart:module",
				Description:      "Compiles the entire dart module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "dartGenModule",
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
				return DartModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// DartPrimaryAction detects the emi content type (dto or module) the same
// way the sibling generators' primary actions do, and dispatches to the
// right sub-compiler.
var DartPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "dart",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "dartGen",
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
			return DartModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			chunk, err := DartCommonObjectGenerator(emiDto.Fields, ctx, DartCommonObjectContext{RootClassName: emiDto.GetClassName()})
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
