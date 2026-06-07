//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

func GetModifyOfferPromotionPackagesActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modifications",
			Type: "array",
		},
		{
			Name: prefix + "additional-marketplaces",
			Type: "array",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReq {
	data := ModifyOfferPromotionPackagesActionReq{}
	if c.IsSet("modifications") {
		data.Modifications = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqModificationsFromCli, "modifications", c)
	}
	if c.IsSet("additional-marketplaces") {
		data.AdditionalMarketplaces = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesFromCli, "additional-marketplaces", c)
	}
	return data
}
func GetModifyOfferPromotionPackagesActionReqModificationsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modification-type",
			Type: "string",
		},
		{
			Name: prefix + "package-type",
			Type: "string",
		},
		{
			Name: prefix + "package-id",
			Type: "string",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqModificationsFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqModifications {
	data := ModifyOfferPromotionPackagesActionReqModifications{}
	if c.IsSet("modification-type") {
		data.ModificationType = c.String("modification-type")
	}
	if c.IsSet("package-type") {
		data.PackageType = c.String("package-type")
	}
	if c.IsSet("package-id") {
		data.PackageId = c.String("package-id")
	}
	return data
}
func GetModifyOfferPromotionPackagesActionReqAdditionalMarketplacesCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "marketplace-id",
			Type: "string",
		},
		{
			Name: prefix + "modifications",
			Type: "array",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces {
	data := ModifyOfferPromotionPackagesActionReqAdditionalMarketplaces{}
	if c.IsSet("marketplace-id") {
		data.MarketplaceId = c.String("marketplace-id")
	}
	if c.IsSet("modifications") {
		data.Modifications = emigo.CapturePossibleArray(CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsFromCli, "modifications", c)
	}
	return data
}
func GetModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "modification-type",
			Type: "string",
		},
		{
			Name: prefix + "package-type",
			Type: "string",
		},
		{
			Name: prefix + "package-id",
			Type: "string",
		},
	}
}
func CastModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModificationsFromCli(c emigo.CliCastable) ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications {
	data := ModifyOfferPromotionPackagesActionReqAdditionalMarketplacesModifications{}
	if c.IsSet("modification-type") {
		data.ModificationType = c.String("modification-type")
	}
	if c.IsSet("package-type") {
		data.PackageType = c.String("package-type")
	}
	if c.IsSet("package-id") {
		data.PackageId = c.String("package-id")
	}
	return data
}
