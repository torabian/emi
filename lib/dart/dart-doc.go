package dart

import "strings"

// DartDoc builds a `///`-style doc comment block, the dartdoc convention.
type DartDoc struct {
	lines  []string
	indent string
}

// NewDartDoc creates a new DartDoc with the given indentation prefix.
func NewDartDoc(indent string) *DartDoc {
	return &DartDoc{indent: indent}
}

// Add appends a new line, ignored later if left blank.
func (x *DartDoc) Add(line string) *DartDoc {
	x.lines = append(x.lines, line)
	return x
}

// String renders the accumulated lines as `///` comments, or "" if empty.
func (x *DartDoc) String() string {
	var b strings.Builder
	wrote := false
	for _, line := range x.lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		b.WriteString(x.indent + "/// " + line + "\n")
		wrote = true
	}
	if !wrote {
		return ""
	}
	return strings.TrimRight(b.String(), "\n")
}
