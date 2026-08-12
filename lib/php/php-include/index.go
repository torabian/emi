// This directory is embedded verbatim into every generated PHP module
// (unless the `no-sdk` tag is passed), providing the small runtime that the
// generated dto/action files use: request/streaming helpers around curl
// (EmiSdk\Runtime\Fetchx) and generic reflection-based (de)serialization
// (EmiSdk\Runtime\Hydrator).
package phpinclude

import "embed"

//go:embed Runtime
var Content embed.FS
