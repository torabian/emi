package core

import (
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// Emi files are written by hand, and a description is prose: "Used by the form: its uniqueId and
// name" is natural to write and is not valid YAML as a plain scalar (a colon followed by a space
// starts a mapping there), which fails the whole file with "mapping values are not allowed in this
// context" and no hint of what to fix.
//
// unmarshalLenient parses content as yaml.Unmarshal does. Only when that fails, it quotes the
// single-line plain values which contain such a colon and parses again; if that does not help
// either, the original error is returned, so a file which is wrong in any other way reports what is
// actually wrong with it. A valid file is never touched.
func unmarshalLenient(content []byte, out interface{}) error {
	err := yaml.Unmarshal(content, out)
	if err == nil {
		return nil
	}
	repaired, changed := quoteColonValues(string(content))
	if !changed {
		return err
	}
	if repairedErr := yaml.Unmarshal([]byte(repaired), out); repairedErr != nil {
		return err
	}
	return nil
}

var (
	// "key: value" or "- key: value" on one line
	keyValueLine = regexp.MustCompile(`^(\s*(?:-\s+)*[A-Za-z_][A-Za-z0-9_.\-]*:)[ \t]+(\S.*?)\s*$`)
	// a value which opens a block scalar: | > |- >- |+ >+ (with an indentation digit too)
	blockScalarStart = regexp.MustCompile(`^[|>][+\-0-9]*(\s+#.*)?$`)
	// a colon which would start a mapping inside a plain scalar
	colonInValue = regexp.MustCompile(`:(\s|$)`)
)

func indentOf(line string) int {
	return len(line) - len(strings.TrimLeft(line, " \t"))
}

// quoteColonValues wraps in single quotes every one-line plain value which has a ": " in it. Values
// that are quoted already, flow collections, anchors, tags and block scalars (and the lines inside
// them, which are plain text) are left alone.
func quoteColonValues(content string) (string, bool) {
	lines := strings.Split(content, "\n")
	changed := false
	blockIndent := -1 // the indent of the key whose block scalar we are inside, -1 when not in one

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if blockIndent >= 0 {
			if trimmed == "" || indentOf(line) > blockIndent {
				continue
			}
			blockIndent = -1
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		match := keyValueLine.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		value := match[2]
		if blockScalarStart.MatchString(value) {
			blockIndent = indentOf(line)
			continue
		}
		switch value[0] {
		case '"', '\'', '[', '{', '&', '*', '!', '|', '>', '%', '@', '`', '#':
			continue
		}
		if !colonInValue.MatchString(value) {
			continue
		}
		lines[i] = match[1] + " '" + strings.ReplaceAll(value, "'", "''") + "'"
		changed = true
	}
	return strings.Join(lines, "\n"), changed
}
