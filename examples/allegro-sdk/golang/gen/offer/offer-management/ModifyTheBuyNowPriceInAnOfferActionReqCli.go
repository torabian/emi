//go:build !wasm

package external

import "github.com/torabian/emi/public/allegro-sdk/golang/emigo"

func GetModifyTheBuyNowPriceInAnOfferActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "string",
		},
		{
			Name:     prefix + "input",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionReqInputCliFlags("input-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionReqFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReq {
	data := ModifyTheBuyNowPriceInAnOfferActionReq{}
	if c.IsSet("id") {
		data.Id = c.String("id")
	}
	if c.IsSet("input") {
		data.Input = CastModifyTheBuyNowPriceInAnOfferActionReqInputFromCli(c)
	}
	return data
}
func GetModifyTheBuyNowPriceInAnOfferActionReqInputCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "buy-now-price",
			Type:     "object",
			Children: GetModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceCliFlags("buy-now-price-"),
		},
	}
}
func CastModifyTheBuyNowPriceInAnOfferActionReqInputFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReqInput {
	data := ModifyTheBuyNowPriceInAnOfferActionReqInput{}
	if c.IsSet("buy-now-price") {
		data.BuyNowPrice = CastModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceFromCli(c)
	}
	return data
}
func GetModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceCliFlags(prefix string) []emigo.CliFlag {
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
func CastModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPriceFromCli(c emigo.CliCastable) ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice {
	data := ModifyTheBuyNowPriceInAnOfferActionReqInputBuyNowPrice{}
	if c.IsSet("amount") {
		data.Amount = c.String("amount")
	}
	if c.IsSet("currency") {
		data.Currency = c.String("currency")
	}
	return data
}
