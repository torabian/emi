// Money is a hand-written complex type used by entity.emi.yml's "complex1" field,
// mirroring examples/emi-entity/sdk/support.go's Go Money struct - a real complex type
// is expected to implement whatever protocol its own wire representation needs, here
// just Codable (client-side Swift has no gorm/database Scan/Value concern of its own).
struct Money: Codable {
	let cents: Int64
}
