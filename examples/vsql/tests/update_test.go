package tests

import (
	"testing"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/vsql_test/dto"
	"github.com/torabian/emi/vsql_test/sqlfiles"
	"github.com/torabian/emi/vsql_test/vsql"
)

// TestUpdateReplaceTags drives the Collection "replace" path: the SQL must
// emit a DELETE on user_tags before the per-row inserts so the existing
// rows are wiped first.
func TestUpdateReplaceTags(t *testing.T) {
	patch := dto.UpdateUserDto{
		Id:   emigo.NullableOf[int64](1),
		Name: emigo.NullableOf("Ali Renamed"),
		Age:  emigo.NullableOf(35),
		Tags: emigo.CollectionReplace([]dto.UpdateUserDtoTags{
			{Key: "plan", Value: "enterprise"},
		}),
	}

	out, err := vsql.Render(sqlfiles.Files, "update_user.sql", patch)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL ---\n%s", out)
}

// TestUpdateAppendTags drives the "append" path: existing rows are kept,
// only the new tag is upserted via ON CONFLICT.
func TestUpdateAppendTags(t *testing.T) {
	patch := dto.UpdateUserDto{
		Id: emigo.NullableOf[int64](1),
		Tags: emigo.CollectionAppend([]dto.UpdateUserDtoTags{
			{Key: "feature-flag", Value: "beta"},
		}),
	}

	out, err := vsql.Render(sqlfiles.Files, "update_user.sql", patch)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL ---\n%s", out)
}

// TestUpdateScalarsOnly demonstrates the cleanest PATCH: only scalars set,
// no Tags wrapper. The user_tags block should disappear entirely.
func TestUpdateScalarsOnly(t *testing.T) {
	patch := dto.UpdateUserDto{
		Id:    emigo.NullableOf[int64](1),
		Email: emigo.NullableOf("new@example.com"),
	}

	out, err := vsql.Render(sqlfiles.Files, "update_user.sql", patch)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	t.Logf("\n--- generated SQL ---\n%s", out)
}
