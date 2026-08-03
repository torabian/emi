import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action Entity4AwareDeletePreviewAction
 */
struct Entity4AwareDeletePreviewActionMeta {
    let name: String = "Entity4AwareDeletePreviewAction"
    let url: String = "/entity4/delete-preview"
    let method: String = "GET"
}
/*
 struct Entity4AwareDeletePreviewActionRequest {
     // reserved
 }
 */
struct Entity4AwareDeletePreviewActionResponse {
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
final class Entity4AwareDeletePreviewActionClient {
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
        if !query.isEmpty {
            components.queryItems = query.map {
                URLQueryItem(name: $0.key, value: $0.value)
            }
        }
        return components.url
    }
    static func compute(
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> Entity4AwareDeletePreviewActionResponse {
        let meta = Entity4AwareDeletePreviewActionMeta()
        let baseUrl = EmiClientConfig.baseUrl
        let resolvedPath = meta.url
        guard let url = buildUrl(
            base: baseUrl,
            path: resolvedPath,
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
        return Entity4AwareDeletePreviewActionResponse(
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
  // The base class definition for entity4AwareDeletePreviewActionRes
struct Entity4AwareDeletePreviewActionRes: Codable {
		let message: String
		let affected: [Entity4AwareDeletePreviewActionResAffected]
}
  // The base class definition for affected
struct Entity4AwareDeletePreviewActionResAffected: Codable {
		let relation: String
		let count: Int64
}