package python

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	pythonInclude "github.com/torabian/emi/lib/python/python-include"
)

// GetPythonPublicActions exposes every python-target action (text and file
// based) - dto/header class generation, full module compilation, and the sdk
// (runtime-only) dump - the exact same shape lib/js, lib/kotlin and
// lib/swift expose, so it plugs into cmd/emi-wasm and lib/gorunner the same
// way.
func GetPythonPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "python:dto",
				Description:      "Generates a python dataclass out of a dto signature (name, fields)",
				WasmFunctionName: "pythonGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}

				chunk, err := PythonCommonObjectGenerator(emiDto.Fields, ctx, PyCommonObjectContext{
					RootClassName: emiDto.GetClassName(),
				})
				if err != nil {
					return "", err
				}

				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "python:headers",
				Description:      "Generates a python dataclass out of typed headers, usable for both request and response",
				WasmFunctionName: "pythonGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}

				chunk, err := PythonHeaderClass("Headers", headers, ctx)
				if err != nil {
					return "", err
				}

				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "python:action",
				Description:      "Creates a complete python client function for a single action, including path/query/request/response shapes",
				WasmFunctionName: "pythonGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}

				chunk, err := PythonActionRender(&action, ctx, nil)
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
		PythonPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "python:sdk",
				Description:      "Writes only the python runtime (fetchx request/streaming helpers + generic dataclass serialization) to disk",
				WasmFunctionName: "pythonGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&pythonInclude.Content, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "python:module",
				Description:      "Compiles the entire python module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "pythonGenModule",
				Flags: []core.FlagDef{
					{
						Name:  "actions",
						Usage: "Actions to be generated - separate the action names using comma (,)",
						Type:  core.FlagString,
					},
					{
						Name:  "remotes",
						Usage: "Remotes to be generated - separate the remote names using comma (,)",
						Type:  core.FlagString,
					},
					{
						Name:  "dtos",
						Usage: "Dtos to be generated - separate the dto names using comma (,)",
						Type:  core.FlagString,
					},
				},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				emiModule, err := core.StringToEmiWithPath(ctx.Content, ctx.Path)
				if err != nil {
					return nil, err
				}

				return PythonModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

// PythonPrimaryAction detects the emi content type (dto or module) the same
// way lib/js/lib/kotlin's primary action does, and dispatches to the right
// sub-compiler.
var PythonPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "python",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "pythonGen",
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

			return PythonModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}

			chunk, err := PythonCommonObjectGenerator(emiDto.Fields, ctx, PyCommonObjectContext{
				RootClassName: emiDto.GetClassName(),
			})
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
