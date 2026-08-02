package external

import (
	"encoding/json"
	"fmt"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/emigorm"
	"gorm.io/gorm"
)

// The base class definition for entity3Entity
type Entity3Entity struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"-" yaml:"-"`
	UniqueId string `gorm:"type:uuid;default:gen_random_uuid();unique" json:"uniqueId" yaml:"uniqueId"`
	Message  string `json:"message" yaml:"message"`
}

func (x *Entity3Entity) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
func GetEntity3EntityCliFlags(prefix string) []emigo.CliFlag {
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
			Name: prefix + "message",
			Type: "string",
		},
	}
}
func CastEntity3EntityFromCli(c emigo.CliCastable) Entity3Entity {
	data := Entity3Entity{}
	if c.IsSet("id") {
		data.Id = int64(c.Int64("id"))
	}
	if c.IsSet("unique-id") {
		data.UniqueId = c.String("unique-id")
	}
	if c.IsSet("message") {
		data.Message = c.String("message")
	}
	return data
}

// Extra entity-specific code (hooks, custom methods, business logic, etc.) can be
// appended here in this template, after the struct GoCommonStructGenerator produced.
// Entity3EntityCreateFn creates a new Entity3Entity row (and its array/collection/one relations,
// including ones nested inside object/object? fields) from dto. dto.Id/dto.UniqueId are
// assigned by the database (see AutoMigrate's column defaults) and populated back onto
// dto once created. Relations are applied in a single transaction: one/one? are
// resolved before the row itself is created (a belongs-to FK doesn't need the parent's
// own id); array/array? and collection/collection? are reconciled afterwards, once
// dto.Id is known.
func Entity3EntityCreateFn(tx *gorm.DB, dto *Entity3Entity) (*Entity3Entity, error) {
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

// Entity3EntityGetFn looks up a single Entity3Entity row by its public uniqueId (e.g. from an API path
// parameter - never the internal auto-increment id).
func Entity3EntityGetFn(tx *gorm.DB, id string) (*Entity3Entity, error) {
	var entity Entity3Entity
	if err := tx.First(&entity, "unique_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Entity3EntityBrowseFn returns Entity3Entity rows matching dsl.Filter (a JSON-logic expression) and
// dsl.Scope (a second, handler-enforced condition - e.g. workspace isolation - see
// emigorm.QueryDSL), sorted/paged per dsl.Sort/StartIndex/ItemsPerPage/Cursor, alongside
// a emigo.QueryResultMeta reporting the total row count matching both filters (ignoring
// paging) and a cursor for fetching the next page.
func Entity3EntityBrowseFn(tx *gorm.DB, dsl emigorm.QueryDSL) ([]*Entity3Entity, *emigo.QueryResultMeta, error) {
	filtered, err := emigorm.ApplyQueryFilter(tx.Model(&Entity3Entity{}), dsl)
	if err != nil {
		return nil, nil, err
	}
	filtered = emigorm.ApplyQueryScope(filtered, dsl)
	var total int64
	if err := filtered.Count(&total).Error; err != nil {
		return nil, nil, err
	}
	var items []*Entity3Entity
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

// Entity3EntityAwareDeleteAffected reports one relation of Entity3Entity that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on Entity3Entity itself, so deleting Entity3Entity doesn't cascade into them.
type Entity3EntityAwareDeleteAffected struct {
	Relation string `json:"relation"`
	Count    int64  `json:"count"`
}

// Entity3EntityAwareDeletePreview is the result of Entity3EntityAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts Entity3EntityAwareDeleteFn would delete/clear
// alongside the Entity3Entity row(s) themselves.
type Entity3EntityAwareDeletePreview struct {
	Message  string                             `json:"message"`
	Affected []Entity3EntityAwareDeleteAffected `json:"affected"`
}

// Entity3EntityAwareDeletePreviewFn looks up the Entity3Entity rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// Entity3EntityAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling Entity3EntityAwareDeleteFn.
func Entity3EntityAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*Entity3EntityAwareDeletePreview, error) {
	var rows []*Entity3Entity
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &Entity3EntityAwareDeletePreview{Message: "No matching Entity3Entity row was found for the given uniqueIds."}, nil
	}
	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}
	affected := []Entity3EntityAwareDeleteAffected{}
	var total int64
	message := fmt.Sprintf("Deleting %d Entity3Entity row(s) will affect %d related record(s) across %d relation(s).", len(rows), total, len(affected))
	return &Entity3EntityAwareDeletePreview{Message: message, Affected: affected}, nil
}

// Entity3EntityAwareDeleteFn deletes the Entity3Entity rows matching uniqueIds, along with every
// array/array?/collection/collection? relation Entity3EntityAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func Entity3EntityAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*Entity3Entity
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
		return tx.Where("id IN ?", ids).Delete(&Entity3Entity{}).Error
	})
}

// Entity3EntityActionsSig bundles the actions available for Entity3Entity. Extend this (and
// Entity3EntityActions below) with more fields as more actions are generated. Which fields are
// present here depends on entity.Features (see Module3EntityFeatures) - a disabled
// feature is omitted entirely rather than left as a nil func.
type Entity3EntityActionsSig struct {
	Create             func(tx *gorm.DB, dto *Entity3Entity) (*Entity3Entity, error)
	Get                func(tx *gorm.DB, id string) (*Entity3Entity, error)
	Browse             func(tx *gorm.DB, dsl emigorm.QueryDSL) ([]*Entity3Entity, *emigo.QueryResultMeta, error)
	AwareDeletePreview func(tx *gorm.DB, uniqueIds []string) (*Entity3EntityAwareDeletePreview, error)
	AwareDelete        func(tx *gorm.DB, uniqueIds []string) error
}

var Entity3EntityActions Entity3EntityActionsSig = Entity3EntityActionsSig{
	Create:             Entity3EntityCreateFn,
	Get:                Entity3EntityGetFn,
	Browse:             Entity3EntityBrowseFn,
	AwareDeletePreview: Entity3EntityAwareDeletePreviewFn,
	AwareDelete:        Entity3EntityAwareDeleteFn,
}
