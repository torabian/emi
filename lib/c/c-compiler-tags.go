package c

import "github.com/torabian/emi/lib/core"

// List of compiler tags the C target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (vendored cJSON + fetchx.h) files
const NoPackage core.CTag = "no-pkg" // skip generating README.md

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime (vendored cJSON + fetchx.h) files"},
	{Tag: NoPackage, Description: "Skip generating README.md"},
}
