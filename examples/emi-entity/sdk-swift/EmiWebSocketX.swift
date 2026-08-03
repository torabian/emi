/**
 * EmiWebSocketX - generated once per module, backs every "method: reactive" action.
 */
import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

final class EmiWebSocketX<Send: Encodable, Receive: Decodable> {
	private let task: URLSessionWebSocketTask
	private let session: URLSession

	var onMessage: ((Receive) -> Void)?
	var onError: ((Error) -> Void)?
	var onClose: ((URLSessionWebSocketTask.CloseCode) -> Void)?

	init(url: URL, session: URLSession = .shared) {
		self.session = session
		self.task = session.webSocketTask(with: url)
	}

	func connect() {
		task.resume()
		listen()
	}

	private func listen() {
		task.receive { [weak self] result in
			guard let self = self else { return }

			switch result {
			case .failure(let error):
				self.onError?(error)
				self.onClose?(self.task.closeCode)
			case .success(let message):
				switch message {
				case .string(let text):
					self.decodeAndDeliver(Data(text.utf8))
				case .data(let data):
					self.decodeAndDeliver(data)
				@unknown default:
					break
				}
				self.listen()
			}
		}
	}

	private func decodeAndDeliver(_ data: Data) {
		do {
			let decoded = try JSONDecoder().decode(Receive.self, from: data)
			onMessage?(decoded)
		} catch {
			onError?(error)
		}
	}

	func send(_ value: Send, completion: ((Error?) -> Void)? = nil) {
		do {
			let data = try JSONEncoder().encode(value)
			let text = String(data: data, encoding: .utf8) ?? ""
			task.send(.string(text)) { [weak self] error in
				if let error = error {
					self?.onError?(error)
				}
				completion?(error)
			}
		} catch {
			onError?(error)
			completion?(error)
		}
	}

	func close(code: URLSessionWebSocketTask.CloseCode = .normalClosure) {
		task.cancel(with: code, reason: nil)
	}
}
