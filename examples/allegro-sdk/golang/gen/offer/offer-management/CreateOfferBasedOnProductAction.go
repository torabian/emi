package external

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action CreateOfferBasedOnProductAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of CreateOfferBasedOnProductAction
func CreateOfferBasedOnProductAction(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error) {
	return &CreateOfferBasedOnProductActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func CreateOfferBasedOnProductActionMeta() struct {
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
		Name:        "CreateOfferBasedOnProductAction",
		CliName:     "create offer based on product-action",
		URL:         "https://api.{environment}/sale/product-offers",
		Method:      "POST",
		Description: ``,
	}
}
func GetCreateOfferBasedOnProductActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Offer title",
		},
		{
			Name:        prefix + "language",
			Type:        "string",
			Description: "Offer language code (e.g., pl-PL)",
		},
		{
			Name:     prefix + "category",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqCategoryCliFlags("category-"),
		},
		{
			Name:        prefix + "product-set",
			Type:        "array",
			Description: "Product details and associated quantities",
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "payments",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqPaymentsCliFlags("payments-"),
		},
		{
			Name:     prefix + "delivery",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqDeliveryCliFlags("delivery-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqPublicationCliFlags("publication-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "map?",
		},
		{
			Name:     prefix + "compatibility-list",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqCompatibilityListCliFlags("compatibility-list-"),
		},
		{
			Name: prefix + "images",
			Type: "slice",
		},
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqDescriptionCliFlags("description-"),
		},
		{
			Name:     prefix + "b2b",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqB2bCliFlags("b2b-"),
		},
		{
			Name: prefix + "attachments",
			Type: "array",
		},
		{
			Name:     prefix + "fundraising-campaign",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqFundraisingCampaignCliFlags("fundraising-campaign-"),
		},
		{
			Name:     prefix + "additional-services",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqAdditionalServicesCliFlags("additional-services-"),
		},
		{
			Name:     prefix + "after-sales-services",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqAfterSalesServicesCliFlags("after-sales-services-"),
		},
		{
			Name:     prefix + "size-table",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqSizeTableCliFlags("size-table-"),
		},
		{
			Name:     prefix + "contact",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqContactCliFlags("contact-"),
		},
		{
			Name:     prefix + "discounts",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqDiscountsCliFlags("discounts-"),
		},
		{
			Name:     prefix + "location",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqLocationCliFlags("location-"),
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqExternalCliFlags("external-"),
		},
		{
			Name:     prefix + "tax-settings",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqTaxSettingsCliFlags("tax-settings-"),
		},
		{
			Name:     prefix + "message-to-seller-settings",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqMessageToSellerSettingsCliFlags("message-to-seller-settings-"),
		},
	}
}
func CastCreateOfferBasedOnProductActionReqFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReq {
	data := CreateOfferBasedOnProductActionReq{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("language") {
		data.Language = c.String("language")
	}
	if c.IsSet("category") {
		data.Category = CastCreateOfferBasedOnProductActionReqCategoryFromCli(c)
	}
	if c.IsSet("product-set") {
		data.ProductSet = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqProductSetFromCli, "product-set", c)
	}
	if c.IsSet("stock") {
		data.Stock = CastCreateOfferBasedOnProductActionReqStockFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastCreateOfferBasedOnProductActionReqSellingModeFromCli(c)
	}
	if c.IsSet("payments") {
		data.Payments = CastCreateOfferBasedOnProductActionReqPaymentsFromCli(c)
	}
	if c.IsSet("delivery") {
		data.Delivery = CastCreateOfferBasedOnProductActionReqDeliveryFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastCreateOfferBasedOnProductActionReqPublicationFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		emigo.ParseNullable(c.String("additional-marketplaces"), &data.AdditionalMarketplaces)
	}
	if c.IsSet("compatibility-list") {
		data.CompatibilityList = CastCreateOfferBasedOnProductActionReqCompatibilityListFromCli(c)
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	if c.IsSet("description") {
		data.Description = CastCreateOfferBasedOnProductActionReqDescriptionFromCli(c)
	}
	if c.IsSet("b2b") {
		data.B2b = CastCreateOfferBasedOnProductActionReqB2bFromCli(c)
	}
	if c.IsSet("attachments") {
		data.Attachments = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqAttachmentsFromCli, "attachments", c)
	}
	if c.IsSet("fundraising-campaign") {
		data.FundraisingCampaign = CastCreateOfferBasedOnProductActionReqFundraisingCampaignFromCli(c)
	}
	if c.IsSet("additional-services") {
		data.AdditionalServices = CastCreateOfferBasedOnProductActionReqAdditionalServicesFromCli(c)
	}
	if c.IsSet("after-sales-services") {
		data.AfterSalesServices = CastCreateOfferBasedOnProductActionReqAfterSalesServicesFromCli(c)
	}
	if c.IsSet("size-table") {
		data.SizeTable = CastCreateOfferBasedOnProductActionReqSizeTableFromCli(c)
	}
	if c.IsSet("contact") {
		data.Contact = CastCreateOfferBasedOnProductActionReqContactFromCli(c)
	}
	if c.IsSet("discounts") {
		data.Discounts = CastCreateOfferBasedOnProductActionReqDiscountsFromCli(c)
	}
	if c.IsSet("location") {
		data.Location = CastCreateOfferBasedOnProductActionReqLocationFromCli(c)
	}
	if c.IsSet("external") {
		data.External = CastCreateOfferBasedOnProductActionReqExternalFromCli(c)
	}
	if c.IsSet("tax-settings") {
		data.TaxSettings = CastCreateOfferBasedOnProductActionReqTaxSettingsFromCli(c)
	}
	if c.IsSet("message-to-seller-settings") {
		data.MessageToSellerSettings = CastCreateOfferBasedOnProductActionReqMessageToSellerSettingsFromCli(c)
	}
	return data
}

// The base class definition for createOfferBasedOnProductActionReq
type CreateOfferBasedOnProductActionReq struct {
	// Offer title
	Name string `json:"name" yaml:"name"`
	// Offer language code (e.g., pl-PL)
	Language string                                     `json:"language" yaml:"language"`
	Category CreateOfferBasedOnProductActionReqCategory `json:"category" yaml:"category"`
	// Product details and associated quantities
	ProductSet              []CreateOfferBasedOnProductActionReqProductSet            `json:"productSet" yaml:"productSet"`
	Stock                   CreateOfferBasedOnProductActionReqStock                   `json:"stock" yaml:"stock"`
	SellingMode             CreateOfferBasedOnProductActionReqSellingMode             `json:"sellingMode" yaml:"sellingMode"`
	Payments                CreateOfferBasedOnProductActionReqPayments                `json:"payments" yaml:"payments"`
	Delivery                CreateOfferBasedOnProductActionReqDelivery                `json:"delivery" yaml:"delivery"`
	Publication             CreateOfferBasedOnProductActionReqPublication             `json:"publication" yaml:"publication"`
	AdditionalMarketplaces  emigo.Nullable[map[any]any]                               `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	CompatibilityList       CreateOfferBasedOnProductActionReqCompatibilityList       `json:"compatibilityList" yaml:"compatibilityList"`
	Images                  []string                                                  `json:"images" yaml:"images"`
	Description             CreateOfferBasedOnProductActionReqDescription             `json:"description" yaml:"description"`
	B2b                     CreateOfferBasedOnProductActionReqB2b                     `json:"b2b" yaml:"b2b"`
	Attachments             []CreateOfferBasedOnProductActionReqAttachments           `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     CreateOfferBasedOnProductActionReqFundraisingCampaign     `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      CreateOfferBasedOnProductActionReqAdditionalServices      `json:"additionalServices" yaml:"additionalServices"`
	AfterSalesServices      CreateOfferBasedOnProductActionReqAfterSalesServices      `json:"afterSalesServices" yaml:"afterSalesServices"`
	SizeTable               CreateOfferBasedOnProductActionReqSizeTable               `json:"sizeTable" yaml:"sizeTable"`
	Contact                 CreateOfferBasedOnProductActionReqContact                 `json:"contact" yaml:"contact"`
	Discounts               CreateOfferBasedOnProductActionReqDiscounts               `json:"discounts" yaml:"discounts"`
	Location                CreateOfferBasedOnProductActionReqLocation                `json:"location" yaml:"location"`
	External                CreateOfferBasedOnProductActionReqExternal                `json:"external" yaml:"external"`
	TaxSettings             CreateOfferBasedOnProductActionReqTaxSettings             `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings CreateOfferBasedOnProductActionReqMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

func GetCreateOfferBasedOnProductActionReqCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqCategoryFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqCategory {
	data := CreateOfferBasedOnProductActionReqCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type CreateOfferBasedOnProductActionReqCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetCreateOfferBasedOnProductActionReqProductSetCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "product",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetProductCliFlags("product-"),
		},
		{
			Name:     prefix + "quantity",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetQuantityCliFlags("quantity-"),
		},
		{
			Name:     prefix + "responsible-person",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetResponsiblePersonCliFlags("responsible-person-"),
		},
		{
			Name:     prefix + "responsible-producer",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetResponsibleProducerCliFlags("responsible-producer-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetSafetyInformationCliFlags("safety-information-"),
		},
		{
			Name: prefix + "marketed-before-gpsr-obligation",
			Type: "bool",
		},
		{
			Name: prefix + "deposits",
			Type: "array",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSet {
	data := CreateOfferBasedOnProductActionReqProductSet{}
	if c.IsSet("product") {
		data.Product = CastCreateOfferBasedOnProductActionReqProductSetProductFromCli(c)
	}
	if c.IsSet("quantity") {
		data.Quantity = CastCreateOfferBasedOnProductActionReqProductSetQuantityFromCli(c)
	}
	if c.IsSet("responsible-person") {
		data.ResponsiblePerson = CastCreateOfferBasedOnProductActionReqProductSetResponsiblePersonFromCli(c)
	}
	if c.IsSet("responsible-producer") {
		data.ResponsibleProducer = CastCreateOfferBasedOnProductActionReqProductSetResponsibleProducerFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastCreateOfferBasedOnProductActionReqProductSetSafetyInformationFromCli(c)
	}
	if c.IsSet("marketed-before-gpsr-obligation") {
		data.MarketedBeforeGPSRObligation = bool(c.Bool("marketed-before-gpsr-obligation"))
	}
	if c.IsSet("deposits") {
		data.Deposits = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqProductSetDepositsFromCli, "deposits", c)
	}
	return data
}

// The base class definition for productSet
type CreateOfferBasedOnProductActionReqProductSet struct {
	Product                      CreateOfferBasedOnProductActionReqProductSetProduct             `json:"product" yaml:"product"`
	Quantity                     CreateOfferBasedOnProductActionReqProductSetQuantity            `json:"quantity" yaml:"quantity"`
	ResponsiblePerson            CreateOfferBasedOnProductActionReqProductSetResponsiblePerson   `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          CreateOfferBasedOnProductActionReqProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            CreateOfferBasedOnProductActionReqProductSetSafetyInformation   `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                            `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     []CreateOfferBasedOnProductActionReqProductSetDeposits          `json:"deposits" yaml:"deposits"`
}

func GetCreateOfferBasedOnProductActionReqProductSetProductCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "id-type",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:     prefix + "category",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetProductCategoryCliFlags("category-"),
		},
		{
			Name: prefix + "parameters",
			Type: "array",
		},
		{
			Name: prefix + "images",
			Type: "slice",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetProductFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetProduct {
	data := CreateOfferBasedOnProductActionReqProductSetProduct{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("id-type") {
		data.IdType = c.String("id-type")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("category") {
		data.Category = CastCreateOfferBasedOnProductActionReqProductSetProductCategoryFromCli(c)
	}
	if c.IsSet("parameters") {
		data.Parameters = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqProductSetProductParametersFromCli, "parameters", c)
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	return data
}

// The base class definition for product
type CreateOfferBasedOnProductActionReqProductSetProduct struct {
	Id         string                                                          `json:"id" yaml:"id"`
	IdType     string                                                          `json:"idType" yaml:"idType"`
	Name       string                                                          `json:"name" yaml:"name"`
	Category   CreateOfferBasedOnProductActionReqProductSetProductCategory     `json:"category" yaml:"category"`
	Parameters []CreateOfferBasedOnProductActionReqProductSetProductParameters `json:"parameters" yaml:"parameters"`
	Images     []string                                                        `json:"images" yaml:"images"`
}

func GetCreateOfferBasedOnProductActionReqProductSetProductCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetProductCategoryFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetProductCategory {
	data := CreateOfferBasedOnProductActionReqProductSetProductCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type CreateOfferBasedOnProductActionReqProductSetProductCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetCreateOfferBasedOnProductActionReqProductSetProductParametersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name:     prefix + "range-value",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqProductSetProductParametersRangeValueCliFlags("range-value-"),
		},
		{
			Name: prefix + "values",
			Type: "slice",
		},
		{
			Name: prefix + "values-ids",
			Type: "slice",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetProductParametersFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetProductParameters {
	data := CreateOfferBasedOnProductActionReqProductSetProductParameters{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("range-value") {
		data.RangeValue = CastCreateOfferBasedOnProductActionReqProductSetProductParametersRangeValueFromCli(c)
	}
	if c.IsSet("values") {
		emigo.InflatePossibleSlice(c.String("values"), &data.Values)
	}
	if c.IsSet("values-ids") {
		emigo.InflatePossibleSlice(c.String("values-ids"), &data.ValuesIds)
	}
	return data
}

// The base class definition for parameters
type CreateOfferBasedOnProductActionReqProductSetProductParameters struct {
	Id         string                                                                  `json:"id" yaml:"id"`
	Name       string                                                                  `json:"name" yaml:"name"`
	RangeValue CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                                `json:"values" yaml:"values"`
	ValuesIds  []string                                                                `json:"valuesIds" yaml:"valuesIds"`
}

func GetCreateOfferBasedOnProductActionReqProductSetProductParametersRangeValueCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "from",
			Type: "string",
		},
		{
			Name: prefix + "to",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetProductParametersRangeValueFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue {
	data := CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue{}
	if c.IsSet("from") {
		data.From = c.String("from")
	}
	if c.IsSet("to") {
		data.To = c.String("to")
	}
	return data
}

// The base class definition for rangeValue
type CreateOfferBasedOnProductActionReqProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

func GetCreateOfferBasedOnProductActionReqProductSetQuantityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "int",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetQuantityFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetQuantity {
	data := CreateOfferBasedOnProductActionReqProductSetQuantity{}
	if c.IsSet("value") {
		data.Value = int(c.Int64("value"))
	}
	return data
}

// The base class definition for quantity
type CreateOfferBasedOnProductActionReqProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

func GetCreateOfferBasedOnProductActionReqProductSetResponsiblePersonCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetResponsiblePersonFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetResponsiblePerson {
	data := CreateOfferBasedOnProductActionReqProductSetResponsiblePerson{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for responsiblePerson
type CreateOfferBasedOnProductActionReqProductSetResponsiblePerson struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqProductSetResponsibleProducerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetResponsibleProducerFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetResponsibleProducer {
	data := CreateOfferBasedOnProductActionReqProductSetResponsibleProducer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for responsibleProducer
type CreateOfferBasedOnProductActionReqProductSetResponsibleProducer struct {
	Id   string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"`
}

func GetCreateOfferBasedOnProductActionReqProductSetSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "description",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetSafetyInformationFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetSafetyInformation {
	data := CreateOfferBasedOnProductActionReqProductSetSafetyInformation{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

// The base class definition for safetyInformation
type CreateOfferBasedOnProductActionReqProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

func GetCreateOfferBasedOnProductActionReqProductSetDepositsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "quantity",
			Type: "int",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqProductSetDepositsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqProductSetDeposits {
	data := CreateOfferBasedOnProductActionReqProductSetDeposits{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("quantity") {
		data.Quantity = int(c.Int64("quantity"))
	}
	return data
}

// The base class definition for deposits
type CreateOfferBasedOnProductActionReqProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

func GetCreateOfferBasedOnProductActionReqStockCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "available",
			Type: "int",
		},
		{
			Name: prefix + "unit",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqStockFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqStock {
	data := CreateOfferBasedOnProductActionReqStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	if c.IsSet("unit") {
		data.Unit = c.String("unit")
	}
	return data
}

// The base class definition for stock
type CreateOfferBasedOnProductActionReqStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

func GetCreateOfferBasedOnProductActionReqSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "format",
			Type: "string",
		},
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqSellingModePriceCliFlags("price-"),
		},
		{
			Name:     prefix + "minimal-price",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqSellingModeMinimalPriceCliFlags("minimal-price-"),
		},
		{
			Name:     prefix + "starting-price",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqSellingModeStartingPriceCliFlags("starting-price-"),
		},
	}
}
func CastCreateOfferBasedOnProductActionReqSellingModeFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqSellingMode {
	data := CreateOfferBasedOnProductActionReqSellingMode{}
	if c.IsSet("format") {
		data.Format = c.String("format")
	}
	if c.IsSet("price") {
		data.Price = CastCreateOfferBasedOnProductActionReqSellingModePriceFromCli(c)
	}
	if c.IsSet("minimal-price") {
		data.MinimalPrice = CastCreateOfferBasedOnProductActionReqSellingModeMinimalPriceFromCli(c)
	}
	if c.IsSet("starting-price") {
		data.StartingPrice = CastCreateOfferBasedOnProductActionReqSellingModeStartingPriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type CreateOfferBasedOnProductActionReqSellingMode struct {
	Format        string                                                     `json:"format" yaml:"format"`
	Price         CreateOfferBasedOnProductActionReqSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  CreateOfferBasedOnProductActionReqSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice CreateOfferBasedOnProductActionReqSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

func GetCreateOfferBasedOnProductActionReqSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqSellingModePriceFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqSellingModePrice {
	data := CreateOfferBasedOnProductActionReqSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type CreateOfferBasedOnProductActionReqSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetCreateOfferBasedOnProductActionReqSellingModeMinimalPriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqSellingModeMinimalPriceFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqSellingModeMinimalPrice {
	data := CreateOfferBasedOnProductActionReqSellingModeMinimalPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for minimalPrice
type CreateOfferBasedOnProductActionReqSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetCreateOfferBasedOnProductActionReqSellingModeStartingPriceCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "amount",
			Type: "string",
		},
		{
			Name: prefix + "currency",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqSellingModeStartingPriceFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqSellingModeStartingPrice {
	data := CreateOfferBasedOnProductActionReqSellingModeStartingPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for startingPrice
type CreateOfferBasedOnProductActionReqSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetCreateOfferBasedOnProductActionReqPaymentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "invoice",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqPaymentsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqPayments {
	data := CreateOfferBasedOnProductActionReqPayments{}
	if c.IsSet("invoice") {
		data.Invoice = c.String("invoice")
	}
	return data
}

// The base class definition for payments
type CreateOfferBasedOnProductActionReqPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

func GetCreateOfferBasedOnProductActionReqDeliveryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "handling-time",
			Type: "string",
		},
		{
			Name: prefix + "additional-info",
			Type: "string",
		},
		{
			Name: prefix + "shipment-date",
			Type: "string",
		},
		{
			Name:        prefix + "shipping-rates",
			Type:        "object",
			Children:    GetCreateOfferBasedOnProductActionReqDeliveryShippingRatesCliFlags("shipping-rates-"),
			Description: "Optional; may be null",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDeliveryFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDelivery {
	data := CreateOfferBasedOnProductActionReqDelivery{}
	if c.IsSet("handling-time") {
		data.HandlingTime = c.String("handling-time")
	}
	if c.IsSet("additional-info") {
		data.AdditionalInfo = c.String("additional-info")
	}
	if c.IsSet("shipment-date") {
		data.ShipmentDate = c.String("shipment-date")
	}
	if c.IsSet("shipping-rates") {
		data.ShippingRates = CastCreateOfferBasedOnProductActionReqDeliveryShippingRatesFromCli(c)
	}
	return data
}

// The base class definition for delivery
type CreateOfferBasedOnProductActionReqDelivery struct {
	HandlingTime   string `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string `json:"shipmentDate" yaml:"shipmentDate"`
	// Optional; may be null
	ShippingRates CreateOfferBasedOnProductActionReqDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

func GetCreateOfferBasedOnProductActionReqDeliveryShippingRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDeliveryShippingRatesFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDeliveryShippingRates {
	data := CreateOfferBasedOnProductActionReqDeliveryShippingRates{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for shippingRates
type CreateOfferBasedOnProductActionReqDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

func GetCreateOfferBasedOnProductActionReqPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "duration",
			Type: "string",
		},
		{
			Name: prefix + "starting-at",
			Type: "string",
		},
		{
			Name: prefix + "ending-at",
			Type: "string",
		},
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "republish",
			Type: "bool",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqPublicationFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqPublication {
	data := CreateOfferBasedOnProductActionReqPublication{}
	if c.IsSet("duration") {
		data.Duration = c.String("duration")
	}
	if c.IsSet("starting-at") {
		data.StartingAt = c.String("starting-at")
	}
	if c.IsSet("ending-at") {
		data.EndingAt = c.String("ending-at")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("republish") {
		data.Republish = bool(c.Bool("republish"))
	}
	return data
}

// The base class definition for publication
type CreateOfferBasedOnProductActionReqPublication struct {
	Duration   string `json:"duration" yaml:"duration"`
	StartingAt string `json:"startingAt" yaml:"startingAt"`
	EndingAt   string `json:"endingAt" yaml:"endingAt"`
	Status     string `json:"status" yaml:"status"`
	Republish  bool   `json:"republish" yaml:"republish"`
}

func GetCreateOfferBasedOnProductActionReqCompatibilityListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqCompatibilityListFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqCompatibilityList {
	data := CreateOfferBasedOnProductActionReqCompatibilityList{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqCompatibilityListItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for compatibilityList
type CreateOfferBasedOnProductActionReqCompatibilityList struct {
	Items []CreateOfferBasedOnProductActionReqCompatibilityListItems `json:"items" yaml:"items"`
}

func GetCreateOfferBasedOnProductActionReqCompatibilityListItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
		{
			Name: prefix + "text",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqCompatibilityListItemsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqCompatibilityListItems {
	data := CreateOfferBasedOnProductActionReqCompatibilityListItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("text") {
		data.Text = c.String("text")
	}
	return data
}

// The base class definition for items
type CreateOfferBasedOnProductActionReqCompatibilityListItems struct {
	Type string `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
}

func GetCreateOfferBasedOnProductActionReqDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDescriptionFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDescription {
	data := CreateOfferBasedOnProductActionReqDescription{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqDescriptionSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for description
type CreateOfferBasedOnProductActionReqDescription struct {
	Sections []CreateOfferBasedOnProductActionReqDescriptionSections `json:"sections" yaml:"sections"`
}

func GetCreateOfferBasedOnProductActionReqDescriptionSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDescriptionSectionsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDescriptionSections {
	data := CreateOfferBasedOnProductActionReqDescriptionSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqDescriptionSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type CreateOfferBasedOnProductActionReqDescriptionSections struct {
	Items []CreateOfferBasedOnProductActionReqDescriptionSectionsItems `json:"items" yaml:"items"`
}

func GetCreateOfferBasedOnProductActionReqDescriptionSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDescriptionSectionsItemsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDescriptionSectionsItems {
	data := CreateOfferBasedOnProductActionReqDescriptionSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for items
type CreateOfferBasedOnProductActionReqDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func GetCreateOfferBasedOnProductActionReqB2bCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "buyable-only-by-business",
			Type: "bool",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqB2bFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqB2b {
	data := CreateOfferBasedOnProductActionReqB2b{}
	if c.IsSet("buyable-only-by-business") {
		data.BuyableOnlyByBusiness = bool(c.Bool("buyable-only-by-business"))
	}
	return data
}

// The base class definition for b2b
type CreateOfferBasedOnProductActionReqB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

func GetCreateOfferBasedOnProductActionReqAttachmentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAttachmentsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAttachments {
	data := CreateOfferBasedOnProductActionReqAttachments{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for attachments
type CreateOfferBasedOnProductActionReqAttachments struct {
	Id string `json:"id" yaml:"id"`
}

func GetCreateOfferBasedOnProductActionReqFundraisingCampaignCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqFundraisingCampaignFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqFundraisingCampaign {
	data := CreateOfferBasedOnProductActionReqFundraisingCampaign{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for fundraisingCampaign
type CreateOfferBasedOnProductActionReqFundraisingCampaign struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqAdditionalServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAdditionalServicesFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAdditionalServices {
	data := CreateOfferBasedOnProductActionReqAdditionalServices{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for additionalServices
type CreateOfferBasedOnProductActionReqAdditionalServices struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqAfterSalesServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "implied-warranty",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarrantyCliFlags("implied-warranty-"),
		},
		{
			Name:     prefix + "return-policy",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicyCliFlags("return-policy-"),
		},
		{
			Name:     prefix + "warranty",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqAfterSalesServicesWarrantyCliFlags("warranty-"),
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAfterSalesServicesFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAfterSalesServices {
	data := CreateOfferBasedOnProductActionReqAfterSalesServices{}
	if c.IsSet("implied-warranty") {
		data.ImpliedWarranty = CastCreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarrantyFromCli(c)
	}
	if c.IsSet("return-policy") {
		data.ReturnPolicy = CastCreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicyFromCli(c)
	}
	if c.IsSet("warranty") {
		data.Warranty = CastCreateOfferBasedOnProductActionReqAfterSalesServicesWarrantyFromCli(c)
	}
	return data
}

// The base class definition for afterSalesServices
type CreateOfferBasedOnProductActionReqAfterSalesServices struct {
	ImpliedWarranty CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

func GetCreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarrantyFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty {
	data := CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for impliedWarranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesImpliedWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicyFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy {
	data := CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for returnPolicy
type CreateOfferBasedOnProductActionReqAfterSalesServicesReturnPolicy struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqAfterSalesServicesWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqAfterSalesServicesWarrantyFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty {
	data := CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for warranty
type CreateOfferBasedOnProductActionReqAfterSalesServicesWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqSizeTableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqSizeTableFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqSizeTable {
	data := CreateOfferBasedOnProductActionReqSizeTable{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for sizeTable
type CreateOfferBasedOnProductActionReqSizeTable struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqContactCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqContactFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqContact {
	data := CreateOfferBasedOnProductActionReqContact{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for contact
type CreateOfferBasedOnProductActionReqContact struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqDiscountsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "wholesale-price-list",
			Type:     "object",
			Children: GetCreateOfferBasedOnProductActionReqDiscountsWholesalePriceListCliFlags("wholesale-price-list-"),
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDiscountsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDiscounts {
	data := CreateOfferBasedOnProductActionReqDiscounts{}
	if c.IsSet("wholesale-price-list") {
		data.WholesalePriceList = CastCreateOfferBasedOnProductActionReqDiscountsWholesalePriceListFromCli(c)
	}
	return data
}

// The base class definition for discounts
type CreateOfferBasedOnProductActionReqDiscounts struct {
	WholesalePriceList CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

func GetCreateOfferBasedOnProductActionReqDiscountsWholesalePriceListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqDiscountsWholesalePriceListFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList {
	data := CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for wholesalePriceList
type CreateOfferBasedOnProductActionReqDiscountsWholesalePriceList struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetCreateOfferBasedOnProductActionReqLocationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "city",
			Type: "string",
		},
		{
			Name: prefix + "country-code",
			Type: "string",
		},
		{
			Name: prefix + "post-code",
			Type: "string",
		},
		{
			Name: prefix + "province",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqLocationFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqLocation {
	data := CreateOfferBasedOnProductActionReqLocation{}
	if c.IsSet("city") {
		data.City = c.String("city")
	}
	if c.IsSet("country-code") {
		data.CountryCode = c.String("country-code")
	}
	if c.IsSet("post-code") {
		data.PostCode = c.String("post-code")
	}
	if c.IsSet("province") {
		data.Province = c.String("province")
	}
	return data
}

// The base class definition for location
type CreateOfferBasedOnProductActionReqLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

func GetCreateOfferBasedOnProductActionReqExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqExternalFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqExternal {
	data := CreateOfferBasedOnProductActionReqExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type CreateOfferBasedOnProductActionReqExternal struct {
	Id string `json:"id" yaml:"id"`
}

func GetCreateOfferBasedOnProductActionReqTaxSettingsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "subject",
			Type: "string",
		},
		{
			Name: prefix + "exemption",
			Type: "string",
		},
		{
			Name: prefix + "rates",
			Type: "array",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqTaxSettingsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqTaxSettings {
	data := CreateOfferBasedOnProductActionReqTaxSettings{}
	if c.IsSet("subject") {
		data.Subject = c.String("subject")
	}
	if c.IsSet("exemption") {
		data.Exemption = c.String("exemption")
	}
	if c.IsSet("rates") {
		data.Rates = emigo.CapturePossibleArray(CastCreateOfferBasedOnProductActionReqTaxSettingsRatesFromCli, "rates", c)
	}
	return data
}

// The base class definition for taxSettings
type CreateOfferBasedOnProductActionReqTaxSettings struct {
	Subject   string                                               `json:"subject" yaml:"subject"`
	Exemption string                                               `json:"exemption" yaml:"exemption"`
	Rates     []CreateOfferBasedOnProductActionReqTaxSettingsRates `json:"rates" yaml:"rates"`
}

func GetCreateOfferBasedOnProductActionReqTaxSettingsRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "rate",
			Type: "string",
		},
		{
			Name: prefix + "country-code",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqTaxSettingsRatesFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqTaxSettingsRates {
	data := CreateOfferBasedOnProductActionReqTaxSettingsRates{}
	if c.IsSet("rate") {
		data.Rate = c.String("rate")
	}
	if c.IsSet("country-code") {
		data.CountryCode = c.String("country-code")
	}
	return data
}

// The base class definition for rates
type CreateOfferBasedOnProductActionReqTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

func GetCreateOfferBasedOnProductActionReqMessageToSellerSettingsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "mode",
			Type: "string",
		},
		{
			Name: prefix + "hint",
			Type: "string",
		},
	}
}
func CastCreateOfferBasedOnProductActionReqMessageToSellerSettingsFromCli(c emigo.CliCastable) CreateOfferBasedOnProductActionReqMessageToSellerSettings {
	data := CreateOfferBasedOnProductActionReqMessageToSellerSettings{}
	if c.IsSet("mode") {
		data.Mode = c.String("mode")
	}
	if c.IsSet("hint") {
		data.Hint = c.String("hint")
	}
	return data
}

// The base class definition for messageToSellerSettings
type CreateOfferBasedOnProductActionReqMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

func (x *CreateOfferBasedOnProductActionReq) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type CreateOfferBasedOnProductActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *CreateOfferBasedOnProductActionResponse) SetContentType(contentType string) *CreateOfferBasedOnProductActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsStream(r io.Reader, contentType string) *CreateOfferBasedOnProductActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsJSON(payload any) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsHTML(payload string) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *CreateOfferBasedOnProductActionResponse) AsBytes(payload []byte) *CreateOfferBasedOnProductActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x CreateOfferBasedOnProductActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x CreateOfferBasedOnProductActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x CreateOfferBasedOnProductActionResponse) GetPayload() interface{} {
	return x.Payload
}

// CreateOfferBasedOnProductActionRaw registers a raw Gin route for the CreateOfferBasedOnProductAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func CreateOfferBasedOnProductActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := CreateOfferBasedOnProductActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type CreateOfferBasedOnProductActionRequestSig = func(c CreateOfferBasedOnProductActionRequest) (*CreateOfferBasedOnProductActionResponse, error)

// CreateOfferBasedOnProductActionHandler returns the HTTP method, route URL, and a typed Gin handler for the CreateOfferBasedOnProductAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func CreateOfferBasedOnProductActionHandler(
	handler CreateOfferBasedOnProductActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := CreateOfferBasedOnProductActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		var body CreateOfferBasedOnProductActionReq
		if err := m.ShouldBindJSON(&body); err != nil {
			m.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
			return
		}
		// Build typed request wrapper
		req := CreateOfferBasedOnProductActionRequest{
			Body:        body,
			QueryParams: m.Request.URL.Query(),
			Headers:     m.Request.Header,
			GinCtx:      m,
		}
		resp, err := handler(req)
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

// CreateOfferBasedOnProductAction is a high-level convenience wrapper around CreateOfferBasedOnProductActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func CreateOfferBasedOnProductActionGin(r gin.IRoutes, handler CreateOfferBasedOnProductActionRequestSig) {
	method, url, h := CreateOfferBasedOnProductActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Create offer based on productAction
 */
// Query wrapper with private fields
type CreateOfferBasedOnProductActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func CreateOfferBasedOnProductActionQueryFromString(rawQuery string) CreateOfferBasedOnProductActionQuery {
	v := CreateOfferBasedOnProductActionQuery{}
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
func CreateOfferBasedOnProductActionQueryFromGin(c *gin.Context) CreateOfferBasedOnProductActionQuery {
	return CreateOfferBasedOnProductActionQueryFromString(c.Request.URL.RawQuery)
}
func CreateOfferBasedOnProductActionQueryFromHttp(r *http.Request) CreateOfferBasedOnProductActionQuery {
	return CreateOfferBasedOnProductActionQueryFromString(r.URL.RawQuery)
}
func (q CreateOfferBasedOnProductActionQuery) Values() url.Values {
	return q.values
}
func (q CreateOfferBasedOnProductActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *CreateOfferBasedOnProductActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *CreateOfferBasedOnProductActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type CreateOfferBasedOnProductActionRequest struct {
	Body        CreateOfferBasedOnProductActionReq
	QueryParams url.Values
	// Automatically casted headers, for purpose of typesafe headers in later versions
	Headers http.Header
	// Gin context for each request in case of a direct access requirement
	GinCtx *gin.Context
	// Urfave context, per each request
	CliCtx *cli.Command
	// Reference to the application instance, in such scenarios that entire
	// application is wrapped into a single struct that holds database connection,
	// routes, etc.
	Application interface{}
}

func (x CreateOfferBasedOnProductActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x CreateOfferBasedOnProductActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type CreateOfferBasedOnProductActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func CreateOfferBasedOnProductActionClientCreateUrl(
	req CreateOfferBasedOnProductActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := CreateOfferBasedOnProductActionMeta()
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
func CreateOfferBasedOnProductActionClientExecuteTyped(httpReq *http.Request) (*CreateOfferBasedOnProductActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result CreateOfferBasedOnProductActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CreateOfferBasedOnProductActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &CreateOfferBasedOnProductActionResponse{Payload: result}, err
	}
	return &CreateOfferBasedOnProductActionResponse{Payload: result}, nil
}
func CreateOfferBasedOnProductActionClientBuildRequest(req CreateOfferBasedOnProductActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := CreateOfferBasedOnProductActionMeta()
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
func CreateOfferBasedOnProductActionCall(
	req CreateOfferBasedOnProductActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*CreateOfferBasedOnProductActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := CreateOfferBasedOnProductActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := CreateOfferBasedOnProductActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return CreateOfferBasedOnProductActionClientExecuteTyped(r)
}
