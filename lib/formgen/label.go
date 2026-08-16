package formgen

import (
	"regexp"
	"strings"
)

var humanizeLabelRe = regexp.MustCompile(`([a-z0-9])([A-Z])`)

// HumanizeLabel turns a camelCase field name into a "Title Case" label, e.g.
// "firstName" -> "First Name". Shared across every renderer (react-form,
// react-jsonschema-form, and whatever comes after them) so field labels stay
// consistent regardless of which compiler produced them.
func HumanizeLabel(name string) string {
	if name == "" {
		return name
	}
	spaced := humanizeLabelRe.ReplaceAllString(name, "$1 $2")
	spaced = strings.ReplaceAll(spaced, "_", " ")
	words := strings.Fields(spaced)
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}
