import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action CreateOfferBasedOnProductAction
 */
struct CreateOfferBasedOnProductActionMeta {
    let name: String = "CreateOfferBasedOnProductAction"
    let url: String = "https://api.{environment}/sale/product-offers"
    let method: String = "Post"
}
/*
 struct CreateOfferBasedOnProductActionRequest {
     // reserved
 }
 */
struct CreateOfferBasedOnProductActionResponse {
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
final class CreateOfferBasedOnProductActionClient {
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
    ) async throws -> CreateOfferBasedOnProductActionResponse {
        let meta = CreateOfferBasedOnProductActionMeta()
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
        return CreateOfferBasedOnProductActionResponse(
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
  // The base class definition for createOfferBasedOnProductActionReq
struct CreateOfferBasedOnProductActionReq: Codable {
		  // Offer title
 let name: String
		  // Offer language code (e.g., pl-PL)
 let language: String
		let category:  CreateOfferBasedOnProductActionReqCategory
		  // Product details and associated quantities
 let productSet: [CreateOfferBasedOnProductActionReqProductSet]
		let stock:  CreateOfferBasedOnProductActionReqStock
		let sellingMode:  CreateOfferBasedOnProductActionReqSellingMode
		let payments:  CreateOfferBasedOnProductActionReqPayments
		let delivery:  CreateOfferBasedOnProductActionReqDelivery
		let publication:  CreateOfferBasedOnProductActionReqPublication
		let additionalMarketplaces: Any
		let compatibilityList:  CreateOfferBasedOnProductActionReqCompatibilityList
		let images: [String]
		let description:  CreateOfferBasedOnProductActionReqDescription
		let b2b:  CreateOfferBasedOnProductActionReqB2b
		let attachments: [CreateOfferBasedOnProductActionReqAttachments]
		let fundraisingCampaign:  CreateOfferBasedOnProductActionReqFundraisingCampaign
		let additionalServices:  CreateOfferBasedOnProductActionReqAdditionalServices
		let afterSalesServices:  CreateOfferBasedOnProductActionReqAfterSalesServices
		let sizeTable:  CreateOfferBasedOnProductActionReqSizeTable
		let contact:  CreateOfferBasedOnProductActionReqContact
		let discounts:  CreateOfferBasedOnProductActionReqDiscounts
		let location:  CreateOfferBasedOnProductActionReqLocation
		let external:  CreateOfferBasedOnProductActionReqExternal
		let taxSettings:  CreateOfferBasedOnProductActionReqTaxSettings
		let messageToSellerSettings:  CreateOfferBasedOnProductActionReqMessageToSellerSettings
}
  // The base class definition for category
struct CreateOfferBasedOnProductActionReqCategory: Codable {
		let id: String
}
  // The base class definition for productSet
struct CreateOfferBasedOnProductActionReqProductSet: Codable {
		let product:  CreateOfferBasedOnProductActionReqProductSetProduct
		let quantity:  CreateOfferBasedOnProductActionReqProductSetQuantity
		let responsiblePerson:  CreateOfferBasedOnProductActionReqProductSetResponsiblePerson
		let responsibleProducer:  CreateOfferBasedOnProductActionReqProductSetResponsibleProducer
		let safetyInformation:  CreateOfferBasedOnProductActionReqProductSetSafetyInformation
		let marketedBeforeGPSRObligation: Boolean
		let deposits: [CreateOfferBasedOnProductActionReqProductSetDeposits]
}
  // The base class definition for product
struct CreateOfferBasedOnProductActionReqProductSetProduct: Codable {
		let id: String
		let idType: String
		let name: String
		let category:  CreateOfferBasedOnProductActionReqProductSetProductCategory
		let parameters: [CreateOfferBasedOnProductActionReqProductSetProductParameters]
		let images: [String]
}
  // The base class definition for category
struct CreateOfferBasedOnProductActionReqProductSetProductCategory: Codable {
		let id: String
}
  // The base class definition for parameters
struct CreateOfferBasedOnProductActionReqProductSetProductParameters: Codable {
		let id: String
		let name: String
		let rangeValue:  CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue
		let values: [String]
		let valuesIds: [String]
}
  // The base class definition for rangeValue
struct CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue: Codable {
		let from: String
		let to: String
}
  // The base class definition for quantity
struct CreateOfferBasedOnProductActionReqProductSetQuantity: Codable {
		let value: Int
}
  // The base class definition for responsiblePerson
struct CreateOfferBasedOnProductActionReqProductSetResponsiblePerson: Codable {
		let id: String
		let name: String
}
  // The base class definition for responsibleProducer
struct CreateOfferBasedOnProductActionReqProductSetResponsibleProducer: Codable {
		let id: String
		let type: String
}
  // The base class definition for safetyInformation
struct CreateOfferBasedOnProductActionReqProductSetSafetyInformation: Codable {
		let type: String
		let description: String
}
  // The base class definition for deposits
struct CreateOfferBasedOnProductActionReqProductSetDeposits: Codable {
		let id: String
		let quantity: Int
}
  // The base class definition for stock
struct CreateOfferBasedOnProductActionReqStock: Codable {
		let available: Int
		let unit: String
}
  // The base class definition for sellingMode
struct CreateOfferBasedOnProductActionReqSellingMode: Codable {
		let format: String
		let price:  CreateOfferBasedOnProductActionReqSellingModePrice
		let minimalPrice:  CreateOfferBasedOnProductActionReqSellingModeMinimalPrice
		let startingPrice:  CreateOfferBasedOnProductActionReqSellingModeStartingPrice
}
  // The base class definition for price
struct CreateOfferBasedOnProductActionReqSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for minimalPrice
struct CreateOfferBasedOnProductActionReqSellingModeMinimalPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for startingPrice
struct CreateOfferBasedOnProductActionReqSellingModeStartingPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for payments
struct CreateOfferBasedOnProductActionReqPayments: Codable {
		let invoice: String
}
  // The base class definition for delivery
struct CreateOfferBasedOnProductActionReqDelivery: Codable {
		let handlingTime: String
		let additionalInfo: String
		let shipmentDate: String
		  // Optional; may be null
 let shippingRates:  CreateOfferBasedOnProductActionReqDeliveryShippingRates
}
  // The base class definition for shippingRates
struct CreateOfferBasedOnProductActionReqDeliveryShippingRates: Codable {
		let id: String
}
  // The base class definition for publication
struct CreateOfferBasedOnProductActionReqPublication: Codable {
		let duration: String
		let startingAt: String
		let endingAt: String
		let status: String
		let republish: Boolean
}
  // The base class definition for compatibilityList
struct CreateOfferBasedOnProductActionReqCompatibilityList: Codable {
		let items: [CreateOfferBasedOnProductActionReqCompatibilityListItems]
}
  // The base class definition for items
struct CreateOfferBasedOnProductActionReqCompatibilityListItems: Codable {
		let type: String
		let text: String
}
  // The base class definition for description
struct CreateOfferBasedOnProductActionReqDescription: Codable {
		let sections: [CreateOfferBasedOnProductActionReqDescriptionSections]
}
  // The base class definition for sections
struct CreateOfferBasedOnProductActionReqDescriptionSections: Codable {
		let items: [CreateOfferBasedOnProductActionReqDescriptionSectionsItems]
}
  // The base class definition for items
struct CreateOfferBasedOnProductActionReqDescriptionSectionsItems: Codable {
		let type: String
}
  // The base class definition for b2b
struct CreateOfferBasedOnProductActionReqB2b: Codable {
		let buyableOnlyByBusiness: Boolean
}
  // The base class definition for attachments
struct CreateOfferBasedOnProductActionReqAttachments: Codable {
		let id: String
}
  // The base class definition for fundraisingCampaign
struct CreateOfferBasedOnProductActionReqFundraisingCampaign: Codable {
		let id: String
		let name: String
}
  // The base class definition for additionalServices
struct CreateOfferBasedOnProductActionReqAdditionalServices: Codable {
		let id: String
		let name: String
}
  // The base class definition for afterSalesServices
struct CreateOfferBasedOnProductActionReqAfterSalesServices: Codable {
		let impliedWarranty:  CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty
		let returnPolicy:  CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy
		let warranty:  CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty
}
  // The base class definition for impliedWarranty
struct CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty: Codable {
		let id: String
		let name: String
}
  // The base class definition for returnPolicy
struct CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy: Codable {
		let id: String
		let name: String
}
  // The base class definition for warranty
struct CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty: Codable {
		let id: String
		let name: String
}
  // The base class definition for sizeTable
struct CreateOfferBasedOnProductActionReqSizeTable: Codable {
		let id: String
		let name: String
}
  // The base class definition for contact
struct CreateOfferBasedOnProductActionReqContact: Codable {
		let id: String
		let name: String
}
  // The base class definition for discounts
struct CreateOfferBasedOnProductActionReqDiscounts: Codable {
		let wholesalePriceList:  CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList
}
  // The base class definition for wholesalePriceList
struct CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList: Codable {
		let id: String
		let name: String
}
  // The base class definition for location
struct CreateOfferBasedOnProductActionReqLocation: Codable {
		let city: String
		let countryCode: String
		let postCode: String
		let province: String
}
  // The base class definition for external
struct CreateOfferBasedOnProductActionReqExternal: Codable {
		let id: String
}
  // The base class definition for taxSettings
struct CreateOfferBasedOnProductActionReqTaxSettings: Codable {
		let subject: String
		let exemption: String
		let rates: [CreateOfferBasedOnProductActionReqTaxSettingsRates]
}
  // The base class definition for rates
struct CreateOfferBasedOnProductActionReqTaxSettingsRates: Codable {
		let rate: String
		let countryCode: String
}
  // The base class definition for messageToSellerSettings
struct CreateOfferBasedOnProductActionReqMessageToSellerSettings: Codable {
		let mode: String
		let hint: String
}