import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetEventsAboutTheSellerSOffersAction
 */
struct GetEventsAboutTheSellerSOffersActionMeta {
    let name: String = "GetEventsAboutTheSellerSOffersAction"
    let url: String = "https://api.{environment}/sale/offer-events"
    let method: String = "Get"
}
/*
 struct GetEventsAboutTheSellerSOffersActionRequest {
     // reserved
 }
 */
struct GetEventsAboutTheSellerSOffersActionResponse {
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
final class GetEventsAboutTheSellerSOffersActionClient {
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
    ) async throws -> GetEventsAboutTheSellerSOffersActionResponse {
        let meta = GetEventsAboutTheSellerSOffersActionMeta()
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
        return GetEventsAboutTheSellerSOffersActionResponse(
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
  // The base class definition for getEventsAboutTheSellerSOffersActionRes
struct GetEventsAboutTheSellerSOffersActionRes: Codable {
		  // List of events related to offer state changes
 let offerEvents: [GetEventsAboutTheSellerSOffersActionResOfferEvents]
}
  // The base class definition for offerEvents
struct GetEventsAboutTheSellerSOffersActionResOfferEvents: Codable {
		  // Unique event identifier (base64 encoded)
 let id: String
		  // ISO8601 timestamp when the event occurred
 let occurredAt: String
		  // Event type (e.g., OFFER_ACTIVATED, OFFER_ENDED, etc.)
 let type: String
		  // Basic offer information for which event occurred
 let offer:  GetEventsAboutTheSellerSOffersActionResOfferEventsOffer
}
  // The base class definition for offer
struct GetEventsAboutTheSellerSOffersActionResOfferEventsOffer: Codable {
		let id: String
		let external:  GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal
}
  // The base class definition for external
struct GetEventsAboutTheSellerSOffersActionResOfferEventsOfferExternal: Codable {
		let id: String
}