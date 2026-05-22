import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetSelectedDataOfTheParticularProductOfferAction
 */
struct GetSelectedDataOfTheParticularProductOfferActionMeta {
    let name: String = "GetSelectedDataOfTheParticularProductOfferAction"
    let url: String = "https://api.{environment}/sale/product-offers/{offerId}/parts"
    let method: String = "Get"
}
/*
 struct GetSelectedDataOfTheParticularProductOfferActionRequest {
     // reserved
 }
 */
struct GetSelectedDataOfTheParticularProductOfferActionResponse {
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
final class GetSelectedDataOfTheParticularProductOfferActionClient {
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
    ) async throws -> GetSelectedDataOfTheParticularProductOfferActionResponse {
        let meta = GetSelectedDataOfTheParticularProductOfferActionMeta()
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
        return GetSelectedDataOfTheParticularProductOfferActionResponse(
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
  // The base class definition for getSelectedDataOfTheParticularProductOfferActionRes
struct GetSelectedDataOfTheParticularProductOfferActionRes: Codable {
		  // Unique offer identifier
 let id: String
		let stock:  GetSelectedDataOfTheParticularProductOfferActionResStock
		let sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResSellingMode
		  // Marketplace-specific price information
 let additionalMarketplaces:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces
}
  // The base class definition for stock
struct GetSelectedDataOfTheParticularProductOfferActionResStock: Codable {
		  // Number of available items in stock
 let available: Int
}
  // The base class definition for sellingMode
struct GetSelectedDataOfTheParticularProductOfferActionResSellingMode: Codable {
		let price:  GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice
}
  // The base class definition for price
struct GetSelectedDataOfTheParticularProductOfferActionResSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for additionalMarketplaces
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplaces: Codable {
		let marketplaceId1:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1
		let marketplaceId2:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2
}
  // The base class definition for marketplaceId1
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1: Codable {
		let sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode
}
  // The base class definition for sellingMode
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingMode: Codable {
		let price:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice
}
  // The base class definition for price
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId1SellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for marketplaceId2
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2: Codable {
		let sellingMode:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode
}
  // The base class definition for sellingMode
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingMode: Codable {
		let price:  GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice
}
  // The base class definition for price
struct GetSelectedDataOfTheParticularProductOfferActionResAdditionalMarketplacesMarketplaceId2SellingModePrice: Codable {
		let amount: String
		let currency: String
}