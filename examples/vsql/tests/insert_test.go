package tests

import (
	"testing"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/examples/vsql/dto"
	"github.com/torabian/emi/examples/vsql/sqlfiles"
	"github.com/torabian/emi/examples/vsql/vsql"
)

func TestInsertUserAddress(t *testing.T) {
	user := dto.CreateUserDto{
		Id:    1,
		Name:  "Ali O'Brien",
		Email: emigo.NullableOf("ali@example.com"),
		Age:   34,
		Preferences: dto.CreateUserDtoPreferences{
			Theme:  "dark",
			Locale: "en-US",
		},
		Tags: emigo.ArrayReplace([]dto.CreateUserDtoTags{
			{Key: "role", Value: "admin"},
			{Key: "plan", Value: "pro"},
		}),
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
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL ---\n%s", out)
}

func TestInsertUserNoTags(t *testing.T) {
	// Empty Tags slice — the {{ if .Tags }} guard should drop the whole
	// user_tags block, leaving only the users + addresses inserts.
	user := dto.CreateUserDto{
		Id:   2,
		Name: "Empty Tags",
		Age:  20,
		Preferences: dto.CreateUserDtoPreferences{
			Theme:  "light",
			Locale: "en-GB",
		},
		Address: dto.CreateUserDtoAddress{
			Id:       200,
			UserId:   2,
			Street:   "1 Empty St",
			City:     "Krakow",
			Location: dto.GeoPoint{Lat: 50.0647, Lon: 19.9450, Valid: true},
		},
	}

	out, err := vsql.Render(sqlfiles.Files, "insert_user_address.sql", user)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL ---\n%s", out)
}
