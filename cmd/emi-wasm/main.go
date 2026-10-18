//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/torabian/emi/lib/c"
	"github.com/torabian/emi/lib/core"
	"github.com/torabian/emi/lib/cpp"
	"github.com/torabian/emi/lib/csharp"
	"github.com/torabian/emi/lib/dart"
	"github.com/torabian/emi/lib/entitysql"
	"github.com/torabian/emi/lib/golang"
	"github.com/torabian/emi/lib/java"
	emijs "github.com/torabian/emi/lib/js"
	"github.com/torabian/emi/lib/kotlin"
	"github.com/torabian/emi/lib/md"
	"github.com/torabian/emi/lib/openapi"
	"github.com/torabian/emi/lib/php"
	"github.com/torabian/emi/lib/postman"
	preprocessor "github.com/torabian/emi/lib/preproceesor"
	"github.com/torabian/emi/lib/python"
	"github.com/torabian/emi/lib/querypredict"
	"github.com/torabian/emi/lib/swift"
)

func main() {

	js.Global().Set("getPublicActions", js.FuncOf(getPublicActions))
	js.Global().Set("getCompilerTags", js.FuncOf(getCompilerTags))

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

	for _, fileAction := range preprocessor.GetPreprocessorPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, fileAction := range entitysql.GetEntitySqlPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, fileAction := range md.GetMdPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, fileAction := range openapi.GetOpenAPIPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, fileAction := range postman.GetPostmanPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range kotlin.GetKotlinPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range kotlin.GetKotlinPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range swift.GetSwiftPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range swift.GetSwiftPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range python.GetPythonPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range python.GetPythonPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range dart.GetDartPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range dart.GetDartPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range csharp.GetCSharpPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range csharp.GetCSharpPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range java.GetJavaPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range java.GetJavaPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range php.GetPhpPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range php.GetPhpPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range c.GetCPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range c.GetCPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	for _, textAction := range cpp.GetCppPublicActions().TextActions {
		js.Global().Set(textAction.WasmFunctionName, js.FuncOf(StringOutFactory(textAction.Run)))
	}

	for _, fileAction := range cpp.GetCppPublicActions().FileActions {
		js.Global().Set(fileAction.WasmFunctionName, js.FuncOf(VirtualFilesFactory(fileAction.Run)))
	}

	select {}
}

// recoverGeneration turns a panic inside a compiler backend into a logged error.
// In wasm an unrecovered panic kills the Go runtime, so every later call from the
// playground would fail until the page is reloaded.
func recoverGeneration(result *any) {
	if r := recover(); r != nil {
		fmt.Println("Generation error: the compiler crashed on this input:", r)
		*result = nil
	}
}

func VirtualFilesFactory(
	callback func(ctx core.MicroGenContext) ([]core.VirtualFile, error),
) func(this js.Value, args []js.Value) any {

	return func(this js.Value, args []js.Value) (result any) {
		defer recoverGeneration(&result)

		var flags map[string]string = map[string]string{}
		json.Unmarshal([]byte(args[1].Get("Flags").String()), &flags)

		content := args[0].String()
		ctx := core.MicroGenContext{
			Tags:    args[1].Get("Tags").String(),
			Content: content,
			Flags:   flags,
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

	return func(this js.Value, args []js.Value) (result any) {
		defer recoverGeneration(&result)

		content := args[0].String()
		var flags map[string]string = map[string]string{}
		json.Unmarshal([]byte(args[1].Get("Flags").String()), &flags)

		ctx := core.MicroGenContext{
			Tags:    args[1].Get("Tags").String(),
			Content: content,
			Flags:   flags,
		}

		compiledChunk, err := callback(ctx)
		if err != nil {
			fmt.Println("Generation error:", err)
			return nil
		}

		return compiledChunk
	}
}

// compilerTagsByTarget maps each playground target (the WasmFunctionName of the
// language's module-level action) to the --tags values its compiler understands.
// The descriptions come straight from each language package's CompilerTags, so the
// playground never keeps its own copy of the list. Languages without compiler tags
// (e.g. swift) are simply absent.
func compilerTagsByTarget() map[string][]core.CompilerTagDoc {
	return map[string][]core.CompilerTagDoc{
		"goGen":           golang.CompilerTags,
		"jsGenModule":     emijs.CompilerTags,
		"kotlinGen":       kotlin.CompilerTags,
		"pythonGenModule": python.CompilerTags,
		"dartGenModule":   dart.CompilerTags,
		"csharpGenModule": csharp.CompilerTags,
		"javaGenModule":   java.CompilerTags,
		"phpGenModule":    php.CompilerTags,
		"cGenModule":      c.CompilerTags,
		"cppGenModule":    cpp.CompilerTags,
		"entitySqlGen":    entitysql.CompilerTags,
	}
}

// getCompilerTags returns { [target]: [{ Tag, Description }] } for the playground's
// compiler options dialog.
func getCompilerTags(this js.Value, args []js.Value) any {
	obj := js.Global().Get("Object").New()
	for target, tags := range compilerTagsByTarget() {
		arr := js.Global().Get("Array").New()
		for _, t := range tags {
			item := js.Global().Get("Object").New()
			item.Set("Tag", string(t.Tag))
			item.Set("Description", t.Description)
			arr.Call("push", item)
		}
		obj.Set(target, arr)
	}
	return obj
}

// Converts Go PublicAPIActions into a JS-friendly object
func getPublicActions(this js.Value, args []js.Value) any {
	actionsJs := emijs.GetJsPublicActions() // from your js package
	actionsGolang := golang.GetGolangPublicActions()
	actionsSwift := swift.GetSwiftPublicActions()
	actionsKotlin := kotlin.GetKotlinPublicActions()
	actionsPython := python.GetPythonPublicActions()
	actionsDart := dart.GetDartPublicActions()
	actionsCSharp := csharp.GetCSharpPublicActions()
	actionsJava := java.GetJavaPublicActions()
	actionsPhp := php.GetPhpPublicActions()
	actionsC := c.GetCPublicActions()
	actionsCpp := cpp.GetCppPublicActions()

	return publicAPIActionsToJS([]core.PublicAPIActions{
		actionsJs,
		actionsGolang,
		actionsSwift,
		actionsKotlin,
		actionsPython,
		actionsDart,
		actionsCSharp,
		actionsJava,
		actionsPhp,
		actionsC,
		actionsCpp,
	})
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
