package js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"strings"

	"github.com/torabian/emi/lib/core"
	jssdk "github.com/torabian/emi/lib/js/ts-sdk"
	tssdk "github.com/torabian/emi/lib/js/ts-sdk"
)

func AsFullDocument(x *core.CodeChunkCompiled) string {
	importsList := CombineImportsJsWorld(*x)
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	finalContent = string(core.EscapeLines([]byte(finalContent)))
	return finalContent
}

// Combines entire features for a module, and creates a virtual map of the files
func JsModuleFullVirtualFiles(module *core.Module3, ctx core.MicroGenContext) ([]core.VirtualFile, error) {

	files := []core.VirtualFile{}

	var actionsRendered []*core.CodeChunkCompiled

	// First, compile the actions, and we might need to use their tokens on the next steps.
	for _, action := range module.Actions {
		actionRendered, err := JsActionClass(action, ctx)
		if err != nil {
			return nil, err
		}
		actionsRendered = append(actionsRendered, actionRendered)
	}

	globalPacakges := []string{"qs", "@types/qs"}
	// Those actions are valid ts or js files, including some helpers for react, axios, fetch
	// and couple of more, directly can be written on the disk
	for _, action := range actionsRendered {

		for _, loc := range action.CodeChunkDependenies {
			if loc.Location == INTERNAL_SDK_LOCATION {
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

	if isTypeScript {
		// step2, write the js library which is needed by the generated code.
		files = append(files, core.FsEmbedToVirtualFile(&tssdk.Content, "sdk")...)

	} else {
		files = append(files, core.FsEmbedToVirtualFile(&jssdk.Content, "sdk")...)

	}

	return files, nil
}
