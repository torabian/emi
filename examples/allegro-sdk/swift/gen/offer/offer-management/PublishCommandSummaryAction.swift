import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action PublishCommandSummaryAction
 */
struct PublishCommandSummaryActionMeta {
    let name: String = "PublishCommandSummaryAction"
    let url: String = "https://api.{environment}/sale/offer-publication-commands/{commandId}"
    let method: String = "Get"
}
/*
 struct PublishCommandSummaryActionRequest {
     // reserved
 }
 */
struct PublishCommandSummaryActionResponse {
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
final class PublishCommandSummaryActionClient {
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
    ) async throws -> PublishCommandSummaryActionResponse {
        let meta = PublishCommandSummaryActionMeta()
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
        return PublishCommandSummaryActionResponse(
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
  // The base class definition for publishCommandSummaryActionRes
struct PublishCommandSummaryActionRes: Codable {
		let id: String
		let createdAt: String
		let completedAt: String
		let taskCount:  PublishCommandSummaryActionResTaskCount
}
  // The base class definition for taskCount
struct PublishCommandSummaryActionResTaskCount: Codable {
		let failed: Int
		let success: Int
		let total: Int
}