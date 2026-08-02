package golang

import (
	"fmt"
	"strings"

	"github.com/torabian/emi/lib/core"
)

// Struct-tag fragments used below - kept as their own constants (rather than typed
// directly into the fmt.Sprintf templates further down) because those templates are Go
// raw string literals, which can't contain a literal backtick themselves; a plain
// double-quoted Go string has no such restriction.
const (
	relationJsonTag = "`json:\"relation\"`"
	countJsonTag    = "`json:\"count\"`"
	messageJsonTag  = "`json:\"message\"`"
	affectedJsonTag = "`json:\"affected\"`"
)

// deleteRelation describes one array/collection relation AwareDelete needs to fold into
// its preview/execute pair - the same has-many (LinkerId) / many-to-many (Association)
// shapes Create/Update already reconcile (see walkCreateFields/walkUpdateFields in
// go-entity-actions.go), just walked purely for awareness/cleanup rather than upsert.
//
// one/one? relations are deliberately excluded: they're a plain FK column on the entity
// itself (see ApplyEntityGormTags), so deleting the entity doesn't cascade into them -
// there's nothing to report or clean up on that side.
type deleteRelation struct {
	kind         string // "array" or "collection"
	relationName string // dotted path for the report, e.g. "nestedContainer.nestedInner.nestedItems"
	childStruct  string // array only: the child row struct name (structPrefix+goName)
	assocName    string // collection only: the {Field}Row sibling's Go field name
}

// walkDeleteRelations recurses through fields - including into object/object?
// containers, at any depth - collecting every array/array?/collection/collection?
// relation. structPrefix mirrors walkCreateFields/walkUpdateFields's own naming (the
// entity's class name at the top level, {parentPrefix}{ParentFieldName} once recursed),
// so childStruct always matches GoCommonStructGenerator's actual generated type name.
func walkDeleteRelations(fields []*core.EmiField, namePrefix, structPrefix string) []deleteRelation {
	var out []deleteRelation
	for _, field := range fields {
		if field == nil {
			continue
		}
		goName := core.ToUpper(field.Name)
		dotted := namePrefix + field.Name

		switch field.Type {
		case core.FieldTypeArray, core.FieldTypeArrayNullable:
			out = append(out, deleteRelation{kind: "array", relationName: dotted, childStruct: structPrefix + goName})

		case core.FieldTypeCollection, core.FieldTypeCollectionNullable:
			out = append(out, deleteRelation{kind: "collection", relationName: dotted, assocName: entityRowFieldName(field)})

		case core.FieldTypeObject, core.FieldTypeObjectNullable:
			out = append(out, walkDeleteRelations(field.Fields, dotted+".", structPrefix+goName)...)
		}
	}
	return out
}

// buildAwareDeleteFns renders {Entity}AwareDeletePreviewFn/{Entity}AwareDeleteFn (plus
// the small types they share) - see the doc comments embedded in the generated code
// itself for their exact contract. Relations are gathered once via walkDeleteRelations
// and used to build both functions' bodies in lockstep, so preview and execute can never
// silently drift apart (a relation added to one is always added to the other).
func buildAwareDeleteFns(className string, fields []*core.EmiField) string {
	relations := walkDeleteRelations(fields, "", className)

	var previewBody, deleteBody strings.Builder
	for i, rel := range relations {
		switch rel.kind {
		case "array":
			fmt.Fprintf(&previewBody, `
	var affected%[1]d int64
	tx.Model(&%[2]s{}).Where("linker_id IN ?", ids).Count(&affected%[1]d)
	if affected%[1]d > 0 {
		affected = append(affected, %[3]sAwareDeleteAffected{Relation: %[4]q, Count: affected%[1]d})
		total += affected%[1]d
	}
`, i, rel.childStruct, className, rel.relationName)

			fmt.Fprintf(&deleteBody, `
	if err := tx.Where("linker_id IN ?", ids).Delete(&%[1]s{}).Error; err != nil {
		return err
	}
`, rel.childStruct)

		case "collection":
			fmt.Fprintf(&previewBody, `
	var affected%[1]d int64
	for i := range rows {
		affected%[1]d += tx.Model(rows[i]).Association(%[2]q).Count()
	}
	if affected%[1]d > 0 {
		affected = append(affected, %[3]sAwareDeleteAffected{Relation: %[4]q, Count: affected%[1]d})
		total += affected%[1]d
	}
`, i, rel.assocName, className, rel.relationName)

			fmt.Fprintf(&deleteBody, `
	for i := range rows {
		if err := tx.Model(rows[i]).Association(%[1]q).Clear(); err != nil {
			return err
		}
	}
`, rel.assocName)
		}
	}

	return fmt.Sprintf(`
// %[1]sAwareDeleteAffected reports one relation of %[1]s that would be affected by
// deleting the matching row(s) - either its has-many child rows are hard-deleted
// (array/array?) or its many-to-many join rows are cleared, leaving the target rows
// themselves untouched (collection/collection?). one/one? relations are never listed:
// they're a plain FK column on %[1]s itself, so deleting %[1]s doesn't cascade into them.
type %[1]sAwareDeleteAffected struct {
	Relation string %[2]s
	Count    int64  %[3]s
}

// %[1]sAwareDeletePreview is the result of %[1]sAwareDeletePreviewFn: a human-readable
// summary plus the exact per-relation counts %[1]sAwareDeleteFn would delete/clear
// alongside the %[1]s row(s) themselves.
type %[1]sAwareDeletePreview struct {
	Message  string %[4]s
	Affected []%[1]sAwareDeleteAffected %[5]s
}

// %[1]sAwareDeletePreviewFn looks up the %[1]s rows matching uniqueIds and reports what
// deleting them would affect - every array/array?/collection/collection? relation (at
// any nesting depth inside object/object? containers), matching exactly what
// %[1]sAwareDeleteFn deletes/clears. Intended as a confirmation step before actually
// calling %[1]sAwareDeleteFn.
func %[1]sAwareDeletePreviewFn(tx *gorm.DB, uniqueIds []string) (*%[1]sAwareDeletePreview, error) {
	var rows []*%[1]s
	if err := tx.Where("unique_id IN ?", uniqueIds).Find(&rows).Error; err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &%[1]sAwareDeletePreview{Message: "No matching %[1]s row was found for the given uniqueIds."}, nil
	}

	ids := make([]int64, len(rows))
	for i := range rows {
		ids[i] = rows[i].Id
	}

	affected := []%[1]sAwareDeleteAffected{}
	var total int64
%[6]s
	message := fmt.Sprintf("Deleting %%d %[1]s row(s) will affect %%d related record(s) across %%d relation(s).", len(rows), total, len(affected))
	return &%[1]sAwareDeletePreview{Message: message, Affected: affected}, nil
}

// %[1]sAwareDeleteFn deletes the %[1]s rows matching uniqueIds, along with every
// array/array?/collection/collection? relation %[1]sAwareDeletePreviewFn reports (see
// its own doc comment for exactly what that means per relation kind).
func %[1]sAwareDeleteFn(tx *gorm.DB, uniqueIds []string) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		var rows []*%[1]s
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
%[7]s
		return tx.Where("id IN ?", ids).Delete(&%[1]s{}).Error
	})
}
`, className, relationJsonTag, countJsonTag, messageJsonTag, affectedJsonTag, previewBody.String(), deleteBody.String())
}
