package js

// Combines multiple parts of an EmiAction definition into a single file and generates
// the webrequestX based class for communication

import (
	"bytes"
	"embed"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/torabian/emi/lib/core"

	ts_envelopes "github.com/torabian/emi/lib/js/ts-envelopes"
	jssdk "github.com/torabian/emi/lib/js/ts-sdk"
	tssdk "github.com/torabian/emi/lib/js/ts-sdk"
)

func AsFullDocument(x *core.CodeChunkCompiled, ctx core.MicroGenContext) string {
	discardTypePrefix := ctx.Flags["discard-type-prefix"] != ""

	importsList := CombineImportsJsWorld(*x, discardTypePrefix)
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	// Assumed it's typescript for both js and ts. Not sure the impact.
	finalContent = FormatWithPrettier(string(core.EscapeLines([]byte(finalContent))), true)
	return finalContent
}

func FormatWithPrettier(code string, isTypeScript bool) string {
	// 1. Check env var
	prettierCmd := os.Getenv("PRETTIER_PATH")
	useNode := false

	// 2. Fallback to local .bin
	if prettierCmd == "" {
		localBin := filepath.Join("node_modules", ".bin", "prettier")
		if _, err := os.Stat(localBin); err == nil {
			prettierCmd = localBin
		}
	}

	// 3. Fallback to local prettier.cjs
	if prettierCmd == "" {
		localCJS := filepath.Join("node_modules", "prettier", "bin", "prettier.cjs")
		if _, err := os.Stat(localCJS); err == nil {
			prettierCmd = localCJS
			useNode = true
		}
	}

	// 4. Fallback to global prettier
	if prettierCmd == "" {
		prettierCmd = "prettier"
	}

	args := []string{"--stdin-filepath", "file.js"}
	if isTypeScript {
		args = []string{"--stdin-filepath", "file.ts", "--parser", "typescript"}
	}

	// If we need node, prepend it to command
	var cmd *exec.Cmd
	if useNode {
		args = append([]string{prettierCmd}, args...)
		cmd = exec.Command("node", args...)
	} else {
		cmd = exec.Command(prettierCmd, args...)
	}

	cmd.Stdin = bytes.NewBufferString(code)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return code
	}
	return out.String()
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
func DiscoverComplexes(module *core.Emi) []RecognizedComplex {
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

	complexes := DiscoverComplexes(module)
	files := []core.VirtualFile{}
	config := JsModuleGenerationFlags{}

	if ctx.Flags["dtos"] != "" {
		str := ctx.Flags["dtos"]
		config.Dtos = &str
	}

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
			ActualScript: AsFullDocument(enumRendered, ctx),
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

	var dtos []*core.CodeChunkCompiled

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
		dtos = append(dtos, actionRendered)
	}

	internalUsage := []string{}
	// Those actions are valid ts or js files, including some helpers for react, fetch
	// and couple of more, directly can be written on the disk
	for _, action := range actionsRendered {

		for _, loc := range action.CodeChunkDependensies {
			if strings.Contains(loc.Location, getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION)) || strings.Contains(loc.Location, getSdkAwareLocation(ctx, INTERNAL_SDK_REACT_LOCATION)) {

				internalUsage = append(internalUsage, loc.Location)
				continue
			}
			globalPacakges = append(globalPacakges, loc.Location)
		}

		files = append(files, core.VirtualFile{
			Name:         action.SuggestedFileName,
			Extension:    action.SuggestedExtension,
			ActualScript: AsFullDocument(action, ctx),
		})
	}

	for _, dtoItem := range dtos {
		for _, loc := range dtoItem.CodeChunkDependensies {
			if strings.Contains(loc.Location, getSdkAwareLocation(ctx, INTERNAL_SDK_JS_LOCATION)) || strings.Contains(loc.Location, getSdkAwareLocation(ctx, INTERNAL_SDK_REACT_LOCATION)) {
				internalUsage = append(internalUsage, loc.Location)
				continue
			}
			globalPacakges = append(globalPacakges, loc.Location)
		}

		files = append(files, core.VirtualFile{
			Name:         dtoItem.SuggestedFileName,
			Extension:    dtoItem.SuggestedExtension,
			ActualScript: AsFullDocument(dtoItem, ctx),
		})
	}

	// Let's add a package.json :)
	pkg, err := GeneratePackageJSON("sdk", globalPacakges, ctx)
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
