package external

import "fmt"

// Money is a consumer-owned complex type (see `type: complex / complex:
// Money` on the `money` column in batch_insert_users.emi.yml) - the same
// pattern as examples/vsql/dto.GeoPoint. It demonstrates a complex type
// backed by more than one physical SQL column: emi's EmiColumn only ever
// carries a single Column projection string, so a multi-column complex type
// just lists every column it needs there ("amount_cents, currency") - it's
// Money's own Go type, not anything emi generates, that stitches the
// scanned values back together (e.g. a hand-written RowScanner that does
// rows.Scan(&row.Id, ..., &row.Money.AmountCents, &row.Money.Currency)).
type Money struct {
	AmountCents int64
	Currency    string
}

// SQLValue is the vsql.SQLValuer contract (see examples/vsql/vsql.SQLValuer)
// - kept duck-typed so this file has no dependency on that package. Only
// relevant if Money is ever used on the *write* side (a Params field); this
// example only reads it back via RETURNING, so nothing calls it, but it's
// what "complex has no wire-level nullability of its own" (see
// core.FieldTypeComplexNullable's doc comment) means in practice: Money
// decides for itself what an absent amount looks like, e.g. by returning
// include=false here, the same way emigo.Nullable[T] opts a primitive out.
func (m Money) SQLValue() (string, bool) {
	if m.Currency == "" {
		return "", false
	}
	return fmt.Sprintf("(%d, %s)", m.AmountCents, m.Currency), true
}
