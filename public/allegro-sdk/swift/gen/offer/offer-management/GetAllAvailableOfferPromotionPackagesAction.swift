import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetAllAvailableOfferPromotionPackagesAction
 */
struct GetAllAvailableOfferPromotionPackagesActionMeta {
    let name: String = "GetAllAvailableOfferPromotionPackagesAction"
    let url: String = "https://api.{environment}/sale/offer-promotion-packages"
    let method: String = "Get"
}
/*
 struct GetAllAvailableOfferPromotionPackagesActionRequest {
     // reserved
 }
 */
struct GetAllAvailableOfferPromotionPackagesActionResponse {
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
final class GetAllAvailableOfferPromotionPackagesActionClient {
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
    ) async throws -> GetAllAvailableOfferPromotionPackagesActionResponse {
        let meta = GetAllAvailableOfferPromotionPackagesActionMeta()
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
        return GetAllAvailableOfferPromotionPackagesActionResponse(
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
  // The base class definition for getAllAvailableOfferPromotionPackagesActionRes
struct GetAllAvailableOfferPromotionPackagesActionRes: Codable {
		let marketplaceId: String
		let basePackages: [GetAllAvailableOfferPromotionPackagesActionResBasePackages]
		let extraPackages: [GetAllAvailableOfferPromotionPackagesActionResExtraPackages]
		let additionalMarketplaces: [GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces]
}
  // The base class definition for basePackages
struct GetAllAvailableOfferPromotionPackagesActionResBasePackages: Codable {
		let id: String
		let name: String
		let cycleDuration: String
}
  // The base class definition for extraPackages
struct GetAllAvailableOfferPromotionPackagesActionResExtraPackages: Codable {
		let id: String
		let name: String
		let cycleDuration: String
}
  // The base class definition for additionalMarketplaces
struct GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let basePackages: [GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages]
		let extraPackages: [GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages]
}
  // The base class definition for basePackages
struct GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesBasePackages: Codable {
		let id: String
		let name: String
		let cycleDuration: String
}
  // The base class definition for extraPackages
struct GetAllAvailableOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages: Codable {
		let id: String
		let name: String
		let cycleDuration: String
}