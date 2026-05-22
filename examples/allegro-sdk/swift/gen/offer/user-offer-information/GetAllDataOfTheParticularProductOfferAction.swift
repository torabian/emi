import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action GetAllDataOfTheParticularProductOfferAction
 */
struct GetAllDataOfTheParticularProductOfferActionMeta {
    let name: String = "GetAllDataOfTheParticularProductOfferAction"
    let url: String = "https://api.{environment}/sale/product-offers/{offerId}"
    let method: String = "Get"
}
/*
 struct GetAllDataOfTheParticularProductOfferActionRequest {
     // reserved
 }
 */
struct GetAllDataOfTheParticularProductOfferActionResponse {
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
final class GetAllDataOfTheParticularProductOfferActionClient {
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
    ) async throws -> GetAllDataOfTheParticularProductOfferActionResponse {
        let meta = GetAllDataOfTheParticularProductOfferActionMeta()
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
        return GetAllDataOfTheParticularProductOfferActionResponse(
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
  // The base class definition for getAllDataOfTheParticularProductOfferActionRes
struct GetAllDataOfTheParticularProductOfferActionRes: Codable {
		  // Unique offer identifier
 let id: String
		  // Offer title
 let name: String
		  // Offer language code (e.g. pl-PL)
 let language: String
		  // Offer creation timestamp (ISO8601)
 let createdAt: String
		  // Offer last update timestamp (ISO8601)
 let updatedAt: String
		let category:  GetAllDataOfTheParticularProductOfferActionResCategory
		let stock:  GetAllDataOfTheParticularProductOfferActionResStock
		let contact:  GetAllDataOfTheParticularProductOfferActionResContact
		let publication:  GetAllDataOfTheParticularProductOfferActionResPublication
		let sellingMode:  GetAllDataOfTheParticularProductOfferActionResSellingMode
		let payments:  GetAllDataOfTheParticularProductOfferActionResPayments
		let delivery:  GetAllDataOfTheParticularProductOfferActionResDelivery
		let afterSalesServices:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServices
		let discounts:  GetAllDataOfTheParticularProductOfferActionResDiscounts
		let description:  GetAllDataOfTheParticularProductOfferActionResDescription
		let images: [String]
		let productSet: [GetAllDataOfTheParticularProductOfferActionResProductSet]
		let attachments: [GetAllDataOfTheParticularProductOfferActionResAttachments]
		let fundraisingCampaign:  GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign
		let additionalServices:  GetAllDataOfTheParticularProductOfferActionResAdditionalServices
		let additionalMarketplaces: Any
		let b2b:  GetAllDataOfTheParticularProductOfferActionResB2b
		let compatibilityList:  GetAllDataOfTheParticularProductOfferActionResCompatibilityList
		let validation:  GetAllDataOfTheParticularProductOfferActionResValidation
		let external:  GetAllDataOfTheParticularProductOfferActionResExternal
		let sizeTable:  GetAllDataOfTheParticularProductOfferActionResSizeTable
		let taxSettings:  GetAllDataOfTheParticularProductOfferActionResTaxSettings
		let messageToSellerSettings:  GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings
}
  // The base class definition for category
struct GetAllDataOfTheParticularProductOfferActionResCategory: Codable {
		let id: String
}
  // The base class definition for stock
struct GetAllDataOfTheParticularProductOfferActionResStock: Codable {
		let available: Int
		let unit: String
}
  // The base class definition for contact
struct GetAllDataOfTheParticularProductOfferActionResContact: Codable {
		let id: String
}
  // The base class definition for publication
struct GetAllDataOfTheParticularProductOfferActionResPublication: Codable {
		let duration: String
		let startingAt: String
		let endingAt: String
		let endedBy: String
		let status: String
		let republish: Boolean
		let marketplaces:  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces
}
  // The base class definition for marketplaces
struct GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces: Codable {
		let base:  GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase
		let additional: [GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional]
}
  // The base class definition for base
struct GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase: Codable {
		let id: String
}
  // The base class definition for additional
struct GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional: Codable {
		let id: String
}
  // The base class definition for sellingMode
struct GetAllDataOfTheParticularProductOfferActionResSellingMode: Codable {
		let format: String
		let price:  GetAllDataOfTheParticularProductOfferActionResSellingModePrice
		let minimalPrice:  GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice
		let startingPrice:  GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice
}
  // The base class definition for price
struct GetAllDataOfTheParticularProductOfferActionResSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for minimalPrice
struct GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for startingPrice
struct GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for payments
struct GetAllDataOfTheParticularProductOfferActionResPayments: Codable {
		let invoice: String
}
  // The base class definition for delivery
struct GetAllDataOfTheParticularProductOfferActionResDelivery: Codable {
		let handlingTime: String
		let additionalInfo: String
		let shipmentDate: String
		let shippingRates:  GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates
}
  // The base class definition for shippingRates
struct GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates: Codable {
		let id: String
}
  // The base class definition for afterSalesServices
struct GetAllDataOfTheParticularProductOfferActionResAfterSalesServices: Codable {
		let impliedWarranty:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty
		let returnPolicy:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy
		let warranty:  GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty
}
  // The base class definition for impliedWarranty
struct GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty: Codable {
		let id: String
}
  // The base class definition for returnPolicy
struct GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy: Codable {
		let id: String
}
  // The base class definition for warranty
struct GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty: Codable {
		let id: String
}
  // The base class definition for discounts
struct GetAllDataOfTheParticularProductOfferActionResDiscounts: Codable {
		let wholesalePriceList:  GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList
}
  // The base class definition for wholesalePriceList
struct GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList: Codable {
		let id: String
}
  // The base class definition for description
struct GetAllDataOfTheParticularProductOfferActionResDescription: Codable {
		let sections: [GetAllDataOfTheParticularProductOfferActionResDescriptionSections]
}
  // The base class definition for sections
struct GetAllDataOfTheParticularProductOfferActionResDescriptionSections: Codable {
		let items: [GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems]
}
  // The base class definition for items
struct GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems: Codable {
		let type: String
}
  // The base class definition for productSet
struct GetAllDataOfTheParticularProductOfferActionResProductSet: Codable {
		let quantity:  GetAllDataOfTheParticularProductOfferActionResProductSetQuantity
		let product:  GetAllDataOfTheParticularProductOfferActionResProductSetProduct
		let responsiblePerson:  GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson
		let responsibleProducer:  GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer
		let safetyInformation:  GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation
		let marketedBeforeGPSRObligation: Boolean
		let deposits: [GetAllDataOfTheParticularProductOfferActionResProductSetDeposits]
}
  // The base class definition for quantity
struct GetAllDataOfTheParticularProductOfferActionResProductSetQuantity: Codable {
		let value: Int
}
  // The base class definition for product
struct GetAllDataOfTheParticularProductOfferActionResProductSetProduct: Codable {
		let id: String
		let isAiCoCreated: Boolean
		let publication:  GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication
		let parameters: [GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters]
}
  // The base class definition for publication
struct GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication: Codable {
		let status: String
}
  // The base class definition for parameters
struct GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters: Codable {
		let id: String
		let name: String
		let rangeValue:  GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue
		let values: [GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues]
		let valuesIds: [GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds]
}
  // The base class definition for rangeValue
struct GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue: Codable {
		let from: String
		let to: String
}
  // The base class definition for responsiblePerson
struct GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson: Codable {
		let id: String
}
  // The base class definition for responsibleProducer
struct GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer: Codable {
		let id: String
}
  // The base class definition for safetyInformation
struct GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation: Codable {
		let type: String
		let description: String
}
  // The base class definition for deposits
struct GetAllDataOfTheParticularProductOfferActionResProductSetDeposits: Codable {
		let id: String
		let quantity: Int
}
  // The base class definition for attachments
struct GetAllDataOfTheParticularProductOfferActionResAttachments: Codable {
		let id: String
}
  // The base class definition for fundraisingCampaign
struct GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign: Codable {
		let id: String
}
  // The base class definition for additionalServices
struct GetAllDataOfTheParticularProductOfferActionResAdditionalServices: Codable {
		let id: String
}
  // The base class definition for b2b
struct GetAllDataOfTheParticularProductOfferActionResB2b: Codable {
		let buyableOnlyByBusiness: Boolean
}
  // The base class definition for compatibilityList
struct GetAllDataOfTheParticularProductOfferActionResCompatibilityList: Codable {
		let type: String
}
  // The base class definition for validation
struct GetAllDataOfTheParticularProductOfferActionResValidation: Codable {
		let validatedAt: String
		let errors: [GetAllDataOfTheParticularProductOfferActionResValidationErrors]
		let warnings: [GetAllDataOfTheParticularProductOfferActionResValidationWarnings]
}
  // The base class definition for errors
struct GetAllDataOfTheParticularProductOfferActionResValidationErrors: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata
}
  // The base class definition for metadata
struct GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata: Codable {
		let productId: String
}
  // The base class definition for warnings
struct GetAllDataOfTheParticularProductOfferActionResValidationWarnings: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata
}
  // The base class definition for metadata
struct GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata: Codable {
		let productId: String
}
  // The base class definition for external
struct GetAllDataOfTheParticularProductOfferActionResExternal: Codable {
		let id: String
}
  // The base class definition for sizeTable
struct GetAllDataOfTheParticularProductOfferActionResSizeTable: Codable {
		let id: String
}
  // The base class definition for taxSettings
struct GetAllDataOfTheParticularProductOfferActionResTaxSettings: Codable {
		let subject: String
		let exemption: String
		let rates: [GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates]
}
  // The base class definition for rates
struct GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates: Codable {
		let rate: String
		let countryCode: String
}
  // The base class definition for messageToSellerSettings
struct GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings: Codable {
		let mode: String
		let hint: String
}