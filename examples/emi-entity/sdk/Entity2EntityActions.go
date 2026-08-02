package external

import "gorm.io/gorm"

// Entity2EntityCreateFn creates a new Entity2Entity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func Entity2EntityCreateFn(tx *gorm.DB, dto *Entity2Entity) (*Entity2Entity, error) {
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(dto).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dto, nil
}

// Entity2EntityUpdateFn applies a partial update to the Entity2Entity row identified by id (its public
// uniqueId, e.g. from an API path parameter - never the internal auto-increment id).
// Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// Entity2EntityCreateFn uses, against entity.Id (the row's real primary key, resolved from id
// up front - gorm's Association API and the has-many reconcile both join on it, not on
// uniqueId).
func Entity2EntityUpdateFn(tx *gorm.DB, id string, input Entity2EntityUpdateInput) (*Entity2Entity, error) {
	var entity Entity2Entity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Label.IsSet() {
			changes["Label"] = input.Label
		}
		if len(changes) > 0 {
			if err := tx.Model(&entity).Updates(changes).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var updated Entity2Entity
	if err := tx.First(&updated, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// Entity2EntityActionsSig bundles the actions available for Entity2Entity. Extend this (and
// Entity2EntityActions below) with more fields as more actions are generated - Create/Update
// are wired to Entity2EntityCreateFn/Entity2EntityUpdateFn by default, but callers can swap either out
// (e.g. in tests, or to layer extra validation/side effects around them).
type Entity2EntityActionsSig struct {
	Create func(tx *gorm.DB, dto *Entity2Entity) (*Entity2Entity, error)
	Update func(tx *gorm.DB, id string, input Entity2EntityUpdateInput) (*Entity2Entity, error)
}

var Entity2EntityActions Entity2EntityActionsSig = Entity2EntityActionsSig{
	Create: Entity2EntityCreateFn,
	Update: Entity2EntityUpdateFn,
}
