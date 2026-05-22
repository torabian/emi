package external

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli/v3"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetAllDataOfTheParticularProductOfferAction
 */
/*
Here is a quick function implementation to make your life easier:
// Actual implementation of GetAllDataOfTheParticularProductOfferAction
func GetAllDataOfTheParticularProductOfferAction(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	return &GetAllDataOfTheParticularProductOfferActionResponse{
		// Payload is an interface. Use it at carefully.
	}, nil
}
*/
func GetAllDataOfTheParticularProductOfferActionMeta() struct {
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
		Name:        "GetAllDataOfTheParticularProductOfferAction",
		CliName:     "get all data of the particular product-offer-action",
		URL:         "https://api.{environment}/sale/product-offers/{offerId}",
		Method:      "GET",
		Description: `Full response model returned by GET /sale/product-offers/{offerId}`,
	}
}
func GetGetAllDataOfTheParticularProductOfferActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "string",
			Description: "Unique offer identifier",
		},
		{
			Name:        prefix + "name",
			Type:        "string",
			Description: "Offer title",
		},
		{
			Name:        prefix + "language",
			Type:        "string",
			Description: "Offer language code (e.g. pl-PL)",
		},
		{
			Name:        prefix + "created-at",
			Type:        "string",
			Description: "Offer creation timestamp (ISO8601)",
		},
		{
			Name:        prefix + "updated-at",
			Type:        "string",
			Description: "Offer last update timestamp (ISO8601)",
		},
		{
			Name:     prefix + "category",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResCategoryCliFlags("category-"),
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "contact",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResContactCliFlags("contact-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResPublicationCliFlags("publication-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "payments",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResPaymentsCliFlags("payments-"),
		},
		{
			Name:     prefix + "delivery",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResDeliveryCliFlags("delivery-"),
		},
		{
			Name:     prefix + "after-sales-services",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesCliFlags("after-sales-services-"),
		},
		{
			Name:     prefix + "discounts",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResDiscountsCliFlags("discounts-"),
		},
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResDescriptionCliFlags("description-"),
		},
		{
			Name: prefix + "images",
			Type: "slice",
		},
		{
			Name: prefix + "product-set",
			Type: "array",
		},
		{
			Name: prefix + "attachments",
			Type: "array",
		},
		{
			Name:     prefix + "fundraising-campaign",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResFundraisingCampaignCliFlags("fundraising-campaign-"),
		},
		{
			Name:     prefix + "additional-services",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResAdditionalServicesCliFlags("additional-services-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "map?",
		},
		{
			Name:     prefix + "b2b",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResB2bCliFlags("b2b-"),
		},
		{
			Name:     prefix + "compatibility-list",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResCompatibilityListCliFlags("compatibility-list-"),
		},
		{
			Name:     prefix + "validation",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResValidationCliFlags("validation-"),
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResExternalCliFlags("external-"),
		},
		{
			Name:     prefix + "size-table",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResSizeTableCliFlags("size-table-"),
		},
		{
			Name:     prefix + "tax-settings",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResTaxSettingsCliFlags("tax-settings-"),
		},
		{
			Name:     prefix + "message-to-seller-settings",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResMessageToSellerSettingsCliFlags("message-to-seller-settings-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionRes {
	data := GetAllDataOfTheParticularProductOfferActionRes{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("language") {
		data.Language = c.String("language")
	}
	if c.IsSet("created-at") {
		data.CreatedAt = c.String("created-at")
	}
	if c.IsSet("updated-at") {
		data.UpdatedAt = c.String("updated-at")
	}
	if c.IsSet("category") {
		data.Category = CastGetAllDataOfTheParticularProductOfferActionResCategoryFromCli(c)
	}
	if c.IsSet("stock") {
		data.Stock = CastGetAllDataOfTheParticularProductOfferActionResStockFromCli(c)
	}
	if c.IsSet("contact") {
		data.Contact = CastGetAllDataOfTheParticularProductOfferActionResContactFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastGetAllDataOfTheParticularProductOfferActionResPublicationFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastGetAllDataOfTheParticularProductOfferActionResSellingModeFromCli(c)
	}
	if c.IsSet("payments") {
		data.Payments = CastGetAllDataOfTheParticularProductOfferActionResPaymentsFromCli(c)
	}
	if c.IsSet("delivery") {
		data.Delivery = CastGetAllDataOfTheParticularProductOfferActionResDeliveryFromCli(c)
	}
	if c.IsSet("after-sales-services") {
		data.AfterSalesServices = CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesFromCli(c)
	}
	if c.IsSet("discounts") {
		data.Discounts = CastGetAllDataOfTheParticularProductOfferActionResDiscountsFromCli(c)
	}
	if c.IsSet("description") {
		data.Description = CastGetAllDataOfTheParticularProductOfferActionResDescriptionFromCli(c)
	}
	if c.IsSet("images") {
		emigo.InflatePossibleSlice(c.String("images"), &data.Images)
	}
	if c.IsSet("product-set") {
		data.ProductSet = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResProductSetFromCli, "product-set", c)
	}
	if c.IsSet("attachments") {
		data.Attachments = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResAttachmentsFromCli, "attachments", c)
	}
	if c.IsSet("fundraising-campaign") {
		data.FundraisingCampaign = CastGetAllDataOfTheParticularProductOfferActionResFundraisingCampaignFromCli(c)
	}
	if c.IsSet("additional-services") {
		data.AdditionalServices = CastGetAllDataOfTheParticularProductOfferActionResAdditionalServicesFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		emigo.ParseNullable(c.String("additional-marketplaces"), &data.AdditionalMarketplaces)
	}
	if c.IsSet("b2b") {
		data.B2b = CastGetAllDataOfTheParticularProductOfferActionResB2bFromCli(c)
	}
	if c.IsSet("compatibility-list") {
		data.CompatibilityList = CastGetAllDataOfTheParticularProductOfferActionResCompatibilityListFromCli(c)
	}
	if c.IsSet("validation") {
		data.Validation = CastGetAllDataOfTheParticularProductOfferActionResValidationFromCli(c)
	}
	if c.IsSet("external") {
		data.External = CastGetAllDataOfTheParticularProductOfferActionResExternalFromCli(c)
	}
	if c.IsSet("size-table") {
		data.SizeTable = CastGetAllDataOfTheParticularProductOfferActionResSizeTableFromCli(c)
	}
	if c.IsSet("tax-settings") {
		data.TaxSettings = CastGetAllDataOfTheParticularProductOfferActionResTaxSettingsFromCli(c)
	}
	if c.IsSet("message-to-seller-settings") {
		data.MessageToSellerSettings = CastGetAllDataOfTheParticularProductOfferActionResMessageToSellerSettingsFromCli(c)
	}
	return data
}

// The base class definition for getAllDataOfTheParticularProductOfferActionRes
type GetAllDataOfTheParticularProductOfferActionRes struct {
	// Unique offer identifier
	Id string `json:"id" yaml:"id"`
	// Offer title
	Name string `json:"name" yaml:"name"`
	// Offer language code (e.g. pl-PL)
	Language string `json:"language" yaml:"language"`
	// Offer creation timestamp (ISO8601)
	CreatedAt string `json:"createdAt" yaml:"createdAt"`
	// Offer last update timestamp (ISO8601)
	UpdatedAt               string                                                                `json:"updatedAt" yaml:"updatedAt"`
	Category                GetAllDataOfTheParticularProductOfferActionResCategory                `json:"category" yaml:"category"`
	Stock                   GetAllDataOfTheParticularProductOfferActionResStock                   `json:"stock" yaml:"stock"`
	Contact                 GetAllDataOfTheParticularProductOfferActionResContact                 `json:"contact" yaml:"contact"`
	Publication             GetAllDataOfTheParticularProductOfferActionResPublication             `json:"publication" yaml:"publication"`
	SellingMode             GetAllDataOfTheParticularProductOfferActionResSellingMode             `json:"sellingMode" yaml:"sellingMode"`
	Payments                GetAllDataOfTheParticularProductOfferActionResPayments                `json:"payments" yaml:"payments"`
	Delivery                GetAllDataOfTheParticularProductOfferActionResDelivery                `json:"delivery" yaml:"delivery"`
	AfterSalesServices      GetAllDataOfTheParticularProductOfferActionResAfterSalesServices      `json:"afterSalesServices" yaml:"afterSalesServices"`
	Discounts               GetAllDataOfTheParticularProductOfferActionResDiscounts               `json:"discounts" yaml:"discounts"`
	Description             GetAllDataOfTheParticularProductOfferActionResDescription             `json:"description" yaml:"description"`
	Images                  []string                                                              `json:"images" yaml:"images"`
	ProductSet              []GetAllDataOfTheParticularProductOfferActionResProductSet            `json:"productSet" yaml:"productSet"`
	Attachments             []GetAllDataOfTheParticularProductOfferActionResAttachments           `json:"attachments" yaml:"attachments"`
	FundraisingCampaign     GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign     `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	AdditionalServices      GetAllDataOfTheParticularProductOfferActionResAdditionalServices      `json:"additionalServices" yaml:"additionalServices"`
	AdditionalMarketplaces  emigo.Nullable[map[any]any]                                           `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
	B2b                     GetAllDataOfTheParticularProductOfferActionResB2b                     `json:"b2b" yaml:"b2b"`
	CompatibilityList       GetAllDataOfTheParticularProductOfferActionResCompatibilityList       `json:"compatibilityList" yaml:"compatibilityList"`
	Validation              GetAllDataOfTheParticularProductOfferActionResValidation              `json:"validation" yaml:"validation"`
	External                GetAllDataOfTheParticularProductOfferActionResExternal                `json:"external" yaml:"external"`
	SizeTable               GetAllDataOfTheParticularProductOfferActionResSizeTable               `json:"sizeTable" yaml:"sizeTable"`
	TaxSettings             GetAllDataOfTheParticularProductOfferActionResTaxSettings             `json:"taxSettings" yaml:"taxSettings"`
	MessageToSellerSettings GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings `json:"messageToSellerSettings" yaml:"messageToSellerSettings"`
}

func GetGetAllDataOfTheParticularProductOfferActionResCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResCategoryFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResCategory {
	data := GetAllDataOfTheParticularProductOfferActionResCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type GetAllDataOfTheParticularProductOfferActionResCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResStockCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResStockFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResStock {
	data := GetAllDataOfTheParticularProductOfferActionResStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	if c.IsSet("unit") {
		data.Unit = c.String("unit")
	}
	return data
}

// The base class definition for stock
type GetAllDataOfTheParticularProductOfferActionResStock struct {
	Available int    `json:"available" yaml:"available"`
	Unit      string `json:"unit" yaml:"unit"`
}

func GetGetAllDataOfTheParticularProductOfferActionResContactCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResContactFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResContact {
	data := GetAllDataOfTheParticularProductOfferActionResContact{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for contact
type GetAllDataOfTheParticularProductOfferActionResContact struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResPublicationCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesCliFlags("marketplaces-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResPublicationFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResPublication {
	data := GetAllDataOfTheParticularProductOfferActionResPublication{}
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
		data.Marketplaces = CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesFromCli(c)
	}
	return data
}

// The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResPublication struct {
	Duration     string                                                                `json:"duration" yaml:"duration"`
	StartingAt   string                                                                `json:"startingAt" yaml:"startingAt"`
	EndingAt     string                                                                `json:"endingAt" yaml:"endingAt"`
	EndedBy      string                                                                `json:"endedBy" yaml:"endedBy"`
	Status       string                                                                `json:"status" yaml:"status"`
	Republish    bool                                                                  `json:"republish" yaml:"republish"`
	Marketplaces GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}

func GetGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBaseCliFlags("base-"),
		},
		{
			Name: prefix + "additional",
			Type: "array",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces {
	data := GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces{}
	if c.IsSet("base") {
		data.Base = CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBaseFromCli(c)
	}
	if c.IsSet("additional") {
		data.Additional = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditionalFromCli, "additional", c)
	}
	return data
}

// The base class definition for marketplaces
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplaces struct {
	Base       GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase         `json:"base" yaml:"base"`
	Additional []GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
}

func GetGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBaseCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBaseFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase {
	data := GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for base
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditionalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditionalFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional {
	data := GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additional
type GetAllDataOfTheParticularProductOfferActionResPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "format",
			Type: "string",
		},
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResSellingModePriceCliFlags("price-"),
		},
		{
			Name:     prefix + "minimal-price",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPriceCliFlags("minimal-price-"),
		},
		{
			Name:     prefix + "starting-price",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResSellingModeStartingPriceCliFlags("starting-price-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResSellingModeFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResSellingMode {
	data := GetAllDataOfTheParticularProductOfferActionResSellingMode{}
	if c.IsSet("format") {
		data.Format = c.String("format")
	}
	if c.IsSet("price") {
		data.Price = CastGetAllDataOfTheParticularProductOfferActionResSellingModePriceFromCli(c)
	}
	if c.IsSet("minimal-price") {
		data.MinimalPrice = CastGetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPriceFromCli(c)
	}
	if c.IsSet("starting-price") {
		data.StartingPrice = CastGetAllDataOfTheParticularProductOfferActionResSellingModeStartingPriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type GetAllDataOfTheParticularProductOfferActionResSellingMode struct {
	Format        string                                                                 `json:"format" yaml:"format"`
	Price         GetAllDataOfTheParticularProductOfferActionResSellingModePrice         `json:"price" yaml:"price"`
	MinimalPrice  GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice  `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice `json:"startingPrice" yaml:"startingPrice"`
}

func GetGetAllDataOfTheParticularProductOfferActionResSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResSellingModePriceFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResSellingModePrice {
	data := GetAllDataOfTheParticularProductOfferActionResSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type GetAllDataOfTheParticularProductOfferActionResSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPriceFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice {
	data := GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for minimalPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetAllDataOfTheParticularProductOfferActionResSellingModeStartingPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResSellingModeStartingPriceFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice {
	data := GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for startingPrice
type GetAllDataOfTheParticularProductOfferActionResSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetAllDataOfTheParticularProductOfferActionResPaymentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "invoice",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResPaymentsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResPayments {
	data := GetAllDataOfTheParticularProductOfferActionResPayments{}
	if c.IsSet("invoice") {
		data.Invoice = c.String("invoice")
	}
	return data
}

// The base class definition for payments
type GetAllDataOfTheParticularProductOfferActionResPayments struct {
	Invoice string `json:"invoice" yaml:"invoice"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDeliveryCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResDeliveryShippingRatesCliFlags("shipping-rates-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDeliveryFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDelivery {
	data := GetAllDataOfTheParticularProductOfferActionResDelivery{}
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
		data.ShippingRates = CastGetAllDataOfTheParticularProductOfferActionResDeliveryShippingRatesFromCli(c)
	}
	return data
}

// The base class definition for delivery
type GetAllDataOfTheParticularProductOfferActionResDelivery struct {
	HandlingTime   string                                                              `json:"handlingTime" yaml:"handlingTime"`
	AdditionalInfo string                                                              `json:"additionalInfo" yaml:"additionalInfo"`
	ShipmentDate   string                                                              `json:"shipmentDate" yaml:"shipmentDate"`
	ShippingRates  GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDeliveryShippingRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDeliveryShippingRatesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates {
	data := GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for shippingRates
type GetAllDataOfTheParticularProductOfferActionResDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "implied-warranty",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarrantyCliFlags("implied-warranty-"),
		},
		{
			Name:     prefix + "return-policy",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicyCliFlags("return-policy-"),
		},
		{
			Name:     prefix + "warranty",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarrantyCliFlags("warranty-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAfterSalesServices {
	data := GetAllDataOfTheParticularProductOfferActionResAfterSalesServices{}
	if c.IsSet("implied-warranty") {
		data.ImpliedWarranty = CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarrantyFromCli(c)
	}
	if c.IsSet("return-policy") {
		data.ReturnPolicy = CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicyFromCli(c)
	}
	if c.IsSet("warranty") {
		data.Warranty = CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarrantyFromCli(c)
	}
	return data
}

// The base class definition for afterSalesServices
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServices struct {
	ImpliedWarranty GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarrantyFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty {
	data := GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for impliedWarranty
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicyFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy {
	data := GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for returnPolicy
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarrantyFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty {
	data := GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for warranty
type GetAllDataOfTheParticularProductOfferActionResAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDiscountsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "wholesale-price-list",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceListCliFlags("wholesale-price-list-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDiscountsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDiscounts {
	data := GetAllDataOfTheParticularProductOfferActionResDiscounts{}
	if c.IsSet("wholesale-price-list") {
		data.WholesalePriceList = CastGetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceListFromCli(c)
	}
	return data
}

// The base class definition for discounts
type GetAllDataOfTheParticularProductOfferActionResDiscounts struct {
	WholesalePriceList GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList `json:"wholesalePriceList" yaml:"wholesalePriceList"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceListFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList {
	data := GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for wholesalePriceList
type GetAllDataOfTheParticularProductOfferActionResDiscountsWholesalePriceList struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDescriptionFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDescription {
	data := GetAllDataOfTheParticularProductOfferActionResDescription{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsFromCli, "sections", c)
	}
	return data
}

// The base class definition for description
type GetAllDataOfTheParticularProductOfferActionResDescription struct {
	Sections []GetAllDataOfTheParticularProductOfferActionResDescriptionSections `json:"sections" yaml:"sections"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDescriptionSections {
	data := GetAllDataOfTheParticularProductOfferActionResDescriptionSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItemsFromCli, "items", c)
	}
	return data
}

// The base class definition for sections
type GetAllDataOfTheParticularProductOfferActionResDescriptionSections struct {
	Items []GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems `json:"items" yaml:"items"`
}

func GetGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItemsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems {
	data := GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for items
type GetAllDataOfTheParticularProductOfferActionResDescriptionSectionsItems struct {
	Type string `json:"type" yaml:"type"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "quantity",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetQuantityCliFlags("quantity-"),
		},
		{
			Name:     prefix + "product",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetProductCliFlags("product-"),
		},
		{
			Name:     prefix + "responsible-person",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePersonCliFlags("responsible-person-"),
		},
		{
			Name:     prefix + "responsible-producer",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducerCliFlags("responsible-producer-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformationCliFlags("safety-information-"),
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
func CastGetAllDataOfTheParticularProductOfferActionResProductSetFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSet {
	data := GetAllDataOfTheParticularProductOfferActionResProductSet{}
	if c.IsSet("quantity") {
		data.Quantity = CastGetAllDataOfTheParticularProductOfferActionResProductSetQuantityFromCli(c)
	}
	if c.IsSet("product") {
		data.Product = CastGetAllDataOfTheParticularProductOfferActionResProductSetProductFromCli(c)
	}
	if c.IsSet("responsible-person") {
		data.ResponsiblePerson = CastGetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePersonFromCli(c)
	}
	if c.IsSet("responsible-producer") {
		data.ResponsibleProducer = CastGetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducerFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastGetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformationFromCli(c)
	}
	if c.IsSet("marketed-before-gpsr-obligation") {
		data.MarketedBeforeGPSRObligation = bool(c.Bool("marketed-before-gpsr-obligation"))
	}
	if c.IsSet("deposits") {
		data.Deposits = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResProductSetDepositsFromCli, "deposits", c)
	}
	return data
}

// The base class definition for productSet
type GetAllDataOfTheParticularProductOfferActionResProductSet struct {
	Quantity                     GetAllDataOfTheParticularProductOfferActionResProductSetQuantity            `json:"quantity" yaml:"quantity"`
	Product                      GetAllDataOfTheParticularProductOfferActionResProductSetProduct             `json:"product" yaml:"product"`
	ResponsiblePerson            GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson   `json:"responsiblePerson" yaml:"responsiblePerson"`
	ResponsibleProducer          GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer `json:"responsibleProducer" yaml:"responsibleProducer"`
	SafetyInformation            GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation   `json:"safetyInformation" yaml:"safetyInformation"`
	MarketedBeforeGPSRObligation bool                                                                        `json:"marketedBeforeGPSRObligation" yaml:"marketedBeforeGPSRObligation"`
	Deposits                     []GetAllDataOfTheParticularProductOfferActionResProductSetDeposits          `json:"deposits" yaml:"deposits"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetQuantityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "value",
			Type: "int",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetQuantityFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetQuantity {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetQuantity{}
	if c.IsSet("value") {
		data.Value = int(c.Int64("value"))
	}
	return data
}

// The base class definition for quantity
type GetAllDataOfTheParticularProductOfferActionResProductSetQuantity struct {
	Value int `json:"value" yaml:"value"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetProductPublicationCliFlags("publication-"),
		},
		{
			Name: prefix + "parameters",
			Type: "array",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProduct {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProduct{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("is-ai-co-created") {
		data.IsAiCoCreated = bool(c.Bool("is-ai-co-created"))
	}
	if c.IsSet("publication") {
		data.Publication = CastGetAllDataOfTheParticularProductOfferActionResProductSetProductPublicationFromCli(c)
	}
	if c.IsSet("parameters") {
		data.Parameters = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersFromCli, "parameters", c)
	}
	return data
}

// The base class definition for product
type GetAllDataOfTheParticularProductOfferActionResProductSetProduct struct {
	Id            string                                                                      `json:"id" yaml:"id"`
	IsAiCoCreated bool                                                                        `json:"isAiCoCreated" yaml:"isAiCoCreated"`
	Publication   GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication  `json:"publication" yaml:"publication"`
	Parameters    []GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters `json:"parameters" yaml:"parameters"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "status",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductPublicationFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication{}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	return data
}

// The base class definition for publication
type GetAllDataOfTheParticularProductOfferActionResProductSetProductPublication struct {
	Status string `json:"status" yaml:"status"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValueCliFlags("range-value-"),
		},
		{
			Name: prefix + "values",
			Type: "array",
		},
		{
			Name: prefix + "values-ids",
			Type: "array",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("range-value") {
		data.RangeValue = CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValueFromCli(c)
	}
	if c.IsSet("values") {
		data.Values = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesFromCli, "values", c)
	}
	if c.IsSet("values-ids") {
		data.ValuesIds = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIdsFromCli, "values-ids", c)
	}
	return data
}

// The base class definition for parameters
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParameters struct {
	Id         string                                                                               `json:"id" yaml:"id"`
	Name       string                                                                               `json:"name" yaml:"name"`
	RangeValue GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue  `json:"rangeValue" yaml:"rangeValue"`
	Values     []GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues    `json:"values" yaml:"values"`
	ValuesIds  []GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds `json:"valuesIds" yaml:"valuesIds"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValueCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValueFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue{}
	if c.IsSet("from") {
		data.From = c.String("from")
	}
	if c.IsSet("to") {
		data.To = c.String("to")
	}
	return data
}

// The base class definition for rangeValue
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersRangeValue struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues{}
	return data
}

// The base class definition for values
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValues struct {
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIdsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIdsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds{}
	return data
}

// The base class definition for valuesIds
type GetAllDataOfTheParticularProductOfferActionResProductSetProductParametersValuesIds struct {
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePersonCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePersonFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for responsiblePerson
type GetAllDataOfTheParticularProductOfferActionResProductSetResponsiblePerson struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducerCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducerFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for responsibleProducer
type GetAllDataOfTheParticularProductOfferActionResProductSetResponsibleProducer struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformationFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	if c.IsSet("description") {
		data.Description = c.String("description")
	}
	return data
}

// The base class definition for safetyInformation
type GetAllDataOfTheParticularProductOfferActionResProductSetSafetyInformation struct {
	Type        string `json:"type" yaml:"type"`
	Description string `json:"description" yaml:"description"`
}

func GetGetAllDataOfTheParticularProductOfferActionResProductSetDepositsCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResProductSetDepositsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResProductSetDeposits {
	data := GetAllDataOfTheParticularProductOfferActionResProductSetDeposits{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("quantity") {
		data.Quantity = int(c.Int64("quantity"))
	}
	return data
}

// The base class definition for deposits
type GetAllDataOfTheParticularProductOfferActionResProductSetDeposits struct {
	Id       string `json:"id" yaml:"id"`
	Quantity int    `json:"quantity" yaml:"quantity"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAttachmentsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAttachmentsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAttachments {
	data := GetAllDataOfTheParticularProductOfferActionResAttachments{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for attachments
type GetAllDataOfTheParticularProductOfferActionResAttachments struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResFundraisingCampaignCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResFundraisingCampaignFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign {
	data := GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for fundraisingCampaign
type GetAllDataOfTheParticularProductOfferActionResFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResAdditionalServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResAdditionalServicesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResAdditionalServices {
	data := GetAllDataOfTheParticularProductOfferActionResAdditionalServices{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additionalServices
type GetAllDataOfTheParticularProductOfferActionResAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResB2bCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "buyable-only-by-business",
			Type: "bool",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResB2bFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResB2b {
	data := GetAllDataOfTheParticularProductOfferActionResB2b{}
	if c.IsSet("buyable-only-by-business") {
		data.BuyableOnlyByBusiness = bool(c.Bool("buyable-only-by-business"))
	}
	return data
}

// The base class definition for b2b
type GetAllDataOfTheParticularProductOfferActionResB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

func GetGetAllDataOfTheParticularProductOfferActionResCompatibilityListCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResCompatibilityListFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResCompatibilityList {
	data := GetAllDataOfTheParticularProductOfferActionResCompatibilityList{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}

// The base class definition for compatibilityList
type GetAllDataOfTheParticularProductOfferActionResCompatibilityList struct {
	Type string `json:"type" yaml:"type"`
}

func GetGetAllDataOfTheParticularProductOfferActionResValidationCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResValidationFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResValidation {
	data := GetAllDataOfTheParticularProductOfferActionResValidation{}
	if c.IsSet("validated-at") {
		data.ValidatedAt = c.String("validated-at")
	}
	if c.IsSet("errors") {
		data.Errors = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResValidationErrorsFromCli, "errors", c)
	}
	if c.IsSet("warnings") {
		data.Warnings = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResValidationWarningsFromCli, "warnings", c)
	}
	return data
}

// The base class definition for validation
type GetAllDataOfTheParticularProductOfferActionResValidation struct {
	ValidatedAt string                                                             `json:"validatedAt" yaml:"validatedAt"`
	Errors      []GetAllDataOfTheParticularProductOfferActionResValidationErrors   `json:"errors" yaml:"errors"`
	Warnings    []GetAllDataOfTheParticularProductOfferActionResValidationWarnings `json:"warnings" yaml:"warnings"`
}

func GetGetAllDataOfTheParticularProductOfferActionResValidationErrorsCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadataCliFlags("metadata-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResValidationErrorsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResValidationErrors {
	data := GetAllDataOfTheParticularProductOfferActionResValidationErrors{}
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
		data.Metadata = CastGetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadataFromCli(c)
	}
	return data
}

// The base class definition for errors
type GetAllDataOfTheParticularProductOfferActionResValidationErrors struct {
	Code        string                                                                 `json:"code" yaml:"code"`
	Details     string                                                                 `json:"details" yaml:"details"`
	Message     string                                                                 `json:"message" yaml:"message"`
	Path        string                                                                 `json:"path" yaml:"path"`
	UserMessage string                                                                 `json:"userMessage" yaml:"userMessage"`
	Metadata    GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata `json:"metadata" yaml:"metadata"`
}

func GetGetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadataFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata {
	data := GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type GetAllDataOfTheParticularProductOfferActionResValidationErrorsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func GetGetAllDataOfTheParticularProductOfferActionResValidationWarningsCliFlags(prefix string) []emigo.CliFlag {
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
			Children: GetGetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadataCliFlags("metadata-"),
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResValidationWarningsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResValidationWarnings {
	data := GetAllDataOfTheParticularProductOfferActionResValidationWarnings{}
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
		data.Metadata = CastGetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadataFromCli(c)
	}
	return data
}

// The base class definition for warnings
type GetAllDataOfTheParticularProductOfferActionResValidationWarnings struct {
	Code        string                                                                   `json:"code" yaml:"code"`
	Details     string                                                                   `json:"details" yaml:"details"`
	Message     string                                                                   `json:"message" yaml:"message"`
	Path        string                                                                   `json:"path" yaml:"path"`
	UserMessage string                                                                   `json:"userMessage" yaml:"userMessage"`
	Metadata    GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata `json:"metadata" yaml:"metadata"`
}

func GetGetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "product-id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadataFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata {
	data := GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata{}
	if c.IsSet("product-id") {
		data.ProductId = c.String("product-id")
	}
	return data
}

// The base class definition for metadata
type GetAllDataOfTheParticularProductOfferActionResValidationWarningsMetadata struct {
	ProductId string `json:"productId" yaml:"productId"`
}

func GetGetAllDataOfTheParticularProductOfferActionResExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResExternalFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResExternal {
	data := GetAllDataOfTheParticularProductOfferActionResExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type GetAllDataOfTheParticularProductOfferActionResExternal struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResSizeTableCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetAllDataOfTheParticularProductOfferActionResSizeTableFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResSizeTable {
	data := GetAllDataOfTheParticularProductOfferActionResSizeTable{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for sizeTable
type GetAllDataOfTheParticularProductOfferActionResSizeTable struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetAllDataOfTheParticularProductOfferActionResTaxSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResTaxSettingsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResTaxSettings {
	data := GetAllDataOfTheParticularProductOfferActionResTaxSettings{}
	if c.IsSet("subject") {
		data.Subject = c.String("subject")
	}
	if c.IsSet("exemption") {
		data.Exemption = c.String("exemption")
	}
	if c.IsSet("rates") {
		data.Rates = emigo.CapturePossibleArray(CastGetAllDataOfTheParticularProductOfferActionResTaxSettingsRatesFromCli, "rates", c)
	}
	return data
}

// The base class definition for taxSettings
type GetAllDataOfTheParticularProductOfferActionResTaxSettings struct {
	Subject   string                                                           `json:"subject" yaml:"subject"`
	Exemption string                                                           `json:"exemption" yaml:"exemption"`
	Rates     []GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates `json:"rates" yaml:"rates"`
}

func GetGetAllDataOfTheParticularProductOfferActionResTaxSettingsRatesCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResTaxSettingsRatesFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates {
	data := GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates{}
	if c.IsSet("rate") {
		data.Rate = c.String("rate")
	}
	if c.IsSet("country-code") {
		data.CountryCode = c.String("country-code")
	}
	return data
}

// The base class definition for rates
type GetAllDataOfTheParticularProductOfferActionResTaxSettingsRates struct {
	Rate        string `json:"rate" yaml:"rate"`
	CountryCode string `json:"countryCode" yaml:"countryCode"`
}

func GetGetAllDataOfTheParticularProductOfferActionResMessageToSellerSettingsCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetAllDataOfTheParticularProductOfferActionResMessageToSellerSettingsFromCli(c emigo.CliCastable) GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings {
	data := GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings{}
	if c.IsSet("mode") {
		data.Mode = c.String("mode")
	}
	if c.IsSet("hint") {
		data.Hint = c.String("hint")
	}
	return data
}

// The base class definition for messageToSellerSettings
type GetAllDataOfTheParticularProductOfferActionResMessageToSellerSettings struct {
	Mode string `json:"mode" yaml:"mode"`
	Hint string `json:"hint" yaml:"hint"`
}

func (x *GetAllDataOfTheParticularProductOfferActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetAllDataOfTheParticularProductOfferActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
	// Do not manually fill this in. It has no effect. This is only useful when you are using
	// client code, and want to get access to the original response. When sending response from your
	// application it will be ignored.
	resp *http.Response
}

func (x *GetAllDataOfTheParticularProductOfferActionResponse) SetContentType(contentType string) *GetAllDataOfTheParticularProductOfferActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsStream(r io.Reader, contentType string) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsJSON(payload any) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}

// When the response is expected as documentation, you call this to get some type
// safety for the action which is happening.
func (x *GetAllDataOfTheParticularProductOfferActionResponse) WithIdeal(payload GetAllDataOfTheParticularProductOfferActionRes) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsHTML(payload string) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetAllDataOfTheParticularProductOfferActionResponse) AsBytes(payload []byte) *GetAllDataOfTheParticularProductOfferActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetAllDataOfTheParticularProductOfferActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetAllDataOfTheParticularProductOfferActionRaw registers a raw Gin route for the GetAllDataOfTheParticularProductOfferAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetAllDataOfTheParticularProductOfferActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetAllDataOfTheParticularProductOfferActionRequestSig = func(c GetAllDataOfTheParticularProductOfferActionRequest) (*GetAllDataOfTheParticularProductOfferActionResponse, error)

// GetAllDataOfTheParticularProductOfferActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetAllDataOfTheParticularProductOfferAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetAllDataOfTheParticularProductOfferActionHandler(
	handler GetAllDataOfTheParticularProductOfferActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetAllDataOfTheParticularProductOfferActionRequest{
			Body:        nil,
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

// GetAllDataOfTheParticularProductOfferAction is a high-level convenience wrapper around GetAllDataOfTheParticularProductOfferActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetAllDataOfTheParticularProductOfferActionGin(r gin.IRoutes, handler GetAllDataOfTheParticularProductOfferActionRequestSig) {
	method, url, h := GetAllDataOfTheParticularProductOfferActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get all data of the particular product-offerAction
 */
// Query wrapper with private fields
type GetAllDataOfTheParticularProductOfferActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetAllDataOfTheParticularProductOfferActionQueryFromString(rawQuery string) GetAllDataOfTheParticularProductOfferActionQuery {
	v := GetAllDataOfTheParticularProductOfferActionQuery{}
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
func GetAllDataOfTheParticularProductOfferActionQueryFromGin(c *gin.Context) GetAllDataOfTheParticularProductOfferActionQuery {
	return GetAllDataOfTheParticularProductOfferActionQueryFromString(c.Request.URL.RawQuery)
}
func GetAllDataOfTheParticularProductOfferActionQueryFromHttp(r *http.Request) GetAllDataOfTheParticularProductOfferActionQuery {
	return GetAllDataOfTheParticularProductOfferActionQueryFromString(r.URL.RawQuery)
}
func (q GetAllDataOfTheParticularProductOfferActionQuery) Values() url.Values {
	return q.values
}
func (q GetAllDataOfTheParticularProductOfferActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetAllDataOfTheParticularProductOfferActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetAllDataOfTheParticularProductOfferActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetAllDataOfTheParticularProductOfferActionRequest struct {
	Body        interface{}
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

func (x GetAllDataOfTheParticularProductOfferActionRequest) IsGin() bool {
	return x.GinCtx != nil
}
func (x GetAllDataOfTheParticularProductOfferActionRequest) IsCli() bool {
	return x.CliCtx != nil
}

// type GetAllDataOfTheParticularProductOfferActionResult struct {
// /resp *http.Response
// /	Payload interface{}
// /}
func GetAllDataOfTheParticularProductOfferActionClientCreateUrl(
	req GetAllDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*url.URL, error) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
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
func GetAllDataOfTheParticularProductOfferActionClientExecuteTyped(httpReq *http.Request) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// At this point, response is valid, and we need to return the results.
	var result GetAllDataOfTheParticularProductOfferActionResponse
	result.resp = resp
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	if err := json.Unmarshal(respBody, &result.Payload); err != nil {
		return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, err
	}
	return &GetAllDataOfTheParticularProductOfferActionResponse{Payload: result}, nil
}
func GetAllDataOfTheParticularProductOfferActionClientBuildRequest(req GetAllDataOfTheParticularProductOfferActionRequest, reqUrl *url.URL, config *emigo.APIClient) (*http.Request, error) {
	meta := GetAllDataOfTheParticularProductOfferActionMeta()
	httpReq, err := http.NewRequest(meta.Method, reqUrl.String(), nil)
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
func GetAllDataOfTheParticularProductOfferActionCall(
	req GetAllDataOfTheParticularProductOfferActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetAllDataOfTheParticularProductOfferActionResponse, error) {
	// This function intentionally is split into 3 different sections, so in case
	// of some modifications that we did not anticipate, at least a part would become quite useful.
	// first we create url, apply all path parameters, query params, etc
	u, err := GetAllDataOfTheParticularProductOfferActionClientCreateUrl(req, config)
	if err != nil {
		return nil, err
	}
	// We create the request from the body in second stage
	r, err := GetAllDataOfTheParticularProductOfferActionClientBuildRequest(req, u, config)
	if err != nil {
		return nil, err
	}
	// This one would execute the request and cast the result.
	return GetAllDataOfTheParticularProductOfferActionClientExecuteTyped(r)
}
