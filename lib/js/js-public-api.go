package js

import "github.com/torabian/emi/lib/core"

func GetJsPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		// {
		// 	BaseAction: core.BaseAction{
		// 		Name:             "js:action:headers",
		// 		Description:      "Generate the javascript class for an action headers extending Headers class",
		// 		WasmFunctionName: "jsGenActionHeaders",
		// 	},
		// 	Run: func(ctx core.MicroGenContext) (string, error) {
		// 		return commonJsActionStringCompiler(ctx, JsActionHeaderClass)
		// 	},
		// },
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
