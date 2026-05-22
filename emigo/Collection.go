package emigo

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Collection describes how an incoming list should be applied to an existing
// one. It exists for PATCH-style update payloads where the caller needs to
// distinguish "replace the whole set" from "add to the existing set."
//
// Collection is self-nullable: a zero value means "unset" (the field was
// never touched by the caller). Use [CollectionReplace] / [CollectionAppend]
// to construct a set value, or call [Collection.Set] / [Collection.Reset].
// This removes the need to wrap it in [Nullable] for PATCH semantics.
//
// JSON wire formats accepted by UnmarshalJSON:
//
//  1. `null` — the field was explicitly cleared. IsSet() reports true,
//     Operation/Items remain zero.
//
//  2. Bare Collection — full replacement (existing rows are deleted first):
//     [{"key": "plan", "value": "pro"}]
//     → Collection{Operation: "replace", Items: [...], isSet: true}
//
//  3. Tagged object — explicit operation:
//     {"__operation": "append", "items": [{"key": "x", "value": "y"}]}
//     → Collection{Operation: "append", Items: [...], isSet: true}
//
// The vsql renderer treats Collection as opted-out via [Collection.SQLValue]
// so unset fields disappear from generated column lists, and set fields are
// JSON-encoded as a jsonb literal when they actually land in a column slot.
// Hand-written templates can still inspect .Operation / .Items / .IsSet
// directly to drive DELETE/INSERT semantics on child tables.
type Collection[T any] struct {
	Operation string `json:"__operation"`
	Items     []T    `json:"items"`
	isSet     bool
}

// CollectionReplace builds a "replace" patch — existing rows in the child
// table should be cleared before the new items are inserted.
func CollectionReplace[T any](items []T) Collection[T] {
	return Collection[T]{Operation: "replace", Items: items, isSet: true}
}

// CollectionAppend builds an "append" patch — existing rows are preserved
// and the new items are added (or upserted) alongside them.
func CollectionAppend[T any](items []T) Collection[T] {
	return Collection[T]{Operation: "append", Items: items, isSet: true}
}

// IsSet reports whether the Collection was explicitly populated (via a
// constructor, Set, or unmarshal). A zero-valued Collection returns false.
func (c Collection[T]) IsSet() bool {
	return c.isSet
}

// Get returns a pointer to the Collection and the isSet flag. Implementing
// this signature lets duck-typed Nullable consumers (e.g. the VSQL renderer's
// derefFunc) treat Collection uniformly with [Nullable].
func (c Collection[T]) Get() (*Collection[T], bool) {
	if !c.isSet {
		return nil, false
	}
	return &c, true
}

// Set marks the Collection as populated with the given operation and items.
func (c *Collection[T]) Set(operation string, items []T) *Collection[T] {
	c.Operation = operation
	c.Items = items
	c.isSet = true
	return c
}

// Reset clears the Collection back to its unset state.
func (c *Collection[T]) Reset() {
	c.Operation = ""
	c.Items = nil
	c.isSet = false
}

// SQLValue makes Collection an SQLValuer-compatible type for the vsql
// renderer. An unset Collection is omitted from generated column lists;
// a set Collection is encoded as a jsonb literal. In practice most
// templates exclude Collection-typed fields from sqlFields/sqlFieldsExcept
// and handle them by hand — this default just keeps the renderer safe.
func (c Collection[T]) SQLValue() (string, bool) {
	if !c.isSet {
		return "", false
	}
	b, err := json.Marshal(c)
	if err != nil {
		return "NULL", true
	}
	return "'" + strings.ReplaceAll(string(b), "'", "''") + "'::jsonb", true
}

// MarshalJSON emits `null` when unset and the tagged object form otherwise.
func (c Collection[T]) MarshalJSON() ([]byte, error) {
	if !c.isSet {
		return []byte("null"), nil
	}
	type alias struct {
		Operation string `json:"__operation"`
		Items     []T    `json:"items"`
	}

	// When operation is replace, we simply return the content
	// directly, because replace is implicit. This way code works both on
	// Client and Backend part of the Golang generator
	if c.Operation == "replace" || c.Operation == "" {
		return json.Marshal(alias{Operation: c.Operation, Items: c.Items})
	}

	return json.Marshal(c.Items)
}

// UnmarshalJSON accepts a bare Collection, a tagged object, or null. Any of these
// flips isSet to true — explicit null still counts as "the caller said
// something about this field."
func (c *Collection[T]) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if string(trimmed) == "null" {
		c.Operation = ""
		c.Items = nil
		c.isSet = true
		return nil
	}
	if len(trimmed) > 0 && trimmed[0] == '[' {
		var items []T
		if err := json.Unmarshal(data, &items); err != nil {
			return err
		}
		c.Operation = "replace"
		c.Items = items
		c.isSet = true
		return nil
	}
	type alias struct {
		Operation string `json:"__operation"`
		Items     []T    `json:"items"`
	}
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	if a.Operation == "" {
		a.Operation = "replace"
	}
	c.Operation = a.Operation
	c.Items = a.Items
	c.isSet = true
	return nil
}

// NewCollection builds a "replace" patch from a variadic list of items. It is a
// convenience over [CollectionReplace] for call sites that already have the items
// individually rather than in a slice.
func NewCollection[T any](items ...T) Collection[T] {
	return Collection[T]{Operation: "replace", Items: items, isSet: true}
}

// Append adds items to the Items slice and marks the Collection as populated.
// If Operation is unset, it defaults to "append"; otherwise the existing
// Operation is preserved so callers can grow either a "replace" or "append"
// payload incrementally.
func (c *Collection[T]) Append(items ...T) *Collection[T] {
	if c.Operation == "" {
		c.Operation = "append"
	}
	c.Items = append(c.Items, items...)
	c.isSet = true
	return c
}

// Len returns the number of items currently held.
func (c Collection[T]) Len() int {
	return len(c.Items)
}

// IsAppend reports whether the Operation is "append".
func (c Collection[T]) IsAppend() bool {
	return c.Operation == "append"
}

// IsReplace reports whether the Operation is "replace".
func (c Collection[T]) IsReplace() bool {
	return c.Operation == "replace"
}

// Clear empties Items while keeping isSet=true. Use this when the caller
// wants to explicitly send an empty list — distinct from [Collection.Reset],
// which returns the Collection to its unset zero value.
func (c *Collection[T]) Clear() *Collection[T] {
	c.Items = nil
	c.isSet = true
	return c
}
