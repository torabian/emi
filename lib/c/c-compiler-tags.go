package c

import "github.com/torabian/emi/lib/core"

// List of compiler tags the C target supports. Pass them via --tags,
// comma separated, e.g. --tags no-sdk
const NoSdk core.CTag = "no-sdk"     // skip embedding the runtime (vendored cJSON + fetchx.h) files
const NoPackage core.CTag = "no-pkg" // skip generating README.md

// EspServer switches the compiler from generating a client (libcurl/websocket
// client functions) to generating the *server* side for ESP-IDF's
// esp_http_server: per action a typed handler type, a register function and,
// for `method: reactive`, websocket sessions with send/broadcast.
// Sqlite renders the entities of the module as C structs with SQLite storage
// (insert / get / update / delete / list), on top of sqlite3.h.
const Sqlite core.CTag = "sqlite"

const EspServer core.CTag = "esp-server"

// Pack gives every entity of the module Entity_pack / Entity_unpack: a compact binary form
// (no field names, varints) for storing it where there is no SQLite - a file, an NVS blob.
// It works with or without the sqlite tag and needs no sqlite3.h.
const Pack core.CTag = "pack"

// Kconfig renders the module's `config:` block as a Kconfig file (Kconfig.<module>)
// instead of leaving it to the Go target, so one yml drives `menuconfig` on
// ESP-IDF (rsource it from main/Kconfig.projbuild) and on host builds.
const Kconfig core.CTag = "kconfig"

// MongooseServer is the same server generation on top of Mongoose's event loop
// (macOS, Linux, Windows, ESP-IDF); the developer supplies mongoose.c/.h.
const MongooseServer core.CTag = "mongoose-server"

// CompilerTags lists every tag this package understands, for `emi tags` to
// display. Keep in sync with the const list above.
var CompilerTags = []core.CompilerTagDoc{
	{Tag: NoSdk, Description: "Skip embedding the runtime (vendored cJSON + fetchx.h) files"},
	{Tag: NoPackage, Description: "Skip generating README.md"},
	{Tag: Sqlite, Description: "Generate the entities of the module as C structs with SQLite storage (schema, insert/get/update/delete/list); needs sqlite3.h"},
	{Tag: Pack, Description: "Generate the entities of the module as C structs with Entity_pack/Entity_unpack, a compact binary form for storage without SQLite (a file, an NVS blob); no sqlite3.h needed"},
	{Tag: MongooseServer, Description: "Generate Mongoose server handlers (incl. websocket sessions for reactive actions) for macOS/Linux/Windows/ESP-IDF; bring your own mongoose.c/.h"},
	{Tag: Kconfig, Description: "Generate Kconfig.<module> from the config: block (bool/int/string, default, min/max, depends); symbols are <MODULE>_<NAME>, so C sees CONFIG_<MODULE>_<NAME>"},
	{Tag: EspServer, Description: "Generate ESP-IDF esp_http_server handlers (incl. websocket sessions for reactive actions) instead of client functions"},
}
