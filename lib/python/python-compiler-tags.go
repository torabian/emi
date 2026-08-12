package python

import "github.com/torabian/emi/lib/core"

// List of compiler tags the python target supports. Pass them via --tags,
// comma separated, e.g. --tags async,no-sdk
const Async core.CTag = "async"      // generate `async def` client functions using httpx.AsyncClient
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (fetchx/serialization) files
const NoPackage core.CTag = "no-pkg" // skip generating requirements.txt / pyproject.toml
