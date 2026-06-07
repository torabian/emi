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
