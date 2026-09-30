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

// ModuleScopedFiles is understood by every language that writes module-level singleton
// files (Permissions, Events). Those files normally have a fixed name, so two modules
// generated into the same folder overwrite each other's; with this tag the module's name
// is put in front (abac -> AbacPermissions). Off by default so existing output and the
// imports pointing at it keep working.
const ModuleScopedFiles CTag = "module-scoped-files"

var ModuleScopedFilesDoc = CompilerTagDoc{
	Tag:         ModuleScopedFiles,
	Description: "Prefix module-level files (Permissions, Events) with the module's name, e.g. AbacPermissions, so several modules can be generated into one folder without overwriting each other",
}

// ModuleFileName returns base, prefixed with the module's name (upper-cased first letter)
// when the ModuleScopedFiles tag is set and the module has a name.
func (x *MicroGenContext) ModuleFileName(module *Emi, base string) string {
	if module == nil || module.Name == "" || !x.HasTag(ModuleScopedFiles) {
		return base
	}
	return ToUpper(module.Name) + base
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
