package dart

import "github.com/torabian/emi/lib/core"

// List of compiler tags the dart target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (fetchx) files
const NoPackage core.CTag = "no-pkg" // skip generating pubspec.yaml
