package cpp

import "github.com/torabian/emi/lib/core"

// List of compiler tags the C++ target supports. Pass them via --tags, comma
// separated, e.g. --tags no-sdk,unreal
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime files
const NoPackage core.CTag = "no-pkg" // skip generating README.md

// TagUnreal (declared in cpp-dialect.go) selects the Unreal dialect the same way
// --dialect unreal does.
