package external

import (
	"bytes"
	"encoding/json"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action EditAnOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of EditAnOfferAction
func EditAnOfferAction(c EditAnOfferActionRequest) (*EditAnOfferActionResponse, error) {
	return &EditAnOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func EditAnOfferActionMeta() struct {
	Name        string
	CliName     string
	URL         string
	Method      string
	Description string
} {
	return struct {
		Name        string
		CliName     string
		URL         string
		Method      string
		Description string
	}{
		Name:        "EditAnOfferAction",
		CliName:     "edit an offer-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}",
		Method:      "PATCH",
		Description: `Use this resource to edit offer. This resource allows you to edit each field independently, so use it if you want to change only, for example, the price or the quantity in an offer. Read more: PL / EN. Note that requests may be limited.`,
	}
}

// The base class definition for editAnOfferActionReq
type EditAnOfferActionReq struct {
	Name                    string                                       `json:"name" yaml:"name"`
	Language                string                                       `json:"language" yaml:"language"`
	Category                EditAnOfferActionReqCategory                 `json:"category" yaml:"category"`
	ProductSet              emigo.Array[EditAnOfferActionReqProductSet]  `json:"productSet" yaml:"productSet"`
	Stock                   EditAnOfferActionReqStock                    `json:"stock" yaml:"stock"`
	SellingMode             EditAnOfferActionReqSellingMode              `json:"sellingMode" yaml:"sellingMode"`
	Payments                EditAnOfferActionReqPayments                 `json:"payments" yaml:"payments"`
	Delivery                EditAnOfferActionReqDelivery                 `json:"delivery" yaml:"delivery"`
	Publication             EditAnOfferActionReqPublication              `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  map[any]any                                  `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	CompatibilityList       EditAnOfferActionReqCompatibilityList        `json:"compatibilityList" yaml:"compatibilityList"`
	Images                  []string                                     `json:"images" yaml:"images"`
	Description             EditAnOfferActionReqDescription              `json:"description" yaml:"description"`
	B2b                     EditAnOfferActionReqB2b                      `json:"b2b" yaml:"b2b"`
	Attachments             emigo.Array[EditAnOfferActionReqAttachments] `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     EditAnOfferActionReqFundraisingCampaign      `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      EditAnOfferActionReqAdditionalServices       `json:"additionalServices" yaml:"additionalServices"`
	AfterSalesServices      EditAnOfferActionReqAfterSalesServices       `json:"afterSalesServices" yaml:"afterSalesServices"`
	SizeTable               EditAnOfferActionReqSizeTable                `json:"sizeTable" yaml:"sizeTable"`
	Contact                 EditAnOfferActionReqContact                  `json:"contact" yaml:"contact"`
	Discounts               EditAnOfferActionReqDiscounts                `json:"discounts" yaml:"discounts"`
	Location                EditAnOfferActionReqLocation                 `json:"location" yaml:"location"`
	External                EditAnOfferActionReqExternal                 `json:"external" yaml:"external"`
	TaxSettings             EditAnOfferActionReqTaxSettings              `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings EditAnOfferActionReqMessageToSellerSettings  `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

// The base class definition for category
type EditAnOfferActionReqCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for productSet
type EditAnOfferActionReqProductSet struct {
	Product                      EditAnOfferActionReqProductSetProduct               `json:"product" yaml:"product"`
	Quantity                     EditAnOfferActionReqProductSetQuantity              `json:"quantity" yaml:"quantity"`
	ResponsiblePerson            EditAnOfferActionReqProductSetResponsiblePerson     `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          EditAnOfferActionReqProductSetResponsibleProducer   `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            EditAnOfferActionReqProductSetSafetyInformation     `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     emigo.Array[EditAnOfferActionReqProductSetDeposits] `json:"deposits" yaml:"deposits"`
}

// The base class definition for product
type EditAnOfferActionReqProductSetProduct struct {
	Id         string                                                       `json:"id" yaml:"id"`
	IdType     string                                                       `json:"idType" yaml:"idType"`
	Name       string                                                       `json:"name" yaml:"name"`
	Category   EditAnOfferActionReqProductSetProductCategory                `json:"category" yaml:"category"`
	Parameters emigo.Array[EditAnOfferActionReqProductSetProductParameters] `json:"parameters" yaml:"parameters"`
	Images     []string                                                     `json:"images" yaml:"images"`
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
	Items emigo.Array[EditAnOfferActionReqCompatibilityListItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type EditAnOfferActionReqCompatibilityListItems struct {
	Type string `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
}

// The base class definition for description
type EditAnOfferActionReqDescription struct {
	Sections emigo.Array[EditAnOfferActionReqDescriptionSections] `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type EditAnOfferActionReqDescriptionSections struct {
	Items emigo.Array[EditAnOfferActionReqDescriptionSectionsItems] `json:"items" yaml:"items"`
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
	Subject   string                                            `json:"subject" yaml:"subject"`
	Exemption string                                            `json:"exemption" yaml:"exemption"`
	Rates     emigo.Array[EditAnOfferActionReqTaxSettingsRates] `json:"rates" yaml:"rates"`
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

func (x *EditAnOfferActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

// The base class definition for editAnOfferActionRes
type EditAnOfferActionRes struct {
	Id                      string                                       `json:"id" yaml:"id"`
	Name                    string                                       `json:"name" yaml:"name"`
	Language                string                                       `json:"language" yaml:"language"`
	Category                EditAnOfferActionResCategory                 `json:"category" yaml:"category"`
	ProductSet              emigo.Array[EditAnOfferActionResProductSet]  `json:"productSet" yaml:"productSet"`
	Stock                   EditAnOfferActionResStock                    `json:"stock" yaml:"stock"`
	Payments                EditAnOfferActionResPayments                 `json:"payments" yaml:"payments"`
	SellingMode             EditAnOfferActionResSellingMode              `json:"sellingMode" yaml:"sellingMode"`
	Delivery                EditAnOfferActionResDelivery                 `json:"delivery" yaml:"delivery"`
	Publication             EditAnOfferActionResPublication              `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  EditAnOfferActionResAdditionalMarketplaces   `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	B2b                     EditAnOfferActionResB2b                      `json:"b2b" yaml:"b2b"`
	CompatibilityList       EditAnOfferActionResCompatibilityList        `json:"compatibilityList" yaml:"compatibilityList"`
	Validation              EditAnOfferActionResValidation               `json:"validation" yaml:"validation"`
	Warnings                []string                                     `json:"warnings" yaml:"warnings"`
	AfterSalesServices      EditAnOfferActionResAfterSalesServices       `json:"afterSalesServices" yaml:"afterSalesServices"`
	Discounts               EditAnOfferActionResDiscounts                `json:"discounts" yaml:"discounts"`
	Contact                 EditAnOfferActionResContact                  `json:"contact" yaml:"contact"`
	Attachments             emigo.Array[EditAnOfferActionResAttachments] `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     EditAnOfferActionResFundraisingCampaign      `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      EditAnOfferActionResAdditionalServices       `json:"additionalServices" yaml:"additionalServices"`
	SizeTable               EditAnOfferActionResSizeTable                `json:"sizeTable" yaml:"sizeTable"`
	Location                EditAnOfferActionResLocation                 `json:"location" yaml:"location"`
	External                EditAnOfferActionResExternal                 `json:"external" yaml:"external"`
	TaxSettings             EditAnOfferActionResTaxSettings              `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings EditAnOfferActionResMessageToSellerSettings  `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
	CreatedAt               string                                       `json:"createdAt" yaml:"createdAt"`
	UpdatedAt               string                                       `json:"updatedAt" yaml:"updatedAt"`
	Images                  []string                                     `json:"images" yaml:"images"`
	Description             EditAnOfferActionResDescription              `json:"description" yaml:"description"`
}

// The base class definition for category
type EditAnOfferActionResCategory struct {
	Id string `json:"id" yaml:"id"`
}

// The base class definition for productSet
type EditAnOfferActionResProductSet struct {
	Quantity                     EditAnOfferActionResProductSetQuantity              `json:"quantity" yaml:"quantity"`
	Product                      EditAnOfferActionResProductSetProduct               `json:"product" yaml:"product"`
	ResponsiblePerson            EditAnOfferActionResProductSetResponsiblePerson     `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          EditAnOfferActionResProductSetResponsibleProducer   `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            EditAnOfferActionResProductSetSafetyInformation     `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     emigo.Array[EditAnOfferActionResProductSetDeposits] `json:"deposits" yaml:"deposits"`
}

// The base class definition for quantity
type EditAnOfferActionResProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

// The base class definition for product
type EditAnOfferActionResProductSetProduct struct {
	Id            string                                                       `json:"id" yaml:"id"`
	IsAiCoCreated bool                                                         `json:"isAiCoCreated" yaml:"isAiCoCreated"`
	Publication   EditAnOfferActionResProductSetProductPublication             `json:"publication" yaml:"publication"`
	Parameters    emigo.Array[EditAnOfferActionResProductSetProductParameters] `json:"parameters" yaml:"parameters"`
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
	Base       EditAnOfferActionResPublicationMarketplacesBase                    `json:"base" yaml:"base"`
	Additional emigo.Array[EditAnOfferActionResPublicationMarketplacesAdditional] `json:"additional" yaml:"additional"`
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
	State          string                                                                           `json:"state" yaml:"state"`
	RefusalReasons emigo.Array[EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons] `json:"refusalReasons" yaml:"refusalReasons"`
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
	ValidatedAt string                                              `json:"validatedAt" yaml:"validatedAt"`
	Errors      emigo.Array[EditAnOfferActionResValidationErrors]   `json:"errors" yaml:"errors"`
	Warnings    emigo.Array[EditAnOfferActionResValidationWarnings] `json:"warnings" yaml:"warnings"`
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
	Subject   string                                            `json:"subject" yaml:"subject"`
	Exemption string                                            `json:"exemption" yaml:"exemption"`
	Rates     emigo.Array[EditAnOfferActionResTaxSettingsRates] `json:"rates" yaml:"rates"`
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
	Sections emigo.Array[EditAnOfferActionResDescriptionSections] `json:"sections" yaml:"sections"`
}

// The base class definition for sections
type EditAnOfferActionResDescriptionSections struct {
	Items emigo.Array[EditAnOfferActionResDescriptionSectionsItems] `json:"items" yaml:"items"`
}

// The base class definition for items
type EditAnOfferActionResDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func (x *EditAnOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type EditAnOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *EditAnOfferActionResponse) SetContentType(contentType string) *EditAnOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *EditAnOfferActionResponse) AsStream(r io.Reader, contentType string) *EditAnOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *EditAnOfferActionResponse) AsJSON(payload any) *EditAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *EditAnOfferActionResponse) WithIdeal(payload EditAnOfferActionRes) *EditAnOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *EditAnOfferActionResponse) AsHTML(payload string) *EditAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *EditAnOfferActionResponse) AsBytes(payload []byte) *EditAnOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x EditAnOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x EditAnOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x EditAnOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// Request signature, which is here for refernece. Now it's inlined, so auto completions suggest the function body.
type EditAnOfferActionRequestSig = func(c EditAnOfferActionRequest) (*EditAnOfferActionResponse, error)

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
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	// Now it's interface, so the code gen doesn't depend on the instance
	// or gin package. Make sure you cast is later into *gin.Context, or whatever
	// your framework is passing when creating a request.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	GinCtx interface{}
	// Cli library helper (urfave) by default. The instance is interface{}, and you
	// need to manually cast it to the *cli.Command, so gives you freedom and independence
	// of external library.
	// Ideally, you should not be needing this, and emi has to provide necessary helper
	// functions to read and write a request.
	CliCtx interface{}
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func EditAnOfferActionClientCreateUrl(
	req EditAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := EditAnOfferActionMeta()
	urlAddr := meta.URL
	urlAddr = config.BaseURL + urlAddr
	// Build final URL with query string
	u, err := url.Parse(urlAddr)
	if err != nil {
		return nil, err
	}
	// if UrlValues present, encode and append
	if len(req.QueryParams) > 0 {
		u.RawQuery = req.QueryParams.Encode()
	}
	return u, nil
}
func EditAnOfferActionClientExecuteTyped(httpReq *http.Request) (*EditAnOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result EditAnOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &EditAnOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &EditAnOfferActionResponse{Payload: result}, err
	}
	return &EditAnOfferActionResponse{Payload: result}, nil
}
func EditAnOfferActionClientBuildRequest(req EditAnOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := EditAnOfferActionMeta()
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header = make(http.Header)
	// copy defaults
	for k, v := range config.Headers {
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	// override with request-specific headers
	for k, v := range req.Headers {
		httpReq.Header.Del(k) // ensure override, not duplicate
		for _, vv := range v {
			httpReq.Header.Add(k, vv)
		}
	}
	return httpReq, nil
}
func EditAnOfferActionCall(
	req EditAnOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*EditAnOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := EditAnOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := EditAnOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return EditAnOfferActionClientExecuteTyped(r)
}

// EditAnOfferActionHttpHandler returns the HTTP method, the ServeMux pattern, and a
// typed net/http handler for the EditAnOfferAction action. Developers implement
// their business logic as a function that receives a typed request object and
// returns either an *EditAnOfferActionResponse or nil. JSON marshalling, headers,
// status codes, and errors are handled automatically.
func EditAnOfferActionHttpHandler(
	handler func(c EditAnOfferActionRequest) (*EditAnOfferActionResponse, error),
) (method, pattern string, h http.HandlerFunc) {
	meta := EditAnOfferActionMeta()
	return meta.Method, meta.URL, func(w http.ResponseWriter, r *http.Request) {
		var body EditAnOfferActionReq
		if r.Body != nil {
			defer r.Body.Close()
			if data, _ := io.ReadAll(r.Body); len(data) > 0 {
				if err := json.Unmarshal(data, &body); err != nil {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON: " + err.Error()})
					return
				}
			}
		}
		// Build typed request wrapper. GinCtx stays nil here (this is not gin),
		// which is what the IsGin() helper keys off.
		req := EditAnOfferActionRequest{
			Body:        body,
			QueryParams: r.URL.Query(),
			Headers:     r.Header,
		}
		resp, err := handler(req)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// If the handler returned nil (and no error), the response was handled
		// manually.
		if resp == nil {
			return
		}
		// Apply headers
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		// Apply status and payload
		status := resp.StatusCode
		if status == 0 {
			status = http.StatusOK
		}
		if resp.Payload != nil {
			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json")
			}
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(resp.Payload)
		} else {
			w.WriteHeader(status)
		}
	}
}

// EditAnOfferActionHttp is a high-level convenience wrapper around
// EditAnOfferActionHttpHandler. It registers the typed route on a standard
// *http.ServeMux using Go 1.22+ method-aware pattern syntax (e.g. "POST /").
// Use this when you don't need custom middleware.
func EditAnOfferActionHttp(
	mux *http.ServeMux,
	handler func(c EditAnOfferActionRequest) (*EditAnOfferActionResponse, error),
) {
	method, pattern, h := EditAnOfferActionHttpHandler(handler)
	mux.HandleFunc(method+" "+pattern, h)
}
