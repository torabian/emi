package kotlin

import "github.com/torabian/emi/lib/core"

// List of all compiler tags kotlin supports. Add them here before using them.
const AndroidForms core.CTag = "android-forms"
const NoSdk core.CTag = "no-sdk"

// CompilerTags lists every tag this package understands, for `emi tags` to display.
// Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: AndroidForms, Description: "Also generate a Compose-friendly <Dto>FormState holder (mutableStateOf fields, per-field validation errors, toDto()/fromDto()) alongside every dto"},
	{Tag: NoSdk, Description: "Skip embedding the runtime (emikot: MaybeField/Maybe/GResponse/EmiWebSocketX/...) files - use when another `emi kotlin` invocation into the same compilation unit (e.g. a sibling package generated for a different module) already provides them, since they're package-global and only need to exist once"},
}
