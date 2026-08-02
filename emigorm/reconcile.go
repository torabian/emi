package emigorm

import (
	"gorm.io/gorm"
)

// resolveExistingId looks up item's real primary key by matching its UniqueId (the only
// identity a caller ever actually has) against an existing row, and stamps it onto item
// if found. gorm's Save() decides insert-vs-update purely from whether the primary key
// (Id) is already populated - a caller never knows or sends back the internal Id
// itself, so without this step Save() would always insert a duplicate row instead of
// updating the one the caller meant to patch.
func resolveExistingId[T any](tx *gorm.DB, item *T) error {
	uid := uniqueIdOf(item)
	if uid == "" {
		return nil
	}

	var existing T
	err := tx.Where("unique_id = ?", uid).First(&existing).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	setId(item, idOf(&existing))
	return nil
}

// ReconcileHasMany applies an emigo.Array/ArrayNullable "array" field against a
// has-many, owned-composition association (LinkerId-linked child rows - i.e. the
// {Field}Row []*Child sibling ApplyEntityGormTags generates alongside the DTO field).
//
// gorm's own Association().Replace() was verified unreliable for this: it silently
// keeps a colliding-primary-key item's *old* content instead of applying the new
// values, and it only clears (never deletes) rows dropped from the list, leaving
// orphans behind forever. This does the reliable thing instead:
//
//   - "replace" (the default, anything other than "append"): every existing child row
//     linked to linkerValue is loaded; any whose UniqueId isn't present in items gets
//     hard-deleted; every item in items is then upserted (Save) with linkerColumnName
//     set to linkerValue, so items that already exist (matched by UniqueId - see
//     resolveExistingId) get their content overwritten and brand new ones get created.
//   - "append": items are upserted the same way, without touching pre-existing rows.
//
// linkerValue is the parent's *id* (its real primary key, not UniqueId - see
// lib/golang/go-entity-default-fields.go for why). linkerColumnDbName is the linker
// column's snake_case DB name (e.g. "linker_id"), used in a raw Where clause - the
// column's Go field name is always "LinkerId" (ApplyEntityGormTags never lets this be
// customized), so that's what gets reflected into on each item.
func ReconcileHasMany[T any](tx *gorm.DB, linkerColumnDbName string, linkerValue int64, operation string, items []*T) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		if operation != "append" {
			var existing []*T
			if err := tx.Where(linkerColumnDbName+" = ?", linkerValue).Find(&existing).Error; err != nil {
				return err
			}

			keep := make(map[string]bool, len(items))
			for _, item := range items {
				if id := uniqueIdOf(item); id != "" {
					keep[id] = true
				}
			}

			for _, row := range existing {
				id := uniqueIdOf(row)
				if id != "" && !keep[id] {
					if err := tx.Unscoped().Delete(row).Error; err != nil {
						return err
					}
				}
			}
		}

		for _, item := range items {
			if err := resolveExistingId(tx, item); err != nil {
				return err
			}
			setInt64Field(item, "LinkerId", linkerValue)
			if err := tx.Save(item).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// ReconcileManyToMany applies an emigo.Collection/CollectionNullable "collection" field
// against a many2many association ({Field}Row []*Target). Unlike has-many, gorm's own
// Association API is reliable here (verified: replacing/appending only ever touches the
// join table, the shared target rows themselves are left alone) - this just upserts any
// inline target values first (so a caller-provided literal, not just an id reference,
// ends up with a real row to associate - matched against an existing row by UniqueId
// via resolveExistingId, same as ReconcileHasMany, so it updates rather than
// duplicates), then delegates to Association().Replace/Append.
//
// model must be a pointer to the owning entity, already loaded with its real Id
// populated (e.g. via a First() lookup - a bare &Entity1Entity{UniqueId: id} will NOT
// work, since gorm's Association API keys off the model's actual primary key), and
// assocName is the Go field name of the {Field}Row sibling (e.g. "Items3Row").
func ReconcileManyToMany[T any](tx *gorm.DB, model any, assocName string, operation string, items []*T) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := resolveExistingId(tx, item); err != nil {
				return err
			}
			if err := tx.Save(item).Error; err != nil {
				return err
			}
		}

		assoc := tx.Model(model).Association(assocName)
		if operation == "append" {
			return assoc.Append(items)
		}
		return assoc.Replace(items)
	})
}

// ReconcileOne resolves the target row for an emigo.One/OneNullable "one" field and
// returns its id, ready to be written into the {field}Id FK column on the parent.
//
//   - operation == "select": selectorId (already extracted by the generated code from
//     the DTO's Selector - its shape is caller/domain-defined, so this package can't
//     interpret it generically - and expected to be the target's UniqueId) is resolved
//     to the target's real id and returned; the target row itself is left untouched.
//   - anything else (an inline value was provided): the target is upserted (Save) as
//     given - matched against an existing row by UniqueId via resolveExistingId, same
//     as the other two reconcile functions, or freshly created if it doesn't match one
//     - and its id is returned.
//
// item may be nil (nothing to do, returns 0), matching a DTO field that was IsSet() but
// had no meaningful payload either way.
func ReconcileOne[T any](tx *gorm.DB, operation string, selectorId string, item *T) (int64, error) {
	if operation == "select" {
		if selectorId == "" {
			return 0, nil
		}
		var existing T
		if err := tx.First(&existing, "unique_id = ?", selectorId).Error; err != nil {
			return 0, err
		}
		return idOf(&existing), nil
	}

	if item == nil {
		return 0, nil
	}

	if err := resolveExistingId(tx, item); err != nil {
		return 0, err
	}
	if err := tx.Save(item).Error; err != nil {
		return 0, err
	}
	return idOf(item), nil
}
