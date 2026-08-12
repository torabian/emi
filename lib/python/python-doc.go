package python

import "strings"

// PyDoc helps build a python triple-quoted docstring out of multiple lines,
// skipping empty ones. Mirrors kotlin.GodocComment / js NewJsDoc in spirit,
// but renders a single `"""..."""` block instead of per-line `//` comments,
// since that's the idiomatic way to document a class/function in python.
type PyDoc struct {
	lines  []string
	indent string
}

// NewPyDoc creates a new PyDoc with the given indentation prefix (e.g. "    ").
func NewPyDoc(indent string) *PyDoc {
	return &PyDoc{indent: indent}
}

// Add appends a new line to the docstring, ignored later if left blank.
func (x *PyDoc) Add(line string) *PyDoc {
	x.lines = append(x.lines, line)
	return x
}

// String renders the accumulated lines as a python docstring block, or an
// empty string when nothing meaningful was added.
func (x *PyDoc) String() string {
	nonEmpty := []string{}
	for _, line := range x.lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		nonEmpty = append(nonEmpty, line)
	}

	if len(nonEmpty) == 0 {
		return ""
	}

	if len(nonEmpty) == 1 {
		return x.indent + `"""` + nonEmpty[0] + `"""`
	}

	var b strings.Builder
	b.WriteString(x.indent + `"""` + "\n")
	for _, line := range nonEmpty {
		b.WriteString(x.indent + line + "\n")
	}
	b.WriteString(x.indent + `"""`)
	return b.String()
}
