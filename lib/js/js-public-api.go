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
							RootClassName: dto.GetClassName(),
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
				WasmFunctionName: "jsGenAction",
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
