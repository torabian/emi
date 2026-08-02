//go:build integration

package external

import (
	"testing"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
)

// These exercise Entity1EntityGetFn/Entity1EntityQueryFn/Entity1EntityAwareDeletePreviewFn/
// Entity1EntityAwareDeleteFn end to end - see actions_postgres_test.go for
// openPostgresTestDB and why these need a real Postgres instance rather than sqlite.

func TestEntity1EntityGetFn_ReturnsMatchingRowByUniqueId(t *testing.T) {
	db := openPostgresTestDB(t)

	created, err := Entity1EntityActions.Create(db, &Entity1Entity{Title: "get-me"})
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	got, err := Entity1EntityActions.Get(db, created.UniqueId)
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	if got.Title != "get-me" {
		t.Fatalf("Title = %q, want %q", got.Title, "get-me")
	}
}

func TestEntity1EntityBrowseFn_FiltersSortsAndPagesWithTotalIgnoringPaging(t *testing.T) {
	db := openPostgresTestDB(t)

	for _, title := range []string{"apple-1", "apple-2", "apple-3", "banana-1"} {
		if _, err := Entity1EntityActions.Create(db, &Entity1Entity{Title: title}); err != nil {
			t.Fatalf("seed Create error: %v", err)
		}
	}

	items, meta, err := Entity1EntityActions.Browse(db, emigorm.QueryDSL{
		Filter:       `{"contains":[{"var":"title"},"apple"]}`,
		Sort:         "title asc",
		ItemsPerPage: 2,
	})
	if err != nil {
		t.Fatalf("Query error: %v", err)
	}
	if meta.TotalItems != 3 {
		t.Fatalf("TotalItems = %d, want 3 (paging must not affect the count)", meta.TotalItems)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items on the first page, got %d", len(items))
	}
	if items[0].Title != "apple-1" || items[1].Title != "apple-2" {
		t.Fatalf("expected sorted apple-1, apple-2, got %q, %q", items[0].Title, items[1].Title)
	}
	if meta.Cursor == nil {
		t.Fatal("expected a cursor since there's a further page to fetch")
	}

	// Following the returned cursor should resume right after the last item already
	// seen, regardless of StartIndex - the remaining "apple-3" row, not a second copy
	// of the first page.
	nextItems, nextMeta, err := Entity1EntityActions.Browse(db, emigorm.QueryDSL{
		Filter:       `{"contains":[{"var":"title"},"apple"]}`,
		Sort:         "title asc",
		ItemsPerPage: 2,
		Cursor:       *meta.Cursor,
	})
	if err != nil {
		t.Fatalf("Query (cursor) error: %v", err)
	}
	if nextMeta.TotalItems != 3 {
		t.Fatalf("TotalItems (cursor page) = %d, want 3", nextMeta.TotalItems)
	}
	if len(nextItems) != 1 || nextItems[0].Title != "apple-3" {
		t.Fatalf("expected exactly [apple-3] on the cursor page, got %+v", nextItems)
	}
}

func TestEntity1EntityAwareDeleteFn_RemovesArrayAndCollectionChildrenAlongsideTheRow(t *testing.T) {
	db := openPostgresTestDB(t)

	tagA := Entity2Entity{UniqueId: "aware-tag-a", Label2: "tag-a"}
	if err := db.Create(&tagA).Error; err != nil {
		t.Fatalf("seed tag error: %v", err)
	}

	dto := &Entity1Entity{
		Title: "to-be-deleted",
		Items: emigo.ArrayReplace([]Entity1EntityItems{
			{Item2: "child-1"},
			{Item2: "child-2"},
		}),
		Items3: emigo.CollectionReplace([]Entity2Entity{tagA}),
	}
	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	preview, err := Entity1EntityActions.AwareDeletePreview(db, []string{created.UniqueId})
	if err != nil {
		t.Fatalf("AwareDeletePreview error: %v", err)
	}
	if len(preview.Affected) != 2 {
		t.Fatalf("expected 2 affected relations (items, items3), got %+v", preview.Affected)
	}
	var itemsCount, items3Count int64
	for _, a := range preview.Affected {
		switch a.Relation {
		case "items":
			itemsCount = a.Count
		case "items3":
			items3Count = a.Count
		}
	}
	if itemsCount != 2 {
		t.Fatalf("expected 2 affected array children, got %d", itemsCount)
	}
	if items3Count != 1 {
		t.Fatalf("expected 1 affected collection row, got %d", items3Count)
	}

	if err := Entity1EntityActions.AwareDelete(db, []string{created.UniqueId}); err != nil {
		t.Fatalf("AwareDelete error: %v", err)
	}

	var remainingEntity int64
	db.Unscoped().Model(&Entity1Entity{}).Where("id = ?", created.Id).Count(&remainingEntity)
	if remainingEntity != 0 {
		t.Fatalf("expected the Entity1Entity row itself to be gone, found %d", remainingEntity)
	}

	var remainingItems int64
	db.Unscoped().Model(&Entity1EntityItems{}).Where("linker_id = ?", created.Id).Count(&remainingItems)
	if remainingItems != 0 {
		t.Fatalf("expected array children to be hard-deleted, found %d", remainingItems)
	}

	// The many-to-many join row must be gone, but the shared target itself must not be.
	var tagStillExists Entity2Entity
	if err := db.First(&tagStillExists, "unique_id = ?", "aware-tag-a").Error; err != nil {
		t.Fatalf("expected tag-a to still exist after AwareDelete: %v", err)
	}
}
