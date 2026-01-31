import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action ModifyTheBuyNowPriceInAnOfferAction
 */
struct ModifyTheBuyNowPriceInAnOfferActionMeta {
    let name: String = "ModifyTheBuyNowPriceInAnOfferAction"
    let url: String = "https://api.{environment}/offers/{offerId}/change-price-commands/{commandId}"
    let method: String = "Put"
}
/*
 struct ModifyTheBuyNowPriceInAnOfferActionRequest {
     // reserved
 }
 */
struct ModifyTheBuyNowPriceInAnOfferActionResponse {
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
final class ModifyTheBuyNowPriceInAnOfferActionClient {
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
    ) async throws -> ModifyTheBuyNowPriceInAnOfferActionResponse {
        let meta = ModifyTheBuyNowPriceInAnOfferActionMeta()
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
        return ModifyTheBuyNowPriceInAnOfferActionResponse(
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
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionReq
struct ModifyTheBuyNowPriceInAnOfferActionReq: Codable {
		let id: String
		let input:  ModifyTheBuyNowPriceInAnOfferActionReqInput
}
  // The base class definition for input
struct ModifyTheBuyNowPriceInAnOfferActionReqInput: Codable {
		let buyNowPrice:  ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice
}
  // The base class definition for buyNowPrice
struct ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for modifyTheBuyNowPriceInAnOfferActionRes
struct ModifyTheBuyNowPriceInAnOfferActionRes: Codable {
		let id: String
		let input:  ModifyTheBuyNowPriceInAnOfferActionResInput
		let output:  ModifyTheBuyNowPriceInAnOfferActionResOutput
}
  // The base class definition for input
struct ModifyTheBuyNowPriceInAnOfferActionResInput: Codable {
		let buyNowPrice:  ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice
}
  // The base class definition for buyNowPrice
struct ModifyTheBuyNowPriceInAnOfferActionResInputBuyNowPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for output
struct ModifyTheBuyNowPriceInAnOfferActionResOutput: Codable {
		let status: String
		let errors: [ModifyTheBuyNowPriceInAnOfferActionResOutputErrors]
}
  // The base class definition for errors
struct ModifyTheBuyNowPriceInAnOfferActionResOutputErrors: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata
}
  // The base class definition for metadata
struct ModifyTheBuyNowPriceInAnOfferActionResOutputErrorsMetadata: Codable {
		let productId: String
}