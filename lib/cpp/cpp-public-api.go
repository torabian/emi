package cpp

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	cppIncludeGeneric "github.com/torabian/emi/lib/cpp/cpp-include/generic"
	cppIncludeUnreal "github.com/torabian/emi/lib/cpp/cpp-include/unreal"
)

var dialectFlag = core.FlagDef{
	Name:  "dialect",
	Usage: "Which C++ world to target: 'generic' (portable ISO C++17 - desktop/POSIX, ESP-IDF, Arduino) or 'unreal' (Unreal Engine 5, USTRUCT/UPROPERTY). Defaults to 'generic'.",
	Type:  core.FlagString,
}

// GetCppPublicActions exposes every C++-target action (text and file based) -
// class/header generation, full module compilation, and the sdk (runtime-only)
// dump - the same shape the sibling generators expose, so it plugs into
// cmd/emi-wasm and lib/gorunner the same way. Every action accepts a `dialect`
// flag (see ResolveDialect, cpp-dialect.go) picking between the generic and
// unreal dialects - defaults to generic.
func GetCppPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "cpp:dto",
				Description:      "Generates a C++ class/struct out of a dto signature (name, fields)",
				WasmFunctionName: "cppGenDto",
				Flags:            []core.FlagDef{dialectFlag},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}
				dialect := ResolveDialect(ctx)
				chunk, err := CppCommonObjectGenerator(emiDto.Fields, ctx, CppCommonObjectContext{
					Dialect:       dialect,
					RootClassName: emiDto.GetClassName(),
				})
				if err != nil {
					return "", err
				}
				return cppAsFullDocument(chunk, dialect), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "cpp:headers",
				Description:      "Generates a C++ class out of typed headers, usable for both request and response",
				WasmFunctionName: "cppGenHeader",
				Flags:            []core.FlagDef{dialectFlag},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}
				dialect := ResolveDialect(ctx)
				chunk, err := CppHeaderClass("Headers", headers, dialect, ctx)
				if err != nil {
					return "", err
				}
				return cppAsFullDocument(chunk, dialect), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "cpp:action",
				Description:      "Creates a complete C++ client method for a single action, including path/query/request/response shapes",
				WasmFunctionName: "cppGenAction",
				Flags:            []core.FlagDef{dialectFlag},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				action, err := core.StringToEmiAction(ctx.Content)
				if err != nil {
					return "", err
				}
				dialect := ResolveDialect(ctx)
				chunk, err := CppActionRender(&action, dialect, ctx, nil)
				if err != nil {
					return "", err
				}
				if chunk == nil {
					return "", errors.New("action does not have a name")
				}
				return cppAsFullDocument(chunk, dialect), nil
			},
		},
	}

	fileActions := []core.ActionFile{
		CppPrimaryAction,
		{
			BaseAction: core.BaseAction{
				Name:             "cpp:sdk",
				Description:      "Writes only the C++ runtime (dialect-dependent: cJSON+EmiHttpTransport*/EmiWebSocketX for generic, EmiHttpClient/EmiWebSocketX for unreal) to disk",
				WasmFunctionName: "cppGenSdk",
				Flags:            []core.FlagDef{dialectFlag},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				if ResolveDialect(ctx) == DialectUnreal {
					return core.FsEmbedToVirtualFile(&cppIncludeUnreal.Content, ""), nil
				}
				return core.FsEmbedToVirtualFile(&cppIncludeGeneric.Content, ""), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "cpp:module",
				Description:      "Compiles the entire C++ module (dtos, enums, actions, remotes) and writes them to disk",
				WasmFunctionName: "cppGenModule",
				Flags: []core.FlagDef{
					dialectFlag,
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
				return CppModuleFullVirtualFiles(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{TextActions: textActions, FileActions: fileActions}
}

// CppPrimaryAction detects the emi content type (dto or module) the same way
// the sibling generators' primary actions do, and dispatches to the right
// sub-compiler.
var CppPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "cpp",
		Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
		WasmFunctionName: "cppGen",
		Flags:            []core.FlagDef{dialectFlag},
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
			return CppModuleFullVirtualFiles(&emiModule, ctx)
		}

		if type_ == "dto" {
			emiDto, err := core.StringToEmiDto(ctx.Content)
			if err != nil {
				return nil, err
			}
			dialect := ResolveDialect(ctx)
			chunk, err := CppCommonObjectGenerator(emiDto.Fields, ctx, CppCommonObjectContext{
				Dialect:       dialect,
				RootClassName: emiDto.GetClassName(),
			})
			if err != nil {
				return nil, err
			}
			return []core.VirtualFile{{
				Name:         chunk.SuggestedFileName,
				Extension:    chunk.SuggestedExtension,
				ActualScript: cppAsFullDocument(chunk, dialect),
			}}, nil
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}

func cppAsFullDocument(chunk *core.CodeChunkCompiled, dialect Dialect) string {
	if dialect == DialectUnreal {
		return CppUnrealAsFullDocument(chunk, chunk.SuggestedFileName)
	}
	return CppGenericAsFullDocument(chunk)
}
