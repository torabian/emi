package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
)

/**
* Action to communicate with the action EditAnOfferAction
 */
func EditAnOfferActionMeta() struct {
	Name   string
	URL    string
	Method string
} {
	return struct {
		Name   string
		URL    string
		Method string
	}{
		Name:   "EditAnOfferAction",
		URL:    "https://api.{environment}/sale/product-offers/{offerId}",
		Method: "PATCH",
	}
}

// The base class definition for editAnOfferActionReq
type EditAnOfferActionReq struct {
	Name                    string                                      `json:"name" yaml:"name"`
	Language                string                                      `json:"language" yaml:"language"`
	Category                EditAnOfferActionReqCategory                `json:"category" yaml:"category"`
	ProductSet              []EditAnOfferActionReqProductSet            `json:"productSet" yaml:"productSet"`
	Stock                   EditAnOfferActionReqStock                   `json:"stock" yaml:"stock"`
	SellingMode             EditAnOfferActionReqSellingMode             `json:"sellingMode" yaml:"sellingMode"`
	Payments                EditAnOfferActionReqPayments                `json:"payments" yaml:"payments"`
	Delivery                EditAnOfferActionReqDelivery                `json:"delivery" yaml:"delivery"`
	Publication             EditAnOfferActionReqPublication             `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  interface{}                                 `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	CompatibilityList       EditAnOfferActionReqCompatibilityList       `json:"compatibilityList" yaml:"compatibilityList"`
	Images                  []string                                    `json:"images" yaml:"images"`
	Description             EditAnOfferActionReqDescription             `json:"description" yaml:"description"`
	B2b                     EditAnOfferActionReqB2b                     `json:"b2b" yaml:"b2b"`
	Attachments             []EditAnOfferActionReqAttachments           `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     EditAnOfferActionReqFundraisingCampaign     `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      EditAnOfferActionReqAdditionalServices      `json:"additionalServices" yaml:"additionalServices"`
	AfterSalesServices      EditAnOfferActionReqAfterSalesServices      `json:"afterSalesServices" yaml:"afterSalesServices"`
	SizeTable               EditAnOfferActionReqSizeTable               `json:"sizeTable" yaml:"sizeTable"`
	Contact                 EditAnOfferActionReqContact                 `json:"contact" yaml:"contact"`
	Discounts               EditAnOfferActionReqDiscounts               `json:"discounts" yaml:"discounts"`
	Location                EditAnOfferActionReqLocation                `json:"location" yaml:"location"`
	External                EditAnOfferActionReqExternal                `json:"external" yaml:"external"`
	TaxSettings             EditAnOfferActionReqTaxSettings             `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings EditAnOfferActionReqMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

// The base class definition for category
type EditAnOfferActionReqCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for productSet
type EditAnOfferActionReqProductSet struct {
	Product                      EditAnOfferActionReqProductSetProduct             `json:"product" yaml:"product"`
	Quantity                     EditAnOfferActionReqProductSetQuantity            `json:"quantity" yaml:"quantity"`
	ResponsiblePerson            EditAnOfferActionReqProductSetResponsiblePerson   `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          EditAnOfferActionReqProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            EditAnOfferActionReqProductSetSafetyInformation   `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                              `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     []EditAnOfferActionReqProductSetDeposits          `json:"deposits" yaml:"deposits"`
}

// The base class definition for product
type EditAnOfferActionReqProductSetProduct struct {
	Id         string                                            `json:"id" yaml:"id"`
	IdType     string                                            `json:"idType" yaml:"idType"`
	Name       string                                            `json:"name" yaml:"name"`
	Category   EditAnOfferActionReqProductSetProductCategory     `json:"category" yaml:"category"`
	Parameters []EditAnOfferActionReqProductSetProductParameters `json:"parameters" yaml:"parameters"`
	Images     []string                                          `json:"images" yaml:"images"`
}

// The base class definition for category
type EditAnOfferActionReqProductSetProductCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for parameters
type EditAnOfferActionReqProductSetProductParameters struct {
	Id         string                                                    `json:"id" yaml:"id"`
	Name       string                                                    `json:"name" yaml:"name"`
	RangeValue EditAnOfferActionReqProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                  `json:"values" yaml:"values"`
	ValuesIds  []string                                                  `json:"valuesIds" yaml:"valuesIds"`
}

// The base class definition for rangeValue
type EditAnOfferActionReqProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

// The base class definition for quantity
type EditAnOfferActionReqProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

// The base class definition for responsiblePerson
type EditAnOfferActionReqProductSetResponsiblePerson struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for responsibleProducer
type EditAnOfferActionReqProductSetResponsibleProducer struct {
	Id   string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"`
}

// The base class definition for safetyInformation
type EditAnOfferActionReqProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

// The base class definition for deposits
type EditAnOfferActionReqProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

// The base class definition for stock
type EditAnOfferActionReqStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

// The base class definition for sellingMode
type EditAnOfferActionReqSellingMode struct {
	Format        string                                       `json:"format" yaml:"format"`
	Price         EditAnOfferActionReqSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  EditAnOfferActionReqSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice EditAnOfferActionReqSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

// The base class definition for price
type EditAnOfferActionReqSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for minimalPrice
type EditAnOfferActionReqSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for startingPrice
type EditAnOfferActionReqSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for payments
type EditAnOfferActionReqPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

// The base class definition for delivery
type EditAnOfferActionReqDelivery struct {
	HandlingTime   string                                    `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                    `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                    `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  EditAnOfferActionReqDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

// The base class definition for shippingRates
type EditAnOfferActionReqDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for publication
type EditAnOfferActionReqPublication struct {
	Duration   string `json:"duration" yaml:"duration"`
	StartingAt string `json:"startingAt" yaml:"startingAt"`
	EndingAt   string `json:"endingAt" yaml:"endingAt"`
	Status     string `json:"status" yaml:"status"`
	Republish  bool   `json:"republish" yaml:"republish"`
}

// The base class definition for compatibilityList
type EditAnOfferActionReqCompatibilityList struct {
	Items []EditAnOfferActionReqCompatibilityListItems `json:"items" yaml:"items"`
}

// The base class definition for items
type EditAnOfferActionReqCompatibilityListItems struct {
	Type string `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
}

// The base class definition for description
type EditAnOfferActionReqDescription struct {
	Sections []EditAnOfferActionReqDescriptionSections `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type EditAnOfferActionReqDescriptionSections struct {
	Items []EditAnOfferActionReqDescriptionSectionsItems `json:"items" yaml:"items"`
}

// The base class definition for items
type EditAnOfferActionReqDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for b2b
type EditAnOfferActionReqB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

// The base class definition for attachments
type EditAnOfferActionReqAttachments struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for fundraisingCampaign
type EditAnOfferActionReqFundraisingCampaign struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for additionalServices
type EditAnOfferActionReqAdditionalServices struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for afterSalesServices
type EditAnOfferActionReqAfterSalesServices struct {
	ImpliedWarranty EditAnOfferActionReqAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    EditAnOfferActionReqAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        EditAnOfferActionReqAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

// The base class definition for impliedWarranty
type EditAnOfferActionReqAfterSalesServicesImpliedWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for returnPolicy
type EditAnOfferActionReqAfterSalesServicesReturnPolicy struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for warranty
type EditAnOfferActionReqAfterSalesServicesWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for sizeTable
type EditAnOfferActionReqSizeTable struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for contact
type EditAnOfferActionReqContact struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for discounts
type EditAnOfferActionReqDiscounts struct {
	WholesalePriceList EditAnOfferActionReqDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

// The base class definition for wholesalePriceList
type EditAnOfferActionReqDiscountsWholesalePriceList struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

// The base class definition for location
type EditAnOfferActionReqLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

// The base class definition for external
type EditAnOfferActionReqExternal struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for taxSettings
type EditAnOfferActionReqTaxSettings struct {
	Subject   string                                 `json:"subject" yaml:"subject"`
	Exemption string                                 `json:"exemption" yaml:"exemption"`
	Rates     []EditAnOfferActionReqTaxSettingsRates `json:"rates" yaml:"rates"`
}

// The base class definition for rates
type EditAnOfferActionReqTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

// The base class definition for messageToSellerSettings
type EditAnOfferActionReqMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

// The base class definition for editAnOfferActionRes
type EditAnOfferActionRes struct {
	Id                      string                                      `json:"id" yaml:"id"`
	Name                    string                                      `json:"name" yaml:"name"`
	Language                string                                      `json:"language" yaml:"language"`
	Category                EditAnOfferActionResCategory                `json:"category" yaml:"category"`
	ProductSet              []EditAnOfferActionResProductSet            `json:"productSet" yaml:"productSet"`
	Stock                   EditAnOfferActionResStock                   `json:"stock" yaml:"stock"`
	Payments                EditAnOfferActionResPayments                `json:"payments" yaml:"payments"`
	SellingMode             EditAnOfferActionResSellingMode             `json:"sellingMode" yaml:"sellingMode"`
	Delivery                EditAnOfferActionResDelivery                `json:"delivery" yaml:"delivery"`
	Publication             EditAnOfferActionResPublication             `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  EditAnOfferActionResAdditionalMarketplaces  `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	B2b                     EditAnOfferActionResB2b                     `json:"b2b" yaml:"b2b"`
	CompatibilityList       EditAnOfferActionResCompatibilityList       `json:"compatibilityList" yaml:"compatibilityList"`
	Validation              EditAnOfferActionResValidation              `json:"validation" yaml:"validation"`
	Warnings                []string                                    `json:"warnings" yaml:"warnings"`
	AfterSalesServices      EditAnOfferActionResAfterSalesServices      `json:"afterSalesServices" yaml:"afterSalesServices"`
	Discounts               EditAnOfferActionResDiscounts               `json:"discounts" yaml:"discounts"`
	Contact                 EditAnOfferActionResContact                 `json:"contact" yaml:"contact"`
	Attachments             []EditAnOfferActionResAttachments           `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     EditAnOfferActionResFundraisingCampaign     `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      EditAnOfferActionResAdditionalServices      `json:"additionalServices" yaml:"additionalServices"`
	SizeTable               EditAnOfferActionResSizeTable               `json:"sizeTable" yaml:"sizeTable"`
	Location                EditAnOfferActionResLocation                `json:"location" yaml:"location"`
	External                EditAnOfferActionResExternal                `json:"external" yaml:"external"`
	TaxSettings             EditAnOfferActionResTaxSettings             `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings EditAnOfferActionResMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
	CreatedAt               string                                      `json:"createdAt" yaml:"createdAt"`
	UpdatedAt               string                                      `json:"updatedAt" yaml:"updatedAt"`
	Images                  []string                                    `json:"images" yaml:"images"`
	Description             EditAnOfferActionResDescription             `json:"description" yaml:"description"`
}

// The base class definition for category
type EditAnOfferActionResCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for productSet
type EditAnOfferActionResProductSet struct {
	Quantity                     EditAnOfferActionResProductSetQuantity            `json:"quantity" yaml:"quantity"`
	Product                      EditAnOfferActionResProductSetProduct             `json:"product" yaml:"product"`
	ResponsiblePerson            EditAnOfferActionResProductSetResponsiblePerson   `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          EditAnOfferActionResProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            EditAnOfferActionResProductSetSafetyInformation   `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                              `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     []EditAnOfferActionResProductSetDeposits          `json:"deposits" yaml:"deposits"`
}

// The base class definition for quantity
type EditAnOfferActionResProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

// The base class definition for product
type EditAnOfferActionResProductSetProduct struct {
	Id            string                                            `json:"id" yaml:"id"`
	IsAiCoCreated bool                                              `json:"isAiCoCreated" yaml:"isAiCoCreated"`
	Publication   EditAnOfferActionResProductSetProductPublication  `json:"publication" yaml:"publication"`
	Parameters    []EditAnOfferActionResProductSetProductParameters `json:"parameters" yaml:"parameters"`
}

// The base class definition for publication
type EditAnOfferActionResProductSetProductPublication struct {
	Status string `json:"status" yaml:"status"`
}

// The base class definition for parameters
type EditAnOfferActionResProductSetProductParameters struct {
	Id         string                                                    `json:"id" yaml:"id"`
	Name       string                                                    `json:"name" yaml:"name"`
	RangeValue EditAnOfferActionResProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                  `json:"values" yaml:"values"`
	ValuesIds  []string                                                  `json:"valuesIds" yaml:"valuesIds"`
}

// The base class definition for rangeValue
type EditAnOfferActionResProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

// The base class definition for responsiblePerson
type EditAnOfferActionResProductSetResponsiblePerson struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for responsibleProducer
type EditAnOfferActionResProductSetResponsibleProducer struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for safetyInformation
type EditAnOfferActionResProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

// The base class definition for deposits
type EditAnOfferActionResProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

// The base class definition for stock
type EditAnOfferActionResStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

// The base class definition for payments
type EditAnOfferActionResPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

// The base class definition for sellingMode
type EditAnOfferActionResSellingMode struct {
	Format        string                                       `json:"format" yaml:"format"`
	Price         EditAnOfferActionResSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  EditAnOfferActionResSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice EditAnOfferActionResSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

// The base class definition for price
type EditAnOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for minimalPrice
type EditAnOfferActionResSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for startingPrice
type EditAnOfferActionResSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for delivery
type EditAnOfferActionResDelivery struct {
	HandlingTime   string                                    `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                    `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                    `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  EditAnOfferActionResDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

// The base class definition for shippingRates
type EditAnOfferActionResDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for publication
type EditAnOfferActionResPublication struct {
	Duration     string                                      `json:"duration" yaml:"duration"`
	StartingAt   string                                      `json:"startingAt" yaml:"startingAt"`
	EndingAt     string                                      `json:"endingAt" yaml:"endingAt"`
	EndedBy      string                                      `json:"endedBy" yaml:"endedBy"`
	Status       string                                      `json:"status" yaml:"status"`
	Republish    bool                                        `json:"republish" yaml:"republish"`
	Marketplaces EditAnOfferActionResPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}

// The base class definition for marketplaces
type EditAnOfferActionResPublicationMarketplaces struct {
	Base       EditAnOfferActionResPublicationMarketplacesBase         `json:"base" yaml:"base"`
	Additional []EditAnOfferActionResPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
}

// The base class definition for base
type EditAnOfferActionResPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additional
type EditAnOfferActionResPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalMarketplaces
type EditAnOfferActionResAdditionalMarketplaces struct {
	SellingMode EditAnOfferActionResAdditionalMarketplacesSellingMode `json:"sellingMode" yaml:"sellingMode"`
	Publication EditAnOfferActionResAdditionalMarketplacesPublication `json:"publication" yaml:"publication"`
}

// The base class definition for sellingMode
type EditAnOfferActionResAdditionalMarketplacesSellingMode struct {
	Price EditAnOfferActionResAdditionalMarketplacesSellingModePrice `json:"price" yaml:"price"`
}

// The base class definition for price
type EditAnOfferActionResAdditionalMarketplacesSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

// The base class definition for publication
type EditAnOfferActionResAdditionalMarketplacesPublication struct {
	State          string                                                                `json:"state" yaml:"state"`
	RefusalReasons []EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons `json:"refusalReasons" yaml:"refusalReasons"`
}

// The base class definition for refusalReasons
type EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons struct {
	Code        string                                                                        `json:"code" yaml:"code"`
	UserMessage string                                                                        `json:"userMessage" yaml:"userMessage"`
	Parameters  EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters `json:"parameters" yaml:"parameters"`
}

// The base class definition for parameters
type EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters struct {
	MaxAllowedPriceDecreasePercent []string `json:"maxAllowedPriceDecreasePercent" yaml:"maxAllowedPriceDecreasePercent"`
}

// The base class definition for b2b
type EditAnOfferActionResB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

// The base class definition for compatibilityList
type EditAnOfferActionResCompatibilityList struct {
	Type string `json:"type" yaml:"type"`
}

// The base class definition for validation
type EditAnOfferActionResValidation struct {
	ValidatedAt string                                   `json:"validatedAt" yaml:"validatedAt"`
	Errors      []EditAnOfferActionResValidationErrors   `json:"errors" yaml:"errors"`
	Warnings    []EditAnOfferActionResValidationWarnings `json:"warnings" yaml:"warnings"`
}

// The base class definition for errors
type EditAnOfferActionResValidationErrors struct {
	Code        string                                       `json:"code" yaml:"code"`
	Details     string                                       `json:"details" yaml:"details"`
	Message     string                                       `json:"message" yaml:"message"`
	Path        string                                       `json:"path" yaml:"path"`
	UserMessage string                                       `json:"userMessage" yaml:"userMessage"`
	Metadata    EditAnOfferActionResValidationErrorsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type EditAnOfferActionResValidationErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

// The base class definition for warnings
type EditAnOfferActionResValidationWarnings struct {
	Code        string                                         `json:"code" yaml:"code"`
	Details     string                                         `json:"details" yaml:"details"`
	Message     string                                         `json:"message" yaml:"message"`
	Path        string                                         `json:"path" yaml:"path"`
	UserMessage string                                         `json:"userMessage" yaml:"userMessage"`
	Metadata    EditAnOfferActionResValidationWarningsMetadata `json:"metadata" yaml:"metadata"`
}

// The base class definition for metadata
type EditAnOfferActionResValidationWarningsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

// The base class definition for afterSalesServices
type EditAnOfferActionResAfterSalesServices struct {
	ImpliedWarranty EditAnOfferActionResAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    EditAnOfferActionResAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        EditAnOfferActionResAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

// The base class definition for impliedWarranty
type EditAnOfferActionResAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for returnPolicy
type EditAnOfferActionResAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for warranty
type EditAnOfferActionResAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for discounts
type EditAnOfferActionResDiscounts struct {
	WholesalePriceList EditAnOfferActionResDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

// The base class definition for wholesalePriceList
type EditAnOfferActionResDiscountsWholesalePriceList struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for contact
type EditAnOfferActionResContact struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for attachments
type EditAnOfferActionResAttachments struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for fundraisingCampaign
type EditAnOfferActionResFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for additionalServices
type EditAnOfferActionResAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for sizeTable
type EditAnOfferActionResSizeTable struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for location
type EditAnOfferActionResLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

// The base class definition for external
type EditAnOfferActionResExternal struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for taxSettings
type EditAnOfferActionResTaxSettings struct {
	Subject   string                                 `json:"subject" yaml:"subject"`
	Exemption string                                 `json:"exemption" yaml:"exemption"`
	Rates     []EditAnOfferActionResTaxSettingsRates `json:"rates" yaml:"rates"`
}

// The base class definition for rates
type EditAnOfferActionResTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

// The base class definition for messageToSellerSettings
type EditAnOfferActionResMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

// The base class definition for description
type EditAnOfferActionResDescription struct {
	Sections []EditAnOfferActionResDescriptionSections `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type EditAnOfferActionResDescriptionSections struct {
	Items []EditAnOfferActionResDescriptionSectionsItems `json:"items" yaml:"items"`
}

// The base class definition for items
type EditAnOfferActionResDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}
type EditAnOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

// EditAnOfferActionRaw registers a raw Gin route for the EditAnOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EditAnOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EditAnOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
} // EditAnOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EditAnOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EditAnOfferActionHandler(
	handler func(c EditAnOfferActionRequest, gin *gin.Context) (*EditAnOfferActionResponse, error),
) (method, url string, h gin.HandlerFunc) {
	meta := EditAnOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body EditAnOfferActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := EditAnOfferActionRequest{
			Body:        body,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
		}
		resp, err := handler(req, m)
		if err != nil {
			m.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), it means the response was handled manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			m.Header(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			m.JSON(status, resp.Payload)
		} else {
			m.Status(status)
		}
	}
}

// EditAnOfferAction is a high-level convenience wrapper around EditAnOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EditAnOfferAction(r gin.IRoutes, handler func(c EditAnOfferActionRequest, gin *gin.Context) (*EditAnOfferActionResponse, error)) {
	method, url, h := EditAnOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Edit an offerAction
 */
// Query wrapper with private fields
type EditAnOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func EditAnOfferActionQueryFromString(rawQuery string) EditAnOfferActionQuery {
	v := EditAnOfferActionQuery{}
	values, _ := url.ParseQuery(rawQuery)
	mapped := map[string]interface{}{}
	if result, err := emigo.UnmarshalQs(rawQuery); err == nil {
		mapped = result
	}
	decoder, err := emigo.NewDecoder(&emigo.DecoderConfig{
		TagName:          "json", // reuse json tags
		WeaklyTypedInput: true,   // "1" -> int, "true" -> bool
		Result:           &v,
	})
	if err == nil {
		_ = decoder.Decode(mapped)
	}
	v.values = values
	v.mapped = mapped
	return v
}
func EditAnOfferActionQueryFromGin(c *gin.Context) EditAnOfferActionQuery {
	return EditAnOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func EditAnOfferActionQueryFromHttp(r *http.Request) EditAnOfferActionQuery {
	return EditAnOfferActionQueryFromString(r.URL.RawQuery)
}
func (q EditAnOfferActionQuery) Values() url.Values {
	return q.values
}
func (q EditAnOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *EditAnOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *EditAnOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type EditAnOfferActionRequest struct {
	Body        EditAnOfferActionReq
	QueryParams url.Values
	Headers     http.Header
}
type EditAnOfferActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func EditAnOfferActionCall(
	req EditAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EditAnOfferActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := EditAnOfferActionMeta()
		baseURL := meta.URL
		// Build final URL with query string
		u, err := url.Parse(baseURL)
		if err != nil {
			return nil, err
		}
		// if UrlValues present, encode and append
		if len(req.QueryParams) > 0 {
			u.RawQuery = req.QueryParams.Encode()
		}
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		req0, err := http.NewRequest(meta.Method, u.String(), bytes.NewReader(bodyBytes))
		if err != nil {
			return nil, err
		}
		httpReq = req0
	} else {
		httpReq = config.Httpr
	}
	httpReq.Header = req.Headers
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	var result EditAnOfferActionResult
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	if resp.StatusCode >= 400 {
		return &result, fmt.Errorf("request failed: %s", respBody)
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &result, err
	}
	return &result, nil
}
