package emigo

import (
	"bytes"
	"encoding/json"
	"strings"
)

// OneNullable holds a single item with optional operation semantics for
// PATCH-style update payloads. Unlike [ArrayNullable] / [CollectionNullable],
// it carries a single T rather than a slice.
//
// OneNullable is self-nullable: a zero value means "unset" (the field was
// never touched by the caller). Use [OneNullableSelect] / [NewOneNullable] to
// construct a set value, or call [OneNullable.Set] / [OneNullable.Reset]. This
// removes the need to wrap it in [Nullable] for PATCH semantics.
//
// JSON wire formats accepted by UnmarshalJSON:
//
//  1. `null` — the field was explicitly cleared. IsSet() reports true;
//     Operation, Item and Selector remain zero.
//
//  2. Tagged object — explicit operation. The only operation OneNullable
//     understands is "select", which carries a "selector" payload that
//     downstream SQL machinery will turn into a row key:
//     {"__operation": "select", "selector": {...}}
//
//  3. Anything else — the JSON value is unmarshalled directly into T. No
//     operation is recorded. This is the common case where the field simply
//     holds a value.
//
// The vsql renderer treats OneNullable as opted-out via [OneNullable.SQLValue]
// so unset fields disappear from generated column lists. A "select"-operation
// OneNullable is left for the selector-aware SQL resolver to fill in; any
// other OneNullable is encoded as a jsonb literal containing the stringified
// item.
type OneNullable[T any] struct {
	Operation string `json:"__operation"`
	Item      T      `json:"__item"`
	Selector  any    `json:"__selector,omitempty"`
	isSet     bool
}

// OneNullableSelect builds a "select" patch — the selector payload identifies
// an existing row that downstream SQL will resolve to a concrete key.
func OneNullableSelect[T any](selector any) OneNullable[T] {
	return OneNullable[T]{Operation: "select", Selector: selector, isSet: true}
}

// NewOneNullable builds an unoperated OneNullable holding the given item. JSON
// marshalling emits the bare T, matching the no-operation wire form.
func NewOneNullable[T any](item T) OneNullable[T] {
	return OneNullable[T]{Item: item, isSet: true}
}

// IsSet reports whether the OneNullable was explicitly populated (via a
// constructor, Set, or unmarshal). A zero-valued OneNullable returns false.
func (c OneNullable[T]) IsSet() bool {
	return c.isSet
}

// Get returns a pointer to the OneNullable and the isSet flag. Implementing
// this signature lets duck-typed Nullable consumers (e.g. the VSQL renderer's
// derefFunc) treat OneNullable uniformly with [Nullable].
func (c OneNullable[T]) Get() (*OneNullable[T], bool) {
	if !c.isSet {
		return nil, false
	}
	return &c, true
}

// Set marks the OneNullable as populated with the given item (no operation).
// Any previously stored selector is cleared.
func (c *OneNullable[T]) Set(item T) *OneNullable[T] {
	c.Operation = ""
	c.Item = item
	c.Selector = nil
	c.isSet = true
	return c
}

// SetSelector marks the OneNullable as populated with the "select" operation
// and the given selector payload. The Item is reset to its zero value.
func (c *OneNullable[T]) SetSelector(selector any) *OneNullable[T] {
	c.Operation = "select"
	c.Selector = selector
	var zero T
	c.Item = zero
	c.isSet = true
	return c
}

// Reset clears the OneNullable back to its unset state.
func (c *OneNullable[T]) Reset() {
	c.Operation = ""
	var zero T
	c.Item = zero
	c.Selector = nil
	c.isSet = false
}

// SQLValue makes OneNullable an SQLValuer-compatible type for the vsql
// renderer. An unset OneNullable is omitted from generated column lists. A
// "select"-operation OneNullable returns a placeholder for now — the resolver
// that turns Selector into a concrete row key will fill this in later. Any
// other OneNullable is encoded as a jsonb literal containing the stringified
// item.
func (c OneNullable[T]) SQLValue() (string, bool) {
	if !c.isSet {
		return "", false
	}
	if c.Operation == "select" {
		return "NULL", true
	}
	b, err := json.Marshal(c.Item)
	if err != nil {
		return "NULL", true
	}
	return "'" + strings.ReplaceAll(string(b), "'", "''") + "'::jsonb", true
}

// MarshalJSON emits `null` when unset, the tagged object form when the
// operation is "select", and the bare item otherwise.
func (c OneNullable[T]) MarshalJSON() ([]byte, error) {
	if !c.isSet {
		return []byte("null"), nil
	}
	if c.Operation == "select" {
		return json.Marshal(struct {
			Operation string `json:"__operation"`
			Selector  any    `json:"__selector,omitempty"`
		}{Operation: c.Operation, Selector: c.Selector})
	}
	return json.Marshal(c.Item)
}

// UnmarshalJSON accepts null, a tagged object carrying `__operation`, or any
// JSON value that maps directly to T. Any of these flips isSet to true —
// explicit null still counts as "the caller said something about this field."
//
// The tagged form is detected by the presence of an `__operation` key; without
// it, the payload is treated as a bare T even if it is a JSON object.
func (c *OneNullable[T]) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if string(trimmed) == "null" {
		c.Operation = ""
		var zero T
		c.Item = zero
		c.Selector = nil
		c.isSet = true
		return nil
	}
	if len(trimmed) > 0 && trimmed[0] == '{' {
		var probe map[string]json.RawMessage
		if err := json.Unmarshal(data, &probe); err == nil {
			if _, hasOp := probe["__operation"]; hasOp {
				type alias struct {
					Operation string          `json:"__operation"`
					Item      T               `json:"__item"`
					Selector  json.RawMessage `json:"__selector"`
				}
				var a alias
				if err := json.Unmarshal(data, &a); err != nil {
					return err
				}
				c.Operation = a.Operation
				c.Item = a.Item
				if len(a.Selector) > 0 {
					var sel any
					if err := json.Unmarshal(a.Selector, &sel); err != nil {
						return err
					}
					c.Selector = sel
				} else {
					c.Selector = nil
				}
				c.isSet = true
				return nil
			}
		}
	}
	var item T
	if err := json.Unmarshal(data, &item); err != nil {
		return err
	}
	c.Operation = ""
	c.Item = item
	c.Selector = nil
	c.isSet = true
	return nil
}

// IsSelect reports whether the Operation is "select".
func (c OneNullable[T]) IsSelect() bool {
	return c.Operation == "select"
}
