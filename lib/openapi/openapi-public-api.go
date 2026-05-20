package openapi

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

func GetOpenAPIPublicActions() core.PublicAPIActions {

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "openapi",
				Description:      "Describes the module as an OpenAPI 3 document (YAML by default; pass --json to also emit JSON).",
				WasmFunctionName: "openapiGen",
				Flags: []core.FlagDef{
					{
						Name:  "json",
						Usage: "Also emit the OpenAPI document as JSON in addition to YAML.",
						Type:  core.FlagBool,
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
					return OpenAPIModuleDescriber(ctx, emiModule)
				}

				return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
			},
		},
	}

	return core.PublicAPIActions{
		FileActions: fileActions,
	}
}
