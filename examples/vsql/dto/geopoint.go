package dto

import "fmt"

// GeoPoint demonstrates a consumer-owned complex type that emi references
// via `type: complex / complex: GeoPoint` in createuser.emi.yml. The
// vsql renderer's SQLValuer hook lets it emit a real Postgres geometry
// literal instead of getting JSON-encoded or flattened.
type GeoPoint struct {
	Lat, Lon float64
	Valid    bool
}

// SQLValue is the vsql.SQLValuer contract — kept duck-typed so this file
// has no dependency on the vsql package.
func (g GeoPoint) SQLValue() (string, bool) {
	if !g.Valid {
		return "", false
	}
	return fmt.Sprintf("ST_GeomFromText('POINT(%g %g)')", g.Lon, g.Lat), true
}
