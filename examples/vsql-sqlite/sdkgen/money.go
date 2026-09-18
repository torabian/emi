package external

import "fmt"

// Money is a consumer-owned complex type (see `type: complex / complex:
// Money` on insertUsersBatch's `money` column in queries.emi.yml) - same
// pattern as examples/vsql/dto.GeoPoint and
// examples/vsql-columns/sdkgen.Money. Backed by two physical SQL columns
// (money_amount_cents, money_currency); Money's own type is what stitches
// them back together after a scan, not anything Emi generates.
type Money struct {
	AmountCents int64
	Currency    string
}

// SQLValue is the vsql.SQLValuer contract (see examples/vsql/vsql.SQLValuer)
// - kept duck-typed so this file has no dependency on that package.
func (m Money) SQLValue() (string, bool) {
	if m.Currency == "" {
		return "", false
	}
	return fmt.Sprintf("(%d, %s)", m.AmountCents, m.Currency), true
}
