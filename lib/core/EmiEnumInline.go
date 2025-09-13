package core

import (
	"regexp"
	"strings"
	"unicode"
)

// Used in EmiField as the definition of enum items
type EmiEnumInline struct {
	// Enum key which will be used in golang generation and validation
	Key string `yaml:"k,omitempty" json:"k,omitempty" jsonschema:"description=Enum key which will be used in golang generation and validation"`

	// Description of the enum for developers. It's not translated or meant to be shown to end users.
	Description string `yaml:"description,omitempty" json:"description,omitempty" jsonschema:"description=Description of the enum for developers. It's not translated or meant to be shown to end users."`
}

func (x EmiEnumInline) GetKey() string {
	return NormaliseKey(x.Key)
}

// ToEnumKey converts any string to UpperCamelCase suitable for Go enum keys
func NormaliseKey(s string) string {
	// Replace non-alphanumeric with spaces
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = re.ReplaceAllString(s, " ")

	// Split into tokens
	parts := strings.Fields(s)

	// Capitalize first rune only, keep rest as-is
	for i, p := range parts {
		if len(p) == 0 {
			continue
		}
		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		parts[i] = string(runes)
	}

	result := strings.Join(parts, "")

	// Ensure it doesn’t start with a digit
	if len(result) > 0 && unicode.IsDigit(rune(result[0])) {
		result = "_" + result
	}

	return result
}
