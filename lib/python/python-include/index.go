// This directory is embedded verbatim into every generated python module
// (unless the `no-sdk` tag is passed), providing the small runtime that the
// generated dto/action files import from: request/streaming helpers around
// httpx, and generic dataclass<->dict (de)serialization.
package pythoninclude

import "embed"

// `all:` is required here (rather than a bare `*`), otherwise go:embed
// silently skips runtime/__init__.py - it starts with `_`, which the
// non-`all:` form always excludes.
//
//go:embed all:runtime
var Content embed.FS
