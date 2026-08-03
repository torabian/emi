  // The base class definition for entity1Dto
struct Entity1Dto: Codable {
		let uniqueId: String?
		let title: String
		let items: [Entity1DtoItems]
		let items2: [Entity1DtoItems2]?
		let items3: [Entity2Dto]?
		let items4: [Entity2Dto]?
		let owner: Entity2Dto?
		let manager: Entity2Dto?
		let content1:  Entity1DtoContent1
		let content2: Entity1DtoContent2?
		let complex1: Money
		let subtitle: String?
		let isActive: Bool
		let isFeatured: Bool?
		let viewCount: Int
		let viewCountOpt: Int?
		let smallCount: Int
		let smallCountOpt: Int?
		let bigCount: Int64
		let bigCountOpt: Int64?
		let ratio32: Float
		let ratio32Opt: Float?
		let ratio64: Double
		let ratio64Opt: Double?
		let status: String
		let statusOpt: String?
		let metadata: [String: String]
		let metadataOpt: [String: String]?
		let rawSettings: [String: String]
		let labels: [String]
		let labelsOpt: [String]?
		let misc: EmiAnyCodable
		let nestedContainer:  Entity1DtoNestedContainer
		let nestedContainerOpt: Entity1DtoNestedContainerOpt?
}
  // The base class definition for items
struct Entity1DtoItems: Codable {
		let item2: String
}
  // The base class definition for items2
struct Entity1DtoItems2: Codable {
		let item2: String
}
  // The base class definition for content1
struct Entity1DtoContent1: Codable {
		let item1: Int64?
}
  // The base class definition for content2
struct Entity1DtoContent2: Codable {
		let item2: Int64?
}
  // The base class definition for nestedContainer
struct Entity1DtoNestedContainer: Codable {
		let nestedInner:  Entity1DtoNestedContainerNestedInner
}
  // The base class definition for nestedInner
struct Entity1DtoNestedContainerNestedInner: Codable {
		let nestedItems: [Entity1DtoNestedContainerNestedInnerNestedItems]
		let nestedOwner: Entity2Dto?
}
  // The base class definition for nestedItems
struct Entity1DtoNestedContainerNestedInnerNestedItems: Codable {
		let label: String
}
  // The base class definition for nestedContainerOpt
struct Entity1DtoNestedContainerOpt: Codable {
		let nestedInner:  Entity1DtoNestedContainerOptNestedInner
}
  // The base class definition for nestedInner
struct Entity1DtoNestedContainerOptNestedInner: Codable {
		let nestedItemsOpt: [Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt]
}
  // The base class definition for nestedItemsOpt
struct Entity1DtoNestedContainerOptNestedInnerNestedItemsOpt: Codable {
		let label: String
}