package emigo

import (
	"crypto/rand"
	"fmt"
)

// NewUUIDv4 generates a random (version 4, variant 10) UUID and formats it as the
// standard 8-4-4-4-12 hex string, using only crypto/rand + encoding/fmt from the
// standard library - no external dependency, so it works everywhere emigo itself does
// (including under wasm, where Go's crypto/rand is backed by the platform's own CSPRNG).
//
// This exists so generated entity code can assign UniqueId in application code (a
// BeforeCreate gorm hook - see go-entity-default-fields.go's uniqueId field doc comment)
// instead of relying on a database-level column default: Postgres has gen_random_uuid()
// built in, but sqlite and MySQL don't have an equivalent, so a DB-level default isn't
// portable across dialects the way this is.
func NewUUIDv4() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// crypto/rand.Read on a real OS/wasm CSPRNG essentially never fails; if it
		// somehow does, panicking is preferable to silently handing back a zero-value
		// UUID that would collide with every other failure at the same instant.
		panic(fmt.Sprintf("emigo.NewUUIDv4: crypto/rand read failed: %v", err))
	}

	// Version 4: the 4 high bits of byte 6 are set to 0100.
	b[6] = (b[6] & 0x0f) | 0x40
	// Variant 10 (RFC 4122): the 2 high bits of byte 8 are set to 10.
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
