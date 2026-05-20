package postman

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

func GetPostmanPublicActions() core.PublicAPIActions {

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "postman",
				Description:      "Generates a Postman v2.1 collection from an Emi module.",
				WasmFunctionName: "postmanGen",
				Flags: []core.FlagDef{
					{
						Name:  "host",
						Usage: "Postman {{HOST}} variable default value.",
						Type:  core.FlagString,
					},
					{
						Name:  "port",
						Usage: "Postman {{PORT}} variable default value.",
						Type:  core.FlagString,
					},
				},
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
					return PostmanModuleDescriber(ctx, emiModule)
				}

				return nil, errors.New("we did not find any matching type for this catalog. set emi: module. type: " + type_)
			},
		},
	}

	return core.PublicAPIActions{
		FileActions: fileActions,
	}
}
