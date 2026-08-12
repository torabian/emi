// This directory is embedded verbatim into every generated C# module
// (unless the `no-sdk` tag is passed), providing the small runtime that the
// generated dto/action files use: request/streaming helpers around
// System.Net.Http.HttpClient.
package csharpinclude

import "embed"

//go:embed all:runtime
var Content embed.FS
