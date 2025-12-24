package main

import (
	"fmt"
	"syscall/js"

	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/golang"
	emijs "github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"
	"github.com/torabian/emi/lib/querypredict"
)

func main() {

	js.Global().Set("getPublicActions", js.FuncOf(getPublicActions))

	for _, textAction := range emijs.GetJsPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range emijs.GetJsPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range golang.GetGolangPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range golang.GetGolangPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range querypredict.GetQPPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range querypredict.GetQPPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range kotlin.GetKotlinPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range kotlin.GetKotlinPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	select {}
}

func VirtualFilesFactory(
	callback func(ctx core.MicroGenContext) ([]core.VirtualFile, error),
) func(this js.Value, args []js.Value) any {

	return func(this js.Value, args []js.Value) any {

		content := args[0].String()
		ctx := core.MicroGenContext{
			Tags:    args[1].Get("Tags").String(),
			Content: content,
			Flags:   args[1].Get("Flags").String(),
		}

		files, err := callback(ctx)
		if err != nil {
			fmt.Println("Generation error:", err.Error())
			return nil
		}

		// Convert to JS array
		jsArray := js.Global().Get("Array").New()
		for _, f := range files {
			obj := map[string]any{
				"Name":         f.Name,
				"MimeType":     f.MimeType,
				"Location":     f.Location,
				"ActualScript": f.ActualScript,
				"Extension":    f.Extension,
			}
			jsArray.Call("push", js.ValueOf(obj))
		}

		return jsArray
	}

}

func StringOutFactory(
	callback func(ctx core.MicroGenContext) (string, error),
) func(this js.Value, args []js.Value) any {

	return func(this js.Value, args []js.Value) any {
		content := args[0].String()

		ctx := core.MicroGenContext{
			Tags:    args[1].Get("Tags").String(),
			Content: content,
			Flags:   args[1].Get("Flags").String(),
		}

		compiledChunk, err := callback(ctx)
		if err != nil {
			fmt.Println("Generation error:", err)
			return nil
		}

		return compiledChunk
	}
}

// Converts Go PublicAPIActions into a JS-friendly object
func getPublicActions(this js.Value, args []js.Value) any {
	actionsJs := emijs.GetJsPublicActions() // from your js package
	actionsGolang := golang.GetGolangPublicActions()
	return publicAPIActionsToJS([]core.PublicAPIActions{actionsGolang, actionsJs})
}

// Helper to convert PublicAPIActions to JS object
func publicAPIActionsToJS(actions []core.PublicAPIActions) js.Value {
	obj := js.Global().Get("Object").New()

	// TextActions
	textArr := js.Global().Get("Array").New()
	for _, action := range actions {
		for _, a := range action.TextActions {
			textArr.Call("push", actionToJS(a.BaseAction))
		}
	}
	obj.Set("TextActions", textArr)

	// FileActions
	fileArr := js.Global().Get("Array").New()
	for _, action := range actions {
		for _, a := range action.FileActions {
			fileArr.Call("push", actionToJS(a.BaseAction))
		}
	}
	obj.Set("FileActions", fileArr)

	return obj
}

// Converts a BaseAction (name, description, flags) into JS object
func actionToJS(a core.BaseAction) js.Value {
	obj := js.Global().Get("Object").New()
	obj.Set("Name", a.Name)
	obj.Set("Description", a.Description)
	obj.Set("WasmFunctionName", a.WasmFunctionName)

	flagsArr := js.Global().Get("Array").New()
	for _, f := range a.Flags {
		fObj := js.Global().Get("Object").New()
		fObj.Set("Name", f.Name)
		fObj.Set("Usage", f.Usage)
		fObj.Set("Required", f.Required)
		fObj.Set("Type", string(f.Type))
		fObj.Set("Default", f.Default)
		flagsArr.Call("push", fObj)
	}
	obj.Set("Flags", flagsArr)

	return obj
}
