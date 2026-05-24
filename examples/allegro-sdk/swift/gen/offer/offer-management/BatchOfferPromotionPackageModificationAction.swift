import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action BatchOfferPromotionPackageModificationAction
 */
struct BatchOfferPromotionPackageModificationActionMeta {
    let name: String = "BatchOfferPromotionPackageModificationAction"
    let url: String = "https://api.{environment}/sale/offers/promo-options-commands/{commandId}"
    let method: String = "Put"
}
/*
 struct BatchOfferPromotionPackageModificationActionRequest {
     // reserved
 }
 */
struct BatchOfferPromotionPackageModificationActionResponse {
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
final class BatchOfferPromotionPackageModificationActionClient {
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
    ) async throws -> BatchOfferPromotionPackageModificationActionResponse {
        let meta = BatchOfferPromotionPackageModificationActionMeta()
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
        return BatchOfferPromotionPackageModificationActionResponse(
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
  // The base class definition for batchOfferPromotionPackageModificationActionReq
struct BatchOfferPromotionPackageModificationActionReq: Codable {
		let offerCriteria: [BatchOfferPromotionPackageModificationActionReqOfferCriteria]
		let modification:  BatchOfferPromotionPackageModificationActionReqModification
		let additionalMarketplaces: [BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces]
}
  // The base class definition for offerCriteria
struct BatchOfferPromotionPackageModificationActionReqOfferCriteria: Codable {
		let offers: [BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers]
		let type: String
}
  // The base class definition for offers
struct BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers: Codable {
		let id: String
}
  // The base class definition for modification
struct BatchOfferPromotionPackageModificationActionReqModification: Codable {
		let basePackage:  BatchOfferPromotionPackageModificationActionReqModificationBasePackage
		let extraPackages: [BatchOfferPromotionPackageModificationActionReqModificationExtraPackages]
		let modificationTime: String
}
  // The base class definition for basePackage
struct BatchOfferPromotionPackageModificationActionReqModificationBasePackage: Codable {
		let id: String
}
  // The base class definition for extraPackages
struct BatchOfferPromotionPackageModificationActionReqModificationExtraPackages: Codable {
		let id: String
}
  // The base class definition for additionalMarketplaces
struct BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let modification:  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification
}
  // The base class definition for modification
struct BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification: Codable {
		let basePackage:  BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage
		let extraPackages: [BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages]
		let modificationTime: String
}
  // The base class definition for basePackage
struct BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage: Codable {
		let id: String
}
  // The base class definition for extraPackages
struct BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages: Codable {
		let id: String
}
  // The base class definition for batchOfferPromotionPackageModificationActionRes
struct BatchOfferPromotionPackageModificationActionRes: Codable {
		let id: String
		let taskCount:  BatchOfferPromotionPackageModificationActionResTaskCount
}
  // The base class definition for taskCount
struct BatchOfferPromotionPackageModificationActionResTaskCount: Codable {
		let failed: Int
		let success: Int
		let total: Int
}