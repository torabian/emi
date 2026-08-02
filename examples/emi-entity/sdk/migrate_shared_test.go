//go:build integration

package external

import (
	"testing"

	"gorm.io/gorm"
)

// runAutoMigrateAssertions is shared by migrate_postgres_test.go and
// migrate_mysql_test.go: both drivers run the exact same checks against the exact same
// generated structs, only the connection differs.
func runAutoMigrateAssertions(t *testing.T, db *gorm.DB) {
	t.Helper()

	// Entity2Entity has to exist before Entity1Entity, since Entity1Entity's OwnerRow/
	// ManagerRow/Items3Row/Items4Row all foreign-key or many2many against it.
	// Entity1EntityItems/Items2 (the has-many array child rows) also need to be listed
	// explicitly - gorm's AutoMigrate does not auto-discover child row types just
	// because a parent has an association field pointing at them (verified against a
	// real run: without this, ReconcileHasMany fails with "no such table").
	if err := db.AutoMigrate(&Entity2Entity{}, &Entity1Entity{}, &Entity1EntityItems{}, &Entity1EntityItems2{}); err != nil {
		t.Fatalf("AutoMigrate error: %v", err)
	}

	m := db.Migrator()

	t.Run("scalar and object/map/slice/any columns exist", func(t *testing.T) {
		for _, col := range []string{
			"UniqueId", "Title", "Subtitle", "IsActive", "IsFeatured",
			"ViewCount", "ViewCountOpt", "SmallCount", "BigCount",
			"Ratio32", "Ratio64", "Status", "StatusOpt",
			"Metadata", "MetadataOpt", "RawSettings", "Labels", "LabelsOpt", "Misc",
		} {
			if !m.HasColumn(&Entity1Entity{}, col) {
				t.Errorf("expected entity1_table to have column for field %s", col)
			}
		}
	})

	t.Run("object/object? are embedded inline (no separate table)", func(t *testing.T) {
		if !m.HasColumn(&Entity1Entity{}, "Item1") {
			t.Error("expected content1/content2's embedded Item1 column on entity1_table")
		}
	})

	t.Run("one/one? produced real FK columns, not the DTO wrapper", func(t *testing.T) {
		if !m.HasColumn(&Entity1Entity{}, "OwnerId") {
			t.Error("expected OwnerId FK column on entity1_table")
		}
		if !m.HasColumn(&Entity1Entity{}, "ManagerId") {
			t.Error("expected ManagerId FK column on entity1_table")
		}
		if m.HasColumn(&Entity1Entity{}, "Owner") {
			t.Error("did not expect a column for the Owner DTO field itself (gorm:\"-\")")
		}
	})

	t.Run("array/array? produced real has-many child tables with LinkerId", func(t *testing.T) {
		if !m.HasTable(&Entity1EntityItems{}) {
			t.Error("expected a table for Entity1EntityItems (has-many child rows)")
		}
		if !m.HasColumn(&Entity1EntityItems{}, "LinkerId") {
			t.Error("expected LinkerId FK column on the Entity1EntityItems child table")
		}
		if !m.HasTable(&Entity1EntityItems2{}) {
			t.Error("expected a table for Entity1EntityItems2 (has-many child rows)")
		}
	})

	t.Run("collection/collection? produced a many2many join table", func(t *testing.T) {
		if !m.HasTable("entity1_items3") {
			t.Error("expected join table entity1_items3 for the Items3 many2many relation")
		}
		if !m.HasTable("entity1_items4") {
			t.Error("expected join table entity1_items4 for the Items4 many2many relation")
		}
	})

	t.Run("explicit tags.gorm override on rawSettings is respected", func(t *testing.T) {
		if !m.HasColumn(&Entity1Entity{}, "RawSettings") {
			t.Error("expected RawSettings column on entity1_table (explicit tags.gorm override still migrates)")
		}
	})

	t.Run("insert and read back a row end to end", func(t *testing.T) {
		row := &Entity1Entity{
			UniqueId:  "e2e-test-1",
			Title:     "hello",
			IsActive:  true,
			ViewCount: 42,
			Ratio64:   3.14,
			Status:    "active",
			Metadata:  map[string]string{"k": "v"},
			Labels:    []string{"a", "b"},
			Misc:      map[string]any{"nested": true},
			Complex1:  Money{Cents: 999},
		}

		if err := db.Create(row).Error; err != nil {
			t.Fatalf("Create error: %v", err)
		}

		var loaded Entity1Entity
		if err := db.First(&loaded, "unique_id = ?", "e2e-test-1").Error; err != nil {
			t.Fatalf("First error: %v", err)
		}

		if loaded.Title != "hello" || !loaded.IsActive || loaded.ViewCount != 42 {
			t.Fatalf("unexpected roundtrip values: %+v", loaded)
		}
		if loaded.Metadata["k"] != "v" {
			t.Fatalf("expected Metadata to roundtrip through serializer:json, got %+v", loaded.Metadata)
		}
		if len(loaded.Labels) != 2 || loaded.Labels[0] != "a" {
			t.Fatalf("expected Labels to roundtrip through serializer:json, got %+v", loaded.Labels)
		}
		if loaded.Complex1.Cents != 999 {
			t.Fatalf("expected Complex1 to roundtrip through Money's Value()/Scan(), got %+v", loaded.Complex1)
		}
	})
}
