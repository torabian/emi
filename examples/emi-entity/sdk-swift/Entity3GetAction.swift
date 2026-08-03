import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action Entity3GetAction
 */
struct Entity3GetActionMeta {
    let name: String = "Entity3GetAction"
    let url: String = "/entity3/:uniqueId"
    let method: String = "GET"
}
/*
 struct Entity3GetActionRequest {
     // reserved
 }
 */
struct Entity3GetActionResponse {
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
/**
 * Path parameters for Entity3GetAction
 */
struct Entity3GetActionPathParameter {
	var UniqueId: String
}
// Converts a placeholder url, and applies the parameters to it.
func Entity3GetActionPathParameterApply(_ params: Entity3GetActionPathParameter, _ templateUrl: String) -> String {
	var url = templateUrl
	url = url.replacingOccurrences(of: ":uniqueId", with: params.UniqueId)
	return url
}
final class Entity3GetActionClient {
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
        path: Entity3GetActionPathParameter,
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> Entity3GetActionResponse {
        let meta = Entity3GetActionMeta()
        let baseUrl = EmiClientConfig.baseUrl
        let resolvedPath = Entity3GetActionPathParameterApply(path, meta.url)
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
        return Entity3GetActionResponse(
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