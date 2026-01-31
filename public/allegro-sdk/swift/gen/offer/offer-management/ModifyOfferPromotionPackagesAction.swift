import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action ModifyOfferPromotionPackagesAction
 */
struct ModifyOfferPromotionPackagesActionMeta {
    let name: String = "ModifyOfferPromotionPackagesAction"
    let url: String = "https://api.{environment}/sale/offers/{offerId}/promo-options-modification"
    let method: String = "Post"
}
/*
 struct ModifyOfferPromotionPackagesActionRequest {
     // reserved
 }
 */
struct ModifyOfferPromotionPackagesActionResponse {
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
final class ModifyOfferPromotionPackagesActionClient {
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
    ) async throws -> ModifyOfferPromotionPackagesActionResponse {
        let meta = ModifyOfferPromotionPackagesActionMeta()
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
        return ModifyOfferPromotionPackagesActionResponse(
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
  // The base class definition for modifyOfferPromotionPackagesActionReq
struct ModifyOfferPromotionPackagesActionReq: Codable {
		let modifications: [ModifyOfferPromotionPackagesActionReqModifications]
		let additionalMarketplaces: [ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces]
}
  // The base class definition for modifications
struct ModifyOfferPromotionPackagesActionReqModifications: Codable {
		let modificationType: String
		let packageType: String
		let packageId: String
}
  // The base class definition for additionalMarketplaces
struct ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let modifications: [ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications]
}
  // The base class definition for modifications
struct ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications: Codable {
		let modificationType: String
		let packageType: String
		let packageId: String
}
  // The base class definition for modifyOfferPromotionPackagesActionRes
struct ModifyOfferPromotionPackagesActionRes: Codable {
		let offerId: String
		let marketplaceId: String
		let basePackage:  ModifyOfferPromotionPackagesActionResBasePackage
		let extraPackages: [ModifyOfferPromotionPackagesActionResExtraPackages]
		let pendingChanges:  ModifyOfferPromotionPackagesActionResPendingChanges
		let additionalMarketplaces: [ModifyOfferPromotionPackagesActionResAdditionalMarketplaces]
}
  // The base class definition for basePackage
struct ModifyOfferPromotionPackagesActionResBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct ModifyOfferPromotionPackagesActionResExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct ModifyOfferPromotionPackagesActionResPendingChanges: Codable {
		let basePackage:  ModifyOfferPromotionPackagesActionResPendingChangesBasePackage
}
  // The base class definition for basePackage
struct ModifyOfferPromotionPackagesActionResPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for additionalMarketplaces
struct ModifyOfferPromotionPackagesActionResAdditionalMarketplaces: Codable {
		let marketplaceId: String
		let basePackage:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage
		let extraPackages: [ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages]
		let pendingChanges:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges
}
  // The base class definition for basePackage
struct ModifyOfferPromotionPackagesActionResAdditionalMarketplacesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for extraPackages
struct ModifyOfferPromotionPackagesActionResAdditionalMarketplacesExtraPackages: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}
  // The base class definition for pendingChanges
struct ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChanges: Codable {
		let basePackage:  ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage
}
  // The base class definition for basePackage
struct ModifyOfferPromotionPackagesActionResAdditionalMarketplacesPendingChangesBasePackage: Codable {
		let id: String
		let validFrom: String
		let validTo: String
		let nextCycleDate: String
}