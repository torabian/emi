// Thin smoke-test binary. Full scenarios live in vsql_test/tests/ —
// run them with `make test`.
package main

import (
	"fmt"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/vsql_test/dto"
	"github.com/torabian/emi/vsql_test/sqlfiles"
	"github.com/torabian/emi/vsql_test/vsql"
)

func main() {
	user := dto.CreateUserDto{
		Id:    1,
		Name:  "Ali O'Brien",
		Email: emigo.NullableOf("ali@example.com"),
		Age:   34,
		Preferences: dto.CreateUserDtoPreferences{
			Theme:  "dark",
			Locale: "en-US",
		},
		Tags: []dto.CreateUserDtoTags{
			{Key: "role", Value: "admin"},
			{Key: "plan", Value: "pro"},
		},
		Address: dto.CreateUserDtoAddress{
			Id:       100,
			UserId:   1,
			Street:   "123 Main St",
			City:     "Warsaw",
			Location: dto.GeoPoint{Lat: 52.2297, Lon: 21.0122, Valid: true},
		},
	}

	out, err := vsql.Render(sqlfiles.Files, "insert_user_address.sql", user)
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}
