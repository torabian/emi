package dart

import "github.com/torabian/emi/lib/core"

// List of compiler tags the dart target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (fetchx) files
const NoPackage core.CTag = "no-pkg" // skip generating pubspec.yaml

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime (fetchx) files"},
	{Tag: NoPackage, Description: "Skip generating pubspec.yaml"},
}
