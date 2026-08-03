import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action Entity4UpdateAction
 */
struct Entity4UpdateActionMeta {
    let name: String = "Entity4UpdateAction"
    let url: String = "/entity4/:uniqueId"
    let method: String = "PATCH"
}
/*
 struct Entity4UpdateActionRequest {
     // reserved
 }
 */
struct Entity4UpdateActionResponse {
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
 * Path parameters for Entity4UpdateAction
 */
struct Entity4UpdateActionPathParameter {
	var UniqueId: String
}
// Converts a placeholder url, and applies the parameters to it.
func Entity4UpdateActionPathParameterApply(_ params: Entity4UpdateActionPathParameter, _ templateUrl: String) -> String {
	var url = templateUrl
	url = url.replacingOccurrences(of: ":uniqueId", with: params.UniqueId)
	return url
}
final class Entity4UpdateActionClient {
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
        path: Entity4UpdateActionPathParameter,
        query: [String: String] = [:],
        headers: [String: String] = [:],
        body: Data? = nil
    ) async throws -> Entity4UpdateActionResponse {
        let meta = Entity4UpdateActionMeta()
        let baseUrl = EmiClientConfig.baseUrl
        let resolvedPath = Entity4UpdateActionPathParameterApply(path, meta.url)
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
        return Entity4UpdateActionResponse(
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