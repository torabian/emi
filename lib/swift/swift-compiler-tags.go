package swift

import "github.com/torabian/emi/lib/core"

// List of all compiler tags swift supports. Add them here before using them.
const NoSdk core.CTag = "no-sdk"

// CompilerTags lists every tag this package understands, for `emi tags` to display.
// Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip emitting the runtime files (EmiAnyCodable/EmiClientConfig/EmiWebSocketX) - use when another `emi swift` invocation into the same Swift module already provides them, since they are module-global and may only exist once"},
}
