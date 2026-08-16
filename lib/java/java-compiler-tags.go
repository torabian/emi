package java

import "github.com/torabian/emi/lib/core"

// List of compiler tags the Java target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (Fetchx) files
const NoPackage core.CTag = "no-pkg" // skip generating the pom.xml

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime (Fetchx) files"},
	{Tag: NoPackage, Description: "Skip generating the pom.xml"},
}
