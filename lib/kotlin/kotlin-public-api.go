package kotlin

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/torabian/emi/lib/core"
	KotlinInclude "github.com/torabian/emi/lib/kotlin/kotlin-include"
)

func GetKotlinPublicActions() core.PublicAPIActions {
	textActions := []core.ActionText{
		{
			BaseAction: core.BaseAction{
				Name:             "kotlin:dto",
				Description:      "Generate kotlin dto objects",
				WasmFunctionName: "genKotlinDto",
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

				res, err := KotlinCommonStructGenerator(emiDto.Fields, ctx, commonClassContext{RootClassName: emiDto.Name, EmiLocation: emiLocation})
				if err != nil {
					return "", err
				}

				return AsFullDocument(res, ctx.Flags["pkg"]), nil

			},
		}, {
			BaseAction: core.BaseAction{
				Name:             "kotlin:headers",
				Description:      "Generates headers, for kotlin, which can be used in client and server",
				WasmFunctionName: "kotlinGenHeader",
				Flags:            []core.FlagDef{},
			},
			Run: func(ctx core.MicroGenContext) (string, error) {

				headers, err := core.StringToEmiHeaders(ctx.Content)
				if err != nil {
					return "", err
				}

				res, err := KotlinHeaderStruct(
					kotlinHeaderStructContext{ClassName: "Anonymouse", Columns: headers, PackageName: ctx.Flags["pkg"]},
					ctx,
				)
				if err != nil {
					return "", err
				}
				return AsFullDocument(res, ctx.Flags["pkg"]), nil

			},
		},
	}

	fileActions := []core.ActionFile{
		KotlinPrimaryAction,
	}

	return core.PublicAPIActions{
		TextActions: textActions,
		FileActions: fileActions,
	}
}

var KotlinPrimaryAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name:             "kotlin",
		Description:      "Compiles kotlin module",
		WasmFunctionName: "kotlinGen",
		Flags: []core.FlagDef{
			{
				Name:  "pkg",
				Usage: "Kotlin package every generated file is written under (e.g. org.example.inventory). Defaults to 'unknownpackage' when unset. Another module's field can cross-reference a dto/entity generated with a different --pkg via 'module: <that pkg>' on a one/collection field - see kotlinCollectTargetDeps.",
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
			emiModule, err := core.StringToEmiWithPath(ctx.Content, ctx.Path)
			if err != nil {
				return nil, err
			}

			files, err := KotlinModuleFull(&emiModule, ctx)

			return files, err
		}

		return nil, errors.New("we did not find any matching type for this catalog. set emi: dto, emi: module, etc. type: " + type_)
	},
}

// DiscoverComplexes finds every module-level `complexes:` entry meant for the kotlin
// compiler (compiler: kotlin), mirroring lib/golang/go-public-api.go's own
// DiscoverComplexes. Location is expected to be the fully-qualified Kotlin import path
// including the class itself (e.g. "org.example.money.Money") - see
// CollectComplexClasses/findComplexLocation in kotlin-class-generator.go for how it's
// consumed.
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {
		if complex.Compiler == "kotlin" {
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
func KotlinModuleFull(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	globalPacakges := []string{"qs", "@types/qs"}

	pkgName := ctx.Flags["pkg"]
	if pkgName == "" {
		pkgName = "unknownpackage"
	}

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

		actionRendered, err := KotlinCommonStructGenerator(dto.Fields, ctx, commonClassContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		entitiesAndDtos = append(entitiesAndDtos, actionRendered)

		// --tags android-forms: also emit a Compose-friendly <Dto>FormState alongside
		// the dto itself - see kotlin-form-state.go.
		if ctx.HasTag(AndroidForms) {
			formState, err := KotlinFormStateGenerator(dto.Fields, dto.GetClassName(), complexes)
			if err != nil {
				return nil, err
			}
			files = append(files, core.VirtualFile{
				Name:         formState.SuggestedFileName,
				Extension:    formState.SuggestedExtension,
				ActualScript: AsFullDocument(formState, pkgName),
			})
		}
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
			ActualScript: AsFullDocument(dtoItem, pkgName),
		})
	}

	// var actionsRendered []*core.CodeChunkCompiled

	for _, action := range module.Actions {

		output, err := KotlinActionRender(action, ctx, complexes)

		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         output.SuggestedFileName,
			Extension:    output.SuggestedExtension,
			ActualScript: AsFullDocument(output, pkgName),
		})
	}

	permissionsOutput, err := KotlinPermissionsGenerate(module.Permissions, ctx)
	if err != nil {
		return nil, err
	}

	if permissionsOutput != nil {
		files = append(files, core.VirtualFile{
			Name:         permissionsOutput.SuggestedFileName,
			Extension:    permissionsOutput.SuggestedExtension,
			ActualScript: AsFullDocument(permissionsOutput, pkgName),
		})
	}

	// Append the sdk include files - skippable via --tags no-sdk when another `emi
	// kotlin` invocation into the same compilation unit already provides them (see
	// CompilerTags in kotlin-compiler-tags.go).
	if !ctx.HasTag(NoSdk) {
		files = append(files, core.GenMoveIncludeDir(&KotlinInclude.KotlinInclude)...)
	}

	return files, nil
}

func AsFullDocument(x *core.CodeChunkCompiled, packageName string) string {
	importsList := CombineJavaImport(*x)
	var finalContent string = "package " + packageName + "\r\n" + importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}

func CombineJavaImport(chunk core.CodeChunkCompiled) string {
	statements := map[string]struct{}{}

	for _, dep := range chunk.CodeChunkDependensies {
		if len(dep.Objects) > 0 {
			// Objects-carrying dependencies come from same-module dto/entity "target:"
			// relations (kotlinCollectTargetDeps in kotlin-class-generator.go, when
			// field.Module is unset). Every generated dto/action/entity class for a
			// single `emi kotlin` invocation lands in the same flat output directory
			// and the same Kotlin package (see AsFullDocument/--pkg), so sibling
			// classes never need an import - Kotlin resolves same-package references
			// on its own. (A cross-module reference, field.Module set, is instead a
			// plain Location-only dependency below, since it needs a real import into
			// a different package - see kotlinCollectTargetDeps.) There used to be an
			// unconditional `import <ClassName> "<location>" //x` here, but that's
			// Go's aliased-import syntax, not valid Kotlin, and broke compilation as
			// soon as a dto/entity actually referenced another one (e.g. `type: one`
			// / `type: collection` fields) - see examples/test-kt.
			continue
		}
		statements[dep.Location] = struct{}{}
	}

	var sorted []string
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
