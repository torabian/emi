package test

import "encoding/json"

type Nullable[T any] struct {
	Value *T   // actual value, can be nil
	IsSet bool // true if user explicitly set it
}

// Marshal JSON: omit if not set, otherwise use the Value
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.IsSet {
		return []byte("null"), nil // or omit if used with `omitempty`
	}
	if n.Value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

// Unmarshal JSON: detect if field was present
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Value = nil
		n.IsSet = true
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	n.Value = &v
	n.IsSet = true
	return nil
}

// The base class definition for singleDto

type SingleDto struct {

	// Minimum number which can be generated
	Min int

	// Maximum number which can be generated
	Max int

	// How many numbers you want to be generated based on maximum and minimum
	Count int

	// This object is optional. Can be passed or not. On typ has ? operator
	NullableObject *SingleDtoNullableObject

	// This object is not nullable. On classes, will be always instantiated, also on types, is not having ? operator
	StaticObject SingleDtoStaticObject
}

// The base class definition for nullableObject

type SingleDtoNullableObject struct {

	//
	FirstName string
}

// The base class definition for staticObject

type SingleDtoStaticObject struct {

	//
	FirstName string
}
