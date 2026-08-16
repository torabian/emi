package js

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

// ReactFormFileAction is the js:form ActionFile - see react-form-writer.go
// for what it actually generates. It's registered from GetJsPublicActions,
// which is enough to also expose it over the CLI (lib/gorunner) and as a wasm
// export (cmd/emi-wasm), both of which loop over that same list generically.
var ReactFormFileAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name: "js:form",
		Description: "Generates a single self-contained React form component (.form.tsx) for one dto, " +
			"covering every emi field type. Nested object/array fields become sibling sub-components " +
			"in the same file; one/collection fields import a {Target}Picker component from elsewhere " +
			"(see --form-relations) instead of being inlined.",
		WasmFunctionName: "jsGenForm",
		Flags: []core.FlagDef{
			{
				Name:  "form-components",
				Usage: "Import path for the primitive field widgets (TextField, NumberField, ...). Defaults to \"" + DefaultFormComponentsImport + "\".",
			},
			{
				Name:  "form-relations",
				Usage: "Import path base for one/collection relation picker components. Defaults to \"" + DefaultFormRelationsImport + "\".",
			},
		},
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}

		if type_ != "dto" {
			return nil, errors.New("js:form only supports a single dto definition (emi: dto) for now, got type: " + type_)
		}

		dto, err := core.StringToEmiDto(ctx.Content)
		if err != nil {
			return nil, err
		}

		toolset := DefaultReactFormToolset(ctx.Flags["form-components"], ctx.Flags["form-relations"])

		file, err := ReactFormGenerator(dto, toolset)
		if err != nil {
			return nil, err
		}

		return []core.VirtualFile{file}, nil
	},
}
