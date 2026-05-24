package emigo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

// Note: This file is referenced in torabian.github.io/trainings/

// Nullable[T] from a literal value of type T.
// It automatically creates a pointer to the value internally, so
// the original value is not modified if the Nullable is changed.
//
// Example:
//
//	n := NullableOf(42)        // value = 42, IsSet = true
//	n := NullableOf("hello")   // value = "hello", IsSet = true
func NullableOf[T any](v T) Nullable[T] {
	return Nullable[T]{value: &v, isSet: true}
}

// NullableOfPtr creates a Nullable[T] from a pointer of type *T.
// The Nullable will store the pointer directly, so modifying the
// Nullable value will affect the original object.
//
// Example:
//
//	s := "hello"
//	n := NullableOfPtr(&s)       // value = &s, IsSet = true
//	*n.Ptr() = "world"
//	fmt.Println(s)              // "world", original modified
func NullableOfPtr[T any](v *T) Nullable[T] {
	return Nullable[T]{value: v, isSet: true}
}

// Nullable[T] represents a generic nullable type that can distinguish
// between "unset", "null", and an actual value.
//
// Example usage:
//
//	var n Nullable[int]
//	n.Set(Ptr(42))         // Set a value
//	fmt.Println(n.OrDefault(0)) // 42
//
//	n.Set(nil)              // Explicitly set null
//	fmt.Println(n.IsNull()) // true
//
//	n.Reset()               // Make it undefined / not set
//	fmt.Println(n.isSet)    // false
type Nullable[T any] struct {
	value *T   // Actual value; can be nil
	isSet bool // True if value has been explicitly set (including nil)
}

func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.isSet || n.value == nil {
		return nil, nil
	}

	switch v := any(*n.value).(type) {
	case string, []byte, int64, float64, bool, int, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8:
		return v, nil
	default:
		// fallback json
		return json.Marshal(v)
	}
}

func (n *Nullable[T]) Scan(value interface{}) error {
	if value == nil {
		n.value = nil
		n.isSet = true
		return nil
	}

	var v T

	switch val := value.(type) {
	case []byte:
		// try json first
		if err := json.Unmarshal(val, &v); err == nil {
			n.value = &v
			n.isSet = true
			return nil
		}

		// fallback string cast
		casted, err := CastPrimitive[T](string(val))
		if err != nil {
			return err
		}

		n.value = &casted
		n.isSet = true
		return nil
	case int64:
		rv := reflect.ValueOf(&v).Elem()

		switch rv.Kind() {
		case reflect.Bool:
			rv.SetBool(val != 0)

		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			rv.SetInt(val)

		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			rv.SetUint(uint64(val))

		case reflect.Float32,
			reflect.Float64:
			rv.SetFloat(float64(val))

		default:
			return fmt.Errorf("cannot convert int64 into %T", v)
		}

		n.value = &v
		n.isSet = true
		return nil

	case string:
		casted, err := CastPrimitive[T](val)
		if err != nil {
			return err
		}

		n.value = &casted
		n.isSet = true
		return nil

	default:
		typed, ok := value.(T)
		if !ok {
			return fmt.Errorf("cannot scan %T into Nullable", value)
		}

		n.value = &typed
		n.isSet = true
		return nil
	}
}

// MarshalJSON implements JSON marshalling for Nullable.
// If the value is not set, returns JSON null.
// If value is set (even nil), marshals the actual value.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.isSet {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

// UnmarshalJSON implements JSON unmarshalling for Nullable.
// Detects whether the field was present in JSON.
//
// Examples:
//
//	var n Nullable[int]
//	_ = json.Unmarshal([]byte("42"), &n)   // n.value = 42, n.isSet = true
//	_ = json.Unmarshal([]byte("null"), &n) // n.value = nil, n.isSet = true
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.value = nil
		n.isSet = true
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	n.value = &v
	n.isSet = true
	return nil
}

// MarshalYAML implements YAML marshalling for Nullable[T].
// If the value is not set, returns nil so that YAML omits the field if "omitempty" is used.
// If the value is set (even nil), marshals the actual value.
func (n Nullable[T]) MarshalYAML() (interface{}, error) {
	if !n.isSet {
		// undefined → YAML omitempty will omit the field
		return nil, nil
	}
	if n.value == nil {
		// explicit null
		return nil, nil
	}
	return *n.value, nil
}

// UnmarshalYAML implements YAML unmarshalling for Nullable[T].
// Detects whether the field was present in YAML.
func (n *Nullable[T]) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind == yaml.ScalarNode && node.Tag == "!!null" {
		// Explicit null
		n.value = nil
		n.isSet = true
		return nil
	}

	var v T
	if err := node.Decode(&v); err != nil {
		return err
	}
	n.value = &v
	n.isSet = true
	return nil
}

// Set assigns a value to the Nullable and marks it as set.
// Can also set nil explicitly to represent JSON null.
//
// Example:
//
//	var n Nullable[string]
//	n.Set(Ptr("hello"))
func (n *Nullable[T]) Set(value *T) *Nullable[T] {
	n.value = value
	n.isSet = true
	return n
}

// IsSet returns true if the Nullable value has been explicitly set.
//
// This includes cases where the value was set to nil (representing JSON null).
// Use this method to distinguish between an unset (undefined) field and one
// that was intentionally set, even to nil.
//
// Example:
//
//	var n Nullable[int]
//	fmt.Println(n.IsSet()) // false, not set yet
//
//	n.Set(ptr(42))
//	fmt.Println(n.IsSet()) // true, value was set
//
//	n.Set(nil)
//	fmt.Println(n.IsSet()) // true, explicitly set to nil
func (n Nullable[T]) IsSet() bool {
	return n.isSet
}

// Get returns the internal pointer to the value and a boolean indicating
// whether the value has been explicitly set.
//
// The returned pointer can be nil if the value was explicitly set to nil.
// The boolean tells you if the field was set at all.
//
// Example:
//
//	var n Nullable[int]
//
//	// Not set yet
//	if v, ok := n.Get(); !ok {
//	    fmt.Println("Value is undefined") // prints this
//	}
//
//	n.Set(ptr(42))
//	if v, ok := n.Get(); ok {
//	    fmt.Println("Value:", *v) // prints "Value: 42"
//	}
//
//	n.Set(nil)
//	if v, ok := n.Get(); ok && v == nil {
//	    fmt.Println("Value explicitly set to nil") // prints this
//	}
func (n Nullable[T]) Get() (*T, bool) {
	return n.value, n.isSet
}

// Ptr returns the internal pointer to the value.
//
// Example:
//
//	var n Nullable[int]
//	v := n.Ptr() // *int or nil
func (n *Nullable[T]) Ptr() *T {
	return n.value
}

// IsNull returns true if the value was explicitly set to nil.
//
// Example:
//
//	var n Nullable[int]
//	n.Set(nil)
//	fmt.Println(n.IsNull()) // true
func (n Nullable[T]) IsNull() bool {
	return n.isSet && n.value == nil
}

// Map applies a function to the value if it is set and not nil.
//
// Example:
//
//	var n Nullable[int]
//	n.Set(Ptr(10))
//	n.Map(func(x int) int { return x * 2 })
//	fmt.Println(*n.value) // 20
func (n *Nullable[T]) Map(f func(T) T) {
	if n.value != nil {
		v := f(*n.value)
		n.value = &v
		n.isSet = true
	}
}

// OrDefault returns the value if present, otherwise returns the provided default.
//
// Example:
//
//	var n Nullable[int]
//	fmt.Println(n.OrDefault(100)) // 100
func (n *Nullable[T]) OrDefault(def T) T {
	if n.value != nil {
		return *n.value
	}
	return def
}

// Reset clears the value and marks it as not set (undefined).
//
// Example:
//
//	var n Nullable[string]
//	n.Set(Ptr("hello"))
//	n.Reset()
//	fmt.Println(n.isSet) // false
func (n *Nullable[T]) Reset() {
	n.value = nil
	n.isSet = false
}

// FromPtr creates a Nullable from a pointer. Marks it as set.
//
// Example:
//
//	n := FromPtr(Ptr(42))
//	fmt.Println(n.isSet)   // true
//	fmt.Println(*n.value)  // 42
func FromPtr[T any](ptr *T) Nullable[T] {
	return Nullable[T]{value: ptr, isSet: true}
}

// ParseNullable parses a string representation into a Nullable[T].
// Handles "", "null", and primitive conversions via CastPrimitive.
//
// Example:
//
//	var n Nullable[int]
//	_ = ParseNullable("42", &n)  // n.value = 42, n.isSet = true
//	_ = ParseNullable("null", &n) // n.value = nil, n.isSet = true
func ParseNullable[T any](s string, n *Nullable[T]) error {
	if s == "" {
		return nil
	}
	if s == "null" {
		n.isSet = true
		n.value = nil
		return nil
	}

	v, err := CastPrimitive[T](s)
	if err != nil {
		return err
	}
	n.value = &v
	n.isSet = true
	return nil
}

// ptr is a small helper to get pointer of a value
// Example: Ptr(42) returns *int pointing to 42
func Ptr[T any](v T) *T {
	return &v
}
