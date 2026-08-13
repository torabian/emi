// This directory is embedded verbatim into every unreal-dialect generated C++
// module (unless the `no-sdk` tag is passed): EmiHttpClient.h (FHttpModule
// wrapper) and EmiWebSocketX.h (IWebSocket wrapper).
package cppincludeunreal

import "embed"

//go:embed *.h
var Content embed.FS
