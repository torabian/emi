package emigo

import (
	"bytes"
	"encoding/json"
	"strings"
)

// Array describes how an incoming list should be applied to an existing
// one. It exists for PATCH-style update payloads where the caller needs to
// distinguish "replace the whole set" from "add to the existing set."
//
// Array is self-nullable: a zero value means "unset" (the field was
// never touched by the caller). Use [ArrayReplace] / [ArrayAppend]
// to construct a set value, or call [Array.Set] / [Array.Reset].
// This removes the need to wrap it in [Nullable] for PATCH semantics.
//
// JSON wire formats accepted by UnmarshalJSON:
//
//  1. `null` — the field was explicitly cleared. IsSet() reports true,
//     Operation/Items remain zero.
//
//  2. Bare array — full replacement (existing rows are deleted first):
//     [{"key": "plan", "value": "pro"}]
//     → Array{Operation: "replace", Items: [...], isSet: true}
//
//  3. Tagged object — explicit operation:
//     {"__operation": "append", "items": [{"key": "x", "value": "y"}]}
//     → Array{Operation: "append", Items: [...], isSet: true}
//
// The vsql renderer treats Array as opted-out via [Array.SQLValue]
// so unset fields disappear from generated column lists, and set fields are
// JSON-encoded as a jsonb literal when they actually land in a column slot.
// Hand-written templates can still inspect .Operation / .Items / .IsSet
// directly to drive DELETE/INSERT semantics on child tables.
type Array[T any] struct {
	Operation string `json:"__operation"`
	Items     []T    `json:"items"`
	isSet     bool
}

// ArrayReplace builds a "replace" patch — existing rows in the child
// table should be cleared before the new items are inserted.
func ArrayReplace[T any](items []T) Array[T] {
	return Array[T]{Operation: "replace", Items: items, isSet: true}
}

// ArrayAppend builds an "append" patch — existing rows are preserved
// and the new items are added (or upserted) alongside them.
func ArrayAppend[T any](items []T) Array[T] {
	return Array[T]{Operation: "append", Items: items, isSet: true}
}

// IsSet reports whether the Array was explicitly populated (via a
// constructor, Set, or unmarshal). A zero-valued Array returns false.
func (c Array[T]) IsSet() bool {
	return c.isSet
}

// Get returns a pointer to the Array and the isSet flag. Implementing
// this signature lets duck-typed Nullable consumers (e.g. the VSQL renderer's
// derefFunc) treat Array uniformly with [Nullable].
func (c Array[T]) Get() (*Array[T], bool) {
	if !c.isSet {
		return nil, false
	}
	return &c, true
}

// Set marks the Array as populated with the given operation and items.
func (c *Array[T]) Set(operation string, items []T) *Array[T] {
	c.Operation = operation
	c.Items = items
	c.isSet = true
	return c
}

// Reset clears the Array back to its unset state.
func (c *Array[T]) Reset() {
	c.Operation = ""
	c.Items = nil
	c.isSet = false
}

// SQLValue makes Array an SQLValuer-compatible type for the vsql
// renderer. An unset Array is omitted from generated column lists;
// a set Array is encoded as a jsonb literal. In practice most
// templates exclude Array-typed fields from sqlFields/sqlFieldsExcept
// and handle them by hand — this default just keeps the renderer safe.
func (c Array[T]) SQLValue() (string, bool) {
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
func (c Array[T]) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON accepts a bare array, a tagged object, or null. Any of these
// flips isSet to true — explicit null still counts as "the caller said
// something about this field."
func (c *Array[T]) UnmarshalJSON(data []byte) error {
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

// NewArray builds a "replace" patch from a variadic list of items. It is a
// convenience over [ArrayReplace] for call sites that already have the items
// individually rather than in a slice.
func NewArray[T any](items ...T) Array[T] {
	return Array[T]{Operation: "replace", Items: items, isSet: true}
}

// Append adds items to the Items slice and marks the Array as populated.
// If Operation is unset, it defaults to "append"; otherwise the existing
// Operation is preserved so callers can grow either a "replace" or "append"
// payload incrementally.
func (c *Array[T]) Append(items ...T) *Array[T] {
	if c.Operation == "" {
		c.Operation = "append"
	}
	c.Items = append(c.Items, items...)
	c.isSet = true
	return c
}

// Len returns the number of items currently held.
func (c Array[T]) Len() int {
	return len(c.Items)
}

// IsAppend reports whether the Operation is "append".
func (c Array[T]) IsAppend() bool {
	return c.Operation == "append"
}

// IsReplace reports whether the Operation is "replace".
func (c Array[T]) IsReplace() bool {
	return c.Operation == "replace"
}

// Clear empties Items while keeping isSet=true. Use this when the caller
// wants to explicitly send an empty list — distinct from [Array.Reset],
// which returns the Array to its unset zero value.
func (c *Array[T]) Clear() *Array[T] {
	c.Items = nil
	c.isSet = true
	return c
}
