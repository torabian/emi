package php

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	phpInclude "github.com/torabian/emi/lib/php/php-include"
)

// GetPhpPublicActions exposes every PHP-target action (text and file based)
// - class/header generation, full module compilation, and the sdk
// (runtime-only) dump - the same shape the sibling generators expose, so it
// plugs into cmd/emi-wasm and lib/gorunner the same way.
func GetPhpPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "php:dto",
				Description:      "Generates a PHP class out of a dto signature (name, fields)",
				WasmFunctionName: "phpGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := PhpCommonObjectGenerator(emiDto.Fields, ctx, PhpCommonObjectContext{RootClassName: emiDto.GetClassName()})
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "php:headers",
				Description:      "Generates a PHP class out of typed headers, usable for both request and response",
				WasmFunctionName: "phpGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := PhpHeaderClass("Headers", headers, ctx)
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "php:action",
				Description:      "Creates a complete PHP client function for a single action, including path/query/request/response shapes",
				WasmFunctionName: "phpGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := PhpActionRender(&action, ctx, nil)
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
		PhpPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "php:sdk",
				Description:      "Writes only the PHP runtime (Fetchx request/streaming helpers + Hydrator reflection-based serialization) to disk",
				WasmFunctionName: "phpGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&phpInclude.Content, "src"), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "php:module",
				Description:      "Compiles the entire PHP module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "phpGenModule",
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
				return PhpModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// PhpPrimaryAction detects the emi content type (dto or module) the same
// way the sibling generators' primary actions do, and dispatches to the
// right sub-compiler.
var PhpPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "php",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "phpGen",
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
			return PhpModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			chunk, err := PhpCommonObjectGenerator(emiDto.Fields, ctx, PhpCommonObjectContext{RootClassName: emiDto.GetClassName()})
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
