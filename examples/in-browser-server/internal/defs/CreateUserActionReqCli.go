//go:build !wasm

package defs

import "github.com/torabian/emi/emigo"

func GetCreateUserActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "first-name",
			Type:        "string",
			Description: "User's first name",
		},
		{
			Name:        prefix + "last-name",
			Type:        "string",
			Description: "User's last name",
		},
		{
			Name:        prefix + "birth-date",
			Type:        "string",
			Description: "User's birth date, ISO format YYYY-MM-DD",
		},
	}
}
func CastCreateUserActionReqFromCli(c emigo.CliCastable) CreateUserActionReq {
	data := CreateUserActionReq{}
	if c.IsSet("first-name") {
		data.FirstName = c.String("first-name")
	}
	if c.IsSet("last-name") {
		data.LastName = c.String("last-name")
	}
	if c.IsSet("birth-date") {
		data.BirthDate = c.String("birth-date")
	}
	return data
}
