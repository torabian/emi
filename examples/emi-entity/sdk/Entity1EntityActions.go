package external

import (
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// Entity1EntityCreateFn creates a new Entity1Entity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func Entity1EntityCreateFn(tx *gorm.DB, dto *Entity1Entity) (*Entity1Entity, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if dto.Owner.IsSet() {
			var selectorId string
			if dto.Owner.Operation == "select" {
				if s, ok := dto.Owner.Selector.(string); ok {
					selectorId = s
				}
			}
			var item *Entity2Entity
			if dto.Owner.Operation != "select" {
				item = &dto.Owner.Item
			}
			resolvedId, err := emigorm.ReconcileOne(tx, dto.Owner.Operation, selectorId, item)
			if err != nil {
				return err
			}
			dto.OwnerId = resolvedId
		}
		if dto.Manager.IsSet() {
			var selectorId string
			if dto.Manager.Operation == "select" {
				if s, ok := dto.Manager.Selector.(string); ok {
					selectorId = s
				}
			}
			var item *Entity2Entity
			if dto.Manager.Operation != "select" {
				item = &dto.Manager.Item
			}
			resolvedId, err := emigorm.ReconcileOne(tx, dto.Manager.Operation, selectorId, item)
			if err != nil {
				return err
			}
			dto.ManagerId = resolvedId
		}
		if v, ok := dto.Content2.Get(); ok && v != nil {
			dto.Content2Row = v
		}
		if dto.NestedContainer.NestedInner.NestedOwner.IsSet() {
			var selectorId string
			if dto.NestedContainer.NestedInner.NestedOwner.Operation == "select" {
				if s, ok := dto.NestedContainer.NestedInner.NestedOwner.Selector.(string); ok {
					selectorId = s
				}
			}
			var item *Entity2Entity
			if dto.NestedContainer.NestedInner.NestedOwner.Operation != "select" {
				item = &dto.NestedContainer.NestedInner.NestedOwner.Item
			}
			resolvedId, err := emigorm.ReconcileOne(tx, dto.NestedContainer.NestedInner.NestedOwner.Operation, selectorId, item)
			if err != nil {
				return err
			}
			dto.NestedContainer.NestedInner.NestedOwnerId = resolvedId
		}
		if v, ok := dto.NestedContainerOpt.Get(); ok && v != nil {
			dto.NestedContainerOptRow = v
		}
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
		if dto.Items.IsSet() {
			items := make([]*Entity1EntityItems, len(dto.Items.Items))
			for i := range dto.Items.Items {
				items[i] = &dto.Items.Items[i]
			}
			if err := emigorm.ReconcileHasMany(tx, "linker_id", dto.Id, dto.Items.Operation, items); err != nil {
				return err
			}
		}
		if dto.Items2.IsSet() {
			items := make([]*Entity1EntityItems2, len(dto.Items2.Items))
			for i := range dto.Items2.Items {
				items[i] = &dto.Items2.Items[i]
			}
			if err := emigorm.ReconcileHasMany(tx, "linker_id", dto.Id, dto.Items2.Operation, items); err != nil {
				return err
			}
		}
		if dto.Items3.IsSet() {
			items := make([]*Entity2Entity, len(dto.Items3.Items))
			for i := range dto.Items3.Items {
				items[i] = &dto.Items3.Items[i]
			}
			if err := emigorm.ReconcileManyToMany(tx, dto, "Items3Row", dto.Items3.Operation, items); err != nil {
				return err
			}
		}
		if dto.Items4.IsSet() {
			items := make([]*Entity2Entity, len(dto.Items4.Items))
			for i := range dto.Items4.Items {
				items[i] = &dto.Items4.Items[i]
			}
			if err := emigorm.ReconcileManyToMany(tx, dto, "Items4Row", dto.Items4.Operation, items); err != nil {
				return err
			}
		}
		if dto.NestedContainer.NestedInner.NestedItems.IsSet() {
			items := make([]*Entity1EntityNestedContainerNestedInnerNestedItems, len(dto.NestedContainer.NestedInner.NestedItems.Items))
			for i := range dto.NestedContainer.NestedInner.NestedItems.Items {
				items[i] = &dto.NestedContainer.NestedInner.NestedItems.Items[i]
			}
			if err := emigorm.ReconcileHasMany(tx, "linker_id", dto.Id, dto.NestedContainer.NestedInner.NestedItems.Operation, items); err != nil {
				return err
			}
		}
		if v, ok := dto.NestedContainerOpt.Get(); ok && v != nil {
			if v.NestedInner.NestedItemsOpt.IsSet() {
				items := make([]*Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt, len(v.NestedInner.NestedItemsOpt.Items))
				for i := range v.NestedInner.NestedItemsOpt.Items {
					items[i] = &v.NestedInner.NestedItemsOpt.Items[i]
				}
				if err := emigorm.ReconcileHasMany(tx, "linker_id", dto.Id, v.NestedInner.NestedItemsOpt.Operation, items); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// Entity1EntityUpdateFn applies a partial update to the Entity1Entity row identified by id (its public
// uniqueId, e.g. from an API path parameter - never the internal auto-increment id).
// Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// Entity1EntityCreateFn uses, against entity.Id (the row's real primary key, resolved from id
// up front - gorm's Association API and the has-many reconcile both join on it, not on
// uniqueId).
func Entity1EntityUpdateFn(tx *gorm.DB, id string, input Entity1EntityUpdateInput) (*Entity1Entity, error) {
	var entity Entity1Entity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Owner.IsSet() {
			var selectorId string
			if input.Owner.Operation == "select" {
				if s, ok := input.Owner.Selector.(string); ok {
					selectorId = s
				}
			}
			var item *Entity2Entity
			if input.Owner.Operation != "select" {
				item = &input.Owner.Item
			}
			resolvedId, err := emigorm.ReconcileOne(tx, input.Owner.Operation, selectorId, item)
			if err != nil {
				return err
			}
			changes["OwnerId"] = resolvedId
		}
		if input.Manager.IsSet() {
			var selectorId string
			if input.Manager.Operation == "select" {
				if s, ok := input.Manager.Selector.(string); ok {
					selectorId = s
				}
			}
			var item *Entity2Entity
			if input.Manager.Operation != "select" {
				item = &input.Manager.Item
			}
			resolvedId, err := emigorm.ReconcileOne(tx, input.Manager.Operation, selectorId, item)
			if err != nil {
				return err
			}
			changes["ManagerId"] = resolvedId
		}
		if input.NestedContainer.IsSet() {
			if v, ok := input.NestedContainer.Get(); ok && v != nil {
				if v.NestedInner.IsSet() {
					if v, ok := v.NestedInner.Get(); ok && v != nil {
						if v.NestedOwner.IsSet() {
							var selectorId string
							if v.NestedOwner.Operation == "select" {
								if s, ok := v.NestedOwner.Selector.(string); ok {
									selectorId = s
								}
							}
							var item *Entity2Entity
							if v.NestedOwner.Operation != "select" {
								item = &v.NestedOwner.Item
							}
							resolvedId, err := emigorm.ReconcileOne(tx, v.NestedOwner.Operation, selectorId, item)
							if err != nil {
								return err
							}
							changes["NestedOwnerId"] = resolvedId
						}
					}
				}
			}
		}
		if input.Title.IsSet() {
			changes["Title"] = input.Title
		}
		if input.Content1.IsSet() {
			if v, ok := input.Content1.Get(); ok && v != nil {
				if v.Item1.IsSet() {
					changes["Item1"] = v.Item1
				}
			}
		}
		if input.Content2.IsSet() {
			if v, ok := input.Content2.Get(); ok && v != nil {
				if v.Item2.IsSet() {
					changes["Item2"] = v.Item2
				}
			}
		}
		if input.Complex1 != nil {
			changes["Complex1"] = input.Complex1
		}
		if input.Subtitle.IsSet() {
			changes["Subtitle"] = input.Subtitle
		}
		if input.IsActive.IsSet() {
			changes["IsActive"] = input.IsActive
		}
		if input.IsFeatured.IsSet() {
			changes["IsFeatured"] = input.IsFeatured
		}
		if input.ViewCount.IsSet() {
			changes["ViewCount"] = input.ViewCount
		}
		if input.ViewCountOpt.IsSet() {
			changes["ViewCountOpt"] = input.ViewCountOpt
		}
		if input.SmallCount.IsSet() {
			changes["SmallCount"] = input.SmallCount
		}
		if input.SmallCountOpt.IsSet() {
			changes["SmallCountOpt"] = input.SmallCountOpt
		}
		if input.BigCount.IsSet() {
			changes["BigCount"] = input.BigCount
		}
		if input.BigCountOpt.IsSet() {
			changes["BigCountOpt"] = input.BigCountOpt
		}
		if input.Ratio32.IsSet() {
			changes["Ratio32"] = input.Ratio32
		}
		if input.Ratio32Opt.IsSet() {
			changes["Ratio32Opt"] = input.Ratio32Opt
		}
		if input.Ratio64.IsSet() {
			changes["Ratio64"] = input.Ratio64
		}
		if input.Ratio64Opt.IsSet() {
			changes["Ratio64Opt"] = input.Ratio64Opt
		}
		if input.Status.IsSet() {
			changes["Status"] = input.Status
		}
		if input.StatusOpt.IsSet() {
			changes["StatusOpt"] = input.StatusOpt
		}
		if input.Metadata.IsSet() {
			changes["Metadata"] = input.Metadata
		}
		if input.MetadataOpt.IsSet() {
			changes["MetadataOpt"] = input.MetadataOpt
		}
		if input.RawSettings.IsSet() {
			changes["RawSettings"] = input.RawSettings
		}
		if input.Labels.IsSet() {
			changes["Labels"] = input.Labels
		}
		if input.LabelsOpt.IsSet() {
			changes["LabelsOpt"] = input.LabelsOpt
		}
		if input.Misc != nil {
			changes["Misc"] = input.Misc
		}
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
		if input.Items.IsSet() {
			items := make([]*Entity1EntityItems, len(input.Items.Items))
			for i := range input.Items.Items {
				items[i] = &input.Items.Items[i]
			}
			if err := emigorm.ReconcileHasMany(tx, "linker_id", entity.Id, input.Items.Operation, items); err != nil {
				return err
			}
		}
		if input.Items2.IsSet() {
			items := make([]*Entity1EntityItems2, len(input.Items2.Items))
			for i := range input.Items2.Items {
				items[i] = &input.Items2.Items[i]
			}
			if err := emigorm.ReconcileHasMany(tx, "linker_id", entity.Id, input.Items2.Operation, items); err != nil {
				return err
			}
		}
		if input.Items3.IsSet() {
			items := make([]*Entity2Entity, len(input.Items3.Items))
			for i := range input.Items3.Items {
				items[i] = &input.Items3.Items[i]
			}
			if err := emigorm.ReconcileManyToMany(tx, &entity, "Items3Row", input.Items3.Operation, items); err != nil {
				return err
			}
		}
		if input.Items4.IsSet() {
			items := make([]*Entity2Entity, len(input.Items4.Items))
			for i := range input.Items4.Items {
				items[i] = &input.Items4.Items[i]
			}
			if err := emigorm.ReconcileManyToMany(tx, &entity, "Items4Row", input.Items4.Operation, items); err != nil {
				return err
			}
		}
		if input.NestedContainer.IsSet() {
			if v, ok := input.NestedContainer.Get(); ok && v != nil {
				if v.NestedInner.IsSet() {
					if v, ok := v.NestedInner.Get(); ok && v != nil {
						if v.NestedItems.IsSet() {
							items := make([]*Entity1EntityNestedContainerNestedInnerNestedItems, len(v.NestedItems.Items))
							for i := range v.NestedItems.Items {
								items[i] = &v.NestedItems.Items[i]
							}
							if err := emigorm.ReconcileHasMany(tx, "linker_id", entity.Id, v.NestedItems.Operation, items); err != nil {
								return err
							}
						}
					}
				}
			}
		}
		if input.NestedContainerOpt.IsSet() {
			if v, ok := input.NestedContainerOpt.Get(); ok && v != nil {
				if v.NestedInner.IsSet() {
					if v, ok := v.NestedInner.Get(); ok && v != nil {
						if v.NestedItemsOpt.IsSet() {
							items := make([]*Entity1EntityNestedContainerOptNestedInnerNestedItemsOpt, len(v.NestedItemsOpt.Items))
							for i := range v.NestedItemsOpt.Items {
								items[i] = &v.NestedItemsOpt.Items[i]
							}
							if err := emigorm.ReconcileHasMany(tx, "linker_id", entity.Id, v.NestedItemsOpt.Operation, items); err != nil {
								return err
							}
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var updated Entity1Entity
	if err := tx.First(&updated, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// Entity1EntityActionsSig bundles the actions available for Entity1Entity. Extend this (and
// Entity1EntityActions below) with more fields as more actions are generated - Create/Update
// are wired to Entity1EntityCreateFn/Entity1EntityUpdateFn by default, but callers can swap either out
// (e.g. in tests, or to layer extra validation/side effects around them).
type Entity1EntityActionsSig struct {
	Create func(tx *gorm.DB, dto *Entity1Entity) (*Entity1Entity, error)
	Update func(tx *gorm.DB, id string, input Entity1EntityUpdateInput) (*Entity1Entity, error)
}

var Entity1EntityActions Entity1EntityActionsSig = Entity1EntityActionsSig{
	Create: Entity1EntityCreateFn,
	Update: Entity1EntityUpdateFn,
}
