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

	var actionsRendered []*core.CodeChunkCompiled

	// First, compile the actions, and we might need to use their tokens on the next steps.
	for _, action := range module.Actions {
		actionRendered, err := JsActionClass(action, ctx)
		if err != nil {
			return nil, err
		}
		actionsRendered = append(actionsRendered, actionRendered)
	}

	// Those actions are valid ts or js files, including some helpers for react, axios, fetch
	// and couple of more, directly can be written on the disk
	for _, action := range actionsRendered {
		files = append(files, core.VirtualFile{
			Name:         action.SuggestedFileName,
			Extension:    action.SuggestedExtension,
			ActualScript: AsFullDocument(action),
		})
	}

	isTypeScript := strings.Contains(ctx.Tags, GEN_TYPESCRIPT_COMPATIBILITY)
	isAngular := strings.Contains(ctx.Tags, GEN_ANGULAR_COMPATIBILITY)

	// For angular, we use the rendered actions, because generated class will be based on the decisions made
	// inside that function, since we will import Meta classes and everything else from that file.
	// On Angular, Typescript is first class citizen, doesn't worth to generate javascript version
	if isTypeScript && isAngular {

		angularService, err := AngularActionsClass(module, actionsRendered, ctx)
		if err != nil {
			return nil, err
		}
		files = append(files, core.VirtualFile{
			Name:         angularService.SuggestedFileName,
			Extension:    angularService.SuggestedExtension,
			ActualScript: AsFullDocument(angularService),
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
