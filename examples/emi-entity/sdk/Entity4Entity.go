package external

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for entity4Entity
type Entity4Entity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Note     string `json:"note" yaml:"note"`
}

func (x *Entity4Entity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity4EntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "note",
			Type: "string",
		},
	}
}
func CastEntity4EntityFromCli(c emigo.CliCastable) Entity4Entity {
	data := Entity4Entity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("note") {
		data.Note = c.String("note")
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// Entity4EntityUpdateFn applies a partial update to the Entity4Entity row identified by id (its public
// uniqueId, e.g. from an API path parameter - never the internal auto-increment id).
// Only fields the caller actually set on input (input.{Field}.IsSet()) are touched -
// anything else is left exactly as it was. one/one? are resolved into their {field}Id
// FK column alongside the rest of the scalar changes; array/array? and
// collection/collection? are reconciled afterwards via the same emigorm helpers
// Entity4EntityCreateFn uses, against entity.Id (the row's real primary key, resolved from id
// up front - gorm's Association API and the has-many reconcile both join on it, not on
// uniqueId).
func Entity4EntityUpdateFn(tx *gorm.DB, id string, input Entity4OptionalDto) (*Entity4Entity, error) {
	var entity Entity4Entity
	err := tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
			return err
		}
		changes := map[string]interface{}{}
		if input.Note.IsSet() {
			changes["Note"] = input.Note
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
	var updated Entity4Entity
	if err := tx.First(&updated, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

// Entity4EntityGetFn looks up a single Entity4Entity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func Entity4EntityGetFn(tx *gorm.DB, id string) (*Entity4Entity, error) {
	var entity Entity4Entity
	if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Entity4EntityBrowseFn returns Entity4Entity rows matching dsl.Filter (a JSON-logic expression) and
// dsl.Scope (a second, handler-enforced condition - e.g. workspace isolation - see
// emigorm.QueryDSL), sorted/paged per dsl.Sort/StartIndex/ItemsPerPage/Cursor, alongside
// a emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func Entity4EntityBrowseFn(tx *gorm.DB, dsl emigorm.QueryDSL) ([]*Entity4Entity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&Entity4Entity{}), dsl)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, dsl)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*Entity4Entity
	paged := emigorm.ApplyQueryPage(emigorm.ApplyQueryCursor(emigorm.ApplyQuerySort(filtered, dsl), dsl), dsl)
	if err := paged.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	meta := &emigo.QueryResultMeta{
		TotalItems: total,
		Cursor:     emigorm.BuildQueryCursor(items),
	}
	return items, meta, nil
}

// Entity4EntityAwareDeleteAffected reports one relation of Entity4Entity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on Entity4Entity itself, so deleting Entity4Entity doesn't cascade into them.
type Entity4EntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// Entity4EntityAwareDeletePreview is the result of Entity4EntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts Entity4EntityAwareDeleteFn would delete/clear
// alongside the Entity4Entity row(s) themselves.
type Entity4EntityAwareDeletePreview struct {
	Message  string                             `json:"message"`
	Affected []Entity4EntityAwareDeleteAffected `json:"affected"`
}

// Entity4EntityAwareDeletePreviewFn looks up the Entity4Entity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// Entity4EntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling Entity4EntityAwareDeleteFn.
func Entity4EntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*Entity4EntityAwareDeletePreview, error) {
	var rows []*Entity4Entity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &Entity4EntityAwareDeletePreview{Message: "No matching Entity4Entity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []Entity4EntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d Entity4Entity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &Entity4EntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// Entity4EntityAwareDeleteFn deletes the Entity4Entity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation Entity4EntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func Entity4EntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*Entity4Entity
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
		return tx.Where("id IN ?", ids).Delete(&Entity4Entity{}).Error
	})
}

// Entity4EntityActionsSig bundles the actions available for Entity4Entity. Extend this (and
// Entity4EntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type Entity4EntityActionsSig struct {
	Update             func(tx *gorm.DB, id string, input Entity4OptionalDto) (*Entity4Entity, error)
	Get                func(tx *gorm.DB, id string) (*Entity4Entity, error)
	Browse             func(tx *gorm.DB, dsl emigorm.QueryDSL) ([]*Entity4Entity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*Entity4EntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var Entity4EntityActions Entity4EntityActionsSig = Entity4EntityActionsSig{
	Update:             Entity4EntityUpdateFn,
	Get:                Entity4EntityGetFn,
	Browse:             Entity4EntityBrowseFn,
	AwareDeletePreview: Entity4EntityAwareDeletePreviewFn,
	AwareDelete:        Entity4EntityAwareDeleteFn,
}
