package swift

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
)

func GetSwiftPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "swift:dto",
				Description:      "Generate swift dto objects",
				WasmFunctionName: "genSwiftDto",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				emiDto, err := core.StringToEmiDto(ctx.Content)
				if err != nil {
					return "", err
				}

				emiLocation := ""
				if val, ok := ctx.Flags["emi-runtime"]; ok && val != "" && val != "<nil>" {
					emiLocation = val
				}

				res, err := SwiftCommonStructGenerator(emiDto.Fields, ctx, commonClassContext{RootClassName: emiDto.Name, EmiLocation: emiLocation})
				if err != nil {
					return "", err
				}

				return AsFullDocument(res, ctx.Flags["pkg"]), nil

			},
		}, {
			BaseAction: core.BaseAction{
				Name:             "swift:headers",
				Description:      "Generates headers, for swift, which can be used in client and server",
				WasmFunctionName: "swiftGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}

				res, err := SwiftHeaderStruct(
					swiftHeaderStructContext{ClassName: "Anonymouse", Columns: headers, PackageName: ctx.Flags["pkg"]},
					ctx,
				)
				if err != nil {
					return "", err
				}
				return AsFullDocument(res, ctx.Flags["pkg"]), nil

			},
		},
	}

	fileActions := []core.ActionFile{SwiftPrimaryAction}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

var SwiftPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "swift",
		Description:      "Compiles swift module",
		WasmFunctionName: "swiftGen",
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

			files, err := SwiftFullModule(&emiModule, ctx)

			return files, err
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}

// Finds the ts/js compatible types.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {

	// Not implemented for swift yet.

	items := []RecognizedComplex{}
	// for _, complex := range module.Complexes {
	// 	if complex.Compiler == "go" {
	// 		items = append(items, RecognizedComplex{
	// 			Symbol:         complex.Name,
	// 			ImportLocation: complex.Location,
	// 		})
	// 	}
	// }

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
func SwiftFullModule(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	globalPacakges := []string{"qs", "@types/qs"}

	complexes := DiscoverComplexes(module)
	files := []core.VirtualFile{}

	config := GoModuleGenerationFlags{}
	if ctx.Flags["dtos"] != "" {
		str := ctx.Flags["dtos"]
		config.Dtos = &str
	}

	var entitiesAndDtos []*core.CodeChunkCompiled

	for _, dto := range module.Dto {
		if config.Dtos != nil && len(*config.Dtos) > 0 && !slices.Contains(config.GetDtos(), dto.Name) {
			continue
		}

		actionRendered, err := SwiftCommonStructGenerator(dto.Fields, ctx, commonClassContext{
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
			ActualScript: AsFullDocument(dtoItem, "unknownpackage"),
		})
	}

	// var actionsRendered []*core.CodeChunkCompiled

	for _, action := range module.Actions {

		output, err := SwiftActionRender(action, ctx, complexes)

		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         output.SuggestedFileName,
			Extension:    output.SuggestedExtension,
			ActualScript: AsFullDocument(output, "unknownpackage"),
		})
	}

	files = append(files, SwiftAnyCodableFile())
	files = append(files, SwiftClientConfigFile())
	files = append(files, SwiftWebSocketRuntimeFile())

	return files, nil
}

func AsFullDocument(x *core.CodeChunkCompiled, packageName string) string {
	importsList := CombineJavaImport(*x)
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}

// CombineJavaImport builds the leading `import` block for a generated Swift file.
//
// A dependency with Objects set (a specific class/type - e.g. another generated Dto
// referenced by a relation field, via castDtoNameToCodeChunk) always refers to a type
// generated into another file in the very same compiled Swift target: unlike JS/TS,
// where every file is its own module and a sibling type genuinely needs a relative
// import, Swift compiles a whole target as one unit, so every file already sees every
// other file's top-level types with zero import needed - emitting one here (as this
// used to, via a JS/TS-shaped "import Name "location" //x" statement) is not just
// unnecessary but invalid Swift syntax.
//
// Only a bare Location with no Objects (a real external Swift module name, e.g.
// "Foundation") is import-worthy - Swift's import statement is always just
// `import ModuleName`, it has no per-symbol form the way JS/TS's does, so Objects is
// never part of the emitted statement even when both are present.
func CombineJavaImport(chunk core.CodeChunkCompiled) string {
	statements := map[string]struct{}{}

	for _, dep := range chunk.CodeChunkDependensies {
		if len(dep.Objects) > 0 || dep.Location == "" {
			continue
		}
		statements[dep.Location] = struct{}{}
	}

	// Sort for deterministic output - iterating the map directly would shuffle import
	// order on every compile run.
	sorted := make([]string, 0, len(statements))
	for stmt := range statements {
		sorted = append(sorted, stmt)
	}
	sort.Strings(sorted)

	statementsX := make([]string, 0, len(sorted))
	for _, v := range sorted {
		statementsX = append(statementsX, fmt.Sprintf("import %v", v))
	}

	return strings.Join(statementsX, "\r\n")
}
