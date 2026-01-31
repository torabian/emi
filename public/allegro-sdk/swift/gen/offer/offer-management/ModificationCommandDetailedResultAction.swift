import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action ModificationCommandDetailedResultAction
 */
struct ModificationCommandDetailedResultActionMeta {
    let name: String = "ModificationCommandDetailedResultAction"
    let url: String = "https://api.{environment}/sale/offers/promo-options-commands/{commandId}/tasks"
    let method: String = "Get"
}
/*
 struct ModificationCommandDetailedResultActionRequest {
     // reserved
 }
 */
struct ModificationCommandDetailedResultActionResponse {
    let statusCode: Int
    let headers: [String: String]
    let payload: Data?
    init(
        statusCode: Int = 200,
        headers: [String: String] = [:],
        payload: Data? = nil
    ) {
        self.statusCode = statusCode
        self.headers = headers
        self.payload = payload
    }
}
final class ModificationCommandDetailedResultActionClient {
    private static let session: URLSession = .shared
    private static func buildUrl(
        base: String,
        path: String,
        query: [String: String]
    ) -> URL? {
        guard var components = URLComponents(string: base) else {
            return nil
        }
        components.path = path
        components.queryItems = query.map {
            URLQueryItem(name: $0.key, value: $0.value)
        }
        return components.url
    }
    static func compute(
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> ModificationCommandDetailedResultActionResponse {
        let meta = ModificationCommandDetailedResultActionMeta()
        let baseUrl = ""
        guard var url = buildUrl(
            base: baseUrl,
            path: meta.url,
            query: query
        ) else {
            throw URLError(.badURL)
        }
        var request = URLRequest(url: url)
        request.httpMethod = meta.method
        request.httpBody = body
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        headers.forEach {
            request.setValue($0.value, forHTTPHeaderField: $0.key)
        }
        let (data, response) = try await session.data(for: request)
        guard let http = response as? HTTPURLResponse else {
            throw URLError(.badServerResponse)
        }
        return ModificationCommandDetailedResultActionResponse(
            statusCode: http.statusCode,
            headers: http.allHeaderFields.reduce(into: [:]) { acc, item in
                if let k = item.key as? String,
                   let v = item.value as? String {
                    acc[k] = v
                }
            },
            payload: data
        )
    }
}
  // The base class definition for modificationCommandDetailedResultActionRes
struct ModificationCommandDetailedResultActionRes: Codable {
		let tasks: [ModificationCommandDetailedResultActionResTasks]
		let modification:  ModificationCommandDetailedResultActionResModification
		let additionalMarketplaces: [ModificationCommandDetailedResultActionResAdditionalMarketplaces]
}
  // The base class definition for tasks
struct ModificationCommandDetailedResultActionResTasks: Codable {
		let offer:  ModificationCommandDetailedResultActionResTasksOffer
		let marketplaceId: String
		let scheduledAt: String
		let finishedAt: String
		let status: String
		let errors: [ModificationCommandDetailedResultActionResTasksErrors]
}
  // The base class definition for offer
struct ModificationCommandDetailedResultActionResTasksOffer: Codable {
		let id: String
}
  // The base class definition for errors
struct ModificationCommandDetailedResultActionResTasksErrors: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  ModificationCommandDetailedResultActionResTasksErrorsMetadata
}
  // The base class definition for metadata
struct ModificationCommandDetailedResultActionResTasksErrorsMetadata: Codable {
		let productId: String
}
  // The base class definition for modification
struct ModificationCommandDetailedResultActionResModification: Codable {
		let basePackage:  ModificationCommandDetailedResultActionResModificationBasePackage
		let extraPackages: [ModificationCommandDetailedResultActionResModificationExtraPackages]
		let modificationTime: String
}
  // The base class definition for basePackage
struct ModificationCommandDetailedResultActionResModificationBasePackage: Codable {
		let id: String
}
  // The base class definition for extraPackages
struct ModificationCommandDetailedResultActionResModificationExtraPackages: Codable {
		let id: String
}
  // The base class definition for additionalMarketplaces
struct ModificationCommandDetailedResultActionResAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let modification:  ModificationCommandDetailedResultActionResAdditionalMarketplacesModification
}
  // The base class definition for modification
struct ModificationCommandDetailedResultActionResAdditionalMarketplacesModification: Codable {
		let basePackage:  ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage
		let extraPackages: [ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages]
		let modificationTime: String
}
  // The base class definition for basePackage
struct ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationBasePackage: Codable {
		let id: String
}
  // The base class definition for extraPackages
struct ModificationCommandDetailedResultActionResAdditionalMarketplacesModificationExtraPackages: Codable {
		let id: String
}