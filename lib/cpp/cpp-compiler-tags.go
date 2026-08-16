package cpp

import "github.com/torabian/emi/lib/core"

// List of compiler tags the C++ target supports. Pass them via --tags, comma
// separated, e.g. --tags no-sdk,unreal
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime files
const NoPackage core.CTag = "no-pkg" // skip generating README.md

// TagUnreal (declared in cpp-dialect.go) selects the Unreal dialect the same way
// --dialect unreal does.

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above and TagUnreal in
// cpp-dialect.go.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime files"},
	{Tag: NoPackage, Description: "Skip generating README.md"},
	{Tag: TagUnreal, Description: "Select the Unreal Engine dialect (same effect as --dialect unreal)"},
}
