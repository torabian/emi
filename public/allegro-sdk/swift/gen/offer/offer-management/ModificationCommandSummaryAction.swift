import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action ModificationCommandSummaryAction
 */
struct ModificationCommandSummaryActionMeta {
    let name: String = "ModificationCommandSummaryAction"
    let url: String = "https://api.{environment}/sale/offers/promo-options-commands/{commandId}"
    let method: String = "Get"
}
/*
 struct ModificationCommandSummaryActionRequest {
     // reserved
 }
 */
struct ModificationCommandSummaryActionResponse {
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
final class ModificationCommandSummaryActionClient {
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
    ) async throws -> ModificationCommandSummaryActionResponse {
        let meta = ModificationCommandSummaryActionMeta()
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
        return ModificationCommandSummaryActionResponse(
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
  // The base class definition for modificationCommandSummaryActionRes
struct ModificationCommandSummaryActionRes: Codable {
		let id: String
		let taskCount:  ModificationCommandSummaryActionResTaskCount
}
  // The base class definition for taskCount
struct ModificationCommandSummaryActionResTaskCount: Codable {
		let failed: Int
		let success: Int
		let total: Int
}