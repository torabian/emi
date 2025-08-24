package main

import (
	"fmt"
	"syscall/js"

	"github.com/torabian/emi/lib/core"
	emijs "github.com/torabian/emi/lib/js"
	"gopkg.in/yaml.v2"
)

// Add two numbers
func add(this js.Value, args []js.Value) any {
	a := args[0].Int()
	b := args[1].Int()
	return a + b
}

func main() {

	// Expose functions to JS global object
	js.Global().Set("js_module", js.FuncOf(generateFiles))

	fmt.Println("Go WASM initialized")
	select {} // keep running
}

func ReadModule3FromString(content string) (*core.Module3, error) {

	var data core.Module3
	err := yaml.Unmarshal([]byte(content), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func generateFiles(this js.Value, args []js.Value) any {
	content := args[0].String()
	ctx := core.MicroGenContext{
		Tags: args[1].Get("Tags").String(),
	}

	// Example: generate multiple files

	module, err := ReadModule3FromString(content)
	if err != nil {
		fmt.Println("Module error")
		return nil
	}

	files, err := emijs.JsModuleFullVirtualFiles(module, ctx)
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
