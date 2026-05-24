package emigo

import (
	"bytes"
	"encoding/json"
	"strings"
)

// ArrayNullable describes how an incoming list should be applied to an existing
// one. It exists for PATCH-style update payloads where the caller needs to
// distinguish "replace the whole set" from "add to the existing set."
//
// ArrayNullable is self-nullable: a zero value means "unset" (the field was
// never touched by the caller). Use [ArrayNullableReplace] / [ArrayNullableAppend]
// to construct a set value, or call [ArrayNullable.Set] / [ArrayNullable.Reset].
// This removes the need to wrap it in [Nullable] for PATCH semantics.
//
// JSON wire formats accepted by UnmarshalJSON:
//
//  1. `null` — the field was explicitly cleared. IsSet() reports true,
//     Operation/Items remain zero.
//
//  2. Bare array — full replacement (existing rows are deleted first):
//     [{"key": "plan", "value": "pro"}]
//     → ArrayNullable{Operation: "replace", Items: [...], isSet: true}
//
//  3. Tagged object — explicit operation:
//     {"__operation": "append", "items": [{"key": "x", "value": "y"}]}
//     → ArrayNullable{Operation: "append", Items: [...], isSet: true}
//
// The vsql renderer treats ArrayNullable as opted-out via [ArrayNullable.SQLValue]
// so unset fields disappear from generated column lists, and set fields are
// JSON-encoded as a jsonb literal when they actually land in a column slot.
// Hand-written templates can still inspect .Operation / .Items / .IsSet
// directly to drive DELETE/INSERT semantics on child tables.
type ArrayNullable[T any] struct {
	Operation string `json:"__operation"`
	Items     []T    `json:"items"`
	isSet     bool
}

// ArrayNullableReplace builds a "replace" patch — existing rows in the child
// table should be cleared before the new items are inserted.
func ArrayNullableReplace[T any](items []T) ArrayNullable[T] {
	return ArrayNullable[T]{Operation: "replace", Items: items, isSet: true}
}

// ArrayNullableAppend builds an "append" patch — existing rows are preserved
// and the new items are added (or upserted) alongside them.
func ArrayNullableAppend[T any](items []T) ArrayNullable[T] {
	return ArrayNullable[T]{Operation: "append", Items: items, isSet: true}
}

// IsSet reports whether the ArrayNullable was explicitly populated (via a
// constructor, Set, or unmarshal). A zero-valued ArrayNullable returns false.
func (c ArrayNullable[T]) IsSet() bool {
	return c.isSet
}

// Get returns a pointer to the ArrayNullable and the isSet flag. Implementing
// this signature lets duck-typed Nullable consumers (e.g. the VSQL renderer's
// derefFunc) treat ArrayNullable uniformly with [Nullable].
func (c ArrayNullable[T]) Get() (*ArrayNullable[T], bool) {
	if !c.isSet {
		return nil, false
	}
	return &c, true
}

// Set marks the ArrayNullable as populated with the given operation and items.
func (c *ArrayNullable[T]) Set(operation string, items []T) *ArrayNullable[T] {
	c.Operation = operation
	c.Items = items
	c.isSet = true
	return c
}

// Reset clears the ArrayNullable back to its unset state.
func (c *ArrayNullable[T]) Reset() {
	c.Operation = ""
	c.Items = nil
	c.isSet = false
}

// SQLValue makes ArrayNullable an SQLValuer-compatible type for the vsql
// renderer. An unset ArrayNullable is omitted from generated column lists;
// a set ArrayNullable is encoded as a jsonb literal. In practice most
// templates exclude ArrayNullable-typed fields from sqlFields/sqlFieldsExcept
// and handle them by hand — this default just keeps the renderer safe.
func (c ArrayNullable[T]) SQLValue() (string, bool) {
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
func (c ArrayNullable[T]) MarshalJSON() ([]byte, error) {
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
		return json.Marshal(c.Items)
	}

	return json.Marshal(alias{Operation: c.Operation, Items: c.Items})
}

// UnmarshalJSON accepts a bare array, a tagged object, or null. Any of these
// flips isSet to true — explicit null still counts as "the caller said
// something about this field."
func (c *ArrayNullable[T]) UnmarshalJSON(data []byte) error {
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

// NewArrayNullable builds a "replace" patch from a variadic list of items. It is a
// convenience over [ArrayNullableReplace] for call sites that already have the items
// individually rather than in a slice.
func NewArrayNullable[T any](items ...T) ArrayNullable[T] {
	return ArrayNullable[T]{Operation: "replace", Items: items, isSet: true}
}

// Append adds items to the Items slice and marks the ArrayNullable as populated.
// If Operation is unset, it defaults to "append"; otherwise the existing
// Operation is preserved so callers can grow either a "replace" or "append"
// payload incrementally.
func (c *ArrayNullable[T]) Append(items ...T) *ArrayNullable[T] {
	if c.Operation == "" {
		c.Operation = "append"
	}
	c.Items = append(c.Items, items...)
	c.isSet = true
	return c
}

// Len returns the number of items currently held.
func (c ArrayNullable[T]) Len() int {
	return len(c.Items)
}

// IsAppend reports whether the Operation is "append".
func (c ArrayNullable[T]) IsAppend() bool {
	return c.Operation == "append"
}

// IsReplace reports whether the Operation is "replace".
func (c ArrayNullable[T]) IsReplace() bool {
	return c.Operation == "replace"
}

// Clear empties Items while keeping isSet=true. Use this when the caller
// wants to explicitly send an empty list — distinct from [ArrayNullable.Reset],
// which returns the ArrayNullable to its unset zero value.
func (c *ArrayNullable[T]) Clear() *ArrayNullable[T] {
	c.Items = nil
	c.isSet = true
	return c
}
