import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetSmartClassificationReportOfTheParticularOfferAction
 */
struct GetSmartClassificationReportOfTheParticularOfferActionMeta {
    let name: String = "GetSmartClassificationReportOfTheParticularOfferAction"
    let url: String = "https://api.{environment}/sale/offers/{offerId}/smart"
    let method: String = "Get"
}
/*
 struct GetSmartClassificationReportOfTheParticularOfferActionRequest {
     // reserved
 }
 */
struct GetSmartClassificationReportOfTheParticularOfferActionResponse {
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
final class GetSmartClassificationReportOfTheParticularOfferActionClient {
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
    ) async throws -> GetSmartClassificationReportOfTheParticularOfferActionResponse {
        let meta = GetSmartClassificationReportOfTheParticularOfferActionMeta()
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
        return GetSmartClassificationReportOfTheParticularOfferActionResponse(
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
  // The base class definition for getSmartClassificationReportOfTheParticularOfferActionRes
struct GetSmartClassificationReportOfTheParticularOfferActionRes: Codable {
		  // Indicates if offer is queued for reclassification
 let scheduledForReclassification: Boolean
		  // Offer classification status and last change date
 let classification:  GetSmartClassificationReportOfTheParticularOfferActionResClassification
		  // List of smart delivery method identifiers
 let smartDeliveryMethods: [GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods]
		  // List of classification conditions with delivery method checks
 let conditions: [GetSmartClassificationReportOfTheParticularOfferActionResConditions]
}
  // The base class definition for classification
struct GetSmartClassificationReportOfTheParticularOfferActionResClassification: Codable {
		  // Whether the classification conditions are fulfilled
 let fulfilled: Boolean
		  // ISO8601 timestamp of last classification change
 let lastChanged: String
}
  // The base class definition for smartDeliveryMethods
struct GetSmartClassificationReportOfTheParticularOfferActionResSmartDeliveryMethods: Codable {
		let id: String
}
  // The base class definition for conditions
struct GetSmartClassificationReportOfTheParticularOfferActionResConditions: Codable {
		  // Condition code identifier
 let code: String
		  // Human-readable condition name
 let name: String
		  // Detailed condition description
 let description: String
		  // Indicates if this condition is fulfilled
 let fulfilled: Boolean
		  // Delivery methods that passed validation for this condition
 let passedDeliveryMethods: [GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods]
		  // Delivery methods that failed validation for this condition
 let failedDeliveryMethods: [GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods]
}
  // The base class definition for passedDeliveryMethods
struct GetSmartClassificationReportOfTheParticularOfferActionResConditionsPassedDeliveryMethods: Codable {
		let id: String
}
  // The base class definition for failedDeliveryMethods
struct GetSmartClassificationReportOfTheParticularOfferActionResConditionsFailedDeliveryMethods: Codable {
		let id: String
}