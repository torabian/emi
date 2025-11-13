package emigo

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
