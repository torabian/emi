package swift

import "github.com/torabian/emi/lib/core"

// emiAnyCodableSource is a minimal Codable wrapper for a "type: any" field. Swift's
// Codable protocol has no built-in way to encode/decode an unconstrained value the way
// Go's interface{} (backed by encoding/json's dynamic marshaling) or TypeScript's any
// can - a bare "Any" field simply cannot satisfy Codable, so every struct containing one
// fails to compile. EmiAnyCodable bridges that gap by hand-implementing init(from:)/
// encode(to:) over the common JSON value shapes (null, bool, number, string, array,
// object).
const emiAnyCodableSource = `/**
 * EmiAnyCodable - generated once per module, backs every "type: any" field.
 */
struct EmiAnyCodable: Codable {
	let value: Any?

	init(_ value: Any?) {
		self.value = value
	}

	init(from decoder: Decoder) throws {
		let container = try decoder.singleValueContainer()
		if container.decodeNil() {
			value = nil
		} else if let v = try? container.decode(Bool.self) {
			value = v
		} else if let v = try? container.decode(Int.self) {
			value = v
		} else if let v = try? container.decode(Double.self) {
			value = v
		} else if let v = try? container.decode(String.self) {
			value = v
		} else if let v = try? container.decode([EmiAnyCodable].self) {
			value = v.map { $0.value }
		} else if let v = try? container.decode([String: EmiAnyCodable].self) {
			value = v.mapValues { $0.value }
		} else {
			value = nil
		}
	}

	func encode(to encoder: Encoder) throws {
		var container = encoder.singleValueContainer()
		switch value {
		case nil:
			try container.encodeNil()
		case let v as Bool:
			try container.encode(v)
		case let v as Int:
			try container.encode(v)
		case let v as Double:
			try container.encode(v)
		case let v as String:
			try container.encode(v)
		case let v as [Any?]:
			try container.encode(v.map { EmiAnyCodable($0) })
		case let v as [String: Any?]:
			try container.encode(v.mapValues { EmiAnyCodable($0) })
		default:
			try container.encodeNil()
		}
	}
}
`

// SwiftAnyCodableFile is EmiAnyCodable rendered as a standalone generated file -
// unconditionally appended by SwiftFullModule (see swift-public-api.go) alongside every
// module's other generated files, so "type: any" fields always have somewhere to resolve
// to without requiring a hand-written companion file the way a "type: complex" field
// does.
func SwiftAnyCodableFile() core.VirtualFile {
	return core.VirtualFile{
		Name:         "EmiAnyCodable",
		Extension:    ".swift",
		ActualScript: emiAnyCodableSource,
	}
}
