package php

import "github.com/torabian/emi/lib/core"

// List of compiler tags the PHP target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (Fetchx/Hydrator) files
const NoPackage core.CTag = "no-pkg" // skip generating composer.json

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime (Fetchx/Hydrator) files"},
	{Tag: NoPackage, Description: "Skip generating composer.json"},
}
