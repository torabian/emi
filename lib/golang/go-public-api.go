package golang

import (
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"
)

func GetGolangPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "go:dto",
				Description:      "Generates dto, for golang, both as client and server",
				WasmFunctionName: "goGenObject",
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}

				res, err := GoCommonStructGenerator(emiDto.Fields, ctx, GoCommonStructContext{RootClassName: emiDto.Name})
				if err != nil {
					return "", err
				}
				return string(res.ActualScript), nil

			},
		},
	}

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "go",
				Description:      "Compiles golang from .emi catalog spec file",
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

					return GoModuleFull(&emiModule, ctx)
				}

				return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
			},
		},
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

// Finds the ts/js compatible types.
func discoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {

		// only pick general or js/ts specific complexes for js-modules
		if complex.Compiler == "go" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}

	return items
}

type GoModuleGenerationFlags struct {
	Dtos *string `json:"dtos"`
}

func (x GoModuleGenerationFlags) GetDtos() []string {
	return strings.Split(*x.Dtos, ",")
}

// Combines entire features for a module, and creates a virtual map of the files
// which is necessary to run entire modules
func GoModuleFull(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	globalPacakges := []string{"qs", "@types/qs"}

	complexes := discoverComplexes(module)
	files := []core.VirtualFile{}

	config := GoModuleGenerationFlags{}
	json.Unmarshal([]byte(ctx.Flags), &config)

	var entitiesAndDtos []*core.CodeChunkCompiled

	for _, dto := range module.Dto {
		if config.Dtos != nil && len(*config.Dtos) > 0 && !slices.Contains(config.GetDtos(), dto.Name) {
			continue
		}

		actionRendered, err := GoCommonStructGenerator(dto.Fields, ctx, GoCommonStructContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		entitiesAndDtos = append(entitiesAndDtos, actionRendered)
	}

	// internalUsage := []string{}

	for _, dtoItem := range entitiesAndDtos {
		for _, loc := range dtoItem.CodeChunkDependensies {
			// I don't remember this
			// if strings.Contains(loc.Location, INTERNAL_SDK_JS_LOCATION) || strings.Contains(loc.Location, INTERNAL_SDK_REACT_LOCATION) {
			// 	internalUsage = append(internalUsage, loc.Location)
			// 	continue
			// }
			globalPacakges = append(globalPacakges, loc.Location)
		}

		files = append(files, core.VirtualFile{
			Name:         dtoItem.SuggestedFileName,
			Extension:    dtoItem.SuggestedExtension,
			ActualScript: AsFullDocument(dtoItem),
		})
	}

	return files, nil
}

func AsFullDocument(x *core.CodeChunkCompiled) string {
	// importsList := CombineImportsJsWorld(*x)
	importsList := ""
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}
