package emigo

import (
	"fmt"
	"strconv"
)

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
