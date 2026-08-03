/**
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
