//go:build !wasm

package defs

import "github.com/torabian/emi/emigo"

func GetSubstringActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "input",
			Type:        "string",
			Description: "The string you want to do substring",
		},
		{
			Name:        prefix + "start",
			Type:        "int",
			Description: "Start position",
		},
		{
			Name:        prefix + "end",
			Type:        "int",
			Description: "End position",
		},
	}
}
func CastSubstringActionReqFromCli(c emigo.CliCastable) SubstringActionReq {
	data := SubstringActionReq{}
	if c.IsSet("input") {
		data.Input = c.String("input")
	}
	if c.IsSet("start") {
		data.Start = int(c.Int64("start"))
	}
	if c.IsSet("end") {
		data.End = int(c.Int64("end"))
	}
	return data
}
