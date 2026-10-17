package c

import (
	"fmt"
	"strings"
)

// escapeDoubleQuoted renders s as a double-quoted, escaped C string literal.
// Not encoding/json: it writes \u003c for "<", and C forbids a universal
// character name for a character of the basic set.
func escapeDoubleQuoted(s string) string {
	var sb strings.Builder
	sb.WriteByte('"')
	for i := 0; i < len(s); i++ {
		switch ch := s[i]; {
		case ch == '\\':
			sb.WriteString(`\\`)
		case ch == '"':
			sb.WriteString(`\"`)
		case ch == '\n':
			sb.WriteString(`\n`)
		case ch == '\r':
			sb.WriteString(`\r`)
		case ch == '\t':
			sb.WriteString(`\t`)
		case ch == '?':
			sb.WriteString(`\?`) // no trigraphs
		case ch < 0x20 || ch == 0x7f:
			fmt.Fprintf(&sb, "\\%03o", ch)
		default:
			sb.WriteByte(ch)
		}
	}
	sb.WriteByte('"')
	return sb.String()
}
