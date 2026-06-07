//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

func GetUpdateOfferTranslationActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "description",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqDescriptionCliFlags("description-"),
		},
		{
			Name:     prefix + "title",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqTitleCliFlags("title-"),
		},
		{
			Name:     prefix + "safety-information",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqSafetyInformationCliFlags("safety-information-"),
		},
	}
}
func CastUpdateOfferTranslationActionReqFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReq {
	data := UpdateOfferTranslationActionReq{}
	if c.IsSet("description") {
		data.Description = CastUpdateOfferTranslationActionReqDescriptionFromCli(c)
	}
	if c.IsSet("title") {
		data.Title = CastUpdateOfferTranslationActionReqTitleFromCli(c)
	}
	if c.IsSet("safety-information") {
		data.SafetyInformation = CastUpdateOfferTranslationActionReqSafetyInformationFromCli(c)
	}
	return data
}
func GetUpdateOfferTranslationActionReqDescriptionCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "translation",
			Type:     "object",
			Children: GetUpdateOfferTranslationActionReqDescriptionTranslationCliFlags("translation-"),
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescription {
	data := UpdateOfferTranslationActionReqDescription{}
	if c.IsSet("translation") {
		data.Translation = CastUpdateOfferTranslationActionReqDescriptionTranslationFromCli(c)
	}
	return data
}
func GetUpdateOfferTranslationActionReqDescriptionTranslationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "sections",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslation {
	data := UpdateOfferTranslationActionReqDescriptionTranslation{}
	if c.IsSet("sections") {
		data.Sections = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsFromCli, "sections", c)
	}
	return data
}
func GetUpdateOfferTranslationActionReqDescriptionTranslationSectionsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "items",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslationSections {
	data := UpdateOfferTranslationActionReqDescriptionTranslationSections{}
	if c.IsSet("items") {
		data.Items = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsFromCli, "items", c)
	}
	return data
}
func GetUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "type",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqDescriptionTranslationSectionsItemsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems {
	data := UpdateOfferTranslationActionReqDescriptionTranslationSectionsItems{}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}
func GetUpdateOfferTranslationActionReqTitleCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "translation",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqTitleFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqTitle {
	data := UpdateOfferTranslationActionReqTitle{}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	return data
}
func GetUpdateOfferTranslationActionReqSafetyInformationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "products",
			Type: "array",
		},
	}
}
func CastUpdateOfferTranslationActionReqSafetyInformationFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqSafetyInformation {
	data := UpdateOfferTranslationActionReqSafetyInformation{}
	if c.IsSet("products") {
		data.Products = emigo.CapturePossibleArray(CastUpdateOfferTranslationActionReqSafetyInformationProductsFromCli, "products", c)
	}
	return data
}
func GetUpdateOfferTranslationActionReqSafetyInformationProductsCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name: prefix + "translation",
			Type: "string",
		},
	}
}
func CastUpdateOfferTranslationActionReqSafetyInformationProductsFromCli(c emigo.CliCastable) UpdateOfferTranslationActionReqSafetyInformationProducts {
	data := UpdateOfferTranslationActionReqSafetyInformationProducts{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("translation") {
		data.Translation = c.String("translation")
	}
	return data
}
