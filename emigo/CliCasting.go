package emigo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Often happens that a dto is being created from cli
// arguments. Here we put some tools to cast them into
// the emi types.

// This is based on urfave, but you can implement
// this for other common frameworks. it should work
// right away with urfave v3
type CliCastable interface {
	String(field string) string
	Int64(field string) int64
	Float64(field string) float64
	Bool(field string) bool
	IsSet(field string) bool
}

// Flags are being defined using this, which is again similar to urfave2/3
// Proper casting functions might be required, to convert them into string flag.
// This is just a representation
type CliFlag struct {
	Name        string
	Required    bool
	Usage       string
	Description string
	Type        string

	Children []CliFlag
}

// When on cli it's passed as array, then we need to get it this way.
func CapturePossibleArray[T any](generator func(c CliCastable) T, fieldName string, c CliCastable) []T {
	var result []T
	json.Unmarshal([]byte(c.String(fieldName)), &result)

	return result
}

func CastPrimitive[T any](s string) (T, error) {
	var zero T

	switch any(zero).(type) {

	case string:
		return any(s).(T), nil

	case int:
		v, err := strconv.Atoi(s)
		return any(v).(T), err

	case int64:
		v, err := strconv.ParseInt(s, 10, 64)
		return any(v).(T), err

	case float64:
		v, err := strconv.ParseFloat(s, 64)
		return any(v).(T), err

	case bool:
		v, err := strconv.ParseBool(s)
		return any(v).(T), err

	default:
		return zero, fmt.Errorf("unsupported slice type")
	}
}
func InflatePossibleSlice[T any](raw string, target *[]T) error {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}

	// JSON array support (still useful)
	if strings.HasPrefix(raw, "[") {
		return json.Unmarshal([]byte(raw), target)
	}

	parts := strings.Split(raw, ",")
	result := make([]T, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)

		v, err := CastPrimitive[T](p)
		if err != nil {
			return err
		}

		result = append(result, v)
	}

	*target = result
	return nil
}
