package js

import (
	"embed"
	"errors"
	"fmt"
	"path"
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"
	tssdk "github.com/torabian/emi/lib/js/ts-sdk"
)

func GetJsPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "js:fields",
				Description:      "Generates the class out of fields",
				WasmFunctionName: "jsGenObject",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				return commonJsObjectStringCompiler(ctx, JsCommonObjectGenerator)

			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "js:action",
				Description:      "Creates a complete js action, with all necessary information, except referenced elements",
				WasmFunctionName: "jsGenAction",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				return commonJsActionStringCompiler(ctx, JsActionManifest)
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "js:headers",
				Description:      "Generate the headers based on the header fields",
				WasmFunctionName: "jsGenActionHeaders",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				fields, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}

				code, err := JsHeaderClass(jsHeaderClassContext{ClassName: ctx.Flags, Columns: fields}, ctx)
				if err != nil {
					return "", err
				}

				return AsFullDocument(code), nil
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "js:dto:class",
				Description:      "Generates a dto class based on the dto signature, (name, fields)",
				WasmFunctionName: "jsGenDtoClass",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				return commonJsDtoStringCompiler(ctx,
					func(
						dto core.EmiDto,
						ctx core.MicroGenContext,
						jsctx JsCommonObjectContext,
					) (*core.CodeChunkCompiled, error) {
						return JsCommonObjectGenerator(dto.Fields, ctx, JsCommonObjectContext{
							RootClassName:       dto.GetClassName(),
							RecognizedComplexes: []RecognizedComplex{},
						})
					},
				)
			},
		},
	}

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "js",
				Description:      "Compiles a definition file catalog, and based on emi tag, it would use an appropriate sub compiler.",
				WasmFunctionName: "jsGen",
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

					return JsModuleFullVirtualFiles(&emiModule, ctx)
				}

				if type_ == "dto" {

					emiDto, err := core.StringToEmiDto(ctx.Content)
					if err != nil {
						return nil, err
					}

					result, err := JsCommonObjectGenerator(emiDto.Fields, ctx, JsCommonObjectContext{
						RootClassName:       emiDto.GetClassName(),
						RecognizedComplexes: []RecognizedComplex{},
					})

					if err != nil {
						return nil, err
					}

					files, err := detectUsedFilesAndImports(result, &tssdk.Content)
					if err != nil {
						return nil, err
					}

					return files, nil
				}

				return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
			},
		},
		{
			BaseAction: core.BaseAction{
				Name:             "js:module",
				Description:      "Compiles the entire javascript modules and writes them to disk",
				WasmFunctionName: "jsGenModule",
				Flags: []core.FlagDef{
					{
						Name:  "actions",
						Usage: "Actions to be generated - separate the action names using comma (,)",
						Type:  core.FlagString,
					},
					{
						Name:  "entities",
						Usage: "Entities which will be generated - separate the action names using comma (,)",
						Type:  core.FlagString,
					},
					{
						Name:  "remotes",
						Usage: "Remotes which will be generated - separate the action names using comma (,)",
						Type:  core.FlagString,
					},
					{
						Name:  "dtos",
						Type:  core.FlagString,
						Usage: "Dtos which will be generated - separate the action names using comma (,) ",
					},
				},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return commonJsModuleFileCompiler(ctx, JsModuleFullVirtualFiles)
			},
		},

		{
			BaseAction: core.BaseAction{
				Name:             "spec",
				Description:      "Generate the emi specs",
				WasmFunctionName: "genEmiSpec",
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
				return core.GenerateJsonSpecForEmi()
			},
		},
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

func detectUsedFilesAndImports(
	result *core.CodeChunkCompiled,
	sdkEmbedContent *embed.FS,
) ([]core.VirtualFile, error) {
	internalUsage := []string{}

	for _, loc := range result.CodeChunkDependenies {
		if strings.Contains(loc.Location, INTERNAL_SDK_JS_LOCATION) || strings.Contains(loc.Location, INTERNAL_SDK_REACT_LOCATION) {
			internalUsage = append(internalUsage, loc.Location)
		}
	}

	files := []core.VirtualFile{{
		Name:         result.SuggestedFileName,
		Extension:    result.SuggestedExtension,
		ActualScript: AsFullDocument(result),
	}}

	// Switch between SDK sources if needed
	sdkFiles := core.FsEmbedToVirtualFile(sdkEmbedContent, "sdk")

	for _, sdkFile := range sdkFiles {
		fmt.Println(sdkFile)
		actual := "./" + path.Join(sdkFile.Location, sdkFile.Name, sdkFile.Extension)
		actual = strings.TrimSuffix(actual, ".ts")
		actual = strings.TrimSuffix(actual, ".js")

		if slices.Contains(internalUsage, actual) {
			files = append(files, sdkFile)
		}
	}

	return files, nil
}
