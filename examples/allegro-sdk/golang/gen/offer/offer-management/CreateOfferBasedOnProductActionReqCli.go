//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

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
