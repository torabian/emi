package external

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/torabian/emi/public/allegro-sdk/golang/emigo"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"net/url"
)

/**
* Action to communicate with the action GetSellersOffersAction
 */
func GetSellersOffersActionMeta() struct {
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
		Name:        "GetSellersOffersAction",
		CliName:     "get sellers offers-action",
		URL:         "https://api.{environment}/sale/offers",
		Method:      "GET",
		Description: `Use this resource to get the list of the seller's offers. You can use different query parameters to filter the list. Read more: PL / EN.`,
	}
}
func GetGetSellersOffersActionResCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "count",
			Type: "int",
		},
		{
			Name: prefix + "total-count",
			Type: "int",
		},
		{
			Name: prefix + "offers",
			Type: "array",
		},
	}
}
func CastGetSellersOffersActionResFromCli(c emigo.CliCastable) GetSellersOffersActionRes {
	data := GetSellersOffersActionRes{}
	if c.IsSet("count") {
		data.Count = int(c.Int64("count"))
	}
	if c.IsSet("total-count") {
		data.TotalCount = int(c.Int64("total-count"))
	}
	if c.IsSet("offers") {
		data.Offers = emigo.CapturePossibleArray(CastGetSellersOffersActionResOffersFromCli, "offers", c)
	}
	return data
}

// The base class definition for getSellersOffersActionRes
type GetSellersOffersActionRes struct {
	// Number of offers in this page
	Count int `json:"count" yaml:"count"`
	// Total number of offers available
	TotalCount int                               `json:"totalCount" yaml:"totalCount"`
	Offers     []GetSellersOffersActionResOffers `json:"offers" yaml:"offers"`
}

func GetGetSellersOffersActionResOffersCliFlags(prefix string) []emigo.CliFlag {
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
			Name:     prefix + "category",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersCategoryCliFlags("category-"),
		},
		{
			Name:     prefix + "primary-image",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersPrimaryImageCliFlags("primary-image-"),
		},
		{
			Name:     prefix + "selling-mode",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModeCliFlags("selling-mode-"),
		},
		{
			Name:     prefix + "sale-info",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSaleInfoCliFlags("sale-info-"),
		},
		{
			Name:     prefix + "stock",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersStockCliFlags("stock-"),
		},
		{
			Name:     prefix + "stats",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersStatsCliFlags("stats-"),
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersPublicationCliFlags("publication-"),
		},
		{
			Name:     prefix + "after-sales-services",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersAfterSalesServicesCliFlags("after-sales-services-"),
		},
		{
			Name:     prefix + "additional-services",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersAdditionalServicesCliFlags("additional-services-"),
		},
		{
			Name:     prefix + "external",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersExternalCliFlags("external-"),
		},
		{
			Name:     prefix + "delivery",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersDeliveryCliFlags("delivery-"),
		},
		{
			Name:     prefix + "b2b",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersB2bCliFlags("b2b-"),
		},
		{
			Name:     prefix + "fundraising-campaign",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersFundraisingCampaignCliFlags("fundraising-campaign-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "map?",
		},
	}
}
func CastGetSellersOffersActionResOffersFromCli(c emigo.CliCastable) GetSellersOffersActionResOffers {
	data := GetSellersOffersActionResOffers{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	if c.IsSet("category") {
		data.Category = CastGetSellersOffersActionResOffersCategoryFromCli(c)
	}
	if c.IsSet("primary-image") {
		data.PrimaryImage = CastGetSellersOffersActionResOffersPrimaryImageFromCli(c)
	}
	if c.IsSet("selling-mode") {
		data.SellingMode = CastGetSellersOffersActionResOffersSellingModeFromCli(c)
	}
	if c.IsSet("sale-info") {
		data.SaleInfo = CastGetSellersOffersActionResOffersSaleInfoFromCli(c)
	}
	if c.IsSet("stock") {
		data.Stock = CastGetSellersOffersActionResOffersStockFromCli(c)
	}
	if c.IsSet("stats") {
		data.Stats = CastGetSellersOffersActionResOffersStatsFromCli(c)
	}
	if c.IsSet("publication") {
		data.Publication = CastGetSellersOffersActionResOffersPublicationFromCli(c)
	}
	if c.IsSet("after-sales-services") {
		data.AfterSalesServices = CastGetSellersOffersActionResOffersAfterSalesServicesFromCli(c)
	}
	if c.IsSet("additional-services") {
		data.AdditionalServices = CastGetSellersOffersActionResOffersAdditionalServicesFromCli(c)
	}
	if c.IsSet("external") {
		data.External = CastGetSellersOffersActionResOffersExternalFromCli(c)
	}
	if c.IsSet("delivery") {
		data.Delivery = CastGetSellersOffersActionResOffersDeliveryFromCli(c)
	}
	if c.IsSet("b2b") {
		data.B2b = CastGetSellersOffersActionResOffersB2bFromCli(c)
	}
	if c.IsSet("fundraising-campaign") {
		data.FundraisingCampaign = CastGetSellersOffersActionResOffersFundraisingCampaignFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		emigo.ParseNullable(c.String("additional-marketplaces"), &data.AdditionalMarketplaces)
	}
	return data
}

// The base class definition for offers
type GetSellersOffersActionResOffers struct {
	// Offer identifier
	Id string `json:"id" yaml:"id"`
	// Offer name or title
	Name                string                                             `json:"name" yaml:"name"`
	Category            GetSellersOffersActionResOffersCategory            `json:"category" yaml:"category"`
	PrimaryImage        GetSellersOffersActionResOffersPrimaryImage        `json:"primaryImage" yaml:"primaryImage"`
	SellingMode         GetSellersOffersActionResOffersSellingMode         `json:"sellingMode" yaml:"sellingMode"`
	SaleInfo            GetSellersOffersActionResOffersSaleInfo            `json:"saleInfo" yaml:"saleInfo"`
	Stock               GetSellersOffersActionResOffersStock               `json:"stock" yaml:"stock"`
	Stats               GetSellersOffersActionResOffersStats               `json:"stats" yaml:"stats"`
	Publication         GetSellersOffersActionResOffersPublication         `json:"publication" yaml:"publication"`
	AfterSalesServices  GetSellersOffersActionResOffersAfterSalesServices  `json:"afterSalesServices" yaml:"afterSalesServices"`
	AdditionalServices  GetSellersOffersActionResOffersAdditionalServices  `json:"additionalServices" yaml:"additionalServices"`
	External            GetSellersOffersActionResOffersExternal            `json:"external" yaml:"external"`
	Delivery            GetSellersOffersActionResOffersDelivery            `json:"delivery" yaml:"delivery"`
	B2b                 GetSellersOffersActionResOffersB2b                 `json:"b2b" yaml:"b2b"`
	FundraisingCampaign GetSellersOffersActionResOffersFundraisingCampaign `json:"fundraisingCampaign" yaml:"fundraisingCampaign"`
	// Marketplace-specific extensions for offer
	AdditionalMarketplaces emigo.Nullable[map[any]any] `json:"additionalMarketplaces" yaml:"additionalMarketplaces"`
}

func GetGetSellersOffersActionResOffersCategoryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersCategoryFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersCategory {
	data := GetSellersOffersActionResOffersCategory{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for category
type GetSellersOffersActionResOffersCategory struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersPrimaryImageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "url",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersPrimaryImageFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersPrimaryImage {
	data := GetSellersOffersActionResOffersPrimaryImage{}
	if c.IsSet("url") {
		data.Url = c.String("url")
	}
	return data
}

// The base class definition for primaryImage
type GetSellersOffersActionResOffersPrimaryImage struct {
	Url string `json:"url" yaml:"url"`
}

func GetGetSellersOffersActionResOffersSellingModeCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "format",
			Type: "string",
		},
		{
			Name:     prefix + "price",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModePriceCliFlags("price-"),
		},
		{
			Name:     prefix + "price-automation",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModePriceAutomationCliFlags("price-automation-"),
		},
		{
			Name:     prefix + "minimal-price",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModeMinimalPriceCliFlags("minimal-price-"),
		},
		{
			Name:     prefix + "starting-price",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModeStartingPriceCliFlags("starting-price-"),
		},
	}
}
func CastGetSellersOffersActionResOffersSellingModeFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingMode {
	data := GetSellersOffersActionResOffersSellingMode{}
	if c.IsSet("format") {
		data.Format = c.String("format")
	}
	if c.IsSet("price") {
		data.Price = CastGetSellersOffersActionResOffersSellingModePriceFromCli(c)
	}
	if c.IsSet("price-automation") {
		data.PriceAutomation = CastGetSellersOffersActionResOffersSellingModePriceAutomationFromCli(c)
	}
	if c.IsSet("minimal-price") {
		data.MinimalPrice = CastGetSellersOffersActionResOffersSellingModeMinimalPriceFromCli(c)
	}
	if c.IsSet("starting-price") {
		data.StartingPrice = CastGetSellersOffersActionResOffersSellingModeStartingPriceFromCli(c)
	}
	return data
}

// The base class definition for sellingMode
type GetSellersOffersActionResOffersSellingMode struct {
	Format          string                                                    `json:"format" yaml:"format"`
	Price           GetSellersOffersActionResOffersSellingModePrice           `json:"price" yaml:"price"`
	PriceAutomation GetSellersOffersActionResOffersSellingModePriceAutomation `json:"priceAutomation" yaml:"priceAutomation"`
	MinimalPrice    GetSellersOffersActionResOffersSellingModeMinimalPrice    `json:"minimalPrice" yaml:"minimalPrice"`
	StartingPrice   GetSellersOffersActionResOffersSellingModeStartingPrice   `json:"startingPrice" yaml:"startingPrice"`
}

func GetGetSellersOffersActionResOffersSellingModePriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetSellersOffersActionResOffersSellingModePriceFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingModePrice {
	data := GetSellersOffersActionResOffersSellingModePrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for price
type GetSellersOffersActionResOffersSellingModePrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSellersOffersActionResOffersSellingModePriceAutomationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "rule",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSellingModePriceAutomationRuleCliFlags("rule-"),
		},
	}
}
func CastGetSellersOffersActionResOffersSellingModePriceAutomationFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingModePriceAutomation {
	data := GetSellersOffersActionResOffersSellingModePriceAutomation{}
	if c.IsSet("rule") {
		data.Rule = CastGetSellersOffersActionResOffersSellingModePriceAutomationRuleFromCli(c)
	}
	return data
}

// The base class definition for priceAutomation
type GetSellersOffersActionResOffersSellingModePriceAutomation struct {
	Rule GetSellersOffersActionResOffersSellingModePriceAutomationRule `json:"rule" yaml:"rule"`
}

func GetGetSellersOffersActionResOffersSellingModePriceAutomationRuleCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersSellingModePriceAutomationRuleFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingModePriceAutomationRule {
	data := GetSellersOffersActionResOffersSellingModePriceAutomationRule{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for rule
type GetSellersOffersActionResOffersSellingModePriceAutomationRule struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersSellingModeMinimalPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetSellersOffersActionResOffersSellingModeMinimalPriceFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingModeMinimalPrice {
	data := GetSellersOffersActionResOffersSellingModeMinimalPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for minimalPrice
type GetSellersOffersActionResOffersSellingModeMinimalPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSellersOffersActionResOffersSellingModeStartingPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetSellersOffersActionResOffersSellingModeStartingPriceFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSellingModeStartingPrice {
	data := GetSellersOffersActionResOffersSellingModeStartingPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for startingPrice
type GetSellersOffersActionResOffersSellingModeStartingPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSellersOffersActionResOffersSaleInfoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "current-price",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersSaleInfoCurrentPriceCliFlags("current-price-"),
		},
		{
			Name: prefix + "bidders-count",
			Type: "int",
		},
	}
}
func CastGetSellersOffersActionResOffersSaleInfoFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSaleInfo {
	data := GetSellersOffersActionResOffersSaleInfo{}
	if c.IsSet("current-price") {
		data.CurrentPrice = CastGetSellersOffersActionResOffersSaleInfoCurrentPriceFromCli(c)
	}
	if c.IsSet("bidders-count") {
		data.BiddersCount = int(c.Int64("bidders-count"))
	}
	return data
}

// The base class definition for saleInfo
type GetSellersOffersActionResOffersSaleInfo struct {
	CurrentPrice GetSellersOffersActionResOffersSaleInfoCurrentPrice `json:"currentPrice" yaml:"currentPrice"`
	BiddersCount int                                                 `json:"biddersCount" yaml:"biddersCount"`
}

func GetGetSellersOffersActionResOffersSaleInfoCurrentPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastGetSellersOffersActionResOffersSaleInfoCurrentPriceFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersSaleInfoCurrentPrice {
	data := GetSellersOffersActionResOffersSaleInfoCurrentPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}

// The base class definition for currentPrice
type GetSellersOffersActionResOffersSaleInfoCurrentPrice struct {
	Amount   string `json:"amount" yaml:"amount"`
	Currency string `json:"currency" yaml:"currency"`
}

func GetGetSellersOffersActionResOffersStockCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "available",
			Type: "int",
		},
		{
			Name: prefix + "sold",
			Type: "int",
		},
	}
}
func CastGetSellersOffersActionResOffersStockFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersStock {
	data := GetSellersOffersActionResOffersStock{}
	if c.IsSet("available") {
		data.Available = int(c.Int64("available"))
	}
	if c.IsSet("sold") {
		data.Sold = int(c.Int64("sold"))
	}
	return data
}

// The base class definition for stock
type GetSellersOffersActionResOffersStock struct {
	Available int `json:"available" yaml:"available"`
	Sold      int `json:"sold" yaml:"sold"`
}

func GetGetSellersOffersActionResOffersStatsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "watchers-count",
			Type: "int",
		},
		{
			Name: prefix + "visits-count",
			Type: "int",
		},
	}
}
func CastGetSellersOffersActionResOffersStatsFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersStats {
	data := GetSellersOffersActionResOffersStats{}
	if c.IsSet("watchers-count") {
		data.WatchersCount = int(c.Int64("watchers-count"))
	}
	if c.IsSet("visits-count") {
		data.VisitsCount = int(c.Int64("visits-count"))
	}
	return data
}

// The base class definition for stats
type GetSellersOffersActionResOffersStats struct {
	WatchersCount int `json:"watchersCount" yaml:"watchersCount"`
	VisitsCount   int `json:"visitsCount" yaml:"visitsCount"`
}

func GetGetSellersOffersActionResOffersPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "status",
			Type: "string",
		},
		{
			Name: prefix + "starting-at",
			Type: "string",
		},
		{
			Name: prefix + "started-at",
			Type: "string",
		},
		{
			Name: prefix + "ending-at",
			Type: "string",
		},
		{
			Name: prefix + "ended-at",
			Type: "string",
		},
		{
			Name:     prefix + "marketplaces",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersPublicationMarketplacesCliFlags("marketplaces-"),
		},
	}
}
func CastGetSellersOffersActionResOffersPublicationFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersPublication {
	data := GetSellersOffersActionResOffersPublication{}
	if c.IsSet("status") {
		data.Status = c.String("status")
	}
	if c.IsSet("starting-at") {
		data.StartingAt = c.String("starting-at")
	}
	if c.IsSet("started-at") {
		data.StartedAt = c.String("started-at")
	}
	if c.IsSet("ending-at") {
		data.EndingAt = c.String("ending-at")
	}
	if c.IsSet("ended-at") {
		data.EndedAt = c.String("ended-at")
	}
	if c.IsSet("marketplaces") {
		data.Marketplaces = CastGetSellersOffersActionResOffersPublicationMarketplacesFromCli(c)
	}
	return data
}

// The base class definition for publication
type GetSellersOffersActionResOffersPublication struct {
	Status       string                                                 `json:"status" yaml:"status"`
	StartingAt   string                                                 `json:"startingAt" yaml:"startingAt"`
	StartedAt    string                                                 `json:"startedAt" yaml:"startedAt"`
	EndingAt     string                                                 `json:"endingAt" yaml:"endingAt"`
	EndedAt      string                                                 `json:"endedAt" yaml:"endedAt"`
	Marketplaces GetSellersOffersActionResOffersPublicationMarketplaces `json:"marketplaces" yaml:"marketplaces"`
}

func GetGetSellersOffersActionResOffersPublicationMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersPublicationMarketplacesBaseCliFlags("base-"),
		},
		{
			Name: prefix + "additional",
			Type: "array",
		},
	}
}
func CastGetSellersOffersActionResOffersPublicationMarketplacesFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersPublicationMarketplaces {
	data := GetSellersOffersActionResOffersPublicationMarketplaces{}
	if c.IsSet("base") {
		data.Base = CastGetSellersOffersActionResOffersPublicationMarketplacesBaseFromCli(c)
	}
	if c.IsSet("additional") {
		data.Additional = emigo.CapturePossibleArray(CastGetSellersOffersActionResOffersPublicationMarketplacesAdditionalFromCli, "additional", c)
	}
	return data
}

// The base class definition for marketplaces
type GetSellersOffersActionResOffersPublicationMarketplaces struct {
	Base       GetSellersOffersActionResOffersPublicationMarketplacesBase         `json:"base" yaml:"base"`
	Additional []GetSellersOffersActionResOffersPublicationMarketplacesAdditional `json:"additional" yaml:"additional"`
}

func GetGetSellersOffersActionResOffersPublicationMarketplacesBaseCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersPublicationMarketplacesBaseFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersPublicationMarketplacesBase {
	data := GetSellersOffersActionResOffersPublicationMarketplacesBase{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for base
type GetSellersOffersActionResOffersPublicationMarketplacesBase struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersPublicationMarketplacesAdditionalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersPublicationMarketplacesAdditionalFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersPublicationMarketplacesAdditional {
	data := GetSellersOffersActionResOffersPublicationMarketplacesAdditional{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additional
type GetSellersOffersActionResOffersPublicationMarketplacesAdditional struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersAfterSalesServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "implied-warranty",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersAfterSalesServicesImpliedWarrantyCliFlags("implied-warranty-"),
		},
		{
			Name:     prefix + "return-policy",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersAfterSalesServicesReturnPolicyCliFlags("return-policy-"),
		},
		{
			Name:     prefix + "warranty",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersAfterSalesServicesWarrantyCliFlags("warranty-"),
		},
	}
}
func CastGetSellersOffersActionResOffersAfterSalesServicesFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersAfterSalesServices {
	data := GetSellersOffersActionResOffersAfterSalesServices{}
	if c.IsSet("implied-warranty") {
		data.ImpliedWarranty = CastGetSellersOffersActionResOffersAfterSalesServicesImpliedWarrantyFromCli(c)
	}
	if c.IsSet("return-policy") {
		data.ReturnPolicy = CastGetSellersOffersActionResOffersAfterSalesServicesReturnPolicyFromCli(c)
	}
	if c.IsSet("warranty") {
		data.Warranty = CastGetSellersOffersActionResOffersAfterSalesServicesWarrantyFromCli(c)
	}
	return data
}

// The base class definition for afterSalesServices
type GetSellersOffersActionResOffersAfterSalesServices struct {
	ImpliedWarranty GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty `json:"impliedWarranty" yaml:"impliedWarranty"`
	ReturnPolicy    GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy    `json:"returnPolicy" yaml:"returnPolicy"`
	Warranty        GetSellersOffersActionResOffersAfterSalesServicesWarranty        `json:"warranty" yaml:"warranty"`
}

func GetGetSellersOffersActionResOffersAfterSalesServicesImpliedWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersAfterSalesServicesImpliedWarrantyFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty {
	data := GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for impliedWarranty
type GetSellersOffersActionResOffersAfterSalesServicesImpliedWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersAfterSalesServicesReturnPolicyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersAfterSalesServicesReturnPolicyFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy {
	data := GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for returnPolicy
type GetSellersOffersActionResOffersAfterSalesServicesReturnPolicy struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersAfterSalesServicesWarrantyCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersAfterSalesServicesWarrantyFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersAfterSalesServicesWarranty {
	data := GetSellersOffersActionResOffersAfterSalesServicesWarranty{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for warranty
type GetSellersOffersActionResOffersAfterSalesServicesWarranty struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersAdditionalServicesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersAdditionalServicesFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersAdditionalServices {
	data := GetSellersOffersActionResOffersAdditionalServices{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for additionalServices
type GetSellersOffersActionResOffersAdditionalServices struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersExternalCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersExternalFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersExternal {
	data := GetSellersOffersActionResOffersExternal{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for external
type GetSellersOffersActionResOffersExternal struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersDeliveryCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "shipping-rates",
			Type:     "object",
			Children: GetGetSellersOffersActionResOffersDeliveryShippingRatesCliFlags("shipping-rates-"),
		},
	}
}
func CastGetSellersOffersActionResOffersDeliveryFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersDelivery {
	data := GetSellersOffersActionResOffersDelivery{}
	if c.IsSet("shipping-rates") {
		data.ShippingRates = CastGetSellersOffersActionResOffersDeliveryShippingRatesFromCli(c)
	}
	return data
}

// The base class definition for delivery
type GetSellersOffersActionResOffersDelivery struct {
	ShippingRates GetSellersOffersActionResOffersDeliveryShippingRates `json:"shippingRates" yaml:"shippingRates"`
}

func GetGetSellersOffersActionResOffersDeliveryShippingRatesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersDeliveryShippingRatesFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersDeliveryShippingRates {
	data := GetSellersOffersActionResOffersDeliveryShippingRates{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for shippingRates
type GetSellersOffersActionResOffersDeliveryShippingRates struct {
	Id string `json:"id" yaml:"id"`
}

func GetGetSellersOffersActionResOffersB2bCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "buyable-only-by-business",
			Type: "bool",
		},
	}
}
func CastGetSellersOffersActionResOffersB2bFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersB2b {
	data := GetSellersOffersActionResOffersB2b{}
	if c.IsSet("buyable-only-by-business") {
		data.BuyableOnlyByBusiness = bool(c.Bool("buyable-only-by-business"))
	}
	return data
}

// The base class definition for b2b
type GetSellersOffersActionResOffersB2b struct {
	BuyableOnlyByBusiness bool `json:"buyableOnlyByBusiness" yaml:"buyableOnlyByBusiness"`
}

func GetGetSellersOffersActionResOffersFundraisingCampaignCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastGetSellersOffersActionResOffersFundraisingCampaignFromCli(c emigo.CliCastable) GetSellersOffersActionResOffersFundraisingCampaign {
	data := GetSellersOffersActionResOffersFundraisingCampaign{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}

// The base class definition for fundraisingCampaign
type GetSellersOffersActionResOffersFundraisingCampaign struct {
	Id string `json:"id" yaml:"id"`
}

func (x *GetSellersOffersActionRes) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}

type GetSellersOffersActionResponse struct {
	StatusCode int
	Headers    map[string]string
	Payload    interface{}
}

func (x *GetSellersOffersActionResponse) SetContentType(contentType string) *GetSellersOffersActionResponse {
	if x.Headers == nil {
		x.Headers = make(map[string]string)
	}
	x.Headers["Content-Type"] = contentType
	return x
}
func (x *GetSellersOffersActionResponse) AsStream(r io.Reader, contentType string) *GetSellersOffersActionResponse {
	x.Payload = r
	x.SetContentType(contentType)
	return x
}
func (x *GetSellersOffersActionResponse) AsJSON(payload any) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/json")
	return x
}
func (x *GetSellersOffersActionResponse) AsHTML(payload string) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("text/html; charset=utf-8")
	return x
}
func (x *GetSellersOffersActionResponse) AsBytes(payload []byte) *GetSellersOffersActionResponse {
	x.Payload = payload
	x.SetContentType("application/octet-stream")
	return x
}
func (x GetSellersOffersActionResponse) GetStatusCode() int {
	return x.StatusCode
}
func (x GetSellersOffersActionResponse) GetRespHeaders() map[string]string {
	return x.Headers
}
func (x GetSellersOffersActionResponse) GetPayload() interface{} {
	return x.Payload
}

// GetSellersOffersActionRaw registers a raw Gin route for the GetSellersOffersAction action.
// This gives the developer full control over middleware, handlers, and response handling.
func GetSellersOffersActionRaw(r *gin.Engine, handlers ...gin.HandlerFunc) {
	meta := GetSellersOffersActionMeta()
	r.Handle(meta.Method, meta.URL, handlers...)
}

type GetSellersOffersActionRequestSig = func(c GetSellersOffersActionRequest) (*GetSellersOffersActionResponse, error)

// GetSellersOffersActionHandler returns the HTTP method, route URL, and a typed Gin handler for the GetSellersOffersAction action.
// Developers implement their business logic as a function that receives a typed request object
// and returns either an *ActionResponse or nil. JSON marshalling, headers, and errors are handled automatically.
func GetSellersOffersActionHandler(
	handler GetSellersOffersActionRequestSig,
) (method, url string, h gin.HandlerFunc) {
	meta := GetSellersOffersActionMeta()
	return meta.Method, meta.URL, func(m *gin.Context) {
		// Build typed request wrapper
		req := GetSellersOffersActionRequest{
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

// GetSellersOffersAction is a high-level convenience wrapper around GetSellersOffersActionHandler.
// It automatically constructs and registers the typed route on the Gin engine.
// Use this when you don't need custom middleware or route grouping.
func GetSellersOffersActionGin(r gin.IRoutes, handler GetSellersOffersActionRequestSig) {
	method, url, h := GetSellersOffersActionHandler(handler)
	r.Handle(method, url, h)
}

/**
 * Query parameters for Get sellers offersAction
 */
// Query wrapper with private fields
type GetSellersOffersActionQuery struct {
	values url.Values
	mapped map[string]interface{}
	// Typesafe fields
}

func GetSellersOffersActionQueryFromString(rawQuery string) GetSellersOffersActionQuery {
	v := GetSellersOffersActionQuery{}
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
func GetSellersOffersActionQueryFromGin(c *gin.Context) GetSellersOffersActionQuery {
	return GetSellersOffersActionQueryFromString(c.Request.URL.RawQuery)
}
func GetSellersOffersActionQueryFromHttp(r *http.Request) GetSellersOffersActionQuery {
	return GetSellersOffersActionQueryFromString(r.URL.RawQuery)
}
func (q GetSellersOffersActionQuery) Values() url.Values {
	return q.values
}
func (q GetSellersOffersActionQuery) Mapped() map[string]interface{} {
	return q.mapped
}
func (q *GetSellersOffersActionQuery) SetValues(v url.Values) {
	q.values = v
}
func (q *GetSellersOffersActionQuery) SetMapped(m map[string]interface{}) {
	q.mapped = m
}

type GetSellersOffersActionRequest struct {
	QueryParams url.Values
	Headers     http.Header
	GinCtx      *gin.Context
	CliCtx      *cli.Context
}
type GetSellersOffersActionResult struct {
	resp    *http.Response // embed original response
	Payload interface{}
}

func GetSellersOffersActionCall(
	req GetSellersOffersActionRequest,
	config *emigo.APIClient, // optional pre-built request
) (*GetSellersOffersActionResult, error) {
	var httpReq *http.Request
	if config == nil || config.Httpr == nil {
		meta := GetSellersOffersActionMeta()
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
		req0, err := http.NewRequest(meta.Method, u.String(), nil)
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
	var result GetSellersOffersActionResult
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
