package js

import (
	"errors"

	"github.com/torabian/emi/lib/core"
)

// ReactJsonSchemaFormFileAction is the js:rjsf ActionFile - see
// react-jsonschema-form-writer.go for what it generates. Registered from
// GetJsPublicActions, which is enough to also expose it over the CLI
// (lib/gorunner) and as a wasm export (cmd/emi-wasm).
var ReactJsonSchemaFormFileAction = core.ActionFile{
	BaseAction: core.BaseAction{
		Name: "js:rjsf",
		Description: "Generates a react-jsonschema-form (RJSF) schema + wrapper component for one dto, covering every " +
			"emi field type. Nested object/array fields need no extra code - JSON Schema expresses them natively. " +
			"one/collection relations and any/complex fields resolve to an unconstrained schema and are left to RJSF's own default rendering.",
		WasmFunctionName: "jsGenRjsfForm",
	},
	Run: func(ctx core.MicroGenContext) ([]core.VirtualFile, error) {
		type_, err := core.DetectEmiStringContentType(ctx.Content)
		if err != nil {
			return nil, err
		}

		if type_ != "dto" {
			return nil, errors.New("js:rjsf only supports a single dto definition (emi: dto) for now, got type: " + type_)
		}

		dto, err := core.StringToEmiDto(ctx.Content)
		if err != nil {
			return nil, err
		}

		return ReactJsonSchemaFormGenerator(dto)
	},
}
