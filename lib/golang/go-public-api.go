package golang

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"sort"
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
				Flags: []core.FlagDef{
					{
						Name:     "pkg",
						Type:     core.FlagString,
						Usage:    "Package name of the golang",
						Required: true,
					},
					{
						Name:  "emi-runtime",
						Type:  core.FlagString,
						Usage: "Location of the emi runtime",
					},
				},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				var m map[string]string = map[string]string{}
				json.Unmarshal([]byte(ctx.Flags), &m)

				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}

				emiLocation := ""
				if val, ok := m["emi-runtime"]; ok && val != "" && val != "<nil>" {
					emiLocation = val
				}

				res, err := GoCommonStructGenerator(emiDto.Fields, ctx, GoCommonStructContext{RootClassName: emiDto.Name, EmiLocation: emiLocation})
				if err != nil {
					return "", err
				}

				return AsFullDocument(res, m["pkg"]), nil

			},
		},
	}

	fileActions := []core.ActionFile{
		{
			BaseAction: core.BaseAction{
				Name:             "go",
				Description:      "Compiles golang from .emi catalog spec file",
				WasmFunctionName: "goGen",
				Flags: []core.FlagDef{
					{
						Name:     "emigo",
						Usage:    "Add location to emigo path folder, can be also github.com/torabian/emi/emigo if you wanted to",
						Required: false,
						Type:     core.FlagString,
						Default:  "github.com/torabian/emi/emigo",
					},
					{
						Name:  "pkg",
						Type:  core.FlagString,
						Usage: "Package name of the golang",
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

					files, err := GoModuleFull(&emiModule, ctx)

					return files, err
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
// Make sure this function is public on later versions
func discoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {

		// only pick general or js/ts specific complexes for js-modules
		if complex.Compiler == "go" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
				Namespace:      complex.Namespace,
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

	type Flags struct {
		Emigo       string `json:"emigo,omitempty"`
		PackageName string `json:"pkg,omitempty"`
	}
	var f Flags = Flags{
		Emigo:       "github.com/torabian/emi/emigo",
		PackageName: DEFAULT_GO_PACKAGE,
	}

	if err := json.Unmarshal([]byte(ctx.Flags), &f); err != nil {
		fmt.Println("Flags provided are not correct:", ctx.Flags)
	}

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
			EmiLocation:         f.Emigo,
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
			ActualScript: AsFullDocument(dtoItem, f.PackageName),
		})
	}

	// var actionsRendered []*core.CodeChunkCompiled

	for _, action := range module.Actions {

		output, err := GoActionRender(action, ctx, complexes)

		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         output.SuggestedFileName,
			Extension:    output.SuggestedExtension,
			ActualScript: AsFullDocument(output, f.PackageName),
		})
	}

	return files, nil
}

func AsFullDocument(x *core.CodeChunkCompiled, packageName string) string {
	importsList := CombineGoImports(*x)
	var finalContent string = "package " + packageName + "\r\n" + importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}
func CombineGoImports(chunk core.CodeChunkCompiled) string {
	statements := map[string]struct{}{}

	// Collect unique import statements
	for _, dep := range chunk.CodeChunkDependensies {
		statement := ""
		if len(dep.Objects) > 0 {
			statement = fmt.Sprintf(`%v "%v"`, dep.Objects[0], dep.Location)
		} else {
			statement = fmt.Sprintf(`"%v"`, dep.Location)
		}
		statements[statement] = struct{}{}
	}

	// Sort statements for deterministic output
	var sorted []string
	for stmt := range statements {
		sorted = append(sorted, stmt)
	}
	sort.Strings(sorted)

	// Combine into final import block
	if len(sorted) == 0 {
		return ""
	} else if len(sorted) == 1 {
		return "import " + sorted[0]
	} else {
		return "import (\n" + strings.Join(sorted, "\n") + "\n)"
	}
}
