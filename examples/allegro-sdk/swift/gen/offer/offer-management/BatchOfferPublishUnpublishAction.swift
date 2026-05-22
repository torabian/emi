import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action BatchOfferPublishUnpublishAction
 */
struct BatchOfferPublishUnpublishActionMeta {
    let name: String = "BatchOfferPublishUnpublishAction"
    let url: String = "https://api.{environment}/sale/offer-publication-commands/{commandId}"
    let method: String = "Put"
}
/*
 struct BatchOfferPublishUnpublishActionRequest {
     // reserved
 }
 */
struct BatchOfferPublishUnpublishActionResponse {
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
final class BatchOfferPublishUnpublishActionClient {
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
    ) async throws -> BatchOfferPublishUnpublishActionResponse {
        let meta = BatchOfferPublishUnpublishActionMeta()
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
        return BatchOfferPublishUnpublishActionResponse(
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
  // The base class definition for batchOfferPublishUnpublishActionReq
struct BatchOfferPublishUnpublishActionReq: Codable {
		let offerCriteria: [BatchOfferPublishUnpublishActionReqOfferCriteria]
		let publication:  BatchOfferPublishUnpublishActionReqPublication
}
  // The base class definition for offerCriteria
struct BatchOfferPublishUnpublishActionReqOfferCriteria: Codable {
		let offers: [BatchOfferPublishUnpublishActionReqOfferCriteriaOffers]
		let type: String
}
  // The base class definition for offers
struct BatchOfferPublishUnpublishActionReqOfferCriteriaOffers: Codable {
		let id: String
}
  // The base class definition for publication
struct BatchOfferPublishUnpublishActionReqPublication: Codable {
		let action: String
		let scheduledFor: String
}
  // The base class definition for batchOfferPublishUnpublishActionRes
struct BatchOfferPublishUnpublishActionRes: Codable {
		let id: String
		let createdAt: String
		let completedAt: String
		let taskCount:  BatchOfferPublishUnpublishActionResTaskCount
}
  // The base class definition for taskCount
struct BatchOfferPublishUnpublishActionResTaskCount: Codable {
		let failed: Int
		let success: Int
		let total: Int
}