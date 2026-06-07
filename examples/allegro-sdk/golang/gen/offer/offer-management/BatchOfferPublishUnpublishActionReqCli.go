//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

func GetBatchOfferPublishUnpublishActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "offer-criteria",
			Type: "array",
		},
		{
			Name:     prefix + "publication",
			Type:     "object",
			Children: GetBatchOfferPublishUnpublishActionReqPublicationCliFlags("publication-"),
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReq {
	data := BatchOfferPublishUnpublishActionReq{}
	if c.IsSet("offer-criteria") {
		data.OfferCriteria = emigo.CapturePossibleArray(CastBatchOfferPublishUnpublishActionReqOfferCriteriaFromCli, "offer-criteria", c)
	}
	if c.IsSet("publication") {
		data.Publication = CastBatchOfferPublishUnpublishActionReqPublicationFromCli(c)
	}
	return data
}
func GetBatchOfferPublishUnpublishActionReqOfferCriteriaCliFlags(prefix string) []emigo.CliFlag {
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
func CastBatchOfferPublishUnpublishActionReqOfferCriteriaFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqOfferCriteria {
	data := BatchOfferPublishUnpublishActionReqOfferCriteria{}
	if c.IsSet("offers") {
		data.Offers = emigo.CapturePossibleArray(CastBatchOfferPublishUnpublishActionReqOfferCriteriaOffersFromCli, "offers", c)
	}
	if c.IsSet("type") {
		data.Type = c.String("type")
	}
	return data
}
func GetBatchOfferPublishUnpublishActionReqOfferCriteriaOffersCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqOfferCriteriaOffersFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqOfferCriteriaOffers {
	data := BatchOfferPublishUnpublishActionReqOfferCriteriaOffers{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	return data
}
func GetBatchOfferPublishUnpublishActionReqPublicationCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "action",
			Type: "string",
		},
		{
			Name: prefix + "scheduled-for",
			Type: "string",
		},
	}
}
func CastBatchOfferPublishUnpublishActionReqPublicationFromCli(c emigo.CliCastable) BatchOfferPublishUnpublishActionReqPublication {
	data := BatchOfferPublishUnpublishActionReqPublication{}
	if c.IsSet("action") {
		data.Action = c.String("action")
	}
	if c.IsSet("scheduled-for") {
		data.ScheduledFor = c.String("scheduled-for")
	}
	return data
}
