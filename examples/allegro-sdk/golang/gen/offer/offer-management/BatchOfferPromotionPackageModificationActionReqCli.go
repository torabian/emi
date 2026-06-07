//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

func GetBatchOfferPromotionPackageModificationActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-criteria",
			Type: "array",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqModificationCliFlags("modification-"),
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReq {
	data := BatchOfferPromotionPackageModificationActionReq{}
	if c.IsSet("offer-criteria") {
		data.OfferCriteria = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaFromCli, "offer-criteria", c)
	}
	if c.IsSet("modification") {
		data.Modification = CastBatchOfferPromotionPackageModificationActionReqModificationFromCli(c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqOfferCriteriaCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offers",
			Type: "array",
		},
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqOfferCriteria {
	data := BatchOfferPromotionPackageModificationActionReqOfferCriteria{}
	if c.IsSet("offers") {
		data.Offers = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersFromCli, "offers", c)
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqOfferCriteriaOffersFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers {
	data := BatchOfferPromotionPackageModificationActionReqOfferCriteriaOffers{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqModificationBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name: prefix + "modification-time",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModification {
	data := BatchOfferPromotionPackageModificationActionReqModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastBatchOfferPromotionPackageModificationActionReqModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationBasePackageFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModificationBasePackage {
	data := BatchOfferPromotionPackageModificationActionReqModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqModificationExtraPackagesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqModificationExtraPackages {
	data := BatchOfferPromotionPackageModificationActionReqModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name:     prefix + "modification",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationCliFlags("modification-"),
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("modification") {
		data.Modification = CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationFromCli(c)
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "base-package",
			Type:     "object",
			Children: GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageCliFlags("base-package-"),
		},
		{
			Name: prefix + "extra-packages",
			Type: "array",
		},
		{
			Name: prefix + "modification-time",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModification{}
	if c.IsSet("base-package") {
		data.BasePackage = CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageFromCli(c)
	}
	if c.IsSet("extra-packages") {
		data.ExtraPackages = emigo.CapturePossibleArray(CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesFromCli, "extra-packages", c)
	}
	if c.IsSet("modification-time") {
		data.ModificationTime = c.String("modification-time")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackageFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationBasePackage{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
func GetBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackagesFromCli(c emigo.CliCastable) BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages {
	data := BatchOfferPromotionPackageModificationActionReqAdditionalMarketplacesModificationExtraPackages{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
