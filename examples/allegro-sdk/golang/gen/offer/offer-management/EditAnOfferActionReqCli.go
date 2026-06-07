//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

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
