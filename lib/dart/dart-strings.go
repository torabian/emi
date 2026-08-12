package dart

import "encoding/json"

// escapeDoubleQuoted renders s as a double-quoted, escaped string literal.
// JSON's string-escaping rules (backslash, quote, control chars) are a valid
// subset of Dart's, so this is both a safe and simple way to get it right.
func escapeDoubleQuoted(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}
