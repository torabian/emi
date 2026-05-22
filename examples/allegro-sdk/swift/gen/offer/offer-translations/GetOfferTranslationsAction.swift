import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetOfferTranslationsAction
 */
struct GetOfferTranslationsActionMeta {
    let name: String = "GetOfferTranslationsAction"
    let url: String = "https://api.{environment}/sale/offers/{offerId}/translations"
    let method: String = "Get"
}
/*
 struct GetOfferTranslationsActionRequest {
     // reserved
 }
 */
struct GetOfferTranslationsActionResponse {
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
final class GetOfferTranslationsActionClient {
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
    ) async throws -> GetOfferTranslationsActionResponse {
        let meta = GetOfferTranslationsActionMeta()
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
        return GetOfferTranslationsActionResponse(
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
  // The base class definition for getOfferTranslationsActionRes
struct GetOfferTranslationsActionRes: Codable {
		let translations: [GetOfferTranslationsActionResTranslations]
}
  // The base class definition for translations
struct GetOfferTranslationsActionResTranslations: Codable {
		let language: String
		let title:  GetOfferTranslationsActionResTranslationsTitle
		let description:  GetOfferTranslationsActionResTranslationsDescription
		let safetyInformation:  GetOfferTranslationsActionResTranslationsSafetyInformation
}
  // The base class definition for title
struct GetOfferTranslationsActionResTranslationsTitle: Codable {
		let translation: String
		let type: String
}
  // The base class definition for description
struct GetOfferTranslationsActionResTranslationsDescription: Codable {
		let translation:  GetOfferTranslationsActionResTranslationsDescriptionTranslation
		let type: String
}
  // The base class definition for translation
struct GetOfferTranslationsActionResTranslationsDescriptionTranslation: Codable {
		let sections: [GetOfferTranslationsActionResTranslationsDescriptionTranslationSections]
}
  // The base class definition for sections
struct GetOfferTranslationsActionResTranslationsDescriptionTranslationSections: Codable {
		let items: [GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems]
}
  // The base class definition for items
struct GetOfferTranslationsActionResTranslationsDescriptionTranslationSectionsItems: Codable {
		let type: String
}
  // The base class definition for safetyInformation
struct GetOfferTranslationsActionResTranslationsSafetyInformation: Codable {
		let products: [GetOfferTranslationsActionResTranslationsSafetyInformationProducts]
}
  // The base class definition for products
struct GetOfferTranslationsActionResTranslationsSafetyInformationProducts: Codable {
		let id: String
		let translation: String
		let type: String
}