package preprocessor

import (
	"errors"

	"github.com/oasdiff/yaml"
	"github.com/torabian/emi/lib/core"
)

func GetPreprocessorPublicActions() core.PublicAPIActions {

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "preprocessor",
				Description:      "Runs the core preprocessor only, and generates the final yaml that compiler sees before genering for any possible target.",
				WasmFunctionName: "preprocessorGen",
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
					err = emiModule.Preprocess()
					if err != nil {
						return nil, err
					}

					content, err := yaml.Marshal(emiModule)
					if err != nil {
						return nil, err
					}

					files := []core.VirtualFile{
						{
							Name:         "final-module",
							MimeType:     "application/yaml",
							Extension:    ".yml",
							ActualScript: string(content),
						},
					}

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
