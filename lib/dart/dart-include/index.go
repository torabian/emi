// This directory is embedded verbatim into every generated dart module
// (unless the `no-sdk` tag is passed), providing the small runtime that the
// generated dto/action files import from: request/streaming helpers around
// package:http.
package dartinclude

import "embed"

//go:embed all:runtime
var Content embed.FS
