package md

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

func GetMdPublicActions() core.PublicAPIActions {

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "md",
				Description:      "Describes the module in a markdown format",
				WasmFunctionName: "goGen",
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
					files, err := MdModuleDescriber(ctx, emiModule)

					return files, err
				}

				return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
			},
		},
	}

	return core.PublicAPIActions{
		FileActions: fileActions,
	}
}
