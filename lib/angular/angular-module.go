package angular

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"errors"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/js"
)

// Combines entire features for a module, and creates a virtual map of the files
func AngularModuleFullVirtualFiles(module *core.Module3, ctx core.MicroGenContext) ([]core.VirtualFile, error) {

	files := []core.VirtualFile{}

	// step 1 - create the actions services
	if result, err := AngularActionsClass(module, ctx); err != nil {
		return nil, errors.New("angular actions class generation failed")
	} else {

		importsList := js.CombineImportsJsWorld(*result)

		files = append(files, core.VirtualFile{
			Name:         result.SuggestedFileName,
			ActualScript: importsList + "\r\n" + string(result.ActualScript),
			Extension:    result.SuggestedExtension,
		})
	}

	return files, nil
}
