package external

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for entity2Entity
type Entity2Entity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Label2   string `json:"label2" yaml:"label2"`
}

func (x *Entity2Entity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity2EntityCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "id",
			Type: "int64",
		},
		{
			Name: prefix + "unique-id",
			Type: "string",
		},
		{
			Name: prefix + "label2",
			Type: "string",
		},
	}
}
func CastEntity2EntityFromCli(c emigo.CliCastable) Entity2Entity {
	data := Entity2Entity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("label2") {
		data.Label2 = c.String("label2")
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
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

// Entity2EntityUpdateFn applies a partial update to the Entity2Entity row identified by uniqueId (its
// public identity, e.g. from an API path parameter - never the internal auto-increment
// id). Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// Entity2EntityCreateFn uses, against entity.Id (the row's real primary key, resolved from
// uniqueId up front - gorm's Association API and the has-many reconcile both join on
// it, not on uniqueId).
func Entity2EntityUpdateFn(tx *gorm.DB, uniqueId string, input Entity2OptionalDto) (*Entity2Entity, error) {
	var entity Entity2Entity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Label2.IsSet() {
			changes["Label2"] = input.Label2
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
	if err := tx.First(&updated, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// Entity2EntityGetFn looks up a single Entity2Entity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func Entity2EntityGetFn(tx *gorm.DB, uniqueId string) (*Entity2Entity, error) {
	var entity Entity2Entity
	if err := tx.First(&entity, "unique_id = ?", uniqueId).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Entity2EntityBrowseFn returns Entity2Entity rows matching qs.Filter (a JSON-logic expression) and
// scope/scopeArgs (a second, handler-enforced condition - e.g. workspace isolation),
// sorted/paged per qs.Sort/StartIndex/ItemsPerPage/Cursor, alongside a
// emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func Entity2EntityBrowseFn(tx *gorm.DB, qs Entity2BrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*Entity2Entity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&Entity2Entity{}), qs.Filter)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, scope, scopeArgs...)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*Entity2Entity
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, qs.Sort), qs.Cursor), qs.StartIndex, qs.ItemsPerPage)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}
	return items, meta, nil
}

// Entity2EntityAwareDeleteAffected reports one relation of Entity2Entity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on Entity2Entity itself, so deleting Entity2Entity doesn't cascade into them.
type Entity2EntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// Entity2EntityAwareDeletePreview is the result of Entity2EntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts Entity2EntityAwareDeleteFn would delete/clear
// alongside the Entity2Entity row(s) themselves.
type Entity2EntityAwareDeletePreview struct {
	Message  string                             `json:"message"`
	Affected []Entity2EntityAwareDeleteAffected `json:"affected"`
}

// Entity2EntityAwareDeletePreviewFn looks up the Entity2Entity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// Entity2EntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling Entity2EntityAwareDeleteFn.
func Entity2EntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*Entity2EntityAwareDeletePreview, error) {
	var rows []*Entity2Entity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &Entity2EntityAwareDeletePreview{Message: "No matching Entity2Entity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []Entity2EntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d Entity2Entity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &Entity2EntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// Entity2EntityAwareDeleteFn deletes the Entity2Entity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation Entity2EntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func Entity2EntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*Entity2Entity
		if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
			return err
		}
		if len(rows) == 0 {
			return nil
		}
		ids := make([]int64, len(rows))
		for i := range rows {
			ids[i] = rows[i].Id
		}
		return tx.Where("id IN ?", ids).Delete(&Entity2Entity{}).Error
	})
}

// Entity2EntityActionsSig bundles the actions available for Entity2Entity. Extend this (and
// Entity2EntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type Entity2EntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *Entity2Entity) (*Entity2Entity, error)
	Update             func(tx *gorm.DB, uniqueId string, input Entity2OptionalDto) (*Entity2Entity, error)
	Get                func(tx *gorm.DB, uniqueId string) (*Entity2Entity, error)
	Browse             func(tx *gorm.DB, qs Entity2BrowseActionQuery, scope string, scopeArgs ...interface{}) ([]*Entity2Entity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*Entity2EntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var Entity2EntityActions Entity2EntityActionsSig = Entity2EntityActionsSig{
	Create:             Entity2EntityCreateFn,
	Update:             Entity2EntityUpdateFn,
	Get:                Entity2EntityGetFn,
	Browse:             Entity2EntityBrowseFn,
	AwareDeletePreview: Entity2EntityAwareDeletePreviewFn,
	AwareDelete:        Entity2EntityAwareDeleteFn,
}
