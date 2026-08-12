package php

import "encoding/json"

// escapeDoubleQuoted renders s as a double-quoted, escaped PHP string
// literal. JSON's string-escaping rules (backslash, quote, control chars)
// are a valid subset of PHP's double-quoted string escaping, so this is
// both a safe and simple way to get it right. PHP additionally treats `$`
// specially inside double-quoted strings (variable interpolation), so it's
// escaped separately.
func escapeDoubleQuoted(s string) string {
	b, _ := json.Marshal(s)
	escaped := string(b)
	// json.Marshal never introduces a bare `$`, but the source string might
	// contain one - escape it so it can't be misread as interpolation.
	out := make([]byte, 0, len(escaped))
	for i := 0; i < len(escaped); i++ {
		if escaped[i] == '$' {
			out = append(out, '\\', '$')
			continue
		}
		out = append(out, escaped[i])
	}
	return string(out)
}
