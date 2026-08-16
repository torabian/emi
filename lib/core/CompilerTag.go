package core

import (
	"slices"
	"strings"
)

// Helpers to keep compiler tags organized
type CTag string

func (x *MicroGenContext) HasTag(tag CTag) bool {
	if slices.Contains(strings.Split(x.Tags, ","), string(tag)) {
		return true
	}

	return false
}

// CompilerTagDoc documents a single --tags value a target compiler
// understands: the literal string passed on the CLI, and a human-readable
// explanation of what it does. Every language package that supports tags
// exposes its list via a `var CompilerTags = []core.CompilerTagDoc{...}`
// (see e.g. lib/golang/go-compiler-tags.go, lib/js/js-compiler-tags.go),
// which the `emi tags` CLI command (lib/gorunner) collects and prints -
// this is what makes tags discoverable instead of something you can only
// find by reading source.
type CompilerTagDoc struct {
	// The literal value passed via --tags, e.g. "no-sdk".
	Tag CTag

	// One-line, human readable explanation of what passing this tag does.
	Description string
}
