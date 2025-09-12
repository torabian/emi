package js

// Combines multiple parts of an EmiAction definition into a single file and generates
// the webrequestX based class for communication

import (
	"embed"
	"encoding/json"
	"path"
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"

	ts_envelopes "github.com/torabian/emi/lib/js/ts-envelopes"
	jssdk "github.com/torabian/emi/lib/js/ts-sdk"
	tssdk "github.com/torabian/emi/lib/js/ts-sdk"
)

func AsFullDocument(x *core.CodeChunkCompiled) string {
	importsList := CombineImportsJsWorld(*x)
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}

type JsModuleGenerationConfig struct {
	Actions []string
	Remotes []string
}

type JsModuleGenerationFlags struct {
	Actions  *string `json:"actions"`
	Remotes  *string `json:"remotes"`
	Dtos     *string `json:"dtos"`
	Entities *string `json:"entities"`
}

func (x JsModuleGenerationFlags) GetActions() []string {
	return strings.Split(*x.Actions, ",")
}

func (x JsModuleGenerationFlags) GetRemotes() []string {
	return strings.Split(*x.Remotes, ",")
}

func (x JsModuleGenerationFlags) GetDtos() []string {
	return strings.Split(*x.Dtos, ",")
}

func (x JsModuleGenerationFlags) GetEntities() []string {
	return strings.Split(*x.Entities, ",")
}

// Finds the ts/js compatible types.
func discoverComplexes(module *core.Emi) []RecognizedComplex {
	items := []RecognizedComplex{}
	for _, complex := range module.Complexes {

		// only pick general or js/ts specific complexes for js-modules
		if complex.Compiler == "js" || complex.Compiler == "ts" || complex.Compiler == "" {
			items = append(items, RecognizedComplex{
				Symbol:         complex.Name,
				ImportLocation: complex.Location,
			})
		}
	}

	return items
}

// Combines entire features for a module, and creates a virtual map of the files
// which is necessary to run entire modules
func JsModuleFullVirtualFiles(module *core.Emi, ctx core.MicroGenContext) ([]core.VirtualFile, error) {
	globalPacakges := []string{"qs", "@types/qs"}

	complexes := discoverComplexes(module)
	files := []core.VirtualFile{}

	config := JsModuleGenerationFlags{}
	json.Unmarshal([]byte(ctx.Flags), &config)

	var actionsRendered []*core.CodeChunkCompiled

	for _, action := range module.Actions {
		// Filter the actions.
		if config.Actions != nil && len(*config.Actions) > 0 && !slices.Contains(config.GetActions(), action.Name) {
			continue
		}

		actionRendered, err := JsActionManifest(action, ctx, complexes)
		if err != nil {
			return nil, err
		}
		actionsRendered = append(actionsRendered, actionRendered)
	}

	for _, enum := range module.Enums {

		enumRendered, err := JsStandaloneEnum(enum, ctx)
		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         enumRendered.SuggestedFileName,
			Extension:    enumRendered.SuggestedExtension,
			ActualScript: AsFullDocument(enumRendered),
		})
	}

	for _, remote := range module.Remotes {
		if config.Remotes != nil && len(*config.Remotes) > 0 && !slices.Contains(config.GetRemotes(), remote.Name) {
			continue
		}

		actionRendered, err := JsActionManifest(remote, ctx, complexes)
		if err != nil {
			return nil, err
		}
		actionsRendered = append(actionsRendered, actionRendered)
	}

	var entitiesAndDtos []*core.CodeChunkCompiled
	for _, entity := range module.Entities {
		if config.Entities != nil && len(*config.Entities) > 0 && !slices.Contains(config.GetEntities(), entity.Name) {
			continue
		}

		actionRendered, err := JsCommonObjectGenerator(entity.Fields, ctx, JsCommonObjectContext{
			RootClassName:       entity.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		entitiesAndDtos = append(entitiesAndDtos, actionRendered)
	}

	for _, dto := range module.Dto {
		if config.Dtos != nil && len(*config.Dtos) > 0 && !slices.Contains(config.GetDtos(), dto.Name) {
			continue
		}

		actionRendered, err := JsCommonObjectGenerator(dto.Fields, ctx, JsCommonObjectContext{
			RootClassName:       dto.GetClassName(),
			RecognizedComplexes: complexes,
		})
		if err != nil {
			return nil, err
		}
		entitiesAndDtos = append(entitiesAndDtos, actionRendered)
	}

	internalUsage := []string{}
	// Those actions are valid ts or js files, including some helpers for react, axios, fetch
	// and couple of more, directly can be written on the disk
	for _, action := range actionsRendered {

		for _, loc := range action.CodeChunkDependenies {
			if strings.Contains(loc.Location, INTERNAL_SDK_JS_LOCATION) {

				internalUsage = append(internalUsage, loc.Location)
				continue
			}
			globalPacakges = append(globalPacakges, loc.Location)
		}

		files = append(files, core.VirtualFile{
			Name:         action.SuggestedFileName,
			Extension:    action.SuggestedExtension,
			ActualScript: AsFullDocument(action),
		})
	}

	for _, dtoItem := range entitiesAndDtos {
		for _, loc := range dtoItem.CodeChunkDependenies {
			if strings.Contains(loc.Location, INTERNAL_SDK_JS_LOCATION) {
				internalUsage = append(internalUsage, loc.Location)
				continue
			}
			globalPacakges = append(globalPacakges, loc.Location)
		}

		files = append(files, core.VirtualFile{
			Name:         dtoItem.SuggestedFileName,
			Extension:    dtoItem.SuggestedExtension,
			ActualScript: AsFullDocument(dtoItem),
		})
	}

	// Let's add a package.json :)
	pkg, err := GeneratePackageJSON("sdk", globalPacakges)
	if err != nil {
		return nil, err
	}

	files = append(files, core.VirtualFile{
		Name:         "package.json",
		Extension:    "",
		ActualScript: string(pkg),
	})

	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	skipEnvelopes := strings.Contains(ctx.Tags, GEN_SKIP_ENVELOPES)
	isAxiosBundle := strings.Contains(ctx.Tags, GEN_AXIOS_BUNDLE_COMPATIBILITY)

	if isAxiosBundle {
		axiosBundle, err := AxiosBundleClass(module, actionsRendered, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         axiosBundle.SuggestedFileName,
			Extension:    axiosBundle.SuggestedExtension,
			ActualScript: AsFullDocument(axiosBundle),
		})
	}

	/// Add the core sdk first

	if !skipEnvelopes {
		var source *embed.FS = &ts_envelopes.Content
		files = append(files, core.FsEmbedToVirtualFile(source, "sdk/envelopes")...)
	}

	sdkFiles := []core.VirtualFile{}
	if isTypeScript {
		sdkFiles = append(sdkFiles, core.FsEmbedToVirtualFile(&tssdk.Content, "sdk")...)
	} else {
		sdkFiles = append(sdkFiles, core.FsEmbedToVirtualFile(&jssdk.Content, "sdk")...)
	}

	for _, sdkFile := range sdkFiles {
		actual := "./" + path.Join(sdkFile.Location, sdkFile.Name, sdkFile.Extension)
		actual = strings.ReplaceAll(actual, ".ts", "")
		actual = strings.ReplaceAll(actual, ".js", "")

		if slices.Contains(internalUsage, actual) {
			files = append(files, sdkFile)
		}
	}

	return files, nil
}
