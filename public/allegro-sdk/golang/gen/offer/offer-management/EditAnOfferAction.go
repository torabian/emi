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
func GetEditAnOfferActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
		{
			Name: prefix + "language",
			Type: "string",
		},
		{
			Name:     prefix + "category",
			Type:     "object",
			Children: GetEditAnOfferActionReqCategoryCliFlags("category-"),
		},
		{
			Name: prefix + "product-set",
			Type: "array",
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetEditAnOfferActionReqStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetEditAnOfferActionReqSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "payments",
			Type:     "object",
			Children: GetEditAnOfferActionReqPaymentsCliFlags("payments-"),
		},
		{
			Name:     prefix + "delivery",
			Type:     "object",
			Children: GetEditAnOfferActionReqDeliveryCliFlags("delivery-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetEditAnOfferActionReqPublicationCliFlags("publication-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "map",
		},
		{
			Name:     prefix + "compatibility-list",
			Type:     "object",
			Children: GetEditAnOfferActionReqCompatibilityListCliFlags("compatibility-list-"),
		},
		{
			Name: prefix + "images",
			Type: "slice",
		},
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetEditAnOfferActionReqDescriptionCliFlags("description-"),
		},
		{
			Name:     prefix + "b2b",
			Type:     "object",
			Children: GetEditAnOfferActionReqB2bCliFlags("b2b-"),
		},
		{
			Name: prefix + "attachments",
			Type: "array",
		},
		{
			Name:     prefix + "fundraising-campaign",
			Type:     "object",
			Children: GetEditAnOfferActionReqFundraisingCampaignCliFlags("fundraising-campaign-"),
		},
		{
			Name:     prefix + "additional-services",
			Type:     "object",
			Children: GetEditAnOfferActionReqAdditionalServicesCliFlags("additional-services-"),
		},
		{
			Name:     prefix + "after-sales-services",
			Type:     "object",
			Children: GetEditAnOfferActionReqAfterSalesServicesCliFlags("after-sales-services-"),
		},
		{
			Name:     prefix + "size-table",
			Type:     "object",
			Children: GetEditAnOfferActionReqSizeTableCliFlags("size-table-"),
		},
		{
			Name:     prefix + "contact",
			Type:     "object",
			Children: GetEditAnOfferActionReqContactCliFlags("contact-"),
		},
		{
			Name:     prefix + "discounts",
			Type:     "object",
			Children: GetEditAnOfferActionReqDiscountsCliFlags("discounts-"),
		},
		{
			Name:     prefix + "location",
			Type:     "object",
			Children: GetEditAnOfferActionReqLocationCliFlags("location-"),
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetEditAnOfferActionReqExternalCliFlags("external-"),
		},
		{
			Name:     prefix + "tax-settings",
			Type:     "object",
			Children: GetEditAnOfferActionReqTaxSettingsCliFlags("tax-settings-"),
		},
		{
			Name:     prefix + "message-to-seller-settings",
			Type:     "object",
			Children: GetEditAnOfferActionReqMessageToSellerSettingsCliFlags("message-to-seller-settings-"),
		},
	}
}
func CastEditAnOfferActionReqFromCli(c emigo.CliCastable) EditAnOfferActionReq {
	data := EditAnOfferActionReq{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("language") {
		data.Language = c.String("language")
	}
	if c.IsSet("category") {
		data.Category = CastEditAnOfferActionReqCategoryFromCli(c)
	}
	if c.IsSet("product-set") {
		data.ProductSet = emigo.CapturePossibleArray(CastEditAnOfferActionReqProductSetFromCli, "product-set", c)
	}
	if c.IsSet("stock") {
		data.Stock = CastEditAnOfferActionReqStockFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastEditAnOfferActionReqSellingModeFromCli(c)
	}
	if c.IsSet("payments") {
		data.Payments = CastEditAnOfferActionReqPaymentsFromCli(c)
	}
	if c.IsSet("delivery") {
		data.Delivery = CastEditAnOfferActionReqDeliveryFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastEditAnOfferActionReqPublicationFromCli(c)
	}
	if c.IsSet("compatibility-list") {
		data.CompatibilityList = CastEditAnOfferActionReqCompatibilityListFromCli(c)
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	if c.IsSet("description") {
		data.Description = CastEditAnOfferActionReqDescriptionFromCli(c)
	}
	if c.IsSet("b2b") {
		data.B2b = CastEditAnOfferActionReqB2bFromCli(c)
	}
	if c.IsSet("attachments") {
		data.Attachments = emigo.CapturePossibleArray(CastEditAnOfferActionReqAttachmentsFromCli, "attachments", c)
	}
	if c.IsSet("fundraising-campaign") {
		data.FundraisingCampaign = CastEditAnOfferActionReqFundraisingCampaignFromCli(c)
	}
	if c.IsSet("additional-services") {
		data.AdditionalServices = CastEditAnOfferActionReqAdditionalServicesFromCli(c)
	}
	if c.IsSet("after-sales-services") {
		data.AfterSalesServices = CastEditAnOfferActionReqAfterSalesServicesFromCli(c)
	}
	if c.IsSet("size-table") {
		data.SizeTable = CastEditAnOfferActionReqSizeTableFromCli(c)
	}
	if c.IsSet("contact") {
		data.Contact = CastEditAnOfferActionReqContactFromCli(c)
	}
	if c.IsSet("discounts") {
		data.Discounts = CastEditAnOfferActionReqDiscountsFromCli(c)
	}
	if c.IsSet("location") {
		data.Location = CastEditAnOfferActionReqLocationFromCli(c)
	}
	if c.IsSet("external") {
		data.External = CastEditAnOfferActionReqExternalFromCli(c)
	}
	if c.IsSet("tax-settings") {
		data.TaxSettings = CastEditAnOfferActionReqTaxSettingsFromCli(c)
	}
	if c.IsSet("message-to-seller-settings") {
		data.MessageToSellerSettings = CastEditAnOfferActionReqMessageToSellerSettingsFromCli(c)
	}
	return data
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
	AdditionalMarketplaces  map[any]any                                 `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
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

func GetEditAnOfferActionReqCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqCategoryFromCli(c emigo.CliCastable) EditAnOfferActionReqCategory {
	data := EditAnOfferActionReqCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type EditAnOfferActionReqCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionReqProductSetCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "product",
			Type:     "object",
			Children: GetEditAnOfferActionReqProductSetProductCliFlags("product-"),
		},
		{
			Name:     prefix + "quantity",
			Type:     "object",
			Children: GetEditAnOfferActionReqProductSetQuantityCliFlags("quantity-"),
		},
		{
			Name:     prefix + "responsible-person",
			Type:     "object",
			Children: GetEditAnOfferActionReqProductSetResponsiblePersonCliFlags("responsible-person-"),
		},
		{
			Name:     prefix + "responsible-producer",
			Type:     "object",
			Children: GetEditAnOfferActionReqProductSetResponsibleProducerCliFlags("responsible-producer-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetEditAnOfferActionReqProductSetSafetyInformationCliFlags("safety-information-"),
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
func CastEditAnOfferActionReqProductSetFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSet {
	data := EditAnOfferActionReqProductSet{}
	if c.IsSet("product") {
		data.Product = CastEditAnOfferActionReqProductSetProductFromCli(c)
	}
	if c.IsSet("quantity") {
		data.Quantity = CastEditAnOfferActionReqProductSetQuantityFromCli(c)
	}
	if c.IsSet("responsible-person") {
		data.ResponsiblePerson = CastEditAnOfferActionReqProductSetResponsiblePersonFromCli(c)
	}
	if c.IsSet("responsible-producer") {
		data.ResponsibleProducer = CastEditAnOfferActionReqProductSetResponsibleProducerFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastEditAnOfferActionReqProductSetSafetyInformationFromCli(c)
	}
	if c.IsSet("marketed-before-gpsr-obligation") {
		data.MarketedBeforeGPSRObligation = bool(c.Bool("marketed-before-gpsr-obligation"))
	}
	if c.IsSet("deposits") {
		data.Deposits = emigo.CapturePossibleArray(CastEditAnOfferActionReqProductSetDepositsFromCli, "deposits", c)
	}
	return data
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

func GetEditAnOfferActionReqProductSetProductCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetEditAnOfferActionReqProductSetProductCategoryCliFlags("category-"),
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
func CastEditAnOfferActionReqProductSetProductFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetProduct {
	data := EditAnOfferActionReqProductSetProduct{}
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
		data.Category = CastEditAnOfferActionReqProductSetProductCategoryFromCli(c)
	}
	if c.IsSet("parameters") {
		data.Parameters = emigo.CapturePossibleArray(CastEditAnOfferActionReqProductSetProductParametersFromCli, "parameters", c)
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	return data
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

func GetEditAnOfferActionReqProductSetProductCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqProductSetProductCategoryFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetProductCategory {
	data := EditAnOfferActionReqProductSetProductCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type EditAnOfferActionReqProductSetProductCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionReqProductSetProductParametersCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetEditAnOfferActionReqProductSetProductParametersRangeValueCliFlags("range-value-"),
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
func CastEditAnOfferActionReqProductSetProductParametersFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetProductParameters {
	data := EditAnOfferActionReqProductSetProductParameters{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("range-value") {
		data.RangeValue = CastEditAnOfferActionReqProductSetProductParametersRangeValueFromCli(c)
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
type EditAnOfferActionReqProductSetProductParameters struct {
	Id         string                                                    `json:"id" yaml:"id"`
	Name       string                                                    `json:"name" yaml:"name"`
	RangeValue EditAnOfferActionReqProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                  `json:"values" yaml:"values"`
	ValuesIds  []string                                                  `json:"valuesIds" yaml:"valuesIds"`
}

func GetEditAnOfferActionReqProductSetProductParametersRangeValueCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqProductSetProductParametersRangeValueFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetProductParametersRangeValue {
	data := EditAnOfferActionReqProductSetProductParametersRangeValue{}
	if c.IsSet("from") {
		data.From = c.String("from")
	}
	if c.IsSet("to") {
		data.To = c.String("to")
	}
	return data
}

// The base class definition for rangeValue
type EditAnOfferActionReqProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

func GetEditAnOfferActionReqProductSetQuantityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "int",
		},
	}
}
func CastEditAnOfferActionReqProductSetQuantityFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetQuantity {
	data := EditAnOfferActionReqProductSetQuantity{}
	if c.IsSet("value") {
		data.Value = int(c.Int64("value"))
	}
	return data
}

// The base class definition for quantity
type EditAnOfferActionReqProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

func GetEditAnOfferActionReqProductSetResponsiblePersonCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqProductSetResponsiblePersonFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetResponsiblePerson {
	data := EditAnOfferActionReqProductSetResponsiblePerson{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for responsiblePerson
type EditAnOfferActionReqProductSetResponsiblePerson struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqProductSetResponsibleProducerCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqProductSetResponsibleProducerFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetResponsibleProducer {
	data := EditAnOfferActionReqProductSetResponsibleProducer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for responsibleProducer
type EditAnOfferActionReqProductSetResponsibleProducer struct {
	Id   string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"`
}

func GetEditAnOfferActionReqProductSetSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqProductSetSafetyInformationFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetSafetyInformation {
	data := EditAnOfferActionReqProductSetSafetyInformation{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

// The base class definition for safetyInformation
type EditAnOfferActionReqProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

func GetEditAnOfferActionReqProductSetDepositsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqProductSetDepositsFromCli(c emigo.CliCastable) EditAnOfferActionReqProductSetDeposits {
	data := EditAnOfferActionReqProductSetDeposits{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("quantity") {
		data.Quantity = int(c.Int64("quantity"))
	}
	return data
}

// The base class definition for deposits
type EditAnOfferActionReqProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

func GetEditAnOfferActionReqStockCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqStockFromCli(c emigo.CliCastable) EditAnOfferActionReqStock {
	data := EditAnOfferActionReqStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	if c.IsSet("unit") {
		data.Unit = c.String("unit")
	}
	return data
}

// The base class definition for stock
type EditAnOfferActionReqStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

func GetEditAnOfferActionReqSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "format",
			Type: "string",
		},
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetEditAnOfferActionReqSellingModePriceCliFlags("price-"),
		},
		{
			Name:     prefix + "minimal-price",
			Type:     "object",
			Children: GetEditAnOfferActionReqSellingModeMinimalPriceCliFlags("minimal-price-"),
		},
		{
			Name:     prefix + "starting-price",
			Type:     "object",
			Children: GetEditAnOfferActionReqSellingModeStartingPriceCliFlags("starting-price-"),
		},
	}
}
func CastEditAnOfferActionReqSellingModeFromCli(c emigo.CliCastable) EditAnOfferActionReqSellingMode {
	data := EditAnOfferActionReqSellingMode{}
	if c.IsSet("format") {
		data.Format = c.String("format")
	}
	if c.IsSet("price") {
		data.Price = CastEditAnOfferActionReqSellingModePriceFromCli(c)
	}
	if c.IsSet("minimal-price") {
		data.MinimalPrice = CastEditAnOfferActionReqSellingModeMinimalPriceFromCli(c)
	}
	if c.IsSet("starting-price") {
		data.StartingPrice = CastEditAnOfferActionReqSellingModeStartingPriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type EditAnOfferActionReqSellingMode struct {
	Format        string                                       `json:"format" yaml:"format"`
	Price         EditAnOfferActionReqSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  EditAnOfferActionReqSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice EditAnOfferActionReqSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

func GetEditAnOfferActionReqSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqSellingModePriceFromCli(c emigo.CliCastable) EditAnOfferActionReqSellingModePrice {
	data := EditAnOfferActionReqSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type EditAnOfferActionReqSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionReqSellingModeMinimalPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqSellingModeMinimalPriceFromCli(c emigo.CliCastable) EditAnOfferActionReqSellingModeMinimalPrice {
	data := EditAnOfferActionReqSellingModeMinimalPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for minimalPrice
type EditAnOfferActionReqSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionReqSellingModeStartingPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqSellingModeStartingPriceFromCli(c emigo.CliCastable) EditAnOfferActionReqSellingModeStartingPrice {
	data := EditAnOfferActionReqSellingModeStartingPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for startingPrice
type EditAnOfferActionReqSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionReqPaymentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "invoice",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqPaymentsFromCli(c emigo.CliCastable) EditAnOfferActionReqPayments {
	data := EditAnOfferActionReqPayments{}
	if c.IsSet("invoice") {
		data.Invoice = c.String("invoice")
	}
	return data
}

// The base class definition for payments
type EditAnOfferActionReqPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

func GetEditAnOfferActionReqDeliveryCliFlags(prefix string) []emigo.CliFlag {
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
			Name:     prefix + "shipping-rates",
			Type:     "object",
			Children: GetEditAnOfferActionReqDeliveryShippingRatesCliFlags("shipping-rates-"),
		},
	}
}
func CastEditAnOfferActionReqDeliveryFromCli(c emigo.CliCastable) EditAnOfferActionReqDelivery {
	data := EditAnOfferActionReqDelivery{}
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
		data.ShippingRates = CastEditAnOfferActionReqDeliveryShippingRatesFromCli(c)
	}
	return data
}

// The base class definition for delivery
type EditAnOfferActionReqDelivery struct {
	HandlingTime   string                                    `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                    `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                    `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  EditAnOfferActionReqDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

func GetEditAnOfferActionReqDeliveryShippingRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqDeliveryShippingRatesFromCli(c emigo.CliCastable) EditAnOfferActionReqDeliveryShippingRates {
	data := EditAnOfferActionReqDeliveryShippingRates{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for shippingRates
type EditAnOfferActionReqDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionReqPublicationCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqPublicationFromCli(c emigo.CliCastable) EditAnOfferActionReqPublication {
	data := EditAnOfferActionReqPublication{}
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
type EditAnOfferActionReqPublication struct {
	Duration   string `json:"duration" yaml:"duration"`
	StartingAt string `json:"startingAt" yaml:"startingAt"`
	EndingAt   string `json:"endingAt" yaml:"endingAt"`
	Status     string `json:"status" yaml:"status"`
	Republish  bool   `json:"republish" yaml:"republish"`
}

func GetEditAnOfferActionReqCompatibilityListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionReqCompatibilityListFromCli(c emigo.CliCastable) EditAnOfferActionReqCompatibilityList {
	data := EditAnOfferActionReqCompatibilityList{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEditAnOfferActionReqCompatibilityListItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for compatibilityList
type EditAnOfferActionReqCompatibilityList struct {
	Items []EditAnOfferActionReqCompatibilityListItems `json:"items" yaml:"items"`
}

func GetEditAnOfferActionReqCompatibilityListItemsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqCompatibilityListItemsFromCli(c emigo.CliCastable) EditAnOfferActionReqCompatibilityListItems {
	data := EditAnOfferActionReqCompatibilityListItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("text") {
		data.Text = c.String("text")
	}
	return data
}

// The base class definition for items
type EditAnOfferActionReqCompatibilityListItems struct {
	Type string `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
}

func GetEditAnOfferActionReqDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionReqDescriptionFromCli(c emigo.CliCastable) EditAnOfferActionReqDescription {
	data := EditAnOfferActionReqDescription{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastEditAnOfferActionReqDescriptionSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for description
type EditAnOfferActionReqDescription struct {
	Sections []EditAnOfferActionReqDescriptionSections `json:"sections" yaml:"sections"`
}

func GetEditAnOfferActionReqDescriptionSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionReqDescriptionSectionsFromCli(c emigo.CliCastable) EditAnOfferActionReqDescriptionSections {
	data := EditAnOfferActionReqDescriptionSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEditAnOfferActionReqDescriptionSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type EditAnOfferActionReqDescriptionSections struct {
	Items []EditAnOfferActionReqDescriptionSectionsItems `json:"items" yaml:"items"`
}

func GetEditAnOfferActionReqDescriptionSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqDescriptionSectionsItemsFromCli(c emigo.CliCastable) EditAnOfferActionReqDescriptionSectionsItems {
	data := EditAnOfferActionReqDescriptionSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for items
type EditAnOfferActionReqDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func GetEditAnOfferActionReqB2bCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "buyable-only-by-business",
			Type: "bool",
		},
	}
}
func CastEditAnOfferActionReqB2bFromCli(c emigo.CliCastable) EditAnOfferActionReqB2b {
	data := EditAnOfferActionReqB2b{}
	if c.IsSet("buyable-only-by-business") {
		data.BuyableOnlyByBusiness = bool(c.Bool("buyable-only-by-business"))
	}
	return data
}

// The base class definition for b2b
type EditAnOfferActionReqB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

func GetEditAnOfferActionReqAttachmentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqAttachmentsFromCli(c emigo.CliCastable) EditAnOfferActionReqAttachments {
	data := EditAnOfferActionReqAttachments{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for attachments
type EditAnOfferActionReqAttachments struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionReqFundraisingCampaignCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqFundraisingCampaignFromCli(c emigo.CliCastable) EditAnOfferActionReqFundraisingCampaign {
	data := EditAnOfferActionReqFundraisingCampaign{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for fundraisingCampaign
type EditAnOfferActionReqFundraisingCampaign struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqAdditionalServicesCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqAdditionalServicesFromCli(c emigo.CliCastable) EditAnOfferActionReqAdditionalServices {
	data := EditAnOfferActionReqAdditionalServices{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for additionalServices
type EditAnOfferActionReqAdditionalServices struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqAfterSalesServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "implied-warranty",
			Type:     "object",
			Children: GetEditAnOfferActionReqAfterSalesServicesImpliedWarrantyCliFlags("implied-warranty-"),
		},
		{
			Name:     prefix + "return-policy",
			Type:     "object",
			Children: GetEditAnOfferActionReqAfterSalesServicesReturnPolicyCliFlags("return-policy-"),
		},
		{
			Name:     prefix + "warranty",
			Type:     "object",
			Children: GetEditAnOfferActionReqAfterSalesServicesWarrantyCliFlags("warranty-"),
		},
	}
}
func CastEditAnOfferActionReqAfterSalesServicesFromCli(c emigo.CliCastable) EditAnOfferActionReqAfterSalesServices {
	data := EditAnOfferActionReqAfterSalesServices{}
	if c.IsSet("implied-warranty") {
		data.ImpliedWarranty = CastEditAnOfferActionReqAfterSalesServicesImpliedWarrantyFromCli(c)
	}
	if c.IsSet("return-policy") {
		data.ReturnPolicy = CastEditAnOfferActionReqAfterSalesServicesReturnPolicyFromCli(c)
	}
	if c.IsSet("warranty") {
		data.Warranty = CastEditAnOfferActionReqAfterSalesServicesWarrantyFromCli(c)
	}
	return data
}

// The base class definition for afterSalesServices
type EditAnOfferActionReqAfterSalesServices struct {
	ImpliedWarranty EditAnOfferActionReqAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    EditAnOfferActionReqAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        EditAnOfferActionReqAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

func GetEditAnOfferActionReqAfterSalesServicesImpliedWarrantyCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqAfterSalesServicesImpliedWarrantyFromCli(c emigo.CliCastable) EditAnOfferActionReqAfterSalesServicesImpliedWarranty {
	data := EditAnOfferActionReqAfterSalesServicesImpliedWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for impliedWarranty
type EditAnOfferActionReqAfterSalesServicesImpliedWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqAfterSalesServicesReturnPolicyCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqAfterSalesServicesReturnPolicyFromCli(c emigo.CliCastable) EditAnOfferActionReqAfterSalesServicesReturnPolicy {
	data := EditAnOfferActionReqAfterSalesServicesReturnPolicy{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for returnPolicy
type EditAnOfferActionReqAfterSalesServicesReturnPolicy struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqAfterSalesServicesWarrantyCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqAfterSalesServicesWarrantyFromCli(c emigo.CliCastable) EditAnOfferActionReqAfterSalesServicesWarranty {
	data := EditAnOfferActionReqAfterSalesServicesWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for warranty
type EditAnOfferActionReqAfterSalesServicesWarranty struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqSizeTableCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqSizeTableFromCli(c emigo.CliCastable) EditAnOfferActionReqSizeTable {
	data := EditAnOfferActionReqSizeTable{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for sizeTable
type EditAnOfferActionReqSizeTable struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqContactCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqContactFromCli(c emigo.CliCastable) EditAnOfferActionReqContact {
	data := EditAnOfferActionReqContact{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for contact
type EditAnOfferActionReqContact struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqDiscountsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "wholesale-price-list",
			Type:     "object",
			Children: GetEditAnOfferActionReqDiscountsWholesalePriceListCliFlags("wholesale-price-list-"),
		},
	}
}
func CastEditAnOfferActionReqDiscountsFromCli(c emigo.CliCastable) EditAnOfferActionReqDiscounts {
	data := EditAnOfferActionReqDiscounts{}
	if c.IsSet("wholesale-price-list") {
		data.WholesalePriceList = CastEditAnOfferActionReqDiscountsWholesalePriceListFromCli(c)
	}
	return data
}

// The base class definition for discounts
type EditAnOfferActionReqDiscounts struct {
	WholesalePriceList EditAnOfferActionReqDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

func GetEditAnOfferActionReqDiscountsWholesalePriceListCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqDiscountsWholesalePriceListFromCli(c emigo.CliCastable) EditAnOfferActionReqDiscountsWholesalePriceList {
	data := EditAnOfferActionReqDiscountsWholesalePriceList{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}

// The base class definition for wholesalePriceList
type EditAnOfferActionReqDiscountsWholesalePriceList struct {
	Id   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func GetEditAnOfferActionReqLocationCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqLocationFromCli(c emigo.CliCastable) EditAnOfferActionReqLocation {
	data := EditAnOfferActionReqLocation{}
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
type EditAnOfferActionReqLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

func GetEditAnOfferActionReqExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionReqExternalFromCli(c emigo.CliCastable) EditAnOfferActionReqExternal {
	data := EditAnOfferActionReqExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type EditAnOfferActionReqExternal struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionReqTaxSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqTaxSettingsFromCli(c emigo.CliCastable) EditAnOfferActionReqTaxSettings {
	data := EditAnOfferActionReqTaxSettings{}
	if c.IsSet("subject") {
		data.Subject = c.String("subject")
	}
	if c.IsSet("exemption") {
		data.Exemption = c.String("exemption")
	}
	if c.IsSet("rates") {
		data.Rates = emigo.CapturePossibleArray(CastEditAnOfferActionReqTaxSettingsRatesFromCli, "rates", c)
	}
	return data
}

// The base class definition for taxSettings
type EditAnOfferActionReqTaxSettings struct {
	Subject   string                                 `json:"subject" yaml:"subject"`
	Exemption string                                 `json:"exemption" yaml:"exemption"`
	Rates     []EditAnOfferActionReqTaxSettingsRates `json:"rates" yaml:"rates"`
}

func GetEditAnOfferActionReqTaxSettingsRatesCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqTaxSettingsRatesFromCli(c emigo.CliCastable) EditAnOfferActionReqTaxSettingsRates {
	data := EditAnOfferActionReqTaxSettingsRates{}
	if c.IsSet("rate") {
		data.Rate = c.String("rate")
	}
	if c.IsSet("country-code") {
		data.CountryCode = c.String("country-code")
	}
	return data
}

// The base class definition for rates
type EditAnOfferActionReqTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

func GetEditAnOfferActionReqMessageToSellerSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionReqMessageToSellerSettingsFromCli(c emigo.CliCastable) EditAnOfferActionReqMessageToSellerSettings {
	data := EditAnOfferActionReqMessageToSellerSettings{}
	if c.IsSet("mode") {
		data.Mode = c.String("mode")
	}
	if c.IsSet("hint") {
		data.Hint = c.String("hint")
	}
	return data
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
func GetEditAnOfferActionResCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "language",
			Type: "string",
		},
		{
			Name:     prefix + "category",
			Type:     "object",
			Children: GetEditAnOfferActionResCategoryCliFlags("category-"),
		},
		{
			Name: prefix + "product-set",
			Type: "array",
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetEditAnOfferActionResStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "payments",
			Type:     "object",
			Children: GetEditAnOfferActionResPaymentsCliFlags("payments-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetEditAnOfferActionResSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "delivery",
			Type:     "object",
			Children: GetEditAnOfferActionResDeliveryCliFlags("delivery-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetEditAnOfferActionResPublicationCliFlags("publication-"),
		},
		{
			Name:     prefix + "additional-marketplaces",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalMarketplacesCliFlags("additional-marketplaces-"),
		},
		{
			Name:     prefix + "b2b",
			Type:     "object",
			Children: GetEditAnOfferActionResB2bCliFlags("b2b-"),
		},
		{
			Name:     prefix + "compatibility-list",
			Type:     "object",
			Children: GetEditAnOfferActionResCompatibilityListCliFlags("compatibility-list-"),
		},
		{
			Name:     prefix + "validation",
			Type:     "object",
			Children: GetEditAnOfferActionResValidationCliFlags("validation-"),
		},
		{
			Name: prefix + "warnings",
			Type: "slice",
		},
		{
			Name:     prefix + "after-sales-services",
			Type:     "object",
			Children: GetEditAnOfferActionResAfterSalesServicesCliFlags("after-sales-services-"),
		},
		{
			Name:     prefix + "discounts",
			Type:     "object",
			Children: GetEditAnOfferActionResDiscountsCliFlags("discounts-"),
		},
		{
			Name:     prefix + "contact",
			Type:     "object",
			Children: GetEditAnOfferActionResContactCliFlags("contact-"),
		},
		{
			Name: prefix + "attachments",
			Type: "array",
		},
		{
			Name:     prefix + "fundraising-campaign",
			Type:     "object",
			Children: GetEditAnOfferActionResFundraisingCampaignCliFlags("fundraising-campaign-"),
		},
		{
			Name:     prefix + "additional-services",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalServicesCliFlags("additional-services-"),
		},
		{
			Name:     prefix + "size-table",
			Type:     "object",
			Children: GetEditAnOfferActionResSizeTableCliFlags("size-table-"),
		},
		{
			Name:     prefix + "location",
			Type:     "object",
			Children: GetEditAnOfferActionResLocationCliFlags("location-"),
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetEditAnOfferActionResExternalCliFlags("external-"),
		},
		{
			Name:     prefix + "tax-settings",
			Type:     "object",
			Children: GetEditAnOfferActionResTaxSettingsCliFlags("tax-settings-"),
		},
		{
			Name:     prefix + "message-to-seller-settings",
			Type:     "object",
			Children: GetEditAnOfferActionResMessageToSellerSettingsCliFlags("message-to-seller-settings-"),
		},
		{
			Name: prefix + "created-at",
			Type: "string",
		},
		{
			Name: prefix + "updated-at",
			Type: "string",
		},
		{
			Name: prefix + "images",
			Type: "slice",
		},
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetEditAnOfferActionResDescriptionCliFlags("description-"),
		},
	}
}
func CastEditAnOfferActionResFromCli(c emigo.CliCastable) EditAnOfferActionRes {
	data := EditAnOfferActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("language") {
		data.Language = c.String("language")
	}
	if c.IsSet("category") {
		data.Category = CastEditAnOfferActionResCategoryFromCli(c)
	}
	if c.IsSet("product-set") {
		data.ProductSet = emigo.CapturePossibleArray(CastEditAnOfferActionResProductSetFromCli, "product-set", c)
	}
	if c.IsSet("stock") {
		data.Stock = CastEditAnOfferActionResStockFromCli(c)
	}
	if c.IsSet("payments") {
		data.Payments = CastEditAnOfferActionResPaymentsFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastEditAnOfferActionResSellingModeFromCli(c)
	}
	if c.IsSet("delivery") {
		data.Delivery = CastEditAnOfferActionResDeliveryFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastEditAnOfferActionResPublicationFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = CastEditAnOfferActionResAdditionalMarketplacesFromCli(c)
	}
	if c.IsSet("b2b") {
		data.B2b = CastEditAnOfferActionResB2bFromCli(c)
	}
	if c.IsSet("compatibility-list") {
		data.CompatibilityList = CastEditAnOfferActionResCompatibilityListFromCli(c)
	}
	if c.IsSet("validation") {
		data.Validation = CastEditAnOfferActionResValidationFromCli(c)
	}
	if c.IsSet("warnings") {
		emigo.InflatePossibleSlice(c.String("warnings"), &data.Warnings)
	}
	if c.IsSet("after-sales-services") {
		data.AfterSalesServices = CastEditAnOfferActionResAfterSalesServicesFromCli(c)
	}
	if c.IsSet("discounts") {
		data.Discounts = CastEditAnOfferActionResDiscountsFromCli(c)
	}
	if c.IsSet("contact") {
		data.Contact = CastEditAnOfferActionResContactFromCli(c)
	}
	if c.IsSet("attachments") {
		data.Attachments = emigo.CapturePossibleArray(CastEditAnOfferActionResAttachmentsFromCli, "attachments", c)
	}
	if c.IsSet("fundraising-campaign") {
		data.FundraisingCampaign = CastEditAnOfferActionResFundraisingCampaignFromCli(c)
	}
	if c.IsSet("additional-services") {
		data.AdditionalServices = CastEditAnOfferActionResAdditionalServicesFromCli(c)
	}
	if c.IsSet("size-table") {
		data.SizeTable = CastEditAnOfferActionResSizeTableFromCli(c)
	}
	if c.IsSet("location") {
		data.Location = CastEditAnOfferActionResLocationFromCli(c)
	}
	if c.IsSet("external") {
		data.External = CastEditAnOfferActionResExternalFromCli(c)
	}
	if c.IsSet("tax-settings") {
		data.TaxSettings = CastEditAnOfferActionResTaxSettingsFromCli(c)
	}
	if c.IsSet("message-to-seller-settings") {
		data.MessageToSellerSettings = CastEditAnOfferActionResMessageToSellerSettingsFromCli(c)
	}
	if c.IsSet("created-at") {
		data.CreatedAt = c.String("created-at")
	}
	if c.IsSet("updated-at") {
		data.UpdatedAt = c.String("updated-at")
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	if c.IsSet("description") {
		data.Description = CastEditAnOfferActionResDescriptionFromCli(c)
	}
	return data
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

func GetEditAnOfferActionResCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResCategoryFromCli(c emigo.CliCastable) EditAnOfferActionResCategory {
	data := EditAnOfferActionResCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type EditAnOfferActionResCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResProductSetCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "quantity",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetQuantityCliFlags("quantity-"),
		},
		{
			Name:     prefix + "product",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetProductCliFlags("product-"),
		},
		{
			Name:     prefix + "responsible-person",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetResponsiblePersonCliFlags("responsible-person-"),
		},
		{
			Name:     prefix + "responsible-producer",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetResponsibleProducerCliFlags("responsible-producer-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetSafetyInformationCliFlags("safety-information-"),
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
func CastEditAnOfferActionResProductSetFromCli(c emigo.CliCastable) EditAnOfferActionResProductSet {
	data := EditAnOfferActionResProductSet{}
	if c.IsSet("quantity") {
		data.Quantity = CastEditAnOfferActionResProductSetQuantityFromCli(c)
	}
	if c.IsSet("product") {
		data.Product = CastEditAnOfferActionResProductSetProductFromCli(c)
	}
	if c.IsSet("responsible-person") {
		data.ResponsiblePerson = CastEditAnOfferActionResProductSetResponsiblePersonFromCli(c)
	}
	if c.IsSet("responsible-producer") {
		data.ResponsibleProducer = CastEditAnOfferActionResProductSetResponsibleProducerFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastEditAnOfferActionResProductSetSafetyInformationFromCli(c)
	}
	if c.IsSet("marketed-before-gpsr-obligation") {
		data.MarketedBeforeGPSRObligation = bool(c.Bool("marketed-before-gpsr-obligation"))
	}
	if c.IsSet("deposits") {
		data.Deposits = emigo.CapturePossibleArray(CastEditAnOfferActionResProductSetDepositsFromCli, "deposits", c)
	}
	return data
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

func GetEditAnOfferActionResProductSetQuantityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "int",
		},
	}
}
func CastEditAnOfferActionResProductSetQuantityFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetQuantity {
	data := EditAnOfferActionResProductSetQuantity{}
	if c.IsSet("value") {
		data.Value = int(c.Int64("value"))
	}
	return data
}

// The base class definition for quantity
type EditAnOfferActionResProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

func GetEditAnOfferActionResProductSetProductCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "is-ai-co-created",
			Type: "bool",
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetEditAnOfferActionResProductSetProductPublicationCliFlags("publication-"),
		},
		{
			Name: prefix + "parameters",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResProductSetProductFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetProduct {
	data := EditAnOfferActionResProductSetProduct{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("is-ai-co-created") {
		data.IsAiCoCreated = bool(c.Bool("is-ai-co-created"))
	}
	if c.IsSet("publication") {
		data.Publication = CastEditAnOfferActionResProductSetProductPublicationFromCli(c)
	}
	if c.IsSet("parameters") {
		data.Parameters = emigo.CapturePossibleArray(CastEditAnOfferActionResProductSetProductParametersFromCli, "parameters", c)
	}
	return data
}

// The base class definition for product
type EditAnOfferActionResProductSetProduct struct {
	Id            string                                            `json:"id" yaml:"id"`
	IsAiCoCreated bool                                              `json:"isAiCoCreated" yaml:"isAiCoCreated"`
	Publication   EditAnOfferActionResProductSetProductPublication  `json:"publication" yaml:"publication"`
	Parameters    []EditAnOfferActionResProductSetProductParameters `json:"parameters" yaml:"parameters"`
}

func GetEditAnOfferActionResProductSetProductPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "status",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResProductSetProductPublicationFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetProductPublication {
	data := EditAnOfferActionResProductSetProductPublication{}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	return data
}

// The base class definition for publication
type EditAnOfferActionResProductSetProductPublication struct {
	Status string `json:"status" yaml:"status"`
}

func GetEditAnOfferActionResProductSetProductParametersCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetEditAnOfferActionResProductSetProductParametersRangeValueCliFlags("range-value-"),
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
func CastEditAnOfferActionResProductSetProductParametersFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetProductParameters {
	data := EditAnOfferActionResProductSetProductParameters{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("range-value") {
		data.RangeValue = CastEditAnOfferActionResProductSetProductParametersRangeValueFromCli(c)
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
type EditAnOfferActionResProductSetProductParameters struct {
	Id         string                                                    `json:"id" yaml:"id"`
	Name       string                                                    `json:"name" yaml:"name"`
	RangeValue EditAnOfferActionResProductSetProductParametersRangeValue `json:"rangeValue" yaml:"rangeValue"`
	Values     []string                                                  `json:"values" yaml:"values"`
	ValuesIds  []string                                                  `json:"valuesIds" yaml:"valuesIds"`
}

func GetEditAnOfferActionResProductSetProductParametersRangeValueCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResProductSetProductParametersRangeValueFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetProductParametersRangeValue {
	data := EditAnOfferActionResProductSetProductParametersRangeValue{}
	if c.IsSet("from") {
		data.From = c.String("from")
	}
	if c.IsSet("to") {
		data.To = c.String("to")
	}
	return data
}

// The base class definition for rangeValue
type EditAnOfferActionResProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

func GetEditAnOfferActionResProductSetResponsiblePersonCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResProductSetResponsiblePersonFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetResponsiblePerson {
	data := EditAnOfferActionResProductSetResponsiblePerson{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for responsiblePerson
type EditAnOfferActionResProductSetResponsiblePerson struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResProductSetResponsibleProducerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResProductSetResponsibleProducerFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetResponsibleProducer {
	data := EditAnOfferActionResProductSetResponsibleProducer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for responsibleProducer
type EditAnOfferActionResProductSetResponsibleProducer struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResProductSetSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResProductSetSafetyInformationFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetSafetyInformation {
	data := EditAnOfferActionResProductSetSafetyInformation{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

// The base class definition for safetyInformation
type EditAnOfferActionResProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

func GetEditAnOfferActionResProductSetDepositsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResProductSetDepositsFromCli(c emigo.CliCastable) EditAnOfferActionResProductSetDeposits {
	data := EditAnOfferActionResProductSetDeposits{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("quantity") {
		data.Quantity = int(c.Int64("quantity"))
	}
	return data
}

// The base class definition for deposits
type EditAnOfferActionResProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

func GetEditAnOfferActionResStockCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResStockFromCli(c emigo.CliCastable) EditAnOfferActionResStock {
	data := EditAnOfferActionResStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	if c.IsSet("unit") {
		data.Unit = c.String("unit")
	}
	return data
}

// The base class definition for stock
type EditAnOfferActionResStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

func GetEditAnOfferActionResPaymentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "invoice",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResPaymentsFromCli(c emigo.CliCastable) EditAnOfferActionResPayments {
	data := EditAnOfferActionResPayments{}
	if c.IsSet("invoice") {
		data.Invoice = c.String("invoice")
	}
	return data
}

// The base class definition for payments
type EditAnOfferActionResPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

func GetEditAnOfferActionResSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "format",
			Type: "string",
		},
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetEditAnOfferActionResSellingModePriceCliFlags("price-"),
		},
		{
			Name:     prefix + "minimal-price",
			Type:     "object",
			Children: GetEditAnOfferActionResSellingModeMinimalPriceCliFlags("minimal-price-"),
		},
		{
			Name:     prefix + "starting-price",
			Type:     "object",
			Children: GetEditAnOfferActionResSellingModeStartingPriceCliFlags("starting-price-"),
		},
	}
}
func CastEditAnOfferActionResSellingModeFromCli(c emigo.CliCastable) EditAnOfferActionResSellingMode {
	data := EditAnOfferActionResSellingMode{}
	if c.IsSet("format") {
		data.Format = c.String("format")
	}
	if c.IsSet("price") {
		data.Price = CastEditAnOfferActionResSellingModePriceFromCli(c)
	}
	if c.IsSet("minimal-price") {
		data.MinimalPrice = CastEditAnOfferActionResSellingModeMinimalPriceFromCli(c)
	}
	if c.IsSet("starting-price") {
		data.StartingPrice = CastEditAnOfferActionResSellingModeStartingPriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type EditAnOfferActionResSellingMode struct {
	Format        string                                       `json:"format" yaml:"format"`
	Price         EditAnOfferActionResSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  EditAnOfferActionResSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice EditAnOfferActionResSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

func GetEditAnOfferActionResSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResSellingModePriceFromCli(c emigo.CliCastable) EditAnOfferActionResSellingModePrice {
	data := EditAnOfferActionResSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type EditAnOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionResSellingModeMinimalPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResSellingModeMinimalPriceFromCli(c emigo.CliCastable) EditAnOfferActionResSellingModeMinimalPrice {
	data := EditAnOfferActionResSellingModeMinimalPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for minimalPrice
type EditAnOfferActionResSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionResSellingModeStartingPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResSellingModeStartingPriceFromCli(c emigo.CliCastable) EditAnOfferActionResSellingModeStartingPrice {
	data := EditAnOfferActionResSellingModeStartingPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for startingPrice
type EditAnOfferActionResSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionResDeliveryCliFlags(prefix string) []emigo.CliFlag {
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
			Name:     prefix + "shipping-rates",
			Type:     "object",
			Children: GetEditAnOfferActionResDeliveryShippingRatesCliFlags("shipping-rates-"),
		},
	}
}
func CastEditAnOfferActionResDeliveryFromCli(c emigo.CliCastable) EditAnOfferActionResDelivery {
	data := EditAnOfferActionResDelivery{}
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
		data.ShippingRates = CastEditAnOfferActionResDeliveryShippingRatesFromCli(c)
	}
	return data
}

// The base class definition for delivery
type EditAnOfferActionResDelivery struct {
	HandlingTime   string                                    `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                    `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                    `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  EditAnOfferActionResDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

func GetEditAnOfferActionResDeliveryShippingRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResDeliveryShippingRatesFromCli(c emigo.CliCastable) EditAnOfferActionResDeliveryShippingRates {
	data := EditAnOfferActionResDeliveryShippingRates{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for shippingRates
type EditAnOfferActionResDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResPublicationCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "ended-by",
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
		{
			Name:     prefix + "marketplaces",
			Type:     "object",
			Children: GetEditAnOfferActionResPublicationMarketplacesCliFlags("marketplaces-"),
		},
	}
}
func CastEditAnOfferActionResPublicationFromCli(c emigo.CliCastable) EditAnOfferActionResPublication {
	data := EditAnOfferActionResPublication{}
	if c.IsSet("duration") {
		data.Duration = c.String("duration")
	}
	if c.IsSet("starting-at") {
		data.StartingAt = c.String("starting-at")
	}
	if c.IsSet("ending-at") {
		data.EndingAt = c.String("ending-at")
	}
	if c.IsSet("ended-by") {
		data.EndedBy = c.String("ended-by")
	}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("republish") {
		data.Republish = bool(c.Bool("republish"))
	}
	if c.IsSet("marketplaces") {
		data.Marketplaces = CastEditAnOfferActionResPublicationMarketplacesFromCli(c)
	}
	return data
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

func GetEditAnOfferActionResPublicationMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base",
			Type:     "object",
			Children: GetEditAnOfferActionResPublicationMarketplacesBaseCliFlags("base-"),
		},
		{
			Name: prefix + "additional",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResPublicationMarketplacesFromCli(c emigo.CliCastable) EditAnOfferActionResPublicationMarketplaces {
	data := EditAnOfferActionResPublicationMarketplaces{}
	if c.IsSet("base") {
		data.Base = CastEditAnOfferActionResPublicationMarketplacesBaseFromCli(c)
	}
	if c.IsSet("additional") {
		data.Additional = emigo.CapturePossibleArray(CastEditAnOfferActionResPublicationMarketplacesAdditionalFromCli, "additional", c)
	}
	return data
}

// The base class definition for marketplaces
type EditAnOfferActionResPublicationMarketplaces struct {
	Base       EditAnOfferActionResPublicationMarketplacesBase         `json:"base" yaml:"base"`
	Additional []EditAnOfferActionResPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
}

func GetEditAnOfferActionResPublicationMarketplacesBaseCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResPublicationMarketplacesBaseFromCli(c emigo.CliCastable) EditAnOfferActionResPublicationMarketplacesBase {
	data := EditAnOfferActionResPublicationMarketplacesBase{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for base
type EditAnOfferActionResPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResPublicationMarketplacesAdditionalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResPublicationMarketplacesAdditionalFromCli(c emigo.CliCastable) EditAnOfferActionResPublicationMarketplacesAdditional {
	data := EditAnOfferActionResPublicationMarketplacesAdditional{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additional
type EditAnOfferActionResPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalMarketplacesSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalMarketplacesPublicationCliFlags("publication-"),
		},
	}
}
func CastEditAnOfferActionResAdditionalMarketplacesFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplaces {
	data := EditAnOfferActionResAdditionalMarketplaces{}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastEditAnOfferActionResAdditionalMarketplacesSellingModeFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastEditAnOfferActionResAdditionalMarketplacesPublicationFromCli(c)
	}
	return data
}

// The base class definition for additionalMarketplaces
type EditAnOfferActionResAdditionalMarketplaces struct {
	SellingMode EditAnOfferActionResAdditionalMarketplacesSellingMode `json:"sellingMode" yaml:"sellingMode"`
	Publication EditAnOfferActionResAdditionalMarketplacesPublication `json:"publication" yaml:"publication"`
}

func GetEditAnOfferActionResAdditionalMarketplacesSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalMarketplacesSellingModePriceCliFlags("price-"),
		},
	}
}
func CastEditAnOfferActionResAdditionalMarketplacesSellingModeFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplacesSellingMode {
	data := EditAnOfferActionResAdditionalMarketplacesSellingMode{}
	if c.IsSet("price") {
		data.Price = CastEditAnOfferActionResAdditionalMarketplacesSellingModePriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type EditAnOfferActionResAdditionalMarketplacesSellingMode struct {
	Price EditAnOfferActionResAdditionalMarketplacesSellingModePrice `json:"price" yaml:"price"`
}

func GetEditAnOfferActionResAdditionalMarketplacesSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResAdditionalMarketplacesSellingModePriceFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplacesSellingModePrice {
	data := EditAnOfferActionResAdditionalMarketplacesSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type EditAnOfferActionResAdditionalMarketplacesSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetEditAnOfferActionResAdditionalMarketplacesPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "state",
			Type: "string",
		},
		{
			Name: prefix + "refusal-reasons",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResAdditionalMarketplacesPublicationFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplacesPublication {
	data := EditAnOfferActionResAdditionalMarketplacesPublication{}
	if c.IsSet("state") {
		data.State = c.String("state")
	}
	if c.IsSet("refusal-reasons") {
		data.RefusalReasons = emigo.CapturePossibleArray(CastEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsFromCli, "refusal-reasons", c)
	}
	return data
}

// The base class definition for publication
type EditAnOfferActionResAdditionalMarketplacesPublication struct {
	State          string                                                                `json:"state" yaml:"state"`
	RefusalReasons []EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons `json:"refusalReasons" yaml:"refusalReasons"`
}

func GetEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "code",
			Type: "string",
		},
		{
			Name: prefix + "user-message",
			Type: "string",
		},
		{
			Name:     prefix + "parameters",
			Type:     "object",
			Children: GetEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParametersCliFlags("parameters-"),
		},
	}
}
func CastEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons {
	data := EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons{}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("user-message") {
		data.UserMessage = c.String("user-message")
	}
	if c.IsSet("parameters") {
		data.Parameters = CastEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParametersFromCli(c)
	}
	return data
}

// The base class definition for refusalReasons
type EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasons struct {
	Code        string                                                                        `json:"code" yaml:"code"`
	UserMessage string                                                                        `json:"userMessage" yaml:"userMessage"`
	Parameters  EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters `json:"parameters" yaml:"parameters"`
}

func GetEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParametersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "max-allowed-price-decrease-percent",
			Type: "slice",
		},
	}
}
func CastEditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParametersFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters {
	data := EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters{}
	if c.IsSet("max-allowed-price-decrease-percent") {
		emigo.InflatePossibleSlice(c.String("max-allowed-price-decrease-percent"), &data.MaxAllowedPriceDecreasePercent)
	}
	return data
}

// The base class definition for parameters
type EditAnOfferActionResAdditionalMarketplacesPublicationRefusalReasonsParameters struct {
	MaxAllowedPriceDecreasePercent []string `json:"maxAllowedPriceDecreasePercent" yaml:"maxAllowedPriceDecreasePercent"`
}

func GetEditAnOfferActionResB2bCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "buyable-only-by-business",
			Type: "bool",
		},
	}
}
func CastEditAnOfferActionResB2bFromCli(c emigo.CliCastable) EditAnOfferActionResB2b {
	data := EditAnOfferActionResB2b{}
	if c.IsSet("buyable-only-by-business") {
		data.BuyableOnlyByBusiness = bool(c.Bool("buyable-only-by-business"))
	}
	return data
}

// The base class definition for b2b
type EditAnOfferActionResB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

func GetEditAnOfferActionResCompatibilityListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResCompatibilityListFromCli(c emigo.CliCastable) EditAnOfferActionResCompatibilityList {
	data := EditAnOfferActionResCompatibilityList{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for compatibilityList
type EditAnOfferActionResCompatibilityList struct {
	Type string `json:"type" yaml:"type"`
}

func GetEditAnOfferActionResValidationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "validated-at",
			Type: "string",
		},
		{
			Name: prefix + "errors",
			Type: "array",
		},
		{
			Name: prefix + "warnings",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResValidationFromCli(c emigo.CliCastable) EditAnOfferActionResValidation {
	data := EditAnOfferActionResValidation{}
	if c.IsSet("validated-at") {
		data.ValidatedAt = c.String("validated-at")
	}
	if c.IsSet("errors") {
		data.Errors = emigo.CapturePossibleArray(CastEditAnOfferActionResValidationErrorsFromCli, "errors", c)
	}
	if c.IsSet("warnings") {
		data.Warnings = emigo.CapturePossibleArray(CastEditAnOfferActionResValidationWarningsFromCli, "warnings", c)
	}
	return data
}

// The base class definition for validation
type EditAnOfferActionResValidation struct {
	ValidatedAt string                                   `json:"validatedAt" yaml:"validatedAt"`
	Errors      []EditAnOfferActionResValidationErrors   `json:"errors" yaml:"errors"`
	Warnings    []EditAnOfferActionResValidationWarnings `json:"warnings" yaml:"warnings"`
}

func GetEditAnOfferActionResValidationErrorsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "code",
			Type: "string",
		},
		{
			Name: prefix + "details",
			Type: "string",
		},
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name: prefix + "path",
			Type: "string",
		},
		{
			Name: prefix + "user-message",
			Type: "string",
		},
		{
			Name:     prefix + "metadata",
			Type:     "object",
			Children: GetEditAnOfferActionResValidationErrorsMetadataCliFlags("metadata-"),
		},
	}
}
func CastEditAnOfferActionResValidationErrorsFromCli(c emigo.CliCastable) EditAnOfferActionResValidationErrors {
	data := EditAnOfferActionResValidationErrors{}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("details") {
		data.Details = c.String("details")
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("path") {
		data.Path = c.String("path")
	}
	if c.IsSet("user-message") {
		data.UserMessage = c.String("user-message")
	}
	if c.IsSet("metadata") {
		data.Metadata = CastEditAnOfferActionResValidationErrorsMetadataFromCli(c)
	}
	return data
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

func GetEditAnOfferActionResValidationErrorsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResValidationErrorsMetadataFromCli(c emigo.CliCastable) EditAnOfferActionResValidationErrorsMetadata {
	data := EditAnOfferActionResValidationErrorsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type EditAnOfferActionResValidationErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func GetEditAnOfferActionResValidationWarningsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "code",
			Type: "string",
		},
		{
			Name: prefix + "details",
			Type: "string",
		},
		{
			Name: prefix + "message",
			Type: "string",
		},
		{
			Name: prefix + "path",
			Type: "string",
		},
		{
			Name: prefix + "user-message",
			Type: "string",
		},
		{
			Name:     prefix + "metadata",
			Type:     "object",
			Children: GetEditAnOfferActionResValidationWarningsMetadataCliFlags("metadata-"),
		},
	}
}
func CastEditAnOfferActionResValidationWarningsFromCli(c emigo.CliCastable) EditAnOfferActionResValidationWarnings {
	data := EditAnOfferActionResValidationWarnings{}
	if c.IsSet("code") {
		data.Code = c.String("code")
	}
	if c.IsSet("details") {
		data.Details = c.String("details")
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	if c.IsSet("path") {
		data.Path = c.String("path")
	}
	if c.IsSet("user-message") {
		data.UserMessage = c.String("user-message")
	}
	if c.IsSet("metadata") {
		data.Metadata = CastEditAnOfferActionResValidationWarningsMetadataFromCli(c)
	}
	return data
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

func GetEditAnOfferActionResValidationWarningsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResValidationWarningsMetadataFromCli(c emigo.CliCastable) EditAnOfferActionResValidationWarningsMetadata {
	data := EditAnOfferActionResValidationWarningsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type EditAnOfferActionResValidationWarningsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func GetEditAnOfferActionResAfterSalesServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "implied-warranty",
			Type:     "object",
			Children: GetEditAnOfferActionResAfterSalesServicesImpliedWarrantyCliFlags("implied-warranty-"),
		},
		{
			Name:     prefix + "return-policy",
			Type:     "object",
			Children: GetEditAnOfferActionResAfterSalesServicesReturnPolicyCliFlags("return-policy-"),
		},
		{
			Name:     prefix + "warranty",
			Type:     "object",
			Children: GetEditAnOfferActionResAfterSalesServicesWarrantyCliFlags("warranty-"),
		},
	}
}
func CastEditAnOfferActionResAfterSalesServicesFromCli(c emigo.CliCastable) EditAnOfferActionResAfterSalesServices {
	data := EditAnOfferActionResAfterSalesServices{}
	if c.IsSet("implied-warranty") {
		data.ImpliedWarranty = CastEditAnOfferActionResAfterSalesServicesImpliedWarrantyFromCli(c)
	}
	if c.IsSet("return-policy") {
		data.ReturnPolicy = CastEditAnOfferActionResAfterSalesServicesReturnPolicyFromCli(c)
	}
	if c.IsSet("warranty") {
		data.Warranty = CastEditAnOfferActionResAfterSalesServicesWarrantyFromCli(c)
	}
	return data
}

// The base class definition for afterSalesServices
type EditAnOfferActionResAfterSalesServices struct {
	ImpliedWarranty EditAnOfferActionResAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    EditAnOfferActionResAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        EditAnOfferActionResAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

func GetEditAnOfferActionResAfterSalesServicesImpliedWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResAfterSalesServicesImpliedWarrantyFromCli(c emigo.CliCastable) EditAnOfferActionResAfterSalesServicesImpliedWarranty {
	data := EditAnOfferActionResAfterSalesServicesImpliedWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for impliedWarranty
type EditAnOfferActionResAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResAfterSalesServicesReturnPolicyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResAfterSalesServicesReturnPolicyFromCli(c emigo.CliCastable) EditAnOfferActionResAfterSalesServicesReturnPolicy {
	data := EditAnOfferActionResAfterSalesServicesReturnPolicy{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for returnPolicy
type EditAnOfferActionResAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResAfterSalesServicesWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResAfterSalesServicesWarrantyFromCli(c emigo.CliCastable) EditAnOfferActionResAfterSalesServicesWarranty {
	data := EditAnOfferActionResAfterSalesServicesWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for warranty
type EditAnOfferActionResAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResDiscountsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "wholesale-price-list",
			Type:     "object",
			Children: GetEditAnOfferActionResDiscountsWholesalePriceListCliFlags("wholesale-price-list-"),
		},
	}
}
func CastEditAnOfferActionResDiscountsFromCli(c emigo.CliCastable) EditAnOfferActionResDiscounts {
	data := EditAnOfferActionResDiscounts{}
	if c.IsSet("wholesale-price-list") {
		data.WholesalePriceList = CastEditAnOfferActionResDiscountsWholesalePriceListFromCli(c)
	}
	return data
}

// The base class definition for discounts
type EditAnOfferActionResDiscounts struct {
	WholesalePriceList EditAnOfferActionResDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

func GetEditAnOfferActionResDiscountsWholesalePriceListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResDiscountsWholesalePriceListFromCli(c emigo.CliCastable) EditAnOfferActionResDiscountsWholesalePriceList {
	data := EditAnOfferActionResDiscountsWholesalePriceList{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for wholesalePriceList
type EditAnOfferActionResDiscountsWholesalePriceList struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResContactCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResContactFromCli(c emigo.CliCastable) EditAnOfferActionResContact {
	data := EditAnOfferActionResContact{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for contact
type EditAnOfferActionResContact struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResAttachmentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResAttachmentsFromCli(c emigo.CliCastable) EditAnOfferActionResAttachments {
	data := EditAnOfferActionResAttachments{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for attachments
type EditAnOfferActionResAttachments struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResFundraisingCampaignCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResFundraisingCampaignFromCli(c emigo.CliCastable) EditAnOfferActionResFundraisingCampaign {
	data := EditAnOfferActionResFundraisingCampaign{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for fundraisingCampaign
type EditAnOfferActionResFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResAdditionalServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResAdditionalServicesFromCli(c emigo.CliCastable) EditAnOfferActionResAdditionalServices {
	data := EditAnOfferActionResAdditionalServices{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additionalServices
type EditAnOfferActionResAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResSizeTableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResSizeTableFromCli(c emigo.CliCastable) EditAnOfferActionResSizeTable {
	data := EditAnOfferActionResSizeTable{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for sizeTable
type EditAnOfferActionResSizeTable struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResLocationCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResLocationFromCli(c emigo.CliCastable) EditAnOfferActionResLocation {
	data := EditAnOfferActionResLocation{}
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
type EditAnOfferActionResLocation struct {
	City        string `json:"city" yaml:"city"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
	PostCode    string `json:"postCode" yaml:"postCode"`
	Province    string `json:"province" yaml:"province"`
}

func GetEditAnOfferActionResExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResExternalFromCli(c emigo.CliCastable) EditAnOfferActionResExternal {
	data := EditAnOfferActionResExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type EditAnOfferActionResExternal struct {
	Id string `json:"id" yaml:"id"`
}

func GetEditAnOfferActionResTaxSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResTaxSettingsFromCli(c emigo.CliCastable) EditAnOfferActionResTaxSettings {
	data := EditAnOfferActionResTaxSettings{}
	if c.IsSet("subject") {
		data.Subject = c.String("subject")
	}
	if c.IsSet("exemption") {
		data.Exemption = c.String("exemption")
	}
	if c.IsSet("rates") {
		data.Rates = emigo.CapturePossibleArray(CastEditAnOfferActionResTaxSettingsRatesFromCli, "rates", c)
	}
	return data
}

// The base class definition for taxSettings
type EditAnOfferActionResTaxSettings struct {
	Subject   string                                 `json:"subject" yaml:"subject"`
	Exemption string                                 `json:"exemption" yaml:"exemption"`
	Rates     []EditAnOfferActionResTaxSettingsRates `json:"rates" yaml:"rates"`
}

func GetEditAnOfferActionResTaxSettingsRatesCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResTaxSettingsRatesFromCli(c emigo.CliCastable) EditAnOfferActionResTaxSettingsRates {
	data := EditAnOfferActionResTaxSettingsRates{}
	if c.IsSet("rate") {
		data.Rate = c.String("rate")
	}
	if c.IsSet("country-code") {
		data.CountryCode = c.String("country-code")
	}
	return data
}

// The base class definition for rates
type EditAnOfferActionResTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

func GetEditAnOfferActionResMessageToSellerSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastEditAnOfferActionResMessageToSellerSettingsFromCli(c emigo.CliCastable) EditAnOfferActionResMessageToSellerSettings {
	data := EditAnOfferActionResMessageToSellerSettings{}
	if c.IsSet("mode") {
		data.Mode = c.String("mode")
	}
	if c.IsSet("hint") {
		data.Hint = c.String("hint")
	}
	return data
}

// The base class definition for messageToSellerSettings
type EditAnOfferActionResMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

func GetEditAnOfferActionResDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResDescriptionFromCli(c emigo.CliCastable) EditAnOfferActionResDescription {
	data := EditAnOfferActionResDescription{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastEditAnOfferActionResDescriptionSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for description
type EditAnOfferActionResDescription struct {
	Sections []EditAnOfferActionResDescriptionSections `json:"sections" yaml:"sections"`
}

func GetEditAnOfferActionResDescriptionSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastEditAnOfferActionResDescriptionSectionsFromCli(c emigo.CliCastable) EditAnOfferActionResDescriptionSections {
	data := EditAnOfferActionResDescriptionSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastEditAnOfferActionResDescriptionSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type EditAnOfferActionResDescriptionSections struct {
	Items []EditAnOfferActionResDescriptionSectionsItems `json:"items" yaml:"items"`
}

func GetEditAnOfferActionResDescriptionSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastEditAnOfferActionResDescriptionSectionsItemsFromCli(c emigo.CliCastable) EditAnOfferActionResDescriptionSectionsItems {
	data := EditAnOfferActionResDescriptionSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
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

// EditAnOfferActionRaw registers a raw Gin route for the EditAnOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func EditAnOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := EditAnOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type EditAnOfferActionRequestSig = func(c EditAnOfferActionRequest) (*EditAnOfferActionResponse, error)

// EditAnOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the EditAnOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func EditAnOfferActionHandler(
	handler EditAnOfferActionRequestSig,
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

// EditAnOfferAction is a high-level convenience wrapper around EditAnOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func EditAnOfferActionGin(r gin.IRoutes, handler EditAnOfferActionRequestSig) {
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

func (x EditAnOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x EditAnOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type EditAnOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
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
