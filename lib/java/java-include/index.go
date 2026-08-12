// This directory is embedded verbatim into every generated Java module
// (unless the `no-sdk` tag is passed), providing the small runtime that the
// generated dto/action files use: request/streaming helpers around
// java.net.http.HttpClient, plus a shared Jackson ObjectMapper.
package javainclude

import "embed"

//go:embed *.java
var Content embed.FS
