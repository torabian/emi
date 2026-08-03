import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Reactive (WebSocket) action ReactiveEchoAction
 */
struct ReactiveEchoActionMeta {
    let name: String = "ReactiveEchoAction"
    let url: String = "/reactive/echo/:channel"
}
/**
 * Path parameters for ReactiveEchoAction
 */
struct ReactiveEchoActionPathParameter {
	var Channel: String
}
// Converts a placeholder url, and applies the parameters to it.
func ReactiveEchoActionPathParameterApply(_ params: ReactiveEchoActionPathParameter, _ templateUrl: String) -> String {
	var url = templateUrl
	url = url.replacingOccurrences(of: ":channel", with: params.Channel)
	return url
}
typealias ReactiveEchoActionSocket = EmiWebSocketX<ReactiveEchoActionReq, ReactiveEchoActionRes>
enum ReactiveEchoAction {
    private static func webSocketBaseUrl() -> String {
        let base = EmiClientConfig.baseUrl
        if base.hasPrefix("https://") {
            return "wss://" + base.dropFirst("https://".count)
        }
        if base.hasPrefix("http://") {
            return "ws://" + base.dropFirst("http://".count)
        }
        return base
    }
    // Opens a new connection. Nil only if EmiClientConfig.baseUrl (plus path/query) is
    // not a valid URL - call .connect() on the result to actually open the socket.
    static func Create(
        path: ReactiveEchoActionPathParameter,
        query: [String: String] = [:]
    ) -> ReactiveEchoActionSocket? {
        let meta = ReactiveEchoActionMeta()
        guard var components = URLComponents(string: webSocketBaseUrl()) else {
            return nil
        }
        components.path = ReactiveEchoActionPathParameterApply(path, meta.url)
        if !query.isEmpty {
            components.queryItems = query.map {
                URLQueryItem(name: $0.key, value: $0.value)
            }
        }
        guard let url = components.url else {
            return nil
        }
        return ReactiveEchoActionSocket(url: url)
    }
}
  // The base class definition for reactiveEchoActionReq
struct ReactiveEchoActionReq: Codable {
		let message: String
}
  // The base class definition for reactiveEchoActionRes
struct ReactiveEchoActionRes: Codable {
		let message: String
		let echoedAt: String
}