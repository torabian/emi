import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action Entity4GetAction
 */
struct Entity4GetActionMeta {
    let name: String = "Entity4GetAction"
    let url: String = "/entity4/:uniqueId"
    let method: String = "GET"
}
/*
 struct Entity4GetActionRequest {
     // reserved
 }
 */
struct Entity4GetActionResponse {
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
 * Path parameters for Entity4GetAction
 */
struct Entity4GetActionPathParameter {
	var UniqueId: String
}
// Converts a placeholder url, and applies the parameters to it.
func Entity4GetActionPathParameterApply(_ params: Entity4GetActionPathParameter, _ templateUrl: String) -> String {
	var url = templateUrl
	url = url.replacingOccurrences(of: ":uniqueId", with: params.UniqueId)
	return url
}
final class Entity4GetActionClient {
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
        path: Entity4GetActionPathParameter,
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> Entity4GetActionResponse {
        let meta = Entity4GetActionMeta()
        let baseUrl = EmiClientConfig.baseUrl
        let resolvedPath = Entity4GetActionPathParameterApply(path, meta.url)
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
        return Entity4GetActionResponse(
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