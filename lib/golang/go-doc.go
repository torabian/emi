package golang

import "strings"

// GodocComment helps build structured Go-style doc comments.
type GodocComment struct {
	lines  []string
	indent string
}

// Add appends a new line to the comment.
func (x *GodocComment) Add(line string) *GodocComment {
	x.lines = append(x.lines, line)
	return x
}

// String renders the comment as proper Go-style lines, like:
func (x *GodocComment) String() string {
	var data string
	for _, line := range x.lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		data += x.indent + "// " + line + "\n"
	}
	return data
}

// NewGoDoc creates a new GodocComment with the given indentation prefix.
func NewGoDoc(indent string) *GodocComment {
	return &GodocComment{indent: indent}
}
