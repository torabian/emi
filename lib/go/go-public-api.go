package gogen

import (
	"github.com/torabian/emi/lib/core"
)

func GetGoPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{}

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "go:query",
				Description:      "Compiles the queries of a module",
				WasmFunctionName: "goGenQuery",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {

				emiModule, err := core.StringToEmi(ctx.Content)
				if err != nil {
					return nil, err
				}

				return GoGenQueries(&emiModule, ctx)
			},
		},
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}
