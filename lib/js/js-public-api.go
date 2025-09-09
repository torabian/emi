package js

import "github.com/torabian/emi/lib/core"

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
		{
			BaseAction: core.BaseAction{
				Name:             "emi:spec",
				Description:      "Generate the emi specs",
				WasmFunctionName: "genEmiSpec",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {
				return core.GenerateJsonSpecForEmi(), nil
			},
		},
	}

	fileActions := []core.ActionFile{
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
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}
