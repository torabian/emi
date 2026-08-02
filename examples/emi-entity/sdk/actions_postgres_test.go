//go:build integration

package external

import (
	"os"
	"testing"

	"github.com/torabian/emi/emigo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// These exercise the actual generated Entity1EntityCreateFn/Entity1EntityUpdateFn (via
// Entity1EntityActions) end to end. They need a real Postgres instance rather than the
// sqlite one lib/golang/go-entity-default-fields.go's uniqueId column relies on
// (gorm:"type:uuid;default:gen_random_uuid()") - Postgres 13+ has gen_random_uuid()
// built in, sqlite/mysql don't have an equivalent, so there's no fast local-only tier
// for these anymore (see ../docker-compose.yml (make entity-db-up) or override
// EMI_ENTITY_POSTGRES_DSN). Run with: make entity-migration-test-postgres

func openPostgresTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := os.Getenv("EMI_ENTITY_POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost port=55432 user=emi password=emi dbname=emi_entity_test sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		t.Fatalf("failed to connect to postgres (is `make entity-db-up` running? dsn=%q): %v", dsn, err)
	}

	// gorm's AutoMigrate does *not* auto-discover has-many child row types through an
	// association field (Entity1Entity.Items) - each has to be listed explicitly, or
	// their tables never get created. This includes relations nested inside
	// object/object? containers (NestedContainer.NestedInner.NestedItems and friends).
	if err := db.AutoMigrate(
		&Entity2Entity{},
		&Entity1Entity{},
		&Entity1EntityItems{},
		&Entity1EntityItems2{},
		&Entity1EntityNestedContainerNestedInnerNestedItems{},
		&Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt{},
	); err != nil {
		t.Fatalf("AutoMigrate error: %v", err)
	}

	return db
}

func TestEntity1EntityCreateFn_CascadesArrayCollectionOne(t *testing.T) {
	db := openPostgresTestDB(t)

	owner := &Entity2Entity{UniqueId: "owner-1", Label2: "owner-one"}
	tagA := &Entity2Entity{UniqueId: "tag-a", Label2: "tag-a"}
	tagB := &Entity2Entity{UniqueId: "tag-b", Label2: "tag-b"}
	if err := db.Create(owner).Error; err != nil {
		t.Fatalf("seed owner error: %v", err)
	}
	if err := db.Create(tagA).Error; err != nil {
		t.Fatalf("seed tagA error: %v", err)
	}
	if err := db.Create(tagB).Error; err != nil {
		t.Fatalf("seed tagB error: %v", err)
	}

	dto := &Entity1Entity{
		Title:    "hello",
		IsActive: true,
		// Items is now a plain, gorm-native has-many ([]*Entity1EntityItems) on the
		// entity struct itself - no more emigo.ArrayReplace wrapper - Operation only
		// ever mattered for Update's reconcile-against-existing-rows semantics, which
		// a fresh Create has nothing to reconcile against anyway.
		Items: []*Entity1EntityItems{
			{Item2: "child-1"},
			{Item2: "child-2"},
		},
		Items3: emigo.CollectionReplace([]Entity2Entity{*tagA, *tagB}),
		// Owner is now a plain, gorm-native belongs-to (*Entity2Entity) on the entity
		// struct itself - no more emigo.OneSelect wrapper. Passing the
		// already-persisted owner (its Id already set from the db.Create above) lets
		// gorm's own Create() cascade recognize the existing row and just wire
		// OwnerId to it, rather than inserting a duplicate.
		Owner: owner,
	}

	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}
	if created.UniqueId == "" {
		t.Fatal("expected the database to assign a UniqueId via gen_random_uuid()")
	}
	if created.Id == 0 {
		t.Fatal("expected the database to assign an Id via autoIncrement")
	}

	var reloaded Entity1Entity
	if err := db.Preload("Items").Preload("Items3Row").Preload("Owner").
		First(&reloaded, "unique_id = ?", created.UniqueId).Error; err != nil {
		t.Fatalf("reload error: %v", err)
	}

	if len(reloaded.Items) != 2 {
		t.Fatalf("expected 2 array children, got %d", len(reloaded.Items))
	}
	if len(reloaded.Items3Row) != 2 {
		t.Fatalf("expected 2 collection targets, got %d", len(reloaded.Items3Row))
	}
	if reloaded.Owner == nil || reloaded.Owner.Label2 != "owner-one" {
		t.Fatalf("expected owner to resolve to owner-one, got %+v", reloaded.Owner)
	}
	if reloaded.OwnerId != owner.Id {
		t.Fatalf("expected OwnerId = owner.Id (%d), got %d", owner.Id, reloaded.OwnerId)
	}
}

func TestEntity1EntityUpdateFn_PartialScalarLeavesOtherFieldsAlone(t *testing.T) {
	db := openPostgresTestDB(t)

	dto := &Entity1Entity{Title: "original", IsActive: true, ViewCount: 5}
	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	var input Entity1OptionalDto
	input.Title.Set(emigo.Ptr("updated"))

	updated, err := Entity1EntityActions.Update(db, created.UniqueId, input)
	if err != nil {
		t.Fatalf("Update error: %v", err)
	}

	if updated.Title != "updated" {
		t.Fatalf("Title = %q, want %q", updated.Title, "updated")
	}
	if !updated.IsActive {
		t.Fatal("expected IsActive to remain true (untouched field)")
	}
	if updated.ViewCount != 5 {
		t.Fatalf("expected ViewCount to remain 5 (untouched field), got %d", updated.ViewCount)
	}
}

func TestEntity1EntityUpdateFn_ArrayReplaceAppliesContentAndDeletesOrphans(t *testing.T) {
	db := openPostgresTestDB(t)

	dto := &Entity1Entity{
		Title: "hello",
		Items: []*Entity1EntityItems{
			{UniqueId: "c1", Item2: "child-1"},
			{UniqueId: "c2", Item2: "child-2"},
		},
	}
	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	var input Entity1OptionalDto
	input.Items.Set("replace", []Entity1OptionalDtoItems{
		{UniqueId: emigo.NullableOf("c2"), Item2: "child-2-updated"},
		{UniqueId: emigo.NullableOf("c3"), Item2: "child-3"},
	})

	if _, err := Entity1EntityActions.Update(db, created.UniqueId, input); err != nil {
		t.Fatalf("Update error: %v", err)
	}

	var children []Entity1EntityItems
	if err := db.Where("linker_id = ?", created.Id).Order("unique_id").Find(&children).Error; err != nil {
		t.Fatalf("query children error: %v", err)
	}
	if len(children) != 2 {
		t.Fatalf("expected 2 children after replace, got %d: %+v", len(children), children)
	}
	if children[0].UniqueId != "c2" || children[0].Item2 != "child-2-updated" {
		t.Fatalf("expected c2 to be updated in place, got %+v", children[0])
	}
	if children[1].UniqueId != "c3" || children[1].Item2 != "child-3" {
		t.Fatalf("expected c3 to be newly created, got %+v", children[1])
	}

	// c1 must be actually gone, not just orphaned (this is the exact gap
	// gorm's own Association().Replace() has - verified unreliable, see reconcile.go).
	var c1Check []Entity1EntityItems
	db.Unscoped().Where("unique_id = ?", "c1").Find(&c1Check)
	if len(c1Check) != 0 {
		t.Fatalf("expected c1 to be deleted, found %d rows", len(c1Check))
	}
}

func TestEntity1EntityUpdateFn_CollectionAppendKeepsExistingAndAddsNew(t *testing.T) {
	db := openPostgresTestDB(t)

	tagA := Entity2Entity{UniqueId: "tag-a", Label2: "tag-a"}
	tagB := Entity2Entity{UniqueId: "tag-b", Label2: "tag-b"}
	if err := db.Create(&tagA).Error; err != nil {
		t.Fatalf("seed tagA error: %v", err)
	}
	if err := db.Create(&tagB).Error; err != nil {
		t.Fatalf("seed tagB error: %v", err)
	}

	dto := &Entity1Entity{Title: "hello", Items3: emigo.CollectionReplace([]Entity2Entity{tagA})}
	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	// Entity1OptionalDto.Items3 targets Entity2Dto (the portable dto), not
	// Entity2Entity - and Entity1EntityUpdateFn only supports referencing an existing
	// row by its uniqueId (see go-entity-actions.go's walkUpdateFields), so tagB is
	// referenced by UniqueId alone; its other field values are looked up fresh rather
	// than trusted from the payload.
	var input Entity1OptionalDto
	input.Items3.Set("append", []Entity2Dto{{UniqueId: emigo.NullableOf(tagB.UniqueId)}})
	if _, err := Entity1EntityActions.Update(db, created.UniqueId, input); err != nil {
		t.Fatalf("Update error: %v", err)
	}

	var reloaded Entity1Entity
	if err := db.Preload("Items3Row").First(&reloaded, "unique_id = ?", created.UniqueId).Error; err != nil {
		t.Fatalf("reload error: %v", err)
	}
	if len(reloaded.Items3Row) != 2 {
		t.Fatalf("expected 2 tags after append, got %d", len(reloaded.Items3Row))
	}

	// tagA (the pre-existing, shared target) must still exist on its own - append/
	// replace on a many2many must never delete the target entity itself.
	var tagACheck Entity2Entity
	if err := db.First(&tagACheck, "unique_id = ?", "tag-a").Error; err != nil {
		t.Fatalf("expected tag-a to still exist: %v", err)
	}
}

func TestEntity1EntityUpdateFn_OneChangesOwnerViaSelect(t *testing.T) {
	db := openPostgresTestDB(t)

	owner1 := Entity2Entity{UniqueId: "owner-1", Label2: "owner-one"}
	owner2 := Entity2Entity{UniqueId: "owner-2", Label2: "owner-two"}
	if err := db.Create(&owner1).Error; err != nil {
		t.Fatalf("seed owner1 error: %v", err)
	}
	if err := db.Create(&owner2).Error; err != nil {
		t.Fatalf("seed owner2 error: %v", err)
	}

	dto := &Entity1Entity{Title: "hello", Owner: &owner1}
	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	var input Entity1OptionalDto
	input.Owner.SetSelector("owner-2")
	updated, err := Entity1EntityActions.Update(db, created.UniqueId, input)
	if err != nil {
		t.Fatalf("Update error: %v", err)
	}
	if updated.OwnerId != owner2.Id {
		t.Fatalf("expected OwnerId = owner2.Id (%d), got %d", owner2.Id, updated.OwnerId)
	}
}

// Exercises relations (array/one) nested 2 objects deep - the scenario that used to be
// silently dropped by Create/Update entirely (see nested_relations_test.go for the
// structural half of this verification, which doesn't need a live database).
func TestEntity1EntityCreateAndUpdateFn_NestedRelationsInsideObjectContainers(t *testing.T) {
	db := openPostgresTestDB(t)

	owner := Entity2Entity{UniqueId: "nested-owner-1", Label2: "nested-owner"}
	if err := db.Create(&owner).Error; err != nil {
		t.Fatalf("seed owner error: %v", err)
	}

	dto := &Entity1Entity{
		Title: "hello",
		NestedContainer: Entity1EntityNestedContainer{
			NestedInner: Entity1EntityNestedContainerNestedInner{
				NestedItems: []*Entity1EntityNestedContainerNestedInnerNestedItems{
					{Label: "nested-child-1"},
					{Label: "nested-child-2"},
				},
				NestedOwner: &owner,
			},
		},
	}

	created, err := Entity1EntityActions.Create(db, dto)
	if err != nil {
		t.Fatalf("Create error: %v", err)
	}

	var reloaded Entity1Entity
	if err := db.Preload("NestedContainer.NestedInner.NestedItems").
		Preload("NestedContainer.NestedInner.NestedOwner").
		First(&reloaded, "unique_id = ?", created.UniqueId).Error; err != nil {
		t.Fatalf("reload error: %v", err)
	}

	nested := reloaded.NestedContainer.NestedInner
	if len(nested.NestedItems) != 2 {
		t.Fatalf("expected 2 nested array children, got %d", len(nested.NestedItems))
	}
	if nested.NestedOwner == nil || nested.NestedOwner.Label2 != "nested-owner" {
		t.Fatalf("expected nested owner to resolve to nested-owner, got %+v", nested.NestedOwner)
	}

	// Update: replace the nested array (drop one, modify one, add one) - the same
	// reconcile guarantees as a top-level array, just reached through two objects.
	var input Entity1OptionalDto
	input.NestedContainer.Set(&Entity1OptionalDtoNestedContainer{})
	nc, _ := input.NestedContainer.Get()
	nc.NestedInner.Set(&Entity1OptionalDtoNestedContainerNestedInner{})
	ni, _ := nc.NestedInner.Get()
	ni.NestedItems.Set("replace", []Entity1OptionalDtoNestedContainerNestedInnerNestedItems{
		{UniqueId: emigo.NullableOf(nested.NestedItems[1].UniqueId), Label: "nested-child-2-updated"},
		{Label: "nested-child-3"},
	})

	if _, err := Entity1EntityActions.Update(db, created.UniqueId, input); err != nil {
		t.Fatalf("Update error: %v", err)
	}

	var afterUpdate Entity1Entity
	if err := db.Preload("NestedContainer.NestedInner.NestedItems").
		First(&afterUpdate, "unique_id = ?", created.UniqueId).Error; err != nil {
		t.Fatalf("reload error: %v", err)
	}
	items := afterUpdate.NestedContainer.NestedInner.NestedItems
	if len(items) != 2 {
		t.Fatalf("expected 2 nested children after replace, got %d", len(items))
	}
}
