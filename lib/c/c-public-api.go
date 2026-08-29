package c

import (
	"errors"

	cInclude "github.com/torabian/emi/lib/c/c-include"
	"github.com/torabian/emi/lib/core"
)

// GetCPublicActions exposes every C-target action (text and file based) -
// struct/header generation, full module compilation, and the sdk
// (runtime-only) dump - the same shape the sibling generators expose, so it
// plugs into cmd/emi-wasm and lib/gorunner the same way.
func GetCPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "c:dto",
				Description:      "Generates a C struct (+ cJSON alloc/free/to_json/from_json) out of a dto signature (name, fields)",
				WasmFunctionName: "cGenDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CCommonObjectGenerator(emiDto.Fields, ctx, CCommonObjectContext{RootTypeName: emiDto.GetClassName()})
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "c:headers",
				Description:      "Generates a C struct out of typed headers, with a helper to serialize it onto a curl_slist",
				WasmFunctionName: "cGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CHeaderStruct("Headers", headers, ctx)
				if err != nil {
					return "", err
				}
				return AsFullDocument(chunk), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "c:action",
				Description:      "Creates a complete C client function for a single action, including path/query/request/response shapes",
				WasmFunctionName: "cGenAction",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}
				chunk, err := CActionRender(&action, ctx, nil)
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
		CPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "c:sdk",
				Description:      "Writes only the C runtime (vendored cJSON + fetchx.h request/streaming helpers) to disk",
				WasmFunctionName: "cGenSdk",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.FsEmbedToVirtualFile(&cInclude.Content, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "c:module",
				Description:      "Compiles the entire C module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "cGenModule",
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
				return CModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// CPrimaryAction detects the emi content type (dto or module) the same way
// the sibling generators' primary actions do, and dispatches to the right
// sub-compiler.
var CPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "c",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "cGen",
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
			return CModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			chunk, err := CCommonObjectGenerator(emiDto.Fields, ctx, CCommonObjectContext{RootTypeName: emiDto.GetClassName()})
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
