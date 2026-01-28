import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetPromoOptionsForSellerSOffersAction
 */
struct GetPromoOptionsForSellerSOffersActionMeta {
    let name: String = "GetPromoOptionsForSellerSOffersAction"
    let url: String = "https://api.{environment}/sale/offers/promo-options"
    let method: String = "Get"
}
/*
 struct GetPromoOptionsForSellerSOffersActionRequest {
     // reserved
 }
 */
struct GetPromoOptionsForSellerSOffersActionResponse {
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
final class GetPromoOptionsForSellerSOffersActionClient {
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
    ) async throws -> GetPromoOptionsForSellerSOffersActionResponse {
        let meta = GetPromoOptionsForSellerSOffersActionMeta()
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
        return GetPromoOptionsForSellerSOffersActionResponse(
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
  // The base class definition for getPromoOptionsForSellerSOffersActionRes
struct GetPromoOptionsForSellerSOffersActionRes: Codable {
		let promoOptions: [GetPromoOptionsForSellerSOffersActionResPromoOptions]
		let count: Int
		let totalCount: Int
}
  // The base class definition for promoOptions
struct GetPromoOptionsForSellerSOffersActionResPromoOptions: Codable {
		let offerId: String
		let marketplaceId: String
		let basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage
		let extraPackages: [GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages]
		let pendingChanges:  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges
		let additionalMarketplaces: [GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces]
}
  // The base class definition for basePackage
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChanges: Codable {
		let basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage
}
  // The base class definition for basePackage
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for additionalMarketplaces
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage
		let extraPackages: [GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages]
		let pendingChanges:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges
}
  // The base class definition for basePackage
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChanges: Codable {
		let basePackage:  GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage
}
  // The base class definition for basePackage
struct GetPromoOptionsForSellerSOffersActionResPromoOptionsAdditionalMarketplacesPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}