//go:build !wasm

package defs

import "github.com/torabian/emi/emigo"

func GetDeleteUserActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:        prefix + "id",
			Type:        "int",
			Description: "Id of the user to delete",
		},
	}
}
func CastDeleteUserActionReqFromCli(c emigo.CliCastable) DeleteUserActionReq {
	data := DeleteUserActionReq{}
	if c.IsSet("id") {
		data.Id = int(c.Int64("id"))
	}
	return data
}
