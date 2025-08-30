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
