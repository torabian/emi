package swift

import "strings"

// DocCComment helps build structured Go-style doc comments.
type DocCComment struct {
	lines  []string
	indent string
}

// Add appends a new line to the comment.
func (x *DocCComment) Add(line string) *DocCComment {
	x.lines = append(x.lines, line)
	return x
}

// String renders the comment as proper Go-style lines, like:
func (x *DocCComment) String() string {
	var data string
	for _, line := range x.lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		data += x.indent + "// " + line + "\n"
	}
	return data
}

// NewDocC creates a new DocCComment with the given indentation prefix.
func NewDocC(indent string) *DocCComment {
	return &DocCComment{indent: indent}
}
