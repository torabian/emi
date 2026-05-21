// Package sqlfiles owns the SQL template files and exposes them as an
// embed.FS so both the main demo and the test suite can read them without
// duplicating //go:embed directives.
package sqlfiles

import "embed"

//go:embed *.sql
var Files embed.FS
