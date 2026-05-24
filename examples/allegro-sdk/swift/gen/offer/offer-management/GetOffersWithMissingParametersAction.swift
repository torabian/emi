import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetOffersWithMissingParametersAction
 */
struct GetOffersWithMissingParametersActionMeta {
    let name: String = "GetOffersWithMissingParametersAction"
    let url: String = "https://api.{environment}/sale/offers/unfilled-parameters"
    let method: String = "Get"
}
/*
 struct GetOffersWithMissingParametersActionRequest {
     // reserved
 }
 */
struct GetOffersWithMissingParametersActionResponse {
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
final class GetOffersWithMissingParametersActionClient {
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
    ) async throws -> GetOffersWithMissingParametersActionResponse {
        let meta = GetOffersWithMissingParametersActionMeta()
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
        return GetOffersWithMissingParametersActionResponse(
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
  // The base class definition for getOffersWithMissingParametersActionRes
struct GetOffersWithMissingParametersActionRes: Codable {
		let offers: [GetOffersWithMissingParametersActionResOffers]
		let count: Int
		let totalCount: Int
}
  // The base class definition for offers
struct GetOffersWithMissingParametersActionResOffers: Codable {
		let id: String
		let parameters: [GetOffersWithMissingParametersActionResOffersParameters]
		let category:  GetOffersWithMissingParametersActionResOffersCategory
}
  // The base class definition for parameters
struct GetOffersWithMissingParametersActionResOffersParameters: Codable {
		let id: String
}
  // The base class definition for category
struct GetOffersWithMissingParametersActionResOffersCategory: Codable {
		let id: String
}