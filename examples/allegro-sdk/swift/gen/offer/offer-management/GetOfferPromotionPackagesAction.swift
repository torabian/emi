import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetOfferPromotionPackagesAction
 */
struct GetOfferPromotionPackagesActionMeta {
    let name: String = "GetOfferPromotionPackagesAction"
    let url: String = "https://api.{environment}/sale/offers/{offerId}/promo-options"
    let method: String = "Get"
}
/*
 struct GetOfferPromotionPackagesActionRequest {
     // reserved
 }
 */
struct GetOfferPromotionPackagesActionResponse {
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
final class GetOfferPromotionPackagesActionClient {
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
    ) async throws -> GetOfferPromotionPackagesActionResponse {
        let meta = GetOfferPromotionPackagesActionMeta()
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
        return GetOfferPromotionPackagesActionResponse(
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
  // The base class definition for getOfferPromotionPackagesActionRes
struct GetOfferPromotionPackagesActionRes: Codable {
		let offerId: String
		let marketplaceId: String
		let basePackage:  GetOfferPromotionPackagesActionResBasePackage
		let extraPackages: [GetOfferPromotionPackagesActionResExtraPackages]
		let pendingChanges:  GetOfferPromotionPackagesActionResPendingChanges
		let additionalMarketplaces: [GetOfferPromotionPackagesActionResAdditionalMarketplaces]
}
  // The base class definition for basePackage
struct GetOfferPromotionPackagesActionResBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct GetOfferPromotionPackagesActionResExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct GetOfferPromotionPackagesActionResPendingChanges: Codable {
		let basePackage:  GetOfferPromotionPackagesActionResPendingChangesBasePackage
}
  // The base class definition for basePackage
struct GetOfferPromotionPackagesActionResPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for additionalMarketplaces
struct GetOfferPromotionPackagesActionResAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let basePackage:  GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage
		let extraPackages: [GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages]
		let pendingChanges:  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges
}
  // The base class definition for basePackage
struct GetOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct GetOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges: Codable {
		let basePackage:  GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage
}
  // The base class definition for basePackage
struct GetOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}