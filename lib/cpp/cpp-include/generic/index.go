// This directory is embedded verbatim into every generic-dialect generated C++
// module (unless the `no-sdk` tag is passed): the vendored cJSON library
// (cjson/cJSON.h, cjson/cJSON.c - MIT licensed, see the header for the original
// copyright notice; same library the C target uses), plus the runtime headers
// generated dto/action code depends on (EmiJson, EmiHttpTransport* and
// EmiWebSocketX/EmiByteStream*).
package cppincludegeneric

import "embed"

//go:embed cjson *.hpp
var Content embed.FS
