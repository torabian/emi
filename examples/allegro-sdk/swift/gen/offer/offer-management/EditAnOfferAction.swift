import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif
/**
 Action to communicate with the action EditAnOfferAction
 */
struct EditAnOfferActionMeta {
    let name: String = "EditAnOfferAction"
    let url: String = "https://api.{environment}/sale/product-offers/{offerId}"
    let method: String = "Patch"
}
/*
 struct EditAnOfferActionRequest {
     // reserved
 }
 */
struct EditAnOfferActionResponse {
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
final class EditAnOfferActionClient {
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
    ) async throws -> EditAnOfferActionResponse {
        let meta = EditAnOfferActionMeta()
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
        return EditAnOfferActionResponse(
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
  // The base class definition for editAnOfferActionReq
struct EditAnOfferActionReq: Codable {
		let name: String
		let language: String
		let category:  EditAnOfferActionReqCategory
		let productSet: [EditAnOfferActionReqProductSet]
		let stock:  EditAnOfferActionReqStock
		let sellingMode:  EditAnOfferActionReqSellingMode
		let payments:  EditAnOfferActionReqPayments
		let delivery:  EditAnOfferActionReqDelivery
		let publication:  EditAnOfferActionReqPublication
		let additionalMarketplaces: Any
		let compatibilityList:  EditAnOfferActionReqCompatibilityList
		let images: [String]
		let description:  EditAnOfferActionReqDescription
		let b2b:  EditAnOfferActionReqB2b
		let attachments: [EditAnOfferActionReqAttachments]
		let fundraisingCampaign:  EditAnOfferActionReqFundraisingCampaign
		let additionalServices:  EditAnOfferActionReqAdditionalServices
		let afterSalesServices:  EditAnOfferActionReqAfterSalesServices
		let sizeTable:  EditAnOfferActionReqSizeTable
		let contact:  EditAnOfferActionReqContact
		let discounts:  EditAnOfferActionReqDiscounts
		let location:  EditAnOfferActionReqLocation
		let external:  EditAnOfferActionReqExternal
		let taxSettings:  EditAnOfferActionReqTaxSettings
		let messageToSellerSettings:  EditAnOfferActionReqMessageToSellerSettings
}
  // The base class definition for category
struct EditAnOfferActionReqCategory: Codable {
		let id: String
}
  // The base class definition for productSet
struct EditAnOfferActionReqProductSet: Codable {
		let product:  EditAnOfferActionReqProductSetProduct
		let quantity:  EditAnOfferActionReqProductSetQuantity
		let responsiblePerson:  EditAnOfferActionReqProductSetResponsiblePerson
		let responsibleProducer:  EditAnOfferActionReqProductSetResponsibleProducer
		let safetyInformation:  EditAnOfferActionReqProductSetSafetyInformation
		let marketedBeforeGPSRObligation: Boolean
		let deposits: [EditAnOfferActionReqProductSetDeposits]
}
  // The base class definition for product
struct EditAnOfferActionReqProductSetProduct: Codable {
		let id: String
		let idType: String
		let name: String
		let category:  EditAnOfferActionReqProductSetProductCategory
		let parameters: [EditAnOfferActionReqProductSetProductParameters]
		let images: [String]
}
  // The base class definition for category
struct EditAnOfferActionReqProductSetProductCategory: Codable {
		let id: String
}
  // The base class definition for parameters
struct EditAnOfferActionReqProductSetProductParameters: Codable {
		let id: String
		let name: String
		let rangeValue:  EditAnOfferActionReqProductSetProductParametersRangeValue
		let values: [String]
		let valuesIds: [String]
}
  // The base class definition for rangeValue
struct EditAnOfferActionReqProductSetProductParametersRangeValue: Codable {
		let from: String
		let to: String
}
  // The base class definition for quantity
struct EditAnOfferActionReqProductSetQuantity: Codable {
		let value: Int
}
  // The base class definition for responsiblePerson
struct EditAnOfferActionReqProductSetResponsiblePerson: Codable {
		let id: String
		let name: String
}
  // The base class definition for responsibleProducer
struct EditAnOfferActionReqProductSetResponsibleProducer: Codable {
		let id: String
		let type: String
}
  // The base class definition for safetyInformation
struct EditAnOfferActionReqProductSetSafetyInformation: Codable {
		let type: String
		let description: String
}
  // The base class definition for deposits
struct EditAnOfferActionReqProductSetDeposits: Codable {
		let id: String
		let quantity: Int
}
  // The base class definition for stock
struct EditAnOfferActionReqStock: Codable {
		let available: Int
		let unit: String
}
  // The base class definition for sellingMode
struct EditAnOfferActionReqSellingMode: Codable {
		let format: String
		let price:  EditAnOfferActionReqSellingModePrice
		let minimalPrice:  EditAnOfferActionReqSellingModeMinimalPrice
		let startingPrice:  EditAnOfferActionReqSellingModeStartingPrice
}
  // The base class definition for price
struct EditAnOfferActionReqSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for minimalPrice
struct EditAnOfferActionReqSellingModeMinimalPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for startingPrice
struct EditAnOfferActionReqSellingModeStartingPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for payments
struct EditAnOfferActionReqPayments: Codable {
		let invoice: String
}
  // The base class definition for delivery
struct EditAnOfferActionReqDelivery: Codable {
		let handlingTime: String
		let additionalInfo: String
		let shipmentDate: String
		let shippingRates:  EditAnOfferActionReqDeliveryShippingRates
}
  // The base class definition for shippingRates
struct EditAnOfferActionReqDeliveryShippingRates: Codable {
		let id: String
}
  // The base class definition for publication
struct EditAnOfferActionReqPublication: Codable {
		let duration: String
		let startingAt: String
		let endingAt: String
		let status: String
		let republish: Boolean
}
  // The base class definition for compatibilityList
struct EditAnOfferActionReqCompatibilityList: Codable {
		let items: [EditAnOfferActionReqCompatibilityListItems]
}
  // The base class definition for items
struct EditAnOfferActionReqCompatibilityListItems: Codable {
		let type: String
		let text: String
}
  // The base class definition for description
struct EditAnOfferActionReqDescription: Codable {
		let sections: [EditAnOfferActionReqDescriptionSections]
}
  // The base class definition for sections
struct EditAnOfferActionReqDescriptionSections: Codable {
		let items: [EditAnOfferActionReqDescriptionSectionsItems]
}
  // The base class definition for items
struct EditAnOfferActionReqDescriptionSectionsItems: Codable {
		let type: String
}
  // The base class definition for b2b
struct EditAnOfferActionReqB2b: Codable {
		let buyableOnlyByBusiness: Boolean
}
  // The base class definition for attachments
struct EditAnOfferActionReqAttachments: Codable {
		let id: String
}
  // The base class definition for fundraisingCampaign
struct EditAnOfferActionReqFundraisingCampaign: Codable {
		let id: String
		let name: String
}
  // The base class definition for additionalServices
struct EditAnOfferActionReqAdditionalServices: Codable {
		let id: String
		let name: String
}
  // The base class definition for afterSalesServices
struct EditAnOfferActionReqAfterSalesServices: Codable {
		let impliedWarranty:  EditAnOfferActionReqAfterSalesServicesImpliedWarranty
		let returnPolicy:  EditAnOfferActionReqAfterSalesServicesReturnPolicy
		let warranty:  EditAnOfferActionReqAfterSalesServicesWarranty
}
  // The base class definition for impliedWarranty
struct EditAnOfferActionReqAfterSalesServicesImpliedWarranty: Codable {
		let id: String
		let name: String
}
  // The base class definition for returnPolicy
struct EditAnOfferActionReqAfterSalesServicesReturnPolicy: Codable {
		let id: String
		let name: String
}
  // The base class definition for warranty
struct EditAnOfferActionReqAfterSalesServicesWarranty: Codable {
		let id: String
		let name: String
}
  // The base class definition for sizeTable
struct EditAnOfferActionReqSizeTable: Codable {
		let id: String
		let name: String
}
  // The base class definition for contact
struct EditAnOfferActionReqContact: Codable {
		let id: String
		let name: String
}
  // The base class definition for discounts
struct EditAnOfferActionReqDiscounts: Codable {
		let wholesalePriceList:  EditAnOfferActionReqDiscountsWholesalePriceList
}
  // The base class definition for wholesalePriceList
struct EditAnOfferActionReqDiscountsWholesalePriceList: Codable {
		let id: String
		let name: String
}
  // The base class definition for location
struct EditAnOfferActionReqLocation: Codable {
		let city: String
		let countryCode: String
		let postCode: String
		let province: String
}
  // The base class definition for external
struct EditAnOfferActionReqExternal: Codable {
		let id: String
}
  // The base class definition for taxSettings
struct EditAnOfferActionReqTaxSettings: Codable {
		let subject: String
		let exemption: String
		let rates: [EditAnOfferActionReqTaxSettingsRates]
}
  // The base class definition for rates
struct EditAnOfferActionReqTaxSettingsRates: Codable {
		let rate: String
		let countryCode: String
}
  // The base class definition for messageToSellerSettings
struct EditAnOfferActionReqMessageToSellerSettings: Codable {
		let mode: String
		let hint: String
}
  // The base class definition for editAnOfferActionRes
struct EditAnOfferActionRes: Codable {
		let id: String
		let name: String
		let language: String
		let category:  EditAnOfferActionResCategory
		let productSet: [EditAnOfferActionResProductSet]
		let stock:  EditAnOfferActionResStock
		let payments:  EditAnOfferActionResPayments
		let sellingMode:  EditAnOfferActionResSellingMode
		let delivery:  EditAnOfferActionResDelivery
		let publication:  EditAnOfferActionResPublication
		let additionalMarketplaces:  EditAnOfferActionResAdditionalMarketplaces
		let b2b:  EditAnOfferActionResB2b
		let compatibilityList:  EditAnOfferActionResCompatibilityList
		let validation:  EditAnOfferActionResValidation
		let warnings: [String]
		let afterSalesServices:  EditAnOfferActionResAfterSalesServices
		let discounts:  EditAnOfferActionResDiscounts
		let contact:  EditAnOfferActionResContact
		let attachments: [EditAnOfferActionResAttachments]
		let fundraisingCampaign:  EditAnOfferActionResFundraisingCampaign
		let additionalServices:  EditAnOfferActionResAdditionalServices
		let sizeTable:  EditAnOfferActionResSizeTable
		let location:  EditAnOfferActionResLocation
		let external:  EditAnOfferActionResExternal
		let taxSettings:  EditAnOfferActionResTaxSettings
		let messageToSellerSettings:  EditAnOfferActionResMessageToSellerSettings
		let createdAt: String
		let updatedAt: String
		let images: [String]
		let description:  EditAnOfferActionResDescription
}
  // The base class definition for category
struct EditAnOfferActionResCategory: Codable {
		let id: String
}
  // The base class definition for productSet
struct EditAnOfferActionResProductSet: Codable {
		let quantity:  EditAnOfferActionResProductSetQuantity
		let product:  EditAnOfferActionResProductSetProduct
		let responsiblePerson:  EditAnOfferActionResProductSetResponsiblePerson
		let responsibleProducer:  EditAnOfferActionResProductSetResponsibleProducer
		let safetyInformation:  EditAnOfferActionResProductSetSafetyInformation
		let marketedBeforeGPSRObligation: Boolean
		let deposits: [EditAnOfferActionResProductSetDeposits]
}
  // The base class definition for quantity
struct EditAnOfferActionResProductSetQuantity: Codable {
		let value: Int
}
  // The base class definition for product
struct EditAnOfferActionResProductSetProduct: Codable {
		let id: String
		let isAiCoCreated: Boolean
		let publication:  EditAnOfferActionResProductSetProductPublication
		let parameters: [EditAnOfferActionResProductSetProductParameters]
}
  // The base class definition for publication
struct EditAnOfferActionResProductSetProductPublication: Codable {
		let status: String
}
  // The base class definition for parameters
struct EditAnOfferActionResProductSetProductParameters: Codable {
		let id: String
		let name: String
		let rangeValue:  EditAnOfferActionResProductSetProductParametersRangeValue
		let values: [String]
		let valuesIds: [String]
}
  // The base class definition for rangeValue
struct EditAnOfferActionResProductSetProductParametersRangeValue: Codable {
		let from: String
		let to: String
}
  // The base class definition for responsiblePerson
struct EditAnOfferActionResProductSetResponsiblePerson: Codable {
		let id: String
}
  // The base class definition for responsibleProducer
struct EditAnOfferActionResProductSetResponsibleProducer: Codable {
		let id: String
}
  // The base class definition for safetyInformation
struct EditAnOfferActionResProductSetSafetyInformation: Codable {
		let type: String
		let description: String
}
  // The base class definition for deposits
struct EditAnOfferActionResProductSetDeposits: Codable {
		let id: String
		let quantity: Int
}
  // The base class definition for stock
struct EditAnOfferActionResStock: Codable {
		let available: Int
		let unit: String
}
  // The base class definition for payments
struct EditAnOfferActionResPayments: Codable {
		let invoice: String
}
  // The base class definition for sellingMode
struct EditAnOfferActionResSellingMode: Codable {
		let format: String
		let price:  EditAnOfferActionResSellingModePrice
		let minimalPrice:  EditAnOfferActionResSellingModeMinimalPrice
		let startingPrice:  EditAnOfferActionResSellingModeStartingPrice
}
  // The base class definition for price
struct EditAnOfferActionResSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for minimalPrice
struct EditAnOfferActionResSellingModeMinimalPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for startingPrice
struct EditAnOfferActionResSellingModeStartingPrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for delivery
struct EditAnOfferActionResDelivery: Codable {
		let handlingTime: String
		let additionalInfo: String
		let shipmentDate: String
		let shippingRates:  EditAnOfferActionResDeliveryShippingRates
}
  // The base class definition for shippingRates
struct EditAnOfferActionResDeliveryShippingRates: Codable {
		let id: String
}
  // The base class definition for publication
struct EditAnOfferActionResPublication: Codable {
		let duration: String
		let startingAt: String
		let endingAt: String
		let endedBy: String
		let status: String
		let republish: Boolean
		let marketplaces:  EditAnOfferActionResPublicationMarketplaces
}
  // The base class definition for marketplaces
struct EditAnOfferActionResPublicationMarketplaces: Codable {
		let base:  EditAnOfferActionResPublicationMarketplacesBase
		let additional: [EditAnOfferActionResPublicationMarketplacesAdditional]
}
  // The base class definition for base
struct EditAnOfferActionResPublicationMarketplacesBase: Codable {
		let id: String
}
  // The base class definition for additional
struct EditAnOfferActionResPublicationMarketplacesAdditional: Codable {
		let id: String
}
  // The base class definition for additionalMarketplaces
struct EditAnOfferActionResAdditionalMarketplaces: Codable {
		let sellingMode:  EditAnOfferActionResAdditionalMarketplacesSellingMode
		let publication:  EditAnOfferActionResAdditionalMarketplacesPublication
}
  // The base class definition for sellingMode
struct EditAnOfferActionResAdditionalMarketplacesSellingMode: Codable {
		let price:  EditAnOfferActionResAdditionalMarketplacesSellingModePrice
}
  // The base class definition for price
struct EditAnOfferActionResAdditionalMarketplacesSellingModePrice: Codable {
		let amount: String
		let currency: String
}
  // The base class definition for publication
struct EditAnOfferActionResAdditionalMarketplacesPublication: Codable {
		let state: String
		let refusalReasons: [EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons]
}
  // The base class definition for refusalReasons
struct EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons: Codable {
		let code: String
		let userMessage: String
		let parameters:  EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters
}
  // The base class definition for parameters
struct EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters: Codable {
		let maxAllowedPriceDecreasePercent: [String]
}
  // The base class definition for b2b
struct EditAnOfferActionResB2b: Codable {
		let buyableOnlyByBusiness: Boolean
}
  // The base class definition for compatibilityList
struct EditAnOfferActionResCompatibilityList: Codable {
		let type: String
}
  // The base class definition for validation
struct EditAnOfferActionResValidation: Codable {
		let validatedAt: String
		let errors: [EditAnOfferActionResValidationErrors]
		let warnings: [EditAnOfferActionResValidationWarnings]
}
  // The base class definition for errors
struct EditAnOfferActionResValidationErrors: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  EditAnOfferActionResValidationErrorsMetadata
}
  // The base class definition for metadata
struct EditAnOfferActionResValidationErrorsMetadata: Codable {
		let productId: String
}
  // The base class definition for warnings
struct EditAnOfferActionResValidationWarnings: Codable {
		let code: String
		let details: String
		let message: String
		let path: String
		let userMessage: String
		let metadata:  EditAnOfferActionResValidationWarningsMetadata
}
  // The base class definition for metadata
struct EditAnOfferActionResValidationWarningsMetadata: Codable {
		let productId: String
}
  // The base class definition for afterSalesServices
struct EditAnOfferActionResAfterSalesServices: Codable {
		let impliedWarranty:  EditAnOfferActionResAfterSalesServicesImpliedWarranty
		let returnPolicy:  EditAnOfferActionResAfterSalesServicesReturnPolicy
		let warranty:  EditAnOfferActionResAfterSalesServicesWarranty
}
  // The base class definition for impliedWarranty
struct EditAnOfferActionResAfterSalesServicesImpliedWarranty: Codable {
		let id: String
}
  // The base class definition for returnPolicy
struct EditAnOfferActionResAfterSalesServicesReturnPolicy: Codable {
		let id: String
}
  // The base class definition for warranty
struct EditAnOfferActionResAfterSalesServicesWarranty: Codable {
		let id: String
}
  // The base class definition for discounts
struct EditAnOfferActionResDiscounts: Codable {
		let wholesalePriceList:  EditAnOfferActionResDiscountsWholesalePriceList
}
  // The base class definition for wholesalePriceList
struct EditAnOfferActionResDiscountsWholesalePriceList: Codable {
		let id: String
}
  // The base class definition for contact
struct EditAnOfferActionResContact: Codable {
		let id: String
}
  // The base class definition for attachments
struct EditAnOfferActionResAttachments: Codable {
		let id: String
}
  // The base class definition for fundraisingCampaign
struct EditAnOfferActionResFundraisingCampaign: Codable {
		let id: String
}
  // The base class definition for additionalServices
struct EditAnOfferActionResAdditionalServices: Codable {
		let id: String
}
  // The base class definition for sizeTable
struct EditAnOfferActionResSizeTable: Codable {
		let id: String
}
  // The base class definition for location
struct EditAnOfferActionResLocation: Codable {
		let city: String
		let countryCode: String
		let postCode: String
		let province: String
}
  // The base class definition for external
struct EditAnOfferActionResExternal: Codable {
		let id: String
}
  // The base class definition for taxSettings
struct EditAnOfferActionResTaxSettings: Codable {
		let subject: String
		let exemption: String
		let rates: [EditAnOfferActionResTaxSettingsRates]
}
  // The base class definition for rates
struct EditAnOfferActionResTaxSettingsRates: Codable {
		let rate: String
		let countryCode: String
}
  // The base class definition for messageToSellerSettings
struct EditAnOfferActionResMessageToSellerSettings: Codable {
		let mode: String
		let hint: String
}
  // The base class definition for description
struct EditAnOfferActionResDescription: Codable {
		let sections: [EditAnOfferActionResDescriptionSections]
}
  // The base class definition for sections
struct EditAnOfferActionResDescriptionSections: Codable {
		let items: [EditAnOfferActionResDescriptionSectionsItems]
}
  // The base class definition for items
struct EditAnOfferActionResDescriptionSectionsItems: Codable {
		let type: String
}