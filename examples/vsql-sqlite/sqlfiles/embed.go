// Package sqlfiles owns the .sql template files that selectUsersPaged
// (queries.emi.yml, `queryName:`) reads at runtime, exposing them as an
// embed.FS so the test suite doesn't duplicate a //go:embed directive.
package sqlfiles

import "embed"

//go:embed *.sql
var Files embed.FS
