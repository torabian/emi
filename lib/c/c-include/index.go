// This directory is embedded verbatim into every generated C module (unless
// the `no-sdk` tag is passed): the vendored cJSON library (cjson/cJSON.h,
// cjson/cJSON.c - MIT licensed, see the header for the original copyright
// notice) used for all (de)serialization, and the small fetchx.h runtime, emiws.h (websocket client) and emimedia.h (mic/camera frame header)
// (request/streaming helpers around libcurl) used by every generated
// action.
package cinclude

import "embed"

//go:embed cjson fetchx.h emiws.h emimedia.h emihttpd_core.h emihttpd_esp.h emihttpd_mongoose.h emisql.h emilayout.h emipack.h
var Content embed.FS
