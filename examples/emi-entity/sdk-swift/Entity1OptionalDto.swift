  // The base class definition for entity1OptionalDto
struct Entity1OptionalDto: Codable {
		let uniqueId: String?
		let title: String?
		let items: [Entity1OptionalDtoItems]?
		let items2: [Entity1OptionalDtoItems2]?
		let items3: [Entity2Dto]?
		let items4: [Entity2Dto]?
		let owner: Entity2Dto?
		let manager: Entity2Dto?
		let content1: Entity1OptionalDtoContent1?
		let content2: Entity1OptionalDtoContent2?
		let complex1: Money
		let subtitle: String?
		let isActive: Bool?
		let isFeatured: Bool?
		let viewCount: Int?
		let viewCountOpt: Int?
		let smallCount: Int?
		let smallCountOpt: Int?
		let bigCount: Int64?
		let bigCountOpt: Int64?
		let ratio32: Float?
		let ratio32Opt: Float?
		let ratio64: Double?
		let ratio64Opt: Double?
		let status: String?
		let statusOpt: String?
		let metadata: [String: String]?
		let metadataOpt: [String: String]?
		let rawSettings: [String: String]?
		let labels: [String]?
		let labelsOpt: [String]?
		let misc: EmiAnyCodable
		let nestedContainer: Entity1OptionalDtoNestedContainer?
		let nestedContainerOpt: Entity1OptionalDtoNestedContainerOpt?
}
  // The base class definition for items
struct Entity1OptionalDtoItems: Codable {
		let uniqueId: String?
		let item2: String
}
  // The base class definition for items2
struct Entity1OptionalDtoItems2: Codable {
		let uniqueId: String?
		let item2: String
}
  // The base class definition for content1
struct Entity1OptionalDtoContent1: Codable {
		let item1: Int64?
}
  // The base class definition for content2
struct Entity1OptionalDtoContent2: Codable {
		let item2: Int64?
}
  // The base class definition for nestedContainer
struct Entity1OptionalDtoNestedContainer: Codable {
		let nestedInner: Entity1OptionalDtoNestedContainerNestedInner?
}
  // The base class definition for nestedInner
struct Entity1OptionalDtoNestedContainerNestedInner: Codable {
		let nestedItems: [Entity1OptionalDtoNestedContainerNestedInnerNestedItems]?
		let nestedOwner: Entity2Dto?
}
  // The base class definition for nestedItems
struct Entity1OptionalDtoNestedContainerNestedInnerNestedItems: Codable {
		let uniqueId: String?
		let label: String
}
  // The base class definition for nestedContainerOpt
struct Entity1OptionalDtoNestedContainerOpt: Codable {
		let nestedInner: Entity1OptionalDtoNestedContainerOptNestedInner?
}
  // The base class definition for nestedInner
struct Entity1OptionalDtoNestedContainerOptNestedInner: Codable {
		let nestedItemsOpt: [Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt]?
}
  // The base class definition for nestedItemsOpt
struct Entity1OptionalDtoNestedContainerOptNestedInnerNestedItemsOpt: Codable {
		let uniqueId: String?
		let label: String
}