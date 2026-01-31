import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetSellersOffersAction
 */
struct GetSellersOffersActionMeta {
    let name: String = "GetSellersOffersAction"
    let url: String = "https://api.{environment}/sale/offers"
    let method: String = "Get"
}
/*
 struct GetSellersOffersActionRequest {
     // reserved
 }
 */
struct GetSellersOffersActionResponse {
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
final class GetSellersOffersActionClient {
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
    ) async throws -> GetSellersOffersActionResponse {
        let meta = GetSellersOffersActionMeta()
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
        return GetSellersOffersActionResponse(
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
  // The base class definition for getSellersOffersActionRes
struct GetSellersOffersActionRes: Codable {
		  // Number of offers in this page
 let count: Int
		  // Total number of offers available
 let totalCount: Int
		let offers: [GetSellersOffersActionResOffers]
}
  // The base class definition for offers
struct GetSellersOffersActionResOffers: Codable {
		  // Offer identifier
 let id: String
		  // Offer name or title
 let name: String
		let category:  GetSellersOffersActionResOffersCategory
		let primaryImage:  GetSellersOffersActionResOffersPrimaryImage
		let sellingMode:  GetSellersOffersActionResOffersSellingMode
		let saleInfo:  GetSellersOffersActionResOffersSaleInfo
		let stock:  GetSellersOffersActionResOffersStock
		let stats:  GetSellersOffersActionResOffersStats
		let publication:  GetSellersOffersActionResOffersPublication
		let afterSalesServices:  GetSellersOffersActionResOffersAfterSalesServices
		let additionalServices:  GetSellersOffersActionResOffersAdditionalServices
		let external:  GetSellersOffersActionResOffersExternal
		let delivery:  GetSellersOffersActionResOffersDelivery
		let b2b:  GetSellersOffersActionResOffersB2b
		let fundraisingCampaign:  GetSellersOffersActionResOffersFundraisingCampaign
		  // Marketplace-specific extensions for offer
 let additionalMarketplaces: Any
}
  // The base class definition for category
struct GetSellersOffersActionResOffersCategory: Codable {
		let id: String
}
  // The base class definition for primaryImage
struct GetSellersOffersActionResOffersPrimaryImage: Codable {
		let url: String
}
  // The base class definition for sellingMode
struct GetSellersOffersActionResOffersSellingMode: Codable {
		let format: String
		let price:  GetSellersOffersActionResOffersSellingModePrice
		let priceAutomation:  GetSellersOffersActionResOffersSellingModePriceAutomation
		let minimalPrice:  GetSellersOffersActionResOffersSellingModeMinimalPrice
		let startingPrice:  GetSellersOffersActionResOffersSellingModeStartingPrice
}
  // The base class definition for price
struct GetSellersOffersActionResOffersSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for priceAutomation
struct GetSellersOffersActionResOffersSellingModePriceAutomation: Codable {
		let rule:  GetSellersOffersActionResOffersSellingModePriceAutomationRule
}
  // The base class definition for rule
struct GetSellersOffersActionResOffersSellingModePriceAutomationRule: Codable {
		let id: String
}
  // The base class definition for minimalPrice
struct GetSellersOffersActionResOffersSellingModeMinimalPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for startingPrice
struct GetSellersOffersActionResOffersSellingModeStartingPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for saleInfo
struct GetSellersOffersActionResOffersSaleInfo: Codable {
		let currentPrice:  GetSellersOffersActionResOffersSaleInfoCurrentPrice
		let biddersCount: Int
}
  // The base class definition for currentPrice
struct GetSellersOffersActionResOffersSaleInfoCurrentPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for stock
struct GetSellersOffersActionResOffersStock: Codable {
		let available: Int
		let sold: Int
}
  // The base class definition for stats
struct GetSellersOffersActionResOffersStats: Codable {
		let watchersCount: Int
		let visitsCount: Int
}
  // The base class definition for publication
struct GetSellersOffersActionResOffersPublication: Codable {
		let status: String
		let startingAt: String
		let startedAt: String
		let endingAt: String
		let endedAt: String
		let marketplaces:  GetSellersOffersActionResOffersPublicationMarketplaces
}
  // The base class definition for marketplaces
struct GetSellersOffersActionResOffersPublicationMarketplaces: Codable {
		let base:  GetSellersOffersActionResOffersPublicationMarketplacesBase
		let additional: [GetSellersOffersActionResOffersPublicationMarketplacesAdditional]
}
  // The base class definition for base
struct GetSellersOffersActionResOffersPublicationMarketplacesBase: Codable {
		let id: String
}
  // The base class definition for additional
struct GetSellersOffersActionResOffersPublicationMarketplacesAdditional: Codable {
		let id: String
}
  // The base class definition for afterSalesServices
struct GetSellersOffersActionResOffersAfterSalesServices: Codable {
		let impliedWarranty:  GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty
		let returnPolicy:  GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy
		let warranty:  GetSellersOffersActionResOffersAfterSalesServicesWarranty
}
  // The base class definition for impliedWarranty
struct GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty: Codable {
		let id: String
}
  // The base class definition for returnPolicy
struct GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy: Codable {
		let id: String
}
  // The base class definition for warranty
struct GetSellersOffersActionResOffersAfterSalesServicesWarranty: Codable {
		let id: String
}
  // The base class definition for additionalServices
struct GetSellersOffersActionResOffersAdditionalServices: Codable {
		let id: String
}
  // The base class definition for external
struct GetSellersOffersActionResOffersExternal: Codable {
		let id: String
}
  // The base class definition for delivery
struct GetSellersOffersActionResOffersDelivery: Codable {
		let shippingRates:  GetSellersOffersActionResOffersDeliveryShippingRates
}
  // The base class definition for shippingRates
struct GetSellersOffersActionResOffersDeliveryShippingRates: Codable {
		let id: String
}
  // The base class definition for b2b
struct GetSellersOffersActionResOffersB2b: Codable {
		let buyableOnlyByBusiness: Boolean
}
  // The base class definition for fundraisingCampaign
struct GetSellersOffersActionResOffersFundraisingCampaign: Codable {
		let id: String
}