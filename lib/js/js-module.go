package js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"strings"

	"github.com/torabian/emi/lib/core"
	jssdk "github.com/torabian/emi/lib/js/prebuilt-sdk"
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

	// step1, is to compile all of the actions, since they are most important.
	for _, action := range module.Actions {

		actionRendered, err := JsActionClass(action, ctx)
		if err != nil {
			return nil, err
		}

		files = append(files, core.VirtualFile{
			Name:         actionRendered.SuggestedFileName,
			Extension:    actionRendered.SuggestedExtension,
			ActualScript: AsFullDocument(actionRendered),
		})

	}

	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)

	if isTypeScript {
		// step2, write the js library which is needed by the generated code.
		files = append(files, core.FsEmbedToVirtualFile(&tssdk.Content, "sdk")...)

	} else {
		files = append(files, core.FsEmbedToVirtualFile(&jssdk.Content, "sdk")...)

	}

	return files, nil
}
